# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeId** | **int32** |  | [readonly] 
**CloudProviderId** | **int32** | The cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider)) | [readonly] 
**DcLocationId** | **int32** |  | [readonly] 
**InstanceType** | **string** |  | [readonly] 
**NumCores** | **int32** |  | [readonly] 
**Memory** | **int32** |  | [readonly] 
**Storage** | **string** |  | [readonly] 
**Status** | **int32** | It will show either instance type is available or not | [readonly] 

## Methods

### NewInstanceType

`func NewInstanceType(instanceTypeId int32, cloudProviderId int32, dcLocationId int32, instanceType string, numCores int32, memory int32, storage string, status int32, ) *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeId

`func (o *InstanceType) GetInstanceTypeId() int32`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *InstanceType) GetInstanceTypeIdOk() (*int32, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *InstanceType) SetInstanceTypeId(v int32)`

SetInstanceTypeId sets InstanceTypeId field to given value.


### GetCloudProviderId

`func (o *InstanceType) GetCloudProviderId() int32`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *InstanceType) GetCloudProviderIdOk() (*int32, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *InstanceType) SetCloudProviderId(v int32)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDcLocationId

`func (o *InstanceType) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *InstanceType) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *InstanceType) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetInstanceType

`func (o *InstanceType) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceType) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceType) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetNumCores

`func (o *InstanceType) GetNumCores() int32`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *InstanceType) GetNumCoresOk() (*int32, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *InstanceType) SetNumCores(v int32)`

SetNumCores sets NumCores field to given value.


### GetMemory

`func (o *InstanceType) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InstanceType) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InstanceType) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetStorage

`func (o *InstanceType) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceType) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceType) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetStatus

`func (o *InstanceType) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceType) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceType) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


