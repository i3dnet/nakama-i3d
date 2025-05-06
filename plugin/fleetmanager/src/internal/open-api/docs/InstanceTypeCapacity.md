# InstanceTypeCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Instance Type Capacity | [readonly] 
**HostCapacityTemplateId** | **string** | The ID of the Host Capacity Template to which this element is linked | [readonly] 
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 
**InstanceType** | **string** | The instance type to which the &#x60;capacity&#x60; applies. This is the name of instance types found at various cloud providers. See GET /cloud/instanceType (for virtual servers) and GET /host/instanceType (for bare metal servers) | 
**IsVirtual** | **int32** | &#x60;1&#x60; if the instance type is a virtual server, &#x60;0&#x60; if it is a bare metal server | [readonly] 
**Capacity** | **int32** | The capacity of the instance type | 
**CreatedAt** | **int32** | The Unix timestamp at which this element is created | [readonly] 

## Methods

### NewInstanceTypeCapacity

`func NewInstanceTypeCapacity(id string, hostCapacityTemplateId string, providerId int32, instanceType string, isVirtual int32, capacity int32, createdAt int32, ) *InstanceTypeCapacity`

NewInstanceTypeCapacity instantiates a new InstanceTypeCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeCapacityWithDefaults

`func NewInstanceTypeCapacityWithDefaults() *InstanceTypeCapacity`

NewInstanceTypeCapacityWithDefaults instantiates a new InstanceTypeCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeCapacity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeCapacity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeCapacity) SetId(v string)`

SetId sets Id field to given value.


### GetHostCapacityTemplateId

`func (o *InstanceTypeCapacity) GetHostCapacityTemplateId() string`

GetHostCapacityTemplateId returns the HostCapacityTemplateId field if non-nil, zero value otherwise.

### GetHostCapacityTemplateIdOk

`func (o *InstanceTypeCapacity) GetHostCapacityTemplateIdOk() (*string, bool)`

GetHostCapacityTemplateIdOk returns a tuple with the HostCapacityTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCapacityTemplateId

`func (o *InstanceTypeCapacity) SetHostCapacityTemplateId(v string)`

SetHostCapacityTemplateId sets HostCapacityTemplateId field to given value.


### GetProviderId

`func (o *InstanceTypeCapacity) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *InstanceTypeCapacity) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *InstanceTypeCapacity) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetInstanceType

`func (o *InstanceTypeCapacity) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceTypeCapacity) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceTypeCapacity) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetIsVirtual

`func (o *InstanceTypeCapacity) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *InstanceTypeCapacity) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *InstanceTypeCapacity) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetCapacity

`func (o *InstanceTypeCapacity) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *InstanceTypeCapacity) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *InstanceTypeCapacity) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *InstanceTypeCapacity) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceTypeCapacity) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceTypeCapacity) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


