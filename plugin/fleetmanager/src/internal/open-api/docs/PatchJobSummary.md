# PatchJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Patch job ID | [readonly] 
**DeploymentEnvironmentId** | **string** | Deployment environment ID | [readonly] 
**ApplicationId** | **string** | Application ID | [readonly] 
**ApplicationName** | **string** | Name of the application | [readonly] 
**NewApplicationBuildName** | **string** | Name of all the application builds that will be patched | [readonly] 
**PatchJobMethodId** | **int32** | Patch job type, you can find the list available patch job type [&#x60;GET /v3/patchJob/type&#x60;](game-publisher#/PatchJob/get_v3_patchJob_type) | [readonly] 
**PatchJobMethodName** | **string** | Name of the patch job chosen | [readonly] 
**PatchJobName** | **string** | Name of the patch job (unique) | [readonly] 
**Status** | **int32** | Status of the patch job | [readonly] 
**Comments** | **string** | Comments for patch job | [readonly] 
**SchedulerStartTime** | **int32** | When the patch job needs to start, default value is 0 this means it’s start directly after creating the patch job, else there needs to give a timestamp in UTC when the job needs to start | [readonly] 
**PatchJobStartTime** | **int32** | Timestamp when the patch job has been started in UTC | [readonly] 
**FinishedAt** | **int32** | Timestamp when the patch job has been finished in UTC | [readonly] 
**IsRevertable** | **int32** | If the patch job can be cancelled | [readonly] 
**PatchJobOverallProgress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | How many application instances are being patched | [readonly] 

## Methods

### NewPatchJobSummary

`func NewPatchJobSummary(id string, deploymentEnvironmentId string, applicationId string, applicationName string, newApplicationBuildName string, patchJobMethodId int32, patchJobMethodName string, patchJobName string, status int32, comments string, schedulerStartTime int32, patchJobStartTime int32, finishedAt int32, isRevertable int32, patchJobOverallProgress NullablePatchJobReportProgress, ) *PatchJobSummary`

NewPatchJobSummary instantiates a new PatchJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobSummaryWithDefaults

`func NewPatchJobSummaryWithDefaults() *PatchJobSummary`

NewPatchJobSummaryWithDefaults instantiates a new PatchJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJobSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJobSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJobSummary) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentEnvironmentId

`func (o *PatchJobSummary) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *PatchJobSummary) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *PatchJobSummary) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetApplicationId

`func (o *PatchJobSummary) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PatchJobSummary) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PatchJobSummary) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *PatchJobSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *PatchJobSummary) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *PatchJobSummary) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetNewApplicationBuildName

`func (o *PatchJobSummary) GetNewApplicationBuildName() string`

GetNewApplicationBuildName returns the NewApplicationBuildName field if non-nil, zero value otherwise.

### GetNewApplicationBuildNameOk

`func (o *PatchJobSummary) GetNewApplicationBuildNameOk() (*string, bool)`

GetNewApplicationBuildNameOk returns a tuple with the NewApplicationBuildName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplicationBuildName

`func (o *PatchJobSummary) SetNewApplicationBuildName(v string)`

SetNewApplicationBuildName sets NewApplicationBuildName field to given value.


### GetPatchJobMethodId

`func (o *PatchJobSummary) GetPatchJobMethodId() int32`

GetPatchJobMethodId returns the PatchJobMethodId field if non-nil, zero value otherwise.

### GetPatchJobMethodIdOk

`func (o *PatchJobSummary) GetPatchJobMethodIdOk() (*int32, bool)`

GetPatchJobMethodIdOk returns a tuple with the PatchJobMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobMethodId

`func (o *PatchJobSummary) SetPatchJobMethodId(v int32)`

SetPatchJobMethodId sets PatchJobMethodId field to given value.


### GetPatchJobMethodName

`func (o *PatchJobSummary) GetPatchJobMethodName() string`

GetPatchJobMethodName returns the PatchJobMethodName field if non-nil, zero value otherwise.

### GetPatchJobMethodNameOk

`func (o *PatchJobSummary) GetPatchJobMethodNameOk() (*string, bool)`

GetPatchJobMethodNameOk returns a tuple with the PatchJobMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobMethodName

`func (o *PatchJobSummary) SetPatchJobMethodName(v string)`

SetPatchJobMethodName sets PatchJobMethodName field to given value.


### GetPatchJobName

`func (o *PatchJobSummary) GetPatchJobName() string`

GetPatchJobName returns the PatchJobName field if non-nil, zero value otherwise.

### GetPatchJobNameOk

`func (o *PatchJobSummary) GetPatchJobNameOk() (*string, bool)`

GetPatchJobNameOk returns a tuple with the PatchJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobName

`func (o *PatchJobSummary) SetPatchJobName(v string)`

SetPatchJobName sets PatchJobName field to given value.


### GetStatus

`func (o *PatchJobSummary) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchJobSummary) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchJobSummary) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetComments

`func (o *PatchJobSummary) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchJobSummary) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchJobSummary) SetComments(v string)`

SetComments sets Comments field to given value.


### GetSchedulerStartTime

`func (o *PatchJobSummary) GetSchedulerStartTime() int32`

GetSchedulerStartTime returns the SchedulerStartTime field if non-nil, zero value otherwise.

### GetSchedulerStartTimeOk

`func (o *PatchJobSummary) GetSchedulerStartTimeOk() (*int32, bool)`

GetSchedulerStartTimeOk returns a tuple with the SchedulerStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerStartTime

`func (o *PatchJobSummary) SetSchedulerStartTime(v int32)`

SetSchedulerStartTime sets SchedulerStartTime field to given value.


### GetPatchJobStartTime

`func (o *PatchJobSummary) GetPatchJobStartTime() int32`

GetPatchJobStartTime returns the PatchJobStartTime field if non-nil, zero value otherwise.

### GetPatchJobStartTimeOk

`func (o *PatchJobSummary) GetPatchJobStartTimeOk() (*int32, bool)`

GetPatchJobStartTimeOk returns a tuple with the PatchJobStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobStartTime

`func (o *PatchJobSummary) SetPatchJobStartTime(v int32)`

SetPatchJobStartTime sets PatchJobStartTime field to given value.


### GetFinishedAt

`func (o *PatchJobSummary) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *PatchJobSummary) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *PatchJobSummary) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### GetIsRevertable

`func (o *PatchJobSummary) GetIsRevertable() int32`

GetIsRevertable returns the IsRevertable field if non-nil, zero value otherwise.

### GetIsRevertableOk

`func (o *PatchJobSummary) GetIsRevertableOk() (*int32, bool)`

GetIsRevertableOk returns a tuple with the IsRevertable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevertable

`func (o *PatchJobSummary) SetIsRevertable(v int32)`

SetIsRevertable sets IsRevertable field to given value.


### GetPatchJobOverallProgress

`func (o *PatchJobSummary) GetPatchJobOverallProgress() PatchJobReportProgress`

GetPatchJobOverallProgress returns the PatchJobOverallProgress field if non-nil, zero value otherwise.

### GetPatchJobOverallProgressOk

`func (o *PatchJobSummary) GetPatchJobOverallProgressOk() (*PatchJobReportProgress, bool)`

GetPatchJobOverallProgressOk returns a tuple with the PatchJobOverallProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobOverallProgress

`func (o *PatchJobSummary) SetPatchJobOverallProgress(v PatchJobReportProgress)`

SetPatchJobOverallProgress sets PatchJobOverallProgress field to given value.


### SetPatchJobOverallProgressNil

`func (o *PatchJobSummary) SetPatchJobOverallProgressNil(b bool)`

 SetPatchJobOverallProgressNil sets the value for PatchJobOverallProgress to be an explicit nil

### UnsetPatchJobOverallProgress
`func (o *PatchJobSummary) UnsetPatchJobOverallProgress()`

UnsetPatchJobOverallProgress ensures that no value is present for PatchJobOverallProgress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


