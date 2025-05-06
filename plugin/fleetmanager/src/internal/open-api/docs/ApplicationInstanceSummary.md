# ApplicationInstanceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application instance ID | [readonly] 
**DeploymentEnvironmentId** | **string** | The deployment environment ID of this application instance | [readonly] 
**DeploymentEnvironmentName** | **string** | The deployment environment name of this application instance | [readonly] 
**FleetId** | **string** | The fleet ID of this application instance | [readonly] 
**FleetName** | **string** | The fleet name of this application instance | [readonly] 
**HostId** | **int32** | The application instance is host on this server. | [readonly] 
**IsVirtual** | **int32** | 0 if this instance runs on a bare metal server, 1 if it runs on a VM | [readonly] 
**ApplicationId** | **string** | The application ID of this application instance | [readonly] 
**ApplicationName** | **string** | The application name of this application instance | [readonly] 
**ApplicationType** | **int32** | The type of the application, the different types can be found in:  GET /application/type | [readonly] 
**ApplicationBuildId** | **string** | The application build ID of this application instance | [readonly] 
**ApplicationBuildName** | **string** | The application build name of this application instance | [readonly] 
**InstallId** | Pointer to **NullableString** | ID of the application install | [optional] 
**DcLocationId** | **int32** | The datacenter ID of where this application instance is located | [readonly] 
**DcLocationName** | **string** | The datacenter name of where this application instance is located | [readonly] 
**RegionId** | **string** | The deployment region ID of where this application instance is located | [readonly] 
**RegionName** | **string** | The deployment region name of where this application instance is located | [readonly] 
**Status** | **int32** | The application instance status, list of application instance status can be found in [&#x60;GET /v3/applicationInstance/status&#x60;](game-publisher#/ApplicationInstance/get_v3_applicationInstance_status) | [readonly] 
**CreatedAt** | **int32** | The Unix timestamp at which this application instance is created | [readonly] 
**StartedAt** | **int32** | The Unix timestamp at which this application instance is started | [readonly] 
**StoppedAt** | **int32** | The Unix timestamp at which this application instance is stopped | [readonly] 
**Pid** | **int32** | The process ID of the application instance running on the host | [readonly] 
**PidChangedAt** | **int32** | The Unix timestamp of the moment the process ID has been changed on the host | [readonly] 
**ManuallyDeployed** | **int32** | If set to 1, the application instance was manually deployed. | [readonly] 
**AutoRestart** | **int32** | If auto restart is 1 than the application instance will auto restart, when 0 the application instance auto restart will be disabled | [readonly] 
**MarkedForDeletion** | **int32** | if an application will be deleted it will be 1 else 0 | [readonly] 
**ArcusAvailable** | **int32** | &#x60;1&#x60; if Arcus is available &amp; running for this application instance, &#x60;0&#x60; otherwise. Note: If your application&#39;s &#x60;managementProtocol&#x60; is NOT set to Arcus, this value will always be &#x60;0&#x60;, even if your game server would support the Arcus protocol. The protocol is only tested, and this value potentially be set to &#x60;1&#x60;, when the &#x60;managementProtocol&#x60; is set to Arcus. | [readonly] 

## Methods

### NewApplicationInstanceSummary

`func NewApplicationInstanceSummary(id string, deploymentEnvironmentId string, deploymentEnvironmentName string, fleetId string, fleetName string, hostId int32, isVirtual int32, applicationId string, applicationName string, applicationType int32, applicationBuildId string, applicationBuildName string, dcLocationId int32, dcLocationName string, regionId string, regionName string, status int32, createdAt int32, startedAt int32, stoppedAt int32, pid int32, pidChangedAt int32, manuallyDeployed int32, autoRestart int32, markedForDeletion int32, arcusAvailable int32, ) *ApplicationInstanceSummary`

NewApplicationInstanceSummary instantiates a new ApplicationInstanceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceSummaryWithDefaults

`func NewApplicationInstanceSummaryWithDefaults() *ApplicationInstanceSummary`

NewApplicationInstanceSummaryWithDefaults instantiates a new ApplicationInstanceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstanceSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstanceSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstanceSummary) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentEnvironmentId

`func (o *ApplicationInstanceSummary) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *ApplicationInstanceSummary) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *ApplicationInstanceSummary) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetDeploymentEnvironmentName

`func (o *ApplicationInstanceSummary) GetDeploymentEnvironmentName() string`

GetDeploymentEnvironmentName returns the DeploymentEnvironmentName field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentNameOk

`func (o *ApplicationInstanceSummary) GetDeploymentEnvironmentNameOk() (*string, bool)`

GetDeploymentEnvironmentNameOk returns a tuple with the DeploymentEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentName

`func (o *ApplicationInstanceSummary) SetDeploymentEnvironmentName(v string)`

SetDeploymentEnvironmentName sets DeploymentEnvironmentName field to given value.


### GetFleetId

`func (o *ApplicationInstanceSummary) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ApplicationInstanceSummary) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ApplicationInstanceSummary) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetName

`func (o *ApplicationInstanceSummary) GetFleetName() string`

GetFleetName returns the FleetName field if non-nil, zero value otherwise.

### GetFleetNameOk

`func (o *ApplicationInstanceSummary) GetFleetNameOk() (*string, bool)`

GetFleetNameOk returns a tuple with the FleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetName

`func (o *ApplicationInstanceSummary) SetFleetName(v string)`

SetFleetName sets FleetName field to given value.


### GetHostId

`func (o *ApplicationInstanceSummary) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ApplicationInstanceSummary) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ApplicationInstanceSummary) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetIsVirtual

`func (o *ApplicationInstanceSummary) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ApplicationInstanceSummary) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ApplicationInstanceSummary) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetApplicationId

`func (o *ApplicationInstanceSummary) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationInstanceSummary) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationInstanceSummary) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *ApplicationInstanceSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApplicationInstanceSummary) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApplicationInstanceSummary) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetApplicationType

`func (o *ApplicationInstanceSummary) GetApplicationType() int32`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationInstanceSummary) GetApplicationTypeOk() (*int32, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationInstanceSummary) SetApplicationType(v int32)`

SetApplicationType sets ApplicationType field to given value.


### GetApplicationBuildId

`func (o *ApplicationInstanceSummary) GetApplicationBuildId() string`

GetApplicationBuildId returns the ApplicationBuildId field if non-nil, zero value otherwise.

### GetApplicationBuildIdOk

`func (o *ApplicationInstanceSummary) GetApplicationBuildIdOk() (*string, bool)`

GetApplicationBuildIdOk returns a tuple with the ApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildId

`func (o *ApplicationInstanceSummary) SetApplicationBuildId(v string)`

SetApplicationBuildId sets ApplicationBuildId field to given value.


### GetApplicationBuildName

`func (o *ApplicationInstanceSummary) GetApplicationBuildName() string`

GetApplicationBuildName returns the ApplicationBuildName field if non-nil, zero value otherwise.

### GetApplicationBuildNameOk

`func (o *ApplicationInstanceSummary) GetApplicationBuildNameOk() (*string, bool)`

GetApplicationBuildNameOk returns a tuple with the ApplicationBuildName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildName

`func (o *ApplicationInstanceSummary) SetApplicationBuildName(v string)`

SetApplicationBuildName sets ApplicationBuildName field to given value.


### GetInstallId

`func (o *ApplicationInstanceSummary) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *ApplicationInstanceSummary) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *ApplicationInstanceSummary) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *ApplicationInstanceSummary) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### SetInstallIdNil

`func (o *ApplicationInstanceSummary) SetInstallIdNil(b bool)`

 SetInstallIdNil sets the value for InstallId to be an explicit nil

### UnsetInstallId
`func (o *ApplicationInstanceSummary) UnsetInstallId()`

UnsetInstallId ensures that no value is present for InstallId, not even an explicit nil
### GetDcLocationId

`func (o *ApplicationInstanceSummary) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *ApplicationInstanceSummary) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *ApplicationInstanceSummary) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetDcLocationName

`func (o *ApplicationInstanceSummary) GetDcLocationName() string`

GetDcLocationName returns the DcLocationName field if non-nil, zero value otherwise.

### GetDcLocationNameOk

`func (o *ApplicationInstanceSummary) GetDcLocationNameOk() (*string, bool)`

GetDcLocationNameOk returns a tuple with the DcLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationName

`func (o *ApplicationInstanceSummary) SetDcLocationName(v string)`

SetDcLocationName sets DcLocationName field to given value.


### GetRegionId

`func (o *ApplicationInstanceSummary) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ApplicationInstanceSummary) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ApplicationInstanceSummary) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetRegionName

`func (o *ApplicationInstanceSummary) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApplicationInstanceSummary) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApplicationInstanceSummary) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetStatus

`func (o *ApplicationInstanceSummary) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationInstanceSummary) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationInstanceSummary) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ApplicationInstanceSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationInstanceSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationInstanceSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *ApplicationInstanceSummary) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApplicationInstanceSummary) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApplicationInstanceSummary) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.


### GetStoppedAt

`func (o *ApplicationInstanceSummary) GetStoppedAt() int32`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *ApplicationInstanceSummary) GetStoppedAtOk() (*int32, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *ApplicationInstanceSummary) SetStoppedAt(v int32)`

SetStoppedAt sets StoppedAt field to given value.


### GetPid

`func (o *ApplicationInstanceSummary) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ApplicationInstanceSummary) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ApplicationInstanceSummary) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetPidChangedAt

`func (o *ApplicationInstanceSummary) GetPidChangedAt() int32`

GetPidChangedAt returns the PidChangedAt field if non-nil, zero value otherwise.

### GetPidChangedAtOk

`func (o *ApplicationInstanceSummary) GetPidChangedAtOk() (*int32, bool)`

GetPidChangedAtOk returns a tuple with the PidChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidChangedAt

`func (o *ApplicationInstanceSummary) SetPidChangedAt(v int32)`

SetPidChangedAt sets PidChangedAt field to given value.


### GetManuallyDeployed

`func (o *ApplicationInstanceSummary) GetManuallyDeployed() int32`

GetManuallyDeployed returns the ManuallyDeployed field if non-nil, zero value otherwise.

### GetManuallyDeployedOk

`func (o *ApplicationInstanceSummary) GetManuallyDeployedOk() (*int32, bool)`

GetManuallyDeployedOk returns a tuple with the ManuallyDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyDeployed

`func (o *ApplicationInstanceSummary) SetManuallyDeployed(v int32)`

SetManuallyDeployed sets ManuallyDeployed field to given value.


### GetAutoRestart

`func (o *ApplicationInstanceSummary) GetAutoRestart() int32`

GetAutoRestart returns the AutoRestart field if non-nil, zero value otherwise.

### GetAutoRestartOk

`func (o *ApplicationInstanceSummary) GetAutoRestartOk() (*int32, bool)`

GetAutoRestartOk returns a tuple with the AutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRestart

`func (o *ApplicationInstanceSummary) SetAutoRestart(v int32)`

SetAutoRestart sets AutoRestart field to given value.


### GetMarkedForDeletion

`func (o *ApplicationInstanceSummary) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *ApplicationInstanceSummary) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *ApplicationInstanceSummary) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetArcusAvailable

`func (o *ApplicationInstanceSummary) GetArcusAvailable() int32`

GetArcusAvailable returns the ArcusAvailable field if non-nil, zero value otherwise.

### GetArcusAvailableOk

`func (o *ApplicationInstanceSummary) GetArcusAvailableOk() (*int32, bool)`

GetArcusAvailableOk returns a tuple with the ArcusAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcusAvailable

`func (o *ApplicationInstanceSummary) SetArcusAvailable(v int32)`

SetArcusAvailable sets ArcusAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


