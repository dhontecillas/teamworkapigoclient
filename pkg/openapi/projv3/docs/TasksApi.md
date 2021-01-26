# \TasksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3TasksMetricsCompleteJson**](TasksApi.md#GETProjectsApiV3TasksMetricsCompleteJson) | **Get** /projects/api/v3/tasks/metrics/complete.json | Total count of completed tasks
[**GETProjectsApiV3TasksMetricsLateJson**](TasksApi.md#GETProjectsApiV3TasksMetricsLateJson) | **Get** /projects/api/v3/tasks/metrics/late.json | Get total count of late tasks



## GETProjectsApiV3TasksMetricsCompleteJson

> CompleteResponse GETProjectsApiV3TasksMetricsCompleteJson(ctx).Execute()

Total count of completed tasks



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
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TasksMetricsCompleteJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TasksMetricsCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksMetricsCompleteJson`: CompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TasksMetricsCompleteJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksMetricsCompleteJsonRequest struct via the builder pattern


### Return type

[**CompleteResponse**](complete.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TasksMetricsLateJson

> LateResponse GETProjectsApiV3TasksMetricsLateJson(ctx).Execute()

Get total count of late tasks



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
    resp, r, err := api_client.TasksApi.GETProjectsApiV3TasksMetricsLateJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.GETProjectsApiV3TasksMetricsLateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksMetricsLateJson`: LateResponse
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.GETProjectsApiV3TasksMetricsLateJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksMetricsLateJsonRequest struct via the builder pattern


### Return type

[**LateResponse**](late.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

