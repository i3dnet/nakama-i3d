package clients

import (
	"context"
	"encoding/json"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/openapi"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func validProviderInstance() openapi.ApplicationInstance {
	return openapi.ApplicationInstance{Id: "instance-1", ApplicationId: "123", Status: 5, IpAddress: []openapi.ApplicationInstanceIP{{IpAddress: "203.0.113.10", Private: 0, IpVersion: 4}}, Properties: []openapi.ApplicationInstanceProperty{{PropertyKey: "port", PropertyValue: "7777"}}}
}
func contractClient(t *testing.T, handler http.HandlerFunc) *OneApiClient {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg := &config.Config{OneApi: config.OneApi{BaseUrl: server.URL, ApplicationId: "123", Token: "test-only"}, Retry: config.Retry{Attempts: 1}}
	return NewOneApiClient(cfg, NewAuthentication(cfg), tests.NewMockLogger())
}
func writeInstances(w http.ResponseWriter, instances []openapi.ApplicationInstance) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(instances)
}

func TestAllocationSendsMetadataAndRawFilterOnce(t *testing.T) {
	var method, path, filter string
	var body openapi.MetadataCollection
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		method, path, filter = r.Method, r.URL.Path, r.URL.Query().Get("filters")
		_ = json.NewDecoder(r.Body).Decode(&body)
		instance := validProviderInstance()
		instance.ApplicationId = "456"
		writeInstances(w, []openapi.ApplicationInstance{instance})
	})
	_, err := client.AllocateApplicationInstance(context.Background(), map[string]any{"map": "arena", ApplicationId: "456", I3dFilters: "routing", "i3d_max_players": 10}, `fleetName="EU \"West\""`)
	require.NoError(t, err)
	require.Equal(t, http.MethodPut, method)
	require.Equal(t, "/v3/applicationInstance/game/456/empty/allocate", path)
	require.Equal(t, `fleetName="EU \"West\""`, filter)
	require.Equal(t, []openapi.Metadata{{Key: "map", Value: "arena"}}, body.Metadata)
}
func TestListSendsPageHeadersAndRetainsCursor(t *testing.T) {
	var cursors, ranges []string
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		cursors = append(cursors, r.Header.Get("PAGE-TOKEN"))
		ranges = append(ranges, r.Header.Get("RANGED-DATA"))
		if len(cursors) == 1 {
			w.Header().Set("PAGE-TOKEN", "second-page")
		}
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
	})
	first, err := client.ListApplicationInstances(context.Background(), "applicationId=123", 7, "")
	require.NoError(t, err)
	second, err := client.ListApplicationInstances(context.Background(), "applicationId=123", 7, first.NextCursor)
	require.NoError(t, err)
	require.Equal(t, []string{"", "second-page"}, cursors)
	require.Equal(t, []string{"results=7", "results=7"}, ranges)
	require.Empty(t, second.NextCursor)
}
func TestUpdateTargetsInstanceForReadAndWrite(t *testing.T) {
	var requests []string
	var update openapi.ApplicationInstance
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		instance := validProviderInstance()
		if r.Method == http.MethodPut {
			_ = json.NewDecoder(r.Body).Decode(&update)
			instance = update
		}
		writeInstances(w, []openapi.ApplicationInstance{instance})
	})
	got, err := client.UpdateApplicationInstance(context.Background(), "instance-1", map[string]any{"map": "arena"})
	require.NoError(t, err)
	require.Equal(t, []string{"GET /v3/applicationInstance/instance-1", "PUT /v3/applicationInstance/instance-1"}, requests)
	require.Equal(t, []openapi.Metadata{{Key: "map", Value: "arena"}}, update.Metadata)
	require.Equal(t, "arena", got.Metadata["map"])
}
func TestEmptyProviderResponsesReturnErrors(t *testing.T) {
	for _, operation := range []string{"get", "allocate", "update-get", "update-put"} {
		t.Run(operation, func(t *testing.T) {
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
				if operation == "update-put" && r.Method == http.MethodGet {
					writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
					return
				}
				writeInstances(w, []openapi.ApplicationInstance{})
			})
			require.NotPanics(t, func() {
				var err error
				switch operation {
				case "get":
					_, err = client.GetApplicationInstance(context.Background(), "instance-1")
				case "allocate":
					_, err = client.AllocateApplicationInstance(context.Background(), nil, "")
				default:
					_, err = client.UpdateApplicationInstance(context.Background(), "instance-1", nil)
				}
				require.Error(t, err)
			})
		})
	}
}
func TestInvalidConnectionDataReturnsErrors(t *testing.T) {
	for _, kind := range []string{"no-address", "private-address", "bad-address", "no-port", "bad-port", "zero-port", "large-port", "no-id"} {
		t.Run(kind, func(t *testing.T) {
			instance := validProviderInstance()
			switch kind {
			case "no-address":
				instance.IpAddress = nil
			case "private-address":
				instance.IpAddress[0].Private = 1
			case "bad-address":
				instance.IpAddress[0].IpAddress = "garbage"
			case "no-port":
				instance.Properties = nil
			case "bad-port":
				instance.Properties[0].PropertyValue = "abc"
			case "zero-port":
				instance.Properties[0].PropertyValue = "0"
			case "large-port":
				instance.Properties[0].PropertyValue = "65536"
			case "no-id":
				instance.Id = ""
			}
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
				writeInstances(w, []openapi.ApplicationInstance{instance})
			})
			require.NotPanics(t, func() {
				_, err := client.GetApplicationInstance(context.Background(), "instance-1")
				require.Error(t, err)
			})
		})
	}
}
func TestAllocationRequiresAllocatedState(t *testing.T) {
	for _, status := range []int32{4, 6} {
		t.Run(strconv.Itoa(int(status)), func(t *testing.T) {
			instance := validProviderInstance()
			instance.Status = status
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
				writeInstances(w, []openapi.ApplicationInstance{instance})
			})
			_, err := client.AllocateApplicationInstance(context.Background(), nil, "")
			require.Error(t, err)
		})
	}
}
func TestListFailsWholePageOnInvalidInstance(t *testing.T) {
	invalid := validProviderInstance()
	invalid.Properties = nil
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance(), invalid})
	})
	got, err := client.ListApplicationInstances(context.Background(), "", 10, "")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetRejectsMismatchedInstanceIdentity(t *testing.T) {
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		instance := validProviderInstance()
		instance.Id = "other"
		writeInstances(w, []openapi.ApplicationInstance{instance})
	})
	got, err := client.GetApplicationInstance(context.Background(), "instance-1")
	require.Error(t, err)
	require.Nil(t, got)
}
func TestAllocateRejectsMismatchedApplication(t *testing.T) {
	for _, metadata := range []map[string]any{nil, {ApplicationId: "456"}, {ApplicationId: int64(6268349608002583795)}} {
		client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
			instance := validProviderInstance()
			instance.ApplicationId = "other"
			writeInstances(w, []openapi.ApplicationInstance{instance})
		})
		got, err := client.AllocateApplicationInstance(context.Background(), metadata, "")
		require.Error(t, err)
		require.Nil(t, got)
	}
}
func TestProviderRejectsNonPublicAddressesDespitePublicFlag(t *testing.T) {
	for _, address := range []string{"127.0.0.1", "10.1.2.3", "172.16.0.1", "192.168.1.2", "169.254.1.1", "0.0.0.0", "224.0.0.1", "::1", "fc00::1", "fe80::1", "::", "ff02::1", "::ffff:192.168.1.2"} {
		t.Run(address, func(t *testing.T) {
			instance := validProviderInstance()
			instance.IpAddress[0].IpAddress = address
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
				writeInstances(w, []openapi.ApplicationInstance{instance})
			})
			got, err := client.GetApplicationInstance(context.Background(), "instance-1")
			require.Error(t, err)
			require.Nil(t, got)
		})
	}
}

func TestAllocationErrorRetainsOnlyConfirmedAllocationIdentity(t *testing.T) {
	for _, kind := range []string{"missing-port", "allocating", "online", "wrong-application", "empty-id", "whitespace-id", "multiple"} {
		t.Run(kind, func(t *testing.T) {
			instance := validProviderInstance()
			known := false
			switch kind {
			case "missing-port":
				instance.Properties = nil
				known = true
			case "allocating":
				instance.Status = 6
				known = true
			case "online":
				instance.Status = 4
			case "wrong-application":
				instance.ApplicationId = "other"
			case "empty-id":
				instance.Id = ""
			case "whitespace-id":
				instance.Id = " "
			}
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
				response := []openapi.ApplicationInstance{instance}
				if kind == "multiple" {
					response = append(response, instance)
				}
				writeInstances(w, response)
			})
			got, err := client.AllocateApplicationInstance(context.Background(), nil, "")
			require.Error(t, err)
			if known {
				require.NotNil(t, got)
				require.Equal(t, instance.Id, got.Id)
			} else {
				require.Nil(t, got, "unconfirmed identity must never authorize restart")
			}
		})
	}
}
