# ApplicationInstanceCpuDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**Cpu** | **int32** | Percentage of cpu used across all cores | [readonly] 

## Methods

### NewApplicationInstanceCpuDataTelemetry

`func NewApplicationInstanceCpuDataTelemetry(timestamp int32, cpu int32, ) *ApplicationInstanceCpuDataTelemetry`

NewApplicationInstanceCpuDataTelemetry instantiates a new ApplicationInstanceCpuDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceCpuDataTelemetryWithDefaults

`func NewApplicationInstanceCpuDataTelemetryWithDefaults() *ApplicationInstanceCpuDataTelemetry`

NewApplicationInstanceCpuDataTelemetryWithDefaults instantiates a new ApplicationInstanceCpuDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstanceCpuDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstanceCpuDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstanceCpuDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetCpu

`func (o *ApplicationInstanceCpuDataTelemetry) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationInstanceCpuDataTelemetry) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationInstanceCpuDataTelemetry) SetCpu(v int32)`

SetCpu sets Cpu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


