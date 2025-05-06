# HostOnlineStatusHistoryHostTypeTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**Data** | [**[]HostOnlineStatusHistoryDataTelemetry**](HostOnlineStatusHistoryDataTelemetry.md) | Contains the data object | [readonly] 

## Methods

### NewHostOnlineStatusHistoryHostTypeTelemetry

`func NewHostOnlineStatusHistoryHostTypeTelemetry(timestamp int32, data []HostOnlineStatusHistoryDataTelemetry, ) *HostOnlineStatusHistoryHostTypeTelemetry`

NewHostOnlineStatusHistoryHostTypeTelemetry instantiates a new HostOnlineStatusHistoryHostTypeTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOnlineStatusHistoryHostTypeTelemetryWithDefaults

`func NewHostOnlineStatusHistoryHostTypeTelemetryWithDefaults() *HostOnlineStatusHistoryHostTypeTelemetry`

NewHostOnlineStatusHistoryHostTypeTelemetryWithDefaults instantiates a new HostOnlineStatusHistoryHostTypeTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) GetData() []HostOnlineStatusHistoryDataTelemetry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) GetDataOk() (*[]HostOnlineStatusHistoryDataTelemetry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostOnlineStatusHistoryHostTypeTelemetry) SetData(v []HostOnlineStatusHistoryDataTelemetry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


