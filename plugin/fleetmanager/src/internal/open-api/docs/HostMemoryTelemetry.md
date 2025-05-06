# HostMemoryTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, which is always &#39;1m&#39; in this endpoint | [readonly] 
**MemMax** | **int32** | The amount of memory that is available on the host (in megabytes) | [readonly] 
**Data** | [**[]HostMemoryDataTelemetry**](HostMemoryDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewHostMemoryTelemetry

`func NewHostMemoryTelemetry(unit string, memMax int32, data []HostMemoryDataTelemetry, ) *HostMemoryTelemetry`

NewHostMemoryTelemetry instantiates a new HostMemoryTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMemoryTelemetryWithDefaults

`func NewHostMemoryTelemetryWithDefaults() *HostMemoryTelemetry`

NewHostMemoryTelemetryWithDefaults instantiates a new HostMemoryTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostMemoryTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostMemoryTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostMemoryTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetMemMax

`func (o *HostMemoryTelemetry) GetMemMax() int32`

GetMemMax returns the MemMax field if non-nil, zero value otherwise.

### GetMemMaxOk

`func (o *HostMemoryTelemetry) GetMemMaxOk() (*int32, bool)`

GetMemMaxOk returns a tuple with the MemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemMax

`func (o *HostMemoryTelemetry) SetMemMax(v int32)`

SetMemMax sets MemMax field to given value.


### GetData

`func (o *HostMemoryTelemetry) GetData() []HostMemoryDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostMemoryTelemetry) GetDataOk() (*[]HostMemoryDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostMemoryTelemetry) SetData(v []HostMemoryDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


