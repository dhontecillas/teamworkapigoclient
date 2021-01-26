# \CompanyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3CompaniesIdJson**](CompanyApi.md#GETProjectsApiV3CompaniesIdJson) | **Get** /projects/api/v3/companies/:id.json | Get a specific company.



## GETProjectsApiV3CompaniesIdJson

> CompanyResponse GETProjectsApiV3CompaniesIdJson(ctx).Execute()

Get a specific company.



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
    resp, r, err := api_client.CompanyApi.GETProjectsApiV3CompaniesIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GETProjectsApiV3CompaniesIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CompaniesIdJson`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GETProjectsApiV3CompaniesIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CompaniesIdJsonRequest struct via the builder pattern


### Return type

[**CompanyResponse**](company.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

