# ApplicationInstanceHostCapacityTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BareMetal** | **int32** | Percentage of game instances on the bare metal, based on the amount set in the host capacity template | [readonly] 
**FlexMetal** | **int32** | Percentage of game instances on the flex metal, based on the amount set in the host capacity template | [readonly] 
**Cloud** | **int32** | Percentage of game instances in the cloud, based on the amount set in the host capacity template | [readonly] 

## Methods

### NewApplicationInstanceHostCapacityTelemetry

`func NewApplicationInstanceHostCapacityTelemetry(bareMetal int32, flexMetal int32, cloud int32, ) *ApplicationInstanceHostCapacityTelemetry`

NewApplicationInstanceHostCapacityTelemetry instantiates a new ApplicationInstanceHostCapacityTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceHostCapacityTelemetryWithDefaults

`func NewApplicationInstanceHostCapacityTelemetryWithDefaults() *ApplicationInstanceHostCapacityTelemetry`

NewApplicationInstanceHostCapacityTelemetryWithDefaults instantiates a new ApplicationInstanceHostCapacityTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBareMetal

`func (o *ApplicationInstanceHostCapacityTelemetry) GetBareMetal() int32`

GetBareMetal returns the BareMetal field if non-nil, zero value otherwise.

### GetBareMetalOk

`func (o *ApplicationInstanceHostCapacityTelemetry) GetBareMetalOk() (*int32, bool)`

GetBareMetalOk returns a tuple with the BareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetal

`func (o *ApplicationInstanceHostCapacityTelemetry) SetBareMetal(v int32)`

SetBareMetal sets BareMetal field to given value.


### GetFlexMetal

`func (o *ApplicationInstanceHostCapacityTelemetry) GetFlexMetal() int32`

GetFlexMetal returns the FlexMetal field if non-nil, zero value otherwise.

### GetFlexMetalOk

`func (o *ApplicationInstanceHostCapacityTelemetry) GetFlexMetalOk() (*int32, bool)`

GetFlexMetalOk returns a tuple with the FlexMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMetal

`func (o *ApplicationInstanceHostCapacityTelemetry) SetFlexMetal(v int32)`

SetFlexMetal sets FlexMetal field to given value.


### GetCloud

`func (o *ApplicationInstanceHostCapacityTelemetry) GetCloud() int32`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ApplicationInstanceHostCapacityTelemetry) GetCloudOk() (*int32, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ApplicationInstanceHostCapacityTelemetry) SetCloud(v int32)`

SetCloud sets Cloud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


