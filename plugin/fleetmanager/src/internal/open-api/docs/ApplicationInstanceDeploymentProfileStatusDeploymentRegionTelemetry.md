# ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRegionId** | **string** | The ID of the deployment region | [readonly] 
**Setup** | **int32** | Average number of application instances with setup status | [readonly] 
**Offline** | **int32** | Average number of application instances with offline status | [readonly] 
**Starting** | **int32** | Average number of application instances with starting status | [readonly] 
**Online** | **int32** | Average number of application instances with online status | [readonly] 
**Allocated** | **int32** | Average number of application instances with allocated status | [readonly] 

## Methods

### NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry

`func NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry(deploymentRegionId string, setup int32, offline int32, starting int32, online int32, allocated int32, ) *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry`

NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry instantiates a new ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetryWithDefaults

`func NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetryWithDefaults() *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry`

NewApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetryWithDefaults instantiates a new ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRegionId

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetDeploymentRegionId() string`

GetDeploymentRegionId returns the DeploymentRegionId field if non-nil, zero value otherwise.

### GetDeploymentRegionIdOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetDeploymentRegionIdOk() (*string, bool)`

GetDeploymentRegionIdOk returns a tuple with the DeploymentRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegionId

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetDeploymentRegionId(v string)`

SetDeploymentRegionId sets DeploymentRegionId field to given value.


### GetSetup

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetSetup() int32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetSetupOk() (*int32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetSetup(v int32)`

SetSetup sets Setup field to given value.


### GetOffline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetOffline() int32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetOfflineOk() (*int32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetOffline(v int32)`

SetOffline sets Offline field to given value.


### GetStarting

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetStarting() int32`

GetStarting returns the Starting field if non-nil, zero value otherwise.

### GetStartingOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetStartingOk() (*int32, bool)`

GetStartingOk returns a tuple with the Starting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarting

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetStarting(v int32)`

SetStarting sets Starting field to given value.


### GetOnline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetOnline() int32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetOnlineOk() (*int32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetOnline(v int32)`

SetOnline sets Online field to given value.


### GetAllocated

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *ApplicationInstanceDeploymentProfileStatusDeploymentRegionTelemetry) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


