# HostOnlineStatusHistoryDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostType** | **int32** | Type of the host e.g 1: Bare metal, 2: Flex metal, 3: Cloud | [readonly] 
**Online** | **int32** | The number of host with status online | [readonly] 
**Install** | **int32** | The number of host with status installing | [readonly] 
**Offline** | **int32** | The number of host with status offline | [readonly] 
**Unknown** | **int32** | The number of host with status unknow | [readonly] 

## Methods

### NewHostOnlineStatusHistoryDataTelemetry

`func NewHostOnlineStatusHistoryDataTelemetry(hostType int32, online int32, install int32, offline int32, unknown int32, ) *HostOnlineStatusHistoryDataTelemetry`

NewHostOnlineStatusHistoryDataTelemetry instantiates a new HostOnlineStatusHistoryDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOnlineStatusHistoryDataTelemetryWithDefaults

`func NewHostOnlineStatusHistoryDataTelemetryWithDefaults() *HostOnlineStatusHistoryDataTelemetry`

NewHostOnlineStatusHistoryDataTelemetryWithDefaults instantiates a new HostOnlineStatusHistoryDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostType

`func (o *HostOnlineStatusHistoryDataTelemetry) GetHostType() int32`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HostOnlineStatusHistoryDataTelemetry) GetHostTypeOk() (*int32, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HostOnlineStatusHistoryDataTelemetry) SetHostType(v int32)`

SetHostType sets HostType field to given value.


### GetOnline

`func (o *HostOnlineStatusHistoryDataTelemetry) GetOnline() int32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *HostOnlineStatusHistoryDataTelemetry) GetOnlineOk() (*int32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *HostOnlineStatusHistoryDataTelemetry) SetOnline(v int32)`

SetOnline sets Online field to given value.


### GetInstall

`func (o *HostOnlineStatusHistoryDataTelemetry) GetInstall() int32`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HostOnlineStatusHistoryDataTelemetry) GetInstallOk() (*int32, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HostOnlineStatusHistoryDataTelemetry) SetInstall(v int32)`

SetInstall sets Install field to given value.


### GetOffline

`func (o *HostOnlineStatusHistoryDataTelemetry) GetOffline() int32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *HostOnlineStatusHistoryDataTelemetry) GetOfflineOk() (*int32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *HostOnlineStatusHistoryDataTelemetry) SetOffline(v int32)`

SetOffline sets Offline field to given value.


### GetUnknown

`func (o *HostOnlineStatusHistoryDataTelemetry) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *HostOnlineStatusHistoryDataTelemetry) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *HostOnlineStatusHistoryDataTelemetry) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


