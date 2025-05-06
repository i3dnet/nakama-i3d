# ApplicationInstanceByRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | The ID of the deployment region | [readonly] 
**AllocatableInstances** | **int32** | Amount of allocatable instances for the deployment region | [readonly] 
**ApplicationInstances** | **[]string** | List of application instances ID&#39;s | [readonly] 

## Methods

### NewApplicationInstanceByRegion

`func NewApplicationInstanceByRegion(regionId string, allocatableInstances int32, applicationInstances []string, ) *ApplicationInstanceByRegion`

NewApplicationInstanceByRegion instantiates a new ApplicationInstanceByRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceByRegionWithDefaults

`func NewApplicationInstanceByRegionWithDefaults() *ApplicationInstanceByRegion`

NewApplicationInstanceByRegionWithDefaults instantiates a new ApplicationInstanceByRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *ApplicationInstanceByRegion) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ApplicationInstanceByRegion) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ApplicationInstanceByRegion) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetAllocatableInstances

`func (o *ApplicationInstanceByRegion) GetAllocatableInstances() int32`

GetAllocatableInstances returns the AllocatableInstances field if non-nil, zero value otherwise.

### GetAllocatableInstancesOk

`func (o *ApplicationInstanceByRegion) GetAllocatableInstancesOk() (*int32, bool)`

GetAllocatableInstancesOk returns a tuple with the AllocatableInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatableInstances

`func (o *ApplicationInstanceByRegion) SetAllocatableInstances(v int32)`

SetAllocatableInstances sets AllocatableInstances field to given value.


### GetApplicationInstances

`func (o *ApplicationInstanceByRegion) GetApplicationInstances() []string`

GetApplicationInstances returns the ApplicationInstances field if non-nil, zero value otherwise.

### GetApplicationInstancesOk

`func (o *ApplicationInstanceByRegion) GetApplicationInstancesOk() (*[]string, bool)`

GetApplicationInstancesOk returns a tuple with the ApplicationInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstances

`func (o *ApplicationInstanceByRegion) SetApplicationInstances(v []string)`

SetApplicationInstances sets ApplicationInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


