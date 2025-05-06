# TaskBatchStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Task Batch ID | [readonly] 
**CreatedAt** | **int32** | The unix timestamp at which this batch was created | [readonly] 
**ExecuteAt** | **int32** | The unix timestamp at which this batch started execution | [readonly] 
**FinishedAt** | **int32** | The unix timestamp at which this batch was finished | [readonly] 
**NumTasks** | **int32** | The total number of tasks in this batch | [readonly] 
**NumTasksOk** | **int32** | The number of successfully executed tasks | [readonly] 
**NumTasksFailed** | **int32** | The number of failed tasks | [readonly] 
**NumTasksCancelled** | **int32** | The number of cancelled tasks | [readonly] 
**TaskIds** | **[]int32** | All the task IDs in this batch | [readonly] 

## Methods

### NewTaskBatchStatus

`func NewTaskBatchStatus(id int32, createdAt int32, executeAt int32, finishedAt int32, numTasks int32, numTasksOk int32, numTasksFailed int32, numTasksCancelled int32, taskIds []int32, ) *TaskBatchStatus`

NewTaskBatchStatus instantiates a new TaskBatchStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBatchStatusWithDefaults

`func NewTaskBatchStatusWithDefaults() *TaskBatchStatus`

NewTaskBatchStatusWithDefaults instantiates a new TaskBatchStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskBatchStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskBatchStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskBatchStatus) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TaskBatchStatus) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskBatchStatus) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskBatchStatus) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExecuteAt

`func (o *TaskBatchStatus) GetExecuteAt() int32`

GetExecuteAt returns the ExecuteAt field if non-nil, zero value otherwise.

### GetExecuteAtOk

`func (o *TaskBatchStatus) GetExecuteAtOk() (*int32, bool)`

GetExecuteAtOk returns a tuple with the ExecuteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteAt

`func (o *TaskBatchStatus) SetExecuteAt(v int32)`

SetExecuteAt sets ExecuteAt field to given value.


### GetFinishedAt

`func (o *TaskBatchStatus) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *TaskBatchStatus) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *TaskBatchStatus) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### GetNumTasks

`func (o *TaskBatchStatus) GetNumTasks() int32`

GetNumTasks returns the NumTasks field if non-nil, zero value otherwise.

### GetNumTasksOk

`func (o *TaskBatchStatus) GetNumTasksOk() (*int32, bool)`

GetNumTasksOk returns a tuple with the NumTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasks

`func (o *TaskBatchStatus) SetNumTasks(v int32)`

SetNumTasks sets NumTasks field to given value.


### GetNumTasksOk

`func (o *TaskBatchStatus) GetNumTasksOk() int32`

GetNumTasksOk returns the NumTasksOk field if non-nil, zero value otherwise.

### GetNumTasksOkOk

`func (o *TaskBatchStatus) GetNumTasksOkOk() (*int32, bool)`

GetNumTasksOkOk returns a tuple with the NumTasksOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasksOk

`func (o *TaskBatchStatus) SetNumTasksOk(v int32)`

SetNumTasksOk sets NumTasksOk field to given value.


### GetNumTasksFailed

`func (o *TaskBatchStatus) GetNumTasksFailed() int32`

GetNumTasksFailed returns the NumTasksFailed field if non-nil, zero value otherwise.

### GetNumTasksFailedOk

`func (o *TaskBatchStatus) GetNumTasksFailedOk() (*int32, bool)`

GetNumTasksFailedOk returns a tuple with the NumTasksFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasksFailed

`func (o *TaskBatchStatus) SetNumTasksFailed(v int32)`

SetNumTasksFailed sets NumTasksFailed field to given value.


### GetNumTasksCancelled

`func (o *TaskBatchStatus) GetNumTasksCancelled() int32`

GetNumTasksCancelled returns the NumTasksCancelled field if non-nil, zero value otherwise.

### GetNumTasksCancelledOk

`func (o *TaskBatchStatus) GetNumTasksCancelledOk() (*int32, bool)`

GetNumTasksCancelledOk returns a tuple with the NumTasksCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasksCancelled

`func (o *TaskBatchStatus) SetNumTasksCancelled(v int32)`

SetNumTasksCancelled sets NumTasksCancelled field to given value.


### GetTaskIds

`func (o *TaskBatchStatus) GetTaskIds() []int32`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *TaskBatchStatus) GetTaskIdsOk() (*[]int32, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *TaskBatchStatus) SetTaskIds(v []int32)`

SetTaskIds sets TaskIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


