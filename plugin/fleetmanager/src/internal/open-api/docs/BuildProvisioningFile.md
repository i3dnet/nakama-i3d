# BuildProvisioningFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Build provisioning file ID | [readonly] 
**BuildProvisionRegistrationId** | **string** | ID of the build provisioning storage registration | [readonly] 
**FileName** | **string** | Name of the file | [readonly] 
**FileSize** | **int32** | Size of the file | [readonly] 
**Md5CheckSum** | **string** | MD5 checksum of the file | [readonly] 
**CreatedAt** | **int32** | UNIX timestamp when the file has been created | [readonly] 
**ChangedAt** | **int32** | UNIX timestamp when the file last has been changed | [readonly] 

## Methods

### NewBuildProvisioningFile

`func NewBuildProvisioningFile(id string, buildProvisionRegistrationId string, fileName string, fileSize int32, md5CheckSum string, createdAt int32, changedAt int32, ) *BuildProvisioningFile`

NewBuildProvisioningFile instantiates a new BuildProvisioningFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildProvisioningFileWithDefaults

`func NewBuildProvisioningFileWithDefaults() *BuildProvisioningFile`

NewBuildProvisioningFileWithDefaults instantiates a new BuildProvisioningFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BuildProvisioningFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildProvisioningFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildProvisioningFile) SetId(v string)`

SetId sets Id field to given value.


### GetBuildProvisionRegistrationId

`func (o *BuildProvisioningFile) GetBuildProvisionRegistrationId() string`

GetBuildProvisionRegistrationId returns the BuildProvisionRegistrationId field if non-nil, zero value otherwise.

### GetBuildProvisionRegistrationIdOk

`func (o *BuildProvisioningFile) GetBuildProvisionRegistrationIdOk() (*string, bool)`

GetBuildProvisionRegistrationIdOk returns a tuple with the BuildProvisionRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildProvisionRegistrationId

`func (o *BuildProvisioningFile) SetBuildProvisionRegistrationId(v string)`

SetBuildProvisionRegistrationId sets BuildProvisionRegistrationId field to given value.


### GetFileName

`func (o *BuildProvisioningFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BuildProvisioningFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BuildProvisioningFile) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *BuildProvisioningFile) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *BuildProvisioningFile) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *BuildProvisioningFile) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetMd5CheckSum

`func (o *BuildProvisioningFile) GetMd5CheckSum() string`

GetMd5CheckSum returns the Md5CheckSum field if non-nil, zero value otherwise.

### GetMd5CheckSumOk

`func (o *BuildProvisioningFile) GetMd5CheckSumOk() (*string, bool)`

GetMd5CheckSumOk returns a tuple with the Md5CheckSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5CheckSum

`func (o *BuildProvisioningFile) SetMd5CheckSum(v string)`

SetMd5CheckSum sets Md5CheckSum field to given value.


### GetCreatedAt

`func (o *BuildProvisioningFile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BuildProvisioningFile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BuildProvisioningFile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *BuildProvisioningFile) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *BuildProvisioningFile) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *BuildProvisioningFile) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


