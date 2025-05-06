# HostHardwareCapacityTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency | [readonly] 
**Data** | [**[]HostHardwareCapacityDataTelemetry**](HostHardwareCapacityDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewHostHardwareCapacityTelemetry

`func NewHostHardwareCapacityTelemetry(unit string, data []HostHardwareCapacityDataTelemetry, ) *HostHardwareCapacityTelemetry`

NewHostHardwareCapacityTelemetry instantiates a new HostHardwareCapacityTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHardwareCapacityTelemetryWithDefaults

`func NewHostHardwareCapacityTelemetryWithDefaults() *HostHardwareCapacityTelemetry`

NewHostHardwareCapacityTelemetryWithDefaults instantiates a new HostHardwareCapacityTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostHardwareCapacityTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostHardwareCapacityTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostHardwareCapacityTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetData

`func (o *HostHardwareCapacityTelemetry) GetData() []HostHardwareCapacityDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostHardwareCapacityTelemetry) GetDataOk() (*[]HostHardwareCapacityDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostHardwareCapacityTelemetry) SetData(v []HostHardwareCapacityDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


