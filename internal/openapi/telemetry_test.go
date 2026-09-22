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
