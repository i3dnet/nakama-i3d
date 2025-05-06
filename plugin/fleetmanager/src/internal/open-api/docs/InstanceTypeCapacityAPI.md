# \InstanceTypeCapacityAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostCapacityTemplateInstanceTypeCapacity**](InstanceTypeCapacityAPI.md#CreateHostCapacityTemplateInstanceTypeCapacity) | **Post** /v3/hostCapacityTemplate/{hostCapacityTemplateId}/instanceTypeCapacity | Create an instance type capacity for the specified hostCapacityTemplateId
[**DeleteInstanceTypeCapacity**](InstanceTypeCapacityAPI.md#DeleteInstanceTypeCapacity) | **Delete** /v3/instanceTypeCapacity/{instanceTypeCapacityId} | Delete given host capacity template
[**GetHostCapacityTemplateInstanceTypeCapacities**](InstanceTypeCapacityAPI.md#GetHostCapacityTemplateInstanceTypeCapacities) | **Get** /v3/hostCapacityTemplate/{hostCapacityTemplateId}/instanceTypeCapacity | Get the details of all your instance type capacities under a host capacity template
[**GetInstanceTypeCapacity**](InstanceTypeCapacityAPI.md#GetInstanceTypeCapacity) | **Get** /v3/instanceTypeCapacity/{instanceTypeCapacityId} | Get the details of the given instance type capacity
[**UpdateInstanceTypeCapacity**](InstanceTypeCapacityAPI.md#UpdateInstanceTypeCapacity) | **Put** /v3/instanceTypeCapacity/{instanceTypeCapacityId} | Update given host capacity template



## CreateHostCapacityTemplateInstanceTypeCapacity

> []InstanceTypeCapacity CreateHostCapacityTemplateInstanceTypeCapacity(ctx, hostCapacityTemplateId).InstanceTypeCapacity(instanceTypeCapacity).Execute()

Create an instance type capacity for the specified hostCapacityTemplateId

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The ID of the hostCapacityTemplate
	instanceTypeCapacity := *openapiclient.NewInstanceTypeCapacity("Id_example", "HostCapacityTemplateId_example", int32(123), "InstanceType_example", int32(123), int32(123), int32(123)) // InstanceTypeCapacity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeCapacityAPI.CreateHostCapacityTemplateInstanceTypeCapacity(context.Background(), hostCapacityTemplateId).InstanceTypeCapacity(instanceTypeCapacity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeCapacityAPI.CreateHostCapacityTemplateInstanceTypeCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostCapacityTemplateInstanceTypeCapacity`: []InstanceTypeCapacity
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeCapacityAPI.CreateHostCapacityTemplateInstanceTypeCapacity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The ID of the hostCapacityTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostCapacityTemplateInstanceTypeCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceTypeCapacity** | [**InstanceTypeCapacity**](InstanceTypeCapacity.md) |  | 

### Return type

[**[]InstanceTypeCapacity**](InstanceTypeCapacity.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceTypeCapacity

> DeleteInstanceTypeCapacity(ctx, instanceTypeCapacityId).Execute()

Delete given host capacity template

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
	instanceTypeCapacityId := "instanceTypeCapacityId_example" // string | The Id of the host capacity template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceTypeCapacityAPI.DeleteInstanceTypeCapacity(context.Background(), instanceTypeCapacityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeCapacityAPI.DeleteInstanceTypeCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeCapacityId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeCapacityRequest struct via the builder pattern


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


## GetHostCapacityTemplateInstanceTypeCapacities

> []InstanceTypeCapacity GetHostCapacityTemplateInstanceTypeCapacities(ctx, hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get the details of all your instance type capacities under a host capacity template

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The ID of the hostCapacityTemplate
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeCapacityAPI.GetHostCapacityTemplateInstanceTypeCapacities(context.Background(), hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeCapacityAPI.GetHostCapacityTemplateInstanceTypeCapacities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCapacityTemplateInstanceTypeCapacities`: []InstanceTypeCapacity
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeCapacityAPI.GetHostCapacityTemplateInstanceTypeCapacities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The ID of the hostCapacityTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCapacityTemplateInstanceTypeCapacitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]InstanceTypeCapacity**](InstanceTypeCapacity.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeCapacity

> []InstanceTypeCapacity GetInstanceTypeCapacity(ctx, instanceTypeCapacityId).Execute()

Get the details of the given instance type capacity

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
	instanceTypeCapacityId := "instanceTypeCapacityId_example" // string | The Id of the instance type capacity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeCapacityAPI.GetInstanceTypeCapacity(context.Background(), instanceTypeCapacityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeCapacityAPI.GetInstanceTypeCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeCapacity`: []InstanceTypeCapacity
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeCapacityAPI.GetInstanceTypeCapacity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeCapacityId** | **string** | The Id of the instance type capacity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstanceTypeCapacity**](InstanceTypeCapacity.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceTypeCapacity

> []InstanceTypeCapacity UpdateInstanceTypeCapacity(ctx, instanceTypeCapacityId).InstanceTypeCapacity(instanceTypeCapacity).Execute()

Update given host capacity template

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
	instanceTypeCapacityId := "instanceTypeCapacityId_example" // string | The Id of the host capacity template
	instanceTypeCapacity := *openapiclient.NewInstanceTypeCapacity("Id_example", "HostCapacityTemplateId_example", int32(123), "InstanceType_example", int32(123), int32(123), int32(123)) // InstanceTypeCapacity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeCapacityAPI.UpdateInstanceTypeCapacity(context.Background(), instanceTypeCapacityId).InstanceTypeCapacity(instanceTypeCapacity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeCapacityAPI.UpdateInstanceTypeCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstanceTypeCapacity`: []InstanceTypeCapacity
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeCapacityAPI.UpdateInstanceTypeCapacity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeCapacityId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceTypeCapacity** | [**InstanceTypeCapacity**](InstanceTypeCapacity.md) |  | 

### Return type

[**[]InstanceTypeCapacity**](InstanceTypeCapacity.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

