# ApplicationInstanceMemoryTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, which is always &#x60;null&#x60; in this endpoint. This property is there for consistency across all telemetry endpoints. It&#39;s &#x60;null&#x60; because the data is not aggregated | [readonly] 
**Data** | [**[]ApplicationInstanceMemoryDataTelemetry**](ApplicationInstanceMemoryDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewApplicationInstanceMemoryTelemetry

`func NewApplicationInstanceMemoryTelemetry(unit string, data []ApplicationInstanceMemoryDataTelemetry, ) *ApplicationInstanceMemoryTelemetry`

NewApplicationInstanceMemoryTelemetry instantiates a new ApplicationInstanceMemoryTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceMemoryTelemetryWithDefaults

`func NewApplicationInstanceMemoryTelemetryWithDefaults() *ApplicationInstanceMemoryTelemetry`

NewApplicationInstanceMemoryTelemetryWithDefaults instantiates a new ApplicationInstanceMemoryTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationInstanceMemoryTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationInstanceMemoryTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationInstanceMemoryTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *ApplicationInstanceMemoryTelemetry) GetData() []ApplicationInstanceMemoryDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationInstanceMemoryTelemetry) GetDataOk() (*[]ApplicationInstanceMemoryDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationInstanceMemoryTelemetry) SetData(v []ApplicationInstanceMemoryDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


