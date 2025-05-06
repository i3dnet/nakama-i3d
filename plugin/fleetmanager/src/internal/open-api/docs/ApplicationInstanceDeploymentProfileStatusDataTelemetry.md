# ApplicationInstanceDeploymentProfileStatusDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Unix timestamp of the data | [readonly] 
**DeploymentProfiles** | [**[]ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry**](ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry.md) | Contains the deployment profiles | [readonly] 

## Methods

### NewApplicationInstanceDeploymentProfileStatusDataTelemetry

`func NewApplicationInstanceDeploymentProfileStatusDataTelemetry(timestamp int32, deploymentProfiles []ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry, ) *ApplicationInstanceDeploymentProfileStatusDataTelemetry`

NewApplicationInstanceDeploymentProfileStatusDataTelemetry instantiates a new ApplicationInstanceDeploymentProfileStatusDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceDeploymentProfileStatusDataTelemetryWithDefaults

`func NewApplicationInstanceDeploymentProfileStatusDataTelemetryWithDefaults() *ApplicationInstanceDeploymentProfileStatusDataTelemetry`

NewApplicationInstanceDeploymentProfileStatusDataTelemetryWithDefaults instantiates a new ApplicationInstanceDeploymentProfileStatusDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetDeploymentProfiles

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) GetDeploymentProfiles() []ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry`

GetDeploymentProfiles returns the DeploymentProfiles field if non-nil, zero value otherwise.

### GetDeploymentProfilesOk

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) GetDeploymentProfilesOk() (*[]ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry, bool)`

GetDeploymentProfilesOk returns a tuple with the DeploymentProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProfiles

`func (o *ApplicationInstanceDeploymentProfileStatusDataTelemetry) SetDeploymentProfiles(v []ApplicationInstanceDeploymentProfileStatusDeploymentProfileTelemetry)`

SetDeploymentProfiles sets DeploymentProfiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


