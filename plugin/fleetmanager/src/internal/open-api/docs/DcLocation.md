# DcLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Datacenter location ID | [readonly] 
**ContinentId** | **int32** | Continent ID * 1: Africa * 2: Antarctica * 3: Asia * 4: Australia * 5: Europe * 6: Middle East * 7: North America * 8: South America | [readonly] 
**Country** | **string** | Country name | [readonly] 
**DisplayName** | **string** | The datacenter location&#39;s name | [readonly] 
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | [readonly] 
**AvailabilityZones** | **[]string** | Data center availability zones (only applies to cloud data centers) | [readonly] 
**RegionName** | **string** | Data center region name (only applies to cloud data centers) | [readonly] 
**PingSiteIds** | **[]string** | Host name for the corresponding data center location. | 

## Methods

### NewDcLocation

`func NewDcLocation(id int32, continentId int32, country string, displayName string, providerId int32, availabilityZones []string, regionName string, pingSiteIds []string, ) *DcLocation`

NewDcLocation instantiates a new DcLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcLocationWithDefaults

`func NewDcLocationWithDefaults() *DcLocation`

NewDcLocationWithDefaults instantiates a new DcLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DcLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DcLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DcLocation) SetId(v int32)`

SetId sets Id field to given value.


### GetContinentId

`func (o *DcLocation) GetContinentId() int32`

GetContinentId returns the ContinentId field if non-nil, zero value otherwise.

### GetContinentIdOk

`func (o *DcLocation) GetContinentIdOk() (*int32, bool)`

GetContinentIdOk returns a tuple with the ContinentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentId

`func (o *DcLocation) SetContinentId(v int32)`

SetContinentId sets ContinentId field to given value.


### GetCountry

`func (o *DcLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DcLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DcLocation) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetDisplayName

`func (o *DcLocation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DcLocation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DcLocation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProviderId

`func (o *DcLocation) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DcLocation) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DcLocation) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetAvailabilityZones

`func (o *DcLocation) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *DcLocation) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *DcLocation) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetRegionName

`func (o *DcLocation) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DcLocation) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DcLocation) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetPingSiteIds

`func (o *DcLocation) GetPingSiteIds() []string`

GetPingSiteIds returns the PingSiteIds field if non-nil, zero value otherwise.

### GetPingSiteIdsOk

`func (o *DcLocation) GetPingSiteIdsOk() (*[]string, bool)`

GetPingSiteIdsOk returns a tuple with the PingSiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSiteIds

`func (o *DcLocation) SetPingSiteIds(v []string)`

SetPingSiteIds sets PingSiteIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


