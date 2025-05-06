# UtilityDeploymentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Utility deployment template ID | [readonly] 
**FleetIds** | **[]string** | IDs of the fleets that use this template | [readonly] 
**Name** | **string** | The name of your utility deployment template | 
**InUse** | **int32** | Indicates whether this utility deployment template is used by one or more fleets. A template cannot be deleted as long as this value is 1. | [readonly] 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 
**UtilityDeploymentTemplateBuild** | [**[]UtilityDeploymentTemplateBuild**](UtilityDeploymentTemplateBuild.md) | A List of utility deployment template build&#39;s | 

## Methods

### NewUtilityDeploymentTemplate

`func NewUtilityDeploymentTemplate(id string, fleetIds []string, name string, inUse int32, createdAt int32, utilityDeploymentTemplateBuild []UtilityDeploymentTemplateBuild, ) *UtilityDeploymentTemplate`

NewUtilityDeploymentTemplate instantiates a new UtilityDeploymentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilityDeploymentTemplateWithDefaults

`func NewUtilityDeploymentTemplateWithDefaults() *UtilityDeploymentTemplate`

NewUtilityDeploymentTemplateWithDefaults instantiates a new UtilityDeploymentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UtilityDeploymentTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UtilityDeploymentTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UtilityDeploymentTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetFleetIds

`func (o *UtilityDeploymentTemplate) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *UtilityDeploymentTemplate) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *UtilityDeploymentTemplate) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetName

`func (o *UtilityDeploymentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UtilityDeploymentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UtilityDeploymentTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetInUse

`func (o *UtilityDeploymentTemplate) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *UtilityDeploymentTemplate) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *UtilityDeploymentTemplate) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetCreatedAt

`func (o *UtilityDeploymentTemplate) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UtilityDeploymentTemplate) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UtilityDeploymentTemplate) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUtilityDeploymentTemplateBuild

`func (o *UtilityDeploymentTemplate) GetUtilityDeploymentTemplateBuild() []UtilityDeploymentTemplateBuild`

GetUtilityDeploymentTemplateBuild returns the UtilityDeploymentTemplateBuild field if non-nil, zero value otherwise.

### GetUtilityDeploymentTemplateBuildOk

`func (o *UtilityDeploymentTemplate) GetUtilityDeploymentTemplateBuildOk() (*[]UtilityDeploymentTemplateBuild, bool)`

GetUtilityDeploymentTemplateBuildOk returns a tuple with the UtilityDeploymentTemplateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityDeploymentTemplateBuild

`func (o *UtilityDeploymentTemplate) SetUtilityDeploymentTemplateBuild(v []UtilityDeploymentTemplateBuild)`

SetUtilityDeploymentTemplateBuild sets UtilityDeploymentTemplateBuild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


