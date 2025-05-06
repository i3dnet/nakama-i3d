# CCUStatusDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**MaxPlayers** | **int32** | Maximum number of players the hosts can support | [readonly] 
**CurrentPlayers** | **int32** | Number of players currently playing on the hosts | [readonly] 

## Methods

### NewCCUStatusDataTelemetry

`func NewCCUStatusDataTelemetry(timestamp int32, maxPlayers int32, currentPlayers int32, ) *CCUStatusDataTelemetry`

NewCCUStatusDataTelemetry instantiates a new CCUStatusDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUStatusDataTelemetryWithDefaults

`func NewCCUStatusDataTelemetryWithDefaults() *CCUStatusDataTelemetry`

NewCCUStatusDataTelemetryWithDefaults instantiates a new CCUStatusDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *CCUStatusDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CCUStatusDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CCUStatusDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetMaxPlayers

`func (o *CCUStatusDataTelemetry) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *CCUStatusDataTelemetry) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *CCUStatusDataTelemetry) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.


### GetCurrentPlayers

`func (o *CCUStatusDataTelemetry) GetCurrentPlayers() int32`

GetCurrentPlayers returns the CurrentPlayers field if non-nil, zero value otherwise.

### GetCurrentPlayersOk

`func (o *CCUStatusDataTelemetry) GetCurrentPlayersOk() (*int32, bool)`

GetCurrentPlayersOk returns a tuple with the CurrentPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlayers

`func (o *CCUStatusDataTelemetry) SetCurrentPlayers(v int32)`

SetCurrentPlayers sets CurrentPlayers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


