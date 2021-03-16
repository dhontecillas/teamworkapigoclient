# \FormsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3FormsIdJson**](FormsApi.md#DELETEProjectsApiV3FormsIdJson) | **Delete** /projects/api/v3/forms/:id.json | Delete an existing form.
[**GETProjectsApiV3FormsIdDraftJson**](FormsApi.md#GETProjectsApiV3FormsIdDraftJson) | **Get** /projects/api/v3/forms/:id/draft.json | Get the draft version of a specific form.
[**GETProjectsApiV3FormsIdJson**](FormsApi.md#GETProjectsApiV3FormsIdJson) | **Get** /projects/api/v3/forms/:id.json | Get a specific form.
[**GETProjectsApiV3FormsJson**](FormsApi.md#GETProjectsApiV3FormsJson) | **Get** /projects/api/v3/forms.json | Get all forms.
[**PATCHProjectsApiV3FormsIdJson**](FormsApi.md#PATCHProjectsApiV3FormsIdJson) | **Patch** /projects/api/v3/forms/:id.json | Update an existing form.
[**POSTProjectsApiV3FormsJson**](FormsApi.md#POSTProjectsApiV3FormsJson) | **Post** /projects/api/v3/forms.json | Create a new form.



## DELETEProjectsApiV3FormsIdJson

> DELETEProjectsApiV3FormsIdJson(ctx).Execute()

Delete an existing form.

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
    resp, r, err := api_client.FormsApi.DELETEProjectsApiV3FormsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.DELETEProjectsApiV3FormsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3FormsIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3FormsIdDraftJson

> FormResponse GETProjectsApiV3FormsIdDraftJson(ctx).Execute()

Get the draft version of a specific form.

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
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsIdDraftJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsIdDraftJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsIdDraftJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsIdDraftJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsIdDraftJsonRequest struct via the builder pattern


### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FormsIdJson

> FormResponse GETProjectsApiV3FormsIdJson(ctx).Execute()

Get a specific form.

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
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsIdJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsIdJsonRequest struct via the builder pattern


### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3FormsJson

> FormFormsResponse GETProjectsApiV3FormsJson(ctx).Execute()

Get all forms.



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
    resp, r, err := api_client.FormsApi.GETProjectsApiV3FormsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.GETProjectsApiV3FormsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsJson`: FormFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.GETProjectsApiV3FormsJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsJsonRequest struct via the builder pattern


### Return type

[**FormFormsResponse**](FormFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3FormsIdJson

> FormResponse PATCHProjectsApiV3FormsIdJson(ctx).FormRequest(formRequest).Execute()

Update an existing form.

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
    formRequest := *openapiclient.NewFormRequest() // FormRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.PATCHProjectsApiV3FormsIdJson(context.Background()).FormRequest(formRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.PATCHProjectsApiV3FormsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3FormsIdJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.PATCHProjectsApiV3FormsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3FormsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formRequest** | [**FormRequest**](FormRequest.md) |  | 

### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FormsJson

> FormResponse POSTProjectsApiV3FormsJson(ctx).FormRequest(formRequest).Execute()

Create a new form.

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
    formRequest := *openapiclient.NewFormRequest() // FormRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.POSTProjectsApiV3FormsJson(context.Background()).FormRequest(formRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.POSTProjectsApiV3FormsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FormsJson`: FormResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.POSTProjectsApiV3FormsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FormsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formRequest** | [**FormRequest**](FormRequest.md) |  | 

### Return type

[**FormResponse**](FormResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

