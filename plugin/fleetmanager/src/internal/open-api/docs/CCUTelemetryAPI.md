# \CCUTelemetryAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTelemetryCcuApplications**](CCUTelemetryAPI.md#GetTelemetryCcuApplications) | **Get** /v3/telemetry/ccu/application | Count of players for an application
[**GetTelemetryCcuStatusAggregates**](CCUTelemetryAPI.md#GetTelemetryCcuStatusAggregates) | **Get** /v3/telemetry/ccu/status/aggregate | Get the telemetry of aggregated average of concurrent connected users
[**GetTelemetryCcuStatuses**](CCUTelemetryAPI.md#GetTelemetryCcuStatuses) | **Get** /v3/telemetry/ccu/status | Get the concurrent connected users telemetry for the current point in time



## GetTelemetryCcuApplications

> []CCUApplicationTelemetry GetTelemetryCcuApplications(ctx).Execute()

Count of players for an application

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
	resp, r, err := apiClient.CCUTelemetryAPI.GetTelemetryCcuApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCUTelemetryAPI.GetTelemetryCcuApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryCcuApplications`: []CCUApplicationTelemetry
	fmt.Fprintf(os.Stdout, "Response from `CCUTelemetryAPI.GetTelemetryCcuApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryCcuApplicationsRequest struct via the builder pattern


### Return type

[**[]CCUApplicationTelemetry**](CCUApplicationTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryCcuStatusAggregates

> []CCUStatusAggregateTelemetry GetTelemetryCcuStatusAggregates(ctx).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

Get the telemetry of aggregated average of concurrent connected users

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/ccu/status/aggregate?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)
	startTime := int32(56) // int32 | Unix timestamp. Filter from time, default value is -1 day (optional)
	endTime := int32(56) // int32 | Unix timestamp. Filter to time, default value is now (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCUTelemetryAPI.GetTelemetryCcuStatusAggregates(context.Background()).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCUTelemetryAPI.GetTelemetryCcuStatusAggregates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryCcuStatusAggregates`: []CCUStatusAggregateTelemetry
	fmt.Fprintf(os.Stdout, "Response from `CCUTelemetryAPI.GetTelemetryCcuStatusAggregates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryCcuStatusAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/ccu/status/aggregate?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 
 **startTime** | **int32** | Unix timestamp. Filter from time, default value is -1 day | 
 **endTime** | **int32** | Unix timestamp. Filter to time, default value is now | 

### Return type

[**[]CCUStatusAggregateTelemetry**](CCUStatusAggregateTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryCcuStatuses

> []CCUStatusTelemetry GetTelemetryCcuStatuses(ctx).Filter(filter).Execute()

Get the concurrent connected users telemetry for the current point in time

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - deploymentRegionId<br /> - dcLocationId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: deploymentRegionId=8881688742569018542,76243584248556249)<br /> A filter example is: deploymentRegionId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /telemetry/ccu/status?filter=deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCUTelemetryAPI.GetTelemetryCcuStatuses(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCUTelemetryAPI.GetTelemetryCcuStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryCcuStatuses`: []CCUStatusTelemetry
	fmt.Fprintf(os.Stdout, "Response from `CCUTelemetryAPI.GetTelemetryCcuStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryCcuStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - deploymentRegionId&lt;br /&gt; - dcLocationId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: deploymentRegionId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: deploymentRegionId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /telemetry/ccu/status?filter&#x3D;deploymentRegionId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 

### Return type

[**[]CCUStatusTelemetry**](CCUStatusTelemetry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

