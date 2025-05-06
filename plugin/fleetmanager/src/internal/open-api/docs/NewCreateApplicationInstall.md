# NewCreateApplicationInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The application&#39;s ID this application install is coupled to | 
**FileId** | **int32** | ID of the assignableFile. Can be obtained from [&#x60;GET /applicationInstall/assignable&#x60;](#/ApplicationAssignableFile/get_applicationInstall_assignable) | 
**OsGroupId** | **int32** | Operating system group ID (1 &#x3D; windows, 2 &#x3D; linux) | 
**Name** | **string** | Application install name | 
**Version** | **string** | Application install version | 
**ReleaseNotes** | Pointer to **NullableString** | Game install release notes | [optional] 
**InstallHasExe** | **int32** | Does the archive contain the game instance executable? | 

## Methods

### NewNewCreateApplicationInstall

`func NewNewCreateApplicationInstall(applicationId string, fileId int32, osGroupId int32, name string, version string, installHasExe int32, ) *NewCreateApplicationInstall`

NewNewCreateApplicationInstall instantiates a new NewCreateApplicationInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCreateApplicationInstallWithDefaults

`func NewNewCreateApplicationInstallWithDefaults() *NewCreateApplicationInstall`

NewNewCreateApplicationInstallWithDefaults instantiates a new NewCreateApplicationInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *NewCreateApplicationInstall) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NewCreateApplicationInstall) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NewCreateApplicationInstall) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetFileId

`func (o *NewCreateApplicationInstall) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *NewCreateApplicationInstall) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *NewCreateApplicationInstall) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetOsGroupId

`func (o *NewCreateApplicationInstall) GetOsGroupId() int32`

GetOsGroupId returns the OsGroupId field if non-nil, zero value otherwise.

### GetOsGroupIdOk

`func (o *NewCreateApplicationInstall) GetOsGroupIdOk() (*int32, bool)`

GetOsGroupIdOk returns a tuple with the OsGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGroupId

`func (o *NewCreateApplicationInstall) SetOsGroupId(v int32)`

SetOsGroupId sets OsGroupId field to given value.


### GetName

`func (o *NewCreateApplicationInstall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCreateApplicationInstall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCreateApplicationInstall) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *NewCreateApplicationInstall) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewCreateApplicationInstall) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NewCreateApplicationInstall) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetReleaseNotes

`func (o *NewCreateApplicationInstall) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *NewCreateApplicationInstall) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *NewCreateApplicationInstall) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *NewCreateApplicationInstall) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### SetReleaseNotesNil

`func (o *NewCreateApplicationInstall) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *NewCreateApplicationInstall) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil
### GetInstallHasExe

`func (o *NewCreateApplicationInstall) GetInstallHasExe() int32`

GetInstallHasExe returns the InstallHasExe field if non-nil, zero value otherwise.

### GetInstallHasExeOk

`func (o *NewCreateApplicationInstall) GetInstallHasExeOk() (*int32, bool)`

GetInstallHasExeOk returns a tuple with the InstallHasExe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallHasExe

`func (o *NewCreateApplicationInstall) SetInstallHasExe(v int32)`

SetInstallHasExe sets InstallHasExe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


