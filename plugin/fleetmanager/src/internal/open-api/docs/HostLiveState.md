# HostLiveState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this host | [readonly] 
**ServerId** | **int32** | The ID of the physical machine | [readonly] 
**CpuLoad** | **float32** | Percentage of cpu used across all cores | [readonly] 
**MemMax** | **int32** | The amount of memory that is available on the host (in megabytes) | [readonly] 
**MemUsed** | **int32** | The amount of memory that is used by the host (in megabytes) | [readonly] 
**MemFree** | **int32** | The amount of free memory that is available on the host (in megabytes) | [readonly] 

## Methods

### NewHostLiveState

`func NewHostLiveState(id int32, serverId int32, cpuLoad float32, memMax int32, memUsed int32, memFree int32, ) *HostLiveState`

NewHostLiveState instantiates a new HostLiveState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostLiveStateWithDefaults

`func NewHostLiveStateWithDefaults() *HostLiveState`

NewHostLiveStateWithDefaults instantiates a new HostLiveState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostLiveState) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostLiveState) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostLiveState) SetId(v int32)`

SetId sets Id field to given value.


### GetServerId

`func (o *HostLiveState) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *HostLiveState) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *HostLiveState) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetCpuLoad

`func (o *HostLiveState) GetCpuLoad() float32`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *HostLiveState) GetCpuLoadOk() (*float32, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *HostLiveState) SetCpuLoad(v float32)`

SetCpuLoad sets CpuLoad field to given value.


### GetMemMax

`func (o *HostLiveState) GetMemMax() int32`

GetMemMax returns the MemMax field if non-nil, zero value otherwise.

### GetMemMaxOk

`func (o *HostLiveState) GetMemMaxOk() (*int32, bool)`

GetMemMaxOk returns a tuple with the MemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemMax

`func (o *HostLiveState) SetMemMax(v int32)`

SetMemMax sets MemMax field to given value.


### GetMemUsed

`func (o *HostLiveState) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *HostLiveState) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *HostLiveState) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.


### GetMemFree

`func (o *HostLiveState) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *HostLiveState) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *HostLiveState) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


