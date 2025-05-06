# ApplicationInstanceDeploymentProfileStatusTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **NullableString** | Unit of data frequency, which is always &#x60;null&#x60; in this endpoint. This property is there for consistency across all telemetry endpoints. It&#39;s &#x60;null&#x60; because the data is not aggregated | [readonly] 
**Data** | [**[]ApplicationInstanceDeploymentProfileStatusDataTelemetry**](ApplicationInstanceDeploymentProfileStatusDataTelemetry.md) | Contains the data object by host type | [readonly] 

## Methods

### NewApplicationInstanceDeploymentProfileStatusTelemetry

`func NewApplicationInstanceDeploymentProfileStatusTelemetry(unit NullableString, data []ApplicationInstanceDeploymentProfileStatusDataTelemetry, ) *ApplicationInstanceDeploymentProfileStatusTelemetry`

NewApplicationInstanceDeploymentProfileStatusTelemetry instantiates a new ApplicationInstanceDeploymentProfileStatusTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceDeploymentProfileStatusTelemetryWithDefaults

`func NewApplicationInstanceDeploymentProfileStatusTelemetryWithDefaults() *ApplicationInstanceDeploymentProfileStatusTelemetry`

NewApplicationInstanceDeploymentProfileStatusTelemetryWithDefaults instantiates a new ApplicationInstanceDeploymentProfileStatusTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetData

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) GetData() []ApplicationInstanceDeploymentProfileStatusDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) GetDataOk() (*[]ApplicationInstanceDeploymentProfileStatusDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationInstanceDeploymentProfileStatusTelemetry) SetData(v []ApplicationInstanceDeploymentProfileStatusDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


