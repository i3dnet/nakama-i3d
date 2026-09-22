package clients

import (
	"context"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	openapi "github.com/i3dnet/nakama-i3d/internal/openapi"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

const I3dFilters = "i3dFilters"
const ApplicationId = "overwriteApplicationId"
const DefaultLimit = 100

type ApplicationInstanceListResponse struct {
	NextCursor string
	Instances  []*runtime.InstanceInfo
}

type ApplicationInstance interface {
	ListApplicationInstances(ctx context.Context, filters string, limit int, previousCursor string) (*ApplicationInstanceListResponse, error)
	GetApplicationInstance(ctx context.Context, instanceID string) (*runtime.InstanceInfo, error)
	AllocateApplicationInstance(ctx context.Context, metaData map[string]any, filters string) (*runtime.InstanceInfo, error)
	RestartApplicationInstance(ctx context.Context, instanceID string) error
	UpdateApplicationInstance(ctx context.Context, instanceID string, metaData map[string]any) (*runtime.InstanceInfo, error)
}

var ApplicationInstanceStatus = map[int32]string{
	0:   "INACTIVE",
	1:   "SETUP",
	2:   "OFFLINE",
	3:   "STARTING",
	4:   "ONLINE",
	5:   "ALLOCATED",
	6:   "ALLOCATING",
	101: "DEPENDENCY_SETUP",
	102: "DEPENDENCY_OFFLINE",
	103: "DEPENDENCY_STARTING",
	104: "DEPENDENCY_RUNNING",
	105: "DEPENDENCY_COMPLETED",
	106: "DEPENDENCY_FAILED",
}

func (o *OneApiClient) ListApplicationInstances(ctx context.Context, filters string, limit int, previousCursor string) (*ApplicationInstanceListResponse, error) {
	client, err := o.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	request := client.ApplicationInstanceAPI.GetApplicationInstances(ctx)
	if filters != "" {
		request = request.Filters(filters)
	}

	request = request.RANGEDDATA(createRangedData(limit))
	request = request.PAGETOKEN(previousCursor)

	applicationInstances, response, err := executeRead(ctx, o, request.Execute)
	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to get Instances")
		return nil, err
	}
	var instances []*runtime.InstanceInfo
	for _, instance := range applicationInstances {
		ins, err := o.mapToInstanceInfo(instance)
		if err != nil {
			o.logger.WithField("error", err.Error()).Error("failed to map instance")
			return nil, fmt.Errorf("invalid instance in provider page: %w", err)
		}

		instances = append(instances, ins)
	}

	return &ApplicationInstanceListResponse{
		NextCursor: response.Header.Get("PAGE-TOKEN"),
		Instances:  instances,
	}, nil
}

func (o *OneApiClient) GetApplicationInstance(ctx context.Context, instanceID string) (*runtime.InstanceInfo, error) {
	client, err := o.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	request := client.ApplicationInstanceAPI.GetApplicationInstance(ctx, instanceID)
	response, _, err := executeRead(ctx, o, request.Execute)
	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to get instance")
		return nil, err
	}

	if len(response) != 1 {
		return nil, fmt.Errorf("expected one instance, got %d", len(response))
	}
	instanceInfo, err := o.mapToInstanceInfo(response[0])
	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to map instance")
		return nil, err
	}

	return instanceInfo, nil
}

func (o *OneApiClient) getApplicationId(metaData map[string]any) string {
	if applicationId, ok := metaData[ApplicationId]; ok {
		return fmt.Sprintf("%v", applicationId)
	}

	return o.cfg.ApplicationId
}

func (o *OneApiClient) AllocateApplicationInstance(ctx context.Context, metaData map[string]any, filters string) (*runtime.InstanceInfo, error) {
	client, err := o.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	request := client.ApplicationInstanceAPI.UpdateApplicationInstanceGameEmptyAllocate(ctx, o.getApplicationId(metaData))
	if filters != "" {
		request = request.Filters(filters)
	}

	request = request.MetadataCollection(createMetaData(metaData))

	// Allocation is not idempotent: an ambiguous failure may already have allocated.
	response, _, err := request.Execute()

	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to allocate instance")
		return nil, err
	}

	if len(response) != 1 {
		return nil, fmt.Errorf("expected one instance, got %d", len(response))
	}
	instanceInfo, err := o.mapToInstanceInfo(response[0])
	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to map instance")
		return nil, err
	}

	if instanceInfo.Status != ApplicationInstanceStatus[5] {
		return nil, fmt.Errorf("allocation is not complete: status %s", instanceInfo.Status)
	}
	return instanceInfo, nil
}

func (o *OneApiClient) RestartApplicationInstance(ctx context.Context, instanceID string) error {
	client, err := o.GetClient(ctx)
	if err != nil {
		return err
	}
	_, _, err = client.ApplicationInstanceAPI.CreateApplicationInstanceRestart(ctx, instanceID).Execute()
	return err
}

func (o *OneApiClient) UpdateApplicationInstance(ctx context.Context, instanceID string, metaData map[string]any) (*runtime.InstanceInfo, error) {
	client, err := o.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	appInstances, _, err := executeRead(ctx, o, client.ApplicationInstanceAPI.GetApplicationInstance(ctx, instanceID).Execute)

	if err != nil {
		return nil, err
	}
	if len(appInstances) != 1 {
		return nil, fmt.Errorf("expected one instance, got %d", len(appInstances))
	}
	appInstance := appInstances[0]
	if appInstance.Id != instanceID {
		return nil, fmt.Errorf("provider returned a different instance")
	}
	appInstance.Metadata = createMetaData(metaData).Metadata

	request := client.ApplicationInstanceAPI.UpdateApplicationInstance(ctx, instanceID).ApplicationInstance(appInstance)
	updated, _, err := request.Execute()

	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to update instance")
		return nil, err
	}

	if len(updated) != 1 || updated[0].Id != instanceID {
		return nil, fmt.Errorf("provider returned an invalid update response")
	}
	return o.mapToInstanceInfo(updated[0])
}

func (o *OneApiClient) mapToInstanceInfo(instance openapi.ApplicationInstance) (*runtime.InstanceInfo, error) {

	if strings.TrimSpace(instance.Id) == "" {
		return nil, fmt.Errorf("instance ID is missing")
	}
	var connectionInfo *runtime.ConnectionInfo
	for _, ip := range instance.IpAddress {
		address := net.ParseIP(ip.IpAddress)
		if ip.Private == 0 && address != nil && !address.IsUnspecified() && !address.IsMulticast() {
			connectionInfo = &runtime.ConnectionInfo{
				IpAddress: ip.IpAddress,
			}
			break
		}
	}
	if connectionInfo == nil {
		return nil, fmt.Errorf("instance has no usable public IP address")
	}
	port, err := parsePort(instance.Properties)
	if err != nil {
		o.logger.WithField("error", err.Error()).Error("failed to parse port")
		return nil, fmt.Errorf("failed to parse port: %w", err)
	}
	connectionInfo.Port = port

	return &runtime.InstanceInfo{
		Id:             instance.Id,
		ConnectionInfo: connectionInfo,
		CreateTime:     parseTimestamp(instance.CreatedAt),
		PlayerCount:    int(instance.NumPlayers),
		Status:         getStatus(instance.Status),
		Metadata:       parseMetadata(instance.Metadata),
	}, nil
}

func findProperty(property []openapi.ApplicationInstanceProperty, key string) *openapi.ApplicationInstanceProperty {
	for _, p := range property {
		if p.PropertyKey == key {
			return &p
		}
	}
	return nil
}

func parseTimestamp(timestamp int32) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func parsePort(property []openapi.ApplicationInstanceProperty) (int, error) {
	port := findProperty(property, "port")
	if port == nil {
		return 0, fmt.Errorf("port not found")
	}

	value, err := strconv.Atoi(port.PropertyValue)
	if err != nil {
		return 0, fmt.Errorf("port is not a number")
	}
	if value < 1 || value > 65535 {
		return 0, fmt.Errorf("port is outside 1-65535")
	}
	return value, nil
}

func getStatus(statusInt int32) string {
	if status, ok := ApplicationInstanceStatus[statusInt]; ok {
		return status
	}
	return "UNKNOWN"
}

func parseMetadata(metadata []openapi.Metadata) map[string]any {
	parsedMetadata := make(map[string]any)
	for _, m := range metadata {
		parsedMetadata[m.Key] = m.Value
	}
	return parsedMetadata
}

func createMetaData(metaData map[string]any) openapi.MetadataCollection {
	parsedMetaData := make([]openapi.Metadata, 0)
	keys := make([]string, 0, len(metaData))
	for key := range metaData {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := metaData[key]
		if key == ApplicationId || key == I3dFilters || strings.HasPrefix(key, "i3d_") {
			continue
		}
		parsedMetaData = append(parsedMetaData, openapi.Metadata{
			Key:   key,
			Value: fmt.Sprintf("%v", value),
		})
	}
	return openapi.MetadataCollection{Metadata: parsedMetaData}
}

func createRangedData(limit int) string {
	if limit <= 0 || limit > DefaultLimit {
		limit = DefaultLimit
	}

	return fmt.Sprintf("results=%d", limit)
}
