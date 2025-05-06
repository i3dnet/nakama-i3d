# HostHddTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, which is always &#39;1m&#39; in this endpoint | [readonly] 
**Data** | [**[]HostHddDataTelemetry**](HostHddDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewHostHddTelemetry

`func NewHostHddTelemetry(unit string, data []HostHddDataTelemetry, ) *HostHddTelemetry`

NewHostHddTelemetry instantiates a new HostHddTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHddTelemetryWithDefaults

`func NewHostHddTelemetryWithDefaults() *HostHddTelemetry`

NewHostHddTelemetryWithDefaults instantiates a new HostHddTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostHddTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostHddTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostHddTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *HostHddTelemetry) GetData() []HostHddDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostHddTelemetry) GetDataOk() (*[]HostHddDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostHddTelemetry) SetData(v []HostHddDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


