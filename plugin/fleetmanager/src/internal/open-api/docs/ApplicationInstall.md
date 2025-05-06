# ApplicationInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Application install ID | [readonly] 
**ApplicationId** | **string** | Application install assigned to applicationId | 
**ApplicationType** | **NullableInt32** | The list of Application types can be found in: GET /application/type | 
**CreatedAt** | **int32** | Application install created at timestamp | [readonly] 
**Name** | **string** | Application install name | 
**Version** | **string** | Application install version | 
**OsGroupId** | **int32** | Operating system group ID (1 &#x3D; windows, 2 &#x3D; linux) | 
**Active** | **int32** | Application install active | 
**ReleaseNotes** | Pointer to **NullableString** | Application install release notes | [optional] 
**FilePath** | **string** | Path of the file on the file server | [readonly] 
**FileSize** | **int32** | File size in bytes | [readonly] 
**FileSHA256** | **string** | Archive file SHA-256 hash | [readonly] 
**ExecutableSHA256** | **string** | SHA-256 hash of the executable in the archive | [readonly] 
**Status** | **int32** | Status of application install. | [readonly] 
**InstallHasExe** | **int32** | Does the archive contain the game instance executable? | 

## Methods

### NewApplicationInstall

`func NewApplicationInstall(id int32, applicationId string, applicationType NullableInt32, createdAt int32, name string, version string, osGroupId int32, active int32, filePath string, fileSize int32, fileSHA256 string, executableSHA256 string, status int32, installHasExe int32, ) *ApplicationInstall`

NewApplicationInstall instantiates a new ApplicationInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstallWithDefaults

`func NewApplicationInstallWithDefaults() *ApplicationInstall`

NewApplicationInstallWithDefaults instantiates a new ApplicationInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstall) SetId(v int32)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *ApplicationInstall) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationInstall) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationInstall) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationType

`func (o *ApplicationInstall) GetApplicationType() int32`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationInstall) GetApplicationTypeOk() (*int32, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationInstall) SetApplicationType(v int32)`

SetApplicationType sets ApplicationType field to given value.


### SetApplicationTypeNil

`func (o *ApplicationInstall) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *ApplicationInstall) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetCreatedAt

`func (o *ApplicationInstall) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationInstall) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationInstall) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *ApplicationInstall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationInstall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationInstall) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ApplicationInstall) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationInstall) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationInstall) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOsGroupId

`func (o *ApplicationInstall) GetOsGroupId() int32`

GetOsGroupId returns the OsGroupId field if non-nil, zero value otherwise.

### GetOsGroupIdOk

`func (o *ApplicationInstall) GetOsGroupIdOk() (*int32, bool)`

GetOsGroupIdOk returns a tuple with the OsGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGroupId

`func (o *ApplicationInstall) SetOsGroupId(v int32)`

SetOsGroupId sets OsGroupId field to given value.


### GetActive

`func (o *ApplicationInstall) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationInstall) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationInstall) SetActive(v int32)`

SetActive sets Active field to given value.


### GetReleaseNotes

`func (o *ApplicationInstall) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ApplicationInstall) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ApplicationInstall) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *ApplicationInstall) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### SetReleaseNotesNil

`func (o *ApplicationInstall) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *ApplicationInstall) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil
### GetFilePath

`func (o *ApplicationInstall) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ApplicationInstall) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ApplicationInstall) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetFileSize

`func (o *ApplicationInstall) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ApplicationInstall) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ApplicationInstall) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFileSHA256

`func (o *ApplicationInstall) GetFileSHA256() string`

GetFileSHA256 returns the FileSHA256 field if non-nil, zero value otherwise.

### GetFileSHA256Ok

`func (o *ApplicationInstall) GetFileSHA256Ok() (*string, bool)`

GetFileSHA256Ok returns a tuple with the FileSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSHA256

`func (o *ApplicationInstall) SetFileSHA256(v string)`

SetFileSHA256 sets FileSHA256 field to given value.


### GetExecutableSHA256

`func (o *ApplicationInstall) GetExecutableSHA256() string`

GetExecutableSHA256 returns the ExecutableSHA256 field if non-nil, zero value otherwise.

### GetExecutableSHA256Ok

`func (o *ApplicationInstall) GetExecutableSHA256Ok() (*string, bool)`

GetExecutableSHA256Ok returns a tuple with the ExecutableSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutableSHA256

`func (o *ApplicationInstall) SetExecutableSHA256(v string)`

SetExecutableSHA256 sets ExecutableSHA256 field to given value.


### GetStatus

`func (o *ApplicationInstall) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationInstall) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationInstall) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetInstallHasExe

`func (o *ApplicationInstall) GetInstallHasExe() int32`

GetInstallHasExe returns the InstallHasExe field if non-nil, zero value otherwise.

### GetInstallHasExeOk

`func (o *ApplicationInstall) GetInstallHasExeOk() (*int32, bool)`

GetInstallHasExeOk returns a tuple with the InstallHasExe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallHasExe

`func (o *ApplicationInstall) SetInstallHasExe(v int32)`

SetInstallHasExe sets InstallHasExe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


