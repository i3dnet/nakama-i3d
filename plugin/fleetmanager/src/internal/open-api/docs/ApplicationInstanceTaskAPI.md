# \ApplicationInstanceTaskAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationInstanceTask**](ApplicationInstanceTaskAPI.md#CreateApplicationInstanceTask) | **Post** /v3/applicationInstance/task | Create a new task batch for given application instance Ids, using the provided task template Id
[**DeleteApplicationInstanceTask**](ApplicationInstanceTaskAPI.md#DeleteApplicationInstanceTask) | **Delete** /v3/applicationInstance/task/{taskId} | Cancel the given task
[**GetApplicationInstanceTask**](ApplicationInstanceTaskAPI.md#GetApplicationInstanceTask) | **Get** /v3/applicationInstance/task/{taskId} | Get the status of a task
[**GetApplicationInstanceTaskBatch**](ApplicationInstanceTaskAPI.md#GetApplicationInstanceTaskBatch) | **Get** /v3/applicationInstance/taskBatch/{taskBatchId} | Get the status of a task batch
[**GetApplicationInstanceTaskBatches**](ApplicationInstanceTaskAPI.md#GetApplicationInstanceTaskBatches) | **Get** /v3/applicationInstance/taskBatch | Get the last 100 task batches you have created
[**GetApplicationInstanceTaskTemplates**](ApplicationInstanceTaskAPI.md#GetApplicationInstanceTaskTemplates) | **Get** /v3/applicationInstance/task/template | Get all available task templates for application instance
[**GetApplicationInstanceTasks**](ApplicationInstanceTaskAPI.md#GetApplicationInstanceTasks) | **Get** /v3/applicationInstance/task | Get the last 100 tasks you have created
[**UpdateApplicationInstanceTask**](ApplicationInstanceTaskAPI.md#UpdateApplicationInstanceTask) | **Put** /v3/applicationInstance/task/{taskId} | Pause the given task if it&#39;s not done yet. Or if the task is paused already, resume it.



## CreateApplicationInstanceTask

> []TaskBatchStatus CreateApplicationInstanceTask(ctx).ApplicationInstanceTaskCreate(applicationInstanceTaskCreate).Execute()

Create a new task batch for given application instance Ids, using the provided task template Id

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
	applicationInstanceTaskCreate := *openapiclient.NewApplicationInstanceTaskCreate("TaskTemplateId_example") // ApplicationInstanceTaskCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.CreateApplicationInstanceTask(context.Background()).ApplicationInstanceTaskCreate(applicationInstanceTaskCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.CreateApplicationInstanceTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInstanceTask`: []TaskBatchStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.CreateApplicationInstanceTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInstanceTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationInstanceTaskCreate** | [**ApplicationInstanceTaskCreate**](ApplicationInstanceTaskCreate.md) |  | 

### Return type

[**[]TaskBatchStatus**](TaskBatchStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationInstanceTask

> []TaskStatus DeleteApplicationInstanceTask(ctx, taskId).Execute()

Cancel the given task

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
	taskId := int32(56) // int32 | The Id of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.DeleteApplicationInstanceTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.DeleteApplicationInstanceTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationInstanceTask`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.DeleteApplicationInstanceTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** | The Id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationInstanceTaskRequest struct via the builder pattern


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


## GetApplicationInstanceTask

> []TaskStatus GetApplicationInstanceTask(ctx, taskId).Execute()

Get the status of a task

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
	taskId := int32(56) // int32 | The Id of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.GetApplicationInstanceTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.GetApplicationInstanceTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTask`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.GetApplicationInstanceTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** | The Id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTaskRequest struct via the builder pattern


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


## GetApplicationInstanceTaskBatch

> []TaskBatchStatus GetApplicationInstanceTaskBatch(ctx, taskBatchId).Execute()

Get the status of a task batch

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
	taskBatchId := int32(56) // int32 | The Id of the task batch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatch(context.Background(), taskBatchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTaskBatch`: []TaskBatchStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskBatchId** | **int32** | The Id of the task batch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTaskBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TaskBatchStatus**](TaskBatchStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceTaskBatches

> []TaskBatchStatus GetApplicationInstanceTaskBatches(ctx).Execute()

Get the last 100 task batches you have created

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
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatches(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTaskBatches`: []TaskBatchStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskBatches`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTaskBatchesRequest struct via the builder pattern


### Return type

[**[]TaskBatchStatus**](TaskBatchStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceTaskTemplates

> []TaskTemplate GetApplicationInstanceTaskTemplates(ctx).Execute()

Get all available task templates for application instance

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
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.GetApplicationInstanceTaskTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTaskTemplates`: []TaskTemplate
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.GetApplicationInstanceTaskTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTaskTemplatesRequest struct via the builder pattern


### Return type

[**[]TaskTemplate**](TaskTemplate.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstanceTasks

> []TaskStatus GetApplicationInstanceTasks(ctx).Execute()

Get the last 100 tasks you have created

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
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.GetApplicationInstanceTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.GetApplicationInstanceTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationInstanceTasks`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.GetApplicationInstanceTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceTasksRequest struct via the builder pattern


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


## UpdateApplicationInstanceTask

> []TaskStatus UpdateApplicationInstanceTask(ctx, taskId).Execute()

Pause the given task if it's not done yet. Or if the task is paused already, resume it.

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
	taskId := int32(56) // int32 | The Id of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInstanceTaskAPI.UpdateApplicationInstanceTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInstanceTaskAPI.UpdateApplicationInstanceTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationInstanceTask`: []TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInstanceTaskAPI.UpdateApplicationInstanceTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** | The Id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationInstanceTaskRequest struct via the builder pattern


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

