# BulkReserveHostByFleetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldFleetId** | **string** | Fleet ID that was assigned to the host prior to bulk reservation. | 
**NewFleetId** | **string** | Fleet ID that will be assigned to the host after completion of bulk reservation. | 
**HostId** | **int32** | ID of the host. | 
**InProgress** | **int32** | When the value is 1 then the host is not changed to the new fleet ID, when value is 0 than the host is changed to the new fleet ID. | 

## Methods

### NewBulkReserveHostByFleetStatus

`func NewBulkReserveHostByFleetStatus(oldFleetId string, newFleetId string, hostId int32, inProgress int32, ) *BulkReserveHostByFleetStatus`

NewBulkReserveHostByFleetStatus instantiates a new BulkReserveHostByFleetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReserveHostByFleetStatusWithDefaults

`func NewBulkReserveHostByFleetStatusWithDefaults() *BulkReserveHostByFleetStatus`

NewBulkReserveHostByFleetStatusWithDefaults instantiates a new BulkReserveHostByFleetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldFleetId

`func (o *BulkReserveHostByFleetStatus) GetOldFleetId() string`

GetOldFleetId returns the OldFleetId field if non-nil, zero value otherwise.

### GetOldFleetIdOk

`func (o *BulkReserveHostByFleetStatus) GetOldFleetIdOk() (*string, bool)`

GetOldFleetIdOk returns a tuple with the OldFleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFleetId

`func (o *BulkReserveHostByFleetStatus) SetOldFleetId(v string)`

SetOldFleetId sets OldFleetId field to given value.


### GetNewFleetId

`func (o *BulkReserveHostByFleetStatus) GetNewFleetId() string`

GetNewFleetId returns the NewFleetId field if non-nil, zero value otherwise.

### GetNewFleetIdOk

`func (o *BulkReserveHostByFleetStatus) GetNewFleetIdOk() (*string, bool)`

GetNewFleetIdOk returns a tuple with the NewFleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFleetId

`func (o *BulkReserveHostByFleetStatus) SetNewFleetId(v string)`

SetNewFleetId sets NewFleetId field to given value.


### GetHostId

`func (o *BulkReserveHostByFleetStatus) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *BulkReserveHostByFleetStatus) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *BulkReserveHostByFleetStatus) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetInProgress

`func (o *BulkReserveHostByFleetStatus) GetInProgress() int32`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *BulkReserveHostByFleetStatus) GetInProgressOk() (*int32, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *BulkReserveHostByFleetStatus) SetInProgress(v int32)`

SetInProgress sets InProgress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


