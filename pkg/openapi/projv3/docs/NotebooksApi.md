# \NotebooksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3NotebooksIdJson**](NotebooksApi.md#DELETEProjectsApiV3NotebooksIdJson) | **Delete** /projects/api/v3/notebooks/:id.json | Delete an existing notebook.
[**DELETEProjectsApiV3NotebooksIdVersionsJson**](NotebooksApi.md#DELETEProjectsApiV3NotebooksIdVersionsJson) | **Delete** /projects/api/v3/notebooks/:id/versions.json | Delete notebook versions.
[**GETProjectsApiV3NotebooksIdJson**](NotebooksApi.md#GETProjectsApiV3NotebooksIdJson) | **Get** /projects/api/v3/notebooks/:id.json | Get a specific notebook.
[**GETProjectsApiV3NotebooksIdVersionsJson**](NotebooksApi.md#GETProjectsApiV3NotebooksIdVersionsJson) | **Get** /projects/api/v3/notebooks/:id/versions.json | Get notebook versions
[**GETProjectsApiV3NotebooksJson**](NotebooksApi.md#GETProjectsApiV3NotebooksJson) | **Get** /projects/api/v3/notebooks.json | Get a list of notebooks.
[**GETProjectsApiV3NotebooksNIdVersionsVIdJson**](NotebooksApi.md#GETProjectsApiV3NotebooksNIdVersionsVIdJson) | **Get** /projects/api/v3/notebooks/:nId/versions/:vId.json | Get a notebook version
[**PATCHProjectsApiV3NotebooksIdJson**](NotebooksApi.md#PATCHProjectsApiV3NotebooksIdJson) | **Patch** /projects/api/v3/notebooks/:id.json | Edits a notebook
[**POSTProjectsApiV3ProjectsProjectIdNotebooksJson**](NotebooksApi.md#POSTProjectsApiV3ProjectsProjectIdNotebooksJson) | **Post** /projects/api/v3/projects/:projectId/notebooks.json | Create a new notebook.
[**PUTProjectsApiV3NotebooksIdLockJson**](NotebooksApi.md#PUTProjectsApiV3NotebooksIdLockJson) | **Put** /projects/api/v3/notebooks/:id/lock.json | Lock a notebook
[**PUTProjectsApiV3NotebooksIdUnlockJson**](NotebooksApi.md#PUTProjectsApiV3NotebooksIdUnlockJson) | **Put** /projects/api/v3/notebooks/:id/unlock.json | Unlock a notebook



## DELETEProjectsApiV3NotebooksIdJson

> DELETEProjectsApiV3NotebooksIdJson(ctx).Execute()

Delete an existing notebook.

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
    resp, r, err := api_client.NotebooksApi.DELETEProjectsApiV3NotebooksIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DELETEProjectsApiV3NotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3NotebooksIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3NotebooksIdVersionsJson

> DELETEProjectsApiV3NotebooksIdVersionsJson(ctx).Execute()

Delete notebook versions.



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
    resp, r, err := api_client.NotebooksApi.DELETEProjectsApiV3NotebooksIdVersionsJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DELETEProjectsApiV3NotebooksIdVersionsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3NotebooksIdVersionsJsonRequest struct via the builder pattern


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


## GETProjectsApiV3NotebooksIdJson

> NotebookResponse GETProjectsApiV3NotebooksIdJson(ctx).Execute()

Get a specific notebook.

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
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksIdJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksIdJsonRequest struct via the builder pattern


### Return type

[**NotebookResponse**](notebook.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksIdVersionsJson

> NotebookVersionsResponse GETProjectsApiV3NotebooksIdVersionsJson(ctx).Include(include).FieldsUsers(fieldsUsers).Execute()

Get notebook versions



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
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksIdVersionsJson(context.Background()).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksIdVersionsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksIdVersionsJson`: NotebookVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksIdVersionsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksIdVersionsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**NotebookVersionsResponse**](notebook.VersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksJson

> NotebookNotebooksResponse GETProjectsApiV3NotebooksJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).CreatedAfterDate(createdAfterDate).CreatedAfter(createdAfter).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).SecureOnly(secureOnly).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).LockedOnly(lockedOnly).GetEmoji(getEmoji).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebookCategories(fieldsNotebookCategories).CreatorIds(creatorIds).CategoryIds(categoryIds).Execute()

Get a list of notebooks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    updatedAfterDate := time.Now() // time.Time | filter by updated after date (deprecated, use updatedAfter) (optional)
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    searchTerm := "searchTerm_example" // string | filter by notebook name and description (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "dateUpdated")
    createdAfterDate := time.Now() // string | filter by created after date (deprecated, use createdAfter) (optional)
    createdAfter := time.Now() // string | filter by created after date (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    secureOnly := true // bool | filter by showing only secure notebooks (optional) (default to false)
    matchAllTags := true // bool | match all notebook tags (optional)
    matchAllProjectTags := true // bool | match all project tags (optional)
    lockedOnly := true // bool | filter by showing only locked notebooks (optional) (default to false)
    getEmoji := true // bool | parse Emoji's (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by notebook tag ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    ids := []int32{int32(123)} // []int32 | filter by ids (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsNotebookCategories := []string{"Inner_example"} // []string |  (optional)
    creatorIds := []int32{int32(123)} // []int32 | filter by creators' ids (optional)
    categoryIds := []int32{int32(123)} // []int32 | filter by notebook category ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).CreatedAfterDate(createdAfterDate).CreatedAfter(createdAfter).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).SecureOnly(secureOnly).MatchAllTags(matchAllTags).MatchAllProjectTags(matchAllProjectTags).LockedOnly(lockedOnly).GetEmoji(getEmoji).TagIds(tagIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).Ids(ids).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsNotebookCategories(fieldsNotebookCategories).CreatorIds(creatorIds).CategoryIds(categoryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksJson`: NotebookNotebooksResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **searchTerm** | **string** | filter by notebook name and description | 
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;dateUpdated&quot;]
 **createdAfterDate** | **string** | filter by created after date (deprecated, use createdAfter) | 
 **createdAfter** | **string** | filter by created after date | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **secureOnly** | **bool** | filter by showing only secure notebooks | [default to false]
 **matchAllTags** | **bool** | match all notebook tags | 
 **matchAllProjectTags** | **bool** | match all project tags | 
 **lockedOnly** | **bool** | filter by showing only locked notebooks | [default to false]
 **getEmoji** | **bool** | parse Emoji&#39;s | 
 **tagIds** | **[]int32** | filter by notebook tag ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **ids** | **[]int32** | filter by ids | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsNotebookCategories** | **[]string** |  | 
 **creatorIds** | **[]int32** | filter by creators&#39; ids | 
 **categoryIds** | **[]int32** | filter by notebook category ids | 

### Return type

[**NotebookNotebooksResponse**](notebook.NotebooksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3NotebooksNIdVersionsVIdJson

> NotebookVersionResponse GETProjectsApiV3NotebooksNIdVersionsVIdJson(ctx).Include(include).FieldsUsers(fieldsUsers).Execute()

Get a notebook version



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
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.GETProjectsApiV3NotebooksNIdVersionsVIdJson(context.Background()).Include(include).FieldsUsers(fieldsUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GETProjectsApiV3NotebooksNIdVersionsVIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3NotebooksNIdVersionsVIdJson`: NotebookVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.GETProjectsApiV3NotebooksNIdVersionsVIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3NotebooksNIdVersionsVIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 

### Return type

[**NotebookVersionResponse**](notebook.VersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3NotebooksIdJson

> NotebookResponse PATCHProjectsApiV3NotebooksIdJson(ctx).NotebookRequest(notebookRequest).Execute()

Edits a notebook



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
    notebookRequest := *openapiclient.NewNotebookRequest() // NotebookRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.PATCHProjectsApiV3NotebooksIdJson(context.Background()).NotebookRequest(notebookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PATCHProjectsApiV3NotebooksIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3NotebooksIdJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.PATCHProjectsApiV3NotebooksIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3NotebooksIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notebookRequest** | [**NotebookRequest**](NotebookRequest.md) |  | 

### Return type

[**NotebookResponse**](notebook.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsProjectIdNotebooksJson

> NotebookResponse POSTProjectsApiV3ProjectsProjectIdNotebooksJson(ctx).NotebookRequest(notebookRequest).Execute()

Create a new notebook.

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
    notebookRequest := *openapiclient.NewNotebookRequest() // NotebookRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotebooksApi.POSTProjectsApiV3ProjectsProjectIdNotebooksJson(context.Background()).NotebookRequest(notebookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.POSTProjectsApiV3ProjectsProjectIdNotebooksJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsProjectIdNotebooksJson`: NotebookResponse
    fmt.Fprintf(os.Stdout, "Response from `NotebooksApi.POSTProjectsApiV3ProjectsProjectIdNotebooksJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsProjectIdNotebooksJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notebookRequest** | [**NotebookRequest**](NotebookRequest.md) |  | 

### Return type

[**NotebookResponse**](notebook.Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3NotebooksIdLockJson

> PUTProjectsApiV3NotebooksIdLockJson(ctx).Execute()

Lock a notebook



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
    resp, r, err := api_client.NotebooksApi.PUTProjectsApiV3NotebooksIdLockJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTProjectsApiV3NotebooksIdLockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3NotebooksIdLockJsonRequest struct via the builder pattern


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


## PUTProjectsApiV3NotebooksIdUnlockJson

> PUTProjectsApiV3NotebooksIdUnlockJson(ctx).Execute()

Unlock a notebook



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
    resp, r, err := api_client.NotebooksApi.PUTProjectsApiV3NotebooksIdUnlockJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.PUTProjectsApiV3NotebooksIdUnlockJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3NotebooksIdUnlockJsonRequest struct via the builder pattern


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

