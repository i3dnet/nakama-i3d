# \ApplicationTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryApplication**](ApplicationTelemetryAPI.md#GetTelemetryApplication) | **Get** /v3/telemetry/application/{applicationId} | Get all your application telemetry history
[**GetTelemetryApplicationCurrents**](ApplicationTelemetryAPI.md#GetTelemetryApplicationCurrents) | **Get** /v3/telemetry/application/{applicationId}/current | Get all your application telemetry



## GetTelemetryApplication

> []ApplicationTelemetryHistoryModel GetTelemetryApplication(ctx).StartTime(startTime).EndTime(endTime).Execute()

Get all your application telemetry history

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
	resp, r, err := apiClient.ApplicationTelemetryAPI.GetTelemetryApplication(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTelemetryAPI.GetTelemetryApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplication`: []ApplicationTelemetryHistoryModel
	fmt.Fprintf(os.Stdout, "Response from `ApplicationTelemetryAPI.GetTelemetryApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]ApplicationTelemetryHistoryModel**](ApplicationTelemetryHistoryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationCurrents

> []ApplicationTelemetryModel GetTelemetryApplicationCurrents(ctx).Execute()

Get all your application telemetry

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
	resp, r, err := apiClient.ApplicationTelemetryAPI.GetTelemetryApplicationCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTelemetryAPI.GetTelemetryApplicationCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationCurrents`: []ApplicationTelemetryModel
	fmt.Fprintf(os.Stdout, "Response from `ApplicationTelemetryAPI.GetTelemetryApplicationCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationCurrentsRequest struct via the builder pattern


### Return type

[**[]ApplicationTelemetryModel**](ApplicationTelemetryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

