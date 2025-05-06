# \DeploymentProfileAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentProfile**](DeploymentProfileAPI.md#CreateDeploymentProfile) | **Post** /v3/deploymentProfile | Create a new deployment profile
[**CreateDeploymentProfileDeploymentRegion**](DeploymentProfileAPI.md#CreateDeploymentProfileDeploymentRegion) | **Post** /v3/deploymentProfile/{deploymentProfileId}/deploymentRegion | Create a new deployment region for the given profile
[**CreateDeploymentRegionProperty**](DeploymentProfileAPI.md#CreateDeploymentRegionProperty) | **Post** /v3/deploymentRegion/{deploymentRegionId}/property | Create a new deployment region property
[**DeleteDeploymentContainer**](DeploymentProfileAPI.md#DeleteDeploymentContainer) | **Delete** /v3/deploymentContainer/{deploymentContainerId} | Delete given deployment profile container (it will delete all the application instances if there are any)
[**DeleteDeploymentContainerLocation**](DeploymentProfileAPI.md#DeleteDeploymentContainerLocation) | **Delete** /v3/deploymentContainerLocation/{deploymentContainerLocationId} | Delete given deployment profile containerLocation (it will delete all the application instances if there are any)
[**DeleteDeploymentProfile**](DeploymentProfileAPI.md#DeleteDeploymentProfile) | **Delete** /v3/deploymentProfile/{deploymentProfileId} | Delete given deployment profile (it will delete all the application instances if there are any)
[**DeleteDeploymentRegion**](DeploymentProfileAPI.md#DeleteDeploymentRegion) | **Delete** /v3/deploymentRegion/{deploymentRegionId} | Delete given deployment profile region (it will delete all the application instances if there are any)
[**DeleteDeploymentRegionProperty**](DeploymentProfileAPI.md#DeleteDeploymentRegionProperty) | **Delete** /v3/deploymentRegion/{deploymentRegionId}/property/{propertyId} | Deletes a deployment region property (does not delete the application property)
[**GetDeploymentProfile**](DeploymentProfileAPI.md#GetDeploymentProfile) | **Get** /v3/deploymentProfile/{deploymentProfileId} | Get the given deployment profile
[**GetDeploymentProfileSummaries**](DeploymentProfileAPI.md#GetDeploymentProfileSummaries) | **Get** /v3/deploymentProfile/summary | Get all the deployment profiles you have created
[**GetDeploymentProfiles**](DeploymentProfileAPI.md#GetDeploymentProfiles) | **Get** /v3/deploymentProfile | Get all the deployment profiles you have created
[**GetDeploymentRegion**](DeploymentProfileAPI.md#GetDeploymentRegion) | **Get** /v3/deploymentRegion/{deploymentRegionId} | Get the given deployment region
[**GetDeploymentRegionProperties**](DeploymentProfileAPI.md#GetDeploymentRegionProperties) | **Get** /v3/deploymentRegion/{deploymentRegionId}/property | Get a list of deployment region properties
[**GetDeploymentRegionProperty**](DeploymentProfileAPI.md#GetDeploymentRegionProperty) | **Get** /v3/deploymentRegion/{deploymentRegionId}/property/{propertyId} | Get a specific deployment region property
[**UpdateDeploymentContainerCanceldelete**](DeploymentProfileAPI.md#UpdateDeploymentContainerCanceldelete) | **Put** /v3/deploymentContainer/{deploymentContainerId}/canceldelete | Cancel delete given cloud deployment container
[**UpdateDeploymentContainerLocationCanceldelete**](DeploymentProfileAPI.md#UpdateDeploymentContainerLocationCanceldelete) | **Put** /v3/deploymentContainerLocation/{deploymentContainerLocationId}/canceldelete | Cancel deletion of given deployment profile containerLocation
[**UpdateDeploymentProfile**](DeploymentProfileAPI.md#UpdateDeploymentProfile) | **Put** /v3/deploymentProfile/{deploymentProfileId} | Update given deployment profile
[**UpdateDeploymentProfileCanceldelete**](DeploymentProfileAPI.md#UpdateDeploymentProfileCanceldelete) | **Put** /v3/deploymentProfile/{deploymentProfileId}/canceldelete | Cancel the deletion of given deployment profile (it will stop deletion but what has been deleted is lost already)
[**UpdateDeploymentRegion**](DeploymentProfileAPI.md#UpdateDeploymentRegion) | **Put** /v3/deploymentRegion/{deploymentRegionId} | Update given deployment region
[**UpdateDeploymentRegionCanceldelete**](DeploymentProfileAPI.md#UpdateDeploymentRegionCanceldelete) | **Put** /v3/deploymentRegion/{deploymentRegionId}/canceldelete | Cancel deletion of given deployment region (it will stop deletion but what has been deleted is lost already)
[**UpdateDeploymentRegionProperty**](DeploymentProfileAPI.md#UpdateDeploymentRegionProperty) | **Put** /v3/deploymentRegion/{deploymentRegionId}/property/{propertyId} | Update a deployment region property



## CreateDeploymentProfile

> []DeploymentProfile CreateDeploymentProfile(ctx).DeploymentProfile(deploymentProfile).Execute()

Create a new deployment profile

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
	deploymentProfile := *openapiclient.NewDeploymentProfile("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), int32(123), int32(123), int32(1), int32(123), []openapiclient.DeploymentRegion{*openapiclient.NewDeploymentRegion("Id_example", "Name_example", []int32{int32(123)}, int32(123), int32(123))}, int32(123)) // DeploymentProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.CreateDeploymentProfile(context.Background()).DeploymentProfile(deploymentProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.CreateDeploymentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentProfile`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.CreateDeploymentProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentProfile** | [**DeploymentProfile**](DeploymentProfile.md) |  | 

### Return type

[**[]DeploymentProfile**](DeploymentProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeploymentProfileDeploymentRegion

> []DeploymentRegion CreateDeploymentProfileDeploymentRegion(ctx, deploymentProfileId).DeploymentRegion(deploymentRegion).Execute()

Create a new deployment region for the given profile

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
	deploymentProfileId := "deploymentProfileId_example" // string | The Id of the deployment profile
	deploymentRegion := *openapiclient.NewDeploymentRegion("Id_example", "Name_example", []int32{int32(123)}, int32(123), int32(123)) // DeploymentRegion | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.CreateDeploymentProfileDeploymentRegion(context.Background(), deploymentProfileId).DeploymentRegion(deploymentRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.CreateDeploymentProfileDeploymentRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentProfileDeploymentRegion`: []DeploymentRegion
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.CreateDeploymentProfileDeploymentRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentProfileId** | **string** | The Id of the deployment profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentProfileDeploymentRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentRegion** | [**DeploymentRegion**](DeploymentRegion.md) |  | 

### Return type

[**[]DeploymentRegion**](DeploymentRegion.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeploymentRegionProperty

> []DeploymentRegionProperty CreateDeploymentRegionProperty(ctx, deploymentRegionId).DeploymentRegionProperty(deploymentRegionProperty).Execute()

Create a new deployment region property

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region
	deploymentRegionProperty := *openapiclient.NewDeploymentRegionProperty("Id_example", "ApplicationPropertyId_example", "PropertyValue_example") // DeploymentRegionProperty | Deployment region property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.CreateDeploymentRegionProperty(context.Background(), deploymentRegionId).DeploymentRegionProperty(deploymentRegionProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.CreateDeploymentRegionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentRegionProperty`: []DeploymentRegionProperty
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.CreateDeploymentRegionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRegionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentRegionProperty** | [**DeploymentRegionProperty**](DeploymentRegionProperty.md) | Deployment region property | 

### Return type

[**[]DeploymentRegionProperty**](DeploymentRegionProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentContainer

> []TaskStatus DeleteDeploymentContainer(ctx, deploymentContainerId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()

Delete given deployment profile container (it will delete all the application instances if there are any)



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
	deploymentContainerId := "deploymentContainerId_example" // string | Id of the deployment profile container to delete
	methodId := int32(56) // int32 | Id for the method of deletion: hard-kill = 1, graceful=2 (optional)
	startTime := int32(56) // int32 | Unix timestamp. When to start deleting the containerLocation and its contents. Default = now (optional)
	endTime := int32(56) // int32 | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default = 0 (no limit) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.DeleteDeploymentContainer(context.Background(), deploymentContainerId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.DeleteDeploymentContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeploymentContainer`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.DeleteDeploymentContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentContainerId** | **string** | Id of the deployment profile container to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **methodId** | **int32** | Id for the method of deletion: hard-kill &#x3D; 1, graceful&#x3D;2 | 
 **startTime** | **int32** | Unix timestamp. When to start deleting the containerLocation and its contents. Default &#x3D; now | 
 **endTime** | **int32** | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default &#x3D; 0 (no limit) | 

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


## DeleteDeploymentContainerLocation

> []TaskStatus DeleteDeploymentContainerLocation(ctx, deploymentContainerLocationId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()

Delete given deployment profile containerLocation (it will delete all the application instances if there are any)



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
	deploymentContainerLocationId := "deploymentContainerLocationId_example" // string | Id of the deployment profile containerLocation to delete
	methodId := int32(56) // int32 | Id for the method of deletion: hard-kill = 1, graceful=2 (optional)
	startTime := int32(56) // int32 | Unix timestamp. When to start deleting the containerLocation and its contents. Default = now (optional)
	endTime := int32(56) // int32 | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default = 0 (no limit) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.DeleteDeploymentContainerLocation(context.Background(), deploymentContainerLocationId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.DeleteDeploymentContainerLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeploymentContainerLocation`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.DeleteDeploymentContainerLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentContainerLocationId** | **string** | Id of the deployment profile containerLocation to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentContainerLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **methodId** | **int32** | Id for the method of deletion: hard-kill &#x3D; 1, graceful&#x3D;2 | 
 **startTime** | **int32** | Unix timestamp. When to start deleting the containerLocation and its contents. Default &#x3D; now | 
 **endTime** | **int32** | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default &#x3D; 0 (no limit) | 

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


## DeleteDeploymentProfile

> []TaskStatus DeleteDeploymentProfile(ctx, deploymentProfileId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()

Delete given deployment profile (it will delete all the application instances if there are any)



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
	deploymentProfileId := "deploymentProfileId_example" // string | Id of the deployment profile to delete
	methodId := int32(56) // int32 | Id for the method of deletion: hard-kill = 1, graceful=2 (optional)
	startTime := int32(56) // int32 | Unix timestamp. When to start deleting the containerLocation and its contents. Default = now (optional)
	endTime := int32(56) // int32 | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default = 0 (no limit) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.DeleteDeploymentProfile(context.Background(), deploymentProfileId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.DeleteDeploymentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeploymentProfile`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.DeleteDeploymentProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentProfileId** | **string** | Id of the deployment profile to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **methodId** | **int32** | Id for the method of deletion: hard-kill &#x3D; 1, graceful&#x3D;2 | 
 **startTime** | **int32** | Unix timestamp. When to start deleting the containerLocation and its contents. Default &#x3D; now | 
 **endTime** | **int32** | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default &#x3D; 0 (no limit) | 

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


## DeleteDeploymentRegion

> []TaskStatus DeleteDeploymentRegion(ctx, deploymentRegionId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()

Delete given deployment profile region (it will delete all the application instances if there are any)



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
	deploymentRegionId := "deploymentRegionId_example" // string | Id of the deployment profile region to delete
	methodId := int32(56) // int32 | Id for the method of deletion: hard-kill = 1, graceful=2 (optional)
	startTime := int32(56) // int32 | Unix timestamp. When to start deleting the containerLocation and its contents. Default = now (optional)
	endTime := int32(56) // int32 | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default = 0 (no limit) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.DeleteDeploymentRegion(context.Background(), deploymentRegionId).MethodId(methodId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.DeleteDeploymentRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeploymentRegion`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.DeleteDeploymentRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | Id of the deployment profile region to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **methodId** | **int32** | Id for the method of deletion: hard-kill &#x3D; 1, graceful&#x3D;2 | 
 **startTime** | **int32** | Unix timestamp. When to start deleting the containerLocation and its contents. Default &#x3D; now | 
 **endTime** | **int32** | Unix timestamp. Time limit in which to delete all remaining game servers immediately (not graceful). Default &#x3D; 0 (no limit) | 

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


## DeleteDeploymentRegionProperty

> DeleteDeploymentRegionProperty(ctx, deploymentRegionId, propertyId).Execute()

Deletes a deployment region property (does not delete the application property)

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region
	propertyId := "propertyId_example" // string | The Id of the deployment region property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentProfileAPI.DeleteDeploymentRegionProperty(context.Background(), deploymentRegionId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.DeleteDeploymentRegionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 
**propertyId** | **string** | The Id of the deployment region property | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRegionPropertyRequest struct via the builder pattern


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


## GetDeploymentProfile

> []DeploymentProfile GetDeploymentProfile(ctx, deploymentProfileId).Execute()

Get the given deployment profile

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
	deploymentProfileId := "deploymentProfileId_example" // string | The Id of the deployment profile

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentProfile(context.Background(), deploymentProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentProfile`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentProfileId** | **string** | The Id of the deployment profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeploymentProfile**](DeploymentProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentProfileSummaries

> []DeploymentProfileSummary GetDeploymentProfileSummaries(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all the deployment profiles you have created

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
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentProfileSummaries(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentProfileSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentProfileSummaries`: []DeploymentProfileSummary
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentProfileSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentProfileSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]DeploymentProfileSummary**](DeploymentProfileSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentProfiles

> []DeploymentProfile GetDeploymentProfiles(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all the deployment profiles you have created

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
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentProfiles(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentProfiles`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]DeploymentProfile**](DeploymentProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRegion

> []DeploymentRegion GetDeploymentRegion(ctx, deploymentRegionId).Execute()

Get the given deployment region

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentRegion(context.Background(), deploymentRegionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentRegion`: []DeploymentRegion
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeploymentRegion**](DeploymentRegion.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRegionProperties

> []DeploymentRegionProperty GetDeploymentRegionProperties(ctx, deploymentRegionId).RANGEDDATA(rANGEDDATA).Execute()

Get a list of deployment region properties

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentRegionProperties(context.Background(), deploymentRegionId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentRegionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentRegionProperties`: []DeploymentRegionProperty
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentRegionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRegionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]DeploymentRegionProperty**](DeploymentRegionProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRegionProperty

> []DeploymentRegionProperty GetDeploymentRegionProperty(ctx, deploymentRegionId, propertyId).Execute()

Get a specific deployment region property

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region
	propertyId := "propertyId_example" // string | The Id of the deployment region property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.GetDeploymentRegionProperty(context.Background(), deploymentRegionId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.GetDeploymentRegionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentRegionProperty`: []DeploymentRegionProperty
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.GetDeploymentRegionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 
**propertyId** | **string** | The Id of the deployment region property | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRegionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DeploymentRegionProperty**](DeploymentRegionProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentContainerCanceldelete

> []DeploymentProfile UpdateDeploymentContainerCanceldelete(ctx, deploymentContainerId).Execute()

Cancel delete given cloud deployment container

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
	deploymentContainerId := "deploymentContainerId_example" // string | Id of the deployment container to stop deleting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentContainerCanceldelete(context.Background(), deploymentContainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentContainerCanceldelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentContainerCanceldelete`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.UpdateDeploymentContainerCanceldelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentContainerId** | **string** | Id of the deployment container to stop deleting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentContainerCanceldeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeploymentProfile**](DeploymentProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentContainerLocationCanceldelete

> UpdateDeploymentContainerLocationCanceldelete(ctx, deploymentContainerLocationId).Execute()

Cancel deletion of given deployment profile containerLocation

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
	deploymentContainerLocationId := "deploymentContainerLocationId_example" // string | Id of the deployment containerLocation to stop deleting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentContainerLocationCanceldelete(context.Background(), deploymentContainerLocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentContainerLocationCanceldelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentContainerLocationId** | **string** | Id of the deployment containerLocation to stop deleting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentContainerLocationCanceldeleteRequest struct via the builder pattern


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


## UpdateDeploymentProfile

> []DeploymentProfile UpdateDeploymentProfile(ctx, deploymentProfileId).DeploymentProfile(deploymentProfile).Execute()

Update given deployment profile

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
	deploymentProfileId := "deploymentProfileId_example" // string | The Id of the deployment profile
	deploymentProfile := *openapiclient.NewDeploymentProfile("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), int32(123), int32(123), int32(1), int32(123), []openapiclient.DeploymentRegion{*openapiclient.NewDeploymentRegion("Id_example", "Name_example", []int32{int32(123)}, int32(123), int32(123))}, int32(123)) // DeploymentProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentProfile(context.Background(), deploymentProfileId).DeploymentProfile(deploymentProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentProfile`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.UpdateDeploymentProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentProfileId** | **string** | The Id of the deployment profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentProfile** | [**DeploymentProfile**](DeploymentProfile.md) |  | 

### Return type

[**[]DeploymentProfile**](DeploymentProfile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentProfileCanceldelete

> UpdateDeploymentProfileCanceldelete(ctx, deploymentProfileId).Execute()

Cancel the deletion of given deployment profile (it will stop deletion but what has been deleted is lost already)

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
	deploymentProfileId := "deploymentProfileId_example" // string | Id of the deployment profile to stop deleting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentProfileCanceldelete(context.Background(), deploymentProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentProfileCanceldelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentProfileId** | **string** | Id of the deployment profile to stop deleting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentProfileCanceldeleteRequest struct via the builder pattern


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


## UpdateDeploymentRegion

> []DeploymentRegion UpdateDeploymentRegion(ctx, deploymentRegionId).DeploymentRegion(deploymentRegion).Execute()

Update given deployment region

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
	deploymentRegionId := "deploymentRegionId_example" // string | Id of the deployment region to update
	deploymentRegion := *openapiclient.NewDeploymentRegion("Id_example", "Name_example", []int32{int32(123)}, int32(123), int32(123)) // DeploymentRegion | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentRegion(context.Background(), deploymentRegionId).DeploymentRegion(deploymentRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentRegion`: []DeploymentRegion
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.UpdateDeploymentRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | Id of the deployment region to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentRegion** | [**DeploymentRegion**](DeploymentRegion.md) |  | 

### Return type

[**[]DeploymentRegion**](DeploymentRegion.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentRegionCanceldelete

> UpdateDeploymentRegionCanceldelete(ctx, deploymentRegionId).Execute()

Cancel deletion of given deployment region (it will stop deletion but what has been deleted is lost already)

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
	deploymentRegionId := int32(56) // int32 | Id of the deployment region to stop deleting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentRegionCanceldelete(context.Background(), deploymentRegionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentRegionCanceldelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **int32** | Id of the deployment region to stop deleting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRegionCanceldeleteRequest struct via the builder pattern


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


## UpdateDeploymentRegionProperty

> []DeploymentRegionProperty UpdateDeploymentRegionProperty(ctx, deploymentRegionId, propertyId).DeploymentRegionProperty(deploymentRegionProperty).Execute()

Update a deployment region property

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
	deploymentRegionId := "deploymentRegionId_example" // string | The Id of the deployment region
	propertyId := "propertyId_example" // string | The Id of the deployment region property
	deploymentRegionProperty := *openapiclient.NewDeploymentRegionProperty("Id_example", "ApplicationPropertyId_example", "PropertyValue_example") // DeploymentRegionProperty | Deployment region property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentProfileAPI.UpdateDeploymentRegionProperty(context.Background(), deploymentRegionId, propertyId).DeploymentRegionProperty(deploymentRegionProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentProfileAPI.UpdateDeploymentRegionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentRegionProperty`: []DeploymentRegionProperty
	fmt.Fprintf(os.Stdout, "Response from `DeploymentProfileAPI.UpdateDeploymentRegionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentRegionId** | **string** | The Id of the deployment region | 
**propertyId** | **string** | The Id of the deployment region property | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRegionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentRegionProperty** | [**DeploymentRegionProperty**](DeploymentRegionProperty.md) | Deployment region property | 

### Return type

[**[]DeploymentRegionProperty**](DeploymentRegionProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

