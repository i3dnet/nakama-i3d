# DependencyDeploymentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | dependency Deployment template ID | [readonly] 
**Name** | **string** | Name of the dependency deployment template | 
**DependencyInstallerBuildId** | **string** | Dependency installer build ID | 
**DependencyUninstallerBuildId** | **string** | Dependency uninstaller build ID | 
**FleetIds** | **[]string** | IDs of the fleets that use this template | [readonly] 
**InUse** | **int32** | Indicates whether this dependency deployment template is used by one or more fleets. A template cannot be deleted as long as this value is 1. | [readonly] 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 

## Methods

### NewDependencyDeploymentTemplate

`func NewDependencyDeploymentTemplate(id string, name string, dependencyInstallerBuildId string, dependencyUninstallerBuildId string, fleetIds []string, inUse int32, createdAt int32, ) *DependencyDeploymentTemplate`

NewDependencyDeploymentTemplate instantiates a new DependencyDeploymentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyDeploymentTemplateWithDefaults

`func NewDependencyDeploymentTemplateWithDefaults() *DependencyDeploymentTemplate`

NewDependencyDeploymentTemplateWithDefaults instantiates a new DependencyDeploymentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DependencyDeploymentTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DependencyDeploymentTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DependencyDeploymentTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DependencyDeploymentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependencyDeploymentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependencyDeploymentTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDependencyInstallerBuildId

`func (o *DependencyDeploymentTemplate) GetDependencyInstallerBuildId() string`

GetDependencyInstallerBuildId returns the DependencyInstallerBuildId field if non-nil, zero value otherwise.

### GetDependencyInstallerBuildIdOk

`func (o *DependencyDeploymentTemplate) GetDependencyInstallerBuildIdOk() (*string, bool)`

GetDependencyInstallerBuildIdOk returns a tuple with the DependencyInstallerBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyInstallerBuildId

`func (o *DependencyDeploymentTemplate) SetDependencyInstallerBuildId(v string)`

SetDependencyInstallerBuildId sets DependencyInstallerBuildId field to given value.


### GetDependencyUninstallerBuildId

`func (o *DependencyDeploymentTemplate) GetDependencyUninstallerBuildId() string`

GetDependencyUninstallerBuildId returns the DependencyUninstallerBuildId field if non-nil, zero value otherwise.

### GetDependencyUninstallerBuildIdOk

`func (o *DependencyDeploymentTemplate) GetDependencyUninstallerBuildIdOk() (*string, bool)`

GetDependencyUninstallerBuildIdOk returns a tuple with the DependencyUninstallerBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyUninstallerBuildId

`func (o *DependencyDeploymentTemplate) SetDependencyUninstallerBuildId(v string)`

SetDependencyUninstallerBuildId sets DependencyUninstallerBuildId field to given value.


### GetFleetIds

`func (o *DependencyDeploymentTemplate) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *DependencyDeploymentTemplate) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *DependencyDeploymentTemplate) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.


### GetInUse

`func (o *DependencyDeploymentTemplate) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DependencyDeploymentTemplate) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DependencyDeploymentTemplate) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetCreatedAt

`func (o *DependencyDeploymentTemplate) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DependencyDeploymentTemplate) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DependencyDeploymentTemplate) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


