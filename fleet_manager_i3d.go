package fleetmanager

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	config "github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/i3dnet/nakama-i3d/internal/storage"
	"strings"
	"sync"
	"time"
)

var _ runtime.FleetManagerInitializer = (*I3dFleetManager)(nil)

type I3dFleetManager struct {
	ctx             context.Context
	cancel          context.CancelFunc
	lifecycleMu     sync.Mutex
	stopping        bool
	operations      sync.WaitGroup
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

	if cfg == nil {
		return nil, ErrInvalidInput
	}
	lifetime, cancel := context.WithCancel(context.WithoutCancel(ctx))
	success := false
	defer func() {
		if !success {
			cancel()
		}
	}()
	client := clients.NewOneApiClient(cfg, clients.NewAuthentication(cfg), logger)

	storageService, err := storage.NewFleetManagerStorageService(nk, initializer, logger)

	if err != nil {
		return nil, err
	}

	fm := &I3dFleetManager{
		ctx:     lifetime,
		cancel:  cancel,
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

	if err = initializer.RegisterShutdown(func(ctx context.Context, _ runtime.Logger, _ *sql.DB, _ runtime.NakamaModule) { fm.shutdown(ctx) }); err != nil {
		return nil, err
	}
	success = true
	return fm, nil
}

// Init initializes the I3dFleetManager instance
func (fm *I3dFleetManager) Init(nk runtime.NakamaModule, callbackHandler runtime.FmCallbackHandler) error {
	fm.logger.WithField("method_name", "Init").Debug("FleetManager - Entered Init Method")

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
		query = fb.Add(applicationId, fm.cfg.ApplicationId).Query()

		result, err := fm.client.ListApplicationInstances(ctx, query, limit, previousCursor)
		if err != nil {
			fm.logger.WithField("error", err.Error()).Error("failed to list instances")
			return nil, "", err
		}

		// One API does not support filtering status. Keep its page cursor while
		// selecting allocated entries locally; a page may contain no matches.
		allocated := make([]*runtime.InstanceInfo, 0, len(result.Instances))
		for _, instance := range result.Instances {
			if instance.Status == clients.ApplicationInstanceStatus[5] {
				allocated = append(allocated, instance)
			}
		}
		if len(allocated) > 0 {
			if err = fm.storage.UpdateStorageGameSession(ctx, allocated); err != nil {
				return nil, "", err
			}
		}
		return allocated, result.NextCursor, nil
	}

	results, err := fm.storage.ListGameSessionsFromStorage(ctx, query, limit, []string{"player_count", "-create_time"}, "")
	if err != nil {
		return nil, "", err
	}
	return results, nextCursor, nil
}

// Create creates a new instance in the Fleet Manager API
func (fm *I3dFleetManager) Create(ctx context.Context, maxPlayers int, userIds []string, latencies []runtime.FleetUserLatencies, metadata map[string]any, callback runtime.FmCreateCallbackFn) (map[string]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if maxPlayers <= 0 || len(userIds) > maxPlayers {
		return nil, ErrInvalidInput
	}
	seen := make(map[string]struct{}, len(userIds))
	for _, id := range userIds {
		if _, duplicate := seen[id]; id == "" || duplicate {
			return nil, ErrInvalidInput
		}
		seen[id] = struct{}{}
	}
	// Own the accepted request; callers may reuse their maps/slices after returning.
	encoded, err := json.Marshal(metadata)
	if err != nil {
		return nil, ErrInvalidInput
	}
	var requestMetadata map[string]any
	if err = json.Unmarshal(encoded, &requestMetadata); err != nil {
		return nil, ErrInvalidInput
	}
	userIds = append([]string(nil), userIds...)

	fm.lifecycleMu.Lock()
	if fm.stopping || fm.ctx.Err() != nil || fm.callbackHandler == nil {
		fm.lifecycleMu.Unlock()
		return nil, runtime.NewError("fleet manager is not running", UNAVAILABLE)
	}
	timeout := fm.cfg.AllocationTimeout
	if timeout == 0 {
		timeout = 120 * time.Second
	}
	operationCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), timeout)
	stopShutdown := context.AfterFunc(fm.ctx, cancel)
	id := ""
	if callback != nil {
		id = fm.callbackHandler.GenerateCallbackId()
		fm.callbackHandler.SetCallback(id, callback)
	}
	fm.operations.Add(2)
	fm.lifecycleMu.Unlock()

	type outcome struct {
		instance *runtime.InstanceInfo
		sessions []*runtime.SessionInfo
		err      error
	}
	result := make(chan outcome, 1)
	go func() {
		defer fm.operations.Done()
		instance, err := fm.client.AllocateApplicationInstance(operationCtx, requestMetadata, GetFilters(requestMetadata))
		if err == nil && instance == nil {
			err = errors.New("allocation returned no instance")
		}
		var sessions []*runtime.SessionInfo
		if err == nil {
			if instance.Metadata == nil {
				instance.Metadata = make(map[string]any)
			}
			instance.PlayerCount = len(userIds)
			instance.Metadata[MaxPlayers] = maxPlayers
			for _, userId := range userIds {
				sessions = append(sessions, &runtime.SessionInfo{UserId: userId})
			}
			if err = operationCtx.Err(); err == nil {
				err = fm.storage.UpdateStorageGameSession(operationCtx, []*runtime.InstanceInfo{instance})
			}
		}
		result <- outcome{instance, sessions, err}
	}()
	go func() {
		defer fm.operations.Done()
		defer cancel()
		defer stopShutdown()
		var completed outcome
		select {
		case completed = <-result:
			if operationCtx.Err() != nil {
				completed.err = operationCtx.Err()
			}
		case <-operationCtx.Done():
			completed.err = operationCtx.Err()
		}
		if callback == nil {
			if completed.err != nil {
				fm.logger.Error("allocation failed: %v", completed.err)
			}
			return
		}
		if completed.err != nil {
			fm.callbackHandler.InvokeCallback(id, runtime.CreateError, nil, nil, nil, completed.err)
			return
		}
		fm.callbackHandler.InvokeCallback(id, runtime.CreateSuccess, completed.instance, completed.sessions, requestMetadata, nil)
	}()
	return nil, nil
}

// shutdown cancels accepted operations and waits within Nakama's grace period.
func (fm *I3dFleetManager) shutdown(ctx context.Context) {
	fm.lifecycleMu.Lock()
	fm.stopping = true
	fm.cancel()
	fm.lifecycleMu.Unlock()
	done := make(chan struct{})
	go func() { fm.operations.Wait(); close(done) }()
	select {
	case <-done:
	case <-ctx.Done():
	}
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
	if ctx.Value(runtime.RUNTIME_CTX_USER_ID) != nil {
		return "", runtime.NewError("server authentication required", PERMISSION_DENIED)
	}

	fm.logger.WithField("method_name", "UpdateInstanceInfo").Debug("FleetManager - Entered UpdateInstanceInfo Method")

	request, err := FromPayloadToRequest[UpdateInstanceInfoRequest](payload)

	if err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal updateInstanceInfo request")
		return "", ErrInvalidInput
	}

	if strings.TrimSpace(request.Id) == "" {
		return "", ErrInvalidInput
	}
	if request.PlayerCount < 0 {
		return "", ErrInvalidInput
	}
	for key := range request.Metadata {
		if strings.HasPrefix(key, "i3d_") || key == clients.I3dFilters || key == clients.ApplicationId {
			return "", ErrInvalidInput
		}
	}

	if err := fm.Update(ctx, request.Id, request.PlayerCount, request.Metadata); err != nil {
		logger.WithField("error", err.Error()).Error("failed to update instance info")
		return "", ErrInternalError
	}

	return "", nil
}

// DeleteInstanceInfo deletes the instance in the Fleet Manager API
func (fm *I3dFleetManager) DeleteInstanceInfo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	if ctx.Value(runtime.RUNTIME_CTX_USER_ID) != nil {
		return "", runtime.NewError("server authentication required", PERMISSION_DENIED)
	}

	fm.logger.WithField("method_name", "DeleteInstanceInfo").Debug("FleetManager - Entered DeleteInstanceInfo Method")

	request, err := FromPayloadToRequest[DeleteInstanceInfoRequest](payload)
	if err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal deleteInstanceInfo request")
		return "", ErrInvalidInput
	}

	if strings.TrimSpace(request.Id) == "" {
		return "", ErrInvalidInput
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
