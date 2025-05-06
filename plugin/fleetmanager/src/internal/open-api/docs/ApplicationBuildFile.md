# ApplicationBuildFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application build file ID | [readonly] 
**BuildProvisioningFileId** | **string** | ID of the uploaded file to the i3D.net storage. Cannot be updated once set | 
**BuildProvisioningRegistrationId** | **string** | ID of the build provisioning registration. Cannot be updated once set | 
**FileName** | **string** | Name of the file | 
**FileSize** | **int32** | Size of the file | [readonly] 
**Version** | **string** | Version of the file | 
**Md5CheckSum** | **string** | MD5 checksum of the file | [readonly] 
**Path** | Pointer to **string** | location of the file | [optional] 
**Domain** | Pointer to **string** | The domain where the files are stored | [optional] 
**DockerRepository** | Pointer to **string** | The docker repository that needs to be used | [optional] 
**CreatedAt** | **int32** | UNIX timestamp when the application build file has been created | [readonly] 
**ChangedAt** | **int32** | UNIX timestamp when the application build file last has been changed | [readonly] 
**AvailableOnCDN** | **int32** | the file is on the CDN when is 1. | [readonly] 

## Methods

### NewApplicationBuildFile

`func NewApplicationBuildFile(id string, buildProvisioningFileId string, buildProvisioningRegistrationId string, fileName string, fileSize int32, version string, md5CheckSum string, createdAt int32, changedAt int32, availableOnCDN int32, ) *ApplicationBuildFile`

NewApplicationBuildFile instantiates a new ApplicationBuildFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBuildFileWithDefaults

`func NewApplicationBuildFileWithDefaults() *ApplicationBuildFile`

NewApplicationBuildFileWithDefaults instantiates a new ApplicationBuildFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationBuildFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationBuildFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationBuildFile) SetId(v string)`

SetId sets Id field to given value.


### GetBuildProvisioningFileId

`func (o *ApplicationBuildFile) GetBuildProvisioningFileId() string`

GetBuildProvisioningFileId returns the BuildProvisioningFileId field if non-nil, zero value otherwise.

### GetBuildProvisioningFileIdOk

`func (o *ApplicationBuildFile) GetBuildProvisioningFileIdOk() (*string, bool)`

GetBuildProvisioningFileIdOk returns a tuple with the BuildProvisioningFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildProvisioningFileId

`func (o *ApplicationBuildFile) SetBuildProvisioningFileId(v string)`

SetBuildProvisioningFileId sets BuildProvisioningFileId field to given value.


### GetBuildProvisioningRegistrationId

`func (o *ApplicationBuildFile) GetBuildProvisioningRegistrationId() string`

GetBuildProvisioningRegistrationId returns the BuildProvisioningRegistrationId field if non-nil, zero value otherwise.

### GetBuildProvisioningRegistrationIdOk

`func (o *ApplicationBuildFile) GetBuildProvisioningRegistrationIdOk() (*string, bool)`

GetBuildProvisioningRegistrationIdOk returns a tuple with the BuildProvisioningRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildProvisioningRegistrationId

`func (o *ApplicationBuildFile) SetBuildProvisioningRegistrationId(v string)`

SetBuildProvisioningRegistrationId sets BuildProvisioningRegistrationId field to given value.


### GetFileName

`func (o *ApplicationBuildFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ApplicationBuildFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ApplicationBuildFile) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *ApplicationBuildFile) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ApplicationBuildFile) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ApplicationBuildFile) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetVersion

`func (o *ApplicationBuildFile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationBuildFile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationBuildFile) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMd5CheckSum

`func (o *ApplicationBuildFile) GetMd5CheckSum() string`

GetMd5CheckSum returns the Md5CheckSum field if non-nil, zero value otherwise.

### GetMd5CheckSumOk

`func (o *ApplicationBuildFile) GetMd5CheckSumOk() (*string, bool)`

GetMd5CheckSumOk returns a tuple with the Md5CheckSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5CheckSum

`func (o *ApplicationBuildFile) SetMd5CheckSum(v string)`

SetMd5CheckSum sets Md5CheckSum field to given value.


### GetPath

`func (o *ApplicationBuildFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApplicationBuildFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApplicationBuildFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApplicationBuildFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDomain

`func (o *ApplicationBuildFile) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApplicationBuildFile) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApplicationBuildFile) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApplicationBuildFile) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDockerRepository

`func (o *ApplicationBuildFile) GetDockerRepository() string`

GetDockerRepository returns the DockerRepository field if non-nil, zero value otherwise.

### GetDockerRepositoryOk

`func (o *ApplicationBuildFile) GetDockerRepositoryOk() (*string, bool)`

GetDockerRepositoryOk returns a tuple with the DockerRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepository

`func (o *ApplicationBuildFile) SetDockerRepository(v string)`

SetDockerRepository sets DockerRepository field to given value.

### HasDockerRepository

`func (o *ApplicationBuildFile) HasDockerRepository() bool`

HasDockerRepository returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationBuildFile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationBuildFile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationBuildFile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *ApplicationBuildFile) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *ApplicationBuildFile) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *ApplicationBuildFile) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.


### GetAvailableOnCDN

`func (o *ApplicationBuildFile) GetAvailableOnCDN() int32`

GetAvailableOnCDN returns the AvailableOnCDN field if non-nil, zero value otherwise.

### GetAvailableOnCDNOk

`func (o *ApplicationBuildFile) GetAvailableOnCDNOk() (*int32, bool)`

GetAvailableOnCDNOk returns a tuple with the AvailableOnCDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOnCDN

`func (o *ApplicationBuildFile) SetAvailableOnCDN(v int32)`

SetAvailableOnCDN sets AvailableOnCDN field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


