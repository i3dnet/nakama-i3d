# CpuPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 
**DcLocationId** | **int32** |  | 
**CpuPlatform** | **NullableString** |  | 

## Methods

### NewCpuPlatform

`func NewCpuPlatform(providerId int32, dcLocationId int32, cpuPlatform NullableString, ) *CpuPlatform`

NewCpuPlatform instantiates a new CpuPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuPlatformWithDefaults

`func NewCpuPlatformWithDefaults() *CpuPlatform`

NewCpuPlatformWithDefaults instantiates a new CpuPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CpuPlatform) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CpuPlatform) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CpuPlatform) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetDcLocationId

`func (o *CpuPlatform) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *CpuPlatform) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *CpuPlatform) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetCpuPlatform

`func (o *CpuPlatform) GetCpuPlatform() string`

GetCpuPlatform returns the CpuPlatform field if non-nil, zero value otherwise.

### GetCpuPlatformOk

`func (o *CpuPlatform) GetCpuPlatformOk() (*string, bool)`

GetCpuPlatformOk returns a tuple with the CpuPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPlatform

`func (o *CpuPlatform) SetCpuPlatform(v string)`

SetCpuPlatform sets CpuPlatform field to given value.


### SetCpuPlatformNil

`func (o *CpuPlatform) SetCpuPlatformNil(b bool)`

 SetCpuPlatformNil sets the value for CpuPlatform to be an explicit nil

### UnsetCpuPlatform
`func (o *CpuPlatform) UnsetCpuPlatform()`

UnsetCpuPlatform ensures that no value is present for CpuPlatform, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


