# \CustomFieldsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3CustomfieldsIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3CustomfieldsIdJson) | **Delete** /projects/api/v3/customfields/:id.json | Delete an existing custom field
[**DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson) | **Delete** /projects/api/v3/projects/:projectId/customfields/:id.json | Delete an existing project custom field value.
[**DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson**](CustomFieldsApi.md#DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson) | **Delete** /projects/api/v3/tasks/:taskId/customfields/:id.json | Delete an existing task custom field value.
[**GETProjectsApiV3CustomfieldsIdJson**](CustomFieldsApi.md#GETProjectsApiV3CustomfieldsIdJson) | **Get** /projects/api/v3/customfields/:id.json | Get a custom field by id.
[**GETProjectsApiV3CustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3CustomfieldsJson) | **Get** /projects/api/v3/customfields.json | Get all custom fields
[**GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson**](CustomFieldsApi.md#GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson) | **Get** /projects/api/v3/projects/:projectId/customfields/:id.json | Project custom field value.
[**GETProjectsApiV3ProjectsProjectIdCustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3ProjectsProjectIdCustomfieldsJson) | **Get** /projects/api/v3/projects/:projectId/customfields.json | Project custom field values.
[**GETProjectsApiV3TasksTaskIdCustomfieldsIdJson**](CustomFieldsApi.md#GETProjectsApiV3TasksTaskIdCustomfieldsIdJson) | **Get** /projects/api/v3/tasks/:taskId/customfields/:id.json | Task custom field value.
[**GETProjectsApiV3TasksTaskIdCustomfieldsJson**](CustomFieldsApi.md#GETProjectsApiV3TasksTaskIdCustomfieldsJson) | **Get** /projects/api/v3/tasks/:taskId/customfields.json | Task custom field values.
[**PATCHProjectsApiV3CustomfieldsIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3CustomfieldsIdJson) | **Patch** /projects/api/v3/customfields/:id.json | Update an existing custom field.
[**PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson) | **Patch** /projects/api/v3/projects/:projectId/customfields/:id.json | Update an existing project custom field value.
[**PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson**](CustomFieldsApi.md#PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson) | **Patch** /projects/api/v3/tasks/:taskId/customfields/:id.json | Update an existing task custom field value.
[**POSTProjectsApiV3CustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3CustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/customfields/bulk/delete.json | Delete many custom fields at once
[**POSTProjectsApiV3CustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3CustomfieldsJson) | **Post** /projects/api/v3/customfields.json | Create a new custom field
[**POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/projects/:projectId/customfields/bulk/delete.json | Delete many project custom fields values at once.
[**POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson) | **Post** /projects/api/v3/projects/:projectId/customfields/bulk/update.json | Update many project custom field values at once.
[**POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson) | **Post** /projects/api/v3/projects/:projectId/customfields.json | Add project custom field value.
[**POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson**](CustomFieldsApi.md#POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson) | **Post** /projects/api/v3/tasks/:taskId/customfields/bulk/delete.json | Delete many task custom fields values at once.
[**POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson**](CustomFieldsApi.md#POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson) | **Post** /projects/api/v3/tasks/:taskId/customfields/bulk/update.json | Update many task custom field values at once.
[**POSTProjectsApiV3TasksTaskIdCustomfieldsJson**](CustomFieldsApi.md#POSTProjectsApiV3TasksTaskIdCustomfieldsJson) | **Post** /projects/api/v3/tasks/:taskId/customfields.json | Add task custom field value.
[**PUTProjectsApiV3CustomfieldsIdJson**](CustomFieldsApi.md#PUTProjectsApiV3CustomfieldsIdJson) | **Put** /projects/api/v3/customfields/:id.json | Update an existing custom field.



## DELETEProjectsApiV3CustomfieldsIdJson

> DELETEProjectsApiV3CustomfieldsIdJson(ctx).Execute()

Delete an existing custom field



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
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3CustomfieldsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3CustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3CustomfieldsIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson

> DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(ctx).Execute()

Delete an existing project custom field value.



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
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3ProjectsProjectIdCustomfieldsIdJsonRequest struct via the builder pattern


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


## DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson

> DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson(ctx).Execute()

Delete an existing task custom field value.



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
    resp, r, err := api_client.CustomFieldsApi.DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.DELETEProjectsApiV3TasksTaskIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3TasksTaskIdCustomfieldsIdJsonRequest struct via the builder pattern


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


## GETProjectsApiV3CustomfieldsIdJson

> CustomfieldResponse GETProjectsApiV3CustomfieldsIdJson(ctx).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()

Get a custom field by id.



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
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    entities := "entities_example" // string | filter by entities (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlySiteLevel := true // bool | only return site-level custom fields (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeSiteLevel := true // bool | include custom fields for site-level in the response (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3CustomfieldsIdJson(context.Background()).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3CustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CustomfieldsIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3CustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CustomfieldsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **entities** | **string** | filter by entities | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlySiteLevel** | **bool** | only return site-level custom fields | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeSiteLevel** | **bool** | include custom fields for site-level in the response | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3CustomfieldsJson

> CustomfieldCustomFieldsResponse GETProjectsApiV3CustomfieldsJson(ctx).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()

Get all custom fields



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
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    entities := "entities_example" // string | filter by entities (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlySiteLevel := true // bool | only return site-level custom fields (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    includeSiteLevel := true // bool | include custom fields for site-level in the response (optional) (default to true)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3CustomfieldsJson(context.Background()).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).Entities(entities).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).OnlySiteLevel(onlySiteLevel).MatchAllProjectTags(matchAllProjectTags).IncludeSiteLevel(includeSiteLevel).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3CustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3CustomfieldsJson`: CustomfieldCustomFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3CustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3CustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **entities** | **string** | filter by entities | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlySiteLevel** | **bool** | only return site-level custom fields | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **includeSiteLevel** | **bool** | include custom fields for site-level in the response | [default to true]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**CustomfieldCustomFieldsResponse**](CustomfieldCustomFieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson

> ProjectResponse GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(ctx).Execute()

Project custom field value.



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
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdCustomfieldsIdJsonRequest struct via the builder pattern


### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdCustomfieldsJson

> ProjectCustomFieldProjectsResponse GETProjectsApiV3ProjectsProjectIdCustomfieldsJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Project custom field values.



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
    updatedAfter := time.Now() // time.Time | updated after (optional)
    searchTerm := "searchTerm_example" // string | filter by project name (optional)
    reportType := "reportType_example" // string | define the format of the report (optional) (default to "project")
    reportFormat := "reportFormat_example" // string | define the format of the report (optional) (default to "pdf")
    projectType := "projectType_example" // string | filter by project type (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "name")
    minLastActivityDate := time.Now() // string | filter by min last activity date (optional)
    maxLastActivityDate := time.Now() // string | filter by max last activity date (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (not used when generating reports) (optional) (default to 50)
    page := int32(56) // int32 | page number (not used when generating reports) (optional) (default to 1)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    onlyProjectsWithExplicitMembership := true // bool | only show projects with explicit membership (optional) (default to false)
    onlyArchivedProjects := true // bool | return only archived projects (optional) (default to false)
    matchAllProjectTags := true // bool | match all project tags (optional)
    isReportDownload := true // bool | generate a report document (optional) (default to false)
    includeProjectUserInfo := true // bool | fetch user-specific data such as isStarred (optional) (default to false)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedStatus := true // bool | optional to include completed projects when filtering by project statuses \"current,late\". (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    hideObservedProjects := true // bool | hide projects where the logged-in user is just an observer (optional) (default to false)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (not used when generating reports) (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectcategories := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectUpdates := []string{"Inner_example"} // []string |  (optional)
    fieldsProjectBudgets := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioColumns := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioCards := []string{"Inner_example"} // []string |  (optional)
    fieldsPortfolioBoards := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldProjects := []string{"Inner_example"} // []string |  (optional)
    fieldsCompanies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsProjectIdCustomfieldsJson`: ProjectCustomFieldProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3ProjectsProjectIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | updated after | 
 **searchTerm** | **string** | filter by project name | 
 **reportType** | **string** | define the format of the report | [default to &quot;project&quot;]
 **reportFormat** | **string** | define the format of the report | [default to &quot;pdf&quot;]
 **projectType** | **string** | filter by project type | 
 **projectStatuses** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;name&quot;]
 **minLastActivityDate** | **string** | filter by min last activity date | 
 **maxLastActivityDate** | **string** | filter by max last activity date | 
 **userId** | **int32** | filter by user id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page (not used when generating reports) | [default to 50]
 **page** | **int32** | page number (not used when generating reports) | [default to 1]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **onlyProjectsWithExplicitMembership** | **bool** | only show projects with explicit membership | [default to false]
 **onlyArchivedProjects** | **bool** | return only archived projects | [default to false]
 **matchAllProjectTags** | **bool** | match all project tags | 
 **isReportDownload** | **bool** | generate a report document | [default to false]
 **includeProjectUserInfo** | **bool** | fetch user-specific data such as isStarred | [default to false]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedStatus** | **bool** | optional to include completed projects when filtering by project statuses \&quot;current,late\&quot;. | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **hideObservedProjects** | **bool** | hide projects where the logged-in user is just an observer | [default to false]
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include (not used when generating reports) | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **fieldsProjectcategories** | **[]string** |  | 
 **fieldsProjectUpdates** | **[]string** |  | 
 **fieldsProjectBudgets** | **[]string** |  | 
 **fieldsPortfolioColumns** | **[]string** |  | 
 **fieldsPortfolioCards** | **[]string** |  | 
 **fieldsPortfolioBoards** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldProjects** | **[]string** |  | 
 **fieldsCompanies** | **[]string** |  | 

### Return type

[**ProjectCustomFieldProjectsResponse**](ProjectCustomFieldProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TasksTaskIdCustomfieldsIdJson

> TaskResponse GETProjectsApiV3TasksTaskIdCustomfieldsIdJson(ctx).Execute()

Task custom field value.



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
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsIdJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksTaskIdCustomfieldsIdJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksTaskIdCustomfieldsIdJsonRequest struct via the builder pattern


### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TasksTaskIdCustomfieldsJson

> TaskCustomFieldTasksResponse GETProjectsApiV3TasksTaskIdCustomfieldsJson(ctx).TaskId(taskId).PageSize(pageSize).Page(page).Include(include).FieldsTasks(fieldsTasks).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldTasks(fieldsCustomfieldTasks).CustomFieldIds(customFieldIds).Execute()

Task custom field values.



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
    taskId := int32(56) // int32 | filter by task id (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfields := []string{"Inner_example"} // []string |  (optional)
    fieldsCustomfieldTasks := []string{"Inner_example"} // []string |  (optional)
    customFieldIds := []int32{int32(123)} // []int32 | filter by custom field ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsJson(context.Background()).TaskId(taskId).PageSize(pageSize).Page(page).Include(include).FieldsTasks(fieldsTasks).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldTasks(fieldsCustomfieldTasks).CustomFieldIds(customFieldIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TasksTaskIdCustomfieldsJson`: TaskCustomFieldTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.GETProjectsApiV3TasksTaskIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TasksTaskIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **int32** | filter by task id | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **include** | **[]string** | include | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsCustomfields** | **[]string** |  | 
 **fieldsCustomfieldTasks** | **[]string** |  | 
 **customFieldIds** | **[]int32** | filter by custom field ids | 

### Return type

[**TaskCustomFieldTasksResponse**](TaskCustomFieldTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3CustomfieldsIdJson

> CustomfieldResponse PATCHProjectsApiV3CustomfieldsIdJson(ctx).CustomfieldRequest(customfieldRequest).Execute()

Update an existing custom field.



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
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3CustomfieldsIdJson(context.Background()).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3CustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3CustomfieldsIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3CustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3CustomfieldsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson

> ProjectResponse PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(ctx).ProjectRequest(projectRequest).Execute()

Update an existing project custom field value.



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
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson(context.Background()).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3ProjectsProjectIdCustomfieldsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson

> TaskResponse PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson(ctx).TaskRequest(taskRequest).Execute()

Update an existing task custom field value.



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
    taskRequest := *openapiclient.NewTaskRequest() // TaskRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson(context.Background()).TaskRequest(taskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PATCHProjectsApiV3TasksTaskIdCustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPATCHProjectsApiV3TasksTaskIdCustomfieldsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskRequest** | [**TaskRequest**](TaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3CustomfieldsBulkDeleteJson

> POSTProjectsApiV3CustomfieldsBulkDeleteJson(ctx).CustomfieldBulkDeleteRequest(customfieldBulkDeleteRequest).Execute()

Delete many custom fields at once



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
    customfieldBulkDeleteRequest := *openapiclient.NewCustomfieldBulkDeleteRequest() // CustomfieldBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3CustomfieldsBulkDeleteJson(context.Background()).CustomfieldBulkDeleteRequest(customfieldBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3CustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3CustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldBulkDeleteRequest** | [**CustomfieldBulkDeleteRequest**](CustomfieldBulkDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3CustomfieldsJson

> CustomfieldResponse POSTProjectsApiV3CustomfieldsJson(ctx).CustomfieldRequest(customfieldRequest).Execute()

Create a new custom field



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
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson(context.Background()).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3CustomfieldsJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3CustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3CustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson

> POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson(ctx).ProjectBulkDeleteRequest(projectBulkDeleteRequest).Execute()

Delete many project custom fields values at once.



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
    projectBulkDeleteRequest := *openapiclient.NewProjectBulkDeleteRequest() // ProjectBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson(context.Background()).ProjectBulkDeleteRequest(projectBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectBulkDeleteRequest** | [**ProjectBulkDeleteRequest**](ProjectBulkDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson

> ProjectCustomFieldProjectsResponse POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson(ctx).ProjectBulkUpdateRequest(projectBulkUpdateRequest).Execute()

Update many project custom field values at once.



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
    projectBulkUpdateRequest := *openapiclient.NewProjectBulkUpdateRequest() // ProjectBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson(context.Background()).ProjectBulkUpdateRequest(projectBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson`: ProjectCustomFieldProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsProjectIdCustomfieldsBulkUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectBulkUpdateRequest** | [**ProjectBulkUpdateRequest**](ProjectBulkUpdateRequest.md) |  | 

### Return type

[**ProjectCustomFieldProjectsResponse**](ProjectCustomFieldProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson

> ProjectResponse POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson(ctx).ProjectRequest(projectRequest).Execute()

Add project custom field value.



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
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson(context.Background()).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3ProjectsProjectIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3ProjectsProjectIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson

> POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson(ctx).TaskBulkDeleteRequest(taskBulkDeleteRequest).Execute()

Delete many task custom fields values at once.



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
    taskBulkDeleteRequest := *openapiclient.NewTaskBulkDeleteRequest() // TaskBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson(context.Background()).TaskBulkDeleteRequest(taskBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TasksTaskIdCustomfieldsBulkDeleteJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskBulkDeleteRequest** | [**TaskBulkDeleteRequest**](TaskBulkDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson

> TaskCustomFieldTasksResponse POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson(ctx).TaskBulkUpdateRequest(taskBulkUpdateRequest).Execute()

Update many task custom field values at once.



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
    taskBulkUpdateRequest := *openapiclient.NewTaskBulkUpdateRequest() // TaskBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson(context.Background()).TaskBulkUpdateRequest(taskBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson`: TaskCustomFieldTasksResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TasksTaskIdCustomfieldsBulkUpdateJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskBulkUpdateRequest** | [**TaskBulkUpdateRequest**](TaskBulkUpdateRequest.md) |  | 

### Return type

[**TaskCustomFieldTasksResponse**](TaskCustomFieldTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3TasksTaskIdCustomfieldsJson

> TaskResponse POSTProjectsApiV3TasksTaskIdCustomfieldsJson(ctx).TaskRequest(taskRequest).Execute()

Add task custom field value.



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
    taskRequest := *openapiclient.NewTaskRequest() // TaskRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsJson(context.Background()).TaskRequest(taskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3TasksTaskIdCustomfieldsJson`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.POSTProjectsApiV3TasksTaskIdCustomfieldsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3TasksTaskIdCustomfieldsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskRequest** | [**TaskRequest**](TaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3CustomfieldsIdJson

> CustomfieldResponse PUTProjectsApiV3CustomfieldsIdJson(ctx).CustomfieldRequest(customfieldRequest).Execute()

Update an existing custom field.



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
    customfieldRequest := *openapiclient.NewCustomfieldRequest() // CustomfieldRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomFieldsApi.PUTProjectsApiV3CustomfieldsIdJson(context.Background()).CustomfieldRequest(customfieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsApi.PUTProjectsApiV3CustomfieldsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3CustomfieldsIdJson`: CustomfieldResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFieldsApi.PUTProjectsApiV3CustomfieldsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3CustomfieldsIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customfieldRequest** | [**CustomfieldRequest**](CustomfieldRequest.md) |  | 

### Return type

[**CustomfieldResponse**](CustomfieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

