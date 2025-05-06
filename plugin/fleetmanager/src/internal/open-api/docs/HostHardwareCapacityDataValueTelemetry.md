# HostHardwareCapacityDataValueTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuPercentage** | **float32** | The amount of used CPU power as a percentage of the total amount of CPU power | [readonly] 
**RamPercentage** | **float32** | The amount of used RAM space as a percentage of the total amount of RAM | [readonly] 
**DiskPercentage** | **float32** | The amount of used disk space as a percentage of the total amount of space | [readonly] 

## Methods

### NewHostHardwareCapacityDataValueTelemetry

`func NewHostHardwareCapacityDataValueTelemetry(cpuPercentage float32, ramPercentage float32, diskPercentage float32, ) *HostHardwareCapacityDataValueTelemetry`

NewHostHardwareCapacityDataValueTelemetry instantiates a new HostHardwareCapacityDataValueTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHardwareCapacityDataValueTelemetryWithDefaults

`func NewHostHardwareCapacityDataValueTelemetryWithDefaults() *HostHardwareCapacityDataValueTelemetry`

NewHostHardwareCapacityDataValueTelemetryWithDefaults instantiates a new HostHardwareCapacityDataValueTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) GetCpuPercentage() float32`

GetCpuPercentage returns the CpuPercentage field if non-nil, zero value otherwise.

### GetCpuPercentageOk

`func (o *HostHardwareCapacityDataValueTelemetry) GetCpuPercentageOk() (*float32, bool)`

GetCpuPercentageOk returns a tuple with the CpuPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) SetCpuPercentage(v float32)`

SetCpuPercentage sets CpuPercentage field to given value.


### GetRamPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) GetRamPercentage() float32`

GetRamPercentage returns the RamPercentage field if non-nil, zero value otherwise.

### GetRamPercentageOk

`func (o *HostHardwareCapacityDataValueTelemetry) GetRamPercentageOk() (*float32, bool)`

GetRamPercentageOk returns a tuple with the RamPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) SetRamPercentage(v float32)`

SetRamPercentage sets RamPercentage field to given value.


### GetDiskPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) GetDiskPercentage() float32`

GetDiskPercentage returns the DiskPercentage field if non-nil, zero value otherwise.

### GetDiskPercentageOk

`func (o *HostHardwareCapacityDataValueTelemetry) GetDiskPercentageOk() (*float32, bool)`

GetDiskPercentageOk returns a tuple with the DiskPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPercentage

`func (o *HostHardwareCapacityDataValueTelemetry) SetDiskPercentage(v float32)`

SetDiskPercentage sets DiskPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


