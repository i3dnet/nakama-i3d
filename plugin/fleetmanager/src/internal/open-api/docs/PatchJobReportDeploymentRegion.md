# PatchJobReportDeploymentRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRegion** | [**GenericModel**](GenericModel.md) | The Deployment Region&#39;s details | [readonly] 
**Progress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | Progress for just this Deployment Region | [readonly] 
**DcLocationData** | [**[]PatchJobReportDcLocation**](PatchJobReportDcLocation.md) | Patch job report data based on DC location | [readonly] 

## Methods

### NewPatchJobReportDeploymentRegion

`func NewPatchJobReportDeploymentRegion(deploymentRegion GenericModel, progress NullablePatchJobReportProgress, dcLocationData []PatchJobReportDcLocation, ) *PatchJobReportDeploymentRegion`

NewPatchJobReportDeploymentRegion instantiates a new PatchJobReportDeploymentRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobReportDeploymentRegionWithDefaults

`func NewPatchJobReportDeploymentRegionWithDefaults() *PatchJobReportDeploymentRegion`

NewPatchJobReportDeploymentRegionWithDefaults instantiates a new PatchJobReportDeploymentRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRegion

`func (o *PatchJobReportDeploymentRegion) GetDeploymentRegion() GenericModel`

GetDeploymentRegion returns the DeploymentRegion field if non-nil, zero value otherwise.

### GetDeploymentRegionOk

`func (o *PatchJobReportDeploymentRegion) GetDeploymentRegionOk() (*GenericModel, bool)`

GetDeploymentRegionOk returns a tuple with the DeploymentRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegion

`func (o *PatchJobReportDeploymentRegion) SetDeploymentRegion(v GenericModel)`

SetDeploymentRegion sets DeploymentRegion field to given value.


### GetProgress

`func (o *PatchJobReportDeploymentRegion) GetProgress() PatchJobReportProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *PatchJobReportDeploymentRegion) GetProgressOk() (*PatchJobReportProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *PatchJobReportDeploymentRegion) SetProgress(v PatchJobReportProgress)`

SetProgress sets Progress field to given value.


### SetProgressNil

`func (o *PatchJobReportDeploymentRegion) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *PatchJobReportDeploymentRegion) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetDcLocationData

`func (o *PatchJobReportDeploymentRegion) GetDcLocationData() []PatchJobReportDcLocation`

GetDcLocationData returns the DcLocationData field if non-nil, zero value otherwise.

### GetDcLocationDataOk

`func (o *PatchJobReportDeploymentRegion) GetDcLocationDataOk() (*[]PatchJobReportDcLocation, bool)`

GetDcLocationDataOk returns a tuple with the DcLocationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationData

`func (o *PatchJobReportDeploymentRegion) SetDcLocationData(v []PatchJobReportDcLocation)`

SetDcLocationData sets DcLocationData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


