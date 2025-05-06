# \DeploymentTemplateAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentTemplateDependency**](DeploymentTemplateAPI.md#CreateDeploymentTemplateDependency) | **Post** /v3/deploymentTemplate/dependency | Create a dependency deployment template
[**CreateDeploymentTemplateGame**](DeploymentTemplateAPI.md#CreateDeploymentTemplateGame) | **Post** /v3/deploymentTemplate/game | Create a game deployment template
[**CreateDeploymentTemplateUtility**](DeploymentTemplateAPI.md#CreateDeploymentTemplateUtility) | **Post** /v3/deploymentTemplate/utility | Create an instance of utility deployment template.
[**DeleteDeploymentTemplateDependency**](DeploymentTemplateAPI.md#DeleteDeploymentTemplateDependency) | **Delete** /v3/deploymentTemplate/dependency/{deploymentTemplateId} | Delete given dependency deployment template.
[**DeleteDeploymentTemplateGame**](DeploymentTemplateAPI.md#DeleteDeploymentTemplateGame) | **Delete** /v3/deploymentTemplate/game/{deploymentTemplateId} | Delete given game deployment template.
[**DeleteDeploymentTemplateUtility**](DeploymentTemplateAPI.md#DeleteDeploymentTemplateUtility) | **Delete** /v3/deploymentTemplate/utility/{deploymentTemplateId} | Delete given utility deployment template
[**GetDeploymentTemplateDependencies**](DeploymentTemplateAPI.md#GetDeploymentTemplateDependencies) | **Get** /v3/deploymentTemplate/dependency | Get all dependency deployment templates
[**GetDeploymentTemplateDependency**](DeploymentTemplateAPI.md#GetDeploymentTemplateDependency) | **Get** /v3/deploymentTemplate/dependency/{deploymentTemplateId} | Get the details of dependency deployment template
[**GetDeploymentTemplateDependencyFleets**](DeploymentTemplateAPI.md#GetDeploymentTemplateDependencyFleets) | **Get** /v3/deploymentTemplate/dependency/{deploymentTemplateId}/fleet | Get all fleets using the given deploymentTemplateId
[**GetDeploymentTemplateGame**](DeploymentTemplateAPI.md#GetDeploymentTemplateGame) | **Get** /v3/deploymentTemplate/game/{deploymentTemplateId} | Get the details of game deployment template
[**GetDeploymentTemplateGameFleets**](DeploymentTemplateAPI.md#GetDeploymentTemplateGameFleets) | **Get** /v3/deploymentTemplate/game/{deploymentTemplateId}/fleet | Get all fleets using the given game deploymentTemplateId
[**GetDeploymentTemplateGames**](DeploymentTemplateAPI.md#GetDeploymentTemplateGames) | **Get** /v3/deploymentTemplate/game | Get all game deployment templates
[**GetDeploymentTemplateUtilities**](DeploymentTemplateAPI.md#GetDeploymentTemplateUtilities) | **Get** /v3/deploymentTemplate/utility | All your utility deployment templates
[**GetDeploymentTemplateUtility**](DeploymentTemplateAPI.md#GetDeploymentTemplateUtility) | **Get** /v3/deploymentTemplate/utility/{deploymentTemplateId} | Get the details of a utility deployment template
[**GetDeploymentTemplateUtilityFleets**](DeploymentTemplateAPI.md#GetDeploymentTemplateUtilityFleets) | **Get** /v3/deploymentTemplate/utility/{deploymentTemplateId}/fleet | Get list of fleet using utility deployment template Id
[**UpdateDeploymentTemplateDependency**](DeploymentTemplateAPI.md#UpdateDeploymentTemplateDependency) | **Put** /v3/deploymentTemplate/dependency/{deploymentTemplateId} | Update given dependency deployment template.
[**UpdateDeploymentTemplateGame**](DeploymentTemplateAPI.md#UpdateDeploymentTemplateGame) | **Put** /v3/deploymentTemplate/game/{deploymentTemplateId} | Update given game deployment template.
[**UpdateDeploymentTemplateUtility**](DeploymentTemplateAPI.md#UpdateDeploymentTemplateUtility) | **Put** /v3/deploymentTemplate/utility/{deploymentTemplateId} | Update an instance of utility deployment template



## CreateDeploymentTemplateDependency

> []DependencyDeploymentTemplate CreateDeploymentTemplateDependency(ctx).DependencyDeploymentTemplate(dependencyDeploymentTemplate).Execute()

Create a dependency deployment template

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
	dependencyDeploymentTemplate := *openapiclient.NewDependencyDeploymentTemplate("Id_example", "Name_example", "DependencyInstallerBuildId_example", "DependencyUninstallerBuildId_example", []string{"FleetIds_example"}, int32(123), int32(123)) // DependencyDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.CreateDeploymentTemplateDependency(context.Background()).DependencyDeploymentTemplate(dependencyDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.CreateDeploymentTemplateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentTemplateDependency`: []DependencyDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.CreateDeploymentTemplateDependency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentTemplateDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dependencyDeploymentTemplate** | [**DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md) |  | 

### Return type

[**[]DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeploymentTemplateGame

> []GameDeploymentTemplate CreateDeploymentTemplateGame(ctx).GameDeploymentTemplate(gameDeploymentTemplate).Execute()

Create a game deployment template

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
	gameDeploymentTemplate := *openapiclient.NewGameDeploymentTemplate("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), []openapiclient.GameDeploymentTemplateBuild{*openapiclient.NewGameDeploymentTemplateBuild("ApplicationBuildId_example")}) // GameDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.CreateDeploymentTemplateGame(context.Background()).GameDeploymentTemplate(gameDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.CreateDeploymentTemplateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentTemplateGame`: []GameDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.CreateDeploymentTemplateGame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentTemplateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameDeploymentTemplate** | [**GameDeploymentTemplate**](GameDeploymentTemplate.md) |  | 

### Return type

[**[]GameDeploymentTemplate**](GameDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeploymentTemplateUtility

> []UtilityDeploymentTemplate CreateDeploymentTemplateUtility(ctx).UtilityDeploymentTemplate(utilityDeploymentTemplate).Execute()

Create an instance of utility deployment template.

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
	utilityDeploymentTemplate := *openapiclient.NewUtilityDeploymentTemplate("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), []openapiclient.UtilityDeploymentTemplateBuild{*openapiclient.NewUtilityDeploymentTemplateBuild("ApplicationBuildId_example")}) // UtilityDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.CreateDeploymentTemplateUtility(context.Background()).UtilityDeploymentTemplate(utilityDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.CreateDeploymentTemplateUtility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeploymentTemplateUtility`: []UtilityDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.CreateDeploymentTemplateUtility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentTemplateUtilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **utilityDeploymentTemplate** | [**UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md) |  | 

### Return type

[**[]UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentTemplateDependency

> DeleteDeploymentTemplateDependency(ctx, deploymentTemplateId).Execute()

Delete given dependency deployment template.

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the dependency deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentTemplateAPI.DeleteDeploymentTemplateDependency(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.DeleteDeploymentTemplateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the dependency deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentTemplateDependencyRequest struct via the builder pattern


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


## DeleteDeploymentTemplateGame

> DeleteDeploymentTemplateGame(ctx, deploymentTemplateId).Execute()

Delete given game deployment template.

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentTemplateAPI.DeleteDeploymentTemplateGame(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.DeleteDeploymentTemplateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentTemplateGameRequest struct via the builder pattern


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


## DeleteDeploymentTemplateUtility

> DeleteDeploymentTemplateUtility(ctx, deploymentTemplateId).Execute()

Delete given utility deployment template

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentTemplateAPI.DeleteDeploymentTemplateUtility(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.DeleteDeploymentTemplateUtility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentTemplateUtilityRequest struct via the builder pattern


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


## GetDeploymentTemplateDependencies

> []DependencyDeploymentTemplate GetDeploymentTemplateDependencies(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all dependency deployment templates

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
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateDependencies(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateDependencies`: []DependencyDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateDependencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateDependency

> []DependencyDeploymentTemplate GetDeploymentTemplateDependency(ctx, deploymentTemplateId).Execute()

Get the details of dependency deployment template

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the dependency deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateDependency(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateDependency`: []DependencyDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the dependency deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateDependencyFleets

> []Fleet GetDeploymentTemplateDependencyFleets(ctx, deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get all fleets using the given deploymentTemplateId

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the dependency deployment template
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateDependencyFleets(context.Background(), deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateDependencyFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateDependencyFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateDependencyFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the dependency deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateDependencyFleetsRequest struct via the builder pattern


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


## GetDeploymentTemplateGame

> []GameDeploymentTemplate GetDeploymentTemplateGame(ctx, deploymentTemplateId).Execute()

Get the details of game deployment template

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateGame(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateGame`: []GameDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GameDeploymentTemplate**](GameDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateGameFleets

> []Fleet GetDeploymentTemplateGameFleets(ctx, deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get all fleets using the given game deploymentTemplateId

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateGameFleets(context.Background(), deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateGameFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateGameFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateGameFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateGameFleetsRequest struct via the builder pattern


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


## GetDeploymentTemplateGames

> []GameDeploymentTemplate GetDeploymentTemplateGames(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all game deployment templates

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
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateGames(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateGames`: []GameDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateGames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]GameDeploymentTemplate**](GameDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateUtilities

> []UtilityDeploymentTemplate GetDeploymentTemplateUtilities(ctx).RANGEDDATA(rANGEDDATA).Execute()

All your utility deployment templates

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
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateUtilities(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateUtilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateUtilities`: []UtilityDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateUtilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateUtilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateUtility

> []UtilityDeploymentTemplate GetDeploymentTemplateUtility(ctx, deploymentTemplateId).Execute()

Get the details of a utility deployment template

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateUtility(context.Background(), deploymentTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateUtility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateUtility`: []UtilityDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateUtility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateUtilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentTemplateUtilityFleets

> []Fleet GetDeploymentTemplateUtilityFleets(ctx, deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()

Get list of fleet using utility deployment template Id

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.GetDeploymentTemplateUtilityFleets(context.Background(), deploymentTemplateId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.GetDeploymentTemplateUtilityFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentTemplateUtilityFleets`: []Fleet
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.GetDeploymentTemplateUtilityFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentTemplateUtilityFleetsRequest struct via the builder pattern


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


## UpdateDeploymentTemplateDependency

> []DependencyDeploymentTemplate UpdateDeploymentTemplateDependency(ctx, deploymentTemplateId).DependencyDeploymentTemplate(dependencyDeploymentTemplate).Execute()

Update given dependency deployment template.

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the dependency deployment template
	dependencyDeploymentTemplate := *openapiclient.NewDependencyDeploymentTemplate("Id_example", "Name_example", "DependencyInstallerBuildId_example", "DependencyUninstallerBuildId_example", []string{"FleetIds_example"}, int32(123), int32(123)) // DependencyDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.UpdateDeploymentTemplateDependency(context.Background(), deploymentTemplateId).DependencyDeploymentTemplate(dependencyDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.UpdateDeploymentTemplateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentTemplateDependency`: []DependencyDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.UpdateDeploymentTemplateDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the dependency deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentTemplateDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dependencyDeploymentTemplate** | [**DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md) |  | 

### Return type

[**[]DependencyDeploymentTemplate**](DependencyDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentTemplateGame

> GameDeploymentTemplate UpdateDeploymentTemplateGame(ctx, deploymentTemplateId).GameDeploymentTemplate(gameDeploymentTemplate).Execute()

Update given game deployment template.

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template
	gameDeploymentTemplate := *openapiclient.NewGameDeploymentTemplate("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), []openapiclient.GameDeploymentTemplateBuild{*openapiclient.NewGameDeploymentTemplateBuild("ApplicationBuildId_example")}) // GameDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.UpdateDeploymentTemplateGame(context.Background(), deploymentTemplateId).GameDeploymentTemplate(gameDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.UpdateDeploymentTemplateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentTemplateGame`: GameDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.UpdateDeploymentTemplateGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentTemplateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameDeploymentTemplate** | [**GameDeploymentTemplate**](GameDeploymentTemplate.md) |  | 

### Return type

[**GameDeploymentTemplate**](GameDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentTemplateUtility

> []UtilityDeploymentTemplate UpdateDeploymentTemplateUtility(ctx, deploymentTemplateId).UtilityDeploymentTemplate(utilityDeploymentTemplate).Execute()

Update an instance of utility deployment template

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
	deploymentTemplateId := "deploymentTemplateId_example" // string | The Id of the deployment template
	utilityDeploymentTemplate := *openapiclient.NewUtilityDeploymentTemplate("Id_example", []string{"FleetIds_example"}, "Name_example", int32(123), int32(123), []openapiclient.UtilityDeploymentTemplateBuild{*openapiclient.NewUtilityDeploymentTemplateBuild("ApplicationBuildId_example")}) // UtilityDeploymentTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentTemplateAPI.UpdateDeploymentTemplateUtility(context.Background(), deploymentTemplateId).UtilityDeploymentTemplate(utilityDeploymentTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentTemplateAPI.UpdateDeploymentTemplateUtility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentTemplateUtility`: []UtilityDeploymentTemplate
	fmt.Fprintf(os.Stdout, "Response from `DeploymentTemplateAPI.UpdateDeploymentTemplateUtility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentTemplateId** | **string** | The Id of the deployment template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentTemplateUtilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilityDeploymentTemplate** | [**UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md) |  | 

### Return type

[**[]UtilityDeploymentTemplate**](UtilityDeploymentTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

