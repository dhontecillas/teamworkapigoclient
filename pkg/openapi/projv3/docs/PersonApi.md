# \PersonApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3PeopleIdJson**](PersonApi.md#GETProjectsApiV3PeopleIdJson) | **Get** /projects/api/v3/people/:id.json | Returns a person



## GETProjectsApiV3PeopleIdJson

> PeopleResponse GETProjectsApiV3PeopleIdJson(ctx).OrderMode(orderMode).OrderBy(orderBy).PageSize(pageSize).Page(page).Include(include).Ids(ids).FieldsCompanies(fieldsCompanies).Execute()

Returns a person



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
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    ids := []int32{int32(123)} // []int32 | filter by user ids (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.GETProjectsApiV3PeopleIdJson(context.Background()).OrderMode(orderMode).OrderBy(orderBy).PageSize(pageSize).Page(page).Include(include).Ids(ids).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.GETProjectsApiV3PeopleIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3PeopleIdJson`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.GETProjectsApiV3PeopleIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3PeopleIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **include** | **[]string** | include (not used when generating reports) | 
 **ids** | **[]int32** | filter by user ids | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**PeopleResponse**](PeopleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

