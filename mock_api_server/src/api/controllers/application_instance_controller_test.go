package controller

import (
	"encoding/json"
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
