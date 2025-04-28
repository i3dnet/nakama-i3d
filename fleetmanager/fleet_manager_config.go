package fleetmanager

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

// FleetManagerConfig is a struct that holds the configuration for the FleetManager.
// The fields are populated from the config.json file and any matching environment variables.
// These fields are just an example and can be modified to suit the needs of the game. Remember to update the default config.json file with the correct values.
type FleetManagerConfig struct {
	FleetManagerApiUrl   string `json:"fleet_manager_api_url"`   // Fleet Manager API Url
	FleetManagerApiToken string `json:"fleet_manager_api_token"` // Fleet Manager API Token
	ApplicationId        string `json:"application_id"`          // Application ID
	MaxPlayers           int    `json:"max_players"`             // Maximum number of players
}

// LoadDefaultConfig loads the default config from a local file and allows the values to be overwritten by the environment variables.
func LoadDefaultConfig(env map[string]string, logger runtime.Logger, configPath string) FleetManagerConfig {

	if configPath == "" {
		configPath = "config.json"
	}

	f, err := os.Open(configPath)
	if err != nil {
		logger.WithField("error", err).Error("Error opening config.json")
	}
	defer f.Close()

	var cfg FleetManagerConfig
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.WithField("error", err).Error("Error decoding config.json")
	}

	fleetManagerApiUrl, ok := env["FLEET_MANAGER_API_URL"]
	if ok {
		cfg.FleetManagerApiUrl = fleetManagerApiUrl
	}

	fleetManagerApiToken, ok := env["FLEET_MANAGER_API_TOKEN"]
	if ok {
		cfg.FleetManagerApiToken = fleetManagerApiToken
	}

	applicationId, ok := env["APPLICATION_ID"]
	if ok {
		cfg.ApplicationId = applicationId
	}

	maxPlayers, ok := env["MAX_PLAYERS"]
	if ok {
		cfg.MaxPlayers, err = strconv.Atoi(maxPlayers)
		if err != nil {
			logger.WithField("error", err).Error("Error converting max_players to int")
		}
	}

	err = cfg.VerifyFleetManagerConfig()
	if err != nil {
		logger.WithField("error", err).Error("Error verifying fleet manager config")
	}

	return cfg
}

func (cfg *FleetManagerConfig) VerifyFleetManagerConfig() error {

	// Verify the Fleet Manager API URL
	if cfg.FleetManagerApiUrl == "" {
		return fmt.Errorf("fleet_manager_api_url is not set")
	}

	// Verify Fleet Manager API Token
	if cfg.FleetManagerApiToken == "" {
		return fmt.Errorf("fleet_manager_api_token is not set")
	}

	// Verify the Application ID
	if cfg.ApplicationId == "" {
		return fmt.Errorf("application_id is not set")
	}

	// Verify the Max Players
	if cfg.MaxPlayers == 0 {
		return fmt.Errorf("max_players is not set")
	}

	return nil
}
