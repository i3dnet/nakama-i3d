# ApplicationTelemetryHistoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency,  possible interval options are as follows - minute - hour - day - month | [readonly] 
**Data** | [**[]ApplicationTelemetryHistoryDataModel**](ApplicationTelemetryHistoryDataModel.md) | Average aggregated data of application instances | [readonly] 

## Methods

### NewApplicationTelemetryHistoryModel

`func NewApplicationTelemetryHistoryModel(unit string, data []ApplicationTelemetryHistoryDataModel, ) *ApplicationTelemetryHistoryModel`

NewApplicationTelemetryHistoryModel instantiates a new ApplicationTelemetryHistoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTelemetryHistoryModelWithDefaults

`func NewApplicationTelemetryHistoryModelWithDefaults() *ApplicationTelemetryHistoryModel`

NewApplicationTelemetryHistoryModelWithDefaults instantiates a new ApplicationTelemetryHistoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationTelemetryHistoryModel) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationTelemetryHistoryModel) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationTelemetryHistoryModel) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *ApplicationTelemetryHistoryModel) GetData() []ApplicationTelemetryHistoryDataModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationTelemetryHistoryModel) GetDataOk() (*[]ApplicationTelemetryHistoryDataModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationTelemetryHistoryModel) SetData(v []ApplicationTelemetryHistoryDataModel)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


