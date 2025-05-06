# HostOsInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Os** | [**OperatingSystemPartial**](OperatingSystemPartial.md) |  | 
**PostInstallScript** | Pointer to **string** |  | [optional] 
**QuickFormat** | Pointer to **bool** | Perform a quick format of all drives. Default: true | [optional] 
**SshKeyUuid** | Pointer to **string** | For an overview of available SSH keys, see /v3/sshKey. Only required for Linux OSs, except Talos. | [optional] 

## Methods

### NewHostOsInstall

`func NewHostOsInstall(os OperatingSystemPartial, ) *HostOsInstall`

NewHostOsInstall instantiates a new HostOsInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOsInstallWithDefaults

`func NewHostOsInstallWithDefaults() *HostOsInstall`

NewHostOsInstallWithDefaults instantiates a new HostOsInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *HostOsInstall) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostOsInstall) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostOsInstall) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HostOsInstall) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOs

`func (o *HostOsInstall) GetOs() OperatingSystemPartial`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *HostOsInstall) GetOsOk() (*OperatingSystemPartial, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *HostOsInstall) SetOs(v OperatingSystemPartial)`

SetOs sets Os field to given value.


### GetPostInstallScript

`func (o *HostOsInstall) GetPostInstallScript() string`

GetPostInstallScript returns the PostInstallScript field if non-nil, zero value otherwise.

### GetPostInstallScriptOk

`func (o *HostOsInstall) GetPostInstallScriptOk() (*string, bool)`

GetPostInstallScriptOk returns a tuple with the PostInstallScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostInstallScript

`func (o *HostOsInstall) SetPostInstallScript(v string)`

SetPostInstallScript sets PostInstallScript field to given value.

### HasPostInstallScript

`func (o *HostOsInstall) HasPostInstallScript() bool`

HasPostInstallScript returns a boolean if a field has been set.

### GetQuickFormat

`func (o *HostOsInstall) GetQuickFormat() bool`

GetQuickFormat returns the QuickFormat field if non-nil, zero value otherwise.

### GetQuickFormatOk

`func (o *HostOsInstall) GetQuickFormatOk() (*bool, bool)`

GetQuickFormatOk returns a tuple with the QuickFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickFormat

`func (o *HostOsInstall) SetQuickFormat(v bool)`

SetQuickFormat sets QuickFormat field to given value.

### HasQuickFormat

`func (o *HostOsInstall) HasQuickFormat() bool`

HasQuickFormat returns a boolean if a field has been set.

### GetSshKeyUuid

`func (o *HostOsInstall) GetSshKeyUuid() string`

GetSshKeyUuid returns the SshKeyUuid field if non-nil, zero value otherwise.

### GetSshKeyUuidOk

`func (o *HostOsInstall) GetSshKeyUuidOk() (*string, bool)`

GetSshKeyUuidOk returns a tuple with the SshKeyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyUuid

`func (o *HostOsInstall) SetSshKeyUuid(v string)`

SetSshKeyUuid sets SshKeyUuid field to given value.

### HasSshKeyUuid

`func (o *HostOsInstall) HasSshKeyUuid() bool`

HasSshKeyUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


