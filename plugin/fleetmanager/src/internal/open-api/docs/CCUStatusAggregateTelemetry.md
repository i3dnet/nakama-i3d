# CCUStatusAggregateTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, possible interval options are as follows - minute - hour - day - month | [readonly] 
**Data** | [**[]CCUStatusAggregateDataContainerTelemetry**](CCUStatusAggregateDataContainerTelemetry.md) | Object containing the data object grouped by timestamp and providerId | [readonly] 

## Methods

### NewCCUStatusAggregateTelemetry

`func NewCCUStatusAggregateTelemetry(unit string, data []CCUStatusAggregateDataContainerTelemetry, ) *CCUStatusAggregateTelemetry`

NewCCUStatusAggregateTelemetry instantiates a new CCUStatusAggregateTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUStatusAggregateTelemetryWithDefaults

`func NewCCUStatusAggregateTelemetryWithDefaults() *CCUStatusAggregateTelemetry`

NewCCUStatusAggregateTelemetryWithDefaults instantiates a new CCUStatusAggregateTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *CCUStatusAggregateTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CCUStatusAggregateTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CCUStatusAggregateTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *CCUStatusAggregateTelemetry) GetData() []CCUStatusAggregateDataContainerTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CCUStatusAggregateTelemetry) GetDataOk() (*[]CCUStatusAggregateDataContainerTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CCUStatusAggregateTelemetry) SetData(v []CCUStatusAggregateDataContainerTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


