# TrafficUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency,  possible interval options are as follows - day | [readonly] 
**Data** | [**[]TrafficUsageData**](TrafficUsageData.md) | Contains the data object for the usage | [readonly] 

## Methods

### NewTrafficUsage

`func NewTrafficUsage(unit string, data []TrafficUsageData, ) *TrafficUsage`

NewTrafficUsage instantiates a new TrafficUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficUsageWithDefaults

`func NewTrafficUsageWithDefaults() *TrafficUsage`

NewTrafficUsageWithDefaults instantiates a new TrafficUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *TrafficUsage) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TrafficUsage) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TrafficUsage) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *TrafficUsage) GetData() []TrafficUsageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrafficUsage) GetDataOk() (*[]TrafficUsageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrafficUsage) SetData(v []TrafficUsageData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


