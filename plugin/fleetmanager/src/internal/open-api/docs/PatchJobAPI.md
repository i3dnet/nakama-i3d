# \PatchJobAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePatchJob**](PatchJobAPI.md#CreatePatchJob) | **Post** /v3/patchJob | Create a patch job
[**CreatePatchJobApplicationInstancePreview**](PatchJobAPI.md#CreatePatchJobApplicationInstancePreview) | **Post** /v3/patchJob/applicationInstancePreview | Show a preview of the number of game application instances that are affected by the selection
[**CreatePatchJobEmail**](PatchJobAPI.md#CreatePatchJobEmail) | **Post** /v3/patchJob/{patchJobId}/email | Create a patch job report email
[**DeletePatchJobEmail**](PatchJobAPI.md#DeletePatchJobEmail) | **Delete** /v3/patchJob/{patchJobId}/email/{patchJobReportEmailId} | Delete a patch job report email
[**GetPatchJob**](PatchJobAPI.md#GetPatchJob) | **Get** /v3/patchJob/{patchJobId} | Get the detail of a patch job
[**GetPatchJobApplicationInstanceFailedApplicationInstances**](PatchJobAPI.md#GetPatchJobApplicationInstanceFailedApplicationInstances) | **Get** /v3/patchJob/{patchJobId}/applicationInstance/{applicationInstanceId}/failedApplicationInstance | Get a failed application instance
[**GetPatchJobDeploymentEnvironmentApplicationFleets**](PatchJobAPI.md#GetPatchJobDeploymentEnvironmentApplicationFleets) | **Get** /v3/patchJob/deploymentEnvironment/{deploymentEnvironmentId}/application/{applicationId}/fleet | Gets the list of patch job fleet based on deploymentEnvironmentId and applicationId
[**GetPatchJobEmails**](PatchJobAPI.md#GetPatchJobEmails) | **Get** /v3/patchJob/{patchJobId}/email | Gets the list of reporting emails for a specific patch job
[**GetPatchJobFailedApplicationInstance**](PatchJobAPI.md#GetPatchJobFailedApplicationInstance) | **Get** /v3/patchJob/{patchJobId}/failedApplicationInstances | Get a list of failed application instances
[**GetPatchJobReportProgress**](PatchJobAPI.md#GetPatchJobReportProgress) | **Get** /v3/patchJob/{patchJobId}/report/progress | Get gives back the progress of the patch job
[**GetPatchJobRuntimeTypes**](PatchJobAPI.md#GetPatchJobRuntimeTypes) | **Get** /v3/patchJob/runtimeType | Get list of patch job runtime types
[**GetPatchJobStatuses**](PatchJobAPI.md#GetPatchJobStatuses) | **Get** /v3/patchJob/status | Get list of patch job status
[**GetPatchJobSummary**](PatchJobAPI.md#GetPatchJobSummary) | **Get** /v3/patchJob/summary | Get summary of all patch jobs
[**GetPatchJobTypes**](PatchJobAPI.md#GetPatchJobTypes) | **Get** /v3/patchJob/type | Get list of patch job types
[**GetPatchJobs**](PatchJobAPI.md#GetPatchJobs) | **Get** /v3/patchJob | Get the list of all patch jobs
[**UpdatePatchJob**](PatchJobAPI.md#UpdatePatchJob) | **Put** /v3/patchJob/{patchJobId} | Update the detail of a patch job
[**UpdatePatchJobCancel**](PatchJobAPI.md#UpdatePatchJobCancel) | **Put** /v3/patchJob/{patchJobId}/cancel | Cancel the given patch job
[**UpdatePatchJobEmailBulk**](PatchJobAPI.md#UpdatePatchJobEmailBulk) | **Put** /v3/patchJob/{patchJobId}/email/bulk | Create, update and remove patch job report email
[**UpdatePatchJobHold**](PatchJobAPI.md#UpdatePatchJobHold) | **Put** /v3/patchJob/{patchJobId}/hold | Put a given patch job on hold
[**UpdatePatchJobStop**](PatchJobAPI.md#UpdatePatchJobStop) | **Put** /v3/patchJob/{patchJobId}/stop | Stops the given patch job and set the patch job to success
[**UpdatePatchJobUnhold**](PatchJobAPI.md#UpdatePatchJobUnhold) | **Put** /v3/patchJob/{patchJobId}/unhold | Unhold the given patch job



## CreatePatchJob

> []PatchJob CreatePatchJob(ctx).PatchJob(patchJob).Execute()

Create a patch job

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
	patchJob := *openapiclient.NewPatchJob("Id_example", "DeploymentEnvironmentId_example", "ApplicationId_example", int32(123), "PatchJobMethod_example", "PatchJobName_example", int32(123), int32(123), "StopMethodName_example", int32(123), "TODO", []openapiclient.PatchJobApplicationBuild{*openapiclient.NewPatchJobApplicationBuild("Id_example", "OldApplicationBuildId_example", "NewApplicationBuildId_example")}, []openapiclient.PatchJobFleet{*openapiclient.NewPatchJobFleet("Id_example", "FleetId_example")}, int32(123), int32(123), int32(123), int32(123), NullableInt32(123), NullableInt32(123)) // PatchJob | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.CreatePatchJob(context.Background()).PatchJob(patchJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.CreatePatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePatchJob`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.CreatePatchJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchJob** | [**PatchJob**](PatchJob.md) |  | 

### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePatchJobApplicationInstancePreview

> []PatchJobApplicationInstanceCount CreatePatchJobApplicationInstancePreview(ctx).PatchJobApplicationInstancePreview(patchJobApplicationInstancePreview).Execute()

Show a preview of the number of game application instances that are affected by the selection

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
	patchJobApplicationInstancePreview := *openapiclient.NewPatchJobApplicationInstancePreview("DeploymentEnvironmentId_example", "ApplicationId_example", []string{"Fleet_example"}, []string{"ApplicationBuild_example"}) // PatchJobApplicationInstancePreview | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.CreatePatchJobApplicationInstancePreview(context.Background()).PatchJobApplicationInstancePreview(patchJobApplicationInstancePreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.CreatePatchJobApplicationInstancePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePatchJobApplicationInstancePreview`: []PatchJobApplicationInstanceCount
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.CreatePatchJobApplicationInstancePreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePatchJobApplicationInstancePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchJobApplicationInstancePreview** | [**PatchJobApplicationInstancePreview**](PatchJobApplicationInstancePreview.md) |  | 

### Return type

[**[]PatchJobApplicationInstanceCount**](PatchJobApplicationInstanceCount.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePatchJobEmail

> []PatchJobEmail CreatePatchJobEmail(ctx, patchJobId).PatchJobEmail(patchJobEmail).Execute()

Create a patch job report email

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job
	patchJobEmail := *openapiclient.NewPatchJobEmail("Id_example", "Email_example", int32(123), int32(123), int32(123), int32(123)) // PatchJobEmail | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.CreatePatchJobEmail(context.Background(), patchJobId).PatchJobEmail(patchJobEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.CreatePatchJobEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePatchJobEmail`: []PatchJobEmail
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.CreatePatchJobEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePatchJobEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchJobEmail** | [**PatchJobEmail**](PatchJobEmail.md) |  | 

### Return type

[**[]PatchJobEmail**](PatchJobEmail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePatchJobEmail

> []PatchJob DeletePatchJobEmail(ctx, patchJobId, patchJobReportEmailId).Execute()

Delete a patch job report email

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job
	patchJobReportEmailId := "patchJobReportEmailId_example" // string | The Id of the patch job report email

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.DeletePatchJobEmail(context.Background(), patchJobId, patchJobReportEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.DeletePatchJobEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePatchJobEmail`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.DeletePatchJobEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 
**patchJobReportEmailId** | **string** | The Id of the patch job report email | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePatchJobEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJob

> []PatchJob GetPatchJob(ctx, patchJobId).Execute()

Get the detail of a patch job

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job for which to fetch the information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.GetPatchJob(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJob`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job for which to fetch the information | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobApplicationInstanceFailedApplicationInstances

> []PatchJobFailedApplicationInstance GetPatchJobApplicationInstanceFailedApplicationInstances(ctx).Execute()

Get a failed application instance

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobApplicationInstanceFailedApplicationInstances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobApplicationInstanceFailedApplicationInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobApplicationInstanceFailedApplicationInstances`: []PatchJobFailedApplicationInstance
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobApplicationInstanceFailedApplicationInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobApplicationInstanceFailedApplicationInstancesRequest struct via the builder pattern


### Return type

[**[]PatchJobFailedApplicationInstance**](PatchJobFailedApplicationInstance.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobDeploymentEnvironmentApplicationFleets

> []PatchJobFleetApplicationBuild GetPatchJobDeploymentEnvironmentApplicationFleets(ctx, deploymentEnvironmentId, applicationId).Execute()

Gets the list of patch job fleet based on deploymentEnvironmentId and applicationId

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
	applicationId := "applicationId_example" // string | The Id of the application (game or utility)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobDeploymentEnvironmentApplicationFleets(context.Background(), deploymentEnvironmentId, applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobDeploymentEnvironmentApplicationFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobDeploymentEnvironmentApplicationFleets`: []PatchJobFleetApplicationBuild
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobDeploymentEnvironmentApplicationFleets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentEnvironmentId** | **string** | The Id of the deployment environment | 
**applicationId** | **string** | The Id of the application (game or utility) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobDeploymentEnvironmentApplicationFleetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PatchJobFleetApplicationBuild**](PatchJobFleetApplicationBuild.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobEmails

> []PatchJobEmail GetPatchJobEmails(ctx, patchJobId).Execute()

Gets the list of reporting emails for a specific patch job

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobEmails(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobEmails`: []PatchJobEmail
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchJobEmail**](PatchJobEmail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobFailedApplicationInstance

> []ApplicationInstanceSummary GetPatchJobFailedApplicationInstance(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get a list of failed application instances

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
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobFailedApplicationInstance(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobFailedApplicationInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobFailedApplicationInstance`: []ApplicationInstanceSummary
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobFailedApplicationInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobFailedApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

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


## GetPatchJobReportProgress

> []PatchJobReport GetPatchJobReportProgress(ctx).Execute()

Get gives back the progress of the patch job

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobReportProgress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobReportProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobReportProgress`: []PatchJobReport
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobReportProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobReportProgressRequest struct via the builder pattern


### Return type

[**[]PatchJobReport**](PatchJobReport.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobRuntimeTypes

> []PatchJobRuntimeType GetPatchJobRuntimeTypes(ctx).Execute()

Get list of patch job runtime types

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobRuntimeTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobRuntimeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobRuntimeTypes`: []PatchJobRuntimeType
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobRuntimeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobRuntimeTypesRequest struct via the builder pattern


### Return type

[**[]PatchJobRuntimeType**](PatchJobRuntimeType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobStatuses

> []PatchJobStatus GetPatchJobStatuses(ctx).Execute()

Get list of patch job status

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobStatuses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobStatuses`: []PatchJobStatus
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobStatusesRequest struct via the builder pattern


### Return type

[**[]PatchJobStatus**](PatchJobStatus.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobSummary

> []PatchJobSummary GetPatchJobSummary(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get summary of all patch jobs

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobSummary(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobSummary`: []PatchJobSummary
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]PatchJobSummary**](PatchJobSummary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobTypes

> []PatchJobType GetPatchJobTypes(ctx).Execute()

Get list of patch job types

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobTypes`: []PatchJobType
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobTypesRequest struct via the builder pattern


### Return type

[**[]PatchJobType**](PatchJobType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatchJobs

> []PatchJob GetPatchJobs(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get the list of all patch jobs

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
	resp, r, err := apiClient.PatchJobAPI.GetPatchJobs(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.GetPatchJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPatchJobs`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.GetPatchJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPatchJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatchJob

> []PatchJob UpdatePatchJob(ctx, patchJobId).PatchJob(patchJob).Execute()

Update the detail of a patch job

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job
	patchJob := *openapiclient.NewPatchJob("Id_example", "DeploymentEnvironmentId_example", "ApplicationId_example", int32(123), "PatchJobMethod_example", "PatchJobName_example", int32(123), int32(123), "StopMethodName_example", int32(123), "TODO", []openapiclient.PatchJobApplicationBuild{*openapiclient.NewPatchJobApplicationBuild("Id_example", "OldApplicationBuildId_example", "NewApplicationBuildId_example")}, []openapiclient.PatchJobFleet{*openapiclient.NewPatchJobFleet("Id_example", "FleetId_example")}, int32(123), int32(123), int32(123), int32(123), NullableInt32(123), NullableInt32(123)) // PatchJob | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.UpdatePatchJob(context.Background(), patchJobId).PatchJob(patchJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePatchJob`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.UpdatePatchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchJob** | [**PatchJob**](PatchJob.md) |  | 

### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatchJobCancel

> []PatchJob UpdatePatchJobCancel(ctx, patchJobId).Execute()

Cancel the given patch job

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.UpdatePatchJobCancel(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJobCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePatchJobCancel`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.UpdatePatchJobCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatchJobEmailBulk

> []PatchJobEmail UpdatePatchJobEmailBulk(ctx, patchJobId).PatchJobEmailBulk(patchJobEmailBulk).Execute()

Create, update and remove patch job report email

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job
	patchJobEmailBulk := *openapiclient.NewPatchJobEmailBulk([]openapiclient.PatchJobEmail{*openapiclient.NewPatchJobEmail("Id_example", "Email_example", int32(123), int32(123), int32(123), int32(123))}) // PatchJobEmailBulk | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.UpdatePatchJobEmailBulk(context.Background(), patchJobId).PatchJobEmailBulk(patchJobEmailBulk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJobEmailBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePatchJobEmailBulk`: []PatchJobEmail
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.UpdatePatchJobEmailBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobEmailBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchJobEmailBulk** | [**PatchJobEmailBulk**](PatchJobEmailBulk.md) |  | 

### Return type

[**[]PatchJobEmail**](PatchJobEmail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatchJobHold

> []PatchJob UpdatePatchJobHold(ctx, patchJobId).Execute()

Put a given patch job on hold

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job to put on hold

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.UpdatePatchJobHold(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJobHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePatchJobHold`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.UpdatePatchJobHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job to put on hold | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatchJobStop

> UpdatePatchJobStop(ctx, patchJobId).Execute()

Stops the given patch job and set the patch job to success

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatchJobAPI.UpdatePatchJobStop(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJobStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobStopRequest struct via the builder pattern


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


## UpdatePatchJobUnhold

> []PatchJob UpdatePatchJobUnhold(ctx, patchJobId).Execute()

Unhold the given patch job

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
	patchJobId := "patchJobId_example" // string | The Id of the patch job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchJobAPI.UpdatePatchJobUnhold(context.Background(), patchJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchJobAPI.UpdatePatchJobUnhold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePatchJobUnhold`: []PatchJob
	fmt.Fprintf(os.Stdout, "Response from `PatchJobAPI.UpdatePatchJobUnhold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchJobId** | **string** | The Id of the patch job | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePatchJobUnholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchJob**](PatchJob.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

