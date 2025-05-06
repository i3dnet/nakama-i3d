# HostHddDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**HardDisk0Free** | **int32** | The first hard disk&#39;s free space in bytes | [readonly] 
**HardDisk1Free** | **int32** | The second hard disk&#39;s free space in bytes | [readonly] 
**HardDisk2Free** | **int32** | The third hard disk&#39;s free space in bytes | [readonly] 
**HardDisk3Free** | **int32** | The fourth hard disk&#39;s free space in bytes | [readonly] 

## Methods

### NewHostHddDataTelemetry

`func NewHostHddDataTelemetry(timestamp int32, hardDisk0Free int32, hardDisk1Free int32, hardDisk2Free int32, hardDisk3Free int32, ) *HostHddDataTelemetry`

NewHostHddDataTelemetry instantiates a new HostHddDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHddDataTelemetryWithDefaults

`func NewHostHddDataTelemetryWithDefaults() *HostHddDataTelemetry`

NewHostHddDataTelemetryWithDefaults instantiates a new HostHddDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostHddDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostHddDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostHddDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetHardDisk0Free

`func (o *HostHddDataTelemetry) GetHardDisk0Free() int32`

GetHardDisk0Free returns the HardDisk0Free field if non-nil, zero value otherwise.

### GetHardDisk0FreeOk

`func (o *HostHddDataTelemetry) GetHardDisk0FreeOk() (*int32, bool)`

GetHardDisk0FreeOk returns a tuple with the HardDisk0Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDisk0Free

`func (o *HostHddDataTelemetry) SetHardDisk0Free(v int32)`

SetHardDisk0Free sets HardDisk0Free field to given value.


### GetHardDisk1Free

`func (o *HostHddDataTelemetry) GetHardDisk1Free() int32`

GetHardDisk1Free returns the HardDisk1Free field if non-nil, zero value otherwise.

### GetHardDisk1FreeOk

`func (o *HostHddDataTelemetry) GetHardDisk1FreeOk() (*int32, bool)`

GetHardDisk1FreeOk returns a tuple with the HardDisk1Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDisk1Free

`func (o *HostHddDataTelemetry) SetHardDisk1Free(v int32)`

SetHardDisk1Free sets HardDisk1Free field to given value.


### GetHardDisk2Free

`func (o *HostHddDataTelemetry) GetHardDisk2Free() int32`

GetHardDisk2Free returns the HardDisk2Free field if non-nil, zero value otherwise.

### GetHardDisk2FreeOk

`func (o *HostHddDataTelemetry) GetHardDisk2FreeOk() (*int32, bool)`

GetHardDisk2FreeOk returns a tuple with the HardDisk2Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDisk2Free

`func (o *HostHddDataTelemetry) SetHardDisk2Free(v int32)`

SetHardDisk2Free sets HardDisk2Free field to given value.


### GetHardDisk3Free

`func (o *HostHddDataTelemetry) GetHardDisk3Free() int32`

GetHardDisk3Free returns the HardDisk3Free field if non-nil, zero value otherwise.

### GetHardDisk3FreeOk

`func (o *HostHddDataTelemetry) GetHardDisk3FreeOk() (*int32, bool)`

GetHardDisk3FreeOk returns a tuple with the HardDisk3Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDisk3Free

`func (o *HostHddDataTelemetry) SetHardDisk3Free(v int32)`

SetHardDisk3Free sets HardDisk3Free field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


