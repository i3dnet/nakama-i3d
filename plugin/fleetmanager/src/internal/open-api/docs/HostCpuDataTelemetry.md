# HostCpuDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**Cpu** | **int32** | Percentage of cpu used across all cores | [readonly] 

## Methods

### NewHostCpuDataTelemetry

`func NewHostCpuDataTelemetry(timestamp int32, cpu int32, ) *HostCpuDataTelemetry`

NewHostCpuDataTelemetry instantiates a new HostCpuDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCpuDataTelemetryWithDefaults

`func NewHostCpuDataTelemetryWithDefaults() *HostCpuDataTelemetry`

NewHostCpuDataTelemetryWithDefaults instantiates a new HostCpuDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostCpuDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostCpuDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostCpuDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetCpu

`func (o *HostCpuDataTelemetry) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *HostCpuDataTelemetry) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *HostCpuDataTelemetry) SetCpu(v int32)`

SetCpu sets Cpu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


