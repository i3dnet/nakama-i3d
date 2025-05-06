# PatchJobApplicationInstanceCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region name of application instances | [readonly] 
**DataCenter** | **string** | Data center location name | [readonly] 
**BuildVersion** | **string** | Application build version number | [readonly] 
**NumberOfInstances** | **int32** | Number of application instances | [readonly] 

## Methods

### NewPatchJobApplicationInstanceCount

`func NewPatchJobApplicationInstanceCount(region string, dataCenter string, buildVersion string, numberOfInstances int32, ) *PatchJobApplicationInstanceCount`

NewPatchJobApplicationInstanceCount instantiates a new PatchJobApplicationInstanceCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobApplicationInstanceCountWithDefaults

`func NewPatchJobApplicationInstanceCountWithDefaults() *PatchJobApplicationInstanceCount`

NewPatchJobApplicationInstanceCountWithDefaults instantiates a new PatchJobApplicationInstanceCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PatchJobApplicationInstanceCount) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchJobApplicationInstanceCount) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchJobApplicationInstanceCount) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDataCenter

`func (o *PatchJobApplicationInstanceCount) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *PatchJobApplicationInstanceCount) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *PatchJobApplicationInstanceCount) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetBuildVersion

`func (o *PatchJobApplicationInstanceCount) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *PatchJobApplicationInstanceCount) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *PatchJobApplicationInstanceCount) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.


### GetNumberOfInstances

`func (o *PatchJobApplicationInstanceCount) GetNumberOfInstances() int32`

GetNumberOfInstances returns the NumberOfInstances field if non-nil, zero value otherwise.

### GetNumberOfInstancesOk

`func (o *PatchJobApplicationInstanceCount) GetNumberOfInstancesOk() (*int32, bool)`

GetNumberOfInstancesOk returns a tuple with the NumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfInstances

`func (o *PatchJobApplicationInstanceCount) SetNumberOfInstances(v int32)`

SetNumberOfInstances sets NumberOfInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


