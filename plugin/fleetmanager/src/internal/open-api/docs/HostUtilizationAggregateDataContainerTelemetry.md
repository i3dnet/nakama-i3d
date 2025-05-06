# HostUtilizationAggregateDataContainerTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Unix time in which the data was saved | [readonly] 
**Hosts** | [**[]HostUtilizationAggregateDataTelemetry**](HostUtilizationAggregateDataTelemetry.md) | List containing the telemetry data grouped by providerId | [readonly] 

## Methods

### NewHostUtilizationAggregateDataContainerTelemetry

`func NewHostUtilizationAggregateDataContainerTelemetry(timestamp int32, hosts []HostUtilizationAggregateDataTelemetry, ) *HostUtilizationAggregateDataContainerTelemetry`

NewHostUtilizationAggregateDataContainerTelemetry instantiates a new HostUtilizationAggregateDataContainerTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUtilizationAggregateDataContainerTelemetryWithDefaults

`func NewHostUtilizationAggregateDataContainerTelemetryWithDefaults() *HostUtilizationAggregateDataContainerTelemetry`

NewHostUtilizationAggregateDataContainerTelemetryWithDefaults instantiates a new HostUtilizationAggregateDataContainerTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *HostUtilizationAggregateDataContainerTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostUtilizationAggregateDataContainerTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostUtilizationAggregateDataContainerTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetHosts

`func (o *HostUtilizationAggregateDataContainerTelemetry) GetHosts() []HostUtilizationAggregateDataTelemetry`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HostUtilizationAggregateDataContainerTelemetry) GetHostsOk() (*[]HostUtilizationAggregateDataTelemetry, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HostUtilizationAggregateDataContainerTelemetry) SetHosts(v []HostUtilizationAggregateDataTelemetry)`

SetHosts sets Hosts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


