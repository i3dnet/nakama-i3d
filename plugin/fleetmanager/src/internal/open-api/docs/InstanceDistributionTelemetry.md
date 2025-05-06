# InstanceDistributionTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency,  possible interval options are as follows - minute - hour - day - month | [readonly] 
**Data** | [**[]InstanceDistributionDataTelemetry**](InstanceDistributionDataTelemetry.md) | Average aggregated data of application instances | [readonly] 

## Methods

### NewInstanceDistributionTelemetry

`func NewInstanceDistributionTelemetry(unit string, data []InstanceDistributionDataTelemetry, ) *InstanceDistributionTelemetry`

NewInstanceDistributionTelemetry instantiates a new InstanceDistributionTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDistributionTelemetryWithDefaults

`func NewInstanceDistributionTelemetryWithDefaults() *InstanceDistributionTelemetry`

NewInstanceDistributionTelemetryWithDefaults instantiates a new InstanceDistributionTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *InstanceDistributionTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InstanceDistributionTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InstanceDistributionTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *InstanceDistributionTelemetry) GetData() []InstanceDistributionDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstanceDistributionTelemetry) GetDataOk() (*[]InstanceDistributionDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstanceDistributionTelemetry) SetData(v []InstanceDistributionDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


