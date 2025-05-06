# BulkReserveHostByFleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **string** | Fleet ID | 
**Hosts** | **[]int32** | List of host ID&#39;s | 
**MethodId** | Pointer to **int32** | Stop method type, list of types can be found in: [GET /v3/application/stopMethod](game-publisher#/Application/getApplicationStopMethods) | [optional] 
**Timeout** | Pointer to **int32** | Stop method timeout in seconds, default value is 1800 seconds | [optional] 

## Methods

### NewBulkReserveHostByFleet

`func NewBulkReserveHostByFleet(fleetId string, hosts []int32, ) *BulkReserveHostByFleet`

NewBulkReserveHostByFleet instantiates a new BulkReserveHostByFleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReserveHostByFleetWithDefaults

`func NewBulkReserveHostByFleetWithDefaults() *BulkReserveHostByFleet`

NewBulkReserveHostByFleetWithDefaults instantiates a new BulkReserveHostByFleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *BulkReserveHostByFleet) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BulkReserveHostByFleet) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BulkReserveHostByFleet) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetHosts

`func (o *BulkReserveHostByFleet) GetHosts() []int32`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *BulkReserveHostByFleet) GetHostsOk() (*[]int32, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *BulkReserveHostByFleet) SetHosts(v []int32)`

SetHosts sets Hosts field to given value.


### GetMethodId

`func (o *BulkReserveHostByFleet) GetMethodId() int32`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *BulkReserveHostByFleet) GetMethodIdOk() (*int32, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *BulkReserveHostByFleet) SetMethodId(v int32)`

SetMethodId sets MethodId field to given value.

### HasMethodId

`func (o *BulkReserveHostByFleet) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.

### GetTimeout

`func (o *BulkReserveHostByFleet) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BulkReserveHostByFleet) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BulkReserveHostByFleet) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BulkReserveHostByFleet) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


