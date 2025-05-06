# ApplicationTelemetryHistoryDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Unix timestamp of data | [readonly] 
**Hosts** | **int32** | The number of hosts | 
**Instances** | **int32** | The number of instances | 
**Players** | **int32** | The number of players | 

## Methods

### NewApplicationTelemetryHistoryDataModel

`func NewApplicationTelemetryHistoryDataModel(timestamp int32, hosts int32, instances int32, players int32, ) *ApplicationTelemetryHistoryDataModel`

NewApplicationTelemetryHistoryDataModel instantiates a new ApplicationTelemetryHistoryDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTelemetryHistoryDataModelWithDefaults

`func NewApplicationTelemetryHistoryDataModelWithDefaults() *ApplicationTelemetryHistoryDataModel`

NewApplicationTelemetryHistoryDataModelWithDefaults instantiates a new ApplicationTelemetryHistoryDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationTelemetryHistoryDataModel) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationTelemetryHistoryDataModel) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationTelemetryHistoryDataModel) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetHosts

`func (o *ApplicationTelemetryHistoryDataModel) GetHosts() int32`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ApplicationTelemetryHistoryDataModel) GetHostsOk() (*int32, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ApplicationTelemetryHistoryDataModel) SetHosts(v int32)`

SetHosts sets Hosts field to given value.


### GetInstances

`func (o *ApplicationTelemetryHistoryDataModel) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ApplicationTelemetryHistoryDataModel) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ApplicationTelemetryHistoryDataModel) SetInstances(v int32)`

SetInstances sets Instances field to given value.


### GetPlayers

`func (o *ApplicationTelemetryHistoryDataModel) GetPlayers() int32`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *ApplicationTelemetryHistoryDataModel) GetPlayersOk() (*int32, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *ApplicationTelemetryHistoryDataModel) SetPlayers(v int32)`

SetPlayers sets Players field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


