# UCSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Ubisoft Cloud network setting id | [readonly] 
**ProjectId** | **string** | Ubisoft Cloud project id | 
**RegionName** | **string** | Ubisoft Cloud region name | 
**PublicNetworkId** | **string** | Ubisoft Cloud public network UUID id e.g: (b091a99e-abdd-11ea-bb37-0242ac130002) | 
**PrivateNetworkId** | **string** | Ubisoft Cloud private network UUID id e.g: (d80e2d7e-aa38-490e-915c-bb56222911c3) | 
**SecurityGroupId** | **string** | Ubisoft Cloud security group id e.g: (0ec61726-405c-4919-b9a1-9dcd6ae56077) | 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 
**ChangedAt** | **int32** | UnixTimestamp | [readonly] 

## Methods

### NewUCSetting

`func NewUCSetting(id string, projectId string, regionName string, publicNetworkId string, privateNetworkId string, securityGroupId string, createdAt int32, changedAt int32, ) *UCSetting`

NewUCSetting instantiates a new UCSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUCSettingWithDefaults

`func NewUCSettingWithDefaults() *UCSetting`

NewUCSettingWithDefaults instantiates a new UCSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UCSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UCSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UCSetting) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *UCSetting) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UCSetting) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UCSetting) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRegionName

`func (o *UCSetting) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *UCSetting) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *UCSetting) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetPublicNetworkId

`func (o *UCSetting) GetPublicNetworkId() string`

GetPublicNetworkId returns the PublicNetworkId field if non-nil, zero value otherwise.

### GetPublicNetworkIdOk

`func (o *UCSetting) GetPublicNetworkIdOk() (*string, bool)`

GetPublicNetworkIdOk returns a tuple with the PublicNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkId

`func (o *UCSetting) SetPublicNetworkId(v string)`

SetPublicNetworkId sets PublicNetworkId field to given value.


### GetPrivateNetworkId

`func (o *UCSetting) GetPrivateNetworkId() string`

GetPrivateNetworkId returns the PrivateNetworkId field if non-nil, zero value otherwise.

### GetPrivateNetworkIdOk

`func (o *UCSetting) GetPrivateNetworkIdOk() (*string, bool)`

GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkId

`func (o *UCSetting) SetPrivateNetworkId(v string)`

SetPrivateNetworkId sets PrivateNetworkId field to given value.


### GetSecurityGroupId

`func (o *UCSetting) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *UCSetting) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *UCSetting) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### GetCreatedAt

`func (o *UCSetting) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UCSetting) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UCSetting) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *UCSetting) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *UCSetting) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *UCSetting) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


