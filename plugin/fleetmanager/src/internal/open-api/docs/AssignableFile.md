# AssignableFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the file | [readonly] 
**FileName** | **string** | Name of the file | [readonly] 
**FileSize** | **int32** | Size in bytes of the file | [readonly] 
**FileSHA256** | **string** | Archive file SHA-256 hash | [readonly] 
**Status** | **int32** | Status of this file, options: - 0: no status. - 1: new file discovered on FTP account. - 2: new file on FTP account has been scanned and added to collection. - 3: application install details have been provided. - 4: transferring the original archive to the main file server. - 5: finished transferring the original archive to the main file server. - 6: re-packing the original archive. - 7: finished re-packing the original archive. - 10: creating a hash of the new archive. - 11: deploying the final archive to all relevant file servers. - 12: deploying the final archive to all relevant file servers - 13: finished deploying the final archive to all relevant file servers. - 14: activated the game install. - 125: cancelled. - 126: cancelled and cleaned up. - 127: the application install creation process has finished. - 128: a permanent error has been encountered. | [readonly] 
**StatusText** | **string** | Status in text of this file | [readonly] 

## Methods

### NewAssignableFile

`func NewAssignableFile(id int32, fileName string, fileSize int32, fileSHA256 string, status int32, statusText string, ) *AssignableFile`

NewAssignableFile instantiates a new AssignableFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignableFileWithDefaults

`func NewAssignableFileWithDefaults() *AssignableFile`

NewAssignableFileWithDefaults instantiates a new AssignableFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssignableFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignableFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignableFile) SetId(v int32)`

SetId sets Id field to given value.


### GetFileName

`func (o *AssignableFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AssignableFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AssignableFile) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *AssignableFile) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AssignableFile) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AssignableFile) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFileSHA256

`func (o *AssignableFile) GetFileSHA256() string`

GetFileSHA256 returns the FileSHA256 field if non-nil, zero value otherwise.

### GetFileSHA256Ok

`func (o *AssignableFile) GetFileSHA256Ok() (*string, bool)`

GetFileSHA256Ok returns a tuple with the FileSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSHA256

`func (o *AssignableFile) SetFileSHA256(v string)`

SetFileSHA256 sets FileSHA256 field to given value.


### GetStatus

`func (o *AssignableFile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignableFile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignableFile) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetStatusText

`func (o *AssignableFile) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *AssignableFile) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *AssignableFile) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


