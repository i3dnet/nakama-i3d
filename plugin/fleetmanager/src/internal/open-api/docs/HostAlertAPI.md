# \HostAlertAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostAlert**](HostAlertAPI.md#CreateHostAlert) | **Post** /v3/host/{hostId}/alert | Create alert for the host
[**DeleteHostAlert**](HostAlertAPI.md#DeleteHostAlert) | **Delete** /v3/host/{hostId}/alert/{hostAlertId} | Delete host alert
[**GetHostAlerts**](HostAlertAPI.md#GetHostAlerts) | **Get** /v3/host/{hostId}/alert | Get alerts defined for the given host
[**UpdateHostAlert**](HostAlertAPI.md#UpdateHostAlert) | **Put** /v3/host/{hostId}/alert/{hostAlertId} | Update host alert
[**UpdateHostNetworkRdn**](HostAlertAPI.md#UpdateHostNetworkRdn) | **Put** /v3/host/{hostId}/network/rdns/{ipAddress} | Set RDNS hostname for IP



## CreateHostAlert

> []HostAlert CreateHostAlert(ctx, hostId).HostAlert(hostAlert).Execute()

Create alert for the host

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
	hostId := int32(56) // int32 | ID of the host to create an alert for
	hostAlert := *openapiclient.NewHostAlert(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)) // HostAlert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAlertAPI.CreateHostAlert(context.Background(), hostId).HostAlert(hostAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAlertAPI.CreateHostAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostAlert`: []HostAlert
	fmt.Fprintf(os.Stdout, "Response from `HostAlertAPI.CreateHostAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to create an alert for | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostAlert** | [**HostAlert**](HostAlert.md) |  | 

### Return type

[**[]HostAlert**](HostAlert.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostAlert

> DeleteHostAlert(ctx, hostId, hostAlertId).Execute()

Delete host alert

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
	hostId := int32(56) // int32 | ID of the host that the alert belongs to
	hostAlertId := int32(56) // int32 | ID of the alert to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAlertAPI.DeleteHostAlert(context.Background(), hostId, hostAlertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAlertAPI.DeleteHostAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host that the alert belongs to | 
**hostAlertId** | **int32** | ID of the alert to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostAlertRequest struct via the builder pattern


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


## GetHostAlerts

> []HostAlert GetHostAlerts(ctx, hostId).RANGEDDATA(rANGEDDATA).Execute()

Get alerts defined for the given host

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
	hostId := int32(56) // int32 | ID of the host to fetch alert for
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAlertAPI.GetHostAlerts(context.Background(), hostId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAlertAPI.GetHostAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostAlerts`: []HostAlert
	fmt.Fprintf(os.Stdout, "Response from `HostAlertAPI.GetHostAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to fetch alert for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]HostAlert**](HostAlert.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostAlert

> []HostAlert UpdateHostAlert(ctx, hostId, hostAlertId).HostAlert(hostAlert).Execute()

Update host alert

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
	hostId := int32(56) // int32 | ID of the host that the alert belongs to
	hostAlertId := int32(56) // int32 | ID of the host alert to update
	hostAlert := *openapiclient.NewHostAlert(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)) // HostAlert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAlertAPI.UpdateHostAlert(context.Background(), hostId, hostAlertId).HostAlert(hostAlert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAlertAPI.UpdateHostAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostAlert`: []HostAlert
	fmt.Fprintf(os.Stdout, "Response from `HostAlertAPI.UpdateHostAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host that the alert belongs to | 
**hostAlertId** | **int32** | ID of the host alert to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostAlert** | [**HostAlert**](HostAlert.md) |  | 

### Return type

[**[]HostAlert**](HostAlert.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostNetworkRdn

> UpdateHostNetworkRdn(ctx, hostId, ipAddress).RDns(rDns).Execute()

Set RDNS hostname for IP

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
	hostId := int32(56) // int32 | ID of the host to which IP address belongs to
	ipAddress := "ipAddress_example" // string | IP address to set RDNS entry for
	rDns := *openapiclient.NewRDns("Hostname_example", "Ip_example") // RDns | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAlertAPI.UpdateHostNetworkRdn(context.Background(), hostId, ipAddress).RDns(rDns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAlertAPI.UpdateHostNetworkRdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | ID of the host to which IP address belongs to | 
**ipAddress** | **string** | IP address to set RDNS entry for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostNetworkRdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rDns** | [**RDns**](RDns.md) |  | 

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

