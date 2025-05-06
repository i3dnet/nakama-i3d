# ApplicationInstanceStatusDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The timestamp on which data is calculated | [readonly] 
**Setup** | **int32** | Average number of application instances with setup status | [readonly] 
**Offline** | **int32** | Average number of application instances with offline status | [readonly] 
**Starting** | **int32** | Average number of application instances with starting status | [readonly] 
**Online** | **int32** | Average number of application instances with online status | [readonly] 
**Allocated** | **int32** | Average number of application instances with allocated status | [readonly] 

## Methods

### NewApplicationInstanceStatusDataTelemetry

`func NewApplicationInstanceStatusDataTelemetry(timestamp int32, setup int32, offline int32, starting int32, online int32, allocated int32, ) *ApplicationInstanceStatusDataTelemetry`

NewApplicationInstanceStatusDataTelemetry instantiates a new ApplicationInstanceStatusDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStatusDataTelemetryWithDefaults

`func NewApplicationInstanceStatusDataTelemetryWithDefaults() *ApplicationInstanceStatusDataTelemetry`

NewApplicationInstanceStatusDataTelemetryWithDefaults instantiates a new ApplicationInstanceStatusDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstanceStatusDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstanceStatusDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSetup

`func (o *ApplicationInstanceStatusDataTelemetry) GetSetup() int32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetSetupOk() (*int32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ApplicationInstanceStatusDataTelemetry) SetSetup(v int32)`

SetSetup sets Setup field to given value.


### GetOffline

`func (o *ApplicationInstanceStatusDataTelemetry) GetOffline() int32`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetOfflineOk() (*int32, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *ApplicationInstanceStatusDataTelemetry) SetOffline(v int32)`

SetOffline sets Offline field to given value.


### GetStarting

`func (o *ApplicationInstanceStatusDataTelemetry) GetStarting() int32`

GetStarting returns the Starting field if non-nil, zero value otherwise.

### GetStartingOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetStartingOk() (*int32, bool)`

GetStartingOk returns a tuple with the Starting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarting

`func (o *ApplicationInstanceStatusDataTelemetry) SetStarting(v int32)`

SetStarting sets Starting field to given value.


### GetOnline

`func (o *ApplicationInstanceStatusDataTelemetry) GetOnline() int32`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetOnlineOk() (*int32, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApplicationInstanceStatusDataTelemetry) SetOnline(v int32)`

SetOnline sets Online field to given value.


### GetAllocated

`func (o *ApplicationInstanceStatusDataTelemetry) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *ApplicationInstanceStatusDataTelemetry) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *ApplicationInstanceStatusDataTelemetry) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


