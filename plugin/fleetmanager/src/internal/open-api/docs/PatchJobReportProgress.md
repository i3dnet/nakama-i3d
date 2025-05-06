# PatchJobReportProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total amount of application instances affected by the patch job | [readonly] 
**Downloaded** | **int32** | The amount of downloaded application instances | [readonly] 
**Stopped** | **int32** | The amount of stopped application instances | [readonly] 
**Success** | **int32** | The amount of successfully patched application instances | [readonly] 
**Failed** | **int32** | The amount of failed application instances | [readonly] 
**Started** | **int32** | The amount of started application instances | [readonly] 
**Canceled** | **int32** | The amount of canceled application instances | [readonly] 
**Remaining** | **int32** | The amount of remaining application instances | [readonly] 

## Methods

### NewPatchJobReportProgress

`func NewPatchJobReportProgress(total int32, downloaded int32, stopped int32, success int32, failed int32, started int32, canceled int32, remaining int32, ) *PatchJobReportProgress`

NewPatchJobReportProgress instantiates a new PatchJobReportProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJobReportProgressWithDefaults

`func NewPatchJobReportProgressWithDefaults() *PatchJobReportProgress`

NewPatchJobReportProgressWithDefaults instantiates a new PatchJobReportProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PatchJobReportProgress) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PatchJobReportProgress) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PatchJobReportProgress) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetDownloaded

`func (o *PatchJobReportProgress) GetDownloaded() int32`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *PatchJobReportProgress) GetDownloadedOk() (*int32, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *PatchJobReportProgress) SetDownloaded(v int32)`

SetDownloaded sets Downloaded field to given value.


### GetStopped

`func (o *PatchJobReportProgress) GetStopped() int32`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *PatchJobReportProgress) GetStoppedOk() (*int32, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *PatchJobReportProgress) SetStopped(v int32)`

SetStopped sets Stopped field to given value.


### GetSuccess

`func (o *PatchJobReportProgress) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PatchJobReportProgress) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PatchJobReportProgress) SetSuccess(v int32)`

SetSuccess sets Success field to given value.


### GetFailed

`func (o *PatchJobReportProgress) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *PatchJobReportProgress) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *PatchJobReportProgress) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetStarted

`func (o *PatchJobReportProgress) GetStarted() int32`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *PatchJobReportProgress) GetStartedOk() (*int32, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *PatchJobReportProgress) SetStarted(v int32)`

SetStarted sets Started field to given value.


### GetCanceled

`func (o *PatchJobReportProgress) GetCanceled() int32`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *PatchJobReportProgress) GetCanceledOk() (*int32, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *PatchJobReportProgress) SetCanceled(v int32)`

SetCanceled sets Canceled field to given value.


### GetRemaining

`func (o *PatchJobReportProgress) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *PatchJobReportProgress) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *PatchJobReportProgress) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


