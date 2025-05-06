# ApplicationInstanceCrashTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency | [readonly] 
**Data** | [**[]ApplicationInstanceCrashDataTelemetry**](ApplicationInstanceCrashDataTelemetry.md) | Contains the data object by host type | [readonly] 

## Methods

### NewApplicationInstanceCrashTelemetry

`func NewApplicationInstanceCrashTelemetry(unit string, data []ApplicationInstanceCrashDataTelemetry, ) *ApplicationInstanceCrashTelemetry`

NewApplicationInstanceCrashTelemetry instantiates a new ApplicationInstanceCrashTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceCrashTelemetryWithDefaults

`func NewApplicationInstanceCrashTelemetryWithDefaults() *ApplicationInstanceCrashTelemetry`

NewApplicationInstanceCrashTelemetryWithDefaults instantiates a new ApplicationInstanceCrashTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationInstanceCrashTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationInstanceCrashTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationInstanceCrashTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *ApplicationInstanceCrashTelemetry) GetData() []ApplicationInstanceCrashDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationInstanceCrashTelemetry) GetDataOk() (*[]ApplicationInstanceCrashDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationInstanceCrashTelemetry) SetData(v []ApplicationInstanceCrashDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


