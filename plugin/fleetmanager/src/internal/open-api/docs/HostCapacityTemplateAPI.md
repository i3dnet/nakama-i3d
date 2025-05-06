# \HostCapacityTemplateAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostCapacityTemplate**](HostCapacityTemplateAPI.md#CreateHostCapacityTemplate) | **Post** /v3/hostCapacityTemplate | Create a host capacity template
[**DeleteHostCapacityTemplate**](HostCapacityTemplateAPI.md#DeleteHostCapacityTemplate) | **Delete** /v3/hostCapacityTemplate/{hostCapacityTemplateId} | Delete given host capacity template
[**GetHostCapacityTemplate**](HostCapacityTemplateAPI.md#GetHostCapacityTemplate) | **Get** /v3/hostCapacityTemplate/{hostCapacityTemplateId} | Get the details of the given host capacity template
[**GetHostCapacityTemplateApplicationBuilds**](HostCapacityTemplateAPI.md#GetHostCapacityTemplateApplicationBuilds) | **Get** /v3/hostCapacityTemplate/{hostCapacityTemplateId}/applicationBuild | Get all application builds that are using the given host capacity template
[**GetHostCapacityTemplateFleets**](HostCapacityTemplateAPI.md#GetHostCapacityTemplateFleets) | **Get** /v3/hostCapacityTemplate/{hostCapacityTemplateId}/fleet | Get all fleets that are using the given host capacity template
[**GetHostCapacityTemplates**](HostCapacityTemplateAPI.md#GetHostCapacityTemplates) | **Get** /v3/hostCapacityTemplate | Get the details of all your host capacity templates
[**UpdateHostCapacityTemplate**](HostCapacityTemplateAPI.md#UpdateHostCapacityTemplate) | **Put** /v3/hostCapacityTemplate/{hostCapacityTemplateId} | Update given host capacity template



## CreateHostCapacityTemplate

> []HostCapacityTemplate CreateHostCapacityTemplate(ctx).HostCapacityTemplate(hostCapacityTemplate).Execute()

Create a host capacity template

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
	hostCapacityTemplate := *openapiclient.NewHostCapacityTemplate("Id_example", []string{"FleetIds_example"}, []string{"ApplicationBuildIds_example"}, "Name_example", int32(123), int32(123)) // HostCapacityTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostCapacityTemplateAPI.CreateHostCapacityTemplate(context.Background()).HostCapacityTemplate(hostCapacityTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.CreateHostCapacityTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostCapacityTemplate`: []HostCapacityTemplate
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.CreateHostCapacityTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostCapacityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostCapacityTemplate** | [**HostCapacityTemplate**](HostCapacityTemplate.md) |  | 

### Return type

[**[]HostCapacityTemplate**](HostCapacityTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostCapacityTemplate

> DeleteHostCapacityTemplate(ctx, hostCapacityTemplateId).Execute()

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The Id of the host capacity template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostCapacityTemplateAPI.DeleteHostCapacityTemplate(context.Background(), hostCapacityTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.DeleteHostCapacityTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostCapacityTemplateRequest struct via the builder pattern


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


## GetHostCapacityTemplate

> []HostCapacityTemplate GetHostCapacityTemplate(ctx, hostCapacityTemplateId).Execute()

Get the details of the given host capacity template

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The Id of the host capacity template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostCapacityTemplateAPI.GetHostCapacityTemplate(context.Background(), hostCapacityTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.GetHostCapacityTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCapacityTemplate`: []HostCapacityTemplate
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.GetHostCapacityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCapacityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HostCapacityTemplate**](HostCapacityTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostCapacityTemplateApplicationBuilds

> []ApplicationBuild GetHostCapacityTemplateApplicationBuilds(ctx, hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get all application builds that are using the given host capacity template

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The Id of the host capacity template
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostCapacityTemplateAPI.GetHostCapacityTemplateApplicationBuilds(context.Background(), hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.GetHostCapacityTemplateApplicationBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCapacityTemplateApplicationBuilds`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.GetHostCapacityTemplateApplicationBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCapacityTemplateApplicationBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

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


## GetHostCapacityTemplateFleets

> []Fleet GetHostCapacityTemplateFleets(ctx, hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get all fleets that are using the given host capacity template

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The Id of the host capacity template
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostCapacityTemplateAPI.GetHostCapacityTemplateFleets(context.Background(), hostCapacityTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.GetHostCapacityTemplateFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCapacityTemplateFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.GetHostCapacityTemplateFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCapacityTemplateFleetsRequest struct via the builder pattern


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


## GetHostCapacityTemplates

> []HostCapacityTemplate GetHostCapacityTemplates(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get the details of all your host capacity templates

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
	resp, r, err := apiClient.HostCapacityTemplateAPI.GetHostCapacityTemplates(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.GetHostCapacityTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCapacityTemplates`: []HostCapacityTemplate
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.GetHostCapacityTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCapacityTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]HostCapacityTemplate**](HostCapacityTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostCapacityTemplate

> []HostCapacityTemplate UpdateHostCapacityTemplate(ctx, hostCapacityTemplateId).HostCapacityTemplate(hostCapacityTemplate).Execute()

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
	hostCapacityTemplateId := "hostCapacityTemplateId_example" // string | The Id of the host capacity template
	hostCapacityTemplate := *openapiclient.NewHostCapacityTemplate("Id_example", []string{"FleetIds_example"}, []string{"ApplicationBuildIds_example"}, "Name_example", int32(123), int32(123)) // HostCapacityTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostCapacityTemplateAPI.UpdateHostCapacityTemplate(context.Background(), hostCapacityTemplateId).HostCapacityTemplate(hostCapacityTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostCapacityTemplateAPI.UpdateHostCapacityTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHostCapacityTemplate`: []HostCapacityTemplate
	fmt.Fprintf(os.Stdout, "Response from `HostCapacityTemplateAPI.UpdateHostCapacityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostCapacityTemplateId** | **string** | The Id of the host capacity template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostCapacityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostCapacityTemplate** | [**HostCapacityTemplate**](HostCapacityTemplate.md) |  | 

### Return type

[**[]HostCapacityTemplate**](HostCapacityTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

