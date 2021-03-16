# \TimeTrackingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEProjectsApiV3MeTimersIdJson**](TimeTrackingApi.md#DELETEProjectsApiV3MeTimersIdJson) | **Delete** /projects/api/v3/me/timers/:id.json | Delete a timer by ID
[**GETProjectsApiV3AllocationsAllocationIdTimeJson**](TimeTrackingApi.md#GETProjectsApiV3AllocationsAllocationIdTimeJson) | **Get** /projects/api/v3/allocations/:allocationId/time.json | Get time entries for a specific allocation
[**GETProjectsApiV3MeTimersJson**](TimeTrackingApi.md#GETProjectsApiV3MeTimersJson) | **Get** /projects/api/v3/me/timers.json | Get all your running timers
[**GETProjectsApiV3ProjectsProjectIdTimeJson**](TimeTrackingApi.md#GETProjectsApiV3ProjectsProjectIdTimeJson) | **Get** /projects/api/v3/projects/:projectId/time.json | Get time entries for a specific project
[**GETProjectsApiV3ProjectsProjectIdTimeTotalJson**](TimeTrackingApi.md#GETProjectsApiV3ProjectsProjectIdTimeTotalJson) | **Get** /projects/api/v3/projects/:projectId/time/total.json | Get timelog totals.
[**GETProjectsApiV3TimeJson**](TimeTrackingApi.md#GETProjectsApiV3TimeJson) | **Get** /projects/api/v3/time.json | Get all time entries
[**GETProjectsApiV3TimeTotalJson**](TimeTrackingApi.md#GETProjectsApiV3TimeTotalJson) | **Get** /projects/api/v3/time/total.json | Get timelog totals.
[**GETProjectsApiV3TimersIdJson**](TimeTrackingApi.md#GETProjectsApiV3TimersIdJson) | **Get** /projects/api/v3/timers/:id.json | Get a specific timer
[**GETProjectsApiV3TimersJson**](TimeTrackingApi.md#GETProjectsApiV3TimersJson) | **Get** /projects/api/v3/timers.json | Get all running timers
[**POSTProjectsApiV3MeTimersJson**](TimeTrackingApi.md#POSTProjectsApiV3MeTimersJson) | **Post** /projects/api/v3/me/timers.json | Create a new timer
[**PUTProjectsApiV3MeTimersIdCompleteJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimersIdCompleteJson) | **Put** /projects/api/v3/me/timers/:id/complete.json | Complete a timer by ID
[**PUTProjectsApiV3MeTimersIdJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimersIdJson) | **Put** /projects/api/v3/me/timers/:id.json | Edits a timer
[**PUTProjectsApiV3MeTimersIdPauseJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimersIdPauseJson) | **Put** /projects/api/v3/me/timers/:id/pause.json | Pause a timer by ID
[**PUTProjectsApiV3MeTimersIdResumeJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimersIdResumeJson) | **Put** /projects/api/v3/me/timers/:id/resume.json | Resume a timer by ID
[**PUTProjectsApiV3MeTimersIdUndeleteJson**](TimeTrackingApi.md#PUTProjectsApiV3MeTimersIdUndeleteJson) | **Put** /projects/api/v3/me/timers/:id/undelete.json | Undelete a timer by ID



## DELETEProjectsApiV3MeTimersIdJson

> DELETEProjectsApiV3MeTimersIdJson(ctx).TimerDeleteRequest(timerDeleteRequest).Execute()

Delete a timer by ID



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
    timerDeleteRequest := *openapiclient.NewTimerDeleteRequest() // TimerDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.DELETEProjectsApiV3MeTimersIdJson(context.Background()).TimerDeleteRequest(timerDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.DELETEProjectsApiV3MeTimersIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDELETEProjectsApiV3MeTimersIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timerDeleteRequest** | [**TimerDeleteRequest**](TimerDeleteRequest.md) |  | 

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


## GETProjectsApiV3AllocationsAllocationIdTimeJson

> TimelogTimelogsResponse GETProjectsApiV3AllocationsAllocationIdTimeJson(ctx).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get time entries for a specific allocation



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
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3AllocationsAllocationIdTimeJson(context.Background()).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3AllocationsAllocationIdTimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3AllocationsAllocationIdTimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3AllocationsAllocationIdTimeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3AllocationsAllocationIdTimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3MeTimersJson

> TimerTimersResponse GETProjectsApiV3MeTimersJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get all your running timers



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
    userId := int32(56) // int32 | filter by user id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    runningTimersOnly := true // bool | enforce running timers only default: false (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3MeTimersJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3MeTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3MeTimersJson`: TimerTimersResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3MeTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3MeTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **userId** | **int32** | filter by user id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **runningTimersOnly** | **bool** | enforce running timers only default: false | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerTimersResponse**](TimerTimersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdTimeJson

> TimelogTimelogsResponse GETProjectsApiV3ProjectsProjectIdTimeJson(ctx).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get time entries for a specific project



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
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeJson(context.Background()).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsProjectIdTimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3ProjectsProjectIdTimeTotalJson

> TimelogTotalsResponse GETProjectsApiV3ProjectsProjectIdTimeTotalJson(ctx).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()

Get timelog totals.

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
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    deskTicketId := int32(56) // int32 | filter by desk ticket id (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeTotalJson(context.Background()).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3ProjectsProjectIdTimeTotalJson`: TimelogTotalsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3ProjectsProjectIdTimeTotalJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3ProjectsProjectIdTimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **userId** | **int32** | filter by user id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **deskTicketId** | **int32** | filter by desk ticket id | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by project company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **companyIds** | **[]int32** | filter by company ids | 

### Return type

[**TimelogTotalsResponse**](TimelogTotalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimeJson

> TimelogTimelogsResponse GETProjectsApiV3TimeJson(ctx).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()

Get all time entries



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
    updatedAfter := time.Now() // time.Time | filter by updated after date (optional)
    startDate := time.Now() // string | filter by a starting date (optional)
    selectedColumns := "selectedColumns_example" // string | customise the report by selecting columns (optional)
    reportFormat := "reportFormat_example" // string | define the format of the report (optional)
    projectStatuses := "projectStatuses_example" // string | filter by project statuses (optional)
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    orderMode := "orderMode_example" // string | order mode (optional) (default to "asc")
    orderBy := "orderBy_example" // string | sort order (optional) (default to "date")
    invoicedType := "invoicedType_example" // string | filter by invoiced type (optional) (default to "all")
    endDate := time.Now() // string | filter by an ending date (optional)
    billableType := "billableType_example" // string | filter by billable type (optional) (default to "all")
    ticketId := int32(56) // int32 | filter by ticket id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    pageSize := int32(56) // int32 | number of items in a page (optional) (default to 50)
    page := int32(56) // int32 | page number (optional) (default to 1)
    allocationId := int32(56) // int32 | filter by allocation id (optional)
    showDeleted := true // bool | include deleted items (optional) (default to false)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    isReportDownload := true // bool | generate a report document (optional)
    includeArchivedProjects := true // bool | include archived projects (optional) (default to false)
    taskTagIds := []int32{int32(123)} // []int32 | filter by task tag ids (optional)
    tagIds := []int32{int32(123)} // []int32 | filter by tag ids (optional)
    projectsFromCompanyId := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimelogs := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsTags := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)
    assignedToUserIds := []int32{int32(123)} // []int32 | filter by assigned user ids (optional)
    assignedToTeamIds := []int32{int32(123)} // []int32 | filter by assigned team ids (optional)
    assignedToCompanyIds := []int32{int32(123)} // []int32 | filter by assigned company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimeJson(context.Background()).UpdatedAfter(updatedAfter).StartDate(startDate).SelectedColumns(selectedColumns).ReportFormat(reportFormat).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).OrderMode(orderMode).OrderBy(orderBy).InvoicedType(invoicedType).EndDate(endDate).BillableType(billableType).TicketId(ticketId).TasklistId(tasklistId).ProjectId(projectId).ProjectHealths(projectHealths).PageSize(pageSize).Page(page).AllocationId(allocationId).ShowDeleted(showDeleted).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).IsReportDownload(isReportDownload).IncludeArchivedProjects(includeArchivedProjects).TaskTagIds(taskTagIds).TagIds(tagIds).ProjectsFromCompanyId(projectsFromCompanyId).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCategoryIds(projectCategoryIds).Include(include).FieldsUsers(fieldsUsers).FieldsTimelogs(fieldsTimelogs).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsTags(fieldsTags).FieldsProjects(fieldsProjects).CompanyIds(companyIds).AssignedToUserIds(assignedToUserIds).AssignedToTeamIds(assignedToTeamIds).AssignedToCompanyIds(assignedToCompanyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimeJson`: TimelogTimelogsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimeJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **startDate** | **string** | filter by a starting date | 
 **selectedColumns** | **string** | customise the report by selecting columns | 
 **reportFormat** | **string** | define the format of the report | 
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **orderMode** | **string** | order mode | [default to &quot;asc&quot;]
 **orderBy** | **string** | sort order | [default to &quot;date&quot;]
 **invoicedType** | **string** | filter by invoiced type | [default to &quot;all&quot;]
 **endDate** | **string** | filter by an ending date | 
 **billableType** | **string** | filter by billable type | [default to &quot;all&quot;]
 **ticketId** | **int32** | filter by ticket id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **pageSize** | **int32** | number of items in a page | [default to 50]
 **page** | **int32** | page number | [default to 1]
 **allocationId** | **int32** | filter by allocation id | 
 **showDeleted** | **bool** | include deleted items | [default to false]
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **isReportDownload** | **bool** | generate a report document | 
 **includeArchivedProjects** | **bool** | include archived projects | [default to false]
 **taskTagIds** | **[]int32** | filter by task tag ids | 
 **tagIds** | **[]int32** | filter by tag ids | 
 **projectsFromCompanyId** | **[]int32** | filter by project company ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimelogs** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsTags** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 
 **companyIds** | **[]int32** | filter by company ids | 
 **assignedToUserIds** | **[]int32** | filter by assigned user ids | 
 **assignedToTeamIds** | **[]int32** | filter by assigned team ids | 
 **assignedToCompanyIds** | **[]int32** | filter by assigned company ids | 

### Return type

[**TimelogTimelogsResponse**](TimelogTimelogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimeTotalJson

> TimelogTotalsResponse GETProjectsApiV3TimeTotalJson(ctx).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()

Get timelog totals.

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
    projectStatus := "projectStatus_example" // string | filter by project status (optional)
    userId := int32(56) // int32 | filter by user id (optional)
    tasklistId := int32(56) // int32 | filter by tasklist id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    projectHealths := int32(56) // int32 | filter by project healths  0: not set 1: bad 2: ok 3: good (optional)
    deskTicketId := int32(56) // int32 | filter by desk ticket id (optional)
    onlyStarredProjects := true // bool | filter by starred projects only (optional)
    matchAllTags := true // bool | match all tags (optional)
    userIds := []int32{int32(123)} // []int32 | filter by user ids (optional)
    projectTagIds := []int32{int32(123)} // []int32 | filter by project tag ids (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 | filter by project owner ids (optional)
    projectIds := []int32{int32(123)} // []int32 | filter by project ids (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 | filter by project company ids (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 | filter by project category ids (optional)
    companyIds := []int32{int32(123)} // []int32 | filter by company ids (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimeTotalJson(context.Background()).ProjectStatuses(projectStatuses).ProjectStatus(projectStatus).UserId(userId).TasklistId(tasklistId).TaskId(taskId).ProjectId(projectId).ProjectHealths(projectHealths).DeskTicketId(deskTicketId).OnlyStarredProjects(onlyStarredProjects).MatchAllTags(matchAllTags).UserIds(userIds).ProjectTagIds(projectTagIds).ProjectOwnerIds(projectOwnerIds).ProjectIds(projectIds).ProjectCompanyIds(projectCompanyIds).ProjectCategoryIds(projectCategoryIds).CompanyIds(companyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimeTotalJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimeTotalJson`: TimelogTotalsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimeTotalJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimeTotalJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectStatuses** | **string** | filter by project statuses | 
 **projectStatus** | **string** | filter by project status | 
 **userId** | **int32** | filter by user id | 
 **tasklistId** | **int32** | filter by tasklist id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **projectHealths** | **int32** | filter by project healths  0: not set 1: bad 2: ok 3: good | 
 **deskTicketId** | **int32** | filter by desk ticket id | 
 **onlyStarredProjects** | **bool** | filter by starred projects only | 
 **matchAllTags** | **bool** | match all tags | 
 **userIds** | **[]int32** | filter by user ids | 
 **projectTagIds** | **[]int32** | filter by project tag ids | 
 **projectOwnerIds** | **[]int32** | filter by project owner ids | 
 **projectIds** | **[]int32** | filter by project ids | 
 **projectCompanyIds** | **[]int32** | filter by project company ids | 
 **projectCategoryIds** | **[]int32** | filter by project category ids | 
 **companyIds** | **[]int32** | filter by company ids | 

### Return type

[**TimelogTotalsResponse**](TimelogTotalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimersIdJson

> TimerResponse GETProjectsApiV3TimersIdJson(ctx).UserId(userId).ShowDeleted(showDeleted).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get a specific timer



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
    userId := int32(56) // int32 | filter by user id (optional)
    showDeleted := true // bool | filter by task id (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimersIdJson(context.Background()).UserId(userId).ShowDeleted(showDeleted).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimersIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimersIdJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimersIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimersIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** | filter by user id | 
 **showDeleted** | **bool** | filter by task id | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV3TimersJson

> TimerTimersResponse GETProjectsApiV3TimersJson(ctx).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()

Get all running timers



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
    userId := int32(56) // int32 | filter by user id (optional)
    taskId := int32(56) // int32 | filter by task id (optional)
    projectId := int32(56) // int32 | filter by project id (optional)
    pageSize := int32(56) // int32 | number of items in a page default: 50 (optional)
    page := int32(56) // int32 | page number default: 1 (optional)
    showDeleted := true // bool | include deleted items default: false (optional)
    runningTimersOnly := true // bool | enforce running timers only default: false (optional)
    include := []string{"Inner_example"} // []string | include (optional)
    fieldsUsers := []string{"Inner_example"} // []string |  (optional)
    fieldsTimers := []string{"Inner_example"} // []string |  (optional)
    fieldsTasks := []string{"Inner_example"} // []string |  (optional)
    fieldsTasklists := []string{"Inner_example"} // []string |  (optional)
    fieldsProjects := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.GETProjectsApiV3TimersJson(context.Background()).UpdatedAfterDate(updatedAfterDate).UpdatedAfter(updatedAfter).UserId(userId).TaskId(taskId).ProjectId(projectId).PageSize(pageSize).Page(page).ShowDeleted(showDeleted).RunningTimersOnly(runningTimersOnly).Include(include).FieldsUsers(fieldsUsers).FieldsTimers(fieldsTimers).FieldsTasks(fieldsTasks).FieldsTasklists(fieldsTasklists).FieldsProjects(fieldsProjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.GETProjectsApiV3TimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV3TimersJson`: TimerTimersResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.GETProjectsApiV3TimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV3TimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **time.Time** | filter by updated after date (deprecated, use updatedAfter) | 
 **updatedAfter** | **time.Time** | filter by updated after date | 
 **userId** | **int32** | filter by user id | 
 **taskId** | **int32** | filter by task id | 
 **projectId** | **int32** | filter by project id | 
 **pageSize** | **int32** | number of items in a page default: 50 | 
 **page** | **int32** | page number default: 1 | 
 **showDeleted** | **bool** | include deleted items default: false | 
 **runningTimersOnly** | **bool** | enforce running timers only default: false | 
 **include** | **[]string** | include | 
 **fieldsUsers** | **[]string** |  | 
 **fieldsTimers** | **[]string** |  | 
 **fieldsTasks** | **[]string** |  | 
 **fieldsTasklists** | **[]string** |  | 
 **fieldsProjects** | **[]string** |  | 

### Return type

[**TimerTimersResponse**](TimerTimersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTProjectsApiV3MeTimersJson

> TimerResponse POSTProjectsApiV3MeTimersJson(ctx).TimerRequest(timerRequest).Execute()

Create a new timer



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
    timerRequest := *openapiclient.NewTimerRequest() // TimerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.POSTProjectsApiV3MeTimersJson(context.Background()).TimerRequest(timerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.POSTProjectsApiV3MeTimersJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTProjectsApiV3MeTimersJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.POSTProjectsApiV3MeTimersJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTProjectsApiV3MeTimersJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timerRequest** | [**TimerRequest**](TimerRequest.md) |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimersIdCompleteJson

> TimerResponse PUTProjectsApiV3MeTimersIdCompleteJson(ctx).Execute()

Complete a timer by ID



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
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimersIdCompleteJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimersIdCompleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimersIdCompleteJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimersIdCompleteJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimersIdCompleteJsonRequest struct via the builder pattern


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimersIdJson

> TimerResponse PUTProjectsApiV3MeTimersIdJson(ctx).TimerRequest(timerRequest).Execute()

Edits a timer



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
    timerRequest := *openapiclient.NewTimerRequest() // TimerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimersIdJson(context.Background()).TimerRequest(timerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimersIdJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimersIdJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimersIdJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimersIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timerRequest** | [**TimerRequest**](TimerRequest.md) |  | 

### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimersIdPauseJson

> TimerResponse PUTProjectsApiV3MeTimersIdPauseJson(ctx).Execute()

Pause a timer by ID



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
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimersIdPauseJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimersIdPauseJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimersIdPauseJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimersIdPauseJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimersIdPauseJsonRequest struct via the builder pattern


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimersIdResumeJson

> TimerResponse PUTProjectsApiV3MeTimersIdResumeJson(ctx).Execute()

Resume a timer by ID



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
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimersIdResumeJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimersIdResumeJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PUTProjectsApiV3MeTimersIdResumeJson`: TimerResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeTrackingApi.PUTProjectsApiV3MeTimersIdResumeJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimersIdResumeJsonRequest struct via the builder pattern


### Return type

[**TimerResponse**](TimerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PUTProjectsApiV3MeTimersIdUndeleteJson

> PUTProjectsApiV3MeTimersIdUndeleteJson(ctx).Execute()

Undelete a timer by ID



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
    resp, r, err := api_client.TimeTrackingApi.PUTProjectsApiV3MeTimersIdUndeleteJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingApi.PUTProjectsApiV3MeTimersIdUndeleteJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPUTProjectsApiV3MeTimersIdUndeleteJsonRequest struct via the builder pattern


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

