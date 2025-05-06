package fleetmanager

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	config "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/fleetmanager_config"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/clients"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/storage"
)

type I3dFleetManager struct {
	ctx             context.Context
	client          clients.ApplicationInstance
	logger          runtime.Logger
	nk              runtime.NakamaModule
	cfg             *config.Config
	callbackHandler runtime.FmCallbackHandler
	storage         storage.FleetManagerStorage
}

// NewI3dFleetManager creates a new I3dFleetManager instance from the given context, logger, initializer, nakama module, and configuration
func NewI3dFleetManager(
	ctx context.Context,
	logger runtime.Logger,
	initializer runtime.Initializer,
	nk runtime.NakamaModule,
	cfg *config.Config,
) (runtime.FleetManagerInitializer, error) {

	client := clients.NewOneApiClient(cfg, clients.NewAuthentication(cfg), logger)

	storageService, err := storage.NewFleetManagerStorageService(nk, initializer, logger)

	if err != nil {
		return nil, err
	}

	fm := &I3dFleetManager{
		ctx:     ctx,
		client:  client,
		logger:  logger,
		nk:      nk,
		cfg:     cfg,
		storage: storageService,
	}

	// NOTE: This RPC is required and must be invoked from the Game Session whenever a player connects/disconnects.
	if err = initializer.RegisterRpc(RpcIdUpdateInstanceInfo, fm.UpdateInstanceInfo); err != nil {
		return nil, err
	}

	// NOTE: This RPC is required and must be invoked from the Game Session whenever a Game Session terminates.
	if err = initializer.RegisterRpc(RpcIdDeleteInstanceInfo, fm.DeleteInstanceInfo); err != nil {
		return nil, err
	}

	return fm, nil
}

// Init initializes the I3dFleetManager instance
func (fm *I3dFleetManager) Init(nk runtime.NakamaModule, callbackHandler runtime.FmCallbackHandler) error {
	fm.logger.WithField("method_name", "Init").Debug("FleetManager - Entered Init Method")

	// do small api test
	foundInstances, err := fm.client.ListApplicationInstances(fm.ctx, "", 1, "")

	if err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to list instances")
	}

	fm.logger.WithField("found_instances", foundInstances).Debug("found instances")
	fm.nk = nk
	fm.callbackHandler = callbackHandler

	return nil
}

// Get retrieves an instance from the Fleet Manager API
func (fm *I3dFleetManager) Get(ctx context.Context, id string) (instance *runtime.InstanceInfo, err error) {
	fm.logger.WithField("method_name", "Get").Debug("FleetManager - Entered Get Method")
	instance, err = fm.client.GetApplicationInstance(ctx, id)

	if err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to get instance")
		return nil, err
	}

	// we only keep instances that are on status 5 (allocated) in the Nakama storage
	if instance.Status != clients.ApplicationInstanceStatus[5] {
		err = fm.storage.DeleteStorageGameSession(ctx, []string{instance.Id})
		if err != nil {
			return nil, err
		}
	} else {
		err = fm.storage.UpdateStorageGameSession(ctx, []*runtime.InstanceInfo{instance})
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

// List lists all instances from the Fleet Manager API
func (fm *I3dFleetManager) List(ctx context.Context, query string, limit int, previousCursor string) (list []*runtime.InstanceInfo, nextCursor string, err error) {
	fm.logger.WithField("method_name", "List").Debug("FleetManager - Entered List Method")

	if query == "" {
		fb := NewFilterBuilder()
		query = fb.Add(status, "5").Add(applicationId, fm.cfg.ApplicationId).Query()

		result, err := fm.client.ListApplicationInstances(ctx, query, limit, previousCursor)
		if err != nil {
			fm.logger.WithField("error", err.Error()).Error("failed to list instances")
			return nil, "", err
		}

		err = fm.storage.UpdateStorageGameSession(ctx, result.Instances)
		return result.Instances, result.NextCursor, nil
	}

	results, err := fm.storage.ListGameSessionsFromStorage(ctx, query, limit, []string{"player_count", "-create_time"}, "")
	if err != nil {
		return nil, "", err
	}
	return results, nextCursor, nil
}

// Create creates a new instance in the Fleet Manager API
func (fm *I3dFleetManager) Create(ctx context.Context, maxPlayers int, userIds []string, latencies []runtime.FleetUserLatencies, metadata map[string]any, callback runtime.FmCreateCallbackFn) error {
	fm.logger.WithField("method_name", "Create").Info("FleetManager - Entered Create Method")

	id := fm.callbackHandler.GenerateCallbackId()

	if callback != nil {
		fm.callbackHandler.SetCallback(id, callback)
	}

	go func(ctx context.Context, callbackId string, callBackHandler runtime.FmCallbackHandler) {
		instance, err := fm.client.AllocateApplicationInstance(ctx, metadata, GetFilters(metadata))
		if err != nil {
			go callBackHandler.InvokeCallback(
				callbackId,
				runtime.CreateError,
				nil, nil,
				metadata,
				err,
			)
			return
		}

		// updating the instance with the concern of Nakama
		instance.PlayerCount = len(userIds)
		instance.Metadata[MaxPlayers] = maxPlayers

		var sessionInfo = make([]*runtime.SessionInfo, 0, len(userIds))

		for _, userId := range userIds {
			sessionInfo = append(sessionInfo, &runtime.SessionInfo{UserId: userId})
		}

		go callBackHandler.InvokeCallback(
			id,
			runtime.CreateSuccess,
			instance,
			sessionInfo,
			metadata,
			nil,
		)

		//  updating the storage
		if err = fm.storage.UpdateStorageGameSession(ctx, []*runtime.InstanceInfo{instance}); err != nil {
			fm.logger.WithField("error", err).Error("error writing to Nakama storage after starting session")
		}

	}(ctx, id, fm.callbackHandler)

	return nil
}

// The Join method is there to allow the Fleet Manager to join a user to an instance but this is not implemented by i3d.
func (fm *I3dFleetManager) Join(ctx context.Context, id string, userIds []string, metadata map[string]string) (joinInfo *runtime.JoinInfo, err error) {
	instance, err := fm.storage.GetGameSessionFromStorage(ctx, id)
	if err != nil {
		return nil, err
	}

	sessionInfo := make([]*runtime.SessionInfo, 0, len(userIds))
	maxPlayers, err := getMaxPlayers(instance)
	if err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to get max players")
		return nil, err
	}

	currentPlayers := instance.PlayerCount
	for _, userId := range userIds {

		// when there are more players that want to join than available slots
		// we will break the join, this is on the implementation of the plugin to handle this
		if currentPlayers >= maxPlayers {
			break
		}

		sessionInfo = append(sessionInfo, &runtime.SessionInfo{
			UserId: userId,
		})

		currentPlayers++
	}

	instance.PlayerCount = currentPlayers
	if err = fm.storage.UpdateStorageGameSession(ctx, []*runtime.InstanceInfo{instance}); err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to update storage")
		return nil, err
	}

	return &runtime.JoinInfo{
		InstanceInfo: instance,
		SessionInfo:  sessionInfo,
	}, nil
}

// UpdateInstanceInfo updates the instance in the Fleet Manager API
func (fm *I3dFleetManager) UpdateInstanceInfo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	fm.logger.WithField("method_name", "UpdateInstanceInfo").Debug("FleetManager - Entered UpdateInstanceInfo Method")

	fm.logger.WithField("payload", payload).Debug("received update from headless instance")

	request, err := FromPayloadToRequest[UpdateInstanceInfoRequest](payload)

	if err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal updateInstanceInfo request")
		return "", ErrInternalError
	}

	if err := fm.Update(ctx, request.Id, request.PlayerCount, request.Metadata); err != nil {
		logger.WithField("error", err.Error()).Error("failed to update instance info")
		return "", ErrInternalError
	}

	return "", nil
}

// DeleteInstanceInfo deletes the instance in the Fleet Manager API
func (fm *I3dFleetManager) DeleteInstanceInfo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	fm.logger.WithField("method_name", "DeleteInstanceInfo").Debug("FleetManager - Entered DeleteInstanceInfo Method")

	request, err := FromPayloadToRequest[DeleteInstanceInfoRequest](payload)
	if err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal deleteInstanceInfo request")
		return "", ErrInternalError
	}

	fm.logger.WithField("instance_id", request.Id).Debug("received delete from headless instance")

	if err := fm.Delete(ctx, request.Id); err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to delete instance info")
		return "", ErrInternalError
	}

	return "", nil
}

// Update updates the instance in the Fleet Manager API
func (fm *I3dFleetManager) Update(ctx context.Context, id string, playerCount int, metadata map[string]any) error {
	fm.logger.WithField("method_name", "Update").Debug("FleetManager - Entered Update Method")
	fm.logger.WithField("instance_id", id).Debug("processing update on api")

	instance, err := fm.client.UpdateApplicationInstance(ctx, id, metadata)
	if err != nil {
		return err
	}

	// updating the instance with the concern of Nakama
	instance.PlayerCount = playerCount

	//  updating the storage
	err = fm.storage.UpdateStorageGameSession(ctx, []*runtime.InstanceInfo{instance})

	if err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to update storage")
		return err
	}

	return nil
}

// Delete deletes the instance in the Fleet Manager API
func (fm *I3dFleetManager) Delete(ctx context.Context, id string) error {
	fm.logger.WithField("method_name", "Delete").Debug("FleetManager - Entered Delete Method")
	fm.logger.WithField("instance_id", id).Debug("processing delete on api")

	// because deleting the instance is slower than restarting it
	err := fm.client.RestartApplicationInstance(ctx, id)
	if err != nil {
		return err
	}

	err = fm.storage.DeleteStorageGameSession(ctx, []string{id})
	if err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to delete storage")
		return err
	}

	return nil
}
