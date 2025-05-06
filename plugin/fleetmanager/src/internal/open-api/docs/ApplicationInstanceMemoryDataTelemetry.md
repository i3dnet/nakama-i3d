# ApplicationInstanceMemoryDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**MemUsed** | **int32** | The amount of memory that is used by the ApplicationInstance (in megabytes) | [readonly] 

## Methods

### NewApplicationInstanceMemoryDataTelemetry

`func NewApplicationInstanceMemoryDataTelemetry(timestamp int32, memUsed int32, ) *ApplicationInstanceMemoryDataTelemetry`

NewApplicationInstanceMemoryDataTelemetry instantiates a new ApplicationInstanceMemoryDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceMemoryDataTelemetryWithDefaults

`func NewApplicationInstanceMemoryDataTelemetryWithDefaults() *ApplicationInstanceMemoryDataTelemetry`

NewApplicationInstanceMemoryDataTelemetryWithDefaults instantiates a new ApplicationInstanceMemoryDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstanceMemoryDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstanceMemoryDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstanceMemoryDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetMemUsed

`func (o *ApplicationInstanceMemoryDataTelemetry) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *ApplicationInstanceMemoryDataTelemetry) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *ApplicationInstanceMemoryDataTelemetry) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


