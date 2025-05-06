# HostCpuTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency, which is always &#39;1m&#39; in this endpoint | [readonly] 
**Data** | [**[]HostCpuDataTelemetry**](HostCpuDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewHostCpuTelemetry

`func NewHostCpuTelemetry(unit string, data []HostCpuDataTelemetry, ) *HostCpuTelemetry`

NewHostCpuTelemetry instantiates a new HostCpuTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCpuTelemetryWithDefaults

`func NewHostCpuTelemetryWithDefaults() *HostCpuTelemetry`

NewHostCpuTelemetryWithDefaults instantiates a new HostCpuTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostCpuTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostCpuTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostCpuTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *HostCpuTelemetry) GetData() []HostCpuDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostCpuTelemetry) GetDataOk() (*[]HostCpuDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostCpuTelemetry) SetData(v []HostCpuDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


