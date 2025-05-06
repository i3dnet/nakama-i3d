# HostHardwareCapacityDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**Value** | [**HostHardwareCapacityDataValueTelemetry**](HostHardwareCapacityDataValueTelemetry.md) | The value of data at this timestamp | [readonly] 

## Methods

### NewHostHardwareCapacityDataTelemetry

`func NewHostHardwareCapacityDataTelemetry(timestamp int32, value HostHardwareCapacityDataValueTelemetry, ) *HostHardwareCapacityDataTelemetry`

NewHostHardwareCapacityDataTelemetry instantiates a new HostHardwareCapacityDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHardwareCapacityDataTelemetryWithDefaults

`func NewHostHardwareCapacityDataTelemetryWithDefaults() *HostHardwareCapacityDataTelemetry`

NewHostHardwareCapacityDataTelemetryWithDefaults instantiates a new HostHardwareCapacityDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostHardwareCapacityDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostHardwareCapacityDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostHardwareCapacityDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetValue

`func (o *HostHardwareCapacityDataTelemetry) GetValue() HostHardwareCapacityDataValueTelemetry`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HostHardwareCapacityDataTelemetry) GetValueOk() (*HostHardwareCapacityDataValueTelemetry, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HostHardwareCapacityDataTelemetry) SetValue(v HostHardwareCapacityDataValueTelemetry)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


