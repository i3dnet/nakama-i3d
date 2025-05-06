# \ApplicationInstanceAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateEmptyGameServersByApplicationId**](ApplicationInstanceAPI.md#AllocateEmptyGameServersByApplicationId) | **Get** /v3/applicationInstance/game/{applicationId}/empty | Get empty game instance for the given applicationId. Currently, implies the ONLINE UNALLOCATED game server status.
[**AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2**](ApplicationInstanceAPI.md#AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2) | **Get** /v3/applicationInstance/{applicationInstanceId}/log | Get the logs for the application instance (for now only the dependency installer logs)
[**CreateApplicationInstanceDeploy**](ApplicationInstanceAPI.md#CreateApplicationInstanceDeploy) | **Post** /v3/applicationInstance/{applicationInstanceId}/deploy | Install a new application instance build
[**CreateApplicationInstanceFleetHost**](ApplicationInstanceAPI.md#CreateApplicationInstanceFleetHost) | **Post** /v3/applicationInstance/fleet/{fleetId}/host/{hostId} | Create a new application instance entity on given host (bare metal or VM) and for given fleetId
[**CreateApplicationInstanceRestart**](ApplicationInstanceAPI.md#CreateApplicationInstanceRestart) | **Post** /v3/applicationInstance/{applicationInstanceId}/restart | Stop and start a single application instance. Also works if it is already stopped. You can give an stop method in the body, if not given we will use the stop method defined in the application build or application. Default is hard kill
[**CreateApplicationInstanceStart**](ApplicationInstanceAPI.md#CreateApplicationInstanceStart) | **Post** /v3/applicationInstance/{applicationInstanceId}/start | Start a single application instance.
[**CreateApplicationInstanceStop**](ApplicationInstanceAPI.md#CreateApplicationInstanceStop) | **Post** /v3/applicationInstance/{applicationInstanceId}/stop | Stop the application instance given, if no stop method is provided we will use the stop method of the application build or application. The default is hard kill
[**DeleteApplicationInstance**](ApplicationInstanceAPI.md#DeleteApplicationInstance) | **Delete** /v3/applicationInstance/{applicationInstanceId} | Remove an application instance
[**GetApplicationInstance**](ApplicationInstanceAPI.md#GetApplicationInstance) | **Get** /v3/applicationInstance/{applicationInstanceId} | Get a single application instance
[**GetApplicationInstanceApplication**](ApplicationInstanceAPI.md#GetApplicationInstanceApplication) | **Get** /v3/applicationInstance/application/{applicationId} | Get all your application instance for the given applicationId.
[**GetApplicationInstanceGameEmptyRegions**](ApplicationInstanceAPI.md#GetApplicationInstanceGameEmptyRegions) | **Get** /v3/applicationInstance/game/{applicationId}/empty/regions | Get empty game instances per region for the given applicationId. Currently implies the ONLINE UNALLOCATED game server status.
[**GetApplicationInstanceHost**](ApplicationInstanceAPI.md#GetApplicationInstanceHost) | **Get** /v3/applicationInstance/host/{hostId} | Get all your application instances on given host (bare metal or VM).
[**GetApplicationInstanceMetadata**](ApplicationInstanceAPI.md#GetApplicationInstanceMetadata) | **Get** /v3/applicationInstance/{applicationInstanceId}/metadata | Get the metadata of given application instance
[**GetApplicationInstanceStatuses**](ApplicationInstanceAPI.md#GetApplicationInstanceStatuses) | **Get** /v3/applicationInstance/status | Get all statuses of application instances
[**GetApplicationInstanceSummaries**](ApplicationInstanceAPI.md#GetApplicationInstanceSummaries) | **Get** /v3/applicationInstance/summary | Get all your application instances and their main details
[**GetApplicationInstanceTasksByApplicationInstanceId**](ApplicationInstanceAPI.md#GetApplicationInstanceTasksByApplicationInstanceId) | **Get** /v3/applicationInstance/{applicationInstanceId}/task | Get the last 50 tasks of the given application instance
[**GetApplicationInstances**](ApplicationInstanceAPI.md#GetApplicationInstances) | **Get** /v3/applicationInstance | Get all your application instances and their main details
[**UpdateApplicationInstance**](ApplicationInstanceAPI.md#UpdateApplicationInstance) | **Put** /v3/applicationInstance/{applicationInstanceId} | Update given application instance
[**UpdateApplicationInstanceGameEmptyAllocate**](ApplicationInstanceAPI.md#UpdateApplicationInstanceGameEmptyAllocate) | **Put** /v3/applicationInstance/game/{applicationId}/empty/allocate | Allocate an ONLINE application instance and return errors if there are no application instances to allocate
[**UpdateApplicationInstanceStatus**](ApplicationInstanceAPI.md#UpdateApplicationInstanceStatus) | **Put** /v3/applicationInstance/{applicationInstanceId}/status/{status} | Set the online status of a single application instance. Only possible if the application instance status is at least \&quot;offline (2)\&quot;



## AllocateEmptyGameServersByApplicationId

> []ApplicationInstance AllocateEmptyGameServersByApplicationId(ctx, applicationId).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()

Get empty game instance for the given applicationId. Currently, implies the ONLINE UNALLOCATED game server status.

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
	applicationId := "applicationId_example" // string | The ID of the application
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/game/{applicationId}/empty?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`.<br /> More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationId(context.Background(), applicationId).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllocateEmptyGameServersByApplicationId`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllocateEmptyGameServersByApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/game/{applicationId}/empty?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;.&lt;br /&gt; More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

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


## AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2

> []ApplicationInstanceLogModel AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2(ctx, applicationInstanceId).RANGEDDATA(rANGEDDATA).Execute()

Get the logs for the application instance (for now only the dependency installer logs)

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2(context.Background(), applicationInstanceId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2`: []ApplicationInstanceLogModel
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.AllocateEmptyGameServersByApplicationIdWithErrorsArcusV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllocateEmptyGameServersByApplicationIdWithErrorsArcusV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]ApplicationInstanceLogModel**](ApplicationInstanceLogModel.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationInstanceDeploy

> []TaskStatus CreateApplicationInstanceDeploy(ctx, applicationInstanceId).ApplicationBuildId(applicationBuildId).Execute()

Install a new application instance build

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	applicationBuildId := *openapiclient.NewApplicationBuildId("ApplicationBuildId_example") // ApplicationBuildId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.CreateApplicationInstanceDeploy(context.Background(), applicationInstanceId).ApplicationBuildId(applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.CreateApplicationInstanceDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceDeploy`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.CreateApplicationInstanceDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationBuildId** | [**ApplicationBuildId**](ApplicationBuildId.md) |  | 

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


## CreateApplicationInstanceFleetHost

> []ApplicationInstance CreateApplicationInstanceFleetHost(ctx, fleetId, hostId).ApplicationInstance(applicationInstance).Execute()

Create a new application instance entity on given host (bare metal or VM) and for given fleetId



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
	applicationInstance := *openapiclient.NewApplicationInstance("Id_example", "DeploymentEnvironmentId_example", "DeploymentEnvironmentName_example", "FleetId_example", "FleetName_example", int32(123), int32(123), "ApplicationId_example", "ApplicationName_example", int32(123), "ApplicationBuildId_example", "ApplicationBuildName_example", int32(123), "DcLocationName_example", "RegionId_example", "RegionName_example", int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "StartupParams_example", "Executable_example", int32(123), []openapiclient.ApplicationInstanceProperty{*openapiclient.NewApplicationInstanceProperty("Id_example", int32(123), "PropertyKey_example", "PropertyValue_example")}, []openapiclient.ApplicationInstanceIP{*openapiclient.NewApplicationInstanceIP("IpAddress_example", int32(123), int32(123))}, []openapiclient.Label{*openapiclient.NewLabel("Key_example", "Value_example")}, int32(123), int32(123), "LiveHostName_example", "LiveMap_example", "LiveGameVersion_example", int32(123), "LiveRules_example", int32(123), int32(123), int32(123)) // ApplicationInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.CreateApplicationInstanceFleetHost(context.Background(), fleetId, hostId).ApplicationInstance(applicationInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.CreateApplicationInstanceFleetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceFleetHost`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.CreateApplicationInstanceFleetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** | The ID of the fleet | 
**hostId** | **int32** | The ID of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceFleetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationInstance** | [**ApplicationInstance**](ApplicationInstance.md) |  | 

### Return type

[**[]ApplicationInstance**](ApplicationInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationInstanceRestart

> []TaskStatus CreateApplicationInstanceRestart(ctx, applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()

Stop and start a single application instance. Also works if it is already stopped. You can give an stop method in the body, if not given we will use the stop method defined in the application build or application. Default is hard kill

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	applicationInstanceStopMethod := *openapiclient.NewApplicationInstanceStopMethod() // ApplicationInstanceStopMethod | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.CreateApplicationInstanceRestart(context.Background(), applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.CreateApplicationInstanceRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceRestart`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.CreateApplicationInstanceRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationInstanceStopMethod** | [**ApplicationInstanceStopMethod**](ApplicationInstanceStopMethod.md) |  | 

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


## CreateApplicationInstanceStart

> []TaskStatus CreateApplicationInstanceStart(ctx, applicationInstanceId).Execute()

Start a single application instance.

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.CreateApplicationInstanceStart(context.Background(), applicationInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.CreateApplicationInstanceStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceStart`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.CreateApplicationInstanceStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceStartRequest struct via the builder pattern


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


## CreateApplicationInstanceStop

> []TaskStatus CreateApplicationInstanceStop(ctx, applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()

Stop the application instance given, if no stop method is provided we will use the stop method of the application build or application. The default is hard kill

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	applicationInstanceStopMethod := *openapiclient.NewApplicationInstanceStopMethod() // ApplicationInstanceStopMethod | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.CreateApplicationInstanceStop(context.Background(), applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.CreateApplicationInstanceStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceStop`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.CreateApplicationInstanceStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationInstanceStopMethod** | [**ApplicationInstanceStopMethod**](ApplicationInstanceStopMethod.md) |  | 

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


## DeleteApplicationInstance

> DeleteApplicationInstance(ctx, applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()

Remove an application instance

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	applicationInstanceStopMethod := *openapiclient.NewApplicationInstanceStopMethod() // ApplicationInstanceStopMethod | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationInstanceAPI.DeleteApplicationInstance(context.Background(), applicationInstanceId).ApplicationInstanceStopMethod(applicationInstanceStopMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.DeleteApplicationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationInstanceStopMethod** | [**ApplicationInstanceStopMethod**](ApplicationInstanceStopMethod.md) |  | 

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


## GetApplicationInstance

> []ApplicationInstance GetApplicationInstance(ctx, applicationInstanceId).Execute()

Get a single application instance

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstance(context.Background(), applicationInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstance`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetApplicationInstanceApplication

> []ApplicationInstance GetApplicationInstanceApplication(ctx, applicationId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()

Get all your application instance for the given applicationId.

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
	applicationId := "applicationId_example" // string | The ID of the application
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/application/{applicationId}?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request<br /> instead. More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, You can only use AND statement You can also search inside nested<br /> objects like `metadata.key='testing' and metadata.value='test'` You can only search on exact matches.<br /> Warning: the `filters` query parameter is ignored when passing a `PAGE-TOKEN` and the original `filters` string will be used for that request<br /> instead.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	sort := "sort_example" // string | Sort expression to sort your results in the order that you like E.G. `sort=fleetId ASC, applicationId DESC` Be aware that all filters<br /> and sorting should be url encoded.<br /> Warning: the `sort` query parameter is ignored when passing a `PAGE-TOKEN` and the original `sort` string will be used for that request instead<br /> Warning: sorting on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA: <span style=\"text-decoration: line-through\">start=0,</span>results=25<br /> Warning: the `start` variable of the `RANGED-DATA` object has been deprecated for this endpoint. It is not possible to supply a start, see<br /> `PAGE-TOKEN` header for a replacement (optional)
	pAGETOKEN := "pAGETOKEN_example" // string | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint<br /> without the `PAGE-TOKEN` you will retrieve the first page of search results. If a `PAGE-TOKEN` is returned, a next page may be requested by<br /> supplying that token in the next request using the same `PAGE-TOKEN` key and value in the request header. If no `PAGE-TOKEN` is returned, the last<br /> page has been reached and no further data can be requested. A returned `PAGE-TOKEN` is valid until the date specified in the `Expires` response<br /> header. Example: `PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)` (truncated for readability)<br /> Expires header example: `Expires: Fri, 08 Jan 2021 09:10:29 GMT` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceApplication(context.Background(), applicationId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceApplication`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/application/{applicationId}?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request&lt;br /&gt; instead. More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, You can only use AND statement You can also search inside nested&lt;br /&gt; objects like &#x60;metadata.key&#x3D;&#39;testing&#39; and metadata.value&#x3D;&#39;test&#39;&#x60; You can only search on exact matches.&lt;br /&gt; Warning: the &#x60;filters&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;filters&#x60; string will be used for that request&lt;br /&gt; instead.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **sort** | **string** | Sort expression to sort your results in the order that you like E.G. &#x60;sort&#x3D;fleetId ASC, applicationId DESC&#x60; Be aware that all filters&lt;br /&gt; and sorting should be url encoded.&lt;br /&gt; Warning: the &#x60;sort&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;sort&#x60; string will be used for that request instead&lt;br /&gt; Warning: sorting on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
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


## GetApplicationInstanceGameEmptyRegions

> []ApplicationInstanceByRegion GetApplicationInstanceGameEmptyRegions(ctx, applicationId).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()

Get empty game instances per region for the given applicationId. Currently implies the ONLINE UNALLOCATED game server status.

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
	applicationId := "applicationId_example" // string | The ID of the application
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/game/{applicationId}/empty/regions?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`.<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`.<br /> More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceGameEmptyRegions(context.Background(), applicationId).Labels(labels).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceGameEmptyRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceGameEmptyRegions`: []ApplicationInstanceByRegion
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceGameEmptyRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceGameEmptyRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/game/{applicationId}/empty/regions?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;.&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;.&lt;br /&gt; More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]ApplicationInstanceByRegion**](ApplicationInstanceByRegion.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceHost

> []ApplicationInstance GetApplicationInstanceHost(ctx, hostId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()

Get all your application instances on given host (bare metal or VM).



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
	hostId := int32(56) // int32 | The ID of the host
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded.<br /> E.g. `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/host/{hostId}?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request<br /> instead. More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, only the `AND` statement is supported at this time. You can also<br /> search inside nested objects like `metadata.key=\"testing\"` and `metadata.value=\"test\"` You can only search on exact matches.<br /> Warning: the `filters` query parameter is ignored when passing a `PAGE-TOKEN` and the original `filters` string will be used for that request<br /> instead.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`,<br /> `numPlayersMax`, `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	sort := "sort_example" // string | Sort expression to sort your results in the order that you like E.G. `sort=fleetId ASC, applicationId DESC` Be aware that all<br /> filters and sorting should be url encoded.<br /> Warning: the `sort` query parameter is ignored when passing a `PAGE-TOKEN` and the original `sort` string will be used for that request instead<br /> Warning: sorting on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`,<br /> `numPlayersMax`, `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA: <span style=\"text-decoration: line-through\">start=0,</span>results=25<br /> Warning: the `start` variable of the `RANGED-DATA` object has been deprecated for this endpoint. It is not possible to supply a start, see<br /> `PAGE-TOKEN` header for a replacement (optional)
	pAGETOKEN := "pAGETOKEN_example" // string | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint<br /> without the `PAGE-TOKEN` you will retrieve the first page of search results. If a `PAGE-TOKEN` is returned, a next page may be requested by<br /> supplying that token in the next request using the same `PAGE-TOKEN` key and value in the request header. If no `PAGE-TOKEN` is returned, the<br /> last page has been reached and no further data can be requested. A returned `PAGE-TOKEN` is valid until the date specified in the `Expires`<br /> response header. Example: `PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)` (truncated for readability) Expires header<br /> example: `Expires: Fri, 08 Jan 2021 09:10:29 GMT` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceHost(context.Background(), hostId).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceHost`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int32** | The ID of the host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded.&lt;br /&gt; E.g. &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/host/{hostId}?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request&lt;br /&gt; instead. More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, only the &#x60;AND&#x60; statement is supported at this time. You can also&lt;br /&gt; search inside nested objects like &#x60;metadata.key&#x3D;\&quot;testing\&quot;&#x60; and &#x60;metadata.value&#x3D;\&quot;test\&quot;&#x60; You can only search on exact matches.&lt;br /&gt; Warning: the &#x60;filters&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;filters&#x60; string will be used for that request&lt;br /&gt; instead.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;,&lt;br /&gt; &#x60;numPlayersMax&#x60;, &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **sort** | **string** | Sort expression to sort your results in the order that you like E.G. &#x60;sort&#x3D;fleetId ASC, applicationId DESC&#x60; Be aware that all&lt;br /&gt; filters and sorting should be url encoded.&lt;br /&gt; Warning: the &#x60;sort&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;sort&#x60; string will be used for that request instead&lt;br /&gt; Warning: sorting on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;,&lt;br /&gt; &#x60;numPlayersMax&#x60;, &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA: &lt;span style&#x3D;\&quot;text-decoration: line-through\&quot;&gt;start&#x3D;0,&lt;/span&gt;results&#x3D;25&lt;br /&gt; Warning: the &#x60;start&#x60; variable of the &#x60;RANGED-DATA&#x60; object has been deprecated for this endpoint. It is not possible to supply a start, see&lt;br /&gt; &#x60;PAGE-TOKEN&#x60; header for a replacement | 
 **pAGETOKEN** | **string** | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint&lt;br /&gt; without the &#x60;PAGE-TOKEN&#x60; you will retrieve the first page of search results. If a &#x60;PAGE-TOKEN&#x60; is returned, a next page may be requested by&lt;br /&gt; supplying that token in the next request using the same &#x60;PAGE-TOKEN&#x60; key and value in the request header. If no &#x60;PAGE-TOKEN&#x60; is returned, the&lt;br /&gt; last page has been reached and no further data can be requested. A returned &#x60;PAGE-TOKEN&#x60; is valid until the date specified in the &#x60;Expires&#x60;&lt;br /&gt; response header. Example: &#x60;PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)&#x60; (truncated for readability) Expires header&lt;br /&gt; example: &#x60;Expires: Fri, 08 Jan 2021 09:10:29 GMT&#x60; | 

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


## GetApplicationInstanceMetadata

> MetadataCollection GetApplicationInstanceMetadata(ctx, applicationInstanceId).Execute()

Get the metadata of given application instance

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceMetadata(context.Background(), applicationInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceMetadata`: MetadataCollection
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataCollection**](MetadataCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceStatuses

> []ApplicationInstanceStatus GetApplicationInstanceStatuses(ctx).Execute()

Get all statuses of application instances

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
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceStatuses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceStatuses`: []ApplicationInstanceStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceStatusesRequest struct via the builder pattern


### Return type

[**[]ApplicationInstanceStatus**](ApplicationInstanceStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceSummaries

> []ApplicationInstanceSummary GetApplicationInstanceSummaries(ctx).Labels(labels).Filters(filters).Sort(sort).Fields(fields).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()

Get all your application instances and their main details

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
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/summary?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request<br /> instead. More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, only the `AND` statement is supported at this time. You can also<br /> search inside nested objects like `metadata.key=\"testing\"` and `metadata.value=\"test\"` You can only search on exact matches.<br /> Warning: the `filters` query parameter is ignored when passing a `PAGE-TOKEN` and the original `filters` string will be used for that request<br /> instead.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	sort := "sort_example" // string | Sort expression to sort your results in the order that you like E.G. `sort=fleetId ASC, applicationId DESC` Be aware that all filters<br /> and sorting should be url encoded.<br /> Warning: the `sort` query parameter is ignored when passing a `PAGE-TOKEN` and the original `sort` string will be used for that request instead<br /> Warning: sorting on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	fields := "fields_example" // string | Fields return in response, only fields available in model can be filter out (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA: <span style=\"text-decoration: line-through\">start=0,</span>results=25<br /> Warning: the `start` variable of the `RANGED-DATA` object has been deprecated for this endpoint. It is not possible to supply a start, see<br /> `PAGE-TOKEN` header for a replacement (optional)
	pAGETOKEN := "pAGETOKEN_example" // string | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint<br /> without the `PAGE-TOKEN` you will retrieve the first page of search results. If a `PAGE-TOKEN` is returned, a next page may be requested by<br /> supplying that token in the next request using the same `PAGE-TOKEN` key and value in the request header. If no `PAGE-TOKEN` is returned, the last<br /> page has been reached and no further data can be requested. A returned `PAGE-TOKEN` is valid until the date specified in the `Expires` response<br /> header. Example: `PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)` (truncated for readability)<br /> Expires header example: `Expires: Fri, 08 Jan 2021 09:10:29 GMT` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceSummaries(context.Background()).Labels(labels).Filters(filters).Sort(sort).Fields(fields).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceSummaries`: []ApplicationInstanceSummary
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/summary?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request&lt;br /&gt; instead. More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, only the &#x60;AND&#x60; statement is supported at this time. You can also&lt;br /&gt; search inside nested objects like &#x60;metadata.key&#x3D;\&quot;testing\&quot;&#x60; and &#x60;metadata.value&#x3D;\&quot;test\&quot;&#x60; You can only search on exact matches.&lt;br /&gt; Warning: the &#x60;filters&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;filters&#x60; string will be used for that request&lt;br /&gt; instead.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **sort** | **string** | Sort expression to sort your results in the order that you like E.G. &#x60;sort&#x3D;fleetId ASC, applicationId DESC&#x60; Be aware that all filters&lt;br /&gt; and sorting should be url encoded.&lt;br /&gt; Warning: the &#x60;sort&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;sort&#x60; string will be used for that request instead&lt;br /&gt; Warning: sorting on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **fields** | **string** | Fields return in response, only fields available in model can be filter out | 
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA: &lt;span style&#x3D;\&quot;text-decoration: line-through\&quot;&gt;start&#x3D;0,&lt;/span&gt;results&#x3D;25&lt;br /&gt; Warning: the &#x60;start&#x60; variable of the &#x60;RANGED-DATA&#x60; object has been deprecated for this endpoint. It is not possible to supply a start, see&lt;br /&gt; &#x60;PAGE-TOKEN&#x60; header for a replacement | 
 **pAGETOKEN** | **string** | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint&lt;br /&gt; without the &#x60;PAGE-TOKEN&#x60; you will retrieve the first page of search results. If a &#x60;PAGE-TOKEN&#x60; is returned, a next page may be requested by&lt;br /&gt; supplying that token in the next request using the same &#x60;PAGE-TOKEN&#x60; key and value in the request header. If no &#x60;PAGE-TOKEN&#x60; is returned, the last&lt;br /&gt; page has been reached and no further data can be requested. A returned &#x60;PAGE-TOKEN&#x60; is valid until the date specified in the &#x60;Expires&#x60; response&lt;br /&gt; header. Example: &#x60;PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)&#x60; (truncated for readability)&lt;br /&gt; Expires header example: &#x60;Expires: Fri, 08 Jan 2021 09:10:29 GMT&#x60; | 

### Return type

[**[]ApplicationInstanceSummary**](ApplicationInstanceSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceTasksByApplicationInstanceId

> []TaskStatus GetApplicationInstanceTasksByApplicationInstanceId(ctx, applicationInstanceId).Execute()

Get the last 50 tasks of the given application instance

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstanceTasksByApplicationInstanceId(context.Background(), applicationInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstanceTasksByApplicationInstanceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTasksByApplicationInstanceId`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstanceTasksByApplicationInstanceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTasksByApplicationInstanceIdRequest struct via the builder pattern


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


## GetApplicationInstances

> []ApplicationInstance GetApplicationInstances(ctx).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()

Get all your application instances and their main details

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
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> Warning: the `labels` query parameter is ignored when passing a `PAGE-TOKEN` and the original `labels` string will be used for that request<br /> instead. More information on the use of labels can be found in the<br /> <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">label documentation</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, only the `AND` statement is supported at this time. You can also<br /> search inside nested objects like `metadata.key=\"testing\"` and `metadata.value=\"test\"` You can only search on exact matches.<br /> Warning: the `filters` query parameter is ignored when passing a `PAGE-TOKEN` and the original `filters` string will be used for that request<br /> instead.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	sort := "sort_example" // string | Sort expression to sort your results in the order that you like E.G. `sort=fleetId ASC, applicationId DESC` Be aware that all filters<br /> and sorting should be url encoded.<br /> Warning: the `sort` query parameter is ignored when passing a `PAGE-TOKEN` and the original `sort` string will be used for that request instead<br /> Warning: sorting on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA: <span style=\"text-decoration: line-through\">start=0,</span>results=25<br /> Warning: the `start` variable of the `RANGED-DATA` object has been deprecated for this endpoint. It is not possible to supply a start, see<br /> `PAGE-TOKEN` header for a replacement (optional)
	pAGETOKEN := "pAGETOKEN_example" // string | This token will be provided by this endpoint to allow for pagination functionality. By performing a request to this endpoint<br /> without the `PAGE-TOKEN` you will retrieve the first page of search results. If a `PAGE-TOKEN` is returned, a next page may be requested by<br /> supplying that token in the next request using the same `PAGE-TOKEN` key and value in the request header. If no `PAGE-TOKEN` is returned, the last<br /> page has been reached and no further data can be requested. A returned `PAGE-TOKEN` is valid until the date specified in the `Expires` response<br /> header. Example: `PAGE-TOKEN: eyJjb25kaXRpb25zIjp7InVzZXJJZCI6NjY2ODI1fSwibGFiZWxz(...)` (truncated for readability)<br /> Expires header example: `Expires: Fri, 08 Jan 2021 09:10:29 GMT` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.GetApplicationInstances(context.Background()).Labels(labels).Filters(filters).Sort(sort).RANGEDDATA(rANGEDDATA).PAGETOKEN(pAGETOKEN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.GetApplicationInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstances`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.GetApplicationInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; Warning: the &#x60;labels&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;labels&#x60; string will be used for that request&lt;br /&gt; instead. More information on the use of labels can be found in the&lt;br /&gt; &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;label documentation&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, only the &#x60;AND&#x60; statement is supported at this time. You can also&lt;br /&gt; search inside nested objects like &#x60;metadata.key&#x3D;\&quot;testing\&quot;&#x60; and &#x60;metadata.value&#x3D;\&quot;test\&quot;&#x60; You can only search on exact matches.&lt;br /&gt; Warning: the &#x60;filters&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;filters&#x60; string will be used for that request&lt;br /&gt; instead.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **sort** | **string** | Sort expression to sort your results in the order that you like E.G. &#x60;sort&#x3D;fleetId ASC, applicationId DESC&#x60; Be aware that all filters&lt;br /&gt; and sorting should be url encoded.&lt;br /&gt; Warning: the &#x60;sort&#x60; query parameter is ignored when passing a &#x60;PAGE-TOKEN&#x60; and the original &#x60;sort&#x60; string will be used for that request instead&lt;br /&gt; Warning: sorting on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
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


## UpdateApplicationInstance

> []ApplicationInstance UpdateApplicationInstance(ctx, applicationInstanceId).ApplicationInstance(applicationInstance).Execute()

Update given application instance

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	applicationInstance := *openapiclient.NewApplicationInstance("Id_example", "DeploymentEnvironmentId_example", "DeploymentEnvironmentName_example", "FleetId_example", "FleetName_example", int32(123), int32(123), "ApplicationId_example", "ApplicationName_example", int32(123), "ApplicationBuildId_example", "ApplicationBuildName_example", int32(123), "DcLocationName_example", "RegionId_example", "RegionName_example", int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), "StartupParams_example", "Executable_example", int32(123), []openapiclient.ApplicationInstanceProperty{*openapiclient.NewApplicationInstanceProperty("Id_example", int32(123), "PropertyKey_example", "PropertyValue_example")}, []openapiclient.ApplicationInstanceIP{*openapiclient.NewApplicationInstanceIP("IpAddress_example", int32(123), int32(123))}, []openapiclient.Label{*openapiclient.NewLabel("Key_example", "Value_example")}, int32(123), int32(123), "LiveHostName_example", "LiveMap_example", "LiveGameVersion_example", int32(123), "LiveRules_example", int32(123), int32(123), int32(123)) // ApplicationInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.UpdateApplicationInstance(context.Background(), applicationInstanceId).ApplicationInstance(applicationInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.UpdateApplicationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationInstance`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.UpdateApplicationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationInstance** | [**ApplicationInstance**](ApplicationInstance.md) |  | 

### Return type

[**[]ApplicationInstance**](ApplicationInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationInstanceGameEmptyAllocate

> []ApplicationInstance UpdateApplicationInstanceGameEmptyAllocate(ctx, applicationId).Labels(labels).Filters(filters).MetadataCollection(metadataCollection).Execute()

Allocate an ONLINE application instance and return errors if there are no application instances to allocate



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
	applicationId := "applicationId_example" // string | The ID of the application
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded.<br /> E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/applicationInstance/game/{applicationId}/empty/allocate?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`<br /> More information on the use of labels can be found <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">here</a>. (optional)
	filters := "filters_example" // string | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like<br /> this E.G. `regionId=123` or multiple filters: `regionId=123 and fleetId=456`, only the `AND` statement is supported at this time. You can only search on exact matches.<br /> Warning: filtering on the following parameters is currently not supported and will give an error:<br /> `autoRestart`, `liveGameVersion`, `liveHostName`, `liveMap`, `liveRules`, `manuallyDeployed`, `markedForDeletion`, `numPlayers`, `numPlayersMax`,<br /> `properties`, `ipAddress`, `labelReadOnly`, `labelPublic`, `metadata`, `pid`, `pidChangedAt`, `startedAt`, `status`, `stoppedAt`, `updatedAt` (optional)
	metadataCollection := *openapiclient.NewMetadataCollection() // MetadataCollection | Custom key/value pairs that form miscellaneous metadata to be stored alongside all game servers selected in this call. Metadata is merged with existing metadata in the game server. Key/value pairs are only deleted if you submit a key with a `null` value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.UpdateApplicationInstanceGameEmptyAllocate(context.Background(), applicationId).Labels(labels).Filters(filters).MetadataCollection(metadataCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.UpdateApplicationInstanceGameEmptyAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationInstanceGameEmptyAllocate`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.UpdateApplicationInstanceGameEmptyAllocate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationInstanceGameEmptyAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded.&lt;br /&gt; E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/applicationInstance/game/{applicationId}/empty/allocate?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;&lt;br /&gt; More information on the use of labels can be found &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;here&lt;/a&gt;. | 
 **filters** | **string** | Filter expression can be used to apply more specific search on most parameters in the application instance, it is written like&lt;br /&gt; this E.G. &#x60;regionId&#x3D;123&#x60; or multiple filters: &#x60;regionId&#x3D;123 and fleetId&#x3D;456&#x60;, only the &#x60;AND&#x60; statement is supported at this time. You can only search on exact matches.&lt;br /&gt; Warning: filtering on the following parameters is currently not supported and will give an error:&lt;br /&gt; &#x60;autoRestart&#x60;, &#x60;liveGameVersion&#x60;, &#x60;liveHostName&#x60;, &#x60;liveMap&#x60;, &#x60;liveRules&#x60;, &#x60;manuallyDeployed&#x60;, &#x60;markedForDeletion&#x60;, &#x60;numPlayers&#x60;, &#x60;numPlayersMax&#x60;,&lt;br /&gt; &#x60;properties&#x60;, &#x60;ipAddress&#x60;, &#x60;labelReadOnly&#x60;, &#x60;labelPublic&#x60;, &#x60;metadata&#x60;, &#x60;pid&#x60;, &#x60;pidChangedAt&#x60;, &#x60;startedAt&#x60;, &#x60;status&#x60;, &#x60;stoppedAt&#x60;, &#x60;updatedAt&#x60; | 
 **metadataCollection** | [**MetadataCollection**](MetadataCollection.md) | Custom key/value pairs that form miscellaneous metadata to be stored alongside all game servers selected in this call. Metadata is merged with existing metadata in the game server. Key/value pairs are only deleted if you submit a key with a &#x60;null&#x60; value | 

### Return type

[**[]ApplicationInstance**](ApplicationInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationInstanceStatus

> []ApplicationInstance UpdateApplicationInstanceStatus(ctx, applicationInstanceId, status).Execute()

Set the online status of a single application instance. Only possible if the application instance status is at least \"offline (2)\"

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
	applicationInstanceId := "applicationInstanceId_example" // string | The ID of the application instance
	status := int32(56) // int32 | Possible values:<br /> - 3: starting<br /> - 4: online<br /> - 5: allocated (only possible if previous status is online)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceAPI.UpdateApplicationInstanceStatus(context.Background(), applicationInstanceId, status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceAPI.UpdateApplicationInstanceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationInstanceStatus`: []ApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceAPI.UpdateApplicationInstanceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationInstanceId** | **string** | The ID of the application instance | 
**status** | **int32** | Possible values:&lt;br /&gt; - 3: starting&lt;br /&gt; - 4: online&lt;br /&gt; - 5: allocated (only possible if previous status is online) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationInstanceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

