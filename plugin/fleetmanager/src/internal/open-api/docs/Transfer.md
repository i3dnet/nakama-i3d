# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileServerId** | **int32** | ID of the file server | 
**Progress** | **int32** | The transfer progress in percentage | 
**ErrorCount** | **int32** | The amount of errors that occurred during this process | 
**Retries** | **int32** | The amount of retries in the process | 

## Methods

### NewTransfer

`func NewTransfer(fileServerId int32, progress int32, errorCount int32, retries int32, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileServerId

`func (o *Transfer) GetFileServerId() int32`

GetFileServerId returns the FileServerId field if non-nil, zero value otherwise.

### GetFileServerIdOk

`func (o *Transfer) GetFileServerIdOk() (*int32, bool)`

GetFileServerIdOk returns a tuple with the FileServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServerId

`func (o *Transfer) SetFileServerId(v int32)`

SetFileServerId sets FileServerId field to given value.


### GetProgress

`func (o *Transfer) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Transfer) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Transfer) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetErrorCount

`func (o *Transfer) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *Transfer) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *Transfer) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetRetries

`func (o *Transfer) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Transfer) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Transfer) SetRetries(v int32)`

SetRetries sets Retries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


