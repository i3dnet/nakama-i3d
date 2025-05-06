# CloudInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ProviderInstanceId** | **string** |  | [readonly] 
**ProviderImageId** | **string** |  | [readonly] 
**ProviderId** | **int32** | Cloud Provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | [readonly] 
**DediServerId** | **int32** |  | [readonly] 
**ApplicationId** | **int32** |  | [readonly] 
**FleetId** | **int32** |  | [readonly] 
**RegionId** | **int32** |  | [readonly] 
**DcLocationId** | **int32** |  | [readonly] 
**InstanceTypeId** | **int32** |  | [readonly] 
**AvailabilityZone** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Tag** | **string** |  | [readonly] 
**PublicIpAddress** | **string** |  | [readonly] 
**Status** | **int32** |  | [readonly] 
**CreatedAt** | **int32** | Unix timestamp | [readonly] 

## Methods

### NewCloudInstance

`func NewCloudInstance(id int32, providerInstanceId string, providerImageId string, providerId int32, dediServerId int32, applicationId int32, fleetId int32, regionId int32, dcLocationId int32, instanceTypeId int32, availabilityZone string, name string, tag string, publicIpAddress string, status int32, createdAt int32, ) *CloudInstance`

NewCloudInstance instantiates a new CloudInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstanceWithDefaults

`func NewCloudInstanceWithDefaults() *CloudInstance`

NewCloudInstanceWithDefaults instantiates a new CloudInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetProviderInstanceId

`func (o *CloudInstance) GetProviderInstanceId() string`

GetProviderInstanceId returns the ProviderInstanceId field if non-nil, zero value otherwise.

### GetProviderInstanceIdOk

`func (o *CloudInstance) GetProviderInstanceIdOk() (*string, bool)`

GetProviderInstanceIdOk returns a tuple with the ProviderInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInstanceId

`func (o *CloudInstance) SetProviderInstanceId(v string)`

SetProviderInstanceId sets ProviderInstanceId field to given value.


### GetProviderImageId

`func (o *CloudInstance) GetProviderImageId() string`

GetProviderImageId returns the ProviderImageId field if non-nil, zero value otherwise.

### GetProviderImageIdOk

`func (o *CloudInstance) GetProviderImageIdOk() (*string, bool)`

GetProviderImageIdOk returns a tuple with the ProviderImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderImageId

`func (o *CloudInstance) SetProviderImageId(v string)`

SetProviderImageId sets ProviderImageId field to given value.


### GetProviderId

`func (o *CloudInstance) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CloudInstance) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CloudInstance) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetDediServerId

`func (o *CloudInstance) GetDediServerId() int32`

GetDediServerId returns the DediServerId field if non-nil, zero value otherwise.

### GetDediServerIdOk

`func (o *CloudInstance) GetDediServerIdOk() (*int32, bool)`

GetDediServerIdOk returns a tuple with the DediServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDediServerId

`func (o *CloudInstance) SetDediServerId(v int32)`

SetDediServerId sets DediServerId field to given value.


### GetApplicationId

`func (o *CloudInstance) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CloudInstance) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CloudInstance) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetFleetId

`func (o *CloudInstance) GetFleetId() int32`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *CloudInstance) GetFleetIdOk() (*int32, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *CloudInstance) SetFleetId(v int32)`

SetFleetId sets FleetId field to given value.


### GetRegionId

`func (o *CloudInstance) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudInstance) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudInstance) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.


### GetDcLocationId

`func (o *CloudInstance) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *CloudInstance) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *CloudInstance) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetInstanceTypeId

`func (o *CloudInstance) GetInstanceTypeId() int32`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *CloudInstance) GetInstanceTypeIdOk() (*int32, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *CloudInstance) SetInstanceTypeId(v int32)`

SetInstanceTypeId sets InstanceTypeId field to given value.


### GetAvailabilityZone

`func (o *CloudInstance) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CloudInstance) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CloudInstance) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetName

`func (o *CloudInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudInstance) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *CloudInstance) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudInstance) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudInstance) SetTag(v string)`

SetTag sets Tag field to given value.


### GetPublicIpAddress

`func (o *CloudInstance) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *CloudInstance) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *CloudInstance) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.


### GetStatus

`func (o *CloudInstance) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudInstance) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudInstance) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CloudInstance) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudInstance) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudInstance) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


