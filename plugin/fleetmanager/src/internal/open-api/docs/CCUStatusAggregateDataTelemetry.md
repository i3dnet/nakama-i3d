# CCUStatusAggregateDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | [readonly] 
**MaxPlayers** | **int32** | Maximum number of players the hosts can support | [readonly] 
**CurrentPlayers** | **int32** | Number of players currently playing on the hosts | [readonly] 

## Methods

### NewCCUStatusAggregateDataTelemetry

`func NewCCUStatusAggregateDataTelemetry(providerId int32, maxPlayers int32, currentPlayers int32, ) *CCUStatusAggregateDataTelemetry`

NewCCUStatusAggregateDataTelemetry instantiates a new CCUStatusAggregateDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUStatusAggregateDataTelemetryWithDefaults

`func NewCCUStatusAggregateDataTelemetryWithDefaults() *CCUStatusAggregateDataTelemetry`

NewCCUStatusAggregateDataTelemetryWithDefaults instantiates a new CCUStatusAggregateDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CCUStatusAggregateDataTelemetry) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CCUStatusAggregateDataTelemetry) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CCUStatusAggregateDataTelemetry) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetMaxPlayers

`func (o *CCUStatusAggregateDataTelemetry) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *CCUStatusAggregateDataTelemetry) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *CCUStatusAggregateDataTelemetry) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.


### GetCurrentPlayers

`func (o *CCUStatusAggregateDataTelemetry) GetCurrentPlayers() int32`

GetCurrentPlayers returns the CurrentPlayers field if non-nil, zero value otherwise.

### GetCurrentPlayersOk

`func (o *CCUStatusAggregateDataTelemetry) GetCurrentPlayersOk() (*int32, bool)`

GetCurrentPlayersOk returns a tuple with the CurrentPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlayers

`func (o *CCUStatusAggregateDataTelemetry) SetCurrentPlayers(v int32)`

SetCurrentPlayers sets CurrentPlayers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


