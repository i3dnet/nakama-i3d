package mockserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// MockServer represents a mock i3D Fleet Manager API server
type MockServer struct {
	server   *http.Server
	mux      *http.ServeMux
	requests map[string][]http.Request
	mutex    sync.Mutex
}

// NewMockServer creates and returns a new MockServer
func NewMockServer() *MockServer {
	mux := http.NewServeMux()

	ms := &MockServer{
		mux:      mux,
		requests: make(map[string][]http.Request),
	}

	// Register all the mock endpoints
	ms.registerHandlers()

	return ms
}

// Start starts the mock server on the given port
func (ms *MockServer) Start(port string) error {
	ms.server = &http.Server{
		Addr:    port,
		Handler: ms.mux,
	}

	log.Printf("Starting mock Fleet Manager API server on %s", port)
	return ms.server.ListenAndServe()
}

// Stop stops the mock server
func (ms *MockServer) Stop() error {
	if ms.server != nil {
		return ms.server.Close()
	}
	return nil
}

// GetRequests returns all requests made to a specific endpoint
func (ms *MockServer) GetRequests(endpoint string) []http.Request {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	return ms.requests[endpoint]
}

// recordRequest records an incoming request for later inspection
func (ms *MockServer) recordRequest(endpoint string, r *http.Request) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.requests[endpoint] = append(ms.requests[endpoint], *r)
}

// registerHandlers sets up all the mock API endpoints
func (ms *MockServer) registerHandlers() {
	// Vehicles endpoint
	ms.mux.HandleFunc("PUT /api/v3/applicationInstance/game/{applicationId}/empty/allocate", func(w http.ResponseWriter, r *http.Request) {
		ms.recordRequest("/api/v3/applicationInstance/game/{applicationId}/empty/allocate", r)

		log.Printf("Received request to allocate an instance")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{
				"id":         fmt.Sprintf("mock-instance-%s", r.PathValue("applicationId")),
				"status":     1,
				"createdAt":  1625097600,
				"numPlayers": 5,
				"ipAddress": []map[string]interface{}{
					{"ipAddress": "127.0.0.1", "ipVersion": 4, "private": 0},
				},
				"properties": []map[string]interface{}{
					{"id": "1", "propertyKey": "port", "propertyValue": "7350", "propertyType": 1},
				},
				"metadata": []map[string]interface{}{
					{"key": "region", "value": "us-west-2"},
				},
			},
		})
	})
}
