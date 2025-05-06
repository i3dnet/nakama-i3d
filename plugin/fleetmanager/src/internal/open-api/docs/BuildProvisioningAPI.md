# \BuildProvisioningAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuildProvisioningStorageFile**](BuildProvisioningAPI.md#GetBuildProvisioningStorageFile) | **Get** /v3/buildProvisioning/storage/file/{id} | Get the details of the file in the storage
[**GetBuildProvisioningStorageFiles**](BuildProvisioningAPI.md#GetBuildProvisioningStorageFiles) | **Get** /v3/buildProvisioning/storage/file | Get the list of all the files storage
[**GetBuildProvisioningStorageFilesByRegistrationId**](BuildProvisioningAPI.md#GetBuildProvisioningStorageFilesByRegistrationId) | **Get** /v3/buildProvisioning/storage/registration/{buildProvisioningRegistrationId}/file | Get the list of all the files in a build provisioning registration
[**GetBuildProvisioningStorageRegistration**](BuildProvisioningAPI.md#GetBuildProvisioningStorageRegistration) | **Get** /v3/buildProvisioning/storage/registration/{id} | Get the details of a build provisioning storage registration
[**GetBuildProvisioningStorageRegistrationTypes**](BuildProvisioningAPI.md#GetBuildProvisioningStorageRegistrationTypes) | **Get** /v3/buildProvisioning/storage/registration/types | Get The build provisioning storage registration types for storing application build files
[**GetBuildProvisioningStorageRegistrations**](BuildProvisioningAPI.md#GetBuildProvisioningStorageRegistrations) | **Get** /v3/buildProvisioning/storage/registration | Get the list of all build provisioning storage registrations for a customer



## GetBuildProvisioningStorageFile

> []BuildProvisioningFile GetBuildProvisioningStorageFile(ctx, id).Execute()

Get the details of the file in the storage

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
	id := "id_example" // string | The ID of the build provisioning file for which to fetch the information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageFile`: []BuildProvisioningFile
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the build provisioning file for which to fetch the information | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BuildProvisioningFile**](BuildProvisioningFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildProvisioningStorageFiles

> []BuildProvisioningFile GetBuildProvisioningStorageFiles(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get the list of all the files storage

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
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageFiles(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageFiles`: []BuildProvisioningFile
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]BuildProvisioningFile**](BuildProvisioningFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildProvisioningStorageFilesByRegistrationId

> []BuildProvisioningFile GetBuildProvisioningStorageFilesByRegistrationId(ctx, buildProvisionRegistrationId).RANGEDDATA(rANGEDDATA).Execute()

Get the list of all the files in a build provisioning registration

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
	buildProvisionRegistrationId := "buildProvisionRegistrationId_example" // string | The ID of the build provisioning registration where you want the files of
	rANGEDDATA := "rANGEDDATA_example" // string | Example header and default range: RANGED-DATA:start=0,results=25 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageFilesByRegistrationId(context.Background(), buildProvisionRegistrationId).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageFilesByRegistrationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageFilesByRegistrationId`: []BuildProvisioningFile
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageFilesByRegistrationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildProvisionRegistrationId** | **string** | The ID of the build provisioning registration where you want the files of | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageFilesByRegistrationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]BuildProvisioningFile**](BuildProvisioningFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildProvisioningStorageRegistration

> []BuildProvisioningRegistration GetBuildProvisioningStorageRegistration(ctx, id).Execute()

Get the details of a build provisioning storage registration

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
	id := "id_example" // string | The ID of the build provisioning storage registration for which to fetch the information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageRegistration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageRegistration`: []BuildProvisioningRegistration
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the build provisioning storage registration for which to fetch the information | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BuildProvisioningRegistration**](BuildProvisioningRegistration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildProvisioningStorageRegistrationTypes

> []BuildProvisioningRegistrationType GetBuildProvisioningStorageRegistrationTypes(ctx).Execute()

Get The build provisioning storage registration types for storing application build files

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
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageRegistrationTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageRegistrationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageRegistrationTypes`: []BuildProvisioningRegistrationType
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageRegistrationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageRegistrationTypesRequest struct via the builder pattern


### Return type

[**[]BuildProvisioningRegistrationType**](BuildProvisioningRegistrationType.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildProvisioningStorageRegistrations

> []BuildProvisioningRegistration GetBuildProvisioningStorageRegistrations(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get the list of all build provisioning storage registrations for a customer

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
	resp, r, err := apiClient.BuildProvisioningAPI.GetBuildProvisioningStorageRegistrations(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildProvisioningAPI.GetBuildProvisioningStorageRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildProvisioningStorageRegistrations`: []BuildProvisioningRegistration
	fmt.Fprintf(os.Stdout, "Response from `BuildProvisioningAPI.GetBuildProvisioningStorageRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildProvisioningStorageRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]BuildProvisioningRegistration**](BuildProvisioningRegistration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

