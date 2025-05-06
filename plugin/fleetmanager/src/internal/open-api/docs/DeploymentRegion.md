# DeploymentRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This deployment region&#39;s ID | [readonly] 
**Name** | **string** | The name of the deployment region | 
**I3dDcLocationIds** | **[]int32** | All the i3D.net DC location IDs configured for this deployment region. You can get available IDs via &#x60;GET /cloud/dcLocation&#x60; | 
**I3dDcLocationIdsToBeRemoved** | Pointer to **[]int32** | i3D.net data center location IDs scheduled for removal from this region | [optional] 
**MinimumCapacity** | Pointer to **NullableInt32** | The minimum amount of game instances that should always be deployed in this region. &#x60;Null&#x60; value means that the &#x60;minimumCapacity&#x60; provided in the deployment profile will persist | [optional] 
**MaximumCapacity** | Pointer to **NullableInt32** | The maximum amount of game instances that can be deployed in this region. &#x60;Null&#x60; value means that the &#x60;maximumCapacity&#x60; provided in the deployment profile will persist | [optional] 
**BufferValue** | Pointer to **NullableInt32** | Override for &#x60;DeploymentProfile.bufferValue&#x60;, to be applied only to this deployment region. &#x60;Null&#x60; value means that the &#x60;bufferValue&#x60; provided in the deployment profile will persist | [optional] 
**BufferValueType** | Pointer to **NullableInt32** | Override for &#x60;DeploymentProfile.bufferValueType&#x60;. Possible values: - 0: Absolute value - 1: Percentage value (must be accompanied by &#x60;bufferValueMin&#x60; and &#x60;bufferValueMax&#x60;) &#x60;Null&#x60; value means that the &#x60;bufferValueType&#x60; provided in the deployment profile will persist | [optional] 
**BufferValueMin** | Pointer to **NullableInt32** | Override for &#x60;DeploymentProfile.bufferValueMin&#x60;. The minimum absolute &#x60;bufferValue&#x60; when using a percentage, to prevent &#x60;bufferValue&#x60; from going too low. &#x60;Null&#x60; value means that the &#x60;bufferValueMin&#x60; provided in the deployment profile will persist. Note: if this value is supplied (not &#x60;null&#x60;) with your request, &#x60;bufferValue&#x60; and &#x60;bufferValueType&#x60; become required parameters. See &#x60;POST /deploymentProfile/{deploymentProfileId}&#x60; for more information | [optional] 
**BufferValueMax** | Pointer to **NullableInt32** | Override for &#x60;DeploymentProfile.bufferValueMax&#x60;. The maximum absolute &#x60;bufferValue&#x60; when using a percentage, to prevent &#x60;bufferValue&#x60; from going too high. &#x60;Null&#x60; value means that the &#x60;bufferValueMax&#x60; provided in the deployment profile will persist. Note: if this value is supplied (not &#x60;null&#x60;) with your request, &#x60;bufferValue&#x60; and &#x60;bufferValueType&#x60; become required parameters. See &#x60;POST /deploymentProfile/{deploymentProfileId}&#x60; for more information | [optional] 
**StrategyType** | Pointer to **NullableInt32** | Override value for &#x60;DeploymentProfile.strategyType&#x60;. Possible values: - &#x60;0&#x60;: Use default - &#x60;1&#x60;: Round robin - &#x60;null&#x60;: The &#x60;strategyType&#x60; provided in the deployment profile will persist | [optional] 
**MarkedForDeletion** | **int32** | If set to &#x60;1&#x60;, all application instances will be gracefully removed inside the deployment region. Afterwards this deployment region will be set to inactive | [readonly] 
**Containers** | Pointer to [**[]DeploymentContainer**](DeploymentContainer.md) |  | [optional] 
**InUse** | **int32** | Possible values: - 0: Deployment region has no active application instances - 1: Deployment region has active application instance(s) | [readonly] 

## Methods

### NewDeploymentRegion

`func NewDeploymentRegion(id string, name string, i3dDcLocationIds []int32, markedForDeletion int32, inUse int32, ) *DeploymentRegion`

NewDeploymentRegion instantiates a new DeploymentRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRegionWithDefaults

`func NewDeploymentRegionWithDefaults() *DeploymentRegion`

NewDeploymentRegionWithDefaults instantiates a new DeploymentRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentRegion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentRegion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentRegion) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeploymentRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentRegion) SetName(v string)`

SetName sets Name field to given value.


### GetI3dDcLocationIds

`func (o *DeploymentRegion) GetI3dDcLocationIds() []int32`

GetI3dDcLocationIds returns the I3dDcLocationIds field if non-nil, zero value otherwise.

### GetI3dDcLocationIdsOk

`func (o *DeploymentRegion) GetI3dDcLocationIdsOk() (*[]int32, bool)`

GetI3dDcLocationIdsOk returns a tuple with the I3dDcLocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI3dDcLocationIds

`func (o *DeploymentRegion) SetI3dDcLocationIds(v []int32)`

SetI3dDcLocationIds sets I3dDcLocationIds field to given value.


### GetI3dDcLocationIdsToBeRemoved

`func (o *DeploymentRegion) GetI3dDcLocationIdsToBeRemoved() []int32`

GetI3dDcLocationIdsToBeRemoved returns the I3dDcLocationIdsToBeRemoved field if non-nil, zero value otherwise.

### GetI3dDcLocationIdsToBeRemovedOk

`func (o *DeploymentRegion) GetI3dDcLocationIdsToBeRemovedOk() (*[]int32, bool)`

GetI3dDcLocationIdsToBeRemovedOk returns a tuple with the I3dDcLocationIdsToBeRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI3dDcLocationIdsToBeRemoved

`func (o *DeploymentRegion) SetI3dDcLocationIdsToBeRemoved(v []int32)`

SetI3dDcLocationIdsToBeRemoved sets I3dDcLocationIdsToBeRemoved field to given value.

### HasI3dDcLocationIdsToBeRemoved

`func (o *DeploymentRegion) HasI3dDcLocationIdsToBeRemoved() bool`

HasI3dDcLocationIdsToBeRemoved returns a boolean if a field has been set.

### GetMinimumCapacity

`func (o *DeploymentRegion) GetMinimumCapacity() int32`

GetMinimumCapacity returns the MinimumCapacity field if non-nil, zero value otherwise.

### GetMinimumCapacityOk

`func (o *DeploymentRegion) GetMinimumCapacityOk() (*int32, bool)`

GetMinimumCapacityOk returns a tuple with the MinimumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCapacity

`func (o *DeploymentRegion) SetMinimumCapacity(v int32)`

SetMinimumCapacity sets MinimumCapacity field to given value.

### HasMinimumCapacity

`func (o *DeploymentRegion) HasMinimumCapacity() bool`

HasMinimumCapacity returns a boolean if a field has been set.

### SetMinimumCapacityNil

`func (o *DeploymentRegion) SetMinimumCapacityNil(b bool)`

 SetMinimumCapacityNil sets the value for MinimumCapacity to be an explicit nil

### UnsetMinimumCapacity
`func (o *DeploymentRegion) UnsetMinimumCapacity()`

UnsetMinimumCapacity ensures that no value is present for MinimumCapacity, not even an explicit nil
### GetMaximumCapacity

`func (o *DeploymentRegion) GetMaximumCapacity() int32`

GetMaximumCapacity returns the MaximumCapacity field if non-nil, zero value otherwise.

### GetMaximumCapacityOk

`func (o *DeploymentRegion) GetMaximumCapacityOk() (*int32, bool)`

GetMaximumCapacityOk returns a tuple with the MaximumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCapacity

`func (o *DeploymentRegion) SetMaximumCapacity(v int32)`

SetMaximumCapacity sets MaximumCapacity field to given value.

### HasMaximumCapacity

`func (o *DeploymentRegion) HasMaximumCapacity() bool`

HasMaximumCapacity returns a boolean if a field has been set.

### SetMaximumCapacityNil

`func (o *DeploymentRegion) SetMaximumCapacityNil(b bool)`

 SetMaximumCapacityNil sets the value for MaximumCapacity to be an explicit nil

### UnsetMaximumCapacity
`func (o *DeploymentRegion) UnsetMaximumCapacity()`

UnsetMaximumCapacity ensures that no value is present for MaximumCapacity, not even an explicit nil
### GetBufferValue

`func (o *DeploymentRegion) GetBufferValue() int32`

GetBufferValue returns the BufferValue field if non-nil, zero value otherwise.

### GetBufferValueOk

`func (o *DeploymentRegion) GetBufferValueOk() (*int32, bool)`

GetBufferValueOk returns a tuple with the BufferValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValue

`func (o *DeploymentRegion) SetBufferValue(v int32)`

SetBufferValue sets BufferValue field to given value.

### HasBufferValue

`func (o *DeploymentRegion) HasBufferValue() bool`

HasBufferValue returns a boolean if a field has been set.

### SetBufferValueNil

`func (o *DeploymentRegion) SetBufferValueNil(b bool)`

 SetBufferValueNil sets the value for BufferValue to be an explicit nil

### UnsetBufferValue
`func (o *DeploymentRegion) UnsetBufferValue()`

UnsetBufferValue ensures that no value is present for BufferValue, not even an explicit nil
### GetBufferValueType

`func (o *DeploymentRegion) GetBufferValueType() int32`

GetBufferValueType returns the BufferValueType field if non-nil, zero value otherwise.

### GetBufferValueTypeOk

`func (o *DeploymentRegion) GetBufferValueTypeOk() (*int32, bool)`

GetBufferValueTypeOk returns a tuple with the BufferValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueType

`func (o *DeploymentRegion) SetBufferValueType(v int32)`

SetBufferValueType sets BufferValueType field to given value.

### HasBufferValueType

`func (o *DeploymentRegion) HasBufferValueType() bool`

HasBufferValueType returns a boolean if a field has been set.

### SetBufferValueTypeNil

`func (o *DeploymentRegion) SetBufferValueTypeNil(b bool)`

 SetBufferValueTypeNil sets the value for BufferValueType to be an explicit nil

### UnsetBufferValueType
`func (o *DeploymentRegion) UnsetBufferValueType()`

UnsetBufferValueType ensures that no value is present for BufferValueType, not even an explicit nil
### GetBufferValueMin

`func (o *DeploymentRegion) GetBufferValueMin() int32`

GetBufferValueMin returns the BufferValueMin field if non-nil, zero value otherwise.

### GetBufferValueMinOk

`func (o *DeploymentRegion) GetBufferValueMinOk() (*int32, bool)`

GetBufferValueMinOk returns a tuple with the BufferValueMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueMin

`func (o *DeploymentRegion) SetBufferValueMin(v int32)`

SetBufferValueMin sets BufferValueMin field to given value.

### HasBufferValueMin

`func (o *DeploymentRegion) HasBufferValueMin() bool`

HasBufferValueMin returns a boolean if a field has been set.

### SetBufferValueMinNil

`func (o *DeploymentRegion) SetBufferValueMinNil(b bool)`

 SetBufferValueMinNil sets the value for BufferValueMin to be an explicit nil

### UnsetBufferValueMin
`func (o *DeploymentRegion) UnsetBufferValueMin()`

UnsetBufferValueMin ensures that no value is present for BufferValueMin, not even an explicit nil
### GetBufferValueMax

`func (o *DeploymentRegion) GetBufferValueMax() int32`

GetBufferValueMax returns the BufferValueMax field if non-nil, zero value otherwise.

### GetBufferValueMaxOk

`func (o *DeploymentRegion) GetBufferValueMaxOk() (*int32, bool)`

GetBufferValueMaxOk returns a tuple with the BufferValueMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferValueMax

`func (o *DeploymentRegion) SetBufferValueMax(v int32)`

SetBufferValueMax sets BufferValueMax field to given value.

### HasBufferValueMax

`func (o *DeploymentRegion) HasBufferValueMax() bool`

HasBufferValueMax returns a boolean if a field has been set.

### SetBufferValueMaxNil

`func (o *DeploymentRegion) SetBufferValueMaxNil(b bool)`

 SetBufferValueMaxNil sets the value for BufferValueMax to be an explicit nil

### UnsetBufferValueMax
`func (o *DeploymentRegion) UnsetBufferValueMax()`

UnsetBufferValueMax ensures that no value is present for BufferValueMax, not even an explicit nil
### GetStrategyType

`func (o *DeploymentRegion) GetStrategyType() int32`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *DeploymentRegion) GetStrategyTypeOk() (*int32, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *DeploymentRegion) SetStrategyType(v int32)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *DeploymentRegion) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### SetStrategyTypeNil

`func (o *DeploymentRegion) SetStrategyTypeNil(b bool)`

 SetStrategyTypeNil sets the value for StrategyType to be an explicit nil

### UnsetStrategyType
`func (o *DeploymentRegion) UnsetStrategyType()`

UnsetStrategyType ensures that no value is present for StrategyType, not even an explicit nil
### GetMarkedForDeletion

`func (o *DeploymentRegion) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *DeploymentRegion) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *DeploymentRegion) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetContainers

`func (o *DeploymentRegion) GetContainers() []DeploymentContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeploymentRegion) GetContainersOk() (*[]DeploymentContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeploymentRegion) SetContainers(v []DeploymentContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeploymentRegion) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetInUse

`func (o *DeploymentRegion) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DeploymentRegion) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DeploymentRegion) SetInUse(v int32)`

SetInUse sets InUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


