# ApplicationInstancesDataTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | The timestamp on which data is calculated | [readonly] 
**ApplicationInstances** | [**[]ApplicationInstancesDataProviderTelemetry**](ApplicationInstancesDataProviderTelemetry.md) | List of application instances | [readonly] 

## Methods

### NewApplicationInstancesDataTelemetry

`func NewApplicationInstancesDataTelemetry(timestamp int32, applicationInstances []ApplicationInstancesDataProviderTelemetry, ) *ApplicationInstancesDataTelemetry`

NewApplicationInstancesDataTelemetry instantiates a new ApplicationInstancesDataTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesDataTelemetryWithDefaults

`func NewApplicationInstancesDataTelemetryWithDefaults() *ApplicationInstancesDataTelemetry`

NewApplicationInstancesDataTelemetryWithDefaults instantiates a new ApplicationInstancesDataTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationInstancesDataTelemetry) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationInstancesDataTelemetry) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationInstancesDataTelemetry) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetApplicationInstances

`func (o *ApplicationInstancesDataTelemetry) GetApplicationInstances() []ApplicationInstancesDataProviderTelemetry`

GetApplicationInstances returns the ApplicationInstances field if non-nil, zero value otherwise.

### GetApplicationInstancesOk

`func (o *ApplicationInstancesDataTelemetry) GetApplicationInstancesOk() (*[]ApplicationInstancesDataProviderTelemetry, bool)`

GetApplicationInstancesOk returns a tuple with the ApplicationInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstances

`func (o *ApplicationInstancesDataTelemetry) SetApplicationInstances(v []ApplicationInstancesDataProviderTelemetry)`

SetApplicationInstances sets ApplicationInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


