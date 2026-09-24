// The smoke client runs only against the disposable test Compose project.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"os"
	"strings"
	"time"
)

type harness struct {
	ctx              context.Context
	nakama, provider string
	http             *http.Client
}

func (h *harness) request(method, address string, body any, auth string) (int, []byte, error) {
	var input io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return 0, nil, err
		}
		input = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(h.ctx, method, address, input)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if auth == "basic" {
		req.SetBasicAuth("defaultkey", "")
	} else if auth == "provider" {
		req.Header.Set("PRIVATE-TOKEN", "test-token")
	} else if auth != "" {
		req.Header.Set("Authorization", "Bearer "+auth)
	}
	response, err := h.http.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	return response.StatusCode, data, err
}
func (h *harness) rpc(name string, payload any) (map[string]any, error) {
	encoded, _ := json.Marshal(payload)
	address := h.nakama + "/v2/rpc/" + name + "?http_key=local-test"
	var body any = string(encoded)
	unwrapped := name == "update_instance_info" || name == "delete_instance_info"
	if unwrapped {
		address += "&unwrap"
		body = payload
	}
	status, data, err := h.request("POST", address, body, "")
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, fmt.Errorf("RPC %s HTTP %d: %s", name, status, data)
	}
	if unwrapped {
		return map[string]any{}, nil
	}
	var envelope struct {
		Payload string `json:"payload"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, err
	}
	result := map[string]any{}
	if envelope.Payload != "" {
		err = json.Unmarshal([]byte(envelope.Payload), &result)
	}
	return result, err
}
func (h *harness) state() (map[string]any, error) {
	status, data, err := h.request("GET", h.provider+"/_test/state", nil, "provider")
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, fmt.Errorf("mock state HTTP %d", status)
	}
	var state map[string]any
	err = json.Unmarshal(data, &state)
	return state, err
}
func (h *harness) read(id string) (map[string]any, error) {
	return h.rpc("i3d_smoke", map[string]any{"op": "read", "id": id})
}
func (h *harness) player(index int) (*websocket.Conn, string, error) {
	status, data, err := h.request("POST", h.nakama+"/v2/account/authenticate/device?create=true", map[string]string{"id": fmt.Sprintf("smoke-%d-%d", time.Now().UnixNano(), index)}, "basic")
	if err != nil {
		return nil, "", err
	}
	if status != 200 {
		return nil, "", fmt.Errorf("authenticate HTTP %d", status)
	}
	var session struct {
		Token string `json:"token"`
	}
	if err = json.Unmarshal(data, &session); err != nil {
		return nil, "", err
	}
	endpoint := "ws" + strings.TrimPrefix(h.nakama, "http") + "/ws?lang=en&status=true&format=json&token=" + url.QueryEscape(session.Token)
	conn, _, err := websocket.Dial(h.ctx, endpoint, nil)
	return conn, session.Token, err
}
func (h *harness) waitNotification(conn *websocket.Conn) error {
	for {
		var envelope map[string]json.RawMessage
		if err := wsjson.Read(h.ctx, conn, &envelope); err != nil {
			return err
		}
		if value := envelope["error"]; value != nil {
			return fmt.Errorf("socket error: %s", value)
		}
		raw := envelope["notifications"]
		if raw == nil {
			continue
		}
		var message struct {
			Notifications []struct {
				Code    int    `json:"code"`
				Content string `json:"content"`
			} `json:"notifications"`
		}
		if err := json.Unmarshal(raw, &message); err != nil {
			return err
		}
		for _, notification := range message.Notifications {
			if notification.Code != 9000 {
				continue
			}
			var content struct {
				IpAddress string
				Port      int
				SessionId string
			}
			if err := json.Unmarshal([]byte(notification.Content), &content); err != nil {
				return err
			}
			if content.IpAddress != "203.0.113.10" || content.Port != 7777 || content.SessionId != "" {
				return fmt.Errorf("invalid connection notification: %+v", content)
			}
			return nil
		}
	}
}
func (h *harness) run() error {
	sockets := make([]*websocket.Conn, 0, 2)
	for i := 0; i < 2; i++ {
		conn, token, err := h.player(i)
		if err != nil {
			return err
		}
		sockets = append(sockets, conn)
		defer conn.Close(websocket.StatusNormalClosure, "done")
		for _, rpc := range []string{"update_instance_info", "delete_instance_info", "i3d_smoke"} {
			status, _, err := h.request("POST", h.nakama+"/v2/rpc/"+rpc, "{}", token)
			if err != nil {
				return err
			}
			if status != 403 {
				return fmt.Errorf("player invoked %s: HTTP %d", rpc, status)
			}
		}
	}
	for _, conn := range sockets {
		if err := wsjson.Write(h.ctx, conn, map[string]any{"cid": "one-ticket", "matchmaker_add": map[string]any{"min_count": 2, "max_count": 2, "query": "*", "string_properties": map[string]string{"duration": "1800"}}}); err != nil {
			return err
		}
	}
	outcomes := make(chan error, 2)
	for _, conn := range sockets {
		go func(conn *websocket.Conn) { outcomes <- h.waitNotification(conn) }(conn)
	}
	for i := 0; i < 2; i++ {
		if err := <-outcomes; err != nil {
			return err
		}
	}
	state, err := h.state()
	if err != nil {
		return err
	}
	if state["allocations"] != float64(1) {
		return fmt.Errorf("wanted one allocation: %v", state["allocations"])
	}
	id := "723709572903"
	stored, err := h.read(id)
	if err != nil {
		return err
	}
	if stored["exists"] != true {
		return fmt.Errorf("notification arrived without stored session")
	}
	instance := stored["instance"].(map[string]any)
	metadata := instance["metadata"].(map[string]any)
	if instance["player_count"] != float64(2) || metadata["i3d_max_players"] != float64(2) || metadata["duration"] != "1800" {
		return fmt.Errorf("incorrect stored allocation: %v", instance)
	}
	if _, err = h.rpc("i3d_smoke", map[string]any{"op": "race"}); err != nil {
		return err
	}
	if _, err = h.rpc("update_instance_info", map[string]any{"id": id, "player_count": 1, "metadata": map[string]string{"map": "arena"}}); err != nil {
		return err
	}
	stored, err = h.read(id)
	if err != nil {
		return err
	}
	instance = stored["instance"].(map[string]any)
	metadata = instance["metadata"].(map[string]any)
	if instance["player_count"] != float64(1) || metadata["map"] != "arena" || metadata["i3d_max_players"] != float64(2) {
		return fmt.Errorf("lifecycle update lost state: %v", instance)
	}
	if _, err = h.rpc("delete_instance_info", map[string]any{"id": id}); err != nil {
		return err
	}
	stored, err = h.read(id)
	if err != nil {
		return err
	}
	if stored["exists"] != false {
		return fmt.Errorf("restart left stale cache")
	}
	state, err = h.state()
	if err != nil {
		return err
	}
	if state["restarts"] != float64(1) || state["updates"] != float64(1) {
		return fmt.Errorf("wrong provider mutation counts")
	}
	created, err := h.rpc("i3d_smoke", map[string]any{"op": "create"})
	if err != nil {
		return err
	}
	if created["success"] != true {
		return fmt.Errorf("recovery allocation failed")
	}
	recoveryID := created["id"].(string)
	status, _, err := h.request("POST", h.provider+"/_test/instances/"+recoveryID+"/status", map[string]int{"status": 4}, "provider")
	if err != nil {
		return err
	}
	if status != 200 {
		return fmt.Errorf("mock termination HTTP %d", status)
	}
	deadline := time.NewTimer(10 * time.Second)
	defer deadline.Stop()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		stored, err = h.read(recoveryID)
		if err != nil {
			return err
		}
		if stored["exists"] == false {
			break
		}
		select {
		case <-ticker.C:
		case <-deadline.C:
			return fmt.Errorf("missed termination was not reconciled")
		case <-h.ctx.Done():
			return h.ctx.Err()
		}
	}
	state, err = h.state()
	if err != nil {
		return err
	}
	if state["allocations"] != float64(2) || state["restarts"] != float64(1) || state["list_pages"].(float64) < 2 {
		return fmt.Errorf("reconciliation restarted a server or missed pagination: %v", state)
	}
	for _, behavior := range []map[string]int{{"response_status": 4}, {"response_status": 6}, {"delay_ms": 1000}} {
		before, err := h.state()
		if err != nil {
			return err
		}
		status, _, err := h.request("POST", h.provider+"/_test/allocation", behavior, "provider")
		if err != nil {
			return err
		}
		if status != 200 {
			return fmt.Errorf("test behavior HTTP %d", status)
		}
		failed, err := h.rpc("i3d_smoke", map[string]any{"op": "create"})
		if err != nil {
			return err
		}
		if failed["success"] != false {
			return fmt.Errorf("invalid readiness/timeout unexpectedly succeeded")
		}
		after, err := h.state()
		if err != nil {
			return err
		}
		if after["allocations"].(float64) != before["allocations"].(float64)+1 {
			return fmt.Errorf("ambiguous allocation was retried")
		}
		if behavior["response_status"] == 6 {
			// A scoped ALLOCATING response identifies the failed allocation.
			// Cleanup runs after the error callback; wait for its one restart.
			until := time.Now().Add(3 * time.Second)
			for after["restarts"] == before["restarts"] && time.Now().Before(until) {
				select {
				case <-time.After(50 * time.Millisecond):
				case <-h.ctx.Done():
					return h.ctx.Err()
				}
				after, err = h.state()
				if err != nil {
					return err
				}
			}
			if after["restarts"].(float64) != before["restarts"].(float64)+1 {
				return fmt.Errorf("known failed allocation was not reclaimed once: %v", after)
			}
		} else if after["restarts"] != before["restarts"] {
			return fmt.Errorf("ambiguous allocation triggered an unsafe restart")
		}
		status, _, err = h.request("POST", h.provider+"/_test/instances/"+id+"/status", map[string]int{"status": 4}, "provider")
		if err != nil || status != 200 {
			return fmt.Errorf("reset fixture failed")
		}
	}
	fmt.Println("Smoke passed: two clients, one matchmaking allocation, persisted metadata before notifications, native storage conflicts, server-only lifecycle update/restart, missed-termination recovery across provider pages, invalid readiness, known-allocation cleanup and bounded allocation timeout.")
	return nil
}
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: smoke NAKAMA_URL MOCK_URL")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 75*time.Second)
	defer cancel()
	h := &harness{ctx: ctx, nakama: os.Args[1], provider: os.Args[2], http: &http.Client{Timeout: 15 * time.Second}}
	if err := h.run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
