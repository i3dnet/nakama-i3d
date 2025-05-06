# \CrashEventAPI

All URIs are relative to *https://api.i3d.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrashes**](CrashEventAPI.md#GetCrashes) | **Get** /v3/crash | Get a list of Application Instance crash events



## GetCrashes

> []CrashEvent GetCrashes(ctx).RANGEDDATA(rANGEDDATA).Filter(filter).Execute()

Get a list of Application Instance crash events

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
	filter := "filter_example" // string | You can filter based on following parameters<br /> - deploymentEnvironmentId<br /> - fleetId<br /> - hostId<br /> - applicationId<br /> - applicationBuildId<br /> - applicationInstanceId<br /> <p>The parameters can have a single value (e.g: fleetId=1734197255512241679) or a list of values (e.g: applicationId=8881688742569018542,76243584248556249)<br /> A filter example is: applicationId=8881688742569018542,76243584248556249&fleetId=1734197255512241679<br /> Remember to urlencode the provided filter<br /> E.g filter=applicationId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679<br /> The final endpoint call would look like: /crash?filter=applicationId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrashEventAPI.GetCrashes(context.Background()).RANGEDDATA(rANGEDDATA).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrashEventAPI.GetCrashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrashes`: []CrashEvent
	fmt.Fprintf(os.Stdout, "Response from `CrashEventAPI.GetCrashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rANGEDDATA** | **string** | Example header and default range: RANGED-DATA:start&#x3D;0,results&#x3D;25 | 
 **filter** | **string** | You can filter based on following parameters&lt;br /&gt; - deploymentEnvironmentId&lt;br /&gt; - fleetId&lt;br /&gt; - hostId&lt;br /&gt; - applicationId&lt;br /&gt; - applicationBuildId&lt;br /&gt; - applicationInstanceId&lt;br /&gt; &lt;p&gt;The parameters can have a single value (e.g: fleetId&#x3D;1734197255512241679) or a list of values (e.g: applicationId&#x3D;8881688742569018542,76243584248556249)&lt;br /&gt; A filter example is: applicationId&#x3D;8881688742569018542,76243584248556249&amp;fleetId&#x3D;1734197255512241679&lt;br /&gt; Remember to urlencode the provided filter&lt;br /&gt; E.g filter&#x3D;applicationId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;br /&gt; The final endpoint call would look like: /crash?filter&#x3D;applicationId%3D8881688742569018542%2C76243584248556249%26fleetId%3D1734197255512241679&lt;/p&gt; | 

### Return type

[**[]CrashEvent**](CrashEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

