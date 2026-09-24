package api

import (
	"bytes"
	"encoding/json"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/internal/configs"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func fixture(t *testing.T) *GinMapping {
	t.Helper()
	m := NewGinMapping(&configs.Config{})
	if err := m.create(); err != nil {
		t.Fatal(err)
	}
	return m
}
func call(m *GinMapping, method, path, body string, headers map[string]string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("PRIVATE-TOKEN", "test-token")
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, r)
	return w
}
func TestLifecyclePersistsStateAndUsesRealRoutes(t *testing.T) {
	m := fixture(t)
	if w := call(m, "PUT", "/v3/applicationInstance/game/6268349608002583795/empty/allocate", `{}`, nil); w.Code != 400 {
		t.Fatalf("missing metadata accepted: %d", w.Code)
	}
	w := call(m, "PUT", "/v3/applicationInstance/game/6268349608002583795/empty/allocate", `{"metadata":[{"key":"mode","value":"arena"}]}`, nil)
	if w.Code != 200 {
		t.Fatal(w.Code, w.Body.String())
	}
	var instances []map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &instances); err != nil {
		t.Fatal(err)
	}
	id := instances[0]["id"].(string)
	w = call(m, "GET", "/v3/applicationInstance/"+id, "", nil)
	if !bytes.Contains(w.Body.Bytes(), []byte(`"status":5`)) || !bytes.Contains(w.Body.Bytes(), []byte("arena")) {
		t.Fatal("allocation not persisted", w.Body.String())
	}
	instances[0]["metadata"] = []map[string]string{{"key": "mode", "value": "updated"}}
	body, _ := json.Marshal(instances[0])
	w = call(m, "PUT", "/v3/applicationInstance/"+id, string(body), nil)
	if w.Code != 200 {
		t.Fatal(w.Code, w.Body.String())
	}
	w = call(m, "GET", "/v3/applicationInstance/"+id, "", nil)
	if !bytes.Contains(w.Body.Bytes(), []byte("updated")) {
		t.Fatal("update not persisted")
	}
	w = call(m, "POST", "/v3/applicationInstance/"+id+"/restart", "", nil)
	var tasks []map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil || len(tasks) != 1 {
		t.Fatal("restart must return task array", w.Body.String())
	}
	if w.Code != 200 {
		t.Fatal(w.Code, w.Body.String())
	}
	w = call(m, "GET", "/v3/applicationInstance/"+id, "", nil)
	if !bytes.Contains(w.Body.Bytes(), []byte(`"status":4`)) {
		t.Fatal("restart not persisted")
	}
}
func TestPaginationAndFilters(t *testing.T) {
	m := fixture(t)
	w := call(m, "GET", "/v3/applicationInstance?filters=applicationId%3D6268349608002583795", "", map[string]string{"RANGED-DATA": "results=1"})
	if w.Code != 200 || w.Header().Get("PAGE-TOKEN") == "" {
		t.Fatal("missing next page", w.Code, w.Body.String())
	}
	first := w.Body.String()
	w = call(m, "GET", "/v3/applicationInstance", "", map[string]string{"RANGED-DATA": "results=1", "PAGE-TOKEN": w.Header().Get("PAGE-TOKEN")})
	if w.Code != 200 || w.Header().Get("PAGE-TOKEN") != "" || w.Body.String() == first {
		t.Fatal("invalid final page", w.Code, w.Body.String())
	}
	for _, path := range []string{"/v3/applicationInstance?filters=status%3D5", "/v3/applicationInstance?filters=bogus%3D1"} {
		if w := call(m, "GET", path, "", nil); w.Code != 400 {
			t.Fatal("unsupported filter accepted", w.Code)
		}
	}
	if w := call(m, "GET", "/v3/applicationInstance?filters=applicationId%3D999", "", nil); w.Code != 200 || w.Body.String() != "[]" {
		t.Fatal(w.Code, w.Body.String())
	}
}
func TestConcurrentReadsAndMutations(t *testing.T) {
	m := fixture(t)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			call(m, http.MethodGet, "/v3/applicationInstance", "", nil)
			call(m, http.MethodPost, "/v3/applicationInstance/723709572903/restart", "", nil)
		}()
	}
	wg.Wait()
}
