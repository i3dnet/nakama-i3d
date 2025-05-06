# \ApplicationBuildAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationBuild**](ApplicationBuildAPI.md#CreateApplicationBuild) | **Post** /v3/applicationBuild | Create a new application build
[**CreateApplicationBuildConfiguration**](ApplicationBuildAPI.md#CreateApplicationBuildConfiguration) | **Post** /v3/applicationBuild/{applicationBuildId}/configuration | Create a new configuration for the given applicationBuildId
[**DeleteApplicationBuild**](ApplicationBuildAPI.md#DeleteApplicationBuild) | **Delete** /v3/applicationBuild/{applicationBuildId} | Delete given application build
[**DeleteApplicationBuildConfiguration**](ApplicationBuildAPI.md#DeleteApplicationBuildConfiguration) | **Delete** /v3/applicationBuild/{applicationBuildId}/configuration/{applicationBuildConfigurationId} | Delete given configuration from given applicationBuildId
[**GetApplicationBuild**](ApplicationBuildAPI.md#GetApplicationBuild) | **Get** /v3/applicationBuild/{applicationBuildId} | Get a specific application build
[**GetApplicationBuildConfiguration**](ApplicationBuildAPI.md#GetApplicationBuildConfiguration) | **Get** /v3/applicationBuild/{applicationBuildId}/configuration/{applicationBuildConfigurationId} | Get a specific application build configuration
[**GetApplicationBuildConfigurations**](ApplicationBuildAPI.md#GetApplicationBuildConfigurations) | **Get** /v3/applicationBuild/{applicationBuildId}/configuration | Get all configurations for the given applicationBuildId
[**GetApplicationBuildDeploymentEnvironments**](ApplicationBuildAPI.md#GetApplicationBuildDeploymentEnvironments) | **Get** /v3/applicationBuild/{applicationBuildId}/deploymentEnvironment | Get a list of deployment environments for a given application build Id
[**GetApplicationBuildDeploymentTemplates**](ApplicationBuildAPI.md#GetApplicationBuildDeploymentTemplates) | **Get** /v3/applicationBuild/{applicationBuildId}/deploymentTemplate | Get a list of deployment templates for a given application build Id
[**GetApplicationBuildFleets**](ApplicationBuildAPI.md#GetApplicationBuildFleets) | **Get** /v3/applicationBuild/{applicationBuildId}/fleet | Get a list of fleets for a given application build Id
[**GetApplicationBuildProperties**](ApplicationBuildAPI.md#GetApplicationBuildProperties) | **Get** /v3/applicationBuild/{applicationBuildId}/property | Get all the properties for the given applicationBuildId
[**GetApplicationBuildProperty**](ApplicationBuildAPI.md#GetApplicationBuildProperty) | **Get** /v3/applicationBuild/{applicationBuildId}/property/{propertyId} | Get given property for the given application build Id
[**GetApplicationBuildStorageTypes**](ApplicationBuildAPI.md#GetApplicationBuildStorageTypes) | **Get** /v3/applicationBuild/storage/types | Get The build provisioning storage types for storing application build file
[**GetApplicationBuilds**](ApplicationBuildAPI.md#GetApplicationBuilds) | **Get** /v3/applicationBuild | Get all your application builds
[**UpdateApplicationBuild**](ApplicationBuildAPI.md#UpdateApplicationBuild) | **Put** /v3/applicationBuild/{applicationBuildId} | Update the given application build
[**UpdateApplicationBuildConfiguration**](ApplicationBuildAPI.md#UpdateApplicationBuildConfiguration) | **Put** /v3/applicationBuild/{applicationBuildId}/configuration/{applicationBuildConfigurationId} | Update an existing configuration for given applicationBuildId
[**UpdateApplicationBuildProperty**](ApplicationBuildAPI.md#UpdateApplicationBuildProperty) | **Put** /v3/applicationBuild/{applicationBuildId}/property/{propertyId} | Update given property for given application build



## CreateApplicationBuild

> []ApplicationBuild CreateApplicationBuild(ctx).ApplicationBuild(applicationBuild).Execute()

Create a new application build

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
	applicationBuild := *openapiclient.NewApplicationBuild("Id_example", "Name_example", "ApplicationId_example", int32(123), "Executable_example", int32(123), int32(123), int32(123)) // ApplicationBuild | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.CreateApplicationBuild(context.Background()).ApplicationBuild(applicationBuild).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.CreateApplicationBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationBuild`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.CreateApplicationBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationBuild** | [**ApplicationBuild**](ApplicationBuild.md) |  | 

### Return type

[**[]ApplicationBuild**](ApplicationBuild.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationBuildConfiguration

> []ApplicationBuildConfiguration CreateApplicationBuildConfiguration(ctx, applicationBuildId).ApplicationBuildConfiguration(applicationBuildConfiguration).Execute()

Create a new configuration for the given applicationBuildId

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	applicationBuildConfiguration := *openapiclient.NewApplicationBuildConfiguration("Id_example", "ConfigPath_example", "ConfigName_example", "ConfigContent_example", int32(123)) // ApplicationBuildConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.CreateApplicationBuildConfiguration(context.Background(), applicationBuildId).ApplicationBuildConfiguration(applicationBuildConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.CreateApplicationBuildConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationBuildConfiguration`: []ApplicationBuildConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.CreateApplicationBuildConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationBuildConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationBuildConfiguration** | [**ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md) |  | 

### Return type

[**[]ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationBuild

> DeleteApplicationBuild(ctx, applicationBuildId).Execute()

Delete given application build

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationBuildAPI.DeleteApplicationBuild(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.DeleteApplicationBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationBuildRequest struct via the builder pattern


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


## DeleteApplicationBuildConfiguration

> DeleteApplicationBuildConfiguration(ctx, applicationBuildId, applicationBuildConfigurationId).Execute()

Delete given configuration from given applicationBuildId

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	applicationBuildConfigurationId := "applicationBuildConfigurationId_example" // string | The Id of the specific configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationBuildAPI.DeleteApplicationBuildConfiguration(context.Background(), applicationBuildId, applicationBuildConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.DeleteApplicationBuildConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 
**applicationBuildConfigurationId** | **string** | The Id of the specific configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationBuildConfigurationRequest struct via the builder pattern


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


## GetApplicationBuild

> []ApplicationBuild GetApplicationBuild(ctx, applicationBuildId).Execute()

Get a specific application build

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuild(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuild`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildRequest struct via the builder pattern


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


## GetApplicationBuildConfiguration

> []ApplicationBuildConfiguration GetApplicationBuildConfiguration(ctx, applicationBuildId, applicationBuildConfigurationId).Execute()

Get a specific application build configuration

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	applicationBuildConfigurationId := "applicationBuildConfigurationId_example" // string | The Id of the specific configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildConfiguration(context.Background(), applicationBuildId, applicationBuildConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildConfiguration`: []ApplicationBuildConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 
**applicationBuildConfigurationId** | **string** | The Id of the specific configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationBuildConfigurations

> []ApplicationBuildConfiguration GetApplicationBuildConfigurations(ctx, applicationBuildId).Execute()

Get all configurations for the given applicationBuildId

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildConfigurations(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildConfigurations`: []ApplicationBuildConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationBuildDeploymentEnvironments

> []DeploymentEnvironment GetApplicationBuildDeploymentEnvironments(ctx, applicationBuildId).Execute()

Get a list of deployment environments for a given application build Id

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
	applicationBuildId := "applicationBuildId_example" // string | The ID of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildDeploymentEnvironments(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildDeploymentEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildDeploymentEnvironments`: []DeploymentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildDeploymentEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The ID of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildDeploymentEnvironmentsRequest struct via the builder pattern


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


## GetApplicationBuildDeploymentTemplates

> []GetApplicationDeploymentTemplates200ResponseInner GetApplicationBuildDeploymentTemplates(ctx, applicationBuildId).Execute()

Get a list of deployment templates for a given application build Id

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
	applicationBuildId := "applicationBuildId_example" // string | The ID of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildDeploymentTemplates(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildDeploymentTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildDeploymentTemplates`: []GetApplicationDeploymentTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildDeploymentTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The ID of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildDeploymentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetApplicationBuildFleets

> []Fleet GetApplicationBuildFleets(ctx, applicationBuildId).Execute()

Get a list of fleets for a given application build Id

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
	applicationBuildId := "applicationBuildId_example" // string | The ID of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildFleets(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The ID of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildFleetsRequest struct via the builder pattern


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


## GetApplicationBuildProperties

> []ApplicationBuildProperty GetApplicationBuildProperties(ctx, applicationBuildId).Execute()

Get all the properties for the given applicationBuildId

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildProperties(context.Background(), applicationBuildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildProperties`: []ApplicationBuildProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationBuildProperty**](ApplicationBuildProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationBuildProperty

> []ApplicationBuildProperty GetApplicationBuildProperty(ctx, applicationBuildId, propertyId).Execute()

Get given property for the given application build Id

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	propertyId := "propertyId_example" // string | The Id of the application build property

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildProperty(context.Background(), applicationBuildId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildProperty`: []ApplicationBuildProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 
**propertyId** | **string** | The Id of the application build property | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApplicationBuildProperty**](ApplicationBuildProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationBuildStorageTypes

> []BuildProvisioningStorageTypes GetApplicationBuildStorageTypes(ctx).Execute()

Get The build provisioning storage types for storing application build file

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
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuildStorageTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuildStorageTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildStorageTypes`: []BuildProvisioningStorageTypes
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuildStorageTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildStorageTypesRequest struct via the builder pattern


### Return type

[**[]BuildProvisioningStorageTypes**](BuildProvisioningStorageTypes.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationBuilds

> []ApplicationBuild GetApplicationBuilds(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all your application builds

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
	resp, r, err := apiClient.ApplicationBuildAPI.GetApplicationBuilds(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.GetApplicationBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuilds`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.GetApplicationBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildsRequest struct via the builder pattern


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


## UpdateApplicationBuild

> []ApplicationBuild UpdateApplicationBuild(ctx, applicationBuildId).ApplicationBuild(applicationBuild).Execute()

Update the given application build

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	applicationBuild := *openapiclient.NewApplicationBuild("Id_example", "Name_example", "ApplicationId_example", int32(123), "Executable_example", int32(123), int32(123), int32(123)) // ApplicationBuild | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.UpdateApplicationBuild(context.Background(), applicationBuildId).ApplicationBuild(applicationBuild).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.UpdateApplicationBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationBuild`: []ApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.UpdateApplicationBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationBuild** | [**ApplicationBuild**](ApplicationBuild.md) |  | 

### Return type

[**[]ApplicationBuild**](ApplicationBuild.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationBuildConfiguration

> []ApplicationBuildConfiguration UpdateApplicationBuildConfiguration(ctx, applicationBuildId, applicationBuildConfigurationId).ApplicationBuildConfiguration(applicationBuildConfiguration).Execute()

Update an existing configuration for given applicationBuildId

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	applicationBuildConfigurationId := "applicationBuildConfigurationId_example" // string | The Id of the specific configuration
	applicationBuildConfiguration := *openapiclient.NewApplicationBuildConfiguration("Id_example", "ConfigPath_example", "ConfigName_example", "ConfigContent_example", int32(123)) // ApplicationBuildConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.UpdateApplicationBuildConfiguration(context.Background(), applicationBuildId, applicationBuildConfigurationId).ApplicationBuildConfiguration(applicationBuildConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.UpdateApplicationBuildConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationBuildConfiguration`: []ApplicationBuildConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.UpdateApplicationBuildConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 
**applicationBuildConfigurationId** | **string** | The Id of the specific configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationBuildConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationBuildConfiguration** | [**ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md) |  | 

### Return type

[**[]ApplicationBuildConfiguration**](ApplicationBuildConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationBuildProperty

> []ApplicationBuildProperty UpdateApplicationBuildProperty(ctx, applicationBuildId, propertyId).ApplicationBuildProperty(applicationBuildProperty).Execute()

Update given property for given application build

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
	applicationBuildId := "applicationBuildId_example" // string | The Id of the application build
	propertyId := "propertyId_example" // string | The Id of the application build property to update
	applicationBuildProperty := *openapiclient.NewApplicationBuildProperty("Id_example", int32(123), "PropertyKey_example", "PropertyValue_example") // ApplicationBuildProperty | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAPI.UpdateApplicationBuildProperty(context.Background(), applicationBuildId, propertyId).ApplicationBuildProperty(applicationBuildProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAPI.UpdateApplicationBuildProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationBuildProperty`: []ApplicationBuildProperty
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAPI.UpdateApplicationBuildProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationBuildId** | **string** | The Id of the application build | 
**propertyId** | **string** | The Id of the application build property to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationBuildPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationBuildProperty** | [**ApplicationBuildProperty**](ApplicationBuildProperty.md) |  | 

### Return type

[**[]ApplicationBuildProperty**](ApplicationBuildProperty.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

