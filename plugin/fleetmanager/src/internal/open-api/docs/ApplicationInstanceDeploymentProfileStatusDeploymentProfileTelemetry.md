# ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentProfileId** | **string** | The ID of the deployment profile | [readonly] 
**Setup** | **int32** | Average number of application instances with setup status | [readonly] 
**Offline** | **int32** | Average number of application instances with offline status | [readonly] 
**Starting** | **int32** | Average number of application instances with starting status | [readonly] 
**Online** | **int32** | Average number of application instances with online status | [readonly] 
**Allocated** | **int32** | Average number of application instances with allocated status | [readonly] 
**DeploymentRegions** | [**[]ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry**](ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry.md) | Contains the deployment regions | [readonly] 

## Methods

### NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry

`func NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry(deploymentProfileId string, setup int32, offline int32, starting int32, online int32, allocated int32, deploymentRegions []ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry, ) *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry`

NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry instantiates a new ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetryWithDefaults

`func NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetryWithDefaults() *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry`

NewApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetryWithDefaults instantiates a new ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentProfileId

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetDeploymentProfileId() string`

GetDeploymentProfileId returns the DeploymentProfileId field if non-nil, zero value otherwise.

### GetDeploymentProfileIdOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetDeploymentProfileIdOk() (*string, bool)`

GetDeploymentProfileIdOk returns a tuple with the DeploymentProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProfileId

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetDeploymentProfileId(v string)`

SetDeploymentProfileId sets DeploymentProfileId field to given value.


### GetSetup

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetSetup() int32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetSetupOk() (*int32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetSetup(v int32)`

SetSetup sets Setup field to given value.


### GetOffline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetOffline() int32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetOfflineOk() (*int32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetOffline(v int32)`

SetOffline sets Offline field to given value.


### GetStarting

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetStarting() int32`

GetStarting returns the Starting field if non-nil, zero value otherwise.

### GetStartingOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetStartingOk() (*int32, bool)`

GetStartingOk returns a tuple with the Starting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarting

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetStarting(v int32)`

SetStarting sets Starting field to given value.


### GetOnline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetOnline() int32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetOnlineOk() (*int32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetOnline(v int32)`

SetOnline sets Online field to given value.


### GetAllocated

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.


### GetDeploymentRegions

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetDeploymentRegions() []ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry`

GetDeploymentRegions returns the DeploymentRegions field if non-nil, zero value otherwise.

### GetDeploymentRegionsOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) GetDeploymentRegionsOk() (*[]ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry, bool)`

GetDeploymentRegionsOk returns a tuple with the DeploymentRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegions

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry) SetDeploymentRegions(v []ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry)`

SetDeploymentRegions sets DeploymentRegions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


