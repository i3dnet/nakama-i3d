# CreateApplicationInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Install ID | [readonly] 
**FileId** | **int32** | ID of the file | 
**FileName** | **string** | Name of the file | 
**FileSize** | **int32** | Size in bytes of the file | 
**FileSHA256** | **string** | Archive file SHA-256 hash | 
**Status** | **int32** | Status of this file&lt;br /&gt; **1**: File has been seen on your FTP account&lt;br /&gt; **2**: Your file has stopped - regarding the upload as finished&lt;br /&gt; **3**: File details have been stored&lt;br /&gt; **5**: File has been moved to our main file server for further processing&lt;br /&gt; **6**: Re-packing your original archive&lt;br /&gt; **7**: Re-packing is done&lt;br /&gt; **10**: Creating an SHA-256 hash of the archive&lt;br /&gt; **11**: Hashing done&lt;br /&gt; **12**: Mirroring the archive to global file servers&lt;br /&gt; **13**: Mirroring is done&lt;br /&gt; **14**: ApplicationInstall has been activated&lt;br /&gt; **125**: Cancelled 1&lt;br /&gt; **126**: Cancelled 2&lt;br /&gt; **127**: ApplicationInstall creation is done&lt;br /&gt; **128**: A permanent error has occurred | 
**StatusText** | **string** | Status in text of this file | 
**MirrorState** | [**[]MirrorState**](MirrorState.md) | An array with the status of the mirroring to all locations process | 

## Methods

### NewCreateApplicationInstall

`func NewCreateApplicationInstall(id int32, fileId int32, fileName string, fileSize int32, fileSHA256 string, status int32, statusText string, mirrorState []MirrorState, ) *CreateApplicationInstall`

NewCreateApplicationInstall instantiates a new CreateApplicationInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationInstallWithDefaults

`func NewCreateApplicationInstallWithDefaults() *CreateApplicationInstall`

NewCreateApplicationInstallWithDefaults instantiates a new CreateApplicationInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateApplicationInstall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateApplicationInstall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateApplicationInstall) SetId(v int32)`

SetId sets Id field to given value.


### GetFileId

`func (o *CreateApplicationInstall) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *CreateApplicationInstall) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *CreateApplicationInstall) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetFileName

`func (o *CreateApplicationInstall) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateApplicationInstall) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateApplicationInstall) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *CreateApplicationInstall) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *CreateApplicationInstall) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *CreateApplicationInstall) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFileSHA256

`func (o *CreateApplicationInstall) GetFileSHA256() string`

GetFileSHA256 returns the FileSHA256 field if non-nil, zero value otherwise.

### GetFileSHA256Ok

`func (o *CreateApplicationInstall) GetFileSHA256Ok() (*string, bool)`

GetFileSHA256Ok returns a tuple with the FileSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSHA256

`func (o *CreateApplicationInstall) SetFileSHA256(v string)`

SetFileSHA256 sets FileSHA256 field to given value.


### GetStatus

`func (o *CreateApplicationInstall) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateApplicationInstall) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateApplicationInstall) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetStatusText

`func (o *CreateApplicationInstall) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *CreateApplicationInstall) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *CreateApplicationInstall) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.


### GetMirrorState

`func (o *CreateApplicationInstall) GetMirrorState() []MirrorState`

GetMirrorState returns the MirrorState field if non-nil, zero value otherwise.

### GetMirrorStateOk

`func (o *CreateApplicationInstall) GetMirrorStateOk() (*[]MirrorState, bool)`

GetMirrorStateOk returns a tuple with the MirrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorState

`func (o *CreateApplicationInstall) SetMirrorState(v []MirrorState)`

SetMirrorState sets MirrorState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


