# CCUStatusAggregateDataContainerTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Unix time in which the data was saved | [readonly] 
**Ccus** | [**[]CCUStatusAggregateDataTelemetry**](CCUStatusAggregateDataTelemetry.md) | List containing the telemetry data grouped by providerId | [readonly] 

## Methods

### NewCCUStatusAggregateDataContainerTelemetry

`func NewCCUStatusAggregateDataContainerTelemetry(timestamp int32, ccus []CCUStatusAggregateDataTelemetry, ) *CCUStatusAggregateDataContainerTelemetry`

NewCCUStatusAggregateDataContainerTelemetry instantiates a new CCUStatusAggregateDataContainerTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUStatusAggregateDataContainerTelemetryWithDefaults

`func NewCCUStatusAggregateDataContainerTelemetryWithDefaults() *CCUStatusAggregateDataContainerTelemetry`

NewCCUStatusAggregateDataContainerTelemetryWithDefaults instantiates a new CCUStatusAggregateDataContainerTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *CCUStatusAggregateDataContainerTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CCUStatusAggregateDataContainerTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CCUStatusAggregateDataContainerTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetCcus

`func (o *CCUStatusAggregateDataContainerTelemetry) GetCcus() []CCUStatusAggregateDataTelemetry`

GetCcus returns the Ccus field if non-nil, zero value otherwise.

### GetCcusOk

`func (o *CCUStatusAggregateDataContainerTelemetry) GetCcusOk() (*[]CCUStatusAggregateDataTelemetry, bool)`

GetCcusOk returns a tuple with the Ccus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcus

`func (o *CCUStatusAggregateDataContainerTelemetry) SetCcus(v []CCUStatusAggregateDataTelemetry)`

SetCcus sets Ccus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


