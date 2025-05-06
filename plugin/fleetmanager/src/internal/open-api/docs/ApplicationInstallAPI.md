# \ApplicationInstallAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationInstall**](ApplicationInstallAPI.md#DeleteApplicationInstall) | **Delete** /v3/applicationInstall/{installId} | Remove (deactivate) an application install
[**GetApplicationInstall**](ApplicationInstallAPI.md#GetApplicationInstall) | **Get** /v3/applicationInstall/{installId} | Get an application install
[**GetApplicationInstalls**](ApplicationInstallAPI.md#GetApplicationInstalls) | **Get** /v3/applicationInstall | Get all your application installs
[**UpdateApplicationInstall**](ApplicationInstallAPI.md#UpdateApplicationInstall) | **Put** /v3/applicationInstall/{installId} | Update an application install



## DeleteApplicationInstall

> DeleteApplicationInstall(ctx, installId).Execute()

Remove (deactivate) an application install

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
	installId := int32(56) // int32 | The Id of the application install you want to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationInstallAPI.DeleteApplicationInstall(context.Background(), installId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallAPI.DeleteApplicationInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installId** | **int32** | The Id of the application install you want to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationInstallRequest struct via the builder pattern


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


## GetApplicationInstall

> []ApplicationInstall GetApplicationInstall(ctx, installId).Execute()

Get an application install

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
	installId := int32(56) // int32 | The Id of the application install to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstallAPI.GetApplicationInstall(context.Background(), installId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallAPI.GetApplicationInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstall`: []ApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallAPI.GetApplicationInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installId** | **int32** | The Id of the application install to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationInstall**](ApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstalls

> []ApplicationInstall GetApplicationInstalls(ctx).Execute()

Get all your application installs

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
	resp, r, err := apiClient.ApplicationInstallAPI.GetApplicationInstalls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallAPI.GetApplicationInstalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstalls`: []ApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallAPI.GetApplicationInstalls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallsRequest struct via the builder pattern


### Return type

[**[]ApplicationInstall**](ApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationInstall

> []ApplicationInstall UpdateApplicationInstall(ctx, installId).ApplicationInstall(applicationInstall).Execute()

Update an application install

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
	installId := int32(56) // int32 | The Id of the application install to update
	applicationInstall := *openapiclient.NewApplicationInstall(int32(123), "ApplicationId_example", NullableInt32(123), int32(123), "Name_example", "Version_example", int32(123), int32(123), "FilePath_example", int32(123), "FileSHA256_example", "ExecutableSHA256_example", int32(123), int32(123)) // ApplicationInstall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstallAPI.UpdateApplicationInstall(context.Background(), installId).ApplicationInstall(applicationInstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstallAPI.UpdateApplicationInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationInstall`: []ApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstallAPI.UpdateApplicationInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installId** | **int32** | The Id of the application install to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationInstall** | [**ApplicationInstall**](ApplicationInstall.md) |  | 

### Return type

[**[]ApplicationInstall**](ApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

