# CloudCredentialsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** |  | 
**SecretAccessKey** | **string** |  | 
**DefaultRegion** | **string** |  | 
**KeyPair** | **string** |  | 
**SecurityGroup** | **string** |  | 
**IsGameLift** | **bool** |  | 
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecretKey** | **string** |  | 
**ResourceName** | **string** |  | 
**ProjectName** | **string** |  | 
**Network** | **string** |  | 
**AuthJson** | [**CloudCredentialsParamsGcpAuthJson**](CloudCredentialsParamsGcpAuthJson.md) |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewCloudCredentialsParams

`func NewCloudCredentialsParams(accessKeyId string, secretAccessKey string, defaultRegion string, keyPair string, securityGroup string, isGameLift bool, tenantId string, subscriptionId string, clientId string, clientSecretKey string, resourceName string, projectName string, network string, authJson CloudCredentialsParamsGcpAuthJson, username string, password string, projectId string, ) *CloudCredentialsParams`

NewCloudCredentialsParams instantiates a new CloudCredentialsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsParamsWithDefaults

`func NewCloudCredentialsParamsWithDefaults() *CloudCredentialsParams`

NewCloudCredentialsParamsWithDefaults instantiates a new CloudCredentialsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *CloudCredentialsParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CloudCredentialsParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CloudCredentialsParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *CloudCredentialsParams) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CloudCredentialsParams) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CloudCredentialsParams) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### GetDefaultRegion

`func (o *CloudCredentialsParams) GetDefaultRegion() string`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *CloudCredentialsParams) GetDefaultRegionOk() (*string, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *CloudCredentialsParams) SetDefaultRegion(v string)`

SetDefaultRegion sets DefaultRegion field to given value.


### GetKeyPair

`func (o *CloudCredentialsParams) GetKeyPair() string`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *CloudCredentialsParams) GetKeyPairOk() (*string, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *CloudCredentialsParams) SetKeyPair(v string)`

SetKeyPair sets KeyPair field to given value.


### GetSecurityGroup

`func (o *CloudCredentialsParams) GetSecurityGroup() string`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *CloudCredentialsParams) GetSecurityGroupOk() (*string, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *CloudCredentialsParams) SetSecurityGroup(v string)`

SetSecurityGroup sets SecurityGroup field to given value.


### GetIsGameLift

`func (o *CloudCredentialsParams) GetIsGameLift() bool`

GetIsGameLift returns the IsGameLift field if non-nil, zero value otherwise.

### GetIsGameLiftOk

`func (o *CloudCredentialsParams) GetIsGameLiftOk() (*bool, bool)`

GetIsGameLiftOk returns a tuple with the IsGameLift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGameLift

`func (o *CloudCredentialsParams) SetIsGameLift(v bool)`

SetIsGameLift sets IsGameLift field to given value.


### GetTenantId

`func (o *CloudCredentialsParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudCredentialsParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudCredentialsParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *CloudCredentialsParams) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CloudCredentialsParams) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CloudCredentialsParams) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetClientId

`func (o *CloudCredentialsParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CloudCredentialsParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CloudCredentialsParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecretKey

`func (o *CloudCredentialsParams) GetClientSecretKey() string`

GetClientSecretKey returns the ClientSecretKey field if non-nil, zero value otherwise.

### GetClientSecretKeyOk

`func (o *CloudCredentialsParams) GetClientSecretKeyOk() (*string, bool)`

GetClientSecretKeyOk returns a tuple with the ClientSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretKey

`func (o *CloudCredentialsParams) SetClientSecretKey(v string)`

SetClientSecretKey sets ClientSecretKey field to given value.


### GetResourceName

`func (o *CloudCredentialsParams) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CloudCredentialsParams) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CloudCredentialsParams) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetProjectName

`func (o *CloudCredentialsParams) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CloudCredentialsParams) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CloudCredentialsParams) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetNetwork

`func (o *CloudCredentialsParams) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CloudCredentialsParams) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CloudCredentialsParams) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAuthJson

`func (o *CloudCredentialsParams) GetAuthJson() CloudCredentialsParamsGcpAuthJson`

GetAuthJson returns the AuthJson field if non-nil, zero value otherwise.

### GetAuthJsonOk

`func (o *CloudCredentialsParams) GetAuthJsonOk() (*CloudCredentialsParamsGcpAuthJson, bool)`

GetAuthJsonOk returns a tuple with the AuthJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthJson

`func (o *CloudCredentialsParams) SetAuthJson(v CloudCredentialsParamsGcpAuthJson)`

SetAuthJson sets AuthJson field to given value.


### GetUsername

`func (o *CloudCredentialsParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudCredentialsParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudCredentialsParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CloudCredentialsParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudCredentialsParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudCredentialsParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProjectId

`func (o *CloudCredentialsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CloudCredentialsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CloudCredentialsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


