# ApplicationInstanceTaskCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskTemplateId** | **string** | Task template ID, as seen in GET /applicationInstance/task/template | 
**ApplicationInstanceIds** | Pointer to **[]string** | Array of application instance IDs to create the tasks for | [optional] 
**FleetId** | Pointer to **string** | If &gt; 0, we gather all game server IDs for this fleet and add those to the gameServerIds property | [optional] 
**RegionId** | Pointer to **string** | If &gt; 0, we gather all game server IDs for this region and add those to the gameServerIds property | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | Optionally select game servers by label key/value pairs. | [optional] 
**ExecuteAt** | Pointer to **int32** | A unix timestamp indicating the time at which to execute the task (to schedule tasks for future execution) | [optional] 
**TaskActionParam** | Pointer to [**[]TaskTemplateActionParameters**](TaskTemplateActionParameters.md) | The list of required action parameters for a template. See the task templates (/applicationInstance/task/template) end point for more information. | [optional] 

## Methods

### NewApplicationInstanceTaskCreate

`func NewApplicationInstanceTaskCreate(taskTemplateId string, ) *ApplicationInstanceTaskCreate`

NewApplicationInstanceTaskCreate instantiates a new ApplicationInstanceTaskCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceTaskCreateWithDefaults

`func NewApplicationInstanceTaskCreateWithDefaults() *ApplicationInstanceTaskCreate`

NewApplicationInstanceTaskCreateWithDefaults instantiates a new ApplicationInstanceTaskCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskTemplateId

`func (o *ApplicationInstanceTaskCreate) GetTaskTemplateId() string`

GetTaskTemplateId returns the TaskTemplateId field if non-nil, zero value otherwise.

### GetTaskTemplateIdOk

`func (o *ApplicationInstanceTaskCreate) GetTaskTemplateIdOk() (*string, bool)`

GetTaskTemplateIdOk returns a tuple with the TaskTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTemplateId

`func (o *ApplicationInstanceTaskCreate) SetTaskTemplateId(v string)`

SetTaskTemplateId sets TaskTemplateId field to given value.


### GetApplicationInstanceIds

`func (o *ApplicationInstanceTaskCreate) GetApplicationInstanceIds() []string`

GetApplicationInstanceIds returns the ApplicationInstanceIds field if non-nil, zero value otherwise.

### GetApplicationInstanceIdsOk

`func (o *ApplicationInstanceTaskCreate) GetApplicationInstanceIdsOk() (*[]string, bool)`

GetApplicationInstanceIdsOk returns a tuple with the ApplicationInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceIds

`func (o *ApplicationInstanceTaskCreate) SetApplicationInstanceIds(v []string)`

SetApplicationInstanceIds sets ApplicationInstanceIds field to given value.

### HasApplicationInstanceIds

`func (o *ApplicationInstanceTaskCreate) HasApplicationInstanceIds() bool`

HasApplicationInstanceIds returns a boolean if a field has been set.

### GetFleetId

`func (o *ApplicationInstanceTaskCreate) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ApplicationInstanceTaskCreate) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ApplicationInstanceTaskCreate) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.

### HasFleetId

`func (o *ApplicationInstanceTaskCreate) HasFleetId() bool`

HasFleetId returns a boolean if a field has been set.

### GetRegionId

`func (o *ApplicationInstanceTaskCreate) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ApplicationInstanceTaskCreate) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ApplicationInstanceTaskCreate) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ApplicationInstanceTaskCreate) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetLabels

`func (o *ApplicationInstanceTaskCreate) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApplicationInstanceTaskCreate) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApplicationInstanceTaskCreate) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApplicationInstanceTaskCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetExecuteAt

`func (o *ApplicationInstanceTaskCreate) GetExecuteAt() int32`

GetExecuteAt returns the ExecuteAt field if non-nil, zero value otherwise.

### GetExecuteAtOk

`func (o *ApplicationInstanceTaskCreate) GetExecuteAtOk() (*int32, bool)`

GetExecuteAtOk returns a tuple with the ExecuteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteAt

`func (o *ApplicationInstanceTaskCreate) SetExecuteAt(v int32)`

SetExecuteAt sets ExecuteAt field to given value.

### HasExecuteAt

`func (o *ApplicationInstanceTaskCreate) HasExecuteAt() bool`

HasExecuteAt returns a boolean if a field has been set.

### GetTaskActionParam

`func (o *ApplicationInstanceTaskCreate) GetTaskActionParam() []TaskTemplateActionParameters`

GetTaskActionParam returns the TaskActionParam field if non-nil, zero value otherwise.

### GetTaskActionParamOk

`func (o *ApplicationInstanceTaskCreate) GetTaskActionParamOk() (*[]TaskTemplateActionParameters, bool)`

GetTaskActionParamOk returns a tuple with the TaskActionParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskActionParam

`func (o *ApplicationInstanceTaskCreate) SetTaskActionParam(v []TaskTemplateActionParameters)`

SetTaskActionParam sets TaskActionParam field to given value.

### HasTaskActionParam

`func (o *ApplicationInstanceTaskCreate) HasTaskActionParam() bool`

HasTaskActionParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


