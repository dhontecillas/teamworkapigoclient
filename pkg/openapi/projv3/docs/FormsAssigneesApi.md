# \FormsAssigneesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3FormsIdAssigneesJson**](FormsAssigneesApi.md#GETProjectsApiV3FormsIdAssigneesJson) | **Get** /projects/api/v3/forms/:id/assignees.json | Get all assignees for a given form.
[**PUTProjectsApiV3FormIdAssigneesJson**](FormsAssigneesApi.md#PUTProjectsApiV3FormIdAssigneesJson) | **Put** /projects/api/v3/form/:id/assignees.json | Update the existing assignees.



## GETProjectsApiV3FormsIdAssigneesJson

> AssigneeFormAssigneesResponse GETProjectsApiV3FormsIdAssigneesJson(ctx).OrderMode(orderMode).UserId(userId).PageSize(pageSize).Page(page).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsFormAssignees(fieldsFormAssignees).FieldsCompanies(fieldsCompanies).Execute()

Get all assignees for a given form.



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
    userId := int32(56) // int32 | filter by user id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTeams := []string{"Inner_example"} // []string |  (optional)
    fieldsFormAssignees := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsAssigneesApi.GETProjectsApiV3FormsIdAssigneesJson(context.Background()).OrderMode(orderMode).UserId(userId).PageSize(pageSize).Page(page).Include(include).FieldsUsers(fieldsUsers).FieldsTeams(fieldsTeams).FieldsFormAssignees(fieldsFormAssignees).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsAssigneesApi.GETProjectsApiV3FormsIdAssigneesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsIdAssigneesJson`: AssigneeFormAssigneesResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsAssigneesApi.GETProjectsApiV3FormsIdAssigneesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsIdAssigneesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **userId** | **int32** | filter by user id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTeams** | **[]string** |  | 
 **fieldsFormAssignees** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**AssigneeFormAssigneesResponse**](AssigneeFormAssigneesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3FormIdAssigneesJson

> AssigneeResponse PUTProjectsApiV3FormIdAssigneesJson(ctx).AssigneeRequest(assigneeRequest).Execute()

Update the existing assignees.



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
    assigneeRequest := *openapiclient.NewAssigneeRequest() // AssigneeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsAssigneesApi.PUTProjectsApiV3FormIdAssigneesJson(context.Background()).AssigneeRequest(assigneeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsAssigneesApi.PUTProjectsApiV3FormIdAssigneesJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3FormIdAssigneesJson`: AssigneeResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsAssigneesApi.PUTProjectsApiV3FormIdAssigneesJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3FormIdAssigneesJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assigneeRequest** | [**AssigneeRequest**](AssigneeRequest.md) |  | 

### Return type

[**AssigneeResponse**](AssigneeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

