# TaskTemplateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idx** | **int32** | Task template action execution order | [readonly] 
**ActionType** | **string** | Task template action type | [readonly] 
**RequiredParams** | **[]string** | Parameters required by this argument | [readonly] 

## Methods

### NewTaskTemplateAction

`func NewTaskTemplateAction(idx int32, actionType string, requiredParams []string, ) *TaskTemplateAction`

NewTaskTemplateAction instantiates a new TaskTemplateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTemplateActionWithDefaults

`func NewTaskTemplateActionWithDefaults() *TaskTemplateAction`

NewTaskTemplateActionWithDefaults instantiates a new TaskTemplateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdx

`func (o *TaskTemplateAction) GetIdx() int32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *TaskTemplateAction) GetIdxOk() (*int32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *TaskTemplateAction) SetIdx(v int32)`

SetIdx sets Idx field to given value.


### GetActionType

`func (o *TaskTemplateAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *TaskTemplateAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *TaskTemplateAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetRequiredParams

`func (o *TaskTemplateAction) GetRequiredParams() []string`

GetRequiredParams returns the RequiredParams field if non-nil, zero value otherwise.

### GetRequiredParamsOk

`func (o *TaskTemplateAction) GetRequiredParamsOk() (*[]string, bool)`

GetRequiredParamsOk returns a tuple with the RequiredParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredParams

`func (o *TaskTemplateAction) SetRequiredParams(v []string)`

SetRequiredParams sets RequiredParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


