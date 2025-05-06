# \CloudCredentialsAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudConfigurationCredential**](CloudCredentialsAPI.md#CreateCloudConfigurationCredential) | **Post** /v3/cloud/configuration/credential | Store new cloud provider credentials
[**CreateCloudConfigurationCredentialValidate**](CloudCredentialsAPI.md#CreateCloudConfigurationCredentialValidate) | **Post** /v3/cloud/configuration/credential/validate | Validate cloud provider credentials
[**DeleteCloudConfigurationCredential**](CloudCredentialsAPI.md#DeleteCloudConfigurationCredential) | **Delete** /v3/cloud/configuration/credential/{credentialName} | Delete a cloud provider credential for a given name
[**GetCloudConfigurationCredential**](CloudCredentialsAPI.md#GetCloudConfigurationCredential) | **Get** /v3/cloud/configuration/credential/{credentialName} | Get a list of cloud provider credentials for a given credential name
[**GetCloudConfigurationCredentials**](CloudCredentialsAPI.md#GetCloudConfigurationCredentials) | **Get** /v3/cloud/configuration/credential | Get list of cloud provider credentials
[**UpdateCloudConfigurationCredential**](CloudCredentialsAPI.md#UpdateCloudConfigurationCredential) | **Put** /v3/cloud/configuration/credential/{credentialName} | Update name and status of credential.



## CreateCloudConfigurationCredential

> []CloudCredentials CreateCloudConfigurationCredential(ctx).CloudCredentials(cloudCredentials).Execute()

Store new cloud provider credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {
	cloudCredentials := *openapiclient.NewCloudCredentials(int32(123), "Name_example", *openapiclient.NewCloudCredentialsParams("AccessKeyId_example", "SecretAccessKey_example", "DefaultRegion_example", "KeyPair_example", "SecurityGroup_example", false, "TenantId_example", "SubscriptionId_example", "ClientId_example", "ClientSecretKey_example", "ResourceName_example", "ProjectName_example", "Network_example", *openapiclient.NewCloudCredentialsParamsGcpAuthJson("Type_example", "ProjectId_example", "PrivateKeyId_example", "PrivateKey_example", "ClientEmail_example", "ClientId_example", "AuthUri_example", "TokenUri_example", "AuthProviderX509CertUrl_example", "ClientX509CertUrl_example"), "Username_example", "Password_example", "ProjectId_example"), int32(123), int32(123), int32(123), int32(123), "InvalidateReason_example", int32(123)) // CloudCredentials | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialsAPI.CreateCloudConfigurationCredential(context.Background()).CloudCredentials(cloudCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.CreateCloudConfigurationCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudConfigurationCredential`: []CloudCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialsAPI.CreateCloudConfigurationCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudConfigurationCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCredentials** | [**CloudCredentials**](CloudCredentials.md) |  | 

### Return type

[**[]CloudCredentials**](CloudCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudConfigurationCredentialValidate

> CreateCloudConfigurationCredentialValidate(ctx).CloudCredentialsValidation(cloudCredentialsValidation).Execute()

Validate cloud provider credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {
	cloudCredentialsValidation := *openapiclient.NewCloudCredentialsValidation(int32(123), *openapiclient.NewCloudCredentialsParams("AccessKeyId_example", "SecretAccessKey_example", "DefaultRegion_example", "KeyPair_example", "SecurityGroup_example", false, "TenantId_example", "SubscriptionId_example", "ClientId_example", "ClientSecretKey_example", "ResourceName_example", "ProjectName_example", "Network_example", *openapiclient.NewCloudCredentialsParamsGcpAuthJson("Type_example", "ProjectId_example", "PrivateKeyId_example", "PrivateKey_example", "ClientEmail_example", "ClientId_example", "AuthUri_example", "TokenUri_example", "AuthProviderX509CertUrl_example", "ClientX509CertUrl_example"), "Username_example", "Password_example", "ProjectId_example")) // CloudCredentialsValidation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudCredentialsAPI.CreateCloudConfigurationCredentialValidate(context.Background()).CloudCredentialsValidation(cloudCredentialsValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.CreateCloudConfigurationCredentialValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudConfigurationCredentialValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCredentialsValidation** | [**CloudCredentialsValidation**](CloudCredentialsValidation.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudConfigurationCredential

> DeleteCloudConfigurationCredential(ctx, credentialName).Execute()

Delete a cloud provider credential for a given name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {
	credentialName := "credentialName_example" // string | The user defined name for cloud credentials.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudCredentialsAPI.DeleteCloudConfigurationCredential(context.Background(), credentialName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.DeleteCloudConfigurationCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialName** | **string** | The user defined name for cloud credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudConfigurationCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudConfigurationCredential

> []CloudCredentials GetCloudConfigurationCredential(ctx, credentialName).Execute()

Get a list of cloud provider credentials for a given credential name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {
	credentialName := "credentialName_example" // string | The user defined name for cloud credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialsAPI.GetCloudConfigurationCredential(context.Background(), credentialName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.GetCloudConfigurationCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudConfigurationCredential`: []CloudCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialsAPI.GetCloudConfigurationCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialName** | **string** | The user defined name for cloud credentials | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConfigurationCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CloudCredentials**](CloudCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudConfigurationCredentials

> []CloudCredentials GetCloudConfigurationCredentials(ctx).Execute()

Get list of cloud provider credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialsAPI.GetCloudConfigurationCredentials(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.GetCloudConfigurationCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudConfigurationCredentials`: []CloudCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialsAPI.GetCloudConfigurationCredentials`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConfigurationCredentialsRequest struct via the builder pattern


### Return type

[**[]CloudCredentials**](CloudCredentials.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudConfigurationCredential

> []CloudCredentialsMetaData UpdateCloudConfigurationCredential(ctx, credentialName).CloudCredentialsMetaData(cloudCredentialsMetaData).Execute()

Update name and status of credential.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/open-api"
)

func main() {
	credentialName := "credentialName_example" // string | The user defined name for cloud credentials.
	cloudCredentialsMetaData := *openapiclient.NewCloudCredentialsMetaData("Name_example", int32(123)) // CloudCredentialsMetaData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialsAPI.UpdateCloudConfigurationCredential(context.Background(), credentialName).CloudCredentialsMetaData(cloudCredentialsMetaData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialsAPI.UpdateCloudConfigurationCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudConfigurationCredential`: []CloudCredentialsMetaData
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialsAPI.UpdateCloudConfigurationCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialName** | **string** | The user defined name for cloud credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudConfigurationCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCredentialsMetaData** | [**CloudCredentialsMetaData**](CloudCredentialsMetaData.md) |  | 

### Return type

[**[]CloudCredentialsMetaData**](CloudCredentialsMetaData.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

