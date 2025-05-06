# TaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Task ID | [readonly] 
**Name** | **string** | Task name | [readonly] 
**BatchId** | **int32** | The ID of the batch the task belongs to | [readonly] 
**CategoryId** | **int32** | The ID of the task category. Valid values are: - 1: game server | [readonly] 
**EntityId** | **int32** | The ID of the entity the task is being executed for. E.g. a game server ID (categoryId 1) | [readonly] 
**CurrentActionIdx** | **int32** | Currently active action index - relates to the action index as found in the template. | [readonly] 
**ExecuteAt** | **int32** | Unix timestamp indicating the start of execution. | [readonly] 
**LastActivityAt** | **int32** | Unix timestamp indicating lastest activity. | [readonly] 
**FinishedAt** | **int32** | Unix timestamp at which the task was finished. | [readonly] 
**ResultCode** | **int32** | The result of the task. 0: pending, 1: in progress, 2: paused, 125: cancelled, 126: done with error, 127: done | [readonly] 
**ResultText** | **string** | The result of the task in message form. | [readonly] 

## Methods

### NewTaskStatus

`func NewTaskStatus(id int32, name string, batchId int32, categoryId int32, entityId int32, currentActionIdx int32, executeAt int32, lastActivityAt int32, finishedAt int32, resultCode int32, resultText string, ) *TaskStatus`

NewTaskStatus instantiates a new TaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusWithDefaults

`func NewTaskStatusWithDefaults() *TaskStatus`

NewTaskStatusWithDefaults instantiates a new TaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskStatus) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TaskStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskStatus) SetName(v string)`

SetName sets Name field to given value.


### GetBatchId

`func (o *TaskStatus) GetBatchId() int32`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *TaskStatus) GetBatchIdOk() (*int32, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *TaskStatus) SetBatchId(v int32)`

SetBatchId sets BatchId field to given value.


### GetCategoryId

`func (o *TaskStatus) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TaskStatus) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TaskStatus) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetEntityId

`func (o *TaskStatus) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TaskStatus) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TaskStatus) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetCurrentActionIdx

`func (o *TaskStatus) GetCurrentActionIdx() int32`

GetCurrentActionIdx returns the CurrentActionIdx field if non-nil, zero value otherwise.

### GetCurrentActionIdxOk

`func (o *TaskStatus) GetCurrentActionIdxOk() (*int32, bool)`

GetCurrentActionIdxOk returns a tuple with the CurrentActionIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActionIdx

`func (o *TaskStatus) SetCurrentActionIdx(v int32)`

SetCurrentActionIdx sets CurrentActionIdx field to given value.


### GetExecuteAt

`func (o *TaskStatus) GetExecuteAt() int32`

GetExecuteAt returns the ExecuteAt field if non-nil, zero value otherwise.

### GetExecuteAtOk

`func (o *TaskStatus) GetExecuteAtOk() (*int32, bool)`

GetExecuteAtOk returns a tuple with the ExecuteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteAt

`func (o *TaskStatus) SetExecuteAt(v int32)`

SetExecuteAt sets ExecuteAt field to given value.


### GetLastActivityAt

`func (o *TaskStatus) GetLastActivityAt() int32`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *TaskStatus) GetLastActivityAtOk() (*int32, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *TaskStatus) SetLastActivityAt(v int32)`

SetLastActivityAt sets LastActivityAt field to given value.


### GetFinishedAt

`func (o *TaskStatus) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *TaskStatus) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *TaskStatus) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### GetResultCode

`func (o *TaskStatus) GetResultCode() int32`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TaskStatus) GetResultCodeOk() (*int32, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TaskStatus) SetResultCode(v int32)`

SetResultCode sets ResultCode field to given value.


### GetResultText

`func (o *TaskStatus) GetResultText() string`

GetResultText returns the ResultText field if non-nil, zero value otherwise.

### GetResultTextOk

`func (o *TaskStatus) GetResultTextOk() (*string, bool)`

GetResultTextOk returns a tuple with the ResultText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultText

`func (o *TaskStatus) SetResultText(v string)`

SetResultText sets ResultText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


