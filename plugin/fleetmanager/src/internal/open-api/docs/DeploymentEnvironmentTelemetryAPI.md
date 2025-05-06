# \DeploymentEnvironmentTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryDeploymentEnvironmentCurrents**](DeploymentEnvironmentTelemetryAPI.md#GetTelemetryDeploymentEnvironmentCurrents) | **Get** /v3/telemetry/deploymentEnvironment/{deploymentEnvironmentId}/current | Get current deployment environment telemetry



## GetTelemetryDeploymentEnvironmentCurrents

> []DeploymentEnvironmentTelemetryModel GetTelemetryDeploymentEnvironmentCurrents(ctx).Execute()

Get current deployment environment telemetry

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
	resp, r, err := apiClient.DeploymentEnvironmentTelemetryAPI.GetTelemetryDeploymentEnvironmentCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentTelemetryAPI.GetTelemetryDeploymentEnvironmentCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryDeploymentEnvironmentCurrents`: []DeploymentEnvironmentTelemetryModel
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentTelemetryAPI.GetTelemetryDeploymentEnvironmentCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryDeploymentEnvironmentCurrentsRequest struct via the builder pattern


### Return type

[**[]DeploymentEnvironmentTelemetryModel**](DeploymentEnvironmentTelemetryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

