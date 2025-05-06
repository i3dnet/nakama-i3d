# ApplicationInstanceStatusTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, possible interval options are as follows - minute - hour - day - week | [readonly] 
**Data** | [**[]ApplicationInstanceStatusDataTelemetry**](ApplicationInstanceStatusDataTelemetry.md) | Average aggregated data of application status | [readonly] 

## Methods

### NewApplicationInstanceStatusTelemetry

`func NewApplicationInstanceStatusTelemetry(unit string, data []ApplicationInstanceStatusDataTelemetry, ) *ApplicationInstanceStatusTelemetry`

NewApplicationInstanceStatusTelemetry instantiates a new ApplicationInstanceStatusTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStatusTelemetryWithDefaults

`func NewApplicationInstanceStatusTelemetryWithDefaults() *ApplicationInstanceStatusTelemetry`

NewApplicationInstanceStatusTelemetryWithDefaults instantiates a new ApplicationInstanceStatusTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationInstanceStatusTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationInstanceStatusTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationInstanceStatusTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *ApplicationInstanceStatusTelemetry) GetData() []ApplicationInstanceStatusDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationInstanceStatusTelemetry) GetDataOk() (*[]ApplicationInstanceStatusDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationInstanceStatusTelemetry) SetData(v []ApplicationInstanceStatusDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


