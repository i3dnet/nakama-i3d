# HostUtilizationAggregateDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | [readonly] 
**HostsOccupied** | **int32** | Number of hosts that are occupied (one or more (but not the maximum amount of) application instances on them) | [readonly] 
**HostsFree** | **int32** | Number of host that are free (no application instances on them) | [readonly] 
**HostsFull** | **int32** | Number of host that are full (the maximum amount of application instances on them) See also [&#x60;GET /v3/hostCapacityTemplate&#x60;](#/HostCapacityTemplate/getHostCapacityTemplates) and [&#x60;GET /v3/hostCapacityTemplate/{hostCapacityTemplateId}/instanceTypeCapacity&#x60;](#/InstanceTypeCapacity/getHostCapacityTemplateInstanceTypeCapacities) | [readonly] 
**HostsTotal** | **int32** | Number of hosts in total | [readonly] 
**Launched** | **int32** | Number of launched instances since the previous entry | [readonly] 
**Destroyed** | **int32** | Number of destroyed instances since the previous entry | [readonly] 

## Methods

### NewHostUtilizationAggregateDataTelemetry

`func NewHostUtilizationAggregateDataTelemetry(providerId int32, hostsOccupied int32, hostsFree int32, hostsFull int32, hostsTotal int32, launched int32, destroyed int32, ) *HostUtilizationAggregateDataTelemetry`

NewHostUtilizationAggregateDataTelemetry instantiates a new HostUtilizationAggregateDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUtilizationAggregateDataTelemetryWithDefaults

`func NewHostUtilizationAggregateDataTelemetryWithDefaults() *HostUtilizationAggregateDataTelemetry`

NewHostUtilizationAggregateDataTelemetryWithDefaults instantiates a new HostUtilizationAggregateDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *HostUtilizationAggregateDataTelemetry) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *HostUtilizationAggregateDataTelemetry) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *HostUtilizationAggregateDataTelemetry) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetHostsOccupied

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsOccupied() int32`

GetHostsOccupied returns the HostsOccupied field if non-nil, zero value otherwise.

### GetHostsOccupiedOk

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsOccupiedOk() (*int32, bool)`

GetHostsOccupiedOk returns a tuple with the HostsOccupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsOccupied

`func (o *HostUtilizationAggregateDataTelemetry) SetHostsOccupied(v int32)`

SetHostsOccupied sets HostsOccupied field to given value.


### GetHostsFree

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsFree() int32`

GetHostsFree returns the HostsFree field if non-nil, zero value otherwise.

### GetHostsFreeOk

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsFreeOk() (*int32, bool)`

GetHostsFreeOk returns a tuple with the HostsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsFree

`func (o *HostUtilizationAggregateDataTelemetry) SetHostsFree(v int32)`

SetHostsFree sets HostsFree field to given value.


### GetHostsFull

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsFull() int32`

GetHostsFull returns the HostsFull field if non-nil, zero value otherwise.

### GetHostsFullOk

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsFullOk() (*int32, bool)`

GetHostsFullOk returns a tuple with the HostsFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsFull

`func (o *HostUtilizationAggregateDataTelemetry) SetHostsFull(v int32)`

SetHostsFull sets HostsFull field to given value.


### GetHostsTotal

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsTotal() int32`

GetHostsTotal returns the HostsTotal field if non-nil, zero value otherwise.

### GetHostsTotalOk

`func (o *HostUtilizationAggregateDataTelemetry) GetHostsTotalOk() (*int32, bool)`

GetHostsTotalOk returns a tuple with the HostsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsTotal

`func (o *HostUtilizationAggregateDataTelemetry) SetHostsTotal(v int32)`

SetHostsTotal sets HostsTotal field to given value.


### GetLaunched

`func (o *HostUtilizationAggregateDataTelemetry) GetLaunched() int32`

GetLaunched returns the Launched field if non-nil, zero value otherwise.

### GetLaunchedOk

`func (o *HostUtilizationAggregateDataTelemetry) GetLaunchedOk() (*int32, bool)`

GetLaunchedOk returns a tuple with the Launched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunched

`func (o *HostUtilizationAggregateDataTelemetry) SetLaunched(v int32)`

SetLaunched sets Launched field to given value.


### GetDestroyed

`func (o *HostUtilizationAggregateDataTelemetry) GetDestroyed() int32`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *HostUtilizationAggregateDataTelemetry) GetDestroyedOk() (*int32, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *HostUtilizationAggregateDataTelemetry) SetDestroyed(v int32)`

SetDestroyed sets Destroyed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


