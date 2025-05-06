# \ApplicationBuildTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryApplicationBuild**](ApplicationBuildTelemetryAPI.md#GetTelemetryApplicationBuild) | **Get** /v3/telemetry/applicationBuild/{applicationBuildId} | Get all your application build telemetry history
[**GetTelemetryApplicationBuildCurrents**](ApplicationBuildTelemetryAPI.md#GetTelemetryApplicationBuildCurrents) | **Get** /v3/telemetry/applicationBuild/{applicationBuildId}/current | Get all your application build telemetry



## GetTelemetryApplicationBuild

> []ApplicationBuildTelemetryHistoryModel GetTelemetryApplicationBuild(ctx).StartTime(startTime).EndTime(endTime).Execute()

Get all your application build telemetry history

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
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuild(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationBuild`: []ApplicationBuildTelemetryHistoryModel
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]ApplicationBuildTelemetryHistoryModel**](ApplicationBuildTelemetryHistoryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationBuildCurrents

> []ApplicationBuildTelemetryModel GetTelemetryApplicationBuildCurrents(ctx).Execute()

Get all your application build telemetry

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
	resp, r, err := apiClient.ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuildCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuildCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationBuildCurrents`: []ApplicationBuildTelemetryModel
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildTelemetryAPI.GetTelemetryApplicationBuildCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationBuildCurrentsRequest struct via the builder pattern


### Return type

[**[]ApplicationBuildTelemetryModel**](ApplicationBuildTelemetryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

