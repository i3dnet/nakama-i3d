# PatchJobEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a patch job email | [readonly] 
**Email** | **string** | Email address for reporting (unique) | 
**ProgressReport** | **int32** | If 1 progress report will be send, if 0 no progress report will be send | 
**ResultReport** | **int32** | If 1 result report will be send, if 0 no result report will be send | 
**CreatedAt** | **int32** | Timestamp when the patch job email has been created | [readonly] 
**ChangedAt** | **int32** | Timestamp when the patch job email last has been changed | [readonly] 

## Methods

### NewPatchJobEmail

`func NewPatchJobEmail(id string, email string, progressReport int32, resultReport int32, createdAt int32, changedAt int32, ) *PatchJobEmail`

NewPatchJobEmail instantiates a new PatchJobEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobEmailWithDefaults

`func NewPatchJobEmailWithDefaults() *PatchJobEmail`

NewPatchJobEmailWithDefaults instantiates a new PatchJobEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJobEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJobEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJobEmail) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *PatchJobEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchJobEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchJobEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProgressReport

`func (o *PatchJobEmail) GetProgressReport() int32`

GetProgressReport returns the ProgressReport field if non-nil, zero value otherwise.

### GetProgressReportOk

`func (o *PatchJobEmail) GetProgressReportOk() (*int32, bool)`

GetProgressReportOk returns a tuple with the ProgressReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressReport

`func (o *PatchJobEmail) SetProgressReport(v int32)`

SetProgressReport sets ProgressReport field to given value.


### GetResultReport

`func (o *PatchJobEmail) GetResultReport() int32`

GetResultReport returns the ResultReport field if non-nil, zero value otherwise.

### GetResultReportOk

`func (o *PatchJobEmail) GetResultReportOk() (*int32, bool)`

GetResultReportOk returns a tuple with the ResultReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReport

`func (o *PatchJobEmail) SetResultReport(v int32)`

SetResultReport sets ResultReport field to given value.


### GetCreatedAt

`func (o *PatchJobEmail) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchJobEmail) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchJobEmail) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *PatchJobEmail) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *PatchJobEmail) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *PatchJobEmail) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


