package fleetmanager

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

type FleetManager_i3d struct {
	ctx             context.Context
	i3dService      *FleetManagerApiClient
	logger          runtime.Logger
	nk              runtime.NakamaModule
	cfg             FleetManagerConfig
	callbackHandler runtime.FmCallbackHandler
}

// UpdateInstanceInfoRequest is the data transfer object for the UpdateInstanceInfo RPC
type UpdateInstanceInfoRequest struct {
	Id          string         `json:"id"`
	PlayerCount int            `json:"player_count"`
	Metadata    map[string]any `json:"metadata"`
}

// DeleteInstanceInfoRequest is the data transfer object for the DeleteInstanceInfo RPC
type DeleteInstanceInfoRequest struct {
	Id string `json:"id"`
}

// Error codes for the Fleet Manager
var (
	ErrInvalidInput  = runtime.NewError("input is invalid", 3)       // INVALID_ARGUMENT
	ErrInternalError = runtime.NewError("internal server error", 13) // INTERNAL
)

// RPC IDs for the Fleet Manager
const (
	RpcIdUpdateInstanceInfo = "update_instance_info"
	RpcIdDeleteInstanceInfo = "delete_instance_info"
)

// NewFleetManager_i3d creates a new FleetManager_i3d instance from the given context, logger, initializer, nakama module, and configuration
func NewFleetManager_i3d(ctx context.Context, logger runtime.Logger, initializer runtime.Initializer, nk runtime.NakamaModule, conf FleetManagerConfig) (runtime.FleetManagerInitializer, error) {

	apiService := NewFleetManagerApiClient(conf.FleetManagerApiToken, conf.FleetManagerApiUrl, conf.ApplicationId, logger)

	fm := &FleetManager_i3d{
		ctx:        ctx,
		i3dService: apiService,
		logger:     logger,
		nk:         nk,
		cfg:        conf,
	}

	// NOTE: This RPC is required and must be invoked from the Game Session whenever a player connects/disconnects.
	if err := initializer.RegisterRpc(RpcIdUpdateInstanceInfo, fm.UpdateInstanceInfo); err != nil {
		return nil, err
	}

	// NOTE: This RPC is required and must be invoked from the Game Session whenever a Game Session terminates.
	if err := initializer.RegisterRpc(RpcIdDeleteInstanceInfo, fm.DeleteInstanceInfo); err != nil {
		return nil, err
	}

	return fm, nil
}

// Init initializes the FleetManager_i3d instance
func (fm *FleetManager_i3d) Init(nk runtime.NakamaModule, callbackHandler runtime.FmCallbackHandler) error {
	fm.logger.WithField("method_name", "Init").Debug("FleetManager - Entered Init Method")

	fm.nk = nk
	fm.callbackHandler = callbackHandler

	return nil
}

// Get retrieves an instance from the Fleet Manager API
func (fm *FleetManager_i3d) Get(ctx context.Context, id string) (instance *runtime.InstanceInfo, err error) {
	fm.logger.WithField("method_name", "Get").Debug("FleetManager - Entered Get Method")

	fm.ctx = ctx
	instance, err = fm.i3dService.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// List lists all instances from the Fleet Manager API
func (fm *FleetManager_i3d) List(ctx context.Context, query string, limit int, previousCursor string) (list []*runtime.InstanceInfo, nextCursor string, err error) {
	fm.logger.WithField("method_name", "List").Debug("FleetManager - Entered List Method")

	fm.ctx = ctx
	if len(fm.i3dService.metadata.values) > 0 {
		fm.logger.Debug("Metadata is not empty")
	}

	return fm.i3dService.List(ctx, query, limit, previousCursor)
}

// Create creates a new instance in the Fleet Manager API
func (fm *FleetManager_i3d) Create(ctx context.Context, maxPlayers int, userIds []string, latencies []runtime.FleetUserLatencies, metadata map[string]any, callback runtime.FmCreateCallbackFn) error {
	fm.logger.WithField("method_name", "Create").Debug("FleetManager - Entered Create Method")

	fm.ctx = ctx
	instance, err := fm.i3dService.Create(ctx, maxPlayers, userIds, latencies, metadata, callback)
	if err != nil {
		return err
	}

	if callback != nil {
		fm.logger.Debug("FM - Calling callback")
		callback(runtime.CreateSuccess, instance, nil, metadata, nil)
	}

	return nil
}

// The Join method is there to allow the Fleet Manager to join a user to an instance but this is not implemented by i3d.
func (fm *FleetManager_i3d) Join(ctx context.Context, id string, userIds []string, metadata map[string]string) (joinInfo *runtime.JoinInfo, err error) {
	return nil, runtime.NewError("Not implemented", 501)
}

// UpdateInstanceInfo updates the instance in the Fleet Manager API
func (fm *FleetManager_i3d) UpdateInstanceInfo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	fm.logger.WithField("method_name", "UpdateInstanceInfo").Debug("FleetManager - Entered UpdateInstanceInfo Method")

	fm.ctx = ctx
	var req *UpdateInstanceInfoRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal updateInstanceInfo request")
		return "", ErrInternalError
	}

	if req.Id == "" {
		return "", ErrInvalidInput
	}

	if err := fm.Update(ctx, req.Id, req.PlayerCount, req.Metadata); err != nil {
		logger.WithField("error", err.Error()).Error("failed to update instance info")
		return "", ErrInternalError
	}

	return "", nil
}

// DeleteInstanceInfo deletes the instance in the Fleet Manager API
func (fm *FleetManager_i3d) DeleteInstanceInfo(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	fm.logger.WithField("method_name", "DeleteInstanceInfo").Debug("FleetManager - Entered DeleteInstanceInfo Method")

	fm.ctx = ctx
	var req *DeleteInstanceInfoRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		logger.WithField("error", err.Error()).Error("failed to unmarshal deleteInstanceInfo request")
		return "", ErrInternalError
	}

	if len(req.Id) < 1 {
		return "", ErrInvalidInput
	}

	fm.logger.WithField("instance_id", req.Id).Debug("received delete from headless instance")

	if err := fm.Delete(ctx, req.Id); err != nil {
		fm.logger.WithField("error", err.Error()).Error("failed to delete instance info")
		return "", ErrInternalError
	}

	return "", nil
}

// Update updates the instance in the Fleet Manager API
func (fm *FleetManager_i3d) Update(ctx context.Context, id string, playerCount int, metadata map[string]any) error {
	fm.logger.WithField("method_name", "Update").Debug("FleetManager - Entered Update Method")
	fm.ctx = ctx

	fm.logger.WithField("instance_id", id).Debug("processing update on api")

	err := fm.i3dService.Update(ctx, id, playerCount, metadata)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes the instance in the Fleet Manager API
func (fm *FleetManager_i3d) Delete(ctx context.Context, id string) error {

	fm.logger.WithField("method_name", "Delete").Debug("FleetManager - Entered Delete Method")
	fm.ctx = ctx

	fm.logger.WithField("instance_id", id).Debug("processing delete on api")

	err := fm.i3dService.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
