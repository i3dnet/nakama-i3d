# \ApplicationAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationAPI.md#CreateApplication) | **Post** /v3/application | Create an application
[**CreateApplicationProperty**](ApplicationAPI.md#CreateApplicationProperty) | **Post** /v3/application/{applicationId}/property | Create a property of given applicationId
[**DeleteApplication**](ApplicationAPI.md#DeleteApplication) | **Delete** /v3/application/{applicationId} | Delete a given application
[**DeleteApplicationProperty**](ApplicationAPI.md#DeleteApplicationProperty) | **Delete** /v3/application/{applicationId}/property/{propertyId} | Delete the given property of the applicationId
[**GetApplication**](ApplicationAPI.md#GetApplication) | **Get** /v3/application/{applicationId} | Get the details of the given application
[**GetApplicationApplicationBuildInUses**](ApplicationAPI.md#GetApplicationApplicationBuildInUses) | **Get** /v3/application/{applicationId}/applicationBuild/inUse | Get all the application builds which have application instances created for the given application
[**GetApplicationApplicationBuilds**](ApplicationAPI.md#GetApplicationApplicationBuilds) | **Get** /v3/application/{applicationId}/applicationBuild | Get all the application builds created for the given application
[**GetApplicationDeploymentEnvironments**](ApplicationAPI.md#GetApplicationDeploymentEnvironments) | **Get** /v3/application/{applicationId}/deploymentEnvironment | Get the list of deployment environment using given applicationId
[**GetApplicationDeploymentTemplates**](ApplicationAPI.md#GetApplicationDeploymentTemplates) | **Get** /v3/application/{applicationId}/deploymentTemplate | Get a list of deployment templates for a given applicationId
[**GetApplicationFleets**](ApplicationAPI.md#GetApplicationFleets) | **Get** /v3/application/{applicationId}/fleet | Get the list of fleets using given applicationId
[**GetApplicationInstallsByApplicationId**](ApplicationAPI.md#GetApplicationInstallsByApplicationId) | **Get** /v3/application/{applicationId}/install | Get all the installs of an application
[**GetApplicationProperties**](ApplicationAPI.md#GetApplicationProperties) | **Get** /v3/application/{applicationId}/property | Get all your application properties for given applicationId
[**GetApplicationProperty**](ApplicationAPI.md#GetApplicationProperty) | **Get** /v3/application/{applicationId}/property/{propertyId} | Get the property for given applicationId and given propertyId
[**GetApplicationPropertyTypes**](ApplicationAPI.md#GetApplicationPropertyTypes) | **Get** /v3/application/property/type | Get a list of all the possible application property types
[**GetApplicationStopMethods**](ApplicationAPI.md#GetApplicationStopMethods) | **Get** /v3/application/stopMethod | Get stop methods of an application
[**GetApplicationTypes**](ApplicationAPI.md#GetApplicationTypes) | **Get** /v3/application/type | Get the list of available application types.
[**GetApplications**](ApplicationAPI.md#GetApplications) | **Get** /v3/application | Get all your applications and their main details
[**UpdateApplication**](ApplicationAPI.md#UpdateApplication) | **Put** /v3/application/{applicationId} | Update a application&#39;s mutable properties
[**UpdateApplicationProperty**](ApplicationAPI.md#UpdateApplicationProperty) | **Put** /v3/application/{applicationId}/property/{propertyId} | Update the given property of given applicationId



## CreateApplication

> []Application CreateApplication(ctx).Application(application).Execute()

Create an application

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
	application := *openapiclient.NewApplication("Id_example", int32(123), int32(123), "Name_example", int32(123), int32(123)) // Application | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.CreateApplication(context.Background()).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CreateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplication`: []Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) |  | 

### Return type

[**[]Application**](Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationProperty

> []ApplicationProperty CreateApplicationProperty(ctx, applicationId).ApplicationProperty(applicationProperty).Execute()

Create a property of given applicationId

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
	applicationId := "applicationId_example" // string | The Id of the application for which the property will be created
	applicationProperty := *openapiclient.NewApplicationProperty("Id_example", int32(123), "PropertyKey_example") // ApplicationProperty | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.CreateApplicationProperty(context.Background(), applicationId).ApplicationProperty(applicationProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CreateApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationProperty`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.CreateApplicationProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which the property will be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationProperty** | [**ApplicationProperty**](ApplicationProperty.md) |  | 

### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, applicationId).Execute()

Delete a given application

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
	applicationId := "applicationId_example" // string | The Id of the application to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.DeleteApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DeleteApplicationProperty

> DeleteApplicationProperty(ctx, applicationId, propertyId).Execute()

Delete the given property of the applicationId

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
	applicationId := "applicationId_example" // string | The Id of the application for which to delete the property
	propertyId := "propertyId_example" // string | Property Id to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.DeleteApplicationProperty(context.Background(), applicationId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to delete the property | 
**propertyId** | **string** | Property Id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationPropertyRequest struct via the builder pattern


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


## GetApplication

> []Application GetApplication(ctx, applicationId).Execute()

Get the details of the given application

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
	applicationId := "applicationId_example" // string | The Id of the application for which to fetch the information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: []Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to fetch the information | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


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


## GetApplicationApplicationBuildInUses

> []ApplicationBuild GetApplicationApplicationBuildInUses(ctx, applicationId).Execute()

Get all the application builds which have application instances created for the given application

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
	applicationId := "applicationId_example" // string | The Id of the application for which to fetch the application builds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationApplicationBuildInUses(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationApplicationBuildInUses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationApplicationBuildInUses`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationApplicationBuildInUses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to fetch the application builds | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApplicationBuildInUsesRequest struct via the builder pattern


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


## GetApplicationApplicationBuilds

> []ApplicationBuild GetApplicationApplicationBuilds(ctx, applicationId).Labels(labels).Execute()

Get all the application builds created for the given application

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
	applicationId := "applicationId_example" // string | The Id of the application for which to fetch the application builds
	labels := "labels_example" // string | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.<br /> E.g. `region_id=123` or multiple filters: `region_id=123 and fleet_id=456 or host_id=46256` The provided filter query needs to be url encoded. E.g.<br /> `region_id%3D123` or multiple filters: `region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256` The total would look like:<br /> `/application/{applicationId}/applicationBuild?labels=region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256`<br /> If you want to filter on a non-numeric label such as `region_name`, you have to wrap the value in double quotes: `region_name=\"Rotterdam\"`.<br /> For more information about labels check <a href=\"https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\">details</a> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationApplicationBuilds(context.Background(), applicationId).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationApplicationBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationApplicationBuilds`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationApplicationBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to fetch the application builds | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApplicationBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labels** | **string** | Label expressions can be used to apply more specific search parameters and can be written in standard SQL query language.&lt;br /&gt; E.g. &#x60;region_id&#x3D;123&#x60; or multiple filters: &#x60;region_id&#x3D;123 and fleet_id&#x3D;456 or host_id&#x3D;46256&#x60; The provided filter query needs to be url encoded. E.g.&lt;br /&gt; &#x60;region_id%3D123&#x60; or multiple filters: &#x60;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60; The total would look like:&lt;br /&gt; &#x60;/application/{applicationId}/applicationBuild?labels&#x3D;region_id%3D123%20and%20fleet_id%3D456%20or%20host_id%3D46256&#x60;&lt;br /&gt; If you want to filter on a non-numeric label such as &#x60;region_name&#x60;, you have to wrap the value in double quotes: &#x60;region_name&#x3D;\&quot;Rotterdam\&quot;&#x60;.&lt;br /&gt; For more information about labels check &lt;a href&#x3D;\&quot;https://www.i3d.net/docs/one/odp/Platform-Elements/Application/Label/#pre-defined-labels\&quot;&gt;details&lt;/a&gt; | 

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


## GetApplicationDeploymentEnvironments

> []DeploymentEnvironment GetApplicationDeploymentEnvironments(ctx).Execute()

Get the list of deployment environment using given applicationId

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationDeploymentEnvironments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationDeploymentEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationDeploymentEnvironments`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationDeploymentEnvironments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentEnvironmentsRequest struct via the builder pattern


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


## GetApplicationDeploymentTemplates

> []GetApplicationDeploymentTemplates200ResponseInner GetApplicationDeploymentTemplates(ctx).Execute()

Get a list of deployment templates for a given applicationId

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationDeploymentTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationDeploymentTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationDeploymentTemplates`: []GetApplicationDeploymentTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationDeploymentTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentTemplatesRequest struct via the builder pattern


### Return type

[**[]GetApplicationDeploymentTemplates200ResponseInner**](GetApplicationDeploymentTemplates200ResponseInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationFleets

> []Fleet GetApplicationFleets(ctx).Execute()

Get the list of fleets using given applicationId

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationFleets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationFleets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationFleetsRequest struct via the builder pattern


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


## GetApplicationInstallsByApplicationId

> []ApplicationInstall GetApplicationInstallsByApplicationId(ctx, applicationId).Execute()

Get all the installs of an application

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
	applicationId := "applicationId_example" // string | The Id of the application for which to fetch the installs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationInstallsByApplicationId(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationInstallsByApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstallsByApplicationId`: []ApplicationInstall
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationInstallsByApplicationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to fetch the installs | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstallsByApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationInstall**](ApplicationInstall.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationProperties

> []ApplicationProperty GetApplicationProperties(ctx, applicationId).Execute()

Get all your application properties for given applicationId

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
	applicationId := "applicationId_example" // string | The Id of the application for which to fetch the properties

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationProperties(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationProperties`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application for which to fetch the properties | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationProperty

> []ApplicationProperty GetApplicationProperty(ctx, applicationId, propertyId).Execute()

Get the property for given applicationId and given propertyId

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
	applicationId := "applicationId_example" // string | The Id of the application the property belongs to
	propertyId := "propertyId_example" // string | Id of the property to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationProperty(context.Background(), applicationId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationProperty`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application the property belongs to | 
**propertyId** | **string** | Id of the property to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPropertyTypes

> []ApplicationPropertyType GetApplicationPropertyTypes(ctx).Execute()

Get a list of all the possible application property types

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationPropertyTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationPropertyTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationPropertyTypes`: []ApplicationPropertyType
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationPropertyTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPropertyTypesRequest struct via the builder pattern


### Return type

[**[]ApplicationPropertyType**](ApplicationPropertyType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationStopMethods

> []ApplicationOSStopMethods GetApplicationStopMethods(ctx).Execute()

Get stop methods of an application

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationStopMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationStopMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationStopMethods`: []ApplicationOSStopMethods
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationStopMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationStopMethodsRequest struct via the builder pattern


### Return type

[**[]ApplicationOSStopMethods**](ApplicationOSStopMethods.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationTypes

> []ApplicationType GetApplicationTypes(ctx).Execute()

Get the list of available application types.

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
	resp, r, err := apiClient.ApplicationAPI.GetApplicationTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationTypes`: []ApplicationType
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTypesRequest struct via the builder pattern


### Return type

[**[]ApplicationType**](ApplicationType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> []Application GetApplications(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all your applications and their main details

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
	resp, r, err := apiClient.ApplicationAPI.GetApplications(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: []Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

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


## UpdateApplication

> []Application UpdateApplication(ctx, applicationId).Application(application).Execute()

Update a application's mutable properties

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
	applicationId := "applicationId_example" // string | The Id of the application to update
	application := *openapiclient.NewApplication("Id_example", int32(123), int32(123), "Name_example", int32(123), int32(123)) // Application | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.UpdateApplication(context.Background(), applicationId).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplication`: []Application
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**Application**](Application.md) |  | 

### Return type

[**[]Application**](Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationProperty

> []ApplicationProperty UpdateApplicationProperty(ctx, applicationId, propertyId).ApplicationProperty(applicationProperty).Execute()

Update the given property of given applicationId

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
	applicationId := "applicationId_example" // string | The Id of the application the property belongs to
	propertyId := "propertyId_example" // string | The Id of the property to update
	applicationProperty := *openapiclient.NewApplicationProperty("Id_example", int32(123), "PropertyKey_example") // ApplicationProperty | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.UpdateApplicationProperty(context.Background(), applicationId, propertyId).ApplicationProperty(applicationProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.UpdateApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationProperty`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.UpdateApplicationProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Id of the application the property belongs to | 
**propertyId** | **string** | The Id of the property to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationProperty** | [**ApplicationProperty**](ApplicationProperty.md) |  | 

### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

