# \AppsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3AppsIdJson**](AppsApi.md#GETProjectsApiV3AppsIdJson) | **Get** /projects/api/v3/apps/:id.json | Get a specific app.
[**GETProjectsApiV3AppsJson**](AppsApi.md#GETProjectsApiV3AppsJson) | **Get** /projects/api/v3/apps.json | Get all apps.
[**POSTProjectsApiV3AppsAppidInstallJson**](AppsApi.md#POSTProjectsApiV3AppsAppidInstallJson) | **Post** /projects/api/v3/apps/:appid/install.json | Install an app onto an installation
[**POSTProjectsApiV3AppsIdUninstallJson**](AppsApi.md#POSTProjectsApiV3AppsIdUninstallJson) | **Post** /projects/api/v3/apps/:id/uninstall.json | Uninstall an app from an installation



## GETProjectsApiV3AppsIdJson

> AppResponse GETProjectsApiV3AppsIdJson(ctx).Execute()

Get a specific app.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.GETProjectsApiV3AppsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GETProjectsApiV3AppsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AppsIdJson`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GETProjectsApiV3AppsIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AppsIdJsonRequest struct via the builder pattern


### Return type

[**AppResponse**](AppResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3AppsJson

> AppAppsResponse GETProjectsApiV3AppsJson(ctx).PageSize(pageSize).Page(page).Include(include).Execute()

Get all apps.

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
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.GETProjectsApiV3AppsJson(context.Background()).PageSize(pageSize).Page(page).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GETProjectsApiV3AppsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AppsJson`: AppAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GETProjectsApiV3AppsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AppsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **include** | **[]string** | include | 

### Return type

[**AppAppsResponse**](AppAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3AppsAppidInstallJson

> AppResponse POSTProjectsApiV3AppsAppidInstallJson(ctx).Execute()

Install an app onto an installation

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.POSTProjectsApiV3AppsAppidInstallJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.POSTProjectsApiV3AppsAppidInstallJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3AppsAppidInstallJson`: AppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.POSTProjectsApiV3AppsAppidInstallJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AppsAppidInstallJsonRequest struct via the builder pattern


### Return type

[**AppResponse**](AppResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3AppsIdUninstallJson

> POSTProjectsApiV3AppsIdUninstallJson(ctx).Execute()

Uninstall an app from an installation

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.POSTProjectsApiV3AppsIdUninstallJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.POSTProjectsApiV3AppsIdUninstallJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3AppsIdUninstallJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

