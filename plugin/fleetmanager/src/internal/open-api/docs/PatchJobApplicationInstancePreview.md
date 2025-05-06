# PatchJobApplicationInstancePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentEnvironmentId** | **string** | Deployment environment ID | 
**ApplicationId** | **string** | Application ID | 
**Fleet** | **[]string** | List of fleet ID | 
**ApplicationBuild** | **[]string** | List of application build ID | 

## Methods

### NewPatchJobApplicationInstancePreview

`func NewPatchJobApplicationInstancePreview(deploymentEnvironmentId string, applicationId string, fleet []string, applicationBuild []string, ) *PatchJobApplicationInstancePreview`

NewPatchJobApplicationInstancePreview instantiates a new PatchJobApplicationInstancePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobApplicationInstancePreviewWithDefaults

`func NewPatchJobApplicationInstancePreviewWithDefaults() *PatchJobApplicationInstancePreview`

NewPatchJobApplicationInstancePreviewWithDefaults instantiates a new PatchJobApplicationInstancePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentEnvironmentId

`func (o *PatchJobApplicationInstancePreview) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *PatchJobApplicationInstancePreview) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *PatchJobApplicationInstancePreview) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetApplicationId

`func (o *PatchJobApplicationInstancePreview) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PatchJobApplicationInstancePreview) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PatchJobApplicationInstancePreview) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetFleet

`func (o *PatchJobApplicationInstancePreview) GetFleet() []string`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *PatchJobApplicationInstancePreview) GetFleetOk() (*[]string, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *PatchJobApplicationInstancePreview) SetFleet(v []string)`

SetFleet sets Fleet field to given value.


### GetApplicationBuild

`func (o *PatchJobApplicationInstancePreview) GetApplicationBuild() []string`

GetApplicationBuild returns the ApplicationBuild field if non-nil, zero value otherwise.

### GetApplicationBuildOk

`func (o *PatchJobApplicationInstancePreview) GetApplicationBuildOk() (*[]string, bool)`

GetApplicationBuildOk returns a tuple with the ApplicationBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuild

`func (o *PatchJobApplicationInstancePreview) SetApplicationBuild(v []string)`

SetApplicationBuild sets ApplicationBuild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


