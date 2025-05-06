# CloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 
**Name** | **string** | User defined name for the cloud credential | 
**Params** | [**CloudCredentialsParams**](CloudCredentialsParams.md) |  | 
**CreatedAt** | **int32** | Unix timestamp of when credentials were first stored (informative only) | [readonly] 
**VerifiedAt** | **int32** | Unix timestamp of credential last verification time (informative only) | [readonly] 
**CredentialStatus** | **int32** | Status of cloud credential * 0: valid * 1: invalid | [readonly] 
**InvalidateAt** | **int32** | Unix timestamp of invalidation | [readonly] 
**InvalidateReason** | **string** | The reason for invalidation of the cloud credential | [readonly] 
**Status** | **int32** | Status of a currently in use cloud credential * 0: inactive * 1: active | 

## Methods

### NewCloudCredentials

`func NewCloudCredentials(providerId int32, name string, params CloudCredentialsParams, createdAt int32, verifiedAt int32, credentialStatus int32, invalidateAt int32, invalidateReason string, status int32, ) *CloudCredentials`

NewCloudCredentials instantiates a new CloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsWithDefaults

`func NewCloudCredentialsWithDefaults() *CloudCredentials`

NewCloudCredentialsWithDefaults instantiates a new CloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CloudCredentials) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CloudCredentials) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CloudCredentials) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetName

`func (o *CloudCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *CloudCredentials) GetParams() CloudCredentialsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CloudCredentials) GetParamsOk() (*CloudCredentialsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CloudCredentials) SetParams(v CloudCredentialsParams)`

SetParams sets Params field to given value.


### GetCreatedAt

`func (o *CloudCredentials) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCredentials) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCredentials) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetVerifiedAt

`func (o *CloudCredentials) GetVerifiedAt() int32`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *CloudCredentials) GetVerifiedAtOk() (*int32, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *CloudCredentials) SetVerifiedAt(v int32)`

SetVerifiedAt sets VerifiedAt field to given value.


### GetCredentialStatus

`func (o *CloudCredentials) GetCredentialStatus() int32`

GetCredentialStatus returns the CredentialStatus field if non-nil, zero value otherwise.

### GetCredentialStatusOk

`func (o *CloudCredentials) GetCredentialStatusOk() (*int32, bool)`

GetCredentialStatusOk returns a tuple with the CredentialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialStatus

`func (o *CloudCredentials) SetCredentialStatus(v int32)`

SetCredentialStatus sets CredentialStatus field to given value.


### GetInvalidateAt

`func (o *CloudCredentials) GetInvalidateAt() int32`

GetInvalidateAt returns the InvalidateAt field if non-nil, zero value otherwise.

### GetInvalidateAtOk

`func (o *CloudCredentials) GetInvalidateAtOk() (*int32, bool)`

GetInvalidateAtOk returns a tuple with the InvalidateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidateAt

`func (o *CloudCredentials) SetInvalidateAt(v int32)`

SetInvalidateAt sets InvalidateAt field to given value.


### GetInvalidateReason

`func (o *CloudCredentials) GetInvalidateReason() string`

GetInvalidateReason returns the InvalidateReason field if non-nil, zero value otherwise.

### GetInvalidateReasonOk

`func (o *CloudCredentials) GetInvalidateReasonOk() (*string, bool)`

GetInvalidateReasonOk returns a tuple with the InvalidateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidateReason

`func (o *CloudCredentials) SetInvalidateReason(v string)`

SetInvalidateReason sets InvalidateReason field to given value.


### GetStatus

`func (o *CloudCredentials) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCredentials) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCredentials) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


