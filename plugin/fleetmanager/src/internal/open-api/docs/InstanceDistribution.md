# InstanceDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BareMetal** | **int32** | Amount of game instances on the bare metal | [readonly] 
**BareMetalPercentage** | **int32** | Percentage of game instances on the bare metal | [readonly] 
**FlexMetal** | **int32** | Amount of game instances on the flex metal | [readonly] 
**FlexMetalPercentage** | **int32** | Percentage of game instances on the flex metal | [readonly] 
**Cloud** | **int32** | Amount of game instances in the cloud | [readonly] 
**CloudPercentage** | **int32** | Percentage of game instances in the cloud | [readonly] 
**Total** | **int32** | Total amount of game instances | [readonly] 

## Methods

### NewInstanceDistribution

`func NewInstanceDistribution(bareMetal int32, bareMetalPercentage int32, flexMetal int32, flexMetalPercentage int32, cloud int32, cloudPercentage int32, total int32, ) *InstanceDistribution`

NewInstanceDistribution instantiates a new InstanceDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDistributionWithDefaults

`func NewInstanceDistributionWithDefaults() *InstanceDistribution`

NewInstanceDistributionWithDefaults instantiates a new InstanceDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBareMetal

`func (o *InstanceDistribution) GetBareMetal() int32`

GetBareMetal returns the BareMetal field if non-nil, zero value otherwise.

### GetBareMetalOk

`func (o *InstanceDistribution) GetBareMetalOk() (*int32, bool)`

GetBareMetalOk returns a tuple with the BareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetal

`func (o *InstanceDistribution) SetBareMetal(v int32)`

SetBareMetal sets BareMetal field to given value.


### GetBareMetalPercentage

`func (o *InstanceDistribution) GetBareMetalPercentage() int32`

GetBareMetalPercentage returns the BareMetalPercentage field if non-nil, zero value otherwise.

### GetBareMetalPercentageOk

`func (o *InstanceDistribution) GetBareMetalPercentageOk() (*int32, bool)`

GetBareMetalPercentageOk returns a tuple with the BareMetalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalPercentage

`func (o *InstanceDistribution) SetBareMetalPercentage(v int32)`

SetBareMetalPercentage sets BareMetalPercentage field to given value.


### GetFlexMetal

`func (o *InstanceDistribution) GetFlexMetal() int32`

GetFlexMetal returns the FlexMetal field if non-nil, zero value otherwise.

### GetFlexMetalOk

`func (o *InstanceDistribution) GetFlexMetalOk() (*int32, bool)`

GetFlexMetalOk returns a tuple with the FlexMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMetal

`func (o *InstanceDistribution) SetFlexMetal(v int32)`

SetFlexMetal sets FlexMetal field to given value.


### GetFlexMetalPercentage

`func (o *InstanceDistribution) GetFlexMetalPercentage() int32`

GetFlexMetalPercentage returns the FlexMetalPercentage field if non-nil, zero value otherwise.

### GetFlexMetalPercentageOk

`func (o *InstanceDistribution) GetFlexMetalPercentageOk() (*int32, bool)`

GetFlexMetalPercentageOk returns a tuple with the FlexMetalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMetalPercentage

`func (o *InstanceDistribution) SetFlexMetalPercentage(v int32)`

SetFlexMetalPercentage sets FlexMetalPercentage field to given value.


### GetCloud

`func (o *InstanceDistribution) GetCloud() int32`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceDistribution) GetCloudOk() (*int32, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceDistribution) SetCloud(v int32)`

SetCloud sets Cloud field to given value.


### GetCloudPercentage

`func (o *InstanceDistribution) GetCloudPercentage() int32`

GetCloudPercentage returns the CloudPercentage field if non-nil, zero value otherwise.

### GetCloudPercentageOk

`func (o *InstanceDistribution) GetCloudPercentageOk() (*int32, bool)`

GetCloudPercentageOk returns a tuple with the CloudPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPercentage

`func (o *InstanceDistribution) SetCloudPercentage(v int32)`

SetCloudPercentage sets CloudPercentage field to given value.


### GetTotal

`func (o *InstanceDistribution) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InstanceDistribution) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InstanceDistribution) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


