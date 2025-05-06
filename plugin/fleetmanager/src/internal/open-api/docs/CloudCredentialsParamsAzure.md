# CloudCredentialsParamsAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**SubscriptionId** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecretKey** | **string** |  | 
**ResourceName** | **string** |  | 
**DefaultRegion** | **string** | (regionName from [&#x60;GET /v3/cloud/provider/28/dcLocation&#x60;](#/Cloud/get_v3_cloud_provider__providerId__dcLocation)) | 

## Methods

### NewCloudCredentialsParamsAzure

`func NewCloudCredentialsParamsAzure(tenantId string, subscriptionId string, clientId string, clientSecretKey string, resourceName string, defaultRegion string, ) *CloudCredentialsParamsAzure`

NewCloudCredentialsParamsAzure instantiates a new CloudCredentialsParamsAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsParamsAzureWithDefaults

`func NewCloudCredentialsParamsAzureWithDefaults() *CloudCredentialsParamsAzure`

NewCloudCredentialsParamsAzureWithDefaults instantiates a new CloudCredentialsParamsAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CloudCredentialsParamsAzure) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudCredentialsParamsAzure) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudCredentialsParamsAzure) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *CloudCredentialsParamsAzure) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CloudCredentialsParamsAzure) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CloudCredentialsParamsAzure) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetClientId

`func (o *CloudCredentialsParamsAzure) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CloudCredentialsParamsAzure) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CloudCredentialsParamsAzure) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecretKey

`func (o *CloudCredentialsParamsAzure) GetClientSecretKey() string`

GetClientSecretKey returns the ClientSecretKey field if non-nil, zero value otherwise.

### GetClientSecretKeyOk

`func (o *CloudCredentialsParamsAzure) GetClientSecretKeyOk() (*string, bool)`

GetClientSecretKeyOk returns a tuple with the ClientSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretKey

`func (o *CloudCredentialsParamsAzure) SetClientSecretKey(v string)`

SetClientSecretKey sets ClientSecretKey field to given value.


### GetResourceName

`func (o *CloudCredentialsParamsAzure) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CloudCredentialsParamsAzure) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CloudCredentialsParamsAzure) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetDefaultRegion

`func (o *CloudCredentialsParamsAzure) GetDefaultRegion() string`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *CloudCredentialsParamsAzure) GetDefaultRegionOk() (*string, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *CloudCredentialsParamsAzure) SetDefaultRegion(v string)`

SetDefaultRegion sets DefaultRegion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


