# \DeploymentEnvironmentAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentEnvironment**](DeploymentEnvironmentAPI.md#CreateDeploymentEnvironment) | **Post** /v3/deploymentEnvironment | Create a deployment environment
[**DeleteDeploymentEnvironment**](DeploymentEnvironmentAPI.md#DeleteDeploymentEnvironment) | **Delete** /v3/deploymentEnvironment/{deploymentEnvironmentId} | Delete given deployment environment
[**GetDeploymentEnvironment**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironment) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId} | Get the details of the given deployment environment
[**GetDeploymentEnvironmentApplicationBuildInUses**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentApplicationBuildInUses) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/applicationBuild/inUse | Get the list of in use application builds for a given deployment environment Id
[**GetDeploymentEnvironmentApplicationFleets**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentApplicationFleets) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/application/{applicationId}/fleet | Get the list of fleet in a deployment environment, having a certain application of type game
[**GetDeploymentEnvironmentApplicationInUses**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentApplicationInUses) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/application/inUse | Get the list of in use applications for given deployment environment Id
[**GetDeploymentEnvironmentApplicationInstances**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentApplicationInstances) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/applicationInstance | Get all the application instances for the given deployment environment
[**GetDeploymentEnvironmentDeploymentProfiles**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentDeploymentProfiles) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/deploymentProfile | Get all the cloud deployment profiles you have created
[**GetDeploymentEnvironmentFleets**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentFleets) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/fleet | Get all the active fleets that belong to the given deployment environment.
[**GetDeploymentEnvironmentHosts**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironmentHosts) | **Get** /v3/deploymentEnvironment/{deploymentEnvironmentId}/host | Get all your hosts for given deployment environment
[**GetDeploymentEnvironments**](DeploymentEnvironmentAPI.md#GetDeploymentEnvironments) | **Get** /v3/deploymentEnvironment | Get the details of all your deployment environment
[**UpdateDeploymentEnvironment**](DeploymentEnvironmentAPI.md#UpdateDeploymentEnvironment) | **Put** /v3/deploymentEnvironment/{deploymentEnvironmentId} | Update given deployment environment



## CreateDeploymentEnvironment

> []DeploymentEnvironment CreateDeploymentEnvironment(ctx).DeploymentEnvironment(deploymentEnvironment).Execute()

Create a deployment environment

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
	deploymentEnvironment := *openapiclient.NewDeploymentEnvironment("Id_example", "Name_example", int32(123)) // DeploymentEnvironment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.CreateDeploymentEnvironment(context.Background()).DeploymentEnvironment(deploymentEnvironment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.CreateDeploymentEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentEnvironment`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.CreateDeploymentEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentEnvironment** | [**DeploymentEnvironment**](DeploymentEnvironment.md) |  | 

### Return type

[**[]DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentEnvironment

> DeleteDeploymentEnvironment(ctx, deploymentEnvironmentId).Execute()

Delete given deployment environment

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentEnvironmentAPI.DeleteDeploymentEnvironment(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.DeleteDeploymentEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentEnvironmentRequest struct via the builder pattern


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


## GetDeploymentEnvironment

> []DeploymentEnvironment GetDeploymentEnvironment(ctx, deploymentEnvironmentId).Execute()

Get the details of the given deployment environment

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironment(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironment`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentEnvironmentApplicationBuildInUses

> []ApplicationBuild GetDeploymentEnvironmentApplicationBuildInUses(ctx, deploymentEnvironmentId).Execute()

Get the list of in use application builds for a given deployment environment Id

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationBuildInUses(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationBuildInUses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentApplicationBuildInUses`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationBuildInUses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentApplicationBuildInUsesRequest struct via the builder pattern


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


## GetDeploymentEnvironmentApplicationFleets

> []Fleet GetDeploymentEnvironmentApplicationFleets(ctx, deploymentEnvironmentId, applicationId).Execute()

Get the list of fleet in a deployment environment, having a certain application of type game

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment
	applicationId := "applicationId_example" // string | The Id of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationFleets(context.Background(), deploymentEnvironmentId, applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentApplicationFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 
**applicationId** | **string** | The Id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentApplicationFleetsRequest struct via the builder pattern


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


## GetDeploymentEnvironmentApplicationInUses

> []Application GetDeploymentEnvironmentApplicationInUses(ctx, deploymentEnvironmentId).Execute()

Get the list of in use applications for given deployment environment Id

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInUses(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInUses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentApplicationInUses`: []Application
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInUses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentApplicationInUsesRequest struct via the builder pattern


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


## GetDeploymentEnvironmentApplicationInstances

> []ApplicationInstance GetDeploymentEnvironmentApplicationInstances(ctx, deploymentEnvironmentId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()

Get all the application instances for the given deployment environment



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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/deploymentEnvironment/{deploymentEnvironmentId}/applicationInstance?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request instead.<br /> More information on the use of labels can be found <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">here</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, only the `AND` statement is supported at this time. You can also<br /> search inside nested objects like `metaData.key=\"testing\"` and `metaData.value=\"test\"` You can only search on exact matches.<br /> Warning: the `filters` query parameter is ignored when passing a `PAGE-TOKEN` and the original `filters` string will be used for that request<br /> instead.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	sort := "sort_example" // string | Sort expression to sort your results in the order that you like E.G. \"sort=fleetId=ASC and applicationId DESC\" Be aware that all filters<br /> and sorting should be url encoded.<br /> Warning: the `sort` query parameter is ignored when passing a `PAGE-TOKEN` and the original `sort` string will be used for that request instead<br /> Warning: sorting on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA: <span style=\"text-decoration: line-through\">start=0,</span>results=25<br /> Warning: the `start` variable of the `RANGED-DATA` object has been deprecated for this endpoint. It is not possible to supply a start, see<br /> `PAGE-TOKEN` header for a replacement (optional)
	pAGETOKEN := "pAGETOKEN_example" // string | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint<br /> without the `PAGE-TOKEN` you will retrieve the first page of search results. If a `PAGE-TOKEN` is returned, a next page may be requested by<br /> supplying that token in the next request using the same `PAGE-TOKEN` key and value in the request header. If no `PAGE-TOKEN` is returned, the last<br /> page has been reached and no further data can be requested. A returned `PAGE-TOKEN` is valid until the date specified in the `Expires` response<br /> header. Example: `PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)` (truncated for readability)<br /> Expires header example: `Expires: Fri, 08 Jan 2021 09:10:29 GMT` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInstances(context.Background(), deploymentEnvironmentId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentApplicationInstances`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentApplicationInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentApplicationInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/deploymentEnvironment/{deploymentEnvironmentId}/applicationInstance?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request instead.&lt;br /&gt; More information on the use of labels can be found &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;here&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, only the &#x60;AND&#x60; statement is supported at this time. You can also&lt;br /&gt; search inside nested objects like &#x60;metaData.key&#x3D;\&quot;testing\&quot;&#x60; and &#x60;metaData.value&#x3D;\&quot;test\&quot;&#x60; You can only search on exact matches.&lt;br /&gt; Warning: the &#x60;filters&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;filters&#x60; string will be used for that request&lt;br /&gt; instead.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **sort** | **string** | Sort expression to sort your results in the order that you like E.G. \&quot;sort&#x3D;fleetId&#x3D;ASC and applicationId DESC\&quot; Be aware that all filters&lt;br /&gt; and sorting should be url encoded.&lt;br /&gt; Warning: the &#x60;sort&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;sort&#x60; string will be used for that request instead&lt;br /&gt; Warning: sorting on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA: &lt;span style&#x3D;\&quot;text-decoration: line-through\&quot;&gt;start&#x3D;0,&lt;/span&gt;results&#x3D;25&lt;br /&gt; Warning: the &#x60;start&#x60; variable of the &#x60;RANGED-DATA&#x60; object has been deprecated for this endpoint. It is not possible to supply a start, see&lt;br /&gt; &#x60;PAGE-TOKEN&#x60; header for a replacement | 
 **pAGETOKEN** | **string** | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint&lt;br /&gt; without the &#x60;PAGE-TOKEN&#x60; you will retrieve the first page of search results. If a &#x60;PAGE-TOKEN&#x60; is returned, a next page may be requested by&lt;br /&gt; supplying that token in the next request using the same &#x60;PAGE-TOKEN&#x60; key and value in the request header. If no &#x60;PAGE-TOKEN&#x60; is returned, the last&lt;br /&gt; page has been reached and no further data can be requested. A returned &#x60;PAGE-TOKEN&#x60; is valid until the date specified in the &#x60;Expires&#x60; response&lt;br /&gt; header. Example: &#x60;PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)&#x60; (truncated for readability)&lt;br /&gt; Expires header example: &#x60;Expires: Fri, 08 Jan 2021 09:10:29 GMT&#x60; | 

### Return type

[**[]ApplicationInstance**](ApplicationInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentEnvironmentDeploymentProfiles

> []DeploymentProfile GetDeploymentEnvironmentDeploymentProfiles(ctx, deploymentEnvironmentId).Execute()

Get all the cloud deployment profiles you have created

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentDeploymentProfiles(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentDeploymentProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentDeploymentProfiles`: []DeploymentProfile
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentDeploymentProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentDeploymentProfilesRequest struct via the builder pattern


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


## GetDeploymentEnvironmentFleets

> []Fleet GetDeploymentEnvironmentFleets(ctx, deploymentEnvironmentId).Execute()

Get all the active fleets that belong to the given deployment environment.

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentFleets(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentFleetsRequest struct via the builder pattern


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


## GetDeploymentEnvironmentHosts

> []Host GetDeploymentEnvironmentHosts(ctx, deploymentEnvironmentId).Execute()

Get all your hosts for given deployment environment

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironmentHosts(context.Background(), deploymentEnvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironmentHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironmentHosts`: []Host
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironmentHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentHostsRequest struct via the builder pattern


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


## GetDeploymentEnvironments

> []DeploymentEnvironment GetDeploymentEnvironments(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get the details of all your deployment environment

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
	resp, r, err := apiClient.DeploymentEnvironmentAPI.GetDeploymentEnvironments(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.GetDeploymentEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentEnvironments`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.GetDeploymentEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentEnvironment

> []DeploymentEnvironment UpdateDeploymentEnvironment(ctx, deploymentEnvironmentId).DeploymentEnvironment(deploymentEnvironment).Execute()

Update given deployment environment

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
	deploymentEnvironmentId := "deploymentEnvironmentId_example" // string | The Id of the deployment environment
	deploymentEnvironment := *openapiclient.NewDeploymentEnvironment("Id_example", "Name_example", int32(123)) // DeploymentEnvironment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentEnvironmentAPI.UpdateDeploymentEnvironment(context.Background(), deploymentEnvironmentId).DeploymentEnvironment(deploymentEnvironment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentEnvironmentAPI.UpdateDeploymentEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentEnvironment`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentEnvironmentAPI.UpdateDeploymentEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentEnvironment** | [**DeploymentEnvironment**](DeploymentEnvironment.md) |  | 

### Return type

[**[]DeploymentEnvironment**](DeploymentEnvironment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

