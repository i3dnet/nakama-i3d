# \HostAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelHost**](HostAPI.md#CancelHost) | **Put** /v3/host/{hostId}/cancel | Cancels a host. If you cancel a host you will continue to have access to the host until the end of the contract
[**CreateHostFormatDisk**](HostAPI.md#CreateHostFormatDisk) | **Post** /v3/host/{hostId}/formatDisks | Format all disks inside the host
[**CreateHostOsInstall**](HostAPI.md#CreateHostOsInstall) | **Post** /v3/host/{hostId}/os/install | Install an operating system on the host
[**GetHost**](HostAPI.md#GetHost) | **Get** /v3/host/{hostId} | Get a host (can be bare metal servers or VMs) and their main details
[**GetHostBmcStatuses**](HostAPI.md#GetHostBmcStatuses) | **Get** /v3/host/{hostId}/bmcStatus | Get the BMC power status of the host
[**GetHostInstanceTypes**](HostAPI.md#GetHostInstanceTypes) | **Get** /v3/host/instanceType | Get all bare metal instance types (and their details) available for your account
[**GetHostInvoices**](HostAPI.md#GetHostInvoices) | **Get** /v3/host/{hostId}/invoice | Get all the invoices for a specific host
[**GetHostLiveStates**](HostAPI.md#GetHostLiveStates) | **Get** /v3/host/liveState | Get live situation of all the gaming hosts
[**GetHostNetworkUsageGraphDays**](HostAPI.md#GetHostNetworkUsageGraphDays) | **Get** /v3/host/{hostId}/network/usage/{upLinkId}/graph/day | Get the traffic usage graph (last 24 hours)
[**GetHostNetworkUsageGraphMonths**](HostAPI.md#GetHostNetworkUsageGraphMonths) | **Get** /v3/host/{hostId}/network/usage/{upLinkId}/graph/month | Get the traffic usage graph (last 30 days)
[**GetHostNetworkUsageGraphWeeks**](HostAPI.md#GetHostNetworkUsageGraphWeeks) | **Get** /v3/host/{hostId}/network/usage/{upLinkId}/graph/week | Get the traffic usage graph (last 7 days)
[**GetHostNetworkUsageGraphYears**](HostAPI.md#GetHostNetworkUsageGraphYears) | **Get** /v3/host/{hostId}/network/usage/{upLinkId}/graph/year | Get the traffic usage graph (last year)
[**GetHostNetworkUsageRaw**](HostAPI.md#GetHostNetworkUsageRaw) | **Get** /v3/host/{hostId}/network/usage/{upLinkId}/raw | Get the traffic usage data for the last day
[**GetHostOsInstallLogs**](HostAPI.md#GetHostOsInstallLogs) | **Get** /v3/host/{hostId}/os/installLog | Get all OS installation logs on this host
[**GetHostStatuses**](HostAPI.md#GetHostStatuses) | **Get** /v3/host/{hostId}/status | Get the power status of the host
[**GetHostSummaries**](HostAPI.md#GetHostSummaries) | **Get** /v3/host/summary | Get all your hosts (can be bare metal servers or VMs) and their main details
[**GetHostTasks**](HostAPI.md#GetHostTasks) | **Get** /v3/host/{hostId}/task | Get the last 50 tasks of the given host
[**GetHosts**](HostAPI.md#GetHosts) | **Get** /v3/host | Get all your hosts (can be bare metal servers or VMs) and their main details
[**RevertHostCancellation**](HostAPI.md#RevertHostCancellation) | **Put** /v3/host/{hostId}/cancel/revert | Reverts the cancellation of a host. You cannot revert the cancellation a host if the current date is equal to or past the end of contract date, nor if the host was never cancelled
[**UpdateHost**](HostAPI.md#UpdateHost) | **Put** /v3/host/{hostId} | Can add, edit or delete host labels. To delete label need to provide key with null value.
[**UpdateHostRestart**](HostAPI.md#UpdateHostRestart) | **Put** /v3/host/{hostId}/restart | Restart the host
[**UpdateHostStart**](HostAPI.md#UpdateHostStart) | **Put** /v3/host/{hostId}/start | Turn the host On
[**UpdateHostStop**](HostAPI.md#UpdateHostStop) | **Put** /v3/host/{hostId}/stop | Turn the host Off



## CancelHost

> Host CancelHost(ctx, hostId).Execute()

Cancels a host. If you cancel a host you will continue to have access to the host until the end of the contract

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
	hostId := int32(56) // int32 | ID of the host to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.CancelHost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.CancelHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelHost`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.CancelHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Host**](Host.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostFormatDisk

> []HostFormatDisks CreateHostFormatDisk(ctx, hostId).HostFormatDisks(hostFormatDisks).Execute()

Format all disks inside the host

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
	hostId := int32(56) // int32 | ID of the host to format disks for
	hostFormatDisks := *openapiclient.NewHostFormatDisks("Method_example") // HostFormatDisks | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.CreateHostFormatDisk(context.Background(), hostId).HostFormatDisks(hostFormatDisks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.CreateHostFormatDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostFormatDisk`: []HostFormatDisks
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.CreateHostFormatDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to format disks for | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostFormatDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostFormatDisks** | [**HostFormatDisks**](HostFormatDisks.md) |  | 

### Return type

[**[]HostFormatDisks**](HostFormatDisks.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostOsInstall

> []HostOsInstall CreateHostOsInstall(ctx, hostId).HostOsInstall(hostOsInstall).Execute()

Install an operating system on the host

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
	hostId := int32(56) // int32 | ID of the host to install
	hostOsInstall := *openapiclient.NewHostOsInstall(*openapiclient.NewOperatingSystemPartial(int32(123))) // HostOsInstall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.CreateHostOsInstall(context.Background(), hostId).HostOsInstall(hostOsInstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.CreateHostOsInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostOsInstall`: []HostOsInstall
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.CreateHostOsInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to install | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostOsInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostOsInstall** | [**HostOsInstall**](HostOsInstall.md) |  | 

### Return type

[**[]HostOsInstall**](HostOsInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> []Host GetHost(ctx, hostId).Execute()

Get a host (can be bare metal servers or VMs) and their main details

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: []Host
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetHostBmcStatuses

> []HostActionResponse GetHostBmcStatuses(ctx, hostId).Execute()

Get the BMC power status of the host

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
	hostId := int32(56) // int32 | ID of the host to fetch BMC status for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostBmcStatuses(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostBmcStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostBmcStatuses`: []HostActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostBmcStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch BMC status for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostBmcStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostActionResponse**](HostActionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostInstanceTypes

> []InstanceTypeBM GetHostInstanceTypes(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all bare metal instance types (and their details) available for your account

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
	resp, r, err := apiClient.HostAPI.GetHostInstanceTypes(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostInstanceTypes`: []InstanceTypeBM
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostInstanceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]InstanceTypeBM**](InstanceTypeBM.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostInvoices

> []Invoice GetHostInvoices(ctx, hostId).RANGEDDATA(rANGEDDATA).Execute()

Get all the invoices for a specific host

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
	hostId := int32(56) // int32 | ID of the host to fetch invoices for
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostInvoices(context.Background(), hostId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostInvoices`: []Invoice
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch invoices for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostLiveStates

> []HostLiveState GetHostLiveStates(ctx).Execute()

Get live situation of all the gaming hosts

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
	resp, r, err := apiClient.HostAPI.GetHostLiveStates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostLiveStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostLiveStates`: []HostLiveState
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostLiveStates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostLiveStatesRequest struct via the builder pattern


### Return type

[**[]HostLiveState**](HostLiveState.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostNetworkUsageGraphDays

> []string GetHostNetworkUsageGraphDays(ctx, hostId, upLinkId).Execute()

Get the traffic usage graph (last 24 hours)

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
	hostId := int32(56) // int32 | ID of the host to fetch traffic for
	upLinkId := int32(56) // int32 | ID of the uplink to fetch traffic for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostNetworkUsageGraphDays(context.Background(), hostId, upLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostNetworkUsageGraphDays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostNetworkUsageGraphDays`: []string
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostNetworkUsageGraphDays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch traffic for | 
**upLinkId** | **int32** | ID of the uplink to fetch traffic for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostNetworkUsageGraphDaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostNetworkUsageGraphMonths

> []string GetHostNetworkUsageGraphMonths(ctx, hostId, upLinkId).Execute()

Get the traffic usage graph (last 30 days)

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
	hostId := int32(56) // int32 | ID of the host to fetch traffic for
	upLinkId := int32(56) // int32 | ID of the uplink to fetch traffic for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostNetworkUsageGraphMonths(context.Background(), hostId, upLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostNetworkUsageGraphMonths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostNetworkUsageGraphMonths`: []string
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostNetworkUsageGraphMonths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch traffic for | 
**upLinkId** | **int32** | ID of the uplink to fetch traffic for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostNetworkUsageGraphMonthsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostNetworkUsageGraphWeeks

> []string GetHostNetworkUsageGraphWeeks(ctx, hostId, upLinkId).Execute()

Get the traffic usage graph (last 7 days)

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
	hostId := int32(56) // int32 | ID of the host to fetch traffic for
	upLinkId := int32(56) // int32 | ID of the uplink to fetch traffic for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostNetworkUsageGraphWeeks(context.Background(), hostId, upLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostNetworkUsageGraphWeeks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostNetworkUsageGraphWeeks`: []string
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostNetworkUsageGraphWeeks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch traffic for | 
**upLinkId** | **int32** | ID of the uplink to fetch traffic for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostNetworkUsageGraphWeeksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostNetworkUsageGraphYears

> []string GetHostNetworkUsageGraphYears(ctx, hostId, upLinkId).Execute()

Get the traffic usage graph (last year)

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
	hostId := int32(56) // int32 | ID of the host to fetch traffic for
	upLinkId := int32(56) // int32 | ID of the uplink to fetch traffic for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostNetworkUsageGraphYears(context.Background(), hostId, upLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostNetworkUsageGraphYears``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostNetworkUsageGraphYears`: []string
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostNetworkUsageGraphYears`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch traffic for | 
**upLinkId** | **int32** | ID of the uplink to fetch traffic for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostNetworkUsageGraphYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostNetworkUsageRaw

> []TrafficUsage GetHostNetworkUsageRaw(ctx, hostId, upLinkId).StartTime(startTime).EndTime(endTime).Execute()

Get the traffic usage data for the last day

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
	hostId := int32(56) // int32 | ID of the host to fetch traffic for
	upLinkId := int32(56) // int32 | ID of the uplink to fetch traffic for
	startTime := int32(56) // int32 | Start unix timestamp (optional)
	endTime := int32(56) // int32 | End unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostNetworkUsageRaw(context.Background(), hostId, upLinkId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostNetworkUsageRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostNetworkUsageRaw`: []TrafficUsage
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostNetworkUsageRaw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch traffic for | 
**upLinkId** | **int32** | ID of the uplink to fetch traffic for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostNetworkUsageRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **int32** | Start unix timestamp | 
 **endTime** | **int32** | End unix timestamp | 

### Return type

[**[]TrafficUsage**](TrafficUsage.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostOsInstallLogs

> []ServerLog GetHostOsInstallLogs(ctx, hostId).RANGEDDATA(rANGEDDATA).Execute()

Get all OS installation logs on this host

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
	hostId := int32(56) // int32 | ID of the host to fetch install logs for
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostOsInstallLogs(context.Background(), hostId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostOsInstallLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostOsInstallLogs`: []ServerLog
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostOsInstallLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch install logs for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostOsInstallLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]ServerLog**](ServerLog.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostStatuses

> []HostActionResponse GetHostStatuses(ctx, hostId).Execute()

Get the power status of the host

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
	hostId := int32(56) // int32 | ID of the host to fetch the status for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostStatuses(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostStatuses`: []HostActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch the status for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostActionResponse**](HostActionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostSummaries

> []HostSummary GetHostSummaries(ctx).Fields(fields).RANGEDDATA(rANGEDDATA).Execute()

Get all your hosts (can be bare metal servers or VMs) and their main details

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
	fields := "fields_example" // string | Fields return in response, only fields available in model can be filter out (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostSummaries(context.Background()).Fields(fields).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostSummaries`: []HostSummary
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Fields return in response, only fields available in model can be filter out | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]HostSummary**](HostSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostTasks

> []TaskStatus GetHostTasks(ctx, hostId).Execute()

Get the last 50 tasks of the given host

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
	hostId := "hostId_example" // string | The ID of the host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHostTasks(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHostTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostTasks`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHostTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The ID of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TaskStatus**](TaskStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosts

> []Host GetHosts(ctx).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()

Get all your hosts (can be bare metal servers or VMs) and their main details

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
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/host?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request instead.<br /> For more information about labels check <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">details</a> (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHosts(context.Background()).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosts`: []Host
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/host?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request instead.&lt;br /&gt; For more information about labels check &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;details&lt;/a&gt; | 
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


## RevertHostCancellation

> Host RevertHostCancellation(ctx, hostId).Execute()

Reverts the cancellation of a host. You cannot revert the cancellation a host if the current date is equal to or past the end of contract date, nor if the host was never cancelled

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
	hostId := int32(56) // int32 | ID of the host to revert the cancellation for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.RevertHostCancellation(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.RevertHostCancellation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertHostCancellation`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.RevertHostCancellation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to revert the cancellation for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertHostCancellationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Host**](Host.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> []Host UpdateHost(ctx, hostId).Host(host).Execute()

Can add, edit or delete host labels. To delete label need to provide key with null value.

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
	host := *openapiclient.NewHost(int32(123), int32(123), int32(123), "ServerName_example", int32(123), "LiveHostName_example", int32(123), "Category_example", int32(123), int32(123), int32(123), "InstanceType_example", "FleetId_example", "NewFleetId_example", int32(123), "RackName_example", "DateStart_example", "DateEnd_example", "DateCancelled_example", "DateEndContract_example", int32(123), int32(123), int32(123), "PurchaseOrder_example", int32(123), "PricePerMonth_example", "PricePerTbOveruse_example", int32(123), int32(123), int32(123), []openapiclient.HostIP{*openapiclient.NewHostIP("IpAddress_example", int32(123), int32(123), int32(123), int32(123), "MacAddress_example", "RDns_example", NullableInt32(123), "Gateway_example", "Netmask_example", NullableInt32(123))}, "Brand_example", "Model_example", int32(123), "CpuInfo_example", "CpuType_example", float32(123), *openapiclient.NewHostCpu(int32(123), int32(123), int32(123), "Info_example", "Type_example"), int32(123), int32(123), int32(123), []openapiclient.HostDisk{*openapiclient.NewHostDisk("DiskType_example", "DiskMedium_example", "Model_example", "Product_example", "DiskSerial_example", "FirmwareVersion_example", int32(123), int32(123), int32(123), int32(123))}, []openapiclient.HostMemory{*openapiclient.NewHostMemory("Brand_example", "Model_example", int32(123), int32(123), int32(123), int32(123), "MemoryType_example", "MemorySlot_example", "MemorySerial_example")}, int32(123), "ServiceTag_example", int32(123), "FmOrderId_example", "InstallStatus_example", "Status_example", int32(123), []openapiclient.HostUplink{*openapiclient.NewHostUplink(int32(123), false)}, false, int32(123), int32(123)) // Host | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.UpdateHost(context.Background(), hostId).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.UpdateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHost`: []Host
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **host** | [**Host**](Host.md) |  | 

### Return type

[**[]Host**](Host.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostRestart

> []HostActionResponse UpdateHostRestart(ctx, hostId).Execute()

Restart the host

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
	hostId := int32(56) // int32 | ID of the host to restart

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.UpdateHostRestart(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.UpdateHostRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostRestart`: []HostActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.UpdateHostRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to restart | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostActionResponse**](HostActionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostStart

> []HostActionResponse UpdateHostStart(ctx, hostId).Execute()

Turn the host On

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
	hostId := int32(56) // int32 | ID of the host to turn on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.UpdateHostStart(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.UpdateHostStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostStart`: []HostActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.UpdateHostStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to turn on | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostActionResponse**](HostActionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostStop

> []HostActionResponse UpdateHostStop(ctx, hostId).Execute()

Turn the host Off

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
	hostId := int32(56) // int32 | ID of the host to turn off

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.UpdateHostStop(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.UpdateHostStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostStop`: []HostActionResponse
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.UpdateHostStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to turn off | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostActionResponse**](HostActionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

