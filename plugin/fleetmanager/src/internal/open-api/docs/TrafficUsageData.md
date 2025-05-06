# TrafficUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The corresponding date | [readonly] 
**Ingress** | **string** | The incoming data usage in bytes | [readonly] 
**Egress** | **string** | The outgoing data usage in bytes | [readonly] 
**Sum** | **string** | The total used data in bytes | [readonly] 

## Methods

### NewTrafficUsageData

`func NewTrafficUsageData(timestamp int32, ingress string, egress string, sum string, ) *TrafficUsageData`

NewTrafficUsageData instantiates a new TrafficUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficUsageDataWithDefaults

`func NewTrafficUsageDataWithDefaults() *TrafficUsageData`

NewTrafficUsageDataWithDefaults instantiates a new TrafficUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TrafficUsageData) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrafficUsageData) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrafficUsageData) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetIngress

`func (o *TrafficUsageData) GetIngress() string`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *TrafficUsageData) GetIngressOk() (*string, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *TrafficUsageData) SetIngress(v string)`

SetIngress sets Ingress field to given value.


### GetEgress

`func (o *TrafficUsageData) GetEgress() string`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *TrafficUsageData) GetEgressOk() (*string, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *TrafficUsageData) SetEgress(v string)`

SetEgress sets Egress field to given value.


### GetSum

`func (o *TrafficUsageData) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *TrafficUsageData) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *TrafficUsageData) SetSum(v string)`

SetSum sets Sum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


