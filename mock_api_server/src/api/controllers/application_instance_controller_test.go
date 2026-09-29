package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/api/models"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestQuotedFilterConjunctions(t *testing.T) {
	for _, name := range []string{"Rock and Roll", "quote \" and slash \\", "(venue)", "123"} {
		instance := models.GetApplicationInstance()
		instance.RegionName = name
		ok, err := matches(instance, "regionName="+strconv.Quote(name)+" and applicationId="+instance.ApplicationID)
		if err != nil || !ok {
			t.Fatalf("%q: match=%v err=%v", name, ok, err)
		}
	}
	for _, filter := range []string{"regionName=\"unterminated", "regionName=\"x\" or fleetId=1", "regionName=x and ", "unknown=1"} {
		if _, err := matches(models.GetApplicationInstance(), filter); err == nil {
			t.Errorf("accepted malformed filter %q", filter)
		}
	}
}
func controllerRequest(controller *ApplicationInstanceController, method, path, body string, headers map[string]string) *httptest.ResponseRecorder {
	router := gin.New()
	router.GET("/v3/applicationInstance", controller.List)
	router.POST("/v3/applicationInstance/:instanceId/restart", controller.Restart)
	router.PUT("/v3/applicationInstance/:instanceId", controller.Update)
	router.PUT("/v3/applicationInstance/game/:applicationId/empty/allocate", controller.Create)
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}
func TestAllocationMergesMetadataAndDeletesOnlyNull(t *testing.T) {
	controller := NewApplicationInstanceController()
	instance := controller.instances["723709572903"]
	if err := json.Unmarshal([]byte(`[{"key":"keep","value":"old"},{"key":"replace","value":"before"},{"key":"remove","value":"gone"},{"key":"empty","value":"before"}]`), &instance.Metadata); err != nil {
		t.Fatal(err)
	}
	response := controllerRequest(controller, "PUT", "/v3/applicationInstance/game/"+instance.ApplicationID+"/empty/allocate",
		`{"metadata":[{"key":"replace","value":"after"},{"key":"remove","value":null},{"key":"empty","value":""},{"key":"new","value":"added"}]}`, nil)
	if response.Code != 200 {
		t.Fatal(response.Code, response.Body.String())
	}
	data, _ := json.Marshal(instance.Metadata)
	var got []struct {
		Key   string
		Value string
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	values := map[string]string{}
	for _, p := range got {
		values[p.Key] = p.Value
	}
	if len(values) != 4 || values["keep"] != "old" || values["replace"] != "after" || values["empty"] != "" || values["new"] != "added" {
		t.Fatalf("merged metadata=%s", data)
	}
	if _, ok := values["remove"]; ok {
		t.Fatal("null key retained")
	}
}
func TestPaginationRetainsOriginalFilteredSnapshot(t *testing.T) {
	for _, filter := range []string{"", "applicationId=other", "not a valid filter"} {
		t.Run(filter, func(t *testing.T) {
			controller := NewApplicationInstanceController()
			third := models.GetApplicationInstance()
			third.Id = "000"
			third.ApplicationID = "other"
			controller.instances[third.Id] = third
			response := controllerRequest(controller, "GET", "/v3/applicationInstance?filters=applicationId%3D6268349608002583795", "", map[string]string{"RANGED-DATA": "results=1"})
			token := response.Header().Get("PAGE-TOKEN")
			if response.Code != 200 || token == "" {
				t.Fatal(response.Code, response.Body.String())
			}
			controller.instances["723709572904"].ApplicationID = "changed"
			added := models.GetApplicationInstance()
			added.Id = "001"
			controller.instances[added.Id] = added
			response = controllerRequest(controller, "GET", "/v3/applicationInstance?filters="+url.QueryEscape(filter), "", map[string]string{"PAGE-TOKEN": token, "RANGED-DATA": "results=1"})
			var instances []models.ApplicationInstance
			if err := json.Unmarshal(response.Body.Bytes(), &instances); err != nil {
				t.Fatal(err)
			}
			if response.Code != 200 || len(instances) != 1 || instances[0].Id != "723709572904" || instances[0].ApplicationID != "6268349608002583795" || response.Header().Get("PAGE-TOKEN") != "" {
				t.Fatal(response.Code, response.Body.String())
			}
		})
	}
	controller := NewApplicationInstanceController()
	if r := controllerRequest(controller, "GET", "/v3/applicationInstance", "", map[string]string{"PAGE-TOKEN": "made-up"}); r.Code != 400 {
		t.Fatal("unknown cursor accepted")
	}
}

func TestAllocationRejectsInvalidMetadataPatch(t *testing.T) {
	for _, body := range []string{
		`{"metadata":[{"key":"missing"}]}`,
		`{"metadata":[{"key":"number","value":42}]}`,
		`{"metadata":[{"key":"same","value":"one"},{"key":"same","value":null}]}`,
	} {
		controller := NewApplicationInstanceController()
		r := controllerRequest(controller, "PUT", "/v3/applicationInstance/game/6268349608002583795/empty/allocate", body, nil)
		if r.Code != 400 || controller.allocations != 0 {
			t.Fatalf("invalid patch allocated: status=%d body=%s", r.Code, body)
		}
	}
}
func TestExpiredPaginationTokenIsRejected(t *testing.T) {
	controller := NewApplicationInstanceController()
	r := controllerRequest(controller, "GET", "/v3/applicationInstance", "", map[string]string{"RANGED-DATA": "results=1"})
	token := r.Header().Get("PAGE-TOKEN")
	if token == "" {
		t.Fatal("missing token")
	}
	page := controller.pages[token]
	page.expires = time.Now().Add(-time.Second)
	controller.pages[token] = page
	r = controllerRequest(controller, "GET", "/v3/applicationInstance", "", map[string]string{"PAGE-TOKEN": token})
	if r.Code != 400 {
		t.Fatal("expired token accepted")
	}
	if _, ok := controller.pages[token]; ok {
		t.Fatal("expired snapshot retained")
	}
}

func TestUpdateMergesMetadataWithoutChangingProviderState(t *testing.T) {
	controller := NewApplicationInstanceController()
	instance := controller.instances["723709572903"]
	instance.Status, instance.NumPlayers = 5, 3
	instance.Metadata = []models.KeyValue{{Key: "keep", Value: "old"}, {Key: "map", Value: "before"}, {Key: "remove", Value: "gone"}}
	response := controllerRequest(controller, "PUT", "/v3/applicationInstance/"+instance.Id,
		`{"metadata":[{"key":"map","value":"arena"},{"key":"remove","value":null},{"key":"empty","value":""}]}`, nil)
	if response.Code != 200 {
		t.Fatal(response.Code, response.Body.String())
	}
	values := map[string]string{}
	for _, pair := range instance.Metadata {
		values[pair.Key] = pair.Value
	}
	if len(values) != 3 || values["keep"] != "old" || values["map"] != "arena" || values["empty"] != "" {
		t.Fatalf("merged metadata=%v", values)
	}
	if _, ok := values["remove"]; ok {
		t.Fatal("null key retained")
	}
	if instance.NumPlayers != 3 || instance.Status != 5 {
		t.Fatalf("metadata update changed provider telemetry/state: %+v", instance)
	}
	var returned []models.ApplicationInstance
	if err := json.Unmarshal(response.Body.Bytes(), &returned); err != nil || len(returned) != 1 || returned[0].NumPlayers != 3 {
		t.Fatalf("invalid returned provider state: %s (%v)", response.Body.String(), err)
	}
	response = controllerRequest(controller, "PUT", "/v3/applicationInstance/"+instance.Id, `{"metadata":[]}`, nil)
	if response.Code != 200 || len(instance.Metadata) != 3 {
		t.Fatal(response.Code, response.Body.String())
	}
}

func TestUpdateRejectsReadOnlyFieldsAndMalformedMetadata(t *testing.T) {
	for _, body := range []string{
		`{"numPlayers":0,"metadata":[]}`,
		`{"status":4,"metadata":[]}`,
		`{"metadata":[{"key":"missing"}]}`,
		`{"metadata":[{"key":"number","value":42}]}`,
		`{"metadata":[{"key":"same","value":"one"},{"key":"same","value":null}]}`,
	} {
		controller := NewApplicationInstanceController()
		instance := controller.instances["723709572903"]
		instance.Status, instance.NumPlayers = 5, 3
		// Reject read-only fields locally; the real API may reject or ignore them.
		response := controllerRequest(controller, "PUT", "/v3/applicationInstance/"+instance.Id, body, nil)
		if response.Code != 400 || controller.updates != 0 || instance.NumPlayers != 3 || instance.Status != 5 {
			t.Fatalf("invalid update mutated provider: status=%d body=%s", response.Code, body)
		}
	}
}

func TestAllocationExcludesOccupiedAndReservedInstances(t *testing.T) {
	for _, state := range []struct{ status, players int }{{4, 2}, {5, 0}, {6, 0}} {
		t.Run(fmt.Sprintf("status%d-players%d", state.status, state.players), func(t *testing.T) {
			controller := NewApplicationInstanceController()
			instance := controller.instances["723709572903"]
			delete(controller.instances, "723709572904")
			instance.Status, instance.NumPlayers = state.status, state.players
			response := controllerRequest(controller, "PUT", "/v3/applicationInstance/game/"+instance.ApplicationID+"/empty/allocate", `{"metadata":[]}`, nil)
			if response.Code < 400 || controller.allocations != 0 || instance.Status != state.status {
				t.Fatalf("occupied/reserved instance allocated: status=%d response=%s", response.Code, response.Body.String())
			}
		})
	}
}

// This exercises the documented selection/reuse contract in the local service.
// It does not emulate host-agent messages or real asynchronous restart timing.
func TestAllocationUpdateRestartAndReuseLifecycle(t *testing.T) {
	controller := NewApplicationInstanceController()
	instance := controller.instances["723709572903"]
	delete(controller.instances, "723709572904")
	allocate := func(body string) *httptest.ResponseRecorder {
		return controllerRequest(controller, "PUT", "/v3/applicationInstance/game/"+instance.ApplicationID+"/empty/allocate", body, nil)
	}
	if r := allocate(`{"metadata":[{"key":"match","value":"A"}]}`); r.Code != 200 || instance.Status != 5 {
		t.Fatal(r.Code, r.Body.String())
	}
	if r := allocate(`{"metadata":[]}`); r.Code < 400 {
		t.Fatal("reserved server allocated twice")
	}
	// Players are reported by Arcus, represented here by the fixture state.
	instance.NumPlayers = 2
	r := controllerRequest(controller, "PUT", "/v3/applicationInstance/"+instance.Id, `{"metadata":[{"key":"map","value":"arena"}]}`, nil)
	if r.Code != 200 || instance.Status != 5 || instance.NumPlayers != 2 {
		t.Fatal("metadata changed allocation/player state", r.Body.String())
	}
	if r := allocate(`{"metadata":[]}`); r.Code < 400 {
		t.Fatal("occupied server allocated")
	}
	// A completed match alone does not release a reserved instance.
	instance.NumPlayers = 0
	if r := allocate(`{"metadata":[]}`); r.Code < 400 {
		t.Fatal("allocated but empty server reused before release")
	}
	r = controllerRequest(controller, "POST", "/v3/applicationInstance/"+instance.Id+"/restart", "", nil)
	if r.Code != 200 || instance.Status != 4 {
		t.Fatal("restart failed", r.Body.String())
	}
	if r := allocate(`{"metadata":[{"key":"match","value":"B"}]}`); r.Code != 200 || instance.Status != 5 || controller.allocations != 2 || controller.restarts != 1 {
		t.Fatal("server could not be reused after release", r.Body.String())
	}
	if len(instance.Metadata) != 1 || instance.Metadata[0].Value != "B" {
		t.Fatal("replacement allocation metadata missing")
	}
}
