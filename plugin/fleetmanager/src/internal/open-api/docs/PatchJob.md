# PatchJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Patch job ID | [readonly] 
**DeploymentEnvironmentId** | **string** | Deployment environment ID | 
**ApplicationId** | **string** | Application ID | 
**PatchJobMethodId** | **int32** | Patch job type, you can find the list available patch job type [&#x60;GET /v3/patchJob/type&#x60;](game-publisher#/PatchJob/get_v3_patchJob_type) | 
**PatchJobMethod** | **string** | Name of the patch job method | [readonly] 
**PatchJobName** | **string** | Name of the patch job (unique) | 
**Status** | **int32** | Status of the patch job, you can find the list available patch job type [&#x60;GET /v3/patchJob/status&#x60;](game-publisher#/PatchJob/get_v3_patchJob_status) | [readonly] 
**StopMethodId** | **int32** | Stop method of an application instance, you can find the list available application instance stop methods [&#x60;GET /v3/application/stopMethod&#x60;](game-publisher#/Application/get_v3_application_stopMethod) | 
**StopMethodName** | **string** | Name of the stop method | [readonly] 
**StopMethodTimeout** | Pointer to **int32** | Stop method timeout of an application instance in seconds | [optional] 
**Comments** | Pointer to **string** | Comments for patch job | [optional] 
**SchedulerStartTime** | Pointer to **int32** | Start time when the patch job will be started. The default value is 0, this means start directly after creation of the patch job, else the patch job will start at the provided UTC timestamp | [optional] 
**PatchJobStartTime** | **int32** | Timestamp when the patch job has been started in UTC | [readonly] 
**PatchJobOverallProgress** | [**NullablePatchJobReportProgress**](PatchJobReportProgress.md) | How many application instances are being patched | [readonly] 
**ApplicationBuild** | [**[]PatchJobApplicationBuild**](PatchJobApplicationBuild.md) | Array of old and new application build needs to be patched | 
**Fleet** | [**[]PatchJobFleet**](PatchJobFleet.md) | Array of fleet that needs to be patched | 
**Email** | Pointer to [**[]PatchJobEmail**](PatchJobEmail.md) | Array of email address for reporting purposes | [optional] 
**CreatedAt** | **int32** | Timestamp when the patch job has been created | [readonly] 
**ChangedAt** | **int32** | Timestamp when the patch job last has been changed | [readonly] 
**FinishedAt** | **int32** | Timestamp when the patch job has been finished in UTC | [readonly] 
**IsRevertable** | **int32** | If the patch job can be cancelled, when is 1 you can cancel the patch job else not | [readonly] 
**RevertsPatchJobId** | **NullableInt32** | If filled in, this ID points to the patch job that has been reverted by this patch job | [readonly] 
**RevertedByPatchJobId** | **NullableInt32** | If filled in, this ID points to the PatchJob that has reverted this patch job | [readonly] 

## Methods

### NewPatchJob

`func NewPatchJob(id string, deploymentEnvironmentId string, applicationId string, patchJobMethodId int32, patchJobMethod string, patchJobName string, status int32, stopMethodId int32, stopMethodName string, patchJobStartTime int32, patchJobOverallProgress NullablePatchJobReportProgress, applicationBuild []PatchJobApplicationBuild, fleet []PatchJobFleet, createdAt int32, changedAt int32, finishedAt int32, isRevertable int32, revertsPatchJobId NullableInt32, revertedByPatchJobId NullableInt32, ) *PatchJob`

NewPatchJob instantiates a new PatchJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobWithDefaults

`func NewPatchJobWithDefaults() *PatchJob`

NewPatchJobWithDefaults instantiates a new PatchJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJob) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentEnvironmentId

`func (o *PatchJob) GetDeploymentEnvironmentId() string`

GetDeploymentEnvironmentId returns the DeploymentEnvironmentId field if non-nil, zero value otherwise.

### GetDeploymentEnvironmentIdOk

`func (o *PatchJob) GetDeploymentEnvironmentIdOk() (*string, bool)`

GetDeploymentEnvironmentIdOk returns a tuple with the DeploymentEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnvironmentId

`func (o *PatchJob) SetDeploymentEnvironmentId(v string)`

SetDeploymentEnvironmentId sets DeploymentEnvironmentId field to given value.


### GetApplicationId

`func (o *PatchJob) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PatchJob) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PatchJob) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetPatchJobMethodId

`func (o *PatchJob) GetPatchJobMethodId() int32`

GetPatchJobMethodId returns the PatchJobMethodId field if non-nil, zero value otherwise.

### GetPatchJobMethodIdOk

`func (o *PatchJob) GetPatchJobMethodIdOk() (*int32, bool)`

GetPatchJobMethodIdOk returns a tuple with the PatchJobMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobMethodId

`func (o *PatchJob) SetPatchJobMethodId(v int32)`

SetPatchJobMethodId sets PatchJobMethodId field to given value.


### GetPatchJobMethod

`func (o *PatchJob) GetPatchJobMethod() string`

GetPatchJobMethod returns the PatchJobMethod field if non-nil, zero value otherwise.

### GetPatchJobMethodOk

`func (o *PatchJob) GetPatchJobMethodOk() (*string, bool)`

GetPatchJobMethodOk returns a tuple with the PatchJobMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobMethod

`func (o *PatchJob) SetPatchJobMethod(v string)`

SetPatchJobMethod sets PatchJobMethod field to given value.


### GetPatchJobName

`func (o *PatchJob) GetPatchJobName() string`

GetPatchJobName returns the PatchJobName field if non-nil, zero value otherwise.

### GetPatchJobNameOk

`func (o *PatchJob) GetPatchJobNameOk() (*string, bool)`

GetPatchJobNameOk returns a tuple with the PatchJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobName

`func (o *PatchJob) SetPatchJobName(v string)`

SetPatchJobName sets PatchJobName field to given value.


### GetStatus

`func (o *PatchJob) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchJob) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchJob) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetStopMethodId

`func (o *PatchJob) GetStopMethodId() int32`

GetStopMethodId returns the StopMethodId field if non-nil, zero value otherwise.

### GetStopMethodIdOk

`func (o *PatchJob) GetStopMethodIdOk() (*int32, bool)`

GetStopMethodIdOk returns a tuple with the StopMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopMethodId

`func (o *PatchJob) SetStopMethodId(v int32)`

SetStopMethodId sets StopMethodId field to given value.


### GetStopMethodName

`func (o *PatchJob) GetStopMethodName() string`

GetStopMethodName returns the StopMethodName field if non-nil, zero value otherwise.

### GetStopMethodNameOk

`func (o *PatchJob) GetStopMethodNameOk() (*string, bool)`

GetStopMethodNameOk returns a tuple with the StopMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopMethodName

`func (o *PatchJob) SetStopMethodName(v string)`

SetStopMethodName sets StopMethodName field to given value.


### GetStopMethodTimeout

`func (o *PatchJob) GetStopMethodTimeout() int32`

GetStopMethodTimeout returns the StopMethodTimeout field if non-nil, zero value otherwise.

### GetStopMethodTimeoutOk

`func (o *PatchJob) GetStopMethodTimeoutOk() (*int32, bool)`

GetStopMethodTimeoutOk returns a tuple with the StopMethodTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopMethodTimeout

`func (o *PatchJob) SetStopMethodTimeout(v int32)`

SetStopMethodTimeout sets StopMethodTimeout field to given value.

### HasStopMethodTimeout

`func (o *PatchJob) HasStopMethodTimeout() bool`

HasStopMethodTimeout returns a boolean if a field has been set.

### GetComments

`func (o *PatchJob) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchJob) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchJob) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchJob) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSchedulerStartTime

`func (o *PatchJob) GetSchedulerStartTime() int32`

GetSchedulerStartTime returns the SchedulerStartTime field if non-nil, zero value otherwise.

### GetSchedulerStartTimeOk

`func (o *PatchJob) GetSchedulerStartTimeOk() (*int32, bool)`

GetSchedulerStartTimeOk returns a tuple with the SchedulerStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerStartTime

`func (o *PatchJob) SetSchedulerStartTime(v int32)`

SetSchedulerStartTime sets SchedulerStartTime field to given value.

### HasSchedulerStartTime

`func (o *PatchJob) HasSchedulerStartTime() bool`

HasSchedulerStartTime returns a boolean if a field has been set.

### GetPatchJobStartTime

`func (o *PatchJob) GetPatchJobStartTime() int32`

GetPatchJobStartTime returns the PatchJobStartTime field if non-nil, zero value otherwise.

### GetPatchJobStartTimeOk

`func (o *PatchJob) GetPatchJobStartTimeOk() (*int32, bool)`

GetPatchJobStartTimeOk returns a tuple with the PatchJobStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobStartTime

`func (o *PatchJob) SetPatchJobStartTime(v int32)`

SetPatchJobStartTime sets PatchJobStartTime field to given value.


### GetPatchJobOverallProgress

`func (o *PatchJob) GetPatchJobOverallProgress() PatchJobReportProgress`

GetPatchJobOverallProgress returns the PatchJobOverallProgress field if non-nil, zero value otherwise.

### GetPatchJobOverallProgressOk

`func (o *PatchJob) GetPatchJobOverallProgressOk() (*PatchJobReportProgress, bool)`

GetPatchJobOverallProgressOk returns a tuple with the PatchJobOverallProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchJobOverallProgress

`func (o *PatchJob) SetPatchJobOverallProgress(v PatchJobReportProgress)`

SetPatchJobOverallProgress sets PatchJobOverallProgress field to given value.


### SetPatchJobOverallProgressNil

`func (o *PatchJob) SetPatchJobOverallProgressNil(b bool)`

 SetPatchJobOverallProgressNil sets the value for PatchJobOverallProgress to be an explicit nil

### UnsetPatchJobOverallProgress
`func (o *PatchJob) UnsetPatchJobOverallProgress()`

UnsetPatchJobOverallProgress ensures that no value is present for PatchJobOverallProgress, not even an explicit nil
### GetApplicationBuild

`func (o *PatchJob) GetApplicationBuild() []PatchJobApplicationBuild`

GetApplicationBuild returns the ApplicationBuild field if non-nil, zero value otherwise.

### GetApplicationBuildOk

`func (o *PatchJob) GetApplicationBuildOk() (*[]PatchJobApplicationBuild, bool)`

GetApplicationBuildOk returns a tuple with the ApplicationBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuild

`func (o *PatchJob) SetApplicationBuild(v []PatchJobApplicationBuild)`

SetApplicationBuild sets ApplicationBuild field to given value.


### GetFleet

`func (o *PatchJob) GetFleet() []PatchJobFleet`

GetFleet returns the Fleet field if non-nil, zero value otherwise.

### GetFleetOk

`func (o *PatchJob) GetFleetOk() (*[]PatchJobFleet, bool)`

GetFleetOk returns a tuple with the Fleet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleet

`func (o *PatchJob) SetFleet(v []PatchJobFleet)`

SetFleet sets Fleet field to given value.


### GetEmail

`func (o *PatchJob) GetEmail() []PatchJobEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchJob) GetEmailOk() (*[]PatchJobEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchJob) SetEmail(v []PatchJobEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchJob) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchJob) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchJob) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchJob) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *PatchJob) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *PatchJob) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *PatchJob) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.


### GetFinishedAt

`func (o *PatchJob) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *PatchJob) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *PatchJob) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### GetIsRevertable

`func (o *PatchJob) GetIsRevertable() int32`

GetIsRevertable returns the IsRevertable field if non-nil, zero value otherwise.

### GetIsRevertableOk

`func (o *PatchJob) GetIsRevertableOk() (*int32, bool)`

GetIsRevertableOk returns a tuple with the IsRevertable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevertable

`func (o *PatchJob) SetIsRevertable(v int32)`

SetIsRevertable sets IsRevertable field to given value.


### GetRevertsPatchJobId

`func (o *PatchJob) GetRevertsPatchJobId() int32`

GetRevertsPatchJobId returns the RevertsPatchJobId field if non-nil, zero value otherwise.

### GetRevertsPatchJobIdOk

`func (o *PatchJob) GetRevertsPatchJobIdOk() (*int32, bool)`

GetRevertsPatchJobIdOk returns a tuple with the RevertsPatchJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertsPatchJobId

`func (o *PatchJob) SetRevertsPatchJobId(v int32)`

SetRevertsPatchJobId sets RevertsPatchJobId field to given value.


### SetRevertsPatchJobIdNil

`func (o *PatchJob) SetRevertsPatchJobIdNil(b bool)`

 SetRevertsPatchJobIdNil sets the value for RevertsPatchJobId to be an explicit nil

### UnsetRevertsPatchJobId
`func (o *PatchJob) UnsetRevertsPatchJobId()`

UnsetRevertsPatchJobId ensures that no value is present for RevertsPatchJobId, not even an explicit nil
### GetRevertedByPatchJobId

`func (o *PatchJob) GetRevertedByPatchJobId() int32`

GetRevertedByPatchJobId returns the RevertedByPatchJobId field if non-nil, zero value otherwise.

### GetRevertedByPatchJobIdOk

`func (o *PatchJob) GetRevertedByPatchJobIdOk() (*int32, bool)`

GetRevertedByPatchJobIdOk returns a tuple with the RevertedByPatchJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedByPatchJobId

`func (o *PatchJob) SetRevertedByPatchJobId(v int32)`

SetRevertedByPatchJobId sets RevertedByPatchJobId field to given value.


### SetRevertedByPatchJobIdNil

`func (o *PatchJob) SetRevertedByPatchJobIdNil(b bool)`

 SetRevertedByPatchJobIdNil sets the value for RevertedByPatchJobId to be an explicit nil

### UnsetRevertedByPatchJobId
`func (o *PatchJob) UnsetRevertedByPatchJobId()`

UnsetRevertedByPatchJobId ensures that no value is present for RevertedByPatchJobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


