# GameDeploymentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Game deployment template ID | [readonly] 
**FleetIds** | **[]string** | IDs of the fleets that use this template | [readonly] 
**Name** | **string** | The public name of your game deployment template | 
**InUse** | **int32** | Indicates whether this game deployment template is used by one or more fleets. A template cannot be deleted as long as this value is 1. | [readonly] 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 
**GameDeploymentTemplateBuild** | [**[]GameDeploymentTemplateBuild**](GameDeploymentTemplateBuild.md) | Application build assigned to the template (can only be one build) | 

## Methods

### NewGameDeploymentTemplate

`func NewGameDeploymentTemplate(id string, fleetIds []string, name string, inUse int32, createdAt int32, gameDeploymentTemplateBuild []GameDeploymentTemplateBuild, ) *GameDeploymentTemplate`

NewGameDeploymentTemplate instantiates a new GameDeploymentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameDeploymentTemplateWithDefaults

`func NewGameDeploymentTemplateWithDefaults() *GameDeploymentTemplate`

NewGameDeploymentTemplateWithDefaults instantiates a new GameDeploymentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameDeploymentTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameDeploymentTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameDeploymentTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetFleetIds

`func (o *GameDeploymentTemplate) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *GameDeploymentTemplate) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *GameDeploymentTemplate) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetName

`func (o *GameDeploymentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameDeploymentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameDeploymentTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetInUse

`func (o *GameDeploymentTemplate) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *GameDeploymentTemplate) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *GameDeploymentTemplate) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetCreatedAt

`func (o *GameDeploymentTemplate) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameDeploymentTemplate) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameDeploymentTemplate) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetGameDeploymentTemplateBuild

`func (o *GameDeploymentTemplate) GetGameDeploymentTemplateBuild() []GameDeploymentTemplateBuild`

GetGameDeploymentTemplateBuild returns the GameDeploymentTemplateBuild field if non-nil, zero value otherwise.

### GetGameDeploymentTemplateBuildOk

`func (o *GameDeploymentTemplate) GetGameDeploymentTemplateBuildOk() (*[]GameDeploymentTemplateBuild, bool)`

GetGameDeploymentTemplateBuildOk returns a tuple with the GameDeploymentTemplateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameDeploymentTemplateBuild

`func (o *GameDeploymentTemplate) SetGameDeploymentTemplateBuild(v []GameDeploymentTemplateBuild)`

SetGameDeploymentTemplateBuild sets GameDeploymentTemplateBuild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


