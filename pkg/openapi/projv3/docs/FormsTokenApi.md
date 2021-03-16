# \FormsTokenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PATCHProjectsApiV3FormsIdTokenJson**](FormsTokenApi.md#PATCHProjectsApiV3FormsIdTokenJson) | **Patch** /projects/api/v3/forms/:id/token.json | Update an existing token.
[**PUTProjectsApiV3FormsIdTokenRefreshJson**](FormsTokenApi.md#PUTProjectsApiV3FormsIdTokenRefreshJson) | **Put** /projects/api/v3/forms/:id/token/refresh.json | Refresh the value of a token



## PATCHProjectsApiV3FormsIdTokenJson

> TokenResponse PATCHProjectsApiV3FormsIdTokenJson(ctx).TokenRequest(tokenRequest).Execute()

Update an existing token.



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
    tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsTokenApi.PATCHProjectsApiV3FormsIdTokenJson(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsTokenApi.PATCHProjectsApiV3FormsIdTokenJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3FormsIdTokenJson`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsTokenApi.PATCHProjectsApiV3FormsIdTokenJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3FormsIdTokenJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3FormsIdTokenRefreshJson

> TokenResponse PUTProjectsApiV3FormsIdTokenRefreshJson(ctx).TokenRequest(tokenRequest).Execute()

Refresh the value of a token



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
    tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsTokenApi.PUTProjectsApiV3FormsIdTokenRefreshJson(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsTokenApi.PUTProjectsApiV3FormsIdTokenRefreshJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3FormsIdTokenRefreshJson`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsTokenApi.PUTProjectsApiV3FormsIdTokenRefreshJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3FormsIdTokenRefreshJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

