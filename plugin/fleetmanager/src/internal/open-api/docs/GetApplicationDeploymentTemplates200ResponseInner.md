# GetApplicationDeploymentTemplates200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | dependency Deployment template ID | [readonly] 
**FleetIds** | **[]string** | IDs of the fleets that use this template | [readonly] 
**Name** | **string** | Name of the dependency deployment template | 
**InUse** | **int32** | Indicates whether this dependency deployment template is used by one or more fleets. A template cannot be deleted as long as this value is 1. | [readonly] 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 
**GameDeploymentTemplateBuild** | [**[]GameDeploymentTemplateBuild**](GameDeploymentTemplateBuild.md) | Application build assigned to the template (can only be one build) | 
**UtilityDeploymentTemplateBuild** | [**[]UtilityDeploymentTemplateBuild**](UtilityDeploymentTemplateBuild.md) | A List of utility deployment template build&#39;s | 
**DependencyInstallerBuildId** | **string** | Dependency installer build ID | 
**DependencyUninstallerBuildId** | **string** | Dependency uninstaller build ID | 

## Methods

### NewGetApplicationDeploymentTemplates200ResponseInner

`func NewGetApplicationDeploymentTemplates200ResponseInner(id string, fleetIds []string, name string, inUse int32, createdAt int32, gameDeploymentTemplateBuild []GameDeploymentTemplateBuild, utilityDeploymentTemplateBuild []UtilityDeploymentTemplateBuild, dependencyInstallerBuildId string, dependencyUninstallerBuildId string, ) *GetApplicationDeploymentTemplates200ResponseInner`

NewGetApplicationDeploymentTemplates200ResponseInner instantiates a new GetApplicationDeploymentTemplates200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationDeploymentTemplates200ResponseInnerWithDefaults

`func NewGetApplicationDeploymentTemplates200ResponseInnerWithDefaults() *GetApplicationDeploymentTemplates200ResponseInner`

NewGetApplicationDeploymentTemplates200ResponseInnerWithDefaults instantiates a new GetApplicationDeploymentTemplates200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetFleetIds

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetName

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetInUse

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetCreatedAt

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetGameDeploymentTemplateBuild

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetGameDeploymentTemplateBuild() []GameDeploymentTemplateBuild`

GetGameDeploymentTemplateBuild returns the GameDeploymentTemplateBuild field if non-nil, zero value otherwise.

### GetGameDeploymentTemplateBuildOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetGameDeploymentTemplateBuildOk() (*[]GameDeploymentTemplateBuild, bool)`

GetGameDeploymentTemplateBuildOk returns a tuple with the GameDeploymentTemplateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameDeploymentTemplateBuild

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetGameDeploymentTemplateBuild(v []GameDeploymentTemplateBuild)`

SetGameDeploymentTemplateBuild sets GameDeploymentTemplateBuild field to given value.


### GetUtilityDeploymentTemplateBuild

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetUtilityDeploymentTemplateBuild() []UtilityDeploymentTemplateBuild`

GetUtilityDeploymentTemplateBuild returns the UtilityDeploymentTemplateBuild field if non-nil, zero value otherwise.

### GetUtilityDeploymentTemplateBuildOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetUtilityDeploymentTemplateBuildOk() (*[]UtilityDeploymentTemplateBuild, bool)`

GetUtilityDeploymentTemplateBuildOk returns a tuple with the UtilityDeploymentTemplateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityDeploymentTemplateBuild

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetUtilityDeploymentTemplateBuild(v []UtilityDeploymentTemplateBuild)`

SetUtilityDeploymentTemplateBuild sets UtilityDeploymentTemplateBuild field to given value.


### GetDependencyInstallerBuildId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetDependencyInstallerBuildId() string`

GetDependencyInstallerBuildId returns the DependencyInstallerBuildId field if non-nil, zero value otherwise.

### GetDependencyInstallerBuildIdOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetDependencyInstallerBuildIdOk() (*string, bool)`

GetDependencyInstallerBuildIdOk returns a tuple with the DependencyInstallerBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyInstallerBuildId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetDependencyInstallerBuildId(v string)`

SetDependencyInstallerBuildId sets DependencyInstallerBuildId field to given value.


### GetDependencyUninstallerBuildId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetDependencyUninstallerBuildId() string`

GetDependencyUninstallerBuildId returns the DependencyUninstallerBuildId field if non-nil, zero value otherwise.

### GetDependencyUninstallerBuildIdOk

`func (o *GetApplicationDeploymentTemplates200ResponseInner) GetDependencyUninstallerBuildIdOk() (*string, bool)`

GetDependencyUninstallerBuildIdOk returns a tuple with the DependencyUninstallerBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyUninstallerBuildId

`func (o *GetApplicationDeploymentTemplates200ResponseInner) SetDependencyUninstallerBuildId(v string)`

SetDependencyUninstallerBuildId sets DependencyUninstallerBuildId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


