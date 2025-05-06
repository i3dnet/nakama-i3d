# \ArcusAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomCommandApplicationInstance**](ArcusAPI.md#CreateCustomCommandApplicationInstance) | **Post** /v3/customCommand/applicationInstance/{applicationInstanceId} | Sends a command to a application instance
[**CreateCustomCommandHost**](ArcusAPI.md#CreateCustomCommandHost) | **Post** /v3/customCommand/host/{hostId} | Sends a command to a host (all application instances will receive the command)



## CreateCustomCommandApplicationInstance

> []TaskStatus CreateCustomCommandApplicationInstance(ctx).CustomCommandCollection(customCommandCollection).Execute()

Sends a command to a application instance

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
	customCommandCollection := *openapiclient.NewCustomCommandCollection() // CustomCommandCollection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcusAPI.CreateCustomCommandApplicationInstance(context.Background()).CustomCommandCollection(customCommandCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcusAPI.CreateCustomCommandApplicationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomCommandApplicationInstance`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ArcusAPI.CreateCustomCommandApplicationInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomCommandApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customCommandCollection** | [**CustomCommandCollection**](CustomCommandCollection.md) |  | 

### Return type

[**[]TaskStatus**](TaskStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomCommandHost

> []TaskStatus CreateCustomCommandHost(ctx).CustomCommandCollection(customCommandCollection).Execute()

Sends a command to a host (all application instances will receive the command)

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
	customCommandCollection := *openapiclient.NewCustomCommandCollection() // CustomCommandCollection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArcusAPI.CreateCustomCommandHost(context.Background()).CustomCommandCollection(customCommandCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArcusAPI.CreateCustomCommandHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomCommandHost`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ArcusAPI.CreateCustomCommandHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomCommandHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customCommandCollection** | [**CustomCommandCollection**](CustomCommandCollection.md) |  | 

### Return type

[**[]TaskStatus**](TaskStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

