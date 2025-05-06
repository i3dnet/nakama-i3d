# HostCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Socket** | **int32** | Number of populated CPU sockets | [readonly] 
**Cores** | **int32** | Total number of cores (sum of all CPU cores) | [readonly] 
**Threads** | **int32** | Total number of threads (sum of all CPU threads). | [readonly] 
**Info** | **string** | CPU information | [readonly] 
**Type** | **string** | CPU type | [readonly] 

## Methods

### NewHostCpu

`func NewHostCpu(socket int32, cores int32, threads int32, info string, type_ string, ) *HostCpu`

NewHostCpu instantiates a new HostCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCpuWithDefaults

`func NewHostCpuWithDefaults() *HostCpu`

NewHostCpuWithDefaults instantiates a new HostCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSocket

`func (o *HostCpu) GetSocket() int32`

GetSocket returns the Socket field if non-nil, zero value otherwise.

### GetSocketOk

`func (o *HostCpu) GetSocketOk() (*int32, bool)`

GetSocketOk returns a tuple with the Socket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocket

`func (o *HostCpu) SetSocket(v int32)`

SetSocket sets Socket field to given value.


### GetCores

`func (o *HostCpu) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *HostCpu) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *HostCpu) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetThreads

`func (o *HostCpu) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *HostCpu) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *HostCpu) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetInfo

`func (o *HostCpu) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *HostCpu) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *HostCpu) SetInfo(v string)`

SetInfo sets Info field to given value.


### GetType

`func (o *HostCpu) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostCpu) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostCpu) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


