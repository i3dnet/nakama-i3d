package openapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Capture actual HTTP requests so an unsubstituted or incorrectly escaped path
// fails without making calls to a live provider. Response mapping is tested elsewhere.
func TestGeneratedRoutesSubstituteRequiredIDs(t *testing.T) {
	for _, tc := range []struct {
		name, want string
		call       func(*APIClient)
	}{
		{"GetPatchJobApplicationInstanceFailedApplicationInstances", "GET /v3/patchJob/patchJobId%2Fone%20two/applicationInstance/applicationInstanceId%2Fone%20two/failedApplicationInstance", func(client *APIClient) {
			_, _, _ = client.PatchJobAPI.GetPatchJobApplicationInstanceFailedApplicationInstances(context.Background(), "patchJobId/one two", "applicationInstanceId/one two").Execute()
		}},
		{"GetPatchJobFailedApplicationInstance", "GET /v3/patchJob/patchJobId%2Fone%20two/failedApplicationInstances", func(client *APIClient) {
			_, _, _ = client.PatchJobAPI.GetPatchJobFailedApplicationInstance(context.Background(), "patchJobId/one two").RANGEDDATA("results=1").Execute()
		}},
		{"GetPatchJobReportProgress", "GET /v3/patchJob/patchJobId%2Fone%20two/report/progress", func(client *APIClient) {
			_, _, _ = client.PatchJobAPI.GetPatchJobReportProgress(context.Background(), "patchJobId/one two").Execute()
		}},
		{"CreateCustomCommandApplicationInstance", "POST /v3/customCommand/applicationInstance/applicationInstanceId%2Fone%20two", func(client *APIClient) {
			_, _, _ = client.ArcusAPI.CreateCustomCommandApplicationInstance(context.Background(), "applicationInstanceId/one two").CustomCommandCollection(CustomCommandCollection{}).Execute()
		}},
		{"CreateCustomCommandHost", "POST /v3/customCommand/host/1234", func(client *APIClient) {
			_, _, _ = client.ArcusAPI.CreateCustomCommandHost(context.Background(), 1234).CustomCommandCollection(CustomCommandCollection{}).Execute()
		}},
		{"GetTelemetryApplication", "GET /v3/telemetry/application/applicationId%2Fone%20two", func(client *APIClient) {
			_, _, _ = client.ApplicationTelemetryAPI.GetTelemetryApplication(context.Background(), "applicationId/one two").Execute()
		}},
		{"GetTelemetryApplicationCurrents", "GET /v3/telemetry/application/applicationId%2Fone%20two/current", func(client *APIClient) {
			_, _, _ = client.ApplicationTelemetryAPI.GetTelemetryApplicationCurrents(context.Background(), "applicationId/one two").Execute()
		}},
		{"GetFleetHostBulkReserveStatuses", "GET /v3/fleet/host/bulkReserve/traceId%2Fone%20two/status", func(client *APIClient) {
			_, _, _ = client.FleetAPI.GetFleetHostBulkReserveStatuses(context.Background(), "traceId/one two").Execute()
		}},
		{"GetTelemetryDeploymentEnvironmentCurrents", "GET /v3/telemetry/deploymentEnvironment/deploymentEnvironmentId%2Fone%20two/current", func(client *APIClient) {
			_, _, _ = client.DeploymentEnvironmentTelemetryAPI.GetTelemetryDeploymentEnvironmentCurrents(context.Background(), "deploymentEnvironmentId/one two").Execute()
		}},
		{"GetTelemetryApplicationBuild", "GET /v3/telemetry/applicationBuild/applicationBuildId%2Fone%20two", func(client *APIClient) {
			_, _, _ = client.ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuild(context.Background(), "applicationBuildId/one two").Execute()
		}},
		{"GetTelemetryApplicationBuildCurrents", "GET /v3/telemetry/applicationBuild/applicationBuildId%2Fone%20two/current", func(client *APIClient) {
			_, _, _ = client.ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuildCurrents(context.Background(), "applicationBuildId/one two").Execute()
		}},
		{"GetApplicationDeploymentEnvironments", "GET /v3/application/applicationId%2Fone%20two/deploymentEnvironment", func(client *APIClient) {
			_, _, _ = client.ApplicationAPI.GetApplicationDeploymentEnvironments(context.Background(), "applicationId/one two").Execute()
		}},
		{"GetApplicationDeploymentTemplates", "GET /v3/application/applicationId%2Fone%20two/deploymentTemplate", func(client *APIClient) {
			_, _, _ = client.ApplicationAPI.GetApplicationDeploymentTemplates(context.Background(), "applicationId/one two").Execute()
		}},
		{"GetApplicationFleets", "GET /v3/application/applicationId%2Fone%20two/fleet", func(client *APIClient) {
			_, _, _ = client.ApplicationAPI.GetApplicationFleets(context.Background(), "applicationId/one two").Execute()
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			requests := make(chan string, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests <- r.Method + " " + r.URL.EscapedPath()
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte("[]"))
			}))
			defer server.Close()
			cfg := NewConfiguration()
			cfg.Servers = ServerConfigurations{{URL: server.URL}}
			tc.call(NewAPIClient(cfg))
			select {
			case got := <-requests:
				if got != tc.want {
					t.Fatalf("request = %q, want %q", got, tc.want)
				}
			default:
				t.Fatal("client did not issue an HTTP request")
			}
		})
	}
}
