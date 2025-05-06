# \CloudAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudCpuPlatform**](CloudAPI.md#GetCloudCpuPlatform) | **Get** /v3/cloud/cpuPlatform/{dcLocationId} | Get available cloud provider CPU platforms for a specific dcLocationId
[**GetCloudCpuPlatforms**](CloudAPI.md#GetCloudCpuPlatforms) | **Get** /v3/cloud/cpuPlatform | Get available cloud provider CPU platforms
[**GetCloudDcLocation**](CloudAPI.md#GetCloudDcLocation) | **Get** /v3/cloud/dcLocation/{dcLocationId} | Get details for a specific data center location
[**GetCloudDcLocationVicinityMappings**](CloudAPI.md#GetCloudDcLocationVicinityMappings) | **Get** /v3/cloud/dcLocation/vicinityMapping | Get a mapping of all cloud datacenter locations in the vicinity of i3D.net dcLocations
[**GetCloudDcLocationVicinityMappingsByDcLocationIds**](CloudAPI.md#GetCloudDcLocationVicinityMappingsByDcLocationIds) | **Get** /v3/cloud/dcLocation/{dcLocationIds}/vicinityMapping | Get a mapping of all cloud data center locations in the vicinity of a i3D.net dcLocation
[**GetCloudDcLocations**](CloudAPI.md#GetCloudDcLocations) | **Get** /v3/cloud/dcLocation | Get all data center locations of all cloud providers
[**GetCloudInstanceType**](CloudAPI.md#GetCloudInstanceType) | **Get** /v3/cloud/instanceType/{dcLocationId} | Get cloud provider instance types for a specific dcLocationId
[**GetCloudInstanceTypes**](CloudAPI.md#GetCloudInstanceTypes) | **Get** /v3/cloud/instanceType | Get cloud provider instance types
[**GetCloudInstances**](CloudAPI.md#GetCloudInstances) | **Get** /v3/cloud/instance | Get list of all the cloud instances
[**GetCloudProviderDcLocations**](CloudAPI.md#GetCloudProviderDcLocations) | **Get** /v3/cloud/provider/{providerId}/dcLocation | Get all available data center locations for a specific cloud provider
[**GetCloudProviders**](CloudAPI.md#GetCloudProviders) | **Get** /v3/cloud/provider | Get all available cloud provider details
[**GetCloudVmErrors**](CloudAPI.md#GetCloudVmErrors) | **Get** /v3/cloud/vm/error | Get all available virtual machine errors
[**GetCloudVmLogs**](CloudAPI.md#GetCloudVmLogs) | **Get** /v3/cloud/vm/log | Get all available virtual machine logs
[**GetCloudVmUptimes**](CloudAPI.md#GetCloudVmUptimes) | **Get** /v3/cloud/vm/uptime | Get all available virtual machine uptime



## GetCloudCpuPlatform

> []CpuPlatform GetCloudCpuPlatform(ctx, dcLocationId).Execute()

Get available cloud provider CPU platforms for a specific dcLocationId

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
	dcLocationId := int32(56) // int32 | The Id of the data center location. For a list of all data center locations see [`GET /v3/cloud/dcLocation`](#/Cloud/get_v3_cloud_dcLocation)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.GetCloudCpuPlatform(context.Background(), dcLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudCpuPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudCpuPlatform`: []CpuPlatform
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudCpuPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dcLocationId** | **int32** | The Id of the data center location. For a list of all data center locations see [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/get_v3_cloud_dcLocation) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCpuPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CpuPlatform**](CpuPlatform.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudCpuPlatforms

> []CpuPlatform GetCloudCpuPlatforms(ctx).Execute()

Get available cloud provider CPU platforms

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
	resp, r, err := apiClient.CloudAPI.GetCloudCpuPlatforms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudCpuPlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudCpuPlatforms`: []CpuPlatform
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudCpuPlatforms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCpuPlatformsRequest struct via the builder pattern


### Return type

[**[]CpuPlatform**](CpuPlatform.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDcLocation

> []DcLocation GetCloudDcLocation(ctx, dcLocationId).Execute()

Get details for a specific data center location

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
	dcLocationId := "dcLocationId_example" // string | i3D.net or cloud data center location ID. For a list of all data center locations see [`GET /v3/cloud/dcLocation`](#/Cloud/get_v3_cloud_dcLocation)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.GetCloudDcLocation(context.Background(), dcLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudDcLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudDcLocation`: []DcLocation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudDcLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dcLocationId** | **string** | i3D.net or cloud data center location ID. For a list of all data center locations see [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/get_v3_cloud_dcLocation) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDcLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DcLocation**](DcLocation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDcLocationVicinityMappings

> []DcLocationVicinityMapping GetCloudDcLocationVicinityMappings(ctx).Execute()

Get a mapping of all cloud datacenter locations in the vicinity of i3D.net dcLocations

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
	resp, r, err := apiClient.CloudAPI.GetCloudDcLocationVicinityMappings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudDcLocationVicinityMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudDcLocationVicinityMappings`: []DcLocationVicinityMapping
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudDcLocationVicinityMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDcLocationVicinityMappingsRequest struct via the builder pattern


### Return type

[**[]DcLocationVicinityMapping**](DcLocationVicinityMapping.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDcLocationVicinityMappingsByDcLocationIds

> []DcLocationVicinityMapping GetCloudDcLocationVicinityMappingsByDcLocationIds(ctx, dcLocationIds).Execute()

Get a mapping of all cloud data center locations in the vicinity of a i3D.net dcLocation

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
	dcLocationIds := "dcLocationIds_example" // string | Comma separated array of data center location Ids, example: id1,id2. For a list of all data center locations see [`GET /v3/cloud/dcLocation`](#/Cloud/get_v3_cloud_dcLocation)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.GetCloudDcLocationVicinityMappingsByDcLocationIds(context.Background(), dcLocationIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudDcLocationVicinityMappingsByDcLocationIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudDcLocationVicinityMappingsByDcLocationIds`: []DcLocationVicinityMapping
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudDcLocationVicinityMappingsByDcLocationIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dcLocationIds** | **string** | Comma separated array of data center location Ids, example: id1,id2. For a list of all data center locations see [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/get_v3_cloud_dcLocation) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDcLocationVicinityMappingsByDcLocationIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DcLocationVicinityMapping**](DcLocationVicinityMapping.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDcLocations

> []DcLocation GetCloudDcLocations(ctx).Execute()

Get all data center locations of all cloud providers

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
	resp, r, err := apiClient.CloudAPI.GetCloudDcLocations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudDcLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudDcLocations`: []DcLocation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudDcLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDcLocationsRequest struct via the builder pattern


### Return type

[**[]DcLocation**](DcLocation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudInstanceType

> []InstanceType GetCloudInstanceType(ctx, dcLocationId).Execute()

Get cloud provider instance types for a specific dcLocationId

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
	dcLocationId := int32(56) // int32 | The Id of the data center location. For a list of all data center locations see [`GET /v3/cloud/dcLocation`](#/Cloud/get_v3_cloud_dcLocation)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.GetCloudInstanceType(context.Background(), dcLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudInstanceType`: []InstanceType
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dcLocationId** | **int32** | The Id of the data center location. For a list of all data center locations see [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/get_v3_cloud_dcLocation) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstanceType**](InstanceType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudInstanceTypes

> []InstanceType GetCloudInstanceTypes(ctx).Execute()

Get cloud provider instance types

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
	resp, r, err := apiClient.CloudAPI.GetCloudInstanceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudInstanceTypes`: []InstanceType
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudInstanceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudInstanceTypesRequest struct via the builder pattern


### Return type

[**[]InstanceType**](InstanceType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudInstances

> []CloudInstance GetCloudInstances(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get list of all the cloud instances

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
	resp, r, err := apiClient.CloudAPI.GetCloudInstances(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudInstances`: []CloudInstance
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]CloudInstance**](CloudInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderDcLocations

> []DcLocation GetCloudProviderDcLocations(ctx, providerId).Execute()

Get all available data center locations for a specific cloud provider

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
	providerId := "providerId_example" // string | Cloud provider ID. For a list of cloud providers see [`GET /v3/cloud/provider`](#/Cloud/get_v3_cloud_provider)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.GetCloudProviderDcLocations(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudProviderDcLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProviderDcLocations`: []DcLocation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudProviderDcLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | Cloud provider ID. For a list of cloud providers see [&#x60;GET /v3/cloud/provider&#x60;](#/Cloud/get_v3_cloud_provider) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderDcLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DcLocation**](DcLocation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviders

> []CloudProvider GetCloudProviders(ctx).Execute()

Get all available cloud provider details

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
	resp, r, err := apiClient.CloudAPI.GetCloudProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProviders`: []CloudProvider
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersRequest struct via the builder pattern


### Return type

[**[]CloudProvider**](CloudProvider.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudVmErrors

> []VmError GetCloudVmErrors(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all available virtual machine errors



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
	resp, r, err := apiClient.CloudAPI.GetCloudVmErrors(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudVmErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudVmErrors`: []VmError
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudVmErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudVmErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]VmError**](VmError.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudVmLogs

> []VmLog GetCloudVmLogs(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all available virtual machine logs



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
	resp, r, err := apiClient.CloudAPI.GetCloudVmLogs(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudVmLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudVmLogs`: []VmLog
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudVmLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudVmLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]VmLog**](VmLog.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudVmUptimes

> []VmUptime GetCloudVmUptimes(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all available virtual machine uptime



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
	resp, r, err := apiClient.CloudAPI.GetCloudVmUptimes(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudVmUptimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudVmUptimes`: []VmUptime
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudVmUptimes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudVmUptimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]VmUptime**](VmUptime.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

