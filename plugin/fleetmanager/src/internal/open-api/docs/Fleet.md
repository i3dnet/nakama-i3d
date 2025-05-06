# Fleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Fleet ID | [readonly] 
**Name** | **string** | The name of the fleet | 
**DeploymentEnvironmentId** | **string** | Deployment environment ID to which the fleet is assigned | 
**DeploymentProfileId** | Pointer to **string** | The &#x60;DeploymentProfileTemplate&#x60;&#39;s ID, used for scaling, that is attached to this &#x60;Fleet&#x60; | [optional] 
**GameDeploymentTemplateId** | Pointer to **string** | The &#x60;GameDeploymentTemplate&#x60;&#39;s ID that is attached to this &#x60;Fleet&#x60; | [optional] 
**UtilityDeploymentTemplateId** | Pointer to **string** | The &#x60;UtilityDeploymentTemplate&#x60;&#39;s ID that is attached to this &#x60;Fleet&#x60; | [optional] 
**DependencyDeploymentTemplateId** | Pointer to **string** | The &#x60;DependencyDeploymentTemplate&#x60;&#39;s ID that is attached to this &#x60;Fleet&#x60; | [optional] 
**HostCapacityTemplateId** | Pointer to **string** | The &#x60;HostCapacityTemplate&#x60;&#39;s ID that is attached to this &#x60;Fleet&#x60; | [optional] 
**OperationalStatus** | **int32** | 0: Manual, 1: Automatic deployment enabled, 2: Automatic scaling enabled (implies automatic deployment) | [readonly] 
**IsGameLift** | **bool** | When it is active it cannot be disabled anymore | 

## Methods

### NewFleet

`func NewFleet(id string, name string, deploymentEnvironmentId string, operationalStatus int32, isGameLift bool, ) *Fleet`

NewFleet instantiates a new Fleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetWithDefaults

`func NewFleetWithDefaults() *Fleet`

NewFleetWithDefaults instantiates a new Fleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fleet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fleet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fleet) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Fleet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fleet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fleet) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentEnvironmentId

`func (o *Fleet) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *Fleet) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *Fleet) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetDeploymentProfileId

`func (o *Fleet) GetDeploymentProfileId() string`

GetDeploymentProfileId returns the DeploymentProfileId field if non-nil, zero value otherwise.

### GetDeploymentProfileIdOk

`func (o *Fleet) GetDeploymentProfileIdOk() (*string, bool)`

GetDeploymentProfileIdOk returns a tuple with the DeploymentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProfileId

`func (o *Fleet) SetDeploymentProfileId(v string)`

SetDeploymentProfileId sets DeploymentProfileId field to given value.

### HasDeploymentProfileId

`func (o *Fleet) HasDeploymentProfileId() bool`

HasDeploymentProfileId returns a boolean if a field has been set.

### GetGameDeploymentTemplateId

`func (o *Fleet) GetGameDeploymentTemplateId() string`

GetGameDeploymentTemplateId returns the GameDeploymentTemplateId field if non-nil, zero value otherwise.

### GetGameDeploymentTemplateIdOk

`func (o *Fleet) GetGameDeploymentTemplateIdOk() (*string, bool)`

GetGameDeploymentTemplateIdOk returns a tuple with the GameDeploymentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameDeploymentTemplateId

`func (o *Fleet) SetGameDeploymentTemplateId(v string)`

SetGameDeploymentTemplateId sets GameDeploymentTemplateId field to given value.

### HasGameDeploymentTemplateId

`func (o *Fleet) HasGameDeploymentTemplateId() bool`

HasGameDeploymentTemplateId returns a boolean if a field has been set.

### GetUtilityDeploymentTemplateId

`func (o *Fleet) GetUtilityDeploymentTemplateId() string`

GetUtilityDeploymentTemplateId returns the UtilityDeploymentTemplateId field if non-nil, zero value otherwise.

### GetUtilityDeploymentTemplateIdOk

`func (o *Fleet) GetUtilityDeploymentTemplateIdOk() (*string, bool)`

GetUtilityDeploymentTemplateIdOk returns a tuple with the UtilityDeploymentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityDeploymentTemplateId

`func (o *Fleet) SetUtilityDeploymentTemplateId(v string)`

SetUtilityDeploymentTemplateId sets UtilityDeploymentTemplateId field to given value.

### HasUtilityDeploymentTemplateId

`func (o *Fleet) HasUtilityDeploymentTemplateId() bool`

HasUtilityDeploymentTemplateId returns a boolean if a field has been set.

### GetDependencyDeploymentTemplateId

`func (o *Fleet) GetDependencyDeploymentTemplateId() string`

GetDependencyDeploymentTemplateId returns the DependencyDeploymentTemplateId field if non-nil, zero value otherwise.

### GetDependencyDeploymentTemplateIdOk

`func (o *Fleet) GetDependencyDeploymentTemplateIdOk() (*string, bool)`

GetDependencyDeploymentTemplateIdOk returns a tuple with the DependencyDeploymentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDeploymentTemplateId

`func (o *Fleet) SetDependencyDeploymentTemplateId(v string)`

SetDependencyDeploymentTemplateId sets DependencyDeploymentTemplateId field to given value.

### HasDependencyDeploymentTemplateId

`func (o *Fleet) HasDependencyDeploymentTemplateId() bool`

HasDependencyDeploymentTemplateId returns a boolean if a field has been set.

### GetHostCapacityTemplateId

`func (o *Fleet) GetHostCapacityTemplateId() string`

GetHostCapacityTemplateId returns the HostCapacityTemplateId field if non-nil, zero value otherwise.

### GetHostCapacityTemplateIdOk

`func (o *Fleet) GetHostCapacityTemplateIdOk() (*string, bool)`

GetHostCapacityTemplateIdOk returns a tuple with the HostCapacityTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCapacityTemplateId

`func (o *Fleet) SetHostCapacityTemplateId(v string)`

SetHostCapacityTemplateId sets HostCapacityTemplateId field to given value.

### HasHostCapacityTemplateId

`func (o *Fleet) HasHostCapacityTemplateId() bool`

HasHostCapacityTemplateId returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *Fleet) GetOperationalStatus() int32`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *Fleet) GetOperationalStatusOk() (*int32, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *Fleet) SetOperationalStatus(v int32)`

SetOperationalStatus sets OperationalStatus field to given value.


### GetIsGameLift

`func (o *Fleet) GetIsGameLift() bool`

GetIsGameLift returns the IsGameLift field if non-nil, zero value otherwise.

### GetIsGameLiftOk

`func (o *Fleet) GetIsGameLiftOk() (*bool, bool)`

GetIsGameLiftOk returns a tuple with the IsGameLift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGameLift

`func (o *Fleet) SetIsGameLift(v bool)`

SetIsGameLift sets IsGameLift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


