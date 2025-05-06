# ApplicationInstancesTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency,  possible interval options are as follows - minute - hour - day - week | [readonly] 
**Data** | [**[]ApplicationInstancesDataTelemetry**](ApplicationInstancesDataTelemetry.md) | Average aggregated data of application instances | [readonly] 

## Methods

### NewApplicationInstancesTelemetry

`func NewApplicationInstancesTelemetry(unit string, data []ApplicationInstancesDataTelemetry, ) *ApplicationInstancesTelemetry`

NewApplicationInstancesTelemetry instantiates a new ApplicationInstancesTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesTelemetryWithDefaults

`func NewApplicationInstancesTelemetryWithDefaults() *ApplicationInstancesTelemetry`

NewApplicationInstancesTelemetryWithDefaults instantiates a new ApplicationInstancesTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationInstancesTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationInstancesTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationInstancesTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *ApplicationInstancesTelemetry) GetData() []ApplicationInstancesDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationInstancesTelemetry) GetDataOk() (*[]ApplicationInstancesDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationInstancesTelemetry) SetData(v []ApplicationInstancesDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


