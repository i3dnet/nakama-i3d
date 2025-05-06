# DeploymentContainerLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This deploymentContainerLocation&#39;s ID | [readonly] 
**CloudProviderId** | **int32** | The cloud provider ID of this location. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 
**DcLocationId** | **int32** | The ID of the data center. Find all possible values from [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/get_v3_cloud_dcLocation) | 
**PrimaryInstanceTypeName** | **NullableString** | The primary instance type to use. Find all possible values from [&#x60;GET /v3/cloud/instanceType&#x60;](#/Cloud/get_v3_cloud_instanceType) | 
**PrimaryInstanceTypeStatus** | **int32** | It will show either primary instance type is available or not | [readonly] 
**SecondaryInstanceTypeName** | Pointer to **NullableString** | The secondary instance type to use. Find all possible values from [&#x60;GET /v3/cloud/instanceType&#x60;](#/Cloud/get_v3_cloud_instanceType) | [optional] 
**SecondaryInstanceTypeStatus** | **int32** | It will show either secondary instance type is available or not | [readonly] 
**CpuPlatform** | Pointer to **NullableString** | The CPU platform to use for the given &#x60;cloudProviderId&#x60;. Find all possible values from [&#x60;GET /v3/cloud/cpuPlatform&#x60;](#/Cloud/get_v3_cloud_cpuPlatform). Null value means that a CPU platform is chosen automatically by the cloud provider. Note that this selector is not available on all cloud platforms | [optional] 
**MarkedForDeletion** | **int32** | If set to 1, all application instances will be gracefully removed in this cloud location. After all application instances are removed this object will also be deleted | [readonly] 

## Methods

### NewDeploymentContainerLocation

`func NewDeploymentContainerLocation(id string, cloudProviderId int32, dcLocationId int32, primaryInstanceTypeName NullableString, primaryInstanceTypeStatus int32, secondaryInstanceTypeStatus int32, markedForDeletion int32, ) *DeploymentContainerLocation`

NewDeploymentContainerLocation instantiates a new DeploymentContainerLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentContainerLocationWithDefaults

`func NewDeploymentContainerLocationWithDefaults() *DeploymentContainerLocation`

NewDeploymentContainerLocationWithDefaults instantiates a new DeploymentContainerLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentContainerLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentContainerLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentContainerLocation) SetId(v string)`

SetId sets Id field to given value.


### GetCloudProviderId

`func (o *DeploymentContainerLocation) GetCloudProviderId() int32`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DeploymentContainerLocation) GetCloudProviderIdOk() (*int32, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DeploymentContainerLocation) SetCloudProviderId(v int32)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDcLocationId

`func (o *DeploymentContainerLocation) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *DeploymentContainerLocation) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *DeploymentContainerLocation) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetPrimaryInstanceTypeName

`func (o *DeploymentContainerLocation) GetPrimaryInstanceTypeName() string`

GetPrimaryInstanceTypeName returns the PrimaryInstanceTypeName field if non-nil, zero value otherwise.

### GetPrimaryInstanceTypeNameOk

`func (o *DeploymentContainerLocation) GetPrimaryInstanceTypeNameOk() (*string, bool)`

GetPrimaryInstanceTypeNameOk returns a tuple with the PrimaryInstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceTypeName

`func (o *DeploymentContainerLocation) SetPrimaryInstanceTypeName(v string)`

SetPrimaryInstanceTypeName sets PrimaryInstanceTypeName field to given value.


### SetPrimaryInstanceTypeNameNil

`func (o *DeploymentContainerLocation) SetPrimaryInstanceTypeNameNil(b bool)`

 SetPrimaryInstanceTypeNameNil sets the value for PrimaryInstanceTypeName to be an explicit nil

### UnsetPrimaryInstanceTypeName
`func (o *DeploymentContainerLocation) UnsetPrimaryInstanceTypeName()`

UnsetPrimaryInstanceTypeName ensures that no value is present for PrimaryInstanceTypeName, not even an explicit nil
### GetPrimaryInstanceTypeStatus

`func (o *DeploymentContainerLocation) GetPrimaryInstanceTypeStatus() int32`

GetPrimaryInstanceTypeStatus returns the PrimaryInstanceTypeStatus field if non-nil, zero value otherwise.

### GetPrimaryInstanceTypeStatusOk

`func (o *DeploymentContainerLocation) GetPrimaryInstanceTypeStatusOk() (*int32, bool)`

GetPrimaryInstanceTypeStatusOk returns a tuple with the PrimaryInstanceTypeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceTypeStatus

`func (o *DeploymentContainerLocation) SetPrimaryInstanceTypeStatus(v int32)`

SetPrimaryInstanceTypeStatus sets PrimaryInstanceTypeStatus field to given value.


### GetSecondaryInstanceTypeName

`func (o *DeploymentContainerLocation) GetSecondaryInstanceTypeName() string`

GetSecondaryInstanceTypeName returns the SecondaryInstanceTypeName field if non-nil, zero value otherwise.

### GetSecondaryInstanceTypeNameOk

`func (o *DeploymentContainerLocation) GetSecondaryInstanceTypeNameOk() (*string, bool)`

GetSecondaryInstanceTypeNameOk returns a tuple with the SecondaryInstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryInstanceTypeName

`func (o *DeploymentContainerLocation) SetSecondaryInstanceTypeName(v string)`

SetSecondaryInstanceTypeName sets SecondaryInstanceTypeName field to given value.

### HasSecondaryInstanceTypeName

`func (o *DeploymentContainerLocation) HasSecondaryInstanceTypeName() bool`

HasSecondaryInstanceTypeName returns a boolean if a field has been set.

### SetSecondaryInstanceTypeNameNil

`func (o *DeploymentContainerLocation) SetSecondaryInstanceTypeNameNil(b bool)`

 SetSecondaryInstanceTypeNameNil sets the value for SecondaryInstanceTypeName to be an explicit nil

### UnsetSecondaryInstanceTypeName
`func (o *DeploymentContainerLocation) UnsetSecondaryInstanceTypeName()`

UnsetSecondaryInstanceTypeName ensures that no value is present for SecondaryInstanceTypeName, not even an explicit nil
### GetSecondaryInstanceTypeStatus

`func (o *DeploymentContainerLocation) GetSecondaryInstanceTypeStatus() int32`

GetSecondaryInstanceTypeStatus returns the SecondaryInstanceTypeStatus field if non-nil, zero value otherwise.

### GetSecondaryInstanceTypeStatusOk

`func (o *DeploymentContainerLocation) GetSecondaryInstanceTypeStatusOk() (*int32, bool)`

GetSecondaryInstanceTypeStatusOk returns a tuple with the SecondaryInstanceTypeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryInstanceTypeStatus

`func (o *DeploymentContainerLocation) SetSecondaryInstanceTypeStatus(v int32)`

SetSecondaryInstanceTypeStatus sets SecondaryInstanceTypeStatus field to given value.


### GetCpuPlatform

`func (o *DeploymentContainerLocation) GetCpuPlatform() string`

GetCpuPlatform returns the CpuPlatform field if non-nil, zero value otherwise.

### GetCpuPlatformOk

`func (o *DeploymentContainerLocation) GetCpuPlatformOk() (*string, bool)`

GetCpuPlatformOk returns a tuple with the CpuPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPlatform

`func (o *DeploymentContainerLocation) SetCpuPlatform(v string)`

SetCpuPlatform sets CpuPlatform field to given value.

### HasCpuPlatform

`func (o *DeploymentContainerLocation) HasCpuPlatform() bool`

HasCpuPlatform returns a boolean if a field has been set.

### SetCpuPlatformNil

`func (o *DeploymentContainerLocation) SetCpuPlatformNil(b bool)`

 SetCpuPlatformNil sets the value for CpuPlatform to be an explicit nil

### UnsetCpuPlatform
`func (o *DeploymentContainerLocation) UnsetCpuPlatform()`

UnsetCpuPlatform ensures that no value is present for CpuPlatform, not even an explicit nil
### GetMarkedForDeletion

`func (o *DeploymentContainerLocation) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *DeploymentContainerLocation) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *DeploymentContainerLocation) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


