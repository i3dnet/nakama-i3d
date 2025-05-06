# PatchJobReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchJobId** | **string** | Patch job ID | [readonly] 
**PatchJobName** | **string** | Name of the patch job | [readonly] 
**ChangedAt** | **int32** | Unix timestamp when the status of the patch job has been changed | [readonly] 
**FinishedAt** | **int32** | Unix timestamp when the patch job finished (if &#x60;0&#x60; then the patch job is still running) | [readonly] 
**StartedAt** | **int32** | Unix timestamp when the patch job started | [readonly] 
**PatchJobMethodName** | **string** | Name of the patch job method (strategy) | [readonly] 
**PatchJobStatus** | **int32** | Status of the patch job | [readonly] 
**DeploymentEnvironment** | [**GenericModel**](GenericModel.md) | The deployment environment | [readonly] 
**OldApplicationBuilds** | [**[]PatchJobReportApplicationBuild**](PatchJobReportApplicationBuild.md) | List of old application builds | [readonly] 
**NewApplicationBuild** | [**PatchJobReportApplicationBuild**](PatchJobReportApplicationBuild.md) | The new application build | [readonly] 
**Progress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | Progress across all fleets - total for the entire patch job | [readonly] 
**FleetData** | [**[]PatchJobReportFleet**](PatchJobReportFleet.md) | Patch job report data based on fleet | [readonly] 

## Methods

### NewPatchJobReport

`func NewPatchJobReport(patchJobId string, patchJobName string, changedAt int32, finishedAt int32, startedAt int32, patchJobMethodName string, patchJobStatus int32, deploymentEnvironment GenericModel, oldApplicationBuilds []PatchJobReportApplicationBuild, newApplicationBuild PatchJobReportApplicationBuild, progress NullablePatchJobReportProgress, fleetData []PatchJobReportFleet, ) *PatchJobReport`

NewPatchJobReport instantiates a new PatchJobReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobReportWithDefaults

`func NewPatchJobReportWithDefaults() *PatchJobReport`

NewPatchJobReportWithDefaults instantiates a new PatchJobReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchJobId

`func (o *PatchJobReport) GetPatchJobId() string`

GetPatchJobId returns the PatchJobId field if non-nil, zero value otherwise.

### GetPatchJobIdOk

`func (o *PatchJobReport) GetPatchJobIdOk() (*string, bool)`

GetPatchJobIdOk returns a tuple with the PatchJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobId

`func (o *PatchJobReport) SetPatchJobId(v string)`

SetPatchJobId sets PatchJobId field to given value.


### GetPatchJobName

`func (o *PatchJobReport) GetPatchJobName() string`

GetPatchJobName returns the PatchJobName field if non-nil, zero value otherwise.

### GetPatchJobNameOk

`func (o *PatchJobReport) GetPatchJobNameOk() (*string, bool)`

GetPatchJobNameOk returns a tuple with the PatchJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobName

`func (o *PatchJobReport) SetPatchJobName(v string)`

SetPatchJobName sets PatchJobName field to given value.


### GetChangedAt

`func (o *PatchJobReport) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *PatchJobReport) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *PatchJobReport) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.


### GetFinishedAt

`func (o *PatchJobReport) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *PatchJobReport) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *PatchJobReport) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### GetStartedAt

`func (o *PatchJobReport) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PatchJobReport) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PatchJobReport) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.


### GetPatchJobMethodName

`func (o *PatchJobReport) GetPatchJobMethodName() string`

GetPatchJobMethodName returns the PatchJobMethodName field if non-nil, zero value otherwise.

### GetPatchJobMethodNameOk

`func (o *PatchJobReport) GetPatchJobMethodNameOk() (*string, bool)`

GetPatchJobMethodNameOk returns a tuple with the PatchJobMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobMethodName

`func (o *PatchJobReport) SetPatchJobMethodName(v string)`

SetPatchJobMethodName sets PatchJobMethodName field to given value.


### GetPatchJobStatus

`func (o *PatchJobReport) GetPatchJobStatus() int32`

GetPatchJobStatus returns the PatchJobStatus field if non-nil, zero value otherwise.

### GetPatchJobStatusOk

`func (o *PatchJobReport) GetPatchJobStatusOk() (*int32, bool)`

GetPatchJobStatusOk returns a tuple with the PatchJobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobStatus

`func (o *PatchJobReport) SetPatchJobStatus(v int32)`

SetPatchJobStatus sets PatchJobStatus field to given value.


### GetDeploymentEnvironment

`func (o *PatchJobReport) GetDeploymentEnvironment() GenericModel`

GetDeploymentEnvironment returns the DeploymentEnvironment field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentOk

`func (o *PatchJobReport) GetDeploymentEnvironmentOk() (*GenericModel, bool)`

GetDeploymentEnvironmentOk returns a tuple with the DeploymentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironment

`func (o *PatchJobReport) SetDeploymentEnvironment(v GenericModel)`

SetDeploymentEnvironment sets DeploymentEnvironment field to given value.


### GetOldApplicationBuilds

`func (o *PatchJobReport) GetOldApplicationBuilds() []PatchJobReportApplicationBuild`

GetOldApplicationBuilds returns the OldApplicationBuilds field if non-nil, zero value otherwise.

### GetOldApplicationBuildsOk

`func (o *PatchJobReport) GetOldApplicationBuildsOk() (*[]PatchJobReportApplicationBuild, bool)`

GetOldApplicationBuildsOk returns a tuple with the OldApplicationBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldApplicationBuilds

`func (o *PatchJobReport) SetOldApplicationBuilds(v []PatchJobReportApplicationBuild)`

SetOldApplicationBuilds sets OldApplicationBuilds field to given value.


### GetNewApplicationBuild

`func (o *PatchJobReport) GetNewApplicationBuild() PatchJobReportApplicationBuild`

GetNewApplicationBuild returns the NewApplicationBuild field if non-nil, zero value otherwise.

### GetNewApplicationBuildOk

`func (o *PatchJobReport) GetNewApplicationBuildOk() (*PatchJobReportApplicationBuild, bool)`

GetNewApplicationBuildOk returns a tuple with the NewApplicationBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplicationBuild

`func (o *PatchJobReport) SetNewApplicationBuild(v PatchJobReportApplicationBuild)`

SetNewApplicationBuild sets NewApplicationBuild field to given value.


### GetProgress

`func (o *PatchJobReport) GetProgress() PatchJobReportProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *PatchJobReport) GetProgressOk() (*PatchJobReportProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *PatchJobReport) SetProgress(v PatchJobReportProgress)`

SetProgress sets Progress field to given value.


### SetProgressNil

`func (o *PatchJobReport) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *PatchJobReport) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetFleetData

`func (o *PatchJobReport) GetFleetData() []PatchJobReportFleet`

GetFleetData returns the FleetData field if non-nil, zero value otherwise.

### GetFleetDataOk

`func (o *PatchJobReport) GetFleetDataOk() (*[]PatchJobReportFleet, bool)`

GetFleetDataOk returns a tuple with the FleetData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetData

`func (o *PatchJobReport) SetFleetData(v []PatchJobReportFleet)`

SetFleetData sets FleetData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


