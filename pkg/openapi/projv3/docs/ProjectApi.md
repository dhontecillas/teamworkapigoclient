# \ProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsIdFeaturesorderJson**](ProjectApi.md#GETProjectsApiV3ProjectsIdFeaturesorderJson) | **Get** /projects/api/v3/projects/:id/featuresorder.json | Returns a project features order to display in tab
[**GETProjectsApiV3ProjectsIdJson**](ProjectApi.md#GETProjectsApiV3ProjectsIdJson) | **Get** /projects/api/v3/projects/:id.json | Returns a project
[**PUTProjectsApiV3ProjectsFeaturesorderJson**](ProjectApi.md#PUTProjectsApiV3ProjectsFeaturesorderJson) | **Put** /projects/api/v3/projects/featuresorder.json | Sets the default features order to display in tab
[**PUTProjectsApiV3ProjectsIdFeaturesorderJson**](ProjectApi.md#PUTProjectsApiV3ProjectsIdFeaturesorderJson) | **Put** /projects/api/v3/projects/:id/featuresorder.json | Sets the the features order to display in tab



## GETProjectsApiV3ProjectsIdFeaturesorderJson

> ProjectFeatureOrderResponse GETProjectsApiV3ProjectsIdFeaturesorderJson(ctx).Execute()

Returns a project features order to display in tab

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
    resp, r, err := api_client.ProjectApi.GETProjectsApiV3ProjectsIdFeaturesorderJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GETProjectsApiV3ProjectsIdFeaturesorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsIdFeaturesorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GETProjectsApiV3ProjectsIdFeaturesorderJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsIdFeaturesorderJsonRequest struct via the builder pattern


### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsIdJson

> ProjectSingleResponse GETProjectsApiV3ProjectsIdJson(ctx).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()

Returns a project



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
    resp, r, err := api_client.ProjectApi.GETProjectsApiV3ProjectsIdJson(context.Background()).UpdatedAfter(updatedAfter).SearchTerm(searchTerm).ReportType(reportType).ReportFormat(reportFormat).ProjectType(projectType).ProjectStatuses(projectStatuses).OrderMode(orderMode).OrderBy(orderBy).MinLastActivityDate(minLastActivityDate).MaxLastActivityDate(maxLastActivityDate).UserId(userId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).OnlyStarredProjects(onlyStarredProjects).OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership).OnlyArchivedProjects(onlyArchivedProjects).MatchAllProjectTags(matchAllProjectTags).IsReportDownload(isReportDownload).IncludeProjectUserInfo(includeProjectUserInfo).IncludeCustomFields(includeCustomFields).IncludeCompletedStatus(includeCompletedStatus).IncludeArchivedProjects(includeArchivedProjects).HideObservedProjects(hideObservedProjects).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).FieldsProjectcategories(fieldsProjectcategories).FieldsProjectUpdates(fieldsProjectUpdates).FieldsProjectBudgets(fieldsProjectBudgets).FieldsPortfolioColumns(fieldsPortfolioColumns).FieldsPortfolioCards(fieldsPortfolioCards).FieldsPortfolioBoards(fieldsPortfolioBoards).FieldsCustomfields(fieldsCustomfields).FieldsCustomfieldProjects(fieldsCustomfieldProjects).FieldsCompanies(fieldsCompanies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GETProjectsApiV3ProjectsIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsIdJson`: ProjectSingleResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GETProjectsApiV3ProjectsIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsIdJsonRequest struct via the builder pattern


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

[**ProjectSingleResponse**](ProjectSingleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3ProjectsFeaturesorderJson

> ProjectFeatureOrderResponse PUTProjectsApiV3ProjectsFeaturesorderJson(ctx).Execute()

Sets the default features order to display in tab

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
    resp, r, err := api_client.ProjectApi.PUTProjectsApiV3ProjectsFeaturesorderJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.PUTProjectsApiV3ProjectsFeaturesorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3ProjectsFeaturesorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.PUTProjectsApiV3ProjectsFeaturesorderJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3ProjectsFeaturesorderJsonRequest struct via the builder pattern


### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3ProjectsIdFeaturesorderJson

> ProjectFeatureOrderResponse PUTProjectsApiV3ProjectsIdFeaturesorderJson(ctx).Execute()

Sets the the features order to display in tab

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
    resp, r, err := api_client.ProjectApi.PUTProjectsApiV3ProjectsIdFeaturesorderJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.PUTProjectsApiV3ProjectsIdFeaturesorderJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3ProjectsIdFeaturesorderJson`: ProjectFeatureOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.PUTProjectsApiV3ProjectsIdFeaturesorderJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3ProjectsIdFeaturesorderJsonRequest struct via the builder pattern


### Return type

[**ProjectFeatureOrderResponse**](ProjectFeatureOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

