# \FormsPublicApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3FormsPublicTokenJson**](FormsPublicApi.md#GETProjectsApiV3FormsPublicTokenJson) | **Get** /projects/api/v3/forms/public/:token.json | Get a forms via its token.
[**POSTProjectsApiV3FormsPublicTokenSubmitJson**](FormsPublicApi.md#POSTProjectsApiV3FormsPublicTokenSubmitJson) | **Post** /projects/api/v3/forms/public/:token/submit.json | Submit a form response



## GETProjectsApiV3FormsPublicTokenJson

> FormPublicResponse GETProjectsApiV3FormsPublicTokenJson(ctx).OrderMode(orderMode).ContentState(contentState).UserId(userId).PageSize(pageSize).Page(page).ProjectIds(projectIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsForms(fieldsForms).Execute()

Get a forms via its token.



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
    contentState := "contentState_example" // string | query by form state (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsForms := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsPublicApi.GETProjectsApiV3FormsPublicTokenJson(context.Background()).OrderMode(orderMode).ContentState(contentState).UserId(userId).PageSize(pageSize).Page(page).ProjectIds(projectIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).FieldsForms(fieldsForms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsPublicApi.GETProjectsApiV3FormsPublicTokenJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3FormsPublicTokenJson`: FormPublicResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsPublicApi.GETProjectsApiV3FormsPublicTokenJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3FormsPublicTokenJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **contentState** | **string** | query by form state | 
 **userId** | **int32** | filter by user id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **projectIds** | **[]int32** | filter by project ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsForms** | **[]string** |  | 

### Return type

[**FormPublicResponse**](FormPublicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3FormsPublicTokenSubmitJson

> FormFormsResponse POSTProjectsApiV3FormsPublicTokenSubmitJson(ctx).FormSubmissionRequest(formSubmissionRequest).Execute()

Submit a form response



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
    formSubmissionRequest := *openapiclient.NewFormSubmissionRequest() // FormSubmissionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsPublicApi.POSTProjectsApiV3FormsPublicTokenSubmitJson(context.Background()).FormSubmissionRequest(formSubmissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsPublicApi.POSTProjectsApiV3FormsPublicTokenSubmitJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3FormsPublicTokenSubmitJson`: FormFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `FormsPublicApi.POSTProjectsApiV3FormsPublicTokenSubmitJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3FormsPublicTokenSubmitJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formSubmissionRequest** | [**FormSubmissionRequest**](FormSubmissionRequest.md) |  | 

### Return type

[**FormFormsResponse**](FormFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

