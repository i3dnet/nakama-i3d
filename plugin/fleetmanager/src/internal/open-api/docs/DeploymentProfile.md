# DeploymentProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The deployment profile ID | [readonly] 
**FleetIds** | **[]string** | IDs of the fleets in which this deployment profile is being used | [readonly] 
**Name** | **string** | This deployment profile&#39;s name | 
**Description** | Pointer to **string** | This deployment profile&#39;s description | [optional] 
**StrategyType** | **int32** | Possible values: - 1: Round robin | 
**MinimumCapacity** | **int32** | The minimum amount of game instances that should always be deployed for each deployment region in this deployment profile. Can be overridden per deployment region | 
**MaximumCapacity** | Pointer to **NullableInt32** | The maximum amount of game instances that can be deployed in each region. Can be overridden per region. &#x60;Null&#x60; value means that there is no &#x60;maximumCapacity&#x60; | [optional] 
**BufferValue** | **int32** | Global buffer value, to be applied to individual deployment regions. Can be overridden per region | 
**BufferValueType** | **int32** | Possible values: - 0: Absolute value - 1: Percentage value (must be accompanied by &#x60;bufferValueMin&#x60; and &#x60;bufferValueMax&#x60;) | 
**BufferValueMin** | Pointer to **NullableInt32** | The minimum absolute &#x60;bufferValue&#x60; when using a percentage, to prevent &#x60;bufferValue&#x60; from going too low | [optional] 
**BufferValueMax** | Pointer to **NullableInt32** | The maximum absolute &#x60;bufferValue&#x60; when using a percentage, to prevent &#x60;bufferValue&#x60; from going too high | [optional] 
**MarkedForDeletion** | **int32** | When &#x60;1&#x60;, &#x60;markedForDeletion&#x60; will propagate to all deployment regions of this deployment profile. Scaling down will happen slowly as application instances get removed naturally. After all application instances are removed, the deployment profile is set to inactive | [readonly] 
**InUse** | **int32** | Possible values: - 0: Deployment profile has no active application instances - 1: Deployment profile has active application instance(s) | [readonly] 
**DeploymentRegions** | [**[]DeploymentRegion**](DeploymentRegion.md) | All deployment regions of this deployment profile | [readonly] 
**CreatedAt** | **int32** | Unix timestamp of creation time | [readonly] 

## Methods

### NewDeploymentProfile

`func NewDeploymentProfile(id string, fleetIds []string, name string, strategyType int32, minimumCapacity int32, bufferValue int32, bufferValueType int32, markedForDeletion int32, inUse int32, deploymentRegions []DeploymentRegion, createdAt int32, ) *DeploymentProfile`

NewDeploymentProfile instantiates a new DeploymentProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProfileWithDefaults

`func NewDeploymentProfileWithDefaults() *DeploymentProfile`

NewDeploymentProfileWithDefaults instantiates a new DeploymentProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentProfile) SetId(v string)`

SetId sets Id field to given value.


### GetFleetIds

`func (o *DeploymentProfile) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *DeploymentProfile) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *DeploymentProfile) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetName

`func (o *DeploymentProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DeploymentProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStrategyType

`func (o *DeploymentProfile) GetStrategyType() int32`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *DeploymentProfile) GetStrategyTypeOk() (*int32, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *DeploymentProfile) SetStrategyType(v int32)`

SetStrategyType sets StrategyType field to given value.


### GetMinimumCapacity

`func (o *DeploymentProfile) GetMinimumCapacity() int32`

GetMinimumCapacity returns the MinimumCapacity field if non-nil, zero value otherwise.

### GetMinimumCapacityOk

`func (o *DeploymentProfile) GetMinimumCapacityOk() (*int32, bool)`

GetMinimumCapacityOk returns a tuple with the MinimumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCapacity

`func (o *DeploymentProfile) SetMinimumCapacity(v int32)`

SetMinimumCapacity sets MinimumCapacity field to given value.


### GetMaximumCapacity

`func (o *DeploymentProfile) GetMaximumCapacity() int32`

GetMaximumCapacity returns the MaximumCapacity field if non-nil, zero value otherwise.

### GetMaximumCapacityOk

`func (o *DeploymentProfile) GetMaximumCapacityOk() (*int32, bool)`

GetMaximumCapacityOk returns a tuple with the MaximumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCapacity

`func (o *DeploymentProfile) SetMaximumCapacity(v int32)`

SetMaximumCapacity sets MaximumCapacity field to given value.

### HasMaximumCapacity

`func (o *DeploymentProfile) HasMaximumCapacity() bool`

HasMaximumCapacity returns a boolean if a field has been set.

### SetMaximumCapacityNil

`func (o *DeploymentProfile) SetMaximumCapacityNil(b bool)`

 SetMaximumCapacityNil sets the value for MaximumCapacity to be an explicit nil

### UnsetMaximumCapacity
`func (o *DeploymentProfile) UnsetMaximumCapacity()`

UnsetMaximumCapacity ensures that no value is present for MaximumCapacity, not even an explicit nil
### GetBufferValue

`func (o *DeploymentProfile) GetBufferValue() int32`

GetBufferValue returns the BufferValue field if non-nil, zero value otherwise.

### GetBufferValueOk

`func (o *DeploymentProfile) GetBufferValueOk() (*int32, bool)`

GetBufferValueOk returns a tuple with the BufferValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValue

`func (o *DeploymentProfile) SetBufferValue(v int32)`

SetBufferValue sets BufferValue field to given value.


### GetBufferValueType

`func (o *DeploymentProfile) GetBufferValueType() int32`

GetBufferValueType returns the BufferValueType field if non-nil, zero value otherwise.

### GetBufferValueTypeOk

`func (o *DeploymentProfile) GetBufferValueTypeOk() (*int32, bool)`

GetBufferValueTypeOk returns a tuple with the BufferValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueType

`func (o *DeploymentProfile) SetBufferValueType(v int32)`

SetBufferValueType sets BufferValueType field to given value.


### GetBufferValueMin

`func (o *DeploymentProfile) GetBufferValueMin() int32`

GetBufferValueMin returns the BufferValueMin field if non-nil, zero value otherwise.

### GetBufferValueMinOk

`func (o *DeploymentProfile) GetBufferValueMinOk() (*int32, bool)`

GetBufferValueMinOk returns a tuple with the BufferValueMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueMin

`func (o *DeploymentProfile) SetBufferValueMin(v int32)`

SetBufferValueMin sets BufferValueMin field to given value.

### HasBufferValueMin

`func (o *DeploymentProfile) HasBufferValueMin() bool`

HasBufferValueMin returns a boolean if a field has been set.

### SetBufferValueMinNil

`func (o *DeploymentProfile) SetBufferValueMinNil(b bool)`

 SetBufferValueMinNil sets the value for BufferValueMin to be an explicit nil

### UnsetBufferValueMin
`func (o *DeploymentProfile) UnsetBufferValueMin()`

UnsetBufferValueMin ensures that no value is present for BufferValueMin, not even an explicit nil
### GetBufferValueMax

`func (o *DeploymentProfile) GetBufferValueMax() int32`

GetBufferValueMax returns the BufferValueMax field if non-nil, zero value otherwise.

### GetBufferValueMaxOk

`func (o *DeploymentProfile) GetBufferValueMaxOk() (*int32, bool)`

GetBufferValueMaxOk returns a tuple with the BufferValueMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueMax

`func (o *DeploymentProfile) SetBufferValueMax(v int32)`

SetBufferValueMax sets BufferValueMax field to given value.

### HasBufferValueMax

`func (o *DeploymentProfile) HasBufferValueMax() bool`

HasBufferValueMax returns a boolean if a field has been set.

### SetBufferValueMaxNil

`func (o *DeploymentProfile) SetBufferValueMaxNil(b bool)`

 SetBufferValueMaxNil sets the value for BufferValueMax to be an explicit nil

### UnsetBufferValueMax
`func (o *DeploymentProfile) UnsetBufferValueMax()`

UnsetBufferValueMax ensures that no value is present for BufferValueMax, not even an explicit nil
### GetMarkedForDeletion

`func (o *DeploymentProfile) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *DeploymentProfile) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *DeploymentProfile) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetInUse

`func (o *DeploymentProfile) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DeploymentProfile) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DeploymentProfile) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetDeploymentRegions

`func (o *DeploymentProfile) GetDeploymentRegions() []DeploymentRegion`

GetDeploymentRegions returns the DeploymentRegions field if non-nil, zero value otherwise.

### GetDeploymentRegionsOk

`func (o *DeploymentProfile) GetDeploymentRegionsOk() (*[]DeploymentRegion, bool)`

GetDeploymentRegionsOk returns a tuple with the DeploymentRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegions

`func (o *DeploymentProfile) SetDeploymentRegions(v []DeploymentRegion)`

SetDeploymentRegions sets DeploymentRegions field to given value.


### GetCreatedAt

`func (o *DeploymentProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


