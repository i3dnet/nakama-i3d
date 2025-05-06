# HostOveruseTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **int32** | The ID of the host | [readonly] 
**HostName** | **string** | The name of the physical machine | [readonly] 
**OrderedBandwidth** | **float32** | Allowed bandwidth of the host | [readonly] 
**UsedBandwidth** | **float32** | Used bandwidth of the host | [readonly] 

## Methods

### NewHostOveruseTelemetry

`func NewHostOveruseTelemetry(hostId int32, hostName string, orderedBandwidth float32, usedBandwidth float32, ) *HostOveruseTelemetry`

NewHostOveruseTelemetry instantiates a new HostOveruseTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOveruseTelemetryWithDefaults

`func NewHostOveruseTelemetryWithDefaults() *HostOveruseTelemetry`

NewHostOveruseTelemetryWithDefaults instantiates a new HostOveruseTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *HostOveruseTelemetry) GetHostId() int32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *HostOveruseTelemetry) GetHostIdOk() (*int32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *HostOveruseTelemetry) SetHostId(v int32)`

SetHostId sets HostId field to given value.


### GetHostName

`func (o *HostOveruseTelemetry) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HostOveruseTelemetry) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HostOveruseTelemetry) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetOrderedBandwidth

`func (o *HostOveruseTelemetry) GetOrderedBandwidth() float32`

GetOrderedBandwidth returns the OrderedBandwidth field if non-nil, zero value otherwise.

### GetOrderedBandwidthOk

`func (o *HostOveruseTelemetry) GetOrderedBandwidthOk() (*float32, bool)`

GetOrderedBandwidthOk returns a tuple with the OrderedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedBandwidth

`func (o *HostOveruseTelemetry) SetOrderedBandwidth(v float32)`

SetOrderedBandwidth sets OrderedBandwidth field to given value.


### GetUsedBandwidth

`func (o *HostOveruseTelemetry) GetUsedBandwidth() float32`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *HostOveruseTelemetry) GetUsedBandwidthOk() (*float32, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *HostOveruseTelemetry) SetUsedBandwidth(v float32)`

SetUsedBandwidth sets UsedBandwidth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


