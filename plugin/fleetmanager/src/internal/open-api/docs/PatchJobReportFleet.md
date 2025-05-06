# PatchJobReportFleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fleet** | [**GenericModel**](GenericModel.md) | The Fleet details | [readonly] 
**Progress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | Progress for just this fleet | [readonly] 
**DeploymentRegionData** | [**[]PatchJobReportDeploymentRegion**](PatchJobReportDeploymentRegion.md) | Patch job report data based on deployment region | [readonly] 

## Methods

### NewPatchJobReportFleet

`func NewPatchJobReportFleet(fleet GenericModel, progress NullablePatchJobReportProgress, deploymentRegionData []PatchJobReportDeploymentRegion, ) *PatchJobReportFleet`

NewPatchJobReportFleet instantiates a new PatchJobReportFleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobReportFleetWithDefaults

`func NewPatchJobReportFleetWithDefaults() *PatchJobReportFleet`

NewPatchJobReportFleetWithDefaults instantiates a new PatchJobReportFleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleet

`func (o *PatchJobReportFleet) GetFleet() GenericModel`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *PatchJobReportFleet) GetFleetOk() (*GenericModel, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *PatchJobReportFleet) SetFleet(v GenericModel)`

SetFleet sets Fleet field to given value.


### GetProgress

`func (o *PatchJobReportFleet) GetProgress() PatchJobReportProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *PatchJobReportFleet) GetProgressOk() (*PatchJobReportProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *PatchJobReportFleet) SetProgress(v PatchJobReportProgress)`

SetProgress sets Progress field to given value.


### SetProgressNil

`func (o *PatchJobReportFleet) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *PatchJobReportFleet) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetDeploymentRegionData

`func (o *PatchJobReportFleet) GetDeploymentRegionData() []PatchJobReportDeploymentRegion`

GetDeploymentRegionData returns the DeploymentRegionData field if non-nil, zero value otherwise.

### GetDeploymentRegionDataOk

`func (o *PatchJobReportFleet) GetDeploymentRegionDataOk() (*[]PatchJobReportDeploymentRegion, bool)`

GetDeploymentRegionDataOk returns a tuple with the DeploymentRegionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRegionData

`func (o *PatchJobReportFleet) SetDeploymentRegionData(v []PatchJobReportDeploymentRegion)`

SetDeploymentRegionData sets DeploymentRegionData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


