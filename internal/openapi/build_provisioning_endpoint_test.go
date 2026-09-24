package openapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildProvisioningFilesUsesRegistrationIDInPath(t *testing.T) {
	requests := make(chan string, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests <- r.Method + " " + r.URL.EscapedPath()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: server.URL}}
	client := NewAPIClient(cfg)
	_, _, err := client.BuildProvisioningAPI.GetBuildProvisioningStorageFilesByRegistrationId(context.Background(), "registration-123").Execute()
	if err != nil {
		t.Fatalf("list registration files: %v", err)
	}
	want := "GET /v3/buildProvisioning/storage/registration/registration-123/file"
	if got := <-requests; got != want {
		t.Fatalf("request = %q, want %q", got, want)
	}
}
