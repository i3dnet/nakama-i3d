# ApplicationInstanceCrashDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Timestamp of the data | [readonly] 
**Crashes** | **int32** | The number of crashes of application instances | [readonly] 

## Methods

### NewApplicationInstanceCrashDataTelemetry

`func NewApplicationInstanceCrashDataTelemetry(timestamp int32, crashes int32, ) *ApplicationInstanceCrashDataTelemetry`

NewApplicationInstanceCrashDataTelemetry instantiates a new ApplicationInstanceCrashDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceCrashDataTelemetryWithDefaults

`func NewApplicationInstanceCrashDataTelemetryWithDefaults() *ApplicationInstanceCrashDataTelemetry`

NewApplicationInstanceCrashDataTelemetryWithDefaults instantiates a new ApplicationInstanceCrashDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstanceCrashDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstanceCrashDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstanceCrashDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetCrashes

`func (o *ApplicationInstanceCrashDataTelemetry) GetCrashes() int32`

GetCrashes returns the Crashes field if non-nil, zero value otherwise.

### GetCrashesOk

`func (o *ApplicationInstanceCrashDataTelemetry) GetCrashesOk() (*int32, bool)`

GetCrashesOk returns a tuple with the Crashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashes

`func (o *ApplicationInstanceCrashDataTelemetry) SetCrashes(v int32)`

SetCrashes sets Crashes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


