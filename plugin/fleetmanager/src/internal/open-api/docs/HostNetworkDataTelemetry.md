# HostNetworkDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The time at which the last update occurred | [readonly] 
**NetworkIngress** | **int32** | The network&#39;s ingress in bytes | [readonly] 
**NetworkEgress** | **int32** | The network&#39;s egress in bytes | [readonly] 

## Methods

### NewHostNetworkDataTelemetry

`func NewHostNetworkDataTelemetry(timestamp int32, networkIngress int32, networkEgress int32, ) *HostNetworkDataTelemetry`

NewHostNetworkDataTelemetry instantiates a new HostNetworkDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostNetworkDataTelemetryWithDefaults

`func NewHostNetworkDataTelemetryWithDefaults() *HostNetworkDataTelemetry`

NewHostNetworkDataTelemetryWithDefaults instantiates a new HostNetworkDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostNetworkDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostNetworkDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostNetworkDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetNetworkIngress

`func (o *HostNetworkDataTelemetry) GetNetworkIngress() int32`

GetNetworkIngress returns the NetworkIngress field if non-nil, zero value otherwise.

### GetNetworkIngressOk

`func (o *HostNetworkDataTelemetry) GetNetworkIngressOk() (*int32, bool)`

GetNetworkIngressOk returns a tuple with the NetworkIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIngress

`func (o *HostNetworkDataTelemetry) SetNetworkIngress(v int32)`

SetNetworkIngress sets NetworkIngress field to given value.


### GetNetworkEgress

`func (o *HostNetworkDataTelemetry) GetNetworkEgress() int32`

GetNetworkEgress returns the NetworkEgress field if non-nil, zero value otherwise.

### GetNetworkEgressOk

`func (o *HostNetworkDataTelemetry) GetNetworkEgressOk() (*int32, bool)`

GetNetworkEgressOk returns a tuple with the NetworkEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEgress

`func (o *HostNetworkDataTelemetry) SetNetworkEgress(v int32)`

SetNetworkEgress sets NetworkEgress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


