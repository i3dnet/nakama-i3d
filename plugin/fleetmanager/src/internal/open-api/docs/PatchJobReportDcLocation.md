# PatchJobReportDcLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DcLocation** | [**GenericIntModel**](GenericIntModel.md) | The DC Location&#39;s details | [readonly] 
**DcLocationRegionName** | **string** | Region name of the DC location | [readonly] 
**Progress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | Progress for just this DC Location | [readonly] 

## Methods

### NewPatchJobReportDcLocation

`func NewPatchJobReportDcLocation(dcLocation GenericIntModel, dcLocationRegionName string, progress NullablePatchJobReportProgress, ) *PatchJobReportDcLocation`

NewPatchJobReportDcLocation instantiates a new PatchJobReportDcLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobReportDcLocationWithDefaults

`func NewPatchJobReportDcLocationWithDefaults() *PatchJobReportDcLocation`

NewPatchJobReportDcLocationWithDefaults instantiates a new PatchJobReportDcLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDcLocation

`func (o *PatchJobReportDcLocation) GetDcLocation() GenericIntModel`

GetDcLocation returns the DcLocation field if non-nil, zero value otherwise.

### GetDcLocationOk

`func (o *PatchJobReportDcLocation) GetDcLocationOk() (*GenericIntModel, bool)`

GetDcLocationOk returns a tuple with the DcLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocation

`func (o *PatchJobReportDcLocation) SetDcLocation(v GenericIntModel)`

SetDcLocation sets DcLocation field to given value.


### GetDcLocationRegionName

`func (o *PatchJobReportDcLocation) GetDcLocationRegionName() string`

GetDcLocationRegionName returns the DcLocationRegionName field if non-nil, zero value otherwise.

### GetDcLocationRegionNameOk

`func (o *PatchJobReportDcLocation) GetDcLocationRegionNameOk() (*string, bool)`

GetDcLocationRegionNameOk returns a tuple with the DcLocationRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationRegionName

`func (o *PatchJobReportDcLocation) SetDcLocationRegionName(v string)`

SetDcLocationRegionName sets DcLocationRegionName field to given value.


### GetProgress

`func (o *PatchJobReportDcLocation) GetProgress() PatchJobReportProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *PatchJobReportDcLocation) GetProgressOk() (*PatchJobReportProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *PatchJobReportDcLocation) SetProgress(v PatchJobReportProgress)`

SetProgress sets Progress field to given value.


### SetProgressNil

`func (o *PatchJobReportDcLocation) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *PatchJobReportDcLocation) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


