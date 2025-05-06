# TaskTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Task template ID | [readonly] 
**Name** | **string** | Task template name | [readonly] 
**Description** | **string** | Task template description | [readonly] 
**Actions** | [**[]TaskTemplateAction**](TaskTemplateAction.md) | Task template actions | [readonly] 

## Methods

### NewTaskTemplate

`func NewTaskTemplate(id int32, name string, description string, actions []TaskTemplateAction, ) *TaskTemplate`

NewTaskTemplate instantiates a new TaskTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTemplateWithDefaults

`func NewTaskTemplateWithDefaults() *TaskTemplate`

NewTaskTemplateWithDefaults instantiates a new TaskTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TaskTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TaskTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetActions

`func (o *TaskTemplate) GetActions() []TaskTemplateAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TaskTemplate) GetActionsOk() (*[]TaskTemplateAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TaskTemplate) SetActions(v []TaskTemplateAction)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


