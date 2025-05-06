# \HostTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryHostCpus**](HostTelemetryAPI.md#GetTelemetryHostCpus) | **Get** /v3/telemetry/host/{hostId}/cpu | Get cpu telemetry for the given host (default &#x3D; last 24 hours)
[**GetTelemetryHostHardwareCapacities**](HostTelemetryAPI.md#GetTelemetryHostHardwareCapacities) | **Get** /v3/telemetry/host/hardwareCapacity | List the historical data of hardware capacity of all your Hosts
[**GetTelemetryHostHardwareCapacityCurrents**](HostTelemetryAPI.md#GetTelemetryHostHardwareCapacityCurrents) | **Get** /v3/telemetry/host/hardwareCapacity/current | List the current hardware capacity of all your Hosts
[**GetTelemetryHostHdds**](HostTelemetryAPI.md#GetTelemetryHostHdds) | **Get** /v3/telemetry/host/{hostId}/hdd | Get hard disk telemetry for the given host (default &#x3D; last 24 hours)
[**GetTelemetryHostMemories**](HostTelemetryAPI.md#GetTelemetryHostMemories) | **Get** /v3/telemetry/host/{hostId}/memory | Get memory telemetry for the given host (default &#x3D; last 24 hours)
[**GetTelemetryHostNetworks**](HostTelemetryAPI.md#GetTelemetryHostNetworks) | **Get** /v3/telemetry/host/{hostId}/network | Get network telemetry for the given host (default &#x3D; last 24 hours)
[**GetTelemetryHostOnlineStatusCurrents**](HostTelemetryAPI.md#GetTelemetryHostOnlineStatusCurrents) | **Get** /v3/telemetry/host/onlineStatus/current | List the statuses of currently available Hosts
[**GetTelemetryHostOnlineStatuses**](HostTelemetryAPI.md#GetTelemetryHostOnlineStatuses) | **Get** /v3/telemetry/host/onlineStatus | List the historical data of status of all your Hosts
[**GetTelemetryHostOveruseCurrents**](HostTelemetryAPI.md#GetTelemetryHostOveruseCurrents) | **Get** /v3/telemetry/host/overuse/current | List of hosts that have traffic overuse
[**GetTelemetryHostUtilizationAggregates**](HostTelemetryAPI.md#GetTelemetryHostUtilizationAggregates) | **Get** /v3/telemetry/host/utilization/aggregate | Get the average of server host status telemetry
[**GetTelemetryHostUtilizations**](HostTelemetryAPI.md#GetTelemetryHostUtilizations) | **Get** /v3/telemetry/host/utilization | Get the telemetry for all hosts across all providers at the current point in time



## GetTelemetryHostCpus

> []HostCpuTelemetry GetTelemetryHostCpus(ctx, hostId).StartTime(startTime).EndTime(endTime).Execute()

Get cpu telemetry for the given host (default = last 24 hours)

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
	hostId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostCpus(context.Background(), hostId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostCpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostCpus`: []HostCpuTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostCpus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostCpusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]HostCpuTelemetry**](HostCpuTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostHardwareCapacities

> []HostHardwareCapacityTelemetry GetTelemetryHostHardwareCapacities(ctx).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

List the historical data of hardware capacity of all your Hosts

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: fleetId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentEnvironmentId=8881688742569018542&fleetId=1734197255512241679,76243584248556249<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249<br /> The final endpoint call would look like: /telemetry/host/hardwareCapacity?filter=deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249</p> (optional)
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostHardwareCapacities(context.Background()).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostHardwareCapacities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostHardwareCapacities`: []HostHardwareCapacityTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostHardwareCapacities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostHardwareCapacitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: fleetId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentEnvironmentId&#x3D;8881688742569018542&amp;fleetId&#x3D;1734197255512241679,76243584248556249&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249&lt;br /&gt; The final endpoint call would look like: /telemetry/host/hardwareCapacity?filter&#x3D;deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249&lt;/p&gt; | 
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]HostHardwareCapacityTelemetry**](HostHardwareCapacityTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostHardwareCapacityCurrents

> []HostHardwareCapacityDataValueTelemetry GetTelemetryHostHardwareCapacityCurrents(ctx).Filter(filter).Execute()

List the current hardware capacity of all your Hosts

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: fleetId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentEnvironmentId=8881688742569018542&fleetId=1734197255512241679,76243584248556249<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249<br /> The final endpoint call would look like: /telemetry/host/hardwareCapacity?filter=deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostHardwareCapacityCurrents(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostHardwareCapacityCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostHardwareCapacityCurrents`: []HostHardwareCapacityDataValueTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostHardwareCapacityCurrents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostHardwareCapacityCurrentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: fleetId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentEnvironmentId&#x3D;8881688742569018542&amp;fleetId&#x3D;1734197255512241679,76243584248556249&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249&lt;br /&gt; The final endpoint call would look like: /telemetry/host/hardwareCapacity?filter&#x3D;deploymentEnvironmentId%3D8881688742569018542%26fleetId%3D1734197255512241679%2C76243584248556249&lt;/p&gt; | 

### Return type

[**[]HostHardwareCapacityDataValueTelemetry**](HostHardwareCapacityDataValueTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostHdds

> []HostHddTelemetry GetTelemetryHostHdds(ctx, hostId).StartTime(startTime).EndTime(endTime).Execute()

Get hard disk telemetry for the given host (default = last 24 hours)

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
	hostId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostHdds(context.Background(), hostId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostHdds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostHdds`: []HostHddTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostHdds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostHddsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]HostHddTelemetry**](HostHddTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostMemories

> []HostMemoryTelemetry GetTelemetryHostMemories(ctx, hostId).StartTime(startTime).EndTime(endTime).Execute()

Get memory telemetry for the given host (default = last 24 hours)

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
	hostId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostMemories(context.Background(), hostId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostMemories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostMemories`: []HostMemoryTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostMemories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostMemoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]HostMemoryTelemetry**](HostMemoryTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostNetworks

> []HostNetworkTelemetry GetTelemetryHostNetworks(ctx, hostId).StartTime(startTime).EndTime(endTime).Execute()

Get network telemetry for the given host (default = last 24 hours)

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
	hostId := int32(56) // int32 | 
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostNetworks(context.Background(), hostId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostNetworks`: []HostNetworkTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]HostNetworkTelemetry**](HostNetworkTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostOnlineStatusCurrents

> []HostOnlineStatusTelemetry GetTelemetryHostOnlineStatusCurrents(ctx).Execute()

List the statuses of currently available Hosts

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
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostOnlineStatusCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostOnlineStatusCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostOnlineStatusCurrents`: []HostOnlineStatusTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostOnlineStatusCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostOnlineStatusCurrentsRequest struct via the builder pattern


### Return type

[**[]HostOnlineStatusTelemetry**](HostOnlineStatusTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostOnlineStatuses

> []HostOnlineStatusHistoryTelemetry GetTelemetryHostOnlineStatuses(ctx).StartTime(startTime).EndTime(endTime).Execute()

List the historical data of status of all your Hosts

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
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostOnlineStatuses(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostOnlineStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostOnlineStatuses`: []HostOnlineStatusHistoryTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostOnlineStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostOnlineStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]HostOnlineStatusHistoryTelemetry**](HostOnlineStatusHistoryTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostOveruseCurrents

> []HostOveruseTelemetry GetTelemetryHostOveruseCurrents(ctx).RANGEDDATA(rANGEDDATA).Execute()

List of hosts that have traffic overuse

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
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostOveruseCurrents(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostOveruseCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostOveruseCurrents`: []HostOveruseTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostOveruseCurrents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostOveruseCurrentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]HostOveruseTelemetry**](HostOveruseTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostUtilizationAggregates

> HostUtilizationAggregateTelemetry GetTelemetryHostUtilizationAggregates(ctx).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

Get the average of server host status telemetry

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/host/status/aggregate?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostUtilizationAggregates(context.Background()).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostUtilizationAggregates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostUtilizationAggregates`: HostUtilizationAggregateTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostUtilizationAggregates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostUtilizationAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/host/status/aggregate?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**HostUtilizationAggregateTelemetry**](HostUtilizationAggregateTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryHostUtilizations

> []HostUtilizationTelemetry GetTelemetryHostUtilizations(ctx).Filter(filter).Execute()

Get the telemetry for all hosts across all providers at the current point in time

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/host/status?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostTelemetryAPI.GetTelemetryHostUtilizations(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostTelemetryAPI.GetTelemetryHostUtilizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryHostUtilizations`: []HostUtilizationTelemetry
	fmt.Fprintf(os.Stdout, "Response from `HostTelemetryAPI.GetTelemetryHostUtilizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryHostUtilizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/host/status?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 

### Return type

[**[]HostUtilizationTelemetry**](HostUtilizationTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

