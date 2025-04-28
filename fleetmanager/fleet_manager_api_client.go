package fleetmanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// FleetManagerApiClient is a struct that holds the settings to allow interaction with the Fleet Manager API
type FleetManagerApiClient struct {
	apiKey        string
	baseUrl       string
	applicationId string
	logger        runtime.Logger
	metadata      Metadata
}

// Set the relative endpoints for the Fleet Manager API
const (
	allocateEndpoint = "/v3/applicationInstance/game/%s/empty/allocate"
	instanceEndpoint = "/v3/applicationInstance"
)

// AllowedFilters is the list of values that are allowed to be set as filters by the i3D Fleet Manager API
var AllowedFilters = []string{"applicationbuildid", "applicationbuildname", "applicationid", "applicationname",
	"applicationtype", "arcusavailable", "createdat", "dclocationid", "dclocationname", "deploymentenvironmentid",
	"deploymentenvironmentname", "executable", "fleetid", "fleetname", "hostid", "isvirtual", "regionid",
	"regionname", "startupparams"}

// i3dServerInstance is the struct that represents the data that is returned from the Fleet Manager API in JSON format
type i3dServerInstance struct {
	Id                        string      `json:"id"`
	DeploymentEnvironmentID   string      `json:"deploymentEnvironmentId"`
	DeploymentEnvironmentName string      `json:"deploymentEnvironmentName"`
	FleetID                   string      `json:"fleetId"`
	FleetName                 string      `json:"fleetName"`
	HostID                    int         `json:"hostId"`
	IsVirtual                 int         `json:"isVirtual"`
	ApplicationID             string      `json:"applicationId"`
	ApplicationName           string      `json:"applicationName"`
	ApplicationType           int         `json:"applicationType"`
	ApplicationBuildID        string      `json:"applicationBuildId"`
	ApplicationBuildName      string      `json:"applicationBuildName"`
	DCLocationID              int         `json:"dcLocationId"`
	DCLocationName            string      `json:"dcLocationName"`
	RegionID                  string      `json:"regionId"`
	RegionName                string      `json:"regionName"`
	Status                    int         `json:"status"`
	CreatedAt                 int         `json:"createdAt"`
	StartedAt                 int         `json:"startedAt"`
	StoppedAt                 int         `json:"stoppedAt"`
	PID                       int         `json:"pid"`
	PIDChangedAt              int         `json:"pidChangedAt"`
	StartupParams             string      `json:"startupParams"`
	Executable                string      `json:"executable"`
	ManuallyDeployed          int         `json:"manuallyDeployed"`
	Properties                []Property  `json:"properties"`
	IPAddresses               []IPAddress `json:"ipAddress"`
	LabelReadOnly             []KeyValue  `json:"labelReadOnly"`
	LabelPublic               []KeyValue  `json:"labelPublic"`
	Metadata                  []KeyValue  `json:"metadata"`
	NumPlayersMax             int         `json:"numPlayersMax"`
	NumPlayers                int         `json:"numPlayers"`
	LiveHostName              string      `json:"liveHostName"`
	LiveMap                   string      `json:"liveMap"`
	LiveGameVersion           string      `json:"liveGameVersion"`
	UpdatedAt                 int         `json:"updatedAt"`
	LiveRules                 string      `json:"liveRules"`
	AutoRestart               int         `json:"autoRestart"`
	MarkedForDeletion         int         `json:"markedForDeletion"`
	ArcusAvailable            int         `json:"arcusAvailable"`
}

// OriginalMetadata is the struct that represents the metadata that is passed to the Fleet Manager API
type OriginalMetadata struct {
	Metadata []map[string]any `json:"metadata"`
}

// KeyValueMetadata is the struct that represents the metadata that is returned from the Fleet Manager API
type KeyValueMetadata struct {
	Metadata []KeyValue `json:"metadata"`
}

// KeyValue is the struct that represents a key and value pair
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Property is the struct that represents a typed key value pair with an Id
type Property struct {
	ID            string `json:"id"`
	PropertyType  int    `json:"propertyType"`
	PropertyKey   string `json:"propertyKey"`
	PropertyValue string `json:"propertyValue"`
}

// IPAddress is the struct that represents an IP address returned by the instance
type IPAddress struct {
	IPAddress string `json:"ipAddress"`
	IPVersion int    `json:"ipVersion"`
	Private   int    `json:"private"`
}

// UntypedProperty is the struct that represents an untyped key value pair
type UntypedProperty struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

// CommandResult is the struct that represents the JSON returned from the Fleet Manager API when a command is executed
type CommandResult struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BatchID          int    `json:"batchId"`
	CategoryID       int    `json:"categoryId"`
	EntityID         int    `json:"entityId"`
	CurrentActionIdx int    `json:"currentActionIdx"`
	ExecuteAt        int    `json:"executeAt"`
	LastActivityAt   int    `json:"lastActivityAt"`
	FinishedAt       int    `json:"finishedAt"`
	ResultCode       int    `json:"resultCode"`
	ResultText       string `json:"resultText"`
}

// Metadata is the struct that represents the key value pairs that are passed to the Fleet Manager API
type Metadata struct {
	values map[interface{}]interface{}
}

// NewFleetManagerApiClient creates a new FleetManagerApiClient to interact with the Fleet Manager API
func NewFleetManagerApiClient(apiKey string, baseUrl string, applicationId string, logger runtime.Logger) *FleetManagerApiClient {
	return &FleetManagerApiClient{apiKey: apiKey, baseUrl: baseUrl, applicationId: applicationId, logger: logger}
}

// Get the instance from the Fleet Manager API
func (c *FleetManagerApiClient) Get(ctx context.Context, id string) (instance *runtime.InstanceInfo, err error) {

	// Create the URL for the get endpoint
	url := c.baseUrl + instanceEndpoint + "/" + id

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}

	// Set the content type to JSON
	req.Header.Add("content-type", "application/json")

	// Set the private token
	req.Header.Add("PRIVATE-TOKEN", c.apiKey)

	// Call the get endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error("ApiClient - Error calling Do: ", err)
		return nil, err
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Read the response body
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response to get the instance info
	var instances []i3dServerInstance
	if err := json.Unmarshal(response, &instances); err != nil {
		return nil, err
	}

	// Convert the instance data to an InstanceInfo
	newinstance, err := convertInstanceToInstanceInfo(&instances[0])
	if err != nil {
		return nil, err
	}

	// Return the new instance
	return newinstance, nil
}

// List the instances in the Fleet Manager API
func (c *FleetManagerApiClient) List(ctx context.Context, query string, limit int, previousCursor string) (list []*runtime.InstanceInfo, nextCursor string, err error) {

	encoded := url.QueryEscape(query)

	// Create the URL for the list endpoint
	url := c.baseUrl + instanceEndpoint + "/" + "?filters=" + encoded

	if limit > 0 {
		url += "&limit=" + strconv.Itoa(limit)
	}

	if previousCursor != "" {
		url += "&previousCursor=" + previousCursor
	}

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, "", err
	}

	// Set the content type to JSON
	req.Header.Add("content-type", "application/json")

	// Set the private token
	req.Header.Add("PRIVATE-TOKEN", c.apiKey)

	// Call the allocate endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error("ApiClient - Error calling Do: ", err)
		return nil, "", err
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Read the response body
	reponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	// Parse the response to get the instance info
	var instances []i3dServerInstance
	if err := json.Unmarshal(reponse, &instances); err != nil {
		return nil, "", err
	}

	// Convert the instance data to a slice of InstanceInfos
	foundInstances := make([]*runtime.InstanceInfo, 0)
	for _, instance := range instances {
		foundInstance, err := convertInstanceToInstanceInfo(&instance)
		if err != nil {
			return nil, "", err
		}
		foundInstances = append(foundInstances, foundInstance)
	}

	// Return the instances
	return foundInstances, "", nil
}

// Create a new instance in the Fleet Manager API
func (c *FleetManagerApiClient) Create(ctx context.Context, maxPlayers int, userIds []string, latencies []runtime.FleetUserLatencies, metadata map[string]any, callback runtime.FmCreateCallbackFn) (instance *runtime.InstanceInfo, err error) {

	// Wrap the original metadata in an appropriately named struct
	passedinMetadata := OriginalMetadata{Metadata: []map[string]any{metadata}}

	// Create the URL for the allocate endpoint
	allocateUrl := c.baseUrl

	// Get the application id
	if applicationId, found := getSpecificMetadata(passedinMetadata, "applicationId"); found {
		// Remove the application id from the metadata
		delete(passedinMetadata.Metadata[0], "applicationId")
		allocateUrl += fmt.Sprintf(allocateEndpoint, applicationId)
	}

	if filters, found := getSpecificMetadata(passedinMetadata, "filters"); found {
		// Remove the filters from the metadata
		delete(passedinMetadata.Metadata[0], "filters")
		if ok := verifyFilters(filters); ok {
			allocateUrl += url.QueryEscape("&filters=" + filters)
		}
	}

	// Convert the original metadata to the desired format
	convertedMetadata := convertFormat(passedinMetadata)

	// Convert to JSON
	jsonData, err := json.Marshal(convertedMetadata)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, allocateUrl, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return nil, err
	}

	// Set the content type to JSON
	req.Header.Add("content-type", "application/json")

	// Set the accept header to JSON
	req.Header.Add("accept", "application/json")

	// Set the private token
	req.Header.Add("PRIVATE-TOKEN", c.apiKey)

	// Call the allocate endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error("ApiClient - Error calling Do: ", err)
		return nil, err
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Read the response body
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response to get the instance info
	var instances []i3dServerInstance
	if err := json.Unmarshal(response, &instances); err != nil {
		return nil, err
	}

	// Convert the instance data to an InstanceInfo
	newinstance, err := convertInstanceToInstanceInfo(&instances[0])
	if err != nil {
		return nil, err
	}

	// Return the new instance
	return newinstance, nil
}

func (c *FleetManagerApiClient) Update(ctx context.Context, id string, playerCount int, metadata map[string]any) error {
	// Wrap the original metadata in an appropriately named struct
	passedinMetadata := OriginalMetadata{Metadata: []map[string]any{metadata}}

	// Convert the original metadata to the desired format
	convertedMetadata := convertFormat(passedinMetadata)

	// Convert to JSON
	jsonData, err := json.Marshal(convertedMetadata)
	if err != nil {
		return err
	}

	// Create the URL for the Update endpoint
	url := c.baseUrl + instanceEndpoint + "/" + id

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return err
	}

	// Set the content type to JSON
	req.Header.Add("content-type", "application/json")

	// Set the accept header to JSON
	req.Header.Add("accept", "application/json")

	// Set the private token
	req.Header.Add("PRIVATE-TOKEN", c.apiKey)

	// Call the allocate endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error("ApiClient - Error calling Do: ", err)
		return err
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Return no error
	return nil
}

// Delete the instance from the Fleet Manager API
func (c *FleetManagerApiClient) Delete(ctx context.Context, id string) error {
	// Create the URL for the Update endpoint
	url := c.baseUrl + instanceEndpoint + "/" + id

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	// Set the content type to JSON
	req.Header.Add("content-type", "application/json")

	// Set the accept header to JSON
	req.Header.Add("accept", "application/json")

	// Set the private token
	req.Header.Add("PRIVATE-TOKEN", c.apiKey)

	// Call the allocate endpoint
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Error("ApiClient - Error calling Do: ", err)
		return err
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Return no error
	return nil
}

// Conversion of the information returned from the Fleet Manager API to the format required by Nakama
func convertInstanceToInstanceInfo(instance *i3dServerInstance) (result *runtime.InstanceInfo, err error) {

	returnedport, ok := findPropertyValue(instance.Properties, "port")
	if !ok {
		return nil, fmt.Errorf("port not found in instance info")
	}

	port, err := strconv.Atoi(returnedport)
	if err != nil {
		return nil, fmt.Errorf("invalid port returned in instance info")
	}

	result = &runtime.InstanceInfo{
		Id:          instance.Id,
		CreateTime:  time.Unix(int64(instance.CreatedAt), 0),
		PlayerCount: int(instance.NumPlayers),
		Status:      strconv.Itoa(instance.Status),
		Metadata:    convertMetadata(instance.Metadata),
		ConnectionInfo: &runtime.ConnectionInfo{
			IpAddress: instance.IPAddresses[0].IPAddress,
			DnsName:   instance.IPAddresses[0].IPAddress,
			Port:      port,
		},
	}

	if len(instance.IPAddresses) > 1 {
		additionalAddresses := instance.IPAddresses[1:]
		for i, address := range additionalAddresses {
			result.Metadata["IpAddresses_"+strconv.Itoa(i)] = address.IPAddress
		}
	}

	return result, nil
}

// Return the value of the given key from the passed in slice of properties, if it exists
func findPropertyValue(properties []Property, key string) (string, bool) {
	for _, prop := range properties {
		if prop.PropertyKey == key {
			return prop.PropertyValue, true
		}
	}
	return "", false
}

// Convert the metadata to a map
func convertMetadata(metadata []KeyValue) map[string]any {
	result := make(map[string]any)
	for _, m := range metadata {
		result[m.Key] = m.Value
	}
	return result
}

// Gets any specific Metadata that is passed in the metadata map by its key
func getSpecificMetadata(original OriginalMetadata, key string) (string, bool) {
	var result string

	for _, metadataMap := range original.Metadata {
		if value, exists := metadataMap[key]; exists {
			result = fmt.Sprintf("%v", value)
			break
		}
	}

	if result == "" {
		return "", false
	}

	return result, true
}

// Convert the original metadata to the desired format
func convertFormat(original OriginalMetadata) KeyValueMetadata {
	var result KeyValueMetadata

	for _, metadataMap := range original.Metadata {
		for key, value := range metadataMap {
			result.Metadata = append(result.Metadata, KeyValue{
				Key:   key,
				Value: fmt.Sprintf("%v", value),
			})
		}
	}

	return result
}

// Utility function to verify the filters are valid
func verifyFilters(filters string) bool {
	if filters == "" {
		return false
	}

	// Split by AND and validate each condition
	conditions := strings.Split(filters, "AND")
	valid := true

	for _, condition := range conditions {
		// Trim spaces and split by "="
		condition = strings.TrimSpace(condition)
		parts := strings.Split(condition, "=")

		// Check if we have exactly two parts (key and value)
		if len(parts) != 2 {
			valid = false
			break
		}

		// Trim spaces from both parts
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])

		// Check if both parts are non-empty
		if key == "" || value == "" {
			valid = false
			break
		}

		// Check if the key is allowed
		if !slices.Contains(AllowedFilters, key) {
			valid = false
			break
		}
	}

	return valid
}
