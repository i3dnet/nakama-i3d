# HostUtilizationAggregateTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, possible interval options are as follows - minute - hour - day - month | [readonly] 
**Data** | [**[]HostUtilizationAggregateDataContainerTelemetry**](HostUtilizationAggregateDataContainerTelemetry.md) | Object containing the data object grouped by timestamp and providerId | [readonly] 

## Methods

### NewHostUtilizationAggregateTelemetry

`func NewHostUtilizationAggregateTelemetry(unit string, data []HostUtilizationAggregateDataContainerTelemetry, ) *HostUtilizationAggregateTelemetry`

NewHostUtilizationAggregateTelemetry instantiates a new HostUtilizationAggregateTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUtilizationAggregateTelemetryWithDefaults

`func NewHostUtilizationAggregateTelemetryWithDefaults() *HostUtilizationAggregateTelemetry`

NewHostUtilizationAggregateTelemetryWithDefaults instantiates a new HostUtilizationAggregateTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostUtilizationAggregateTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostUtilizationAggregateTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostUtilizationAggregateTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *HostUtilizationAggregateTelemetry) GetData() []HostUtilizationAggregateDataContainerTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostUtilizationAggregateTelemetry) GetDataOk() (*[]HostUtilizationAggregateDataContainerTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostUtilizationAggregateTelemetry) SetData(v []HostUtilizationAggregateDataContainerTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


