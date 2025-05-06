# \CloudDetailAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudProviderUcSetting**](CloudDetailAPI.md#CreateCloudProviderUcSetting) | **Post** /v3/cloud/provider/uc/setting | Create Ubisoft cloud settings
[**DeleteCloudProviderUcSetting**](CloudDetailAPI.md#DeleteCloudProviderUcSetting) | **Delete** /v3/cloud/provider/uc/setting/{ucSettingId} | Delete a given Ubisoft cloud setting
[**GetCloudProviderUcSetting**](CloudDetailAPI.md#GetCloudProviderUcSetting) | **Get** /v3/cloud/provider/uc/setting/{ucSettingId} | Get the detail of a Ubisoft cloud setting
[**GetCloudProviderUcSettings**](CloudDetailAPI.md#GetCloudProviderUcSettings) | **Get** /v3/cloud/provider/uc/setting | Get all Ubisoft cloud settings
[**UpdateCloudProviderUcSetting**](CloudDetailAPI.md#UpdateCloudProviderUcSetting) | **Put** /v3/cloud/provider/uc/setting/{ucSettingId} | Update the detail of a Ubisoft cloud setting



## CreateCloudProviderUcSetting

> []UCSetting CreateCloudProviderUcSetting(ctx).UCSetting(uCSetting).Execute()

Create Ubisoft cloud settings

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
	uCSetting := *openapiclient.NewUCSetting("Id_example", "ProjectId_example", "RegionName_example", "PublicNetworkId_example", "PrivateNetworkId_example", "SecurityGroupId_example", int32(123), int32(123)) // UCSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDetailAPI.CreateCloudProviderUcSetting(context.Background()).UCSetting(uCSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDetailAPI.CreateCloudProviderUcSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudProviderUcSetting`: []UCSetting
	fmt.Fprintf(os.Stdout, "Response from `CloudDetailAPI.CreateCloudProviderUcSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudProviderUcSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uCSetting** | [**UCSetting**](UCSetting.md) |  | 

### Return type

[**[]UCSetting**](UCSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudProviderUcSetting

> DeleteCloudProviderUcSetting(ctx, ucSettingId).Execute()

Delete a given Ubisoft cloud setting

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
	ucSettingId := int32(56) // int32 | The Id of the Ubisoft cloud setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudDetailAPI.DeleteCloudProviderUcSetting(context.Background(), ucSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDetailAPI.DeleteCloudProviderUcSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ucSettingId** | **int32** | The Id of the Ubisoft cloud setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudProviderUcSettingRequest struct via the builder pattern


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


## GetCloudProviderUcSetting

> []UCSetting GetCloudProviderUcSetting(ctx, ucSettingId).Execute()

Get the detail of a Ubisoft cloud setting

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
	ucSettingId := int32(56) // int32 | The Id of the Ubisoft cloud setting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDetailAPI.GetCloudProviderUcSetting(context.Background(), ucSettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDetailAPI.GetCloudProviderUcSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProviderUcSetting`: []UCSetting
	fmt.Fprintf(os.Stdout, "Response from `CloudDetailAPI.GetCloudProviderUcSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ucSettingId** | **int32** | The Id of the Ubisoft cloud setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderUcSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UCSetting**](UCSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderUcSettings

> []UCSetting GetCloudProviderUcSettings(ctx).RANGEDDATA(rANGEDDATA).Execute()

Get all Ubisoft cloud settings

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
	resp, r, err := apiClient.CloudDetailAPI.GetCloudProviderUcSettings(context.Background()).RANGEDDATA(rANGEDDATA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDetailAPI.GetCloudProviderUcSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProviderUcSettings`: []UCSetting
	fmt.Fprintf(os.Stdout, "Response from `CloudDetailAPI.GetCloudProviderUcSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderUcSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 

### Return type

[**[]UCSetting**](UCSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudProviderUcSetting

> []UCSetting UpdateCloudProviderUcSetting(ctx, ucSettingId).UCSetting(uCSetting).Execute()

Update the detail of a Ubisoft cloud setting

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
	ucSettingId := int32(56) // int32 | The Id of the Ubisoft cloud setting
	uCSetting := *openapiclient.NewUCSetting("Id_example", "ProjectId_example", "RegionName_example", "PublicNetworkId_example", "PrivateNetworkId_example", "SecurityGroupId_example", int32(123), int32(123)) // UCSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDetailAPI.UpdateCloudProviderUcSetting(context.Background(), ucSettingId).UCSetting(uCSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDetailAPI.UpdateCloudProviderUcSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudProviderUcSetting`: []UCSetting
	fmt.Fprintf(os.Stdout, "Response from `CloudDetailAPI.UpdateCloudProviderUcSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ucSettingId** | **int32** | The Id of the Ubisoft cloud setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudProviderUcSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uCSetting** | [**UCSetting**](UCSetting.md) |  | 

### Return type

[**[]UCSetting**](UCSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

