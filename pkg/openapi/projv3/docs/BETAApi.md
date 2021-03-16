# \BETAApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ColumnsIdJson**](BETAApi.md#GETProjectsApiV3ColumnsIdJson) | **Get** /projects/api/v3/columns/:id.json | Get a specific column.
[**GETProjectsApiV3RatesInstallationJson**](BETAApi.md#GETProjectsApiV3RatesInstallationJson) | **Get** /projects/api/v3/rates/installation.json | Get an installation default rate.
[**GETProjectsApiV3RatesInstallationUsersJson**](BETAApi.md#GETProjectsApiV3RatesInstallationUsersJson) | **Get** /projects/api/v3/rates/installation/users.json | Get a specific rates.
[**GETProjectsApiV3RatesProjectsProjectIdJson**](BETAApi.md#GETProjectsApiV3RatesProjectsProjectIdJson) | **Get** /projects/api/v3/rates/projects/:projectId.json | Get a projects default rate.
[**GETProjectsApiV3RatesUserUserIdJson**](BETAApi.md#GETProjectsApiV3RatesUserUserIdJson) | **Get** /projects/api/v3/rates/user/:userId.json | Get a user default rate.
[**PUTProjectsApiV3CostsUsersIdJson**](BETAApi.md#PUTProjectsApiV3CostsUsersIdJson) | **Put** /projects/api/v3/costs/users/:id.json | set a usercost.
[**PUTProjectsApiV3RatesInstallationJson**](BETAApi.md#PUTProjectsApiV3RatesInstallationJson) | **Put** /projects/api/v3/rates/installation.json | set an installation default rate.
[**PUTProjectsApiV3RatesProjectsProjectIdJson**](BETAApi.md#PUTProjectsApiV3RatesProjectsProjectIdJson) | **Put** /projects/api/v3/rates/projects/:projectId.json | set a project default rate.
[**PUTProjectsApiV3RatesUsersUserIdJson**](BETAApi.md#PUTProjectsApiV3RatesUsersUserIdJson) | **Put** /projects/api/v3/rates/users/:userId.json | set a user default rate.



## GETProjectsApiV3ColumnsIdJson

> ColumnResponse GETProjectsApiV3ColumnsIdJson(ctx).Execute()

Get a specific column.



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
    resp, r, err := api_client.BETAApi.GETProjectsApiV3ColumnsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3ColumnsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ColumnsIdJson`: ColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3ColumnsIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ColumnsIdJsonRequest struct via the builder pattern


### Return type

[**ColumnResponse**](ColumnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesInstallationJson

> RatesRateResponse GETProjectsApiV3RatesInstallationJson(ctx).Execute()

Get an installation default rate.



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
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesInstallationJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesInstallationJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesInstallationJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesInstallationJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesInstallationJsonRequest struct via the builder pattern


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesInstallationUsersJson

> RatesInstallationUserRateResponse GETProjectsApiV3RatesInstallationUsersJson(ctx).PageSize(pageSize).Page(page).Execute()

Get a specific rates.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesInstallationUsersJson(context.Background()).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesInstallationUsersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesInstallationUsersJson`: RatesInstallationUserRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesInstallationUsersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesInstallationUsersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]

### Return type

[**RatesInstallationUserRateResponse**](RatesInstallationUserRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesProjectsProjectIdJson

> RatesRateResponse GETProjectsApiV3RatesProjectsProjectIdJson(ctx).Execute()

Get a projects default rate.



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
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesProjectsProjectIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesProjectsProjectIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesProjectsProjectIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesProjectsProjectIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesProjectsProjectIdJsonRequest struct via the builder pattern


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3RatesUserUserIdJson

> RatesRateResponse GETProjectsApiV3RatesUserUserIdJson(ctx).Execute()

Get a user default rate.



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
    resp, r, err := api_client.BETAApi.GETProjectsApiV3RatesUserUserIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.GETProjectsApiV3RatesUserUserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3RatesUserUserIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.GETProjectsApiV3RatesUserUserIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3RatesUserUserIdJsonRequest struct via the builder pattern


### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3CostsUsersIdJson

> UsercostResponse PUTProjectsApiV3CostsUsersIdJson(ctx).UsercostRequest(usercostRequest).Execute()

set a usercost.



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
    usercostRequest := *openapiclient.NewUsercostRequest() // UsercostRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3CostsUsersIdJson(context.Background()).UsercostRequest(usercostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3CostsUsersIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3CostsUsersIdJson`: UsercostResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3CostsUsersIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3CostsUsersIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usercostRequest** | [**UsercostRequest**](UsercostRequest.md) |  | 

### Return type

[**UsercostResponse**](UsercostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesInstallationJson

> RatesRateResponse PUTProjectsApiV3RatesInstallationJson(ctx).RatesRequest(ratesRequest).Execute()

set an installation default rate.



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
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesInstallationJson(context.Background()).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesInstallationJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesInstallationJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesInstallationJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesInstallationJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesProjectsProjectIdJson

> RatesRateResponse PUTProjectsApiV3RatesProjectsProjectIdJson(ctx).RatesRequest(ratesRequest).Execute()

set a project default rate.



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
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesProjectsProjectIdJson(context.Background()).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesProjectsProjectIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesProjectsProjectIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesProjectsProjectIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesProjectsProjectIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3RatesUsersUserIdJson

> RatesRateResponse PUTProjectsApiV3RatesUsersUserIdJson(ctx).RatesRequest(ratesRequest).Execute()

set a user default rate.



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
    ratesRequest := *openapiclient.NewRatesRequest() // RatesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BETAApi.PUTProjectsApiV3RatesUsersUserIdJson(context.Background()).RatesRequest(ratesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BETAApi.PUTProjectsApiV3RatesUsersUserIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3RatesUsersUserIdJson`: RatesRateResponse
    fmt.Fprintf(os.Stdout, "Response from `BETAApi.PUTProjectsApiV3RatesUsersUserIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3RatesUsersUserIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratesRequest** | [**RatesRequest**](RatesRequest.md) |  | 

### Return type

[**RatesRateResponse**](RatesRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

