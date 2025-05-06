# \ApplicationInstallCreateAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationInstallCreate**](ApplicationInstallCreateAPI.md#CreateApplicationInstallCreate) | **Post** /v3/applicationInstall/create | Create a new application installation process
[**DeleteApplicationInstallCreate**](ApplicationInstallCreateAPI.md#DeleteApplicationInstallCreate) | **Delete** /v3/applicationInstall/create/{fileId} | Cancel a application installation process
[**GetApplicationInstallCreate**](ApplicationInstallCreateAPI.md#GetApplicationInstallCreate) | **Get** /v3/applicationInstall/create/{fileId} | Get a create application install process
[**GetApplicationInstallCreates**](ApplicationInstallCreateAPI.md#GetApplicationInstallCreates) | **Get** /v3/applicationInstall/create | Get a list of create application install processes



## CreateApplicationInstallCreate

> []CreateApplicationInstall CreateApplicationInstallCreate(ctx).NewCreateApplicationInstall(newCreateApplicationInstall).Execute()

Create a new application installation process

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
	newCreateApplicationInstall := *openapiclient.NewNewCreateApplicationInstall("ApplicationId_example", int32(123), int32(123), "Name_example", "Version_example", int32(123)) // NewCreateApplicationInstall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstallCreateAPI.CreateApplicationInstallCreate(context.Background()).NewCreateApplicationInstall(newCreateApplicationInstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallCreateAPI.CreateApplicationInstallCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstallCreate`: []CreateApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallCreateAPI.CreateApplicationInstallCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstallCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newCreateApplicationInstall** | [**NewCreateApplicationInstall**](NewCreateApplicationInstall.md) |  | 

### Return type

[**[]CreateApplicationInstall**](CreateApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationInstallCreate

> DeleteApplicationInstallCreate(ctx, fileId).Execute()

Cancel a application installation process

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
	fileId := int32(56) // int32 | The Id of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationInstallCreateAPI.DeleteApplicationInstallCreate(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallCreateAPI.DeleteApplicationInstallCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The Id of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationInstallCreateRequest struct via the builder pattern


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


## GetApplicationInstallCreate

> []CreateApplicationInstall GetApplicationInstallCreate(ctx, fileId).Execute()

Get a create application install process

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
	fileId := int32(56) // int32 | The Id of the file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstallCreateAPI.GetApplicationInstallCreate(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallCreateAPI.GetApplicationInstallCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstallCreate`: []CreateApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallCreateAPI.GetApplicationInstallCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The Id of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CreateApplicationInstall**](CreateApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstallCreates

> []CreateApplicationInstall GetApplicationInstallCreates(ctx).Execute()

Get a list of create application install processes

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
	resp, r, err := apiClient.ApplicationInstallCreateAPI.GetApplicationInstallCreates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallCreateAPI.GetApplicationInstallCreates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstallCreates`: []CreateApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallCreateAPI.GetApplicationInstallCreates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallCreatesRequest struct via the builder pattern


### Return type

[**[]CreateApplicationInstall**](CreateApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

