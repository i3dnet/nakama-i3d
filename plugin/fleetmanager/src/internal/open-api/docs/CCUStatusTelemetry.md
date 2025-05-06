# CCUStatusTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, which is always &#39;1m&#39; in this endpoint | [readonly] 
**Data** | [**[]CCUStatusDataTelemetry**](CCUStatusDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewCCUStatusTelemetry

`func NewCCUStatusTelemetry(unit string, data []CCUStatusDataTelemetry, ) *CCUStatusTelemetry`

NewCCUStatusTelemetry instantiates a new CCUStatusTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUStatusTelemetryWithDefaults

`func NewCCUStatusTelemetryWithDefaults() *CCUStatusTelemetry`

NewCCUStatusTelemetryWithDefaults instantiates a new CCUStatusTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *CCUStatusTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CCUStatusTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CCUStatusTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *CCUStatusTelemetry) GetData() []CCUStatusDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CCUStatusTelemetry) GetDataOk() (*[]CCUStatusDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CCUStatusTelemetry) SetData(v []CCUStatusDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


