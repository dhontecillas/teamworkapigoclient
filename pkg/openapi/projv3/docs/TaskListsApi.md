# \TaskListsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV3ProjectsProjectIdTasklistsCsv**](TaskListsApi.md#GETProjectsApiV3ProjectsProjectIdTasklistsCsv) | **Get** /projects/api/v3/projects/:projectId/tasklists.csv | Generate tasklist report in CSV format
[**GETProjectsApiV3ProjectsProjectIdTasklistsHtml**](TaskListsApi.md#GETProjectsApiV3ProjectsProjectIdTasklistsHtml) | **Get** /projects/api/v3/projects/:projectId/tasklists.html | Generate tasklist report in HTML format
[**GETProjectsApiV3ProjectsProjectIdTasklistsPdf**](TaskListsApi.md#GETProjectsApiV3ProjectsProjectIdTasklistsPdf) | **Get** /projects/api/v3/projects/:projectId/tasklists.pdf | Generate tasklist report in PDF format
[**GETProjectsApiV3ProjectsProjectIdTasklistsXlsx**](TaskListsApi.md#GETProjectsApiV3ProjectsProjectIdTasklistsXlsx) | **Get** /projects/api/v3/projects/:projectId/tasklists.xlsx | Generate tasklist report in XLSX format



## GETProjectsApiV3ProjectsProjectIdTasklistsCsv

> GETProjectsApiV3ProjectsProjectIdTasklistsCsv(ctx).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate tasklist report in CSV format



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
    taskDate := "taskDate_example" // string | filter by task date (optional) (default to "anytime")
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "manual")
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    commentsOrderMode := "commentsOrderMode_example" // string | order mode of the comments (optional) (default to "asc")
    tasklistId := int32(56) // int32 | filter by task list (optional)
    projectId := int32(56) // int32 | filter by project (optional)
    onlyCompletedTasks := true // bool | only completed tasks (optional) (default to false)
    isReportDownload := true // bool | generate a report document (optional)
    includeTime := true // bool | include time (optional)
    includeTags := true // bool | include tags (optional) (default to true)
    includePrivateItems := true // bool | include private items (optional) (default to true)
    includeLateTasks := true // bool | include late tasks (optional)
    includeEstimatedTime := true // bool | include estimated time (optional) (default to true)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedTasks := true // bool | include completed tasks (optional) (default to true)
    includeComments := true // bool | include comments (optional) (default to false)
    includeCommentAvatars := true // bool | include comment avatars (optional)
    includeColumns := true // bool | include columns (optional) (default to true)
    includeAnytimeTasks := true // bool | include anytime tasks (optional)
    includeAllComments := true // bool | include all comments (optional) (default to false)
    ignoreStartDates := true // bool | ignore task start dates (optional)
    groupByTasklists := true // bool | group by task lists (optional)
    tags := []string{"Inner_example"} // []string | filter by tag names (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsCsv(context.Background()).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTasklistsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDate** | **string** | filter by task date | [default to &quot;anytime&quot;]
 **startDate** | **time.Time** | filter by start datetime | 
 **reportFormat** | **string** | define the format of the report | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;manual&quot;]
 **endDate** | **time.Time** | filter by end datetime | 
 **commentsOrderMode** | **string** | order mode of the comments | [default to &quot;asc&quot;]
 **tasklistId** | **int32** | filter by task list | 
 **projectId** | **int32** | filter by project | 
 **onlyCompletedTasks** | **bool** | only completed tasks | [default to false]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTime** | **bool** | include time | 
 **includeTags** | **bool** | include tags | [default to true]
 **includePrivateItems** | **bool** | include private items | [default to true]
 **includeLateTasks** | **bool** | include late tasks | 
 **includeEstimatedTime** | **bool** | include estimated time | [default to true]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedTasks** | **bool** | include completed tasks | [default to true]
 **includeComments** | **bool** | include comments | [default to false]
 **includeCommentAvatars** | **bool** | include comment avatars | 
 **includeColumns** | **bool** | include columns | [default to true]
 **includeAnytimeTasks** | **bool** | include anytime tasks | 
 **includeAllComments** | **bool** | include all comments | [default to false]
 **ignoreStartDates** | **bool** | ignore task start dates | 
 **groupByTasklists** | **bool** | group by task lists | 
 **tags** | **[]string** | filter by tag names | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdTasklistsHtml

> GETProjectsApiV3ProjectsProjectIdTasklistsHtml(ctx).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate tasklist report in HTML format



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
    taskDate := "taskDate_example" // string | filter by task date (optional) (default to "anytime")
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "manual")
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    commentsOrderMode := "commentsOrderMode_example" // string | order mode of the comments (optional) (default to "asc")
    tasklistId := int32(56) // int32 | filter by task list (optional)
    projectId := int32(56) // int32 | filter by project (optional)
    onlyCompletedTasks := true // bool | only completed tasks (optional) (default to false)
    isReportDownload := true // bool | generate a report document (optional)
    includeTime := true // bool | include time (optional)
    includeTags := true // bool | include tags (optional) (default to true)
    includePrivateItems := true // bool | include private items (optional) (default to true)
    includeLateTasks := true // bool | include late tasks (optional)
    includeEstimatedTime := true // bool | include estimated time (optional) (default to true)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedTasks := true // bool | include completed tasks (optional) (default to true)
    includeComments := true // bool | include comments (optional) (default to false)
    includeCommentAvatars := true // bool | include comment avatars (optional)
    includeColumns := true // bool | include columns (optional) (default to true)
    includeAnytimeTasks := true // bool | include anytime tasks (optional)
    includeAllComments := true // bool | include all comments (optional) (default to false)
    ignoreStartDates := true // bool | ignore task start dates (optional)
    groupByTasklists := true // bool | group by task lists (optional)
    tags := []string{"Inner_example"} // []string | filter by tag names (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsHtml(context.Background()).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsHtml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTasklistsHtmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDate** | **string** | filter by task date | [default to &quot;anytime&quot;]
 **startDate** | **time.Time** | filter by start datetime | 
 **reportFormat** | **string** | define the format of the report | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;manual&quot;]
 **endDate** | **time.Time** | filter by end datetime | 
 **commentsOrderMode** | **string** | order mode of the comments | [default to &quot;asc&quot;]
 **tasklistId** | **int32** | filter by task list | 
 **projectId** | **int32** | filter by project | 
 **onlyCompletedTasks** | **bool** | only completed tasks | [default to false]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTime** | **bool** | include time | 
 **includeTags** | **bool** | include tags | [default to true]
 **includePrivateItems** | **bool** | include private items | [default to true]
 **includeLateTasks** | **bool** | include late tasks | 
 **includeEstimatedTime** | **bool** | include estimated time | [default to true]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedTasks** | **bool** | include completed tasks | [default to true]
 **includeComments** | **bool** | include comments | [default to false]
 **includeCommentAvatars** | **bool** | include comment avatars | 
 **includeColumns** | **bool** | include columns | [default to true]
 **includeAnytimeTasks** | **bool** | include anytime tasks | 
 **includeAllComments** | **bool** | include all comments | [default to false]
 **ignoreStartDates** | **bool** | ignore task start dates | 
 **groupByTasklists** | **bool** | group by task lists | 
 **tags** | **[]string** | filter by tag names | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdTasklistsPdf

> GETProjectsApiV3ProjectsProjectIdTasklistsPdf(ctx).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate tasklist report in PDF format



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
    taskDate := "taskDate_example" // string | filter by task date (optional) (default to "anytime")
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "manual")
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    commentsOrderMode := "commentsOrderMode_example" // string | order mode of the comments (optional) (default to "asc")
    tasklistId := int32(56) // int32 | filter by task list (optional)
    projectId := int32(56) // int32 | filter by project (optional)
    onlyCompletedTasks := true // bool | only completed tasks (optional) (default to false)
    isReportDownload := true // bool | generate a report document (optional)
    includeTime := true // bool | include time (optional)
    includeTags := true // bool | include tags (optional) (default to true)
    includePrivateItems := true // bool | include private items (optional) (default to true)
    includeLateTasks := true // bool | include late tasks (optional)
    includeEstimatedTime := true // bool | include estimated time (optional) (default to true)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedTasks := true // bool | include completed tasks (optional) (default to true)
    includeComments := true // bool | include comments (optional) (default to false)
    includeCommentAvatars := true // bool | include comment avatars (optional)
    includeColumns := true // bool | include columns (optional) (default to true)
    includeAnytimeTasks := true // bool | include anytime tasks (optional)
    includeAllComments := true // bool | include all comments (optional) (default to false)
    ignoreStartDates := true // bool | ignore task start dates (optional)
    groupByTasklists := true // bool | group by task lists (optional)
    tags := []string{"Inner_example"} // []string | filter by tag names (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsPdf(context.Background()).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTasklistsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDate** | **string** | filter by task date | [default to &quot;anytime&quot;]
 **startDate** | **time.Time** | filter by start datetime | 
 **reportFormat** | **string** | define the format of the report | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;manual&quot;]
 **endDate** | **time.Time** | filter by end datetime | 
 **commentsOrderMode** | **string** | order mode of the comments | [default to &quot;asc&quot;]
 **tasklistId** | **int32** | filter by task list | 
 **projectId** | **int32** | filter by project | 
 **onlyCompletedTasks** | **bool** | only completed tasks | [default to false]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTime** | **bool** | include time | 
 **includeTags** | **bool** | include tags | [default to true]
 **includePrivateItems** | **bool** | include private items | [default to true]
 **includeLateTasks** | **bool** | include late tasks | 
 **includeEstimatedTime** | **bool** | include estimated time | [default to true]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedTasks** | **bool** | include completed tasks | [default to true]
 **includeComments** | **bool** | include comments | [default to false]
 **includeCommentAvatars** | **bool** | include comment avatars | 
 **includeColumns** | **bool** | include columns | [default to true]
 **includeAnytimeTasks** | **bool** | include anytime tasks | 
 **includeAllComments** | **bool** | include all comments | [default to false]
 **ignoreStartDates** | **bool** | ignore task start dates | 
 **groupByTasklists** | **bool** | group by task lists | 
 **tags** | **[]string** | filter by tag names | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdTasklistsXlsx

> GETProjectsApiV3ProjectsProjectIdTasklistsXlsx(ctx).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Generate tasklist report in XLSX format



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
    taskDate := "taskDate_example" // string | filter by task date (optional) (default to "anytime")
    startDate := time.Now() // time.Time | filter by start datetime (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | order by (optional) (default to "manual")
    endDate := time.Now() // time.Time | filter by end datetime (optional)
    commentsOrderMode := "commentsOrderMode_example" // string | order mode of the comments (optional) (default to "asc")
    tasklistId := int32(56) // int32 | filter by task list (optional)
    projectId := int32(56) // int32 | filter by project (optional)
    onlyCompletedTasks := true // bool | only completed tasks (optional) (default to false)
    isReportDownload := true // bool | generate a report document (optional)
    includeTime := true // bool | include time (optional)
    includeTags := true // bool | include tags (optional) (default to true)
    includePrivateItems := true // bool | include private items (optional) (default to true)
    includeLateTasks := true // bool | include late tasks (optional)
    includeEstimatedTime := true // bool | include estimated time (optional) (default to true)
    includeCustomFields := true // bool | include custom fields (optional) (default to false)
    includeCompletedTasks := true // bool | include completed tasks (optional) (default to true)
    includeComments := true // bool | include comments (optional) (default to false)
    includeCommentAvatars := true // bool | include comment avatars (optional)
    includeColumns := true // bool | include columns (optional) (default to true)
    includeAnytimeTasks := true // bool | include anytime tasks (optional)
    includeAllComments := true // bool | include all comments (optional) (default to false)
    ignoreStartDates := true // bool | ignore task start dates (optional)
    groupByTasklists := true // bool | group by task lists (optional)
    tags := []string{"Inner_example"} // []string | filter by tag names (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsXlsx(context.Background()).TaskDate(taskDate).StartDate(startDate).ReportFormat(reportFormat).OrderMode(orderMode).OrderBy(orderBy).EndDate(endDate).CommentsOrderMode(commentsOrderMode).TasklistId(tasklistId).ProjectId(projectId).OnlyCompletedTasks(onlyCompletedTasks).IsReportDownload(isReportDownload).IncludeTime(includeTime).IncludeTags(includeTags).IncludePrivateItems(includePrivateItems).IncludeLateTasks(includeLateTasks).IncludeEstimatedTime(includeEstimatedTime).IncludeCustomFields(includeCustomFields).IncludeCompletedTasks(includeCompletedTasks).IncludeComments(includeComments).IncludeCommentAvatars(includeCommentAvatars).IncludeColumns(includeColumns).IncludeAnytimeTasks(includeAnytimeTasks).IncludeAllComments(includeAllComments).IgnoreStartDates(ignoreStartDates).GroupByTasklists(groupByTasklists).Tags(tags).TagIds(tagIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskListsApi.GETProjectsApiV3ProjectsProjectIdTasklistsXlsx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTasklistsXlsxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDate** | **string** | filter by task date | [default to &quot;anytime&quot;]
 **startDate** | **time.Time** | filter by start datetime | 
 **reportFormat** | **string** | define the format of the report | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | order by | [default to &quot;manual&quot;]
 **endDate** | **time.Time** | filter by end datetime | 
 **commentsOrderMode** | **string** | order mode of the comments | [default to &quot;asc&quot;]
 **tasklistId** | **int32** | filter by task list | 
 **projectId** | **int32** | filter by project | 
 **onlyCompletedTasks** | **bool** | only completed tasks | [default to false]
 **isReportDownload** | **bool** | generate a report document | 
 **includeTime** | **bool** | include time | 
 **includeTags** | **bool** | include tags | [default to true]
 **includePrivateItems** | **bool** | include private items | [default to true]
 **includeLateTasks** | **bool** | include late tasks | 
 **includeEstimatedTime** | **bool** | include estimated time | [default to true]
 **includeCustomFields** | **bool** | include custom fields | [default to false]
 **includeCompletedTasks** | **bool** | include completed tasks | [default to true]
 **includeComments** | **bool** | include comments | [default to false]
 **includeCommentAvatars** | **bool** | include comment avatars | 
 **includeColumns** | **bool** | include columns | [default to true]
 **includeAnytimeTasks** | **bool** | include anytime tasks | 
 **includeAllComments** | **bool** | include all comments | [default to false]
 **ignoreStartDates** | **bool** | ignore task start dates | 
 **groupByTasklists** | **bool** | group by task lists | 
 **tags** | **[]string** | filter by tag names | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

