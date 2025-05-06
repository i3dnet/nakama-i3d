# ApplicationInstance

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
**ApplicationType** | **int32** | The type of the application, the different types can be found in: [GET /v3/application/type](game-publisher#/Application/getApplicationTypes) | [readonly] 
**ApplicationBuildId** | **string** | The application build ID of this application instance | [readonly] 
**ApplicationBuildName** | **string** | The application build name of this application instance | [readonly] 
**InstallId** | Pointer to **NullableString** | ID of the application install | [optional] 
**DcLocationId** | **int32** | The datacenter ID of where this application instance is located | [readonly] 
**DcLocationName** | **string** | The datacenter name of where this application instance is located | [readonly] 
**RegionId** | **string** | The deployment region ID of where this application instance is located | [readonly] 
**RegionName** | **string** | The deployment region name of where this application instance is located | [readonly] 
**Status** | **int32** | The application instance status, list of application instance status can be found in [GET /v3/applicationInstance/status](game-publisher#/ApplicationInstance/getApplicationInstanceStatuses) | [readonly] 
**CreatedAt** | **int32** | The Unix timestamp at which this application instance is created | [readonly] 
**StartedAt** | **int32** | The Unix timestamp at which this application instance is started | [readonly] 
**StoppedAt** | **int32** | The Unix timestamp at which this application instance is stopped | [readonly] 
**Pid** | **int32** | The process ID of the application instance running on the host | [readonly] 
**PidChangedAt** | **int32** | The Unix timestamp of the moment the process ID has been changed on the host | [readonly] 
**StartupParams** | **NullableString** | The exact startup parameters that were used to start this application instance with all properties resolved to their respective values. Will be &#x60;null&#x60; if the instance has not been started, but will retain its previous value if it was stopped. Once set it will never be cleared again | [readonly] 
**Executable** | **NullableString** | The name of the executable that was ran when the application instance started. Will be &#x60;null&#x60; if the instance has not been started, but will retain its previous value if it was stopped. Once set it will never be cleared again | [readonly] 
**ManuallyDeployed** | **int32** | If set to &#x60;1&#x60;, the application instance was manually deployed, instead of automatically deployed through the system | [readonly] 
**Properties** | [**[]ApplicationInstanceProperty**](ApplicationInstanceProperty.md) | Properties of application instance | [readonly] 
**IpAddress** | [**[]ApplicationInstanceIP**](ApplicationInstanceIP.md) | List of IP addresses for an application instance | [readonly] 
**LabelReadOnly** | [**[]Label**](Label.md) | Pre-defined key/value pairs that can be used for application selection / identification. Available labelReadOnly keys are: * application_id * fleet_id * host_id * dc_location_id * region_id * application_build_id | [readonly] 
**LabelPublic** | Pointer to [**[]Label**](Label.md) | Custom key/value pairs that can be used for application instance selection / identification | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Custom key/value pairs to be set or adjusted by the allocation call or the application instance itself. This metadata will be passed to the application instance when starting it [needs more clarification] | [optional] 
**NumPlayersMax** | **int32** | The maximum number of players on the application instance (this only apply to application instances of the type game) | [readonly] 
**NumPlayers** | **int32** | The number of online players on the application instance (this only apply to application instances of the type game) | [readonly] 
**LiveHostName** | **string** | The live host name of the application instance (this only apply to application instances of the type game) | [readonly] 
**LiveMap** | **string** | The live map of an application instance (this only apply to application instances of the type game) | [readonly] 
**LiveGameVersion** | **string** | The live game version of an application instance (this only apply to application instances of the type game) | [readonly] 
**UpdatedAt** | **int32** | The Unix timestamp at which this application instance is live data is updated (this only apply to application instances of the type game) | [readonly] 
**LiveRules** | **string** | The live rules of an application instance (this only apply to application instances of the type game) | [readonly] 
**AutoRestart** | **int32** | If auto restart is &#x60;1&#x60; then the application instance will auto restart, when &#x60;0&#x60; the application instance auto restart will be disabled | [readonly] 
**MarkedForDeletion** | **int32** | If an application will be deleted it will be &#x60;1&#x60;, &#x60;0&#x60; otherwise | [readonly] 
**ArcusAvailable** | **int32** | &#x60;1&#x60; if Arcus is available &amp; running for this application instance, &#x60;0&#x60; otherwise. Note: If your application&#39;s &#x60;managementProtocol&#x60; is NOT set to Arcus, this value will always be &#x60;0&#x60;, even if your game server would support the Arcus protocol. The protocol is only tested, and this value potentially be set to &#x60;1&#x60;, when the &#x60;managementProtocol&#x60; is set to Arcus. | [readonly] 

## Methods

### NewApplicationInstance

`func NewApplicationInstance(id string, deploymentEnvironmentId string, deploymentEnvironmentName string, fleetId string, fleetName string, hostId int32, isVirtual int32, applicationId string, applicationName string, applicationType int32, applicationBuildId string, applicationBuildName string, dcLocationId int32, dcLocationName string, regionId string, regionName string, status int32, createdAt int32, startedAt int32, stoppedAt int32, pid int32, pidChangedAt int32, startupParams NullableString, executable NullableString, manuallyDeployed int32, properties []ApplicationInstanceProperty, ipAddress []ApplicationInstanceIP, labelReadOnly []Label, numPlayersMax int32, numPlayers int32, liveHostName string, liveMap string, liveGameVersion string, updatedAt int32, liveRules string, autoRestart int32, markedForDeletion int32, arcusAvailable int32, ) *ApplicationInstance`

NewApplicationInstance instantiates a new ApplicationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceWithDefaults

`func NewApplicationInstanceWithDefaults() *ApplicationInstance`

NewApplicationInstanceWithDefaults instantiates a new ApplicationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstance) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentEnvironmentId

`func (o *ApplicationInstance) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *ApplicationInstance) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *ApplicationInstance) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetDeploymentEnvironmentName

`func (o *ApplicationInstance) GetDeploymentEnvironmentName() string`

GetDeploymentEnvironmentName returns the DeploymentEnvironmentName field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentNameOk

`func (o *ApplicationInstance) GetDeploymentEnvironmentNameOk() (*string, bool)`

GetDeploymentEnvironmentNameOk returns a tuple with the DeploymentEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentName

`func (o *ApplicationInstance) SetDeploymentEnvironmentName(v string)`

SetDeploymentEnvironmentName sets DeploymentEnvironmentName field to given value.


### GetFleetId

`func (o *ApplicationInstance) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ApplicationInstance) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ApplicationInstance) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetName

`func (o *ApplicationInstance) GetFleetName() string`

GetFleetName returns the FleetName field if non-nil, zero value otherwise.

### GetFleetNameOk

`func (o *ApplicationInstance) GetFleetNameOk() (*string, bool)`

GetFleetNameOk returns a tuple with the FleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetName

`func (o *ApplicationInstance) SetFleetName(v string)`

SetFleetName sets FleetName field to given value.


### GetHostId

`func (o *ApplicationInstance) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ApplicationInstance) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ApplicationInstance) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetIsVirtual

`func (o *ApplicationInstance) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ApplicationInstance) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ApplicationInstance) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetApplicationId

`func (o *ApplicationInstance) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationInstance) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationInstance) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *ApplicationInstance) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApplicationInstance) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApplicationInstance) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetApplicationType

`func (o *ApplicationInstance) GetApplicationType() int32`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationInstance) GetApplicationTypeOk() (*int32, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationInstance) SetApplicationType(v int32)`

SetApplicationType sets ApplicationType field to given value.


### GetApplicationBuildId

`func (o *ApplicationInstance) GetApplicationBuildId() string`

GetApplicationBuildId returns the ApplicationBuildId field if non-nil, zero value otherwise.

### GetApplicationBuildIdOk

`func (o *ApplicationInstance) GetApplicationBuildIdOk() (*string, bool)`

GetApplicationBuildIdOk returns a tuple with the ApplicationBuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildId

`func (o *ApplicationInstance) SetApplicationBuildId(v string)`

SetApplicationBuildId sets ApplicationBuildId field to given value.


### GetApplicationBuildName

`func (o *ApplicationInstance) GetApplicationBuildName() string`

GetApplicationBuildName returns the ApplicationBuildName field if non-nil, zero value otherwise.

### GetApplicationBuildNameOk

`func (o *ApplicationInstance) GetApplicationBuildNameOk() (*string, bool)`

GetApplicationBuildNameOk returns a tuple with the ApplicationBuildName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildName

`func (o *ApplicationInstance) SetApplicationBuildName(v string)`

SetApplicationBuildName sets ApplicationBuildName field to given value.


### GetInstallId

`func (o *ApplicationInstance) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *ApplicationInstance) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *ApplicationInstance) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.

### HasInstallId

`func (o *ApplicationInstance) HasInstallId() bool`

HasInstallId returns a boolean if a field has been set.

### SetInstallIdNil

`func (o *ApplicationInstance) SetInstallIdNil(b bool)`

 SetInstallIdNil sets the value for InstallId to be an explicit nil

### UnsetInstallId
`func (o *ApplicationInstance) UnsetInstallId()`

UnsetInstallId ensures that no value is present for InstallId, not even an explicit nil
### GetDcLocationId

`func (o *ApplicationInstance) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *ApplicationInstance) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *ApplicationInstance) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetDcLocationName

`func (o *ApplicationInstance) GetDcLocationName() string`

GetDcLocationName returns the DcLocationName field if non-nil, zero value otherwise.

### GetDcLocationNameOk

`func (o *ApplicationInstance) GetDcLocationNameOk() (*string, bool)`

GetDcLocationNameOk returns a tuple with the DcLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationName

`func (o *ApplicationInstance) SetDcLocationName(v string)`

SetDcLocationName sets DcLocationName field to given value.


### GetRegionId

`func (o *ApplicationInstance) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ApplicationInstance) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ApplicationInstance) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetRegionName

`func (o *ApplicationInstance) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApplicationInstance) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApplicationInstance) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetStatus

`func (o *ApplicationInstance) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationInstance) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationInstance) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ApplicationInstance) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationInstance) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationInstance) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *ApplicationInstance) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApplicationInstance) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApplicationInstance) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.


### GetStoppedAt

`func (o *ApplicationInstance) GetStoppedAt() int32`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *ApplicationInstance) GetStoppedAtOk() (*int32, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *ApplicationInstance) SetStoppedAt(v int32)`

SetStoppedAt sets StoppedAt field to given value.


### GetPid

`func (o *ApplicationInstance) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ApplicationInstance) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ApplicationInstance) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetPidChangedAt

`func (o *ApplicationInstance) GetPidChangedAt() int32`

GetPidChangedAt returns the PidChangedAt field if non-nil, zero value otherwise.

### GetPidChangedAtOk

`func (o *ApplicationInstance) GetPidChangedAtOk() (*int32, bool)`

GetPidChangedAtOk returns a tuple with the PidChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidChangedAt

`func (o *ApplicationInstance) SetPidChangedAt(v int32)`

SetPidChangedAt sets PidChangedAt field to given value.


### GetStartupParams

`func (o *ApplicationInstance) GetStartupParams() string`

GetStartupParams returns the StartupParams field if non-nil, zero value otherwise.

### GetStartupParamsOk

`func (o *ApplicationInstance) GetStartupParamsOk() (*string, bool)`

GetStartupParamsOk returns a tuple with the StartupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupParams

`func (o *ApplicationInstance) SetStartupParams(v string)`

SetStartupParams sets StartupParams field to given value.


### SetStartupParamsNil

`func (o *ApplicationInstance) SetStartupParamsNil(b bool)`

 SetStartupParamsNil sets the value for StartupParams to be an explicit nil

### UnsetStartupParams
`func (o *ApplicationInstance) UnsetStartupParams()`

UnsetStartupParams ensures that no value is present for StartupParams, not even an explicit nil
### GetExecutable

`func (o *ApplicationInstance) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *ApplicationInstance) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *ApplicationInstance) SetExecutable(v string)`

SetExecutable sets Executable field to given value.


### SetExecutableNil

`func (o *ApplicationInstance) SetExecutableNil(b bool)`

 SetExecutableNil sets the value for Executable to be an explicit nil

### UnsetExecutable
`func (o *ApplicationInstance) UnsetExecutable()`

UnsetExecutable ensures that no value is present for Executable, not even an explicit nil
### GetManuallyDeployed

`func (o *ApplicationInstance) GetManuallyDeployed() int32`

GetManuallyDeployed returns the ManuallyDeployed field if non-nil, zero value otherwise.

### GetManuallyDeployedOk

`func (o *ApplicationInstance) GetManuallyDeployedOk() (*int32, bool)`

GetManuallyDeployedOk returns a tuple with the ManuallyDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyDeployed

`func (o *ApplicationInstance) SetManuallyDeployed(v int32)`

SetManuallyDeployed sets ManuallyDeployed field to given value.


### GetProperties

`func (o *ApplicationInstance) GetProperties() []ApplicationInstanceProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ApplicationInstance) GetPropertiesOk() (*[]ApplicationInstanceProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ApplicationInstance) SetProperties(v []ApplicationInstanceProperty)`

SetProperties sets Properties field to given value.


### GetIpAddress

`func (o *ApplicationInstance) GetIpAddress() []ApplicationInstanceIP`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApplicationInstance) GetIpAddressOk() (*[]ApplicationInstanceIP, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApplicationInstance) SetIpAddress(v []ApplicationInstanceIP)`

SetIpAddress sets IpAddress field to given value.


### GetLabelReadOnly

`func (o *ApplicationInstance) GetLabelReadOnly() []Label`

GetLabelReadOnly returns the LabelReadOnly field if non-nil, zero value otherwise.

### GetLabelReadOnlyOk

`func (o *ApplicationInstance) GetLabelReadOnlyOk() (*[]Label, bool)`

GetLabelReadOnlyOk returns a tuple with the LabelReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelReadOnly

`func (o *ApplicationInstance) SetLabelReadOnly(v []Label)`

SetLabelReadOnly sets LabelReadOnly field to given value.


### GetLabelPublic

`func (o *ApplicationInstance) GetLabelPublic() []Label`

GetLabelPublic returns the LabelPublic field if non-nil, zero value otherwise.

### GetLabelPublicOk

`func (o *ApplicationInstance) GetLabelPublicOk() (*[]Label, bool)`

GetLabelPublicOk returns a tuple with the LabelPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPublic

`func (o *ApplicationInstance) SetLabelPublic(v []Label)`

SetLabelPublic sets LabelPublic field to given value.

### HasLabelPublic

`func (o *ApplicationInstance) HasLabelPublic() bool`

HasLabelPublic returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationInstance) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationInstance) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationInstance) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNumPlayersMax

`func (o *ApplicationInstance) GetNumPlayersMax() int32`

GetNumPlayersMax returns the NumPlayersMax field if non-nil, zero value otherwise.

### GetNumPlayersMaxOk

`func (o *ApplicationInstance) GetNumPlayersMaxOk() (*int32, bool)`

GetNumPlayersMaxOk returns a tuple with the NumPlayersMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPlayersMax

`func (o *ApplicationInstance) SetNumPlayersMax(v int32)`

SetNumPlayersMax sets NumPlayersMax field to given value.


### GetNumPlayers

`func (o *ApplicationInstance) GetNumPlayers() int32`

GetNumPlayers returns the NumPlayers field if non-nil, zero value otherwise.

### GetNumPlayersOk

`func (o *ApplicationInstance) GetNumPlayersOk() (*int32, bool)`

GetNumPlayersOk returns a tuple with the NumPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPlayers

`func (o *ApplicationInstance) SetNumPlayers(v int32)`

SetNumPlayers sets NumPlayers field to given value.


### GetLiveHostName

`func (o *ApplicationInstance) GetLiveHostName() string`

GetLiveHostName returns the LiveHostName field if non-nil, zero value otherwise.

### GetLiveHostNameOk

`func (o *ApplicationInstance) GetLiveHostNameOk() (*string, bool)`

GetLiveHostNameOk returns a tuple with the LiveHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHostName

`func (o *ApplicationInstance) SetLiveHostName(v string)`

SetLiveHostName sets LiveHostName field to given value.


### GetLiveMap

`func (o *ApplicationInstance) GetLiveMap() string`

GetLiveMap returns the LiveMap field if non-nil, zero value otherwise.

### GetLiveMapOk

`func (o *ApplicationInstance) GetLiveMapOk() (*string, bool)`

GetLiveMapOk returns a tuple with the LiveMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveMap

`func (o *ApplicationInstance) SetLiveMap(v string)`

SetLiveMap sets LiveMap field to given value.


### GetLiveGameVersion

`func (o *ApplicationInstance) GetLiveGameVersion() string`

GetLiveGameVersion returns the LiveGameVersion field if non-nil, zero value otherwise.

### GetLiveGameVersionOk

`func (o *ApplicationInstance) GetLiveGameVersionOk() (*string, bool)`

GetLiveGameVersionOk returns a tuple with the LiveGameVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveGameVersion

`func (o *ApplicationInstance) SetLiveGameVersion(v string)`

SetLiveGameVersion sets LiveGameVersion field to given value.


### GetUpdatedAt

`func (o *ApplicationInstance) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationInstance) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationInstance) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLiveRules

`func (o *ApplicationInstance) GetLiveRules() string`

GetLiveRules returns the LiveRules field if non-nil, zero value otherwise.

### GetLiveRulesOk

`func (o *ApplicationInstance) GetLiveRulesOk() (*string, bool)`

GetLiveRulesOk returns a tuple with the LiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveRules

`func (o *ApplicationInstance) SetLiveRules(v string)`

SetLiveRules sets LiveRules field to given value.


### GetAutoRestart

`func (o *ApplicationInstance) GetAutoRestart() int32`

GetAutoRestart returns the AutoRestart field if non-nil, zero value otherwise.

### GetAutoRestartOk

`func (o *ApplicationInstance) GetAutoRestartOk() (*int32, bool)`

GetAutoRestartOk returns a tuple with the AutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRestart

`func (o *ApplicationInstance) SetAutoRestart(v int32)`

SetAutoRestart sets AutoRestart field to given value.


### GetMarkedForDeletion

`func (o *ApplicationInstance) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *ApplicationInstance) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *ApplicationInstance) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetArcusAvailable

`func (o *ApplicationInstance) GetArcusAvailable() int32`

GetArcusAvailable returns the ArcusAvailable field if non-nil, zero value otherwise.

### GetArcusAvailableOk

`func (o *ApplicationInstance) GetArcusAvailableOk() (*int32, bool)`

GetArcusAvailableOk returns a tuple with the ArcusAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcusAvailable

`func (o *ApplicationInstance) SetArcusAvailable(v int32)`

SetArcusAvailable sets ArcusAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


