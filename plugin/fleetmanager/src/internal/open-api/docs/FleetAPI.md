# \FleetAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFleet**](FleetAPI.md#CreateFleet) | **Post** /v3/fleet | Create a new fleet
[**CreateFleetHostBulkReserve**](FleetAPI.md#CreateFleetHostBulkReserve) | **Post** /v3/fleet/host/bulkReserve | Reserve a list of hosts for a specific new fleet. The existing application instances on these servers will be automatically scaled down.
[**DeleteFleet**](FleetAPI.md#DeleteFleet) | **Delete** /v3/fleet/{fleetId} | Removes an existing fleet. Note that it may not have any active servers assigned to it.
[**GetFleet**](FleetAPI.md#GetFleet) | **Get** /v3/fleet/{fleetId} | Get a specific fleet by Id
[**GetFleetApplicationBuildInUses**](FleetAPI.md#GetFleetApplicationBuildInUses) | **Get** /v3/fleet/{fleetId}/applicationBuild/inUse | Get the list of application builds for the given fleet
[**GetFleetApplicationInUses**](FleetAPI.md#GetFleetApplicationInUses) | **Get** /v3/fleet/{fleetId}/application/inUse | Get the list of applications for the given fleet
[**GetFleetGameLiftDeploymentEnvironmentVariables**](FleetAPI.md#GetFleetGameLiftDeploymentEnvironmentVariables) | **Get** /v3/fleet/{fleetId}/gameLiftDeploymentEnvironmentVariables | Get a list of environment variables for a fleet
[**GetFleetHostBulkReserveStatuses**](FleetAPI.md#GetFleetHostBulkReserveStatuses) | **Get** /v3/fleet/host/bulkReserve/{traceId}/status | Progress for the bulk reserve you have done, base on the trace id you received in the header when you have created the bulk reserve
[**GetFleetHostReserveds**](FleetAPI.md#GetFleetHostReserveds) | **Get** /v3/fleet/{fleetId}/host/reserved | Get all hosts reserved for the given fleet
[**GetFleetHosts**](FleetAPI.md#GetFleetHosts) | **Get** /v3/fleet/{fleetId}/host | Get all your hosts for a fleet.
[**GetFleets**](FleetAPI.md#GetFleets) | **Get** /v3/fleet | Get a list of fleets
[**GetTelemetryFleetCurrents**](FleetAPI.md#GetTelemetryFleetCurrents) | **Get** /v3/telemetry/fleet/{fleetId}/current | Get current fleet telemetry
[**UpdateFleet**](FleetAPI.md#UpdateFleet) | **Put** /v3/fleet/{fleetId} | Update given fleet
[**UpdateFleetGameLiftDeploymentEnvironmentVariables**](FleetAPI.md#UpdateFleetGameLiftDeploymentEnvironmentVariables) | **Put** /v3/fleet/{fleetId}/gameLiftDeploymentEnvironmentVariables | Add or update environment variables for a fleet
[**UpdateFleetHostReserve**](FleetAPI.md#UpdateFleetHostReserve) | **Put** /v3/fleet/{fleetId}/host/{hostId}/reserve | Reserve the given host for the given fleet Id
[**UpdateFleetHostUnreserve**](FleetAPI.md#UpdateFleetHostUnreserve) | **Put** /v3/fleet/{fleetId}/host/{hostId}/unreserve | Remove a fleet reservation from the given host. Sets &#x60;fleetId&#x60; to &#x60;0&#x60; only if no application instances are deployed on it.
[**UpdateFleetOperationalStatus**](FleetAPI.md#UpdateFleetOperationalStatus) | **Put** /v3/fleet/{fleetId}/operationalStatus | Update the operational status of the given fleet



## CreateFleet

> []Fleet CreateFleet(ctx).Fleet(fleet).Execute()

Create a new fleet

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
	fleet := *openapiclient.NewFleet("Id_example", "Name_example", "DeploymentEnvironmentId_example", int32(123), false) // Fleet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateFleet(context.Background()).Fleet(fleet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFleet`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateFleet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleet** | [**Fleet**](Fleet.md) |  | 

### Return type

[**[]Fleet**](Fleet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleetHostBulkReserve

> CreateFleetHostBulkReserve(ctx).BulkReserveHostByFleet(bulkReserveHostByFleet).XTraceId(xTraceId).Execute()

Reserve a list of hosts for a specific new fleet. The existing application instances on these servers will be automatically scaled down.

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
	bulkReserveHostByFleet := *openapiclient.NewBulkReserveHostByFleet("FleetId_example", []int32{int32(123)}) // BulkReserveHostByFleet | 
	xTraceId := "xTraceId_example" // string | Response header which you can use to fetch the status of the bulk reservation of hosts by using endpoint<br /> [`GET /v3/fleet/host/bulkReserve/{traceId}/status`](#/Fleet/getFleetHostBulkReserveStatuses) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetAPI.CreateFleetHostBulkReserve(context.Background()).BulkReserveHostByFleet(bulkReserveHostByFleet).XTraceId(xTraceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateFleetHostBulkReserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFleetHostBulkReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkReserveHostByFleet** | [**BulkReserveHostByFleet**](BulkReserveHostByFleet.md) |  | 
 **xTraceId** | **string** | Response header which you can use to fetch the status of the bulk reservation of hosts by using endpoint&lt;br /&gt; [&#x60;GET /v3/fleet/host/bulkReserve/{traceId}/status&#x60;](#/Fleet/getFleetHostBulkReserveStatuses) | 

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


## DeleteFleet

> DeleteFleet(ctx, fleetId).Execute()

Removes an existing fleet. Note that it may not have any active servers assigned to it.

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
	fleetId := "fleetId_example" // string | The ID of the fleet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetAPI.DeleteFleet(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.DeleteFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The ID of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFleetRequest struct via the builder pattern


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


## GetFleet

> []Fleet GetFleet(ctx, fleetId).Execute()

Get a specific fleet by Id

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
	fleetId := "fleetId_example" // string | The Id of the fleet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleet(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleet`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Fleet**](Fleet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetApplicationBuildInUses

> []ApplicationBuild GetFleetApplicationBuildInUses(ctx, fleetId).Execute()

Get the list of application builds for the given fleet

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
	fleetId := "fleetId_example" // string | The Id of the fleet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleetApplicationBuildInUses(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetApplicationBuildInUses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetApplicationBuildInUses`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetApplicationBuildInUses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetApplicationBuildInUsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationBuild**](ApplicationBuild.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetApplicationInUses

> []Application GetFleetApplicationInUses(ctx, fleetId).Execute()

Get the list of applications for the given fleet

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
	fleetId := "fleetId_example" // string | The Id of the fleet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleetApplicationInUses(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetApplicationInUses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetApplicationInUses`: []Application
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetApplicationInUses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetApplicationInUsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Application**](Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetGameLiftDeploymentEnvironmentVariables

> []GameLiftEnvironmentVariables GetFleetGameLiftDeploymentEnvironmentVariables(ctx).Execute()

Get a list of environment variables for a fleet

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
	resp, r, err := apiClient.FleetAPI.GetFleetGameLiftDeploymentEnvironmentVariables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetGameLiftDeploymentEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetGameLiftDeploymentEnvironmentVariables`: []GameLiftEnvironmentVariables
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetGameLiftDeploymentEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetGameLiftDeploymentEnvironmentVariablesRequest struct via the builder pattern


### Return type

[**[]GameLiftEnvironmentVariables**](GameLiftEnvironmentVariables.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetHostBulkReserveStatuses

> []BulkReserveHostByFleetStatus GetFleetHostBulkReserveStatuses(ctx).Execute()

Progress for the bulk reserve you have done, base on the trace id you received in the header when you have created the bulk reserve

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
	resp, r, err := apiClient.FleetAPI.GetFleetHostBulkReserveStatuses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetHostBulkReserveStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetHostBulkReserveStatuses`: []BulkReserveHostByFleetStatus
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetHostBulkReserveStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetHostBulkReserveStatusesRequest struct via the builder pattern


### Return type

[**[]BulkReserveHostByFleetStatus**](BulkReserveHostByFleetStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetHostReserveds

> []Host GetFleetHostReserveds(ctx, fleetId).RANGEDDATA(rANGEDDATA).Execute()

Get all hosts reserved for the given fleet

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
	fleetId := "fleetId_example" // string | The Id of the fleet
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleetHostReserveds(context.Background(), fleetId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetHostReserveds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetHostReserveds`: []Host
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetHostReserveds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetHostReservedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]Host**](Host.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetHosts

> []Host GetFleetHosts(ctx, fleetId).RANGEDDATA(rANGEDDATA).Execute()

Get all your hosts for a fleet.

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
	fleetId := "fleetId_example" // string | The Id of the fleet
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleetHosts(context.Background(), fleetId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleetHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetHosts`: []Host
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleetHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]Host**](Host.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleets

> []Fleet GetFleets(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get a list of fleets

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
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetFleets(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetFleets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]Fleet**](Fleet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryFleetCurrents

> []FleetTelemetryModel GetTelemetryFleetCurrents(ctx).Execute()

Get current fleet telemetry

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
	resp, r, err := apiClient.FleetAPI.GetTelemetryFleetCurrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetTelemetryFleetCurrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryFleetCurrents`: []FleetTelemetryModel
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetTelemetryFleetCurrents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryFleetCurrentsRequest struct via the builder pattern


### Return type

[**[]FleetTelemetryModel**](FleetTelemetryModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> []Fleet UpdateFleet(ctx, fleetId).Fleet(fleet).Execute()

Update given fleet

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
	fleetId := "fleetId_example" // string | The Id of the fleet
	fleet := *openapiclient.NewFleet("Id_example", "Name_example", "DeploymentEnvironmentId_example", int32(123), false) // Fleet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.UpdateFleet(context.Background(), fleetId).Fleet(fleet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.UpdateFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFleet`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.UpdateFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleet** | [**Fleet**](Fleet.md) |  | 

### Return type

[**[]Fleet**](Fleet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleetGameLiftDeploymentEnvironmentVariables

> []GameLiftEnvironmentVariables UpdateFleetGameLiftDeploymentEnvironmentVariables(ctx).Execute()

Add or update environment variables for a fleet

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
	resp, r, err := apiClient.FleetAPI.UpdateFleetGameLiftDeploymentEnvironmentVariables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.UpdateFleetGameLiftDeploymentEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFleetGameLiftDeploymentEnvironmentVariables`: []GameLiftEnvironmentVariables
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.UpdateFleetGameLiftDeploymentEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetGameLiftDeploymentEnvironmentVariablesRequest struct via the builder pattern


### Return type

[**[]GameLiftEnvironmentVariables**](GameLiftEnvironmentVariables.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleetHostReserve

> UpdateFleetHostReserve(ctx, fleetId, hostId).Execute()

Reserve the given host for the given fleet Id

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
	fleetId := "fleetId_example" // string | The Id of the fleet
	hostId := int32(56) // int32 | The Id of the host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetAPI.UpdateFleetHostReserve(context.Background(), fleetId, hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.UpdateFleetHostReserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 
**hostId** | **int32** | The Id of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetHostReserveRequest struct via the builder pattern


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


## UpdateFleetHostUnreserve

> UpdateFleetHostUnreserve(ctx, fleetId, hostId).Execute()

Remove a fleet reservation from the given host. Sets `fleetId` to `0` only if no application instances are deployed on it.



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
	fleetId := "fleetId_example" // string | The ID of the fleet
	hostId := int32(56) // int32 | The ID of the host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetAPI.UpdateFleetHostUnreserve(context.Background(), fleetId, hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.UpdateFleetHostUnreserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The ID of the fleet | 
**hostId** | **int32** | The ID of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetHostUnreserveRequest struct via the builder pattern


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


## UpdateFleetOperationalStatus

> []Fleet UpdateFleetOperationalStatus(ctx, fleetId).OperationalStatus(operationalStatus).Execute()

Update the operational status of the given fleet

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
	fleetId := "fleetId_example" // string | The Id of the fleet
	operationalStatus := *openapiclient.NewOperationalStatus(int32(123)) // OperationalStatus | Options: - 0: Manual, no application instances will be deployed nor deleted - 1: Automatic deployment enabled, the initial deploy according to deploymentProfile will be activated - 2: Automatic scaling enabled, scaling based on demand will be activated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.UpdateFleetOperationalStatus(context.Background(), fleetId).OperationalStatus(operationalStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.UpdateFleetOperationalStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFleetOperationalStatus`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.UpdateFleetOperationalStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The Id of the fleet | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFleetOperationalStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operationalStatus** | [**OperationalStatus**](OperationalStatus.md) | Options: - 0: Manual, no application instances will be deployed nor deleted - 1: Automatic deployment enabled, the initial deploy according to deploymentProfile will be activated - 2: Automatic scaling enabled, scaling based on demand will be activated | 

### Return type

[**[]Fleet**](Fleet.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

