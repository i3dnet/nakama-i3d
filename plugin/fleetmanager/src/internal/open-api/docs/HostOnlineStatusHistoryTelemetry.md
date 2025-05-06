# HostOnlineStatusHistoryTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Unit of data frequency | [readonly] 
**Host** | [**[]HostOnlineStatusHistoryHostTypeTelemetry**](HostOnlineStatusHistoryHostTypeTelemetry.md) | Contains the data object by host type | [readonly] 

## Methods

### NewHostOnlineStatusHistoryTelemetry

`func NewHostOnlineStatusHistoryTelemetry(unit string, host []HostOnlineStatusHistoryHostTypeTelemetry, ) *HostOnlineStatusHistoryTelemetry`

NewHostOnlineStatusHistoryTelemetry instantiates a new HostOnlineStatusHistoryTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOnlineStatusHistoryTelemetryWithDefaults

`func NewHostOnlineStatusHistoryTelemetryWithDefaults() *HostOnlineStatusHistoryTelemetry`

NewHostOnlineStatusHistoryTelemetryWithDefaults instantiates a new HostOnlineStatusHistoryTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HostOnlineStatusHistoryTelemetry) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostOnlineStatusHistoryTelemetry) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostOnlineStatusHistoryTelemetry) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetHost

`func (o *HostOnlineStatusHistoryTelemetry) GetHost() []HostOnlineStatusHistoryHostTypeTelemetry`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostOnlineStatusHistoryTelemetry) GetHostOk() (*[]HostOnlineStatusHistoryHostTypeTelemetry, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostOnlineStatusHistoryTelemetry) SetHost(v []HostOnlineStatusHistoryHostTypeTelemetry)`

SetHost sets Host field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


