# CloudCredentialsValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 
**Params** | [**CloudCredentialsParams**](CloudCredentialsParams.md) |  | 

## Methods

### NewCloudCredentialsValidation

`func NewCloudCredentialsValidation(providerId int32, params CloudCredentialsParams, ) *CloudCredentialsValidation`

NewCloudCredentialsValidation instantiates a new CloudCredentialsValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsValidationWithDefaults

`func NewCloudCredentialsValidationWithDefaults() *CloudCredentialsValidation`

NewCloudCredentialsValidationWithDefaults instantiates a new CloudCredentialsValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CloudCredentialsValidation) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CloudCredentialsValidation) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CloudCredentialsValidation) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetParams

`func (o *CloudCredentialsValidation) GetParams() CloudCredentialsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CloudCredentialsValidation) GetParamsOk() (*CloudCredentialsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CloudCredentialsValidation) SetParams(v CloudCredentialsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


