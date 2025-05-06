# CrashEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Crash event ID | [readonly] 
**HostId** | **int32** | Host ID (Dedicated Server ID) | [readonly] 
**DeploymentEnvironmentId** | **string** | Deployment environment ID to which the fleet is assigned | [readonly] 
**FleetId** | **int32** | ID of the fleet | [readonly] 
**ApplicationId** | **string** | Application install assigned to applicationId | [readonly] 
**ApplicationBuildId** | **string** | ID of the application build | [readonly] 
**ApplicationBuildExecutable** | **string** | Executable of the application build | [readonly] 
**ApplicationBuildStartParams** | **string** | Start parameters of the application build | [readonly] 
**ExitCode** | **int32** | Exit code of the application instance | [readonly] 
**DeploymentEnvironment** | [**DeploymentEnvironment**](DeploymentEnvironment.md) | Snapshot of deployment environment | [readonly] 
**Fleet** | [**Fleet**](Fleet.md) | Snapshot of fleet | [readonly] 
**ApplicationInstance** | [**ApplicationInstance**](ApplicationInstance.md) | Snapshot of an application instance | [readonly] 
**ApplicationInstanceConfiguration** | [**[]ApplicationInstanceConfiguration**](ApplicationInstanceConfiguration.md) | Snapshot of an application instance | [readonly] 
**ApplicationInstanceMetadata** | [**[]Metadata**](Metadata.md) | Snapshot of an application instance metadata | [readonly] 
**CreatedAt** | **int32** | Application install created at timestamp | [readonly] 

## Methods

### NewCrashEvent

`func NewCrashEvent(id string, hostId int32, deploymentEnvironmentId string, fleetId int32, applicationId string, applicationBuildId string, applicationBuildExecutable string, applicationBuildStartParams string, exitCode int32, deploymentEnvironment DeploymentEnvironment, fleet Fleet, applicationInstance ApplicationInstance, applicationInstanceConfiguration []ApplicationInstanceConfiguration, applicationInstanceMetadata []Metadata, createdAt int32, ) *CrashEvent`

NewCrashEvent instantiates a new CrashEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrashEventWithDefaults

`func NewCrashEventWithDefaults() *CrashEvent`

NewCrashEventWithDefaults instantiates a new CrashEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CrashEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrashEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrashEvent) SetId(v string)`

SetId sets Id field to given value.


### GetHostId

`func (o *CrashEvent) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CrashEvent) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CrashEvent) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetDeploymentEnvironmentId

`func (o *CrashEvent) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *CrashEvent) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *CrashEvent) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetFleetId

`func (o *CrashEvent) GetFleetId() int32`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *CrashEvent) GetFleetIdOk() (*int32, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *CrashEvent) SetFleetId(v int32)`

SetFleetId sets FleetId field to given value.


### GetApplicationId

`func (o *CrashEvent) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CrashEvent) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CrashEvent) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationBuildId

`func (o *CrashEvent) GetApplicationBuildId() string`

GetApplicationBuildId returns the ApplicationBuildId field if non-nil, zero value otherwise.

### GetApplicationBuildIdOk

`func (o *CrashEvent) GetApplicationBuildIdOk() (*string, bool)`

GetApplicationBuildIdOk returns a tuple with the ApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildId

`func (o *CrashEvent) SetApplicationBuildId(v string)`

SetApplicationBuildId sets ApplicationBuildId field to given value.


### GetApplicationBuildExecutable

`func (o *CrashEvent) GetApplicationBuildExecutable() string`

GetApplicationBuildExecutable returns the ApplicationBuildExecutable field if non-nil, zero value otherwise.

### GetApplicationBuildExecutableOk

`func (o *CrashEvent) GetApplicationBuildExecutableOk() (*string, bool)`

GetApplicationBuildExecutableOk returns a tuple with the ApplicationBuildExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildExecutable

`func (o *CrashEvent) SetApplicationBuildExecutable(v string)`

SetApplicationBuildExecutable sets ApplicationBuildExecutable field to given value.


### GetApplicationBuildStartParams

`func (o *CrashEvent) GetApplicationBuildStartParams() string`

GetApplicationBuildStartParams returns the ApplicationBuildStartParams field if non-nil, zero value otherwise.

### GetApplicationBuildStartParamsOk

`func (o *CrashEvent) GetApplicationBuildStartParamsOk() (*string, bool)`

GetApplicationBuildStartParamsOk returns a tuple with the ApplicationBuildStartParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildStartParams

`func (o *CrashEvent) SetApplicationBuildStartParams(v string)`

SetApplicationBuildStartParams sets ApplicationBuildStartParams field to given value.


### GetExitCode

`func (o *CrashEvent) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *CrashEvent) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *CrashEvent) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetDeploymentEnvironment

`func (o *CrashEvent) GetDeploymentEnvironment() DeploymentEnvironment`

GetDeploymentEnvironment returns the DeploymentEnvironment field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentOk

`func (o *CrashEvent) GetDeploymentEnvironmentOk() (*DeploymentEnvironment, bool)`

GetDeploymentEnvironmentOk returns a tuple with the DeploymentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironment

`func (o *CrashEvent) SetDeploymentEnvironment(v DeploymentEnvironment)`

SetDeploymentEnvironment sets DeploymentEnvironment field to given value.


### GetFleet

`func (o *CrashEvent) GetFleet() Fleet`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *CrashEvent) GetFleetOk() (*Fleet, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *CrashEvent) SetFleet(v Fleet)`

SetFleet sets Fleet field to given value.


### GetApplicationInstance

`func (o *CrashEvent) GetApplicationInstance() ApplicationInstance`

GetApplicationInstance returns the ApplicationInstance field if non-nil, zero value otherwise.

### GetApplicationInstanceOk

`func (o *CrashEvent) GetApplicationInstanceOk() (*ApplicationInstance, bool)`

GetApplicationInstanceOk returns a tuple with the ApplicationInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstance

`func (o *CrashEvent) SetApplicationInstance(v ApplicationInstance)`

SetApplicationInstance sets ApplicationInstance field to given value.


### GetApplicationInstanceConfiguration

`func (o *CrashEvent) GetApplicationInstanceConfiguration() []ApplicationInstanceConfiguration`

GetApplicationInstanceConfiguration returns the ApplicationInstanceConfiguration field if non-nil, zero value otherwise.

### GetApplicationInstanceConfigurationOk

`func (o *CrashEvent) GetApplicationInstanceConfigurationOk() (*[]ApplicationInstanceConfiguration, bool)`

GetApplicationInstanceConfigurationOk returns a tuple with the ApplicationInstanceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceConfiguration

`func (o *CrashEvent) SetApplicationInstanceConfiguration(v []ApplicationInstanceConfiguration)`

SetApplicationInstanceConfiguration sets ApplicationInstanceConfiguration field to given value.


### GetApplicationInstanceMetadata

`func (o *CrashEvent) GetApplicationInstanceMetadata() []Metadata`

GetApplicationInstanceMetadata returns the ApplicationInstanceMetadata field if non-nil, zero value otherwise.

### GetApplicationInstanceMetadataOk

`func (o *CrashEvent) GetApplicationInstanceMetadataOk() (*[]Metadata, bool)`

GetApplicationInstanceMetadataOk returns a tuple with the ApplicationInstanceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceMetadata

`func (o *CrashEvent) SetApplicationInstanceMetadata(v []Metadata)`

SetApplicationInstanceMetadata sets ApplicationInstanceMetadata field to given value.


### GetCreatedAt

`func (o *CrashEvent) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrashEvent) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrashEvent) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


