# \AppsSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PUTProjectsApiV3AppsAppidSettingsSettingidJson**](AppsSettingsApi.md#PUTProjectsApiV3AppsAppidSettingsSettingidJson) | **Put** /projects/api/v3/apps/:appid/settings/:settingid.json | Update an existing setting.



## PUTProjectsApiV3AppsAppidSettingsSettingidJson

> SettingResponse PUTProjectsApiV3AppsAppidSettingsSettingidJson(ctx).SettingRequest(settingRequest).Execute()

Update an existing setting.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    settingRequest := *openapiclient.NewSettingRequest() // SettingRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsSettingsApi.PUTProjectsApiV3AppsAppidSettingsSettingidJson(context.Background()).SettingRequest(settingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsSettingsApi.PUTProjectsApiV3AppsAppidSettingsSettingidJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3AppsAppidSettingsSettingidJson`: SettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsSettingsApi.PUTProjectsApiV3AppsAppidSettingsSettingidJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3AppsAppidSettingsSettingidJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingRequest** | [**SettingRequest**](SettingRequest.md) |  | 

### Return type

[**SettingResponse**](SettingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

