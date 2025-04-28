package fleetmanager

import (
	"os"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestLoadDefaultConfig(t *testing.T) {
	// Create a temporary config file for testing
	tempConfig := `{
		"fleet_manager_api_url": "https://test.api.com",
		"fleet_manager_api_token": "test-token",
		"application_id": "test-app-id"
	}`

	err := os.WriteFile("test-config.json", []byte(tempConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	defer os.Remove("test-config.json")

	tests := []struct {
		name      string
		env       map[string]string
		wantURL   string
		wantToken string
		wantAppID string
		logger    runtime.Logger
	}{
		{
			name:      "Load from config file only",
			env:       map[string]string{},
			wantURL:   "https://test.api.com",
			wantToken: "test-token",
			wantAppID: "test-app-id",
		},
		{
			name: "Override with environment variables",
			env: map[string]string{
				"FLEET_MANAGER_API_URL":   "https://env.api.com",
				"FLEET_MANAGER_API_TOKEN": "env-token",
				"APPLICATION_ID":          "env-app-id",
			},
			wantURL:   "https://env.api.com",
			wantToken: "env-token",
			wantAppID: "env-app-id",
		},
		{
			name: "Partial environment override",
			env: map[string]string{
				"FLEET_MANAGER_API_URL": "https://partial.api.com",
			},
			wantURL:   "https://partial.api.com",
			wantToken: "test-token",
			wantAppID: "test-app-id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := LoadDefaultConfig(tt.env, tt.logger, "test-config.json")

			if cfg.FleetManagerApiUrl != tt.wantURL {
				t.Errorf("LoadDefaultConfig() FleetManagerApiUrl = %v, want %v",
					cfg.FleetManagerApiUrl, tt.wantURL)
			}
			if cfg.FleetManagerApiToken != tt.wantToken {
				t.Errorf("LoadDefaultConfig() FleetManagerApiToken = %v, want %v",
					cfg.FleetManagerApiToken, tt.wantToken)
			}
			if cfg.ApplicationId != tt.wantAppID {
				t.Errorf("LoadDefaultConfig() ApplicationId = %v, want %v",
					cfg.ApplicationId, tt.wantAppID)
			}
		})
	}
}

func TestVerifyFleetManagerConfig(t *testing.T) {
	tests := []struct {
		name      string
		config    FleetManagerConfig
		wantError bool
	}{
		{
			name: "Valid config",
			config: FleetManagerConfig{
				FleetManagerApiUrl:   "https://api.test.com",
				FleetManagerApiToken: "valid-token",
				ApplicationId:        "valid-app-id",
			},
			wantError: false,
		},
		{
			name: "Missing API URL",
			config: FleetManagerConfig{
				FleetManagerApiUrl:   "",
				FleetManagerApiToken: "valid-token",
				ApplicationId:        "valid-app-id",
			},
			wantError: true,
		},
		{
			name: "Missing API token",
			config: FleetManagerConfig{
				FleetManagerApiUrl:   "https://api.test.com",
				FleetManagerApiToken: "",
				ApplicationId:        "valid-app-id",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.VerifyFleetManagerConfig()
			if (err != nil) != tt.wantError {
				t.Errorf("VerifyFleetManagerConfig() error = %v, wantError %v",
					err, tt.wantError)
			}
		})
	}
}
