# InstanceTypeBM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the instance type | [readonly] 
**InstanceType** | **string** | The name of the instance type | [readonly] 
**CpuType** | **string** | The type of CPU of the instance type | [readonly] 
**CpuSpeed** | **int32** | The speed of CPU of the instance type | [readonly] 
**NumCpu** | **int32** | The number of CPUs of the instance type | [readonly] 
**NumCoresPerCpu** | **int32** | The number of cores per CPUs of the instance type | [readonly] 
**MemoryType** | **string** | The type of memory of the instance type | [readonly] 
**MemorySpeed** | **int32** | The memory speed of the instance type | [readonly] 
**MemorySize** | **int32** | The amount of memory of the instance type in MB | [readonly] 

## Methods

### NewInstanceTypeBM

`func NewInstanceTypeBM(id string, instanceType string, cpuType string, cpuSpeed int32, numCpu int32, numCoresPerCpu int32, memoryType string, memorySpeed int32, memorySize int32, ) *InstanceTypeBM`

NewInstanceTypeBM instantiates a new InstanceTypeBM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeBMWithDefaults

`func NewInstanceTypeBMWithDefaults() *InstanceTypeBM`

NewInstanceTypeBMWithDefaults instantiates a new InstanceTypeBM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeBM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeBM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeBM) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *InstanceTypeBM) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceTypeBM) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceTypeBM) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetCpuType

`func (o *InstanceTypeBM) GetCpuType() string`

GetCpuType returns the CpuType field if non-nil, zero value otherwise.

### GetCpuTypeOk

`func (o *InstanceTypeBM) GetCpuTypeOk() (*string, bool)`

GetCpuTypeOk returns a tuple with the CpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuType

`func (o *InstanceTypeBM) SetCpuType(v string)`

SetCpuType sets CpuType field to given value.


### GetCpuSpeed

`func (o *InstanceTypeBM) GetCpuSpeed() int32`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *InstanceTypeBM) GetCpuSpeedOk() (*int32, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *InstanceTypeBM) SetCpuSpeed(v int32)`

SetCpuSpeed sets CpuSpeed field to given value.


### GetNumCpu

`func (o *InstanceTypeBM) GetNumCpu() int32`

GetNumCpu returns the NumCpu field if non-nil, zero value otherwise.

### GetNumCpuOk

`func (o *InstanceTypeBM) GetNumCpuOk() (*int32, bool)`

GetNumCpuOk returns a tuple with the NumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpu

`func (o *InstanceTypeBM) SetNumCpu(v int32)`

SetNumCpu sets NumCpu field to given value.


### GetNumCoresPerCpu

`func (o *InstanceTypeBM) GetNumCoresPerCpu() int32`

GetNumCoresPerCpu returns the NumCoresPerCpu field if non-nil, zero value otherwise.

### GetNumCoresPerCpuOk

`func (o *InstanceTypeBM) GetNumCoresPerCpuOk() (*int32, bool)`

GetNumCoresPerCpuOk returns a tuple with the NumCoresPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresPerCpu

`func (o *InstanceTypeBM) SetNumCoresPerCpu(v int32)`

SetNumCoresPerCpu sets NumCoresPerCpu field to given value.


### GetMemoryType

`func (o *InstanceTypeBM) GetMemoryType() string`

GetMemoryType returns the MemoryType field if non-nil, zero value otherwise.

### GetMemoryTypeOk

`func (o *InstanceTypeBM) GetMemoryTypeOk() (*string, bool)`

GetMemoryTypeOk returns a tuple with the MemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryType

`func (o *InstanceTypeBM) SetMemoryType(v string)`

SetMemoryType sets MemoryType field to given value.


### GetMemorySpeed

`func (o *InstanceTypeBM) GetMemorySpeed() int32`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *InstanceTypeBM) GetMemorySpeedOk() (*int32, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *InstanceTypeBM) SetMemorySpeed(v int32)`

SetMemorySpeed sets MemorySpeed field to given value.


### GetMemorySize

`func (o *InstanceTypeBM) GetMemorySize() int32`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *InstanceTypeBM) GetMemorySizeOk() (*int32, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *InstanceTypeBM) SetMemorySize(v int32)`

SetMemorySize sets MemorySize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


