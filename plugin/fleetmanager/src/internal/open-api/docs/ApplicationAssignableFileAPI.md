# \ApplicationAssignableFileAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationInstallAssignable**](ApplicationAssignableFileAPI.md#GetApplicationInstallAssignable) | **Get** /v3/applicationInstall/assignable/{fileId} | Get details for the given assignable file
[**GetApplicationInstallAssignables**](ApplicationAssignableFileAPI.md#GetApplicationInstallAssignables) | **Get** /v3/applicationInstall/assignable | Get a list of assignable files



## GetApplicationInstallAssignable

> []AssignableFile GetApplicationInstallAssignable(ctx, fileId).Execute()

Get details for the given assignable file

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
	fileId := int32(56) // int32 | The Id of the assignable file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAssignableFileAPI.GetApplicationInstallAssignable(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAssignableFileAPI.GetApplicationInstallAssignable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstallAssignable`: []AssignableFile
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAssignableFileAPI.GetApplicationInstallAssignable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The Id of the assignable file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallAssignableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AssignableFile**](AssignableFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstallAssignables

> []AssignableFile GetApplicationInstallAssignables(ctx).Execute()

Get a list of assignable files

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
	resp, r, err := apiClient.ApplicationAssignableFileAPI.GetApplicationInstallAssignables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAssignableFileAPI.GetApplicationInstallAssignables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstallAssignables`: []AssignableFile
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAssignableFileAPI.GetApplicationInstallAssignables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallAssignablesRequest struct via the builder pattern


### Return type

[**[]AssignableFile**](AssignableFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

