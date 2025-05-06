# CloudCredentialsParamsAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**DefaultRegion** | **string** | (regionName from [&#x60;GET /v3/cloud/provider/27/dcLocation&#x60;](#/Cloud/get_v3_cloud_provider__providerId__dcLocation)) | 
**KeyPair** | **NullableString** |  | 
**SecurityGroup** | **NullableString** |  | 
**IsGameLift** | **bool** |  | 

## Methods

### NewCloudCredentialsParamsAws

`func NewCloudCredentialsParamsAws(accessKeyId string, secretAccessKey string, defaultRegion string, keyPair NullableString, securityGroup NullableString, isGameLift bool, ) *CloudCredentialsParamsAws`

NewCloudCredentialsParamsAws instantiates a new CloudCredentialsParamsAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsParamsAwsWithDefaults

`func NewCloudCredentialsParamsAwsWithDefaults() *CloudCredentialsParamsAws`

NewCloudCredentialsParamsAwsWithDefaults instantiates a new CloudCredentialsParamsAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *CloudCredentialsParamsAws) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CloudCredentialsParamsAws) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CloudCredentialsParamsAws) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *CloudCredentialsParamsAws) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CloudCredentialsParamsAws) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CloudCredentialsParamsAws) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetDefaultRegion

`func (o *CloudCredentialsParamsAws) GetDefaultRegion() string`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *CloudCredentialsParamsAws) GetDefaultRegionOk() (*string, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *CloudCredentialsParamsAws) SetDefaultRegion(v string)`

SetDefaultRegion sets DefaultRegion field to given value.


### GetKeyPair

`func (o *CloudCredentialsParamsAws) GetKeyPair() string`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *CloudCredentialsParamsAws) GetKeyPairOk() (*string, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *CloudCredentialsParamsAws) SetKeyPair(v string)`

SetKeyPair sets KeyPair field to given value.


### SetKeyPairNil

`func (o *CloudCredentialsParamsAws) SetKeyPairNil(b bool)`

 SetKeyPairNil sets the value for KeyPair to be an explicit nil

### UnsetKeyPair
`func (o *CloudCredentialsParamsAws) UnsetKeyPair()`

UnsetKeyPair ensures that no value is present for KeyPair, not even an explicit nil
### GetSecurityGroup

`func (o *CloudCredentialsParamsAws) GetSecurityGroup() string`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *CloudCredentialsParamsAws) GetSecurityGroupOk() (*string, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *CloudCredentialsParamsAws) SetSecurityGroup(v string)`

SetSecurityGroup sets SecurityGroup field to given value.


### SetSecurityGroupNil

`func (o *CloudCredentialsParamsAws) SetSecurityGroupNil(b bool)`

 SetSecurityGroupNil sets the value for SecurityGroup to be an explicit nil

### UnsetSecurityGroup
`func (o *CloudCredentialsParamsAws) UnsetSecurityGroup()`

UnsetSecurityGroup ensures that no value is present for SecurityGroup, not even an explicit nil
### GetIsGameLift

`func (o *CloudCredentialsParamsAws) GetIsGameLift() bool`

GetIsGameLift returns the IsGameLift field if non-nil, zero value otherwise.

### GetIsGameLiftOk

`func (o *CloudCredentialsParamsAws) GetIsGameLiftOk() (*bool, bool)`

GetIsGameLiftOk returns a tuple with the IsGameLift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGameLift

`func (o *CloudCredentialsParamsAws) SetIsGameLift(v bool)`

SetIsGameLift sets IsGameLift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


