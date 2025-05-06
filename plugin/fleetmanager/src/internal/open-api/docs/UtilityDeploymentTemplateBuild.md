# UtilityDeploymentTemplateBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationBuildId** | **string** | The Application Build ID | 
**DeployOn** | Pointer to **int32** | Indicates on which host types to install the utilities on: 1) bare metal, 2) VM, 3) both (default value when not specified) | [optional] 
**InstallsPerLocation** | Pointer to **int32** | If &gt; 0, limit the amount of deploys on hosts to this figure, for each active dcLocation. So if this value is 2, this utility will only be deployed on two hosts per dcLocation. If 0 (default value when not specified), it will be deployed on every host in use. | [optional] 

## Methods

### NewUtilityDeploymentTemplateBuild

`func NewUtilityDeploymentTemplateBuild(applicationBuildId string, ) *UtilityDeploymentTemplateBuild`

NewUtilityDeploymentTemplateBuild instantiates a new UtilityDeploymentTemplateBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilityDeploymentTemplateBuildWithDefaults

`func NewUtilityDeploymentTemplateBuildWithDefaults() *UtilityDeploymentTemplateBuild`

NewUtilityDeploymentTemplateBuildWithDefaults instantiates a new UtilityDeploymentTemplateBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationBuildId

`func (o *UtilityDeploymentTemplateBuild) GetApplicationBuildId() string`

GetApplicationBuildId returns the ApplicationBuildId field if non-nil, zero value otherwise.

### GetApplicationBuildIdOk

`func (o *UtilityDeploymentTemplateBuild) GetApplicationBuildIdOk() (*string, bool)`

GetApplicationBuildIdOk returns a tuple with the ApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildId

`func (o *UtilityDeploymentTemplateBuild) SetApplicationBuildId(v string)`

SetApplicationBuildId sets ApplicationBuildId field to given value.


### GetDeployOn

`func (o *UtilityDeploymentTemplateBuild) GetDeployOn() int32`

GetDeployOn returns the DeployOn field if non-nil, zero value otherwise.

### GetDeployOnOk

`func (o *UtilityDeploymentTemplateBuild) GetDeployOnOk() (*int32, bool)`

GetDeployOnOk returns a tuple with the DeployOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployOn

`func (o *UtilityDeploymentTemplateBuild) SetDeployOn(v int32)`

SetDeployOn sets DeployOn field to given value.

### HasDeployOn

`func (o *UtilityDeploymentTemplateBuild) HasDeployOn() bool`

HasDeployOn returns a boolean if a field has been set.

### GetInstallsPerLocation

`func (o *UtilityDeploymentTemplateBuild) GetInstallsPerLocation() int32`

GetInstallsPerLocation returns the InstallsPerLocation field if non-nil, zero value otherwise.

### GetInstallsPerLocationOk

`func (o *UtilityDeploymentTemplateBuild) GetInstallsPerLocationOk() (*int32, bool)`

GetInstallsPerLocationOk returns a tuple with the InstallsPerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallsPerLocation

`func (o *UtilityDeploymentTemplateBuild) SetInstallsPerLocation(v int32)`

SetInstallsPerLocation sets InstallsPerLocation field to given value.

### HasInstallsPerLocation

`func (o *UtilityDeploymentTemplateBuild) HasInstallsPerLocation() bool`

HasInstallsPerLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


