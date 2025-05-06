# \ApplicationInstanceTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryApplicationInstanceCpus**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceCpus) | **Get** /v3/telemetry/applicationInstance/{applicationInstanceId}/cpu | Get cpu telemetry data for the given application instance (default &#x3D; last 24 hours)
[**GetTelemetryApplicationInstanceCrashes**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceCrashes) | **Get** /v3/telemetry/applicationInstance/crash | Count of crashes within a given time period for all application instances
[**GetTelemetryApplicationInstanceDeploymentProfileStatuses**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceDeploymentProfileStatuses) | **Get** /v3/telemetry/applicationInstance/deploymentProfile/status | Get live status of application instances grouped by deployment profile
[**GetTelemetryApplicationInstanceDistributionCurrents**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceDistributionCurrents) | **Get** /v3/telemetry/applicationInstance/distribution/current | Get the current distribution of application instances (of type Game) for bare metal, flex metal and cloud
[**GetTelemetryApplicationInstanceDistributions**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceDistributions) | **Get** /v3/telemetry/applicationInstance/distribution | Get the historic distribution of application instances (of type Game) for bare metal, flex metal and cloud
[**GetTelemetryApplicationInstanceHostCapacityCurrents**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceHostCapacityCurrents) | **Get** /v3/telemetry/applicationInstance/hostCapacity/current | Get the aggregated list of application instances (of type Game) for bare metal, flex metal and cloud
[**GetTelemetryApplicationInstanceMemories**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceMemories) | **Get** /v3/telemetry/applicationInstance/{applicationInstanceId}/memory | Get memory telemetry data for the given application instance (default &#x3D; last 24 hours)
[**GetTelemetryApplicationInstanceProviderIdAggregates**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceProviderIdAggregates) | **Get** /v3/telemetry/applicationInstance/providerId/aggregate | Get the aggregated list of application instance by providerId
[**GetTelemetryApplicationInstanceProviderIds**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceProviderIds) | **Get** /v3/telemetry/applicationInstance/providerId | Get how many application instances are currently deployed by providerId
[**GetTelemetryApplicationInstanceStatusAggregates**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceStatusAggregates) | **Get** /v3/telemetry/applicationInstance/status/aggregate | Get the aggregated list of application instances status
[**GetTelemetryApplicationInstanceStatuses**](ApplicationInstanceTelemetryAPI.md#GetTelemetryApplicationInstanceStatuses) | **Get** /v3/telemetry/applicationInstance/status | Get live status of application instances



## GetTelemetryApplicationInstanceCpus

> []ApplicationInstanceCpuTelemetry GetTelemetryApplicationInstanceCpus(ctx, applicationInstanceId).StartTime(startTime).EndTime(endTime).Execute()

Get cpu telemetry data for the given application instance (default = last 24 hours)

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
	applicationInstanceId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCpus(context.Background(), applicationInstanceId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceCpus`: []ApplicationInstanceCpuTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCpus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceCpusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]ApplicationInstanceCpuTelemetry**](ApplicationInstanceCpuTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceCrashes

> []ApplicationInstanceCrashTelemetry GetTelemetryApplicationInstanceCrashes(ctx).StartTime(startTime).EndTime(endTime).Execute()

Count of crashes within a given time period for all application instances

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
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCrashes(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCrashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceCrashes`: []ApplicationInstanceCrashTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceCrashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceCrashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]ApplicationInstanceCrashTelemetry**](ApplicationInstanceCrashTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceDeploymentProfileStatuses

> []ApplicationInstanceDeploymentProfileStatusTelemetry GetTelemetryApplicationInstanceDeploymentProfileStatuses(ctx).Type_(type_).Filter(filter).Execute()

Get live status of application instances grouped by deployment profile

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
	type_ := "type__example" // string | - 0: Both game and utilities<br /> - 1: Only game (default)<br /> - 2: Only utilities (optional)
	filter := "filter_example" // string | You can filter based on the following parameters<br /> - deploymentProfileId<br /> <p>The parameters can have a single value (e.g: deploymentProfileId=1734197255512241679) or a list of values (e.g: deploymentProfileId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentProfileId=8881688742569018542,76243584248556249<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentProfileId%3D8881688742569018542%2C76243584248556249<br /> The final endpoint call would look like: /telemetry/applicationInstance/status?filter=deploymentProfileId%3D8881688742569018542%2C76243584248556249</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDeploymentProfileStatuses(context.Background()).Type_(type_).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDeploymentProfileStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceDeploymentProfileStatuses`: []ApplicationInstanceDeploymentProfileStatusTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDeploymentProfileStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceDeploymentProfileStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | - 0: Both game and utilities&lt;br /&gt; - 1: Only game (default)&lt;br /&gt; - 2: Only utilities | 
 **filter** | **string** | You can filter based on the following parameters&lt;br /&gt; - deploymentProfileId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: deploymentProfileId&#x3D;1734197255512241679) or a list of values (e.g: deploymentProfileId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentProfileId&#x3D;8881688742569018542,76243584248556249&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentProfileId%3D8881688742569018542%2C76243584248556249&lt;br /&gt; The final endpoint call would look like: /telemetry/applicationInstance/status?filter&#x3D;deploymentProfileId%3D8881688742569018542%2C76243584248556249&lt;/p&gt; | 

### Return type

[**[]ApplicationInstanceDeploymentProfileStatusTelemetry**](ApplicationInstanceDeploymentProfileStatusTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceDistributionCurrents

> []InstanceDistribution GetTelemetryApplicationInstanceDistributionCurrents(ctx).Execute()

Get the current distribution of application instances (of type Game) for bare metal, flex metal and cloud

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
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributionCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributionCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceDistributionCurrents`: []InstanceDistribution
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributionCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceDistributionCurrentsRequest struct via the builder pattern


### Return type

[**[]InstanceDistribution**](InstanceDistribution.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceDistributions

> []InstanceDistributionTelemetry GetTelemetryApplicationInstanceDistributions(ctx).StartTime(startTime).EndTime(endTime).Execute()

Get the historic distribution of application instances (of type Game) for bare metal, flex metal and cloud



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
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributions(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceDistributions`: []InstanceDistributionTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceDistributions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceDistributionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]InstanceDistributionTelemetry**](InstanceDistributionTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceHostCapacityCurrents

> []ApplicationInstanceHostCapacityTelemetry GetTelemetryApplicationInstanceHostCapacityCurrents(ctx).Execute()

Get the aggregated list of application instances (of type Game) for bare metal, flex metal and cloud

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
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceHostCapacityCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceHostCapacityCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceHostCapacityCurrents`: []ApplicationInstanceHostCapacityTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceHostCapacityCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceHostCapacityCurrentsRequest struct via the builder pattern


### Return type

[**[]ApplicationInstanceHostCapacityTelemetry**](ApplicationInstanceHostCapacityTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceMemories

> []ApplicationInstanceMemoryTelemetry GetTelemetryApplicationInstanceMemories(ctx, applicationInstanceId).StartTime(startTime).EndTime(endTime).Execute()

Get memory telemetry data for the given application instance (default = last 24 hours)

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
	applicationInstanceId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceMemories(context.Background(), applicationInstanceId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceMemories`: []ApplicationInstanceMemoryTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceMemories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceMemoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]ApplicationInstanceMemoryTelemetry**](ApplicationInstanceMemoryTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceProviderIdAggregates

> []ApplicationInstancesTelemetry GetTelemetryApplicationInstanceProviderIdAggregates(ctx).Type_(type_).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

Get the aggregated list of application instance by providerId



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
	type_ := "type__example" // string | - 0: Both game and utilities<br /> - 1: Only game (default)<br /> - 2: Only utilities (optional)
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter.<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/applicationInstance/providerId/aggregate?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIdAggregates(context.Background()).Type_(type_).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIdAggregates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceProviderIdAggregates`: []ApplicationInstancesTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIdAggregates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceProviderIdAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | - 0: Both game and utilities&lt;br /&gt; - 1: Only game (default)&lt;br /&gt; - 2: Only utilities | 
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter.&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/applicationInstance/providerId/aggregate?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]ApplicationInstancesTelemetry**](ApplicationInstancesTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceProviderIds

> []ApplicationInstancesTelemetry GetTelemetryApplicationInstanceProviderIds(ctx).Filter(filter).Execute()

Get how many application instances are currently deployed by providerId

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/applicationInstance/aggregate?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIds(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceProviderIds`: []ApplicationInstancesTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceProviderIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceProviderIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/applicationInstance/aggregate?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 

### Return type

[**[]ApplicationInstancesTelemetry**](ApplicationInstancesTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceStatusAggregates

> []ApplicationInstanceStatusTelemetry GetTelemetryApplicationInstanceStatusAggregates(ctx).Type_(type_).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

Get the aggregated list of application instances status



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
	type_ := "type__example" // string | - 0: Both game and utilities<br /> - 1: Only game (default)<br /> - 2: Only utilities (optional)
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - regionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: regionId=8881688742569018542,76243584248556249)<br /> A filter example is: regionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter.<br /> E.g filter=regionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/applicationInstance/status/aggregate?filter=regionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatusAggregates(context.Background()).Type_(type_).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatusAggregates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceStatusAggregates`: []ApplicationInstanceStatusTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatusAggregates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceStatusAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | - 0: Both game and utilities&lt;br /&gt; - 1: Only game (default)&lt;br /&gt; - 2: Only utilities | 
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - regionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: regionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: regionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter.&lt;br /&gt; E.g filter&#x3D;regionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/applicationInstance/status/aggregate?filter&#x3D;regionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]ApplicationInstanceStatusTelemetry**](ApplicationInstanceStatusTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryApplicationInstanceStatuses

> []ApplicationInstanceStatusTelemetry GetTelemetryApplicationInstanceStatuses(ctx).Type_(type_).Filter(filter).Execute()

Get live status of application instances

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
	type_ := "type__example" // string | - 0: Both game and utilities<br /> - 1: Only game (default)<br /> - 2: Only utilities (optional)
	filter := "filter_example" // string | You can filter based on the following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/applicationInstance/status?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatuses(context.Background()).Type_(type_).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryApplicationInstanceStatuses`: []ApplicationInstanceStatusTelemetry
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTelemetryAPI.GetTelemetryApplicationInstanceStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryApplicationInstanceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | - 0: Both game and utilities&lt;br /&gt; - 1: Only game (default)&lt;br /&gt; - 2: Only utilities | 
 **filter** | **string** | You can filter based on the following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/applicationInstance/status?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 

### Return type

[**[]ApplicationInstanceStatusTelemetry**](ApplicationInstanceStatusTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

