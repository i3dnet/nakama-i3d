# ApplicationBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application build ID | [readonly] 
**Name** | **string** | The name of the application build | 
**ApplicationId** | **NullableString** | The ID of the application. May be null if buildProvisioningStorageType is LegacyFTP | 
**HostCapacityTemplateId** | Pointer to **string** | The ID of the host capacity template | [optional] 
**Type** | **int32** | The application build, list of types can be found in: [GET /application/type](#/Application/get_v3_application_type) | [readonly] 
**Executable** | **string** | The executable name including the path originating from the root directory of the software structure | 
**StartupParameters** | Pointer to **string** | The startup parameters required for an executable | [optional] 
**InstanceDoesReadyCallback** | Pointer to **int32** | Set to 1 if a deployed application instance will inform the platform it&#39;s done with initializing and is ready to accept players. Defaults to 0 if not set, which means an instance&#39;s status will go directly to ONLINE (4) after starting | [optional] 
**InstallId** | Pointer to **int32** | ID of the application install (legacy FTP, will be removed, may not be updated) | [optional] 
**OsId** | **int32** | ID of the operating system, list of operating systems can be found in: [GET /v3/operatingsystem](all#/OperatingSystem/get_v3_operatingsystem) | 
**StopMethod** | Pointer to **NullableInt32** | The stop method that will be used on restart or destroy of an application instance. The default is taken from the value of [GET /v3/application/stopMethod](game-publisher#/Application/getApplicationStopMethods) - any value submitted here overrides that Application default | [optional] 
**StopTimeout** | Pointer to **NullableInt32** | The timeout that will be used on restart or destroy of an application instance. The default is taken from the value of &#x60;Application.stopTimeout&#x60; - any value submitted here overrides that Application default | [optional] 
**CreatedAt** | **int32** | When the application build was created (unix timestamp) | [readonly] 
**Label** | Pointer to [**[]Label**](Label.md) | Custom key/value pairs that can be used for application build selection / identification | [optional] 
**ApplicationBuildFile** | Pointer to [**NullableApplicationBuildFile**](ApplicationBuildFile.md) | The application build file to be used for the application instance | [optional] 
**BuildProvisioningStorageType** | Pointer to **int32** | Type of storage being used [GET /v3/buildProvisioning/storage/registration/types](#/BuildProvisioning/getBuildProvisioningStorageRegistrationTypes). Default: 0 | [optional] 
**RunAsRoot** | **int32** | Controls whether to run this build as root or as a non-privileged user. Warning: This currently only works for builds deployed onto a Linux operating system. Windows builds will always run as administrator * 0: run as non-privileged user * 1: run as root user (default) | 

## Methods

### NewApplicationBuild

`func NewApplicationBuild(id string, name string, applicationId NullableString, type_ int32, executable string, osId int32, createdAt int32, runAsRoot int32, ) *ApplicationBuild`

NewApplicationBuild instantiates a new ApplicationBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBuildWithDefaults

`func NewApplicationBuildWithDefaults() *ApplicationBuild`

NewApplicationBuildWithDefaults instantiates a new ApplicationBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationBuild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationBuild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationBuild) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationBuild) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationId

`func (o *ApplicationBuild) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationBuild) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationBuild) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### SetApplicationIdNil

`func (o *ApplicationBuild) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *ApplicationBuild) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetHostCapacityTemplateId

`func (o *ApplicationBuild) GetHostCapacityTemplateId() string`

GetHostCapacityTemplateId returns the HostCapacityTemplateId field if non-nil, zero value otherwise.

### GetHostCapacityTemplateIdOk

`func (o *ApplicationBuild) GetHostCapacityTemplateIdOk() (*string, bool)`

GetHostCapacityTemplateIdOk returns a tuple with the HostCapacityTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCapacityTemplateId

`func (o *ApplicationBuild) SetHostCapacityTemplateId(v string)`

SetHostCapacityTemplateId sets HostCapacityTemplateId field to given value.

### HasHostCapacityTemplateId

`func (o *ApplicationBuild) HasHostCapacityTemplateId() bool`

HasHostCapacityTemplateId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationBuild) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationBuild) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationBuild) SetType(v int32)`

SetType sets Type field to given value.


### GetExecutable

`func (o *ApplicationBuild) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *ApplicationBuild) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *ApplicationBuild) SetExecutable(v string)`

SetExecutable sets Executable field to given value.


### GetStartupParameters

`func (o *ApplicationBuild) GetStartupParameters() string`

GetStartupParameters returns the StartupParameters field if non-nil, zero value otherwise.

### GetStartupParametersOk

`func (o *ApplicationBuild) GetStartupParametersOk() (*string, bool)`

GetStartupParametersOk returns a tuple with the StartupParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupParameters

`func (o *ApplicationBuild) SetStartupParameters(v string)`

SetStartupParameters sets StartupParameters field to given value.

### HasStartupParameters

`func (o *ApplicationBuild) HasStartupParameters() bool`

HasStartupParameters returns a boolean if a field has been set.

### GetInstanceDoesReadyCallback

`func (o *ApplicationBuild) GetInstanceDoesReadyCallback() int32`

GetInstanceDoesReadyCallback returns the InstanceDoesReadyCallback field if non-nil, zero value otherwise.

### GetInstanceDoesReadyCallbackOk

`func (o *ApplicationBuild) GetInstanceDoesReadyCallbackOk() (*int32, bool)`

GetInstanceDoesReadyCallbackOk returns a tuple with the InstanceDoesReadyCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDoesReadyCallback

`func (o *ApplicationBuild) SetInstanceDoesReadyCallback(v int32)`

SetInstanceDoesReadyCallback sets InstanceDoesReadyCallback field to given value.

### HasInstanceDoesReadyCallback

`func (o *ApplicationBuild) HasInstanceDoesReadyCallback() bool`

HasInstanceDoesReadyCallback returns a boolean if a field has been set.

### GetInstallId

`func (o *ApplicationBuild) GetInstallId() int32`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *ApplicationBuild) GetInstallIdOk() (*int32, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *ApplicationBuild) SetInstallId(v int32)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *ApplicationBuild) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### GetOsId

`func (o *ApplicationBuild) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *ApplicationBuild) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *ApplicationBuild) SetOsId(v int32)`

SetOsId sets OsId field to given value.


### GetStopMethod

`func (o *ApplicationBuild) GetStopMethod() int32`

GetStopMethod returns the StopMethod field if non-nil, zero value otherwise.

### GetStopMethodOk

`func (o *ApplicationBuild) GetStopMethodOk() (*int32, bool)`

GetStopMethodOk returns a tuple with the StopMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopMethod

`func (o *ApplicationBuild) SetStopMethod(v int32)`

SetStopMethod sets StopMethod field to given value.

### HasStopMethod

`func (o *ApplicationBuild) HasStopMethod() bool`

HasStopMethod returns a boolean if a field has been set.

### SetStopMethodNil

`func (o *ApplicationBuild) SetStopMethodNil(b bool)`

 SetStopMethodNil sets the value for StopMethod to be an explicit nil

### UnsetStopMethod
`func (o *ApplicationBuild) UnsetStopMethod()`

UnsetStopMethod ensures that no value is present for StopMethod, not even an explicit nil
### GetStopTimeout

`func (o *ApplicationBuild) GetStopTimeout() int32`

GetStopTimeout returns the StopTimeout field if non-nil, zero value otherwise.

### GetStopTimeoutOk

`func (o *ApplicationBuild) GetStopTimeoutOk() (*int32, bool)`

GetStopTimeoutOk returns a tuple with the StopTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimeout

`func (o *ApplicationBuild) SetStopTimeout(v int32)`

SetStopTimeout sets StopTimeout field to given value.

### HasStopTimeout

`func (o *ApplicationBuild) HasStopTimeout() bool`

HasStopTimeout returns a boolean if a field has been set.

### SetStopTimeoutNil

`func (o *ApplicationBuild) SetStopTimeoutNil(b bool)`

 SetStopTimeoutNil sets the value for StopTimeout to be an explicit nil

### UnsetStopTimeout
`func (o *ApplicationBuild) UnsetStopTimeout()`

UnsetStopTimeout ensures that no value is present for StopTimeout, not even an explicit nil
### GetCreatedAt

`func (o *ApplicationBuild) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationBuild) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationBuild) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabel

`func (o *ApplicationBuild) GetLabel() []Label`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationBuild) GetLabelOk() (*[]Label, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationBuild) SetLabel(v []Label)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApplicationBuild) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetApplicationBuildFile

`func (o *ApplicationBuild) GetApplicationBuildFile() ApplicationBuildFile`

GetApplicationBuildFile returns the ApplicationBuildFile field if non-nil, zero value otherwise.

### GetApplicationBuildFileOk

`func (o *ApplicationBuild) GetApplicationBuildFileOk() (*ApplicationBuildFile, bool)`

GetApplicationBuildFileOk returns a tuple with the ApplicationBuildFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildFile

`func (o *ApplicationBuild) SetApplicationBuildFile(v ApplicationBuildFile)`

SetApplicationBuildFile sets ApplicationBuildFile field to given value.

### HasApplicationBuildFile

`func (o *ApplicationBuild) HasApplicationBuildFile() bool`

HasApplicationBuildFile returns a boolean if a field has been set.

### SetApplicationBuildFileNil

`func (o *ApplicationBuild) SetApplicationBuildFileNil(b bool)`

 SetApplicationBuildFileNil sets the value for ApplicationBuildFile to be an explicit nil

### UnsetApplicationBuildFile
`func (o *ApplicationBuild) UnsetApplicationBuildFile()`

UnsetApplicationBuildFile ensures that no value is present for ApplicationBuildFile, not even an explicit nil
### GetBuildProvisioningStorageType

`func (o *ApplicationBuild) GetBuildProvisioningStorageType() int32`

GetBuildProvisioningStorageType returns the BuildProvisioningStorageType field if non-nil, zero value otherwise.

### GetBuildProvisioningStorageTypeOk

`func (o *ApplicationBuild) GetBuildProvisioningStorageTypeOk() (*int32, bool)`

GetBuildProvisioningStorageTypeOk returns a tuple with the BuildProvisioningStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildProvisioningStorageType

`func (o *ApplicationBuild) SetBuildProvisioningStorageType(v int32)`

SetBuildProvisioningStorageType sets BuildProvisioningStorageType field to given value.

### HasBuildProvisioningStorageType

`func (o *ApplicationBuild) HasBuildProvisioningStorageType() bool`

HasBuildProvisioningStorageType returns a boolean if a field has been set.

### GetRunAsRoot

`func (o *ApplicationBuild) GetRunAsRoot() int32`

GetRunAsRoot returns the RunAsRoot field if non-nil, zero value otherwise.

### GetRunAsRootOk

`func (o *ApplicationBuild) GetRunAsRootOk() (*int32, bool)`

GetRunAsRootOk returns a tuple with the RunAsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsRoot

`func (o *ApplicationBuild) SetRunAsRoot(v int32)`

SetRunAsRoot sets RunAsRoot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


