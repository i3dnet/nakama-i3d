# ApplicationInstancesDataProviderTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | [readonly] 
**ApplicationInstanceCount** | **int32** | Number of application instances | [readonly] 

## Methods

### NewApplicationInstancesDataProviderTelemetry

`func NewApplicationInstancesDataProviderTelemetry(providerId int32, applicationInstanceCount int32, ) *ApplicationInstancesDataProviderTelemetry`

NewApplicationInstancesDataProviderTelemetry instantiates a new ApplicationInstancesDataProviderTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesDataProviderTelemetryWithDefaults

`func NewApplicationInstancesDataProviderTelemetryWithDefaults() *ApplicationInstancesDataProviderTelemetry`

NewApplicationInstancesDataProviderTelemetryWithDefaults instantiates a new ApplicationInstancesDataProviderTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ApplicationInstancesDataProviderTelemetry) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ApplicationInstancesDataProviderTelemetry) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ApplicationInstancesDataProviderTelemetry) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetApplicationInstanceCount

`func (o *ApplicationInstancesDataProviderTelemetry) GetApplicationInstanceCount() int32`

GetApplicationInstanceCount returns the ApplicationInstanceCount field if non-nil, zero value otherwise.

### GetApplicationInstanceCountOk

`func (o *ApplicationInstancesDataProviderTelemetry) GetApplicationInstanceCountOk() (*int32, bool)`

GetApplicationInstanceCountOk returns a tuple with the ApplicationInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceCount

`func (o *ApplicationInstancesDataProviderTelemetry) SetApplicationInstanceCount(v int32)`

SetApplicationInstanceCount sets ApplicationInstanceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


