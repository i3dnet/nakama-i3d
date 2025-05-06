# CloudCredentialsParamsGcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** |  | 
**Network** | **string** |  | 
**AuthJson** | [**CloudCredentialsParamsGcpAuthJson**](CloudCredentialsParamsGcpAuthJson.md) |  | 

## Methods

### NewCloudCredentialsParamsGcp

`func NewCloudCredentialsParamsGcp(projectName string, network string, authJson CloudCredentialsParamsGcpAuthJson, ) *CloudCredentialsParamsGcp`

NewCloudCredentialsParamsGcp instantiates a new CloudCredentialsParamsGcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsParamsGcpWithDefaults

`func NewCloudCredentialsParamsGcpWithDefaults() *CloudCredentialsParamsGcp`

NewCloudCredentialsParamsGcpWithDefaults instantiates a new CloudCredentialsParamsGcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *CloudCredentialsParamsGcp) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CloudCredentialsParamsGcp) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CloudCredentialsParamsGcp) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetNetwork

`func (o *CloudCredentialsParamsGcp) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CloudCredentialsParamsGcp) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CloudCredentialsParamsGcp) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAuthJson

`func (o *CloudCredentialsParamsGcp) GetAuthJson() CloudCredentialsParamsGcpAuthJson`

GetAuthJson returns the AuthJson field if non-nil, zero value otherwise.

### GetAuthJsonOk

`func (o *CloudCredentialsParamsGcp) GetAuthJsonOk() (*CloudCredentialsParamsGcpAuthJson, bool)`

GetAuthJsonOk returns a tuple with the AuthJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthJson

`func (o *CloudCredentialsParamsGcp) SetAuthJson(v CloudCredentialsParamsGcpAuthJson)`

SetAuthJson sets AuthJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


