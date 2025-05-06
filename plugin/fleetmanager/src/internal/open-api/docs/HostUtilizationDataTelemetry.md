# HostUtilizationDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**HostsOccupied** | **int32** | Number of hosts that are occupied (one or more (but not the maximum amount of) application instances on them) | [readonly] 
**HostsFree** | **int32** | Number of host that are free (no application instances on them) | [readonly] 
**HostsFull** | **int32** | Number of host that are full (the maximum amount of application instances on them) See also GET /v3/hostCapacityTemplate and GET /v3/hostCapacityTemplate/{hostCapacityTemplateId}/instanceTypeCapacity | [readonly] 
**HostsTotal** | **int32** | Number of hosts in total | [readonly] 
**Launched** | **int32** | Number of launched instances since the previous entry | [readonly] 
**Destroyed** | **int32** | Number of destroyed instances since the previous entry | [readonly] 

## Methods

### NewHostUtilizationDataTelemetry

`func NewHostUtilizationDataTelemetry(timestamp int32, hostsOccupied int32, hostsFree int32, hostsFull int32, hostsTotal int32, launched int32, destroyed int32, ) *HostUtilizationDataTelemetry`

NewHostUtilizationDataTelemetry instantiates a new HostUtilizationDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUtilizationDataTelemetryWithDefaults

`func NewHostUtilizationDataTelemetryWithDefaults() *HostUtilizationDataTelemetry`

NewHostUtilizationDataTelemetryWithDefaults instantiates a new HostUtilizationDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostUtilizationDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostUtilizationDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostUtilizationDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetHostsOccupied

`func (o *HostUtilizationDataTelemetry) GetHostsOccupied() int32`

GetHostsOccupied returns the HostsOccupied field if non-nil, zero value otherwise.

### GetHostsOccupiedOk

`func (o *HostUtilizationDataTelemetry) GetHostsOccupiedOk() (*int32, bool)`

GetHostsOccupiedOk returns a tuple with the HostsOccupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsOccupied

`func (o *HostUtilizationDataTelemetry) SetHostsOccupied(v int32)`

SetHostsOccupied sets HostsOccupied field to given value.


### GetHostsFree

`func (o *HostUtilizationDataTelemetry) GetHostsFree() int32`

GetHostsFree returns the HostsFree field if non-nil, zero value otherwise.

### GetHostsFreeOk

`func (o *HostUtilizationDataTelemetry) GetHostsFreeOk() (*int32, bool)`

GetHostsFreeOk returns a tuple with the HostsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsFree

`func (o *HostUtilizationDataTelemetry) SetHostsFree(v int32)`

SetHostsFree sets HostsFree field to given value.


### GetHostsFull

`func (o *HostUtilizationDataTelemetry) GetHostsFull() int32`

GetHostsFull returns the HostsFull field if non-nil, zero value otherwise.

### GetHostsFullOk

`func (o *HostUtilizationDataTelemetry) GetHostsFullOk() (*int32, bool)`

GetHostsFullOk returns a tuple with the HostsFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsFull

`func (o *HostUtilizationDataTelemetry) SetHostsFull(v int32)`

SetHostsFull sets HostsFull field to given value.


### GetHostsTotal

`func (o *HostUtilizationDataTelemetry) GetHostsTotal() int32`

GetHostsTotal returns the HostsTotal field if non-nil, zero value otherwise.

### GetHostsTotalOk

`func (o *HostUtilizationDataTelemetry) GetHostsTotalOk() (*int32, bool)`

GetHostsTotalOk returns a tuple with the HostsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsTotal

`func (o *HostUtilizationDataTelemetry) SetHostsTotal(v int32)`

SetHostsTotal sets HostsTotal field to given value.


### GetLaunched

`func (o *HostUtilizationDataTelemetry) GetLaunched() int32`

GetLaunched returns the Launched field if non-nil, zero value otherwise.

### GetLaunchedOk

`func (o *HostUtilizationDataTelemetry) GetLaunchedOk() (*int32, bool)`

GetLaunchedOk returns a tuple with the Launched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunched

`func (o *HostUtilizationDataTelemetry) SetLaunched(v int32)`

SetLaunched sets Launched field to given value.


### GetDestroyed

`func (o *HostUtilizationDataTelemetry) GetDestroyed() int32`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *HostUtilizationDataTelemetry) GetDestroyedOk() (*int32, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *HostUtilizationDataTelemetry) SetDestroyed(v int32)`

SetDestroyed sets Destroyed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


