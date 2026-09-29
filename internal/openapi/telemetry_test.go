package openapi

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

type responseTransport func(*http.Request) (*http.Response, error)

func (f responseTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestProviderTelemetryDoesNotLogBodiesOrCredentials(t *testing.T) {
	var captured bytes.Buffer
	old := log.Writer()
	log.SetOutput(&captured)
	defer log.SetOutput(old)
	cfg := NewConfiguration()
	cfg.Debug = true
	cfg.HTTPClient = &http.Client{Transport: responseTransport(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 503, Header: http.Header{"Set-Cookie": []string{"secret-cookie"}}, Body: io.NopCloser(strings.NewReader("private-response")), Request: r}, nil
	})}
	client := NewAPIClient(cfg)
	request, _ := http.NewRequestWithContext(context.Background(), "PUT", "https://example.test/v3/applicationInstance/game/private-app/empty/allocate?filters=private-filter", strings.NewReader("private-body"))
	request.Header.Set("PRIVATE-TOKEN", "private-token")
	response, err := client.callAPI(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	if string(body) != "private-response" {
		t.Fatal("telemetry consumed the response")
	}
	for _, secret := range []string{"private-response", "secret-cookie", "private-app", "private-filter", "private-body", "private-token"} {
		if strings.Contains(captured.String(), secret) {
			t.Errorf("logged %q", secret)
		}
	}
}

func TestTelemetryRoutesWithDeploymentPrefixes(t *testing.T) {
	for _, prefix := range []string{"", "/proxy/private-tenant", "/proxy%2Fprivate-tenant", "/v3/applicationInstance/proxy"} {
		for _, tc := range []struct{ path, route string }{
			{"/v3/applicationInstance", "/v3/applicationInstance"},
			{"/v3/applicationInstance/", "/v3/applicationInstance"},
			{"/v3/applicationInstance/private-instance", "/v3/applicationInstance/{instanceId}"},
			{"/v3/applicationInstance/private-instance/restart", "/v3/applicationInstance/{instanceId}/restart"},
			{"/v3/applicationInstance/game/private-app/empty/allocate", "/v3/applicationInstance/game/{applicationId}/empty/allocate"},
			{"/v3/applicationInstance/unknown/nested/restart", "/other"},
			{"/v3/applicationInstance/game/extra/private-app/empty/allocate", "/other"},
			{"/v3/applicationInstances", "/other"},
			{"/prefix-v3/applicationInstance/private-instance", "/other"},
			{"/v3/applicationInstanceExtra/private-instance", "/other"},
			{"/prefix%2Fv3/applicationInstance/private-instance", "/other"},
			{"/v3%2FapplicationInstance/private-instance", "/other"},
			{"/v3/applicationInstance/private-instance%2Frestart", "/v3/applicationInstance/{instanceId}"},
			{"/v3/applicationInstance/private-instance%2Frestart/restart", "/v3/applicationInstance/{instanceId}/restart"},
		} {
			t.Run(prefix+tc.path, func(t *testing.T) {
				var captured bytes.Buffer
				old := log.Writer()
				log.SetOutput(&captured)
				defer log.SetOutput(old)
				cfg := NewConfiguration()
				cfg.Debug = true
				cfg.Servers = ServerConfigurations{{URL: "https://example.test" + prefix}}
				cfg.HTTPClient = &http.Client{Transport: responseTransport(func(r *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("[]")), Header: make(http.Header), Request: r}, nil
				})}
				request, err := http.NewRequest("GET", "https://example.test"+prefix+tc.path+"?filters=private-filter", nil)
				if err != nil {
					t.Fatal(err)
				}
				response, err := NewAPIClient(cfg).callAPI(request)
				if err != nil {
					t.Fatal(err)
				}
				response.Body.Close()
				expected := "provider request: GET " + tc.route + "\n"
				if !strings.Contains(captured.String(), expected) {
					t.Fatalf("want %q, got %q", expected, captured.String())
				}
				for _, secret := range []string{"private-tenant", "private-instance", "private-app", "private-filter"} {
					if strings.Contains(captured.String(), secret) {
						t.Errorf("logged %s", secret)
					}
				}
			})
		}
	}
}
