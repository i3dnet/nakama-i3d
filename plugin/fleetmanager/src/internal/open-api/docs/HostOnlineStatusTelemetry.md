# HostOnlineStatusTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The unix time stamp | [readonly] 
**HostType** | **int32** | Type of the host e.g * 1: Bare metal * 2: Flex metal * 3: Cloud | [readonly] 
**Online** | **int32** | The number of hosts with status online | [readonly] 
**OnlinePercentage** | **int32** | The percentage of hosts with status online | [readonly] 
**Install** | **int32** | The number of hosts with status installing | [readonly] 
**InstallPercentage** | **int32** | The percentage of hosts with status installing | [readonly] 
**Offline** | **int32** | The number of hosts with status offline | [readonly] 
**OfflinePercentage** | **int32** | The percentage of hosts with status offline | [readonly] 
**Unknown** | **int32** | The number of hosts with status unknown | [readonly] 
**UnknownPercentage** | **int32** | The percentage of hosts with status unknown | [readonly] 
**Total** | **int32** | The total number of hosts | [readonly] 

## Methods

### NewHostOnlineStatusTelemetry

`func NewHostOnlineStatusTelemetry(timestamp int32, hostType int32, online int32, onlinePercentage int32, install int32, installPercentage int32, offline int32, offlinePercentage int32, unknown int32, unknownPercentage int32, total int32, ) *HostOnlineStatusTelemetry`

NewHostOnlineStatusTelemetry instantiates a new HostOnlineStatusTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOnlineStatusTelemetryWithDefaults

`func NewHostOnlineStatusTelemetryWithDefaults() *HostOnlineStatusTelemetry`

NewHostOnlineStatusTelemetryWithDefaults instantiates a new HostOnlineStatusTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostOnlineStatusTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostOnlineStatusTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostOnlineStatusTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetHostType

`func (o *HostOnlineStatusTelemetry) GetHostType() int32`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HostOnlineStatusTelemetry) GetHostTypeOk() (*int32, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HostOnlineStatusTelemetry) SetHostType(v int32)`

SetHostType sets HostType field to given value.


### GetOnline

`func (o *HostOnlineStatusTelemetry) GetOnline() int32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *HostOnlineStatusTelemetry) GetOnlineOk() (*int32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *HostOnlineStatusTelemetry) SetOnline(v int32)`

SetOnline sets Online field to given value.


### GetOnlinePercentage

`func (o *HostOnlineStatusTelemetry) GetOnlinePercentage() int32`

GetOnlinePercentage returns the OnlinePercentage field if non-nil, zero value otherwise.

### GetOnlinePercentageOk

`func (o *HostOnlineStatusTelemetry) GetOnlinePercentageOk() (*int32, bool)`

GetOnlinePercentageOk returns a tuple with the OnlinePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlinePercentage

`func (o *HostOnlineStatusTelemetry) SetOnlinePercentage(v int32)`

SetOnlinePercentage sets OnlinePercentage field to given value.


### GetInstall

`func (o *HostOnlineStatusTelemetry) GetInstall() int32`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HostOnlineStatusTelemetry) GetInstallOk() (*int32, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HostOnlineStatusTelemetry) SetInstall(v int32)`

SetInstall sets Install field to given value.


### GetInstallPercentage

`func (o *HostOnlineStatusTelemetry) GetInstallPercentage() int32`

GetInstallPercentage returns the InstallPercentage field if non-nil, zero value otherwise.

### GetInstallPercentageOk

`func (o *HostOnlineStatusTelemetry) GetInstallPercentageOk() (*int32, bool)`

GetInstallPercentageOk returns a tuple with the InstallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallPercentage

`func (o *HostOnlineStatusTelemetry) SetInstallPercentage(v int32)`

SetInstallPercentage sets InstallPercentage field to given value.


### GetOffline

`func (o *HostOnlineStatusTelemetry) GetOffline() int32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *HostOnlineStatusTelemetry) GetOfflineOk() (*int32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *HostOnlineStatusTelemetry) SetOffline(v int32)`

SetOffline sets Offline field to given value.


### GetOfflinePercentage

`func (o *HostOnlineStatusTelemetry) GetOfflinePercentage() int32`

GetOfflinePercentage returns the OfflinePercentage field if non-nil, zero value otherwise.

### GetOfflinePercentageOk

`func (o *HostOnlineStatusTelemetry) GetOfflinePercentageOk() (*int32, bool)`

GetOfflinePercentageOk returns a tuple with the OfflinePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflinePercentage

`func (o *HostOnlineStatusTelemetry) SetOfflinePercentage(v int32)`

SetOfflinePercentage sets OfflinePercentage field to given value.


### GetUnknown

`func (o *HostOnlineStatusTelemetry) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *HostOnlineStatusTelemetry) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *HostOnlineStatusTelemetry) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.


### GetUnknownPercentage

`func (o *HostOnlineStatusTelemetry) GetUnknownPercentage() int32`

GetUnknownPercentage returns the UnknownPercentage field if non-nil, zero value otherwise.

### GetUnknownPercentageOk

`func (o *HostOnlineStatusTelemetry) GetUnknownPercentageOk() (*int32, bool)`

GetUnknownPercentageOk returns a tuple with the UnknownPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownPercentage

`func (o *HostOnlineStatusTelemetry) SetUnknownPercentage(v int32)`

SetUnknownPercentage sets UnknownPercentage field to given value.


### GetTotal

`func (o *HostOnlineStatusTelemetry) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HostOnlineStatusTelemetry) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HostOnlineStatusTelemetry) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


