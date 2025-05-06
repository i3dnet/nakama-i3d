// This is the example package that uses the I3d Nakama plugin to create a fleet manager.
package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/config"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/fleetmanager"

	"time"
)

var localconfig *config.Config

// InitModule is the entry point for the module. and creates the fleet manager with the given configuration.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	//Create the config
	var runTimeError *runtime.Error

	// getting config from the local.yml inside the runtime environment variables
	localconfig, runTimeError = config.NewConfigFromRuntime(ctx)
	if runTimeError != nil {

		// fallback to configuration from .env file
		// or set as environment variables on the operating system
		// or as settings.json set in the app directory
		localconfig, runTimeError = config.NewConfig()
		if runTimeError != nil {
			logger.WithField("error", runTimeError).Error("failed to create config")
			return runTimeError
		}
	}

	//Create the fleet manager
	fm, err := fleetmanager.NewI3dFleetManager(ctx, logger, initializer, nk, localconfig)
	if err != nil {
		return err
	}

	//Register the fleet manager with Nakama
	if err = initializer.RegisterFleetManager(fm); err != nil {
		logger.WithField("error", err).Error("failed to register fleet manager")
		return err
	}

	//Register the healthcheck RPC
	err = initializer.RegisterRpc("healthcheck", fleetmanager.RpcHealthCheck)
	if err != nil {
		logger.Error("Error registering healthcheck RPC: %v", err)
		return err
	}

	//Register the matchmaker matched function
	if err = initializer.RegisterMatchmakerMatched(MatchmakerMatched); err != nil {
		logger.Error("Error registering MatchmakerMatched: %v", err)
		return err
	}

	logger.Debug("Module loaded in %dms", time.Since(initStart).Milliseconds())
	return nil
}

// MatchmakerMatched is called when a match is made by Nakama due to the registration of this method with the initializer in the InitModule function.
func MatchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	logger.Debug("MatchmakerMatched called")

	// Get the passed in properties of the first entry
	properties := entries[0].GetProperties()
	logger.Debug("Properties: %v", properties)

	// Get the user ids from the matchmaker entries
	userIds := make([]string, 0, len(entries))
	for _, entry := range entries {
		userIds = append(userIds, entry.GetPresence().GetUserId())
	}

	// Retrieve the fleet manager
	fm := nk.GetFleetManager()

	// Attach a callback function to the Fleet Manager Create
	var callback runtime.FmCreateCallbackFn = func(status runtime.FmCreateStatus, instanceInfo *runtime.InstanceInfo, sessionInfo []*runtime.SessionInfo, metadata map[string]any, err error) {
		logger.Debug("Callback called")

		if err != nil {
			logger.Error("Error creating game session: %v", err)
			return
		}

		for _, userId := range userIds {
			// Use the Nakama Instance to Notify each user that the game session has been created and supply IpAddress
			sessionId, found := getSessionForUserId(sessionInfo, userId)
			if found {

				sessionInfo := map[string]any{
					"IpAddress": instanceInfo.ConnectionInfo.IpAddress,
					"Port":      instanceInfo.ConnectionInfo.Port,
					"SessionId": sessionId,
				}

				err = nk.NotificationsSend(ctx, []*runtime.NotificationSend{
					{
						Code:    9000,
						UserID:  userId,
						Subject: "Game Session Created",
						Content: sessionInfo,
					},
				})

				if err != nil {
					logger.WithField("error", err).Error("error sending notification")
				}
			}
		}

		logger.Debug("Callback successfully called with status: %v", status)
	}

	metadata := make(map[string]any)

	// Add any properties passed in from the matchmaker
	if properties["duration"] != "" {
		metadata["duration"] = 1800
	}

	logger.Debug("Metadata: %v", metadata)

	maxPlayers := 2

	// Create the game session
	err := fm.Create(ctx, maxPlayers, userIds, nil, metadata, callback)
	if err != nil {
		logger.Error("Error creating game session: %v", err)
		return "", err
	}

	return "", nil
}

// getSessionForUserId returns the session id for the given user id or an empty string if the user id is not found
func getSessionForUserId(sessions []*runtime.SessionInfo, userId string) (string, bool) {
	for _, session := range sessions {
		if session.UserId == userId {
			return session.SessionId, true
		}
	}

	return "", false
}
