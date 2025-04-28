package fleetmanager

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/codewizards/nakama-i3d/mockserver"
)

func TestFleetManager_i3d_MockServer_Benchmark_Test(t *testing.T) {
	mockServer := mockserver.NewMockServer()
	go mockServer.Start(":8765")
	defer mockServer.Stop()

	// Number of concurrent requests to make
	numRequests := 20

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Create multiple calls in separate goroutines
	for i := 0; i < numRequests; i++ {
		go func(requestNum int) {
			defer wg.Done()

			// Create an HTTP client
			client := &http.Client{}

			// Make a request to the mock server
			url := "http://localhost:8765/api/v3/applicationInstance/game/12345/empty/allocate"
			req, err := http.NewRequest(http.MethodPut, url, nil)
			if err != nil {
				t.Errorf("Failed to create request %d: %v", requestNum, err)
				return
			}

			// Add headers similar to what the real client would use
			req.Header.Add("content-type", "application/json")
			req.Header.Add("PRIVATE-TOKEN", "test-token")

			// Execute the request
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Request %d failed: %v", requestNum, err)
				return
			}

			defer resp.Body.Close()

			// You can add assertions here if needed
			t.Logf("Request %d completed with status: %s", requestNum, resp.Status)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Allow time for log messages to print
	time.Sleep(1 * time.Second)

	t.Logf("All requests completed")
}

func TestFleetManager_i3d_Benchmark_Full_Api_Client_Test(t *testing.T) {
	mockServer := mockserver.NewMockServer()
	go mockServer.Start(":8765")
	defer mockServer.Stop()

	client := NewFleetManagerApiClient("test-token", "http://localhost:8765/api", "12345", nil)

	// Number of concurrent requests to make
	numRequests := 20

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Create multiple calls in separate goroutines
	for i := 0; i < numRequests; i++ {
		go func(requestNum int) {
			defer wg.Done()

			metadata := make(map[string]any)
			metadata["applicationId"] = strconv.Itoa(requestNum + 1)
			instance, err := client.Create(context.Background(), 10, []string{"1", "2", "3", "4", "5"}, nil, metadata, nil)
			if err != nil {
				t.Errorf("Failed to create instance %d: %v", requestNum, err)
				return
			}

			t.Logf("created instance with id: %v", instance.Id)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Allow time for log messages to print
	time.Sleep(1 * time.Second)

	t.Logf("All requests completed")
}
