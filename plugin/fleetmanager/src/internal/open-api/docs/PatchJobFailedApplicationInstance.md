# PatchJobFailedApplicationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application instance ID | [readonly] 
**FleetId** | **string** | The application&#39;s fleet ID | [readonly] 
**FleetName** | **string** | The application&#39;s fleet name | [readonly] 
**HostId** | **int32** | The application instance is host on this server. | [readonly] 
**IsVirtual** | **int32** | 0 if this instance runs on a bare metal server, 1 if it runs on a VM | [readonly] 
**ApplicationId** | **string** | The ID of the application | [readonly] 
**ApplicationName** | **string** | The application&#39;s application name | [readonly] 
**ApplicationType** | **int32** | The type of the application, the different types can be found in:  [GET /v3/application/type](game-publisher#/Application/getApplicationTypes)] | [readonly] 
**ApplicationBuildId** | **string** | The ID of the application build | [readonly] 
**ApplicationBuildName** | **string** | The application&#39;s application build name | [readonly] 
**InstallId** | Pointer to **int32** | ID of the application install | [optional] 
**DcLocationId** | **int32** | The datacenter ID of where this application instance is located | [readonly] 
**DcLocationName** | **string** | The application&#39;s dc location name | [readonly] 
**RegionId** | **string** | The deployment region ID of where this application instance is located | [readonly] 
**RegionName** | **string** | The application&#39;s region name | [readonly] 
**Status** | **int32** | The application instance status, list of application instance status can be found in [&#x60;GET /v3/applicationInstance/status&#x60;](game-publisher#/ApplicationInstance/get_v3_applicationInstance_status) | [readonly] 
**CreatedAt** | **int32** | The Unix timestamp at which this application instance is created | [readonly] 
**StartedAt** | **int32** | The Unix timestamp at which this application instance is started | [readonly] 
**StoppedAt** | **int32** | The Unix timestamp at which this application instance is stopped | [readonly] 
**Pid** | **int32** | The process ID of the application instance running on the host | [readonly] 
**PidChangedAt** | **int32** | The Unix timestamp of the moment the process ID has been changed on the host | [readonly] 
**ManuallyDeployed** | **int32** | If set to 1, the application instance was manually deployed. | [readonly] 
**Properties** | [**[]ApplicationInstanceProperty**](ApplicationInstanceProperty.md) | Properties of application instance | [readonly] 
**IpAddress** | [**[]ApplicationInstanceIP**](ApplicationInstanceIP.md) | List of IP addresses for an application instance. | [readonly] 
**LabelReadOnly** | [**[]Label**](Label.md) | Pre-defined key/value pairs that can be used for application selection / identification. List of available labelReadOnly keys (application_id, fleet_id, host_id, dc_location_id, region_id, application_build_id) | [readonly] 
**LabelPublic** | Pointer to [**[]Label**](Label.md) | Custom key/value pairs that can be used for application instance selection / identification | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Custom key/value pairs to be set or adjusted by the allocation call or the application instance itself. This metadata will be passed to the application instance when starting it [needs more clarification]. | [optional] 
**NumPlayersMax** | **int32** | The maximum number of players on the application instance (this only apply to application instances of the type game) | [readonly] 
**NumPlayers** | **int32** | The number of online players on the application instance (this only apply to application instances of the type game) | [readonly] 
**LiveHostName** | **string** | The live host name of the application instance (this only apply to application instances of the type game) | [readonly] 
**LiveMap** | **string** | The live map of an application instance (this only apply to application instances of the type game) | [readonly] 
**LiveGameVersion** | **string** | The live game version of an application instance (this only apply to application instances of the type game) | [readonly] 
**LiveRules** | **string** | The live rules of an application instance (this only apply to application instances of the type game) | [readonly] 
**AutoRestart** | **int32** | If auto restart is 1 than the application instance will auto restart, when 0 the application instance auto restart will be disabled | [readonly] 
**MarkedForDeletion** | **int32** | if an application will be deleted it will be 1 else 0 | [readonly] 
**ExitCode** | **int32** | The exit code of the application, if it has crashed during patching else it will be 0 | [readonly] 
**FailingReason** | **string** | The reason why the application instance failed to patch | [readonly] 
**StartParameters** | **string** | The start parameters that were used to start the application | [readonly] 

## Methods

### NewPatchJobFailedApplicationInstance

`func NewPatchJobFailedApplicationInstance(id string, fleetId string, fleetName string, hostId int32, isVirtual int32, applicationId string, applicationName string, applicationType int32, applicationBuildId string, applicationBuildName string, dcLocationId int32, dcLocationName string, regionId string, regionName string, status int32, createdAt int32, startedAt int32, stoppedAt int32, pid int32, pidChangedAt int32, manuallyDeployed int32, properties []ApplicationInstanceProperty, ipAddress []ApplicationInstanceIP, labelReadOnly []Label, numPlayersMax int32, numPlayers int32, liveHostName string, liveMap string, liveGameVersion string, liveRules string, autoRestart int32, markedForDeletion int32, exitCode int32, failingReason string, startParameters string, ) *PatchJobFailedApplicationInstance`

NewPatchJobFailedApplicationInstance instantiates a new PatchJobFailedApplicationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobFailedApplicationInstanceWithDefaults

`func NewPatchJobFailedApplicationInstanceWithDefaults() *PatchJobFailedApplicationInstance`

NewPatchJobFailedApplicationInstanceWithDefaults instantiates a new PatchJobFailedApplicationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJobFailedApplicationInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJobFailedApplicationInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJobFailedApplicationInstance) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *PatchJobFailedApplicationInstance) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *PatchJobFailedApplicationInstance) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *PatchJobFailedApplicationInstance) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetName

`func (o *PatchJobFailedApplicationInstance) GetFleetName() string`

GetFleetName returns the FleetName field if non-nil, zero value otherwise.

### GetFleetNameOk

`func (o *PatchJobFailedApplicationInstance) GetFleetNameOk() (*string, bool)`

GetFleetNameOk returns a tuple with the FleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetName

`func (o *PatchJobFailedApplicationInstance) SetFleetName(v string)`

SetFleetName sets FleetName field to given value.


### GetHostId

`func (o *PatchJobFailedApplicationInstance) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *PatchJobFailedApplicationInstance) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *PatchJobFailedApplicationInstance) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetIsVirtual

`func (o *PatchJobFailedApplicationInstance) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *PatchJobFailedApplicationInstance) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *PatchJobFailedApplicationInstance) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetApplicationId

`func (o *PatchJobFailedApplicationInstance) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PatchJobFailedApplicationInstance) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PatchJobFailedApplicationInstance) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *PatchJobFailedApplicationInstance) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *PatchJobFailedApplicationInstance) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *PatchJobFailedApplicationInstance) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetApplicationType

`func (o *PatchJobFailedApplicationInstance) GetApplicationType() int32`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *PatchJobFailedApplicationInstance) GetApplicationTypeOk() (*int32, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *PatchJobFailedApplicationInstance) SetApplicationType(v int32)`

SetApplicationType sets ApplicationType field to given value.


### GetApplicationBuildId

`func (o *PatchJobFailedApplicationInstance) GetApplicationBuildId() string`

GetApplicationBuildId returns the ApplicationBuildId field if non-nil, zero value otherwise.

### GetApplicationBuildIdOk

`func (o *PatchJobFailedApplicationInstance) GetApplicationBuildIdOk() (*string, bool)`

GetApplicationBuildIdOk returns a tuple with the ApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildId

`func (o *PatchJobFailedApplicationInstance) SetApplicationBuildId(v string)`

SetApplicationBuildId sets ApplicationBuildId field to given value.


### GetApplicationBuildName

`func (o *PatchJobFailedApplicationInstance) GetApplicationBuildName() string`

GetApplicationBuildName returns the ApplicationBuildName field if non-nil, zero value otherwise.

### GetApplicationBuildNameOk

`func (o *PatchJobFailedApplicationInstance) GetApplicationBuildNameOk() (*string, bool)`

GetApplicationBuildNameOk returns a tuple with the ApplicationBuildName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildName

`func (o *PatchJobFailedApplicationInstance) SetApplicationBuildName(v string)`

SetApplicationBuildName sets ApplicationBuildName field to given value.


### GetInstallId

`func (o *PatchJobFailedApplicationInstance) GetInstallId() int32`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *PatchJobFailedApplicationInstance) GetInstallIdOk() (*int32, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *PatchJobFailedApplicationInstance) SetInstallId(v int32)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *PatchJobFailedApplicationInstance) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### GetDcLocationId

`func (o *PatchJobFailedApplicationInstance) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *PatchJobFailedApplicationInstance) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *PatchJobFailedApplicationInstance) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetDcLocationName

`func (o *PatchJobFailedApplicationInstance) GetDcLocationName() string`

GetDcLocationName returns the DcLocationName field if non-nil, zero value otherwise.

### GetDcLocationNameOk

`func (o *PatchJobFailedApplicationInstance) GetDcLocationNameOk() (*string, bool)`

GetDcLocationNameOk returns a tuple with the DcLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationName

`func (o *PatchJobFailedApplicationInstance) SetDcLocationName(v string)`

SetDcLocationName sets DcLocationName field to given value.


### GetRegionId

`func (o *PatchJobFailedApplicationInstance) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *PatchJobFailedApplicationInstance) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *PatchJobFailedApplicationInstance) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetRegionName

`func (o *PatchJobFailedApplicationInstance) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *PatchJobFailedApplicationInstance) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *PatchJobFailedApplicationInstance) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetStatus

`func (o *PatchJobFailedApplicationInstance) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchJobFailedApplicationInstance) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchJobFailedApplicationInstance) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *PatchJobFailedApplicationInstance) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchJobFailedApplicationInstance) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchJobFailedApplicationInstance) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *PatchJobFailedApplicationInstance) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PatchJobFailedApplicationInstance) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PatchJobFailedApplicationInstance) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.


### GetStoppedAt

`func (o *PatchJobFailedApplicationInstance) GetStoppedAt() int32`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *PatchJobFailedApplicationInstance) GetStoppedAtOk() (*int32, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *PatchJobFailedApplicationInstance) SetStoppedAt(v int32)`

SetStoppedAt sets StoppedAt field to given value.


### GetPid

`func (o *PatchJobFailedApplicationInstance) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *PatchJobFailedApplicationInstance) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *PatchJobFailedApplicationInstance) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetPidChangedAt

`func (o *PatchJobFailedApplicationInstance) GetPidChangedAt() int32`

GetPidChangedAt returns the PidChangedAt field if non-nil, zero value otherwise.

### GetPidChangedAtOk

`func (o *PatchJobFailedApplicationInstance) GetPidChangedAtOk() (*int32, bool)`

GetPidChangedAtOk returns a tuple with the PidChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidChangedAt

`func (o *PatchJobFailedApplicationInstance) SetPidChangedAt(v int32)`

SetPidChangedAt sets PidChangedAt field to given value.


### GetManuallyDeployed

`func (o *PatchJobFailedApplicationInstance) GetManuallyDeployed() int32`

GetManuallyDeployed returns the ManuallyDeployed field if non-nil, zero value otherwise.

### GetManuallyDeployedOk

`func (o *PatchJobFailedApplicationInstance) GetManuallyDeployedOk() (*int32, bool)`

GetManuallyDeployedOk returns a tuple with the ManuallyDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyDeployed

`func (o *PatchJobFailedApplicationInstance) SetManuallyDeployed(v int32)`

SetManuallyDeployed sets ManuallyDeployed field to given value.


### GetProperties

`func (o *PatchJobFailedApplicationInstance) GetProperties() []ApplicationInstanceProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchJobFailedApplicationInstance) GetPropertiesOk() (*[]ApplicationInstanceProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchJobFailedApplicationInstance) SetProperties(v []ApplicationInstanceProperty)`

SetProperties sets Properties field to given value.


### GetIpAddress

`func (o *PatchJobFailedApplicationInstance) GetIpAddress() []ApplicationInstanceIP`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PatchJobFailedApplicationInstance) GetIpAddressOk() (*[]ApplicationInstanceIP, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PatchJobFailedApplicationInstance) SetIpAddress(v []ApplicationInstanceIP)`

SetIpAddress sets IpAddress field to given value.


### GetLabelReadOnly

`func (o *PatchJobFailedApplicationInstance) GetLabelReadOnly() []Label`

GetLabelReadOnly returns the LabelReadOnly field if non-nil, zero value otherwise.

### GetLabelReadOnlyOk

`func (o *PatchJobFailedApplicationInstance) GetLabelReadOnlyOk() (*[]Label, bool)`

GetLabelReadOnlyOk returns a tuple with the LabelReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelReadOnly

`func (o *PatchJobFailedApplicationInstance) SetLabelReadOnly(v []Label)`

SetLabelReadOnly sets LabelReadOnly field to given value.


### GetLabelPublic

`func (o *PatchJobFailedApplicationInstance) GetLabelPublic() []Label`

GetLabelPublic returns the LabelPublic field if non-nil, zero value otherwise.

### GetLabelPublicOk

`func (o *PatchJobFailedApplicationInstance) GetLabelPublicOk() (*[]Label, bool)`

GetLabelPublicOk returns a tuple with the LabelPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPublic

`func (o *PatchJobFailedApplicationInstance) SetLabelPublic(v []Label)`

SetLabelPublic sets LabelPublic field to given value.

### HasLabelPublic

`func (o *PatchJobFailedApplicationInstance) HasLabelPublic() bool`

HasLabelPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchJobFailedApplicationInstance) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchJobFailedApplicationInstance) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchJobFailedApplicationInstance) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchJobFailedApplicationInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNumPlayersMax

`func (o *PatchJobFailedApplicationInstance) GetNumPlayersMax() int32`

GetNumPlayersMax returns the NumPlayersMax field if non-nil, zero value otherwise.

### GetNumPlayersMaxOk

`func (o *PatchJobFailedApplicationInstance) GetNumPlayersMaxOk() (*int32, bool)`

GetNumPlayersMaxOk returns a tuple with the NumPlayersMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPlayersMax

`func (o *PatchJobFailedApplicationInstance) SetNumPlayersMax(v int32)`

SetNumPlayersMax sets NumPlayersMax field to given value.


### GetNumPlayers

`func (o *PatchJobFailedApplicationInstance) GetNumPlayers() int32`

GetNumPlayers returns the NumPlayers field if non-nil, zero value otherwise.

### GetNumPlayersOk

`func (o *PatchJobFailedApplicationInstance) GetNumPlayersOk() (*int32, bool)`

GetNumPlayersOk returns a tuple with the NumPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPlayers

`func (o *PatchJobFailedApplicationInstance) SetNumPlayers(v int32)`

SetNumPlayers sets NumPlayers field to given value.


### GetLiveHostName

`func (o *PatchJobFailedApplicationInstance) GetLiveHostName() string`

GetLiveHostName returns the LiveHostName field if non-nil, zero value otherwise.

### GetLiveHostNameOk

`func (o *PatchJobFailedApplicationInstance) GetLiveHostNameOk() (*string, bool)`

GetLiveHostNameOk returns a tuple with the LiveHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHostName

`func (o *PatchJobFailedApplicationInstance) SetLiveHostName(v string)`

SetLiveHostName sets LiveHostName field to given value.


### GetLiveMap

`func (o *PatchJobFailedApplicationInstance) GetLiveMap() string`

GetLiveMap returns the LiveMap field if non-nil, zero value otherwise.

### GetLiveMapOk

`func (o *PatchJobFailedApplicationInstance) GetLiveMapOk() (*string, bool)`

GetLiveMapOk returns a tuple with the LiveMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveMap

`func (o *PatchJobFailedApplicationInstance) SetLiveMap(v string)`

SetLiveMap sets LiveMap field to given value.


### GetLiveGameVersion

`func (o *PatchJobFailedApplicationInstance) GetLiveGameVersion() string`

GetLiveGameVersion returns the LiveGameVersion field if non-nil, zero value otherwise.

### GetLiveGameVersionOk

`func (o *PatchJobFailedApplicationInstance) GetLiveGameVersionOk() (*string, bool)`

GetLiveGameVersionOk returns a tuple with the LiveGameVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveGameVersion

`func (o *PatchJobFailedApplicationInstance) SetLiveGameVersion(v string)`

SetLiveGameVersion sets LiveGameVersion field to given value.


### GetLiveRules

`func (o *PatchJobFailedApplicationInstance) GetLiveRules() string`

GetLiveRules returns the LiveRules field if non-nil, zero value otherwise.

### GetLiveRulesOk

`func (o *PatchJobFailedApplicationInstance) GetLiveRulesOk() (*string, bool)`

GetLiveRulesOk returns a tuple with the LiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveRules

`func (o *PatchJobFailedApplicationInstance) SetLiveRules(v string)`

SetLiveRules sets LiveRules field to given value.


### GetAutoRestart

`func (o *PatchJobFailedApplicationInstance) GetAutoRestart() int32`

GetAutoRestart returns the AutoRestart field if non-nil, zero value otherwise.

### GetAutoRestartOk

`func (o *PatchJobFailedApplicationInstance) GetAutoRestartOk() (*int32, bool)`

GetAutoRestartOk returns a tuple with the AutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRestart

`func (o *PatchJobFailedApplicationInstance) SetAutoRestart(v int32)`

SetAutoRestart sets AutoRestart field to given value.


### GetMarkedForDeletion

`func (o *PatchJobFailedApplicationInstance) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *PatchJobFailedApplicationInstance) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *PatchJobFailedApplicationInstance) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetExitCode

`func (o *PatchJobFailedApplicationInstance) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *PatchJobFailedApplicationInstance) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *PatchJobFailedApplicationInstance) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetFailingReason

`func (o *PatchJobFailedApplicationInstance) GetFailingReason() string`

GetFailingReason returns the FailingReason field if non-nil, zero value otherwise.

### GetFailingReasonOk

`func (o *PatchJobFailedApplicationInstance) GetFailingReasonOk() (*string, bool)`

GetFailingReasonOk returns a tuple with the FailingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingReason

`func (o *PatchJobFailedApplicationInstance) SetFailingReason(v string)`

SetFailingReason sets FailingReason field to given value.


### GetStartParameters

`func (o *PatchJobFailedApplicationInstance) GetStartParameters() string`

GetStartParameters returns the StartParameters field if non-nil, zero value otherwise.

### GetStartParametersOk

`func (o *PatchJobFailedApplicationInstance) GetStartParametersOk() (*string, bool)`

GetStartParametersOk returns a tuple with the StartParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartParameters

`func (o *PatchJobFailedApplicationInstance) SetStartParameters(v string)`

SetStartParameters sets StartParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


