# CCUApplicationTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | Application ID | [readonly] 
**ApplicationName** | **string** | The public name of your application | 
**PlayerCount** | **int32** | Number of players currently playing on the hosts | [readonly] 
**MaxPlayers** | **int32** | Maximum number of players the hosts can support | [readonly] 

## Methods

### NewCCUApplicationTelemetry

`func NewCCUApplicationTelemetry(applicationId string, applicationName string, playerCount int32, maxPlayers int32, ) *CCUApplicationTelemetry`

NewCCUApplicationTelemetry instantiates a new CCUApplicationTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCUApplicationTelemetryWithDefaults

`func NewCCUApplicationTelemetryWithDefaults() *CCUApplicationTelemetry`

NewCCUApplicationTelemetryWithDefaults instantiates a new CCUApplicationTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CCUApplicationTelemetry) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CCUApplicationTelemetry) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CCUApplicationTelemetry) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *CCUApplicationTelemetry) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *CCUApplicationTelemetry) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *CCUApplicationTelemetry) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetPlayerCount

`func (o *CCUApplicationTelemetry) GetPlayerCount() int32`

GetPlayerCount returns the PlayerCount field if non-nil, zero value otherwise.

### GetPlayerCountOk

`func (o *CCUApplicationTelemetry) GetPlayerCountOk() (*int32, bool)`

GetPlayerCountOk returns a tuple with the PlayerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerCount

`func (o *CCUApplicationTelemetry) SetPlayerCount(v int32)`

SetPlayerCount sets PlayerCount field to given value.


### GetMaxPlayers

`func (o *CCUApplicationTelemetry) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *CCUApplicationTelemetry) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *CCUApplicationTelemetry) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


