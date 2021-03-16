# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETProjectsApiV2PeopleUserIdAvatar**](DefaultApi.md#GETProjectsApiV2PeopleUserIdAvatar) | **Get** /projects/api/v2/people/:userId/avatar | Get user avatar
[**GETProjectsApiV2ProjectsProjectIdPeopleUserIds**](DefaultApi.md#GETProjectsApiV2ProjectsProjectIdPeopleUserIds) | **Get** /projects/api/v2/projects/:projectId/people/:userIds | Get people project
[**GETProjectsApiV2ProjectsProjectIdTasksGantt**](DefaultApi.md#GETProjectsApiV2ProjectsProjectIdTasksGantt) | **Get** /projects/api/v2/projects/:projectId/tasks/gantt | Will return the total for all sub-tasks



## GETProjectsApiV2PeopleUserIdAvatar

> GETProjectsApiV2PeopleUserIdAvatar(ctx).Execute()

Get user avatar

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
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2PeopleUserIdAvatar(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2PeopleUserIdAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2PeopleUserIdAvatarRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETProjectsApiV2ProjectsProjectIdPeopleUserIds

> PeopleResponse GETProjectsApiV2ProjectsProjectIdPeopleUserIds(ctx).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()

Get people project

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
    updatedAfterDateTime := "updatedAfterDateTime_example" // string |  (optional)
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    searchTerm := "searchTerm_example" // string |  (optional)
    projectType := "projectType_example" // string |  (optional)
    inviteStatus := "inviteStatus_example" // string |  (optional)
    firstLetter := "firstLetter_example" // string |  (optional)
    emailAddress := "emailAddress_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    projectId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    excludeProjectId := int32(56) // int32 |  (optional)
    showDeleted := true // bool |  (optional)
    returnTeams := true // bool |  (optional)
    returnProjectIds := true // bool |  (optional)
    returnLetters := true // bool |  (optional)
    ownerCompanyFirst := true // bool |  (optional)
    onlyids := true // bool |  (optional)
    includeTags := true // bool |  (optional)
    includeObservers := true // bool |  (optional)
    includeCompanyDetails := true // bool |  (optional)
    includeClockIn := true // bool |  (optional)
    groupByCompany := true // bool |  (optional)
    getProjectRoles := true // bool |  (optional)
    getCounts := true // bool |  (optional)
    fullprofile := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    userIds := []int32{int32(123)} // []int32 |  (optional)
    companyId := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2ProjectsProjectIdPeopleUserIds(context.Background()).UpdatedAfterDateTime(updatedAfterDateTime).UpdatedAfterDate(updatedAfterDate).Type_(type_).SortOrder(sortOrder).Sort(sort).SearchTerm(searchTerm).ProjectType(projectType).InviteStatus(inviteStatus).FirstLetter(firstLetter).EmailAddress(emailAddress).DataSet(dataSet).ProjectId(projectId).PageSize(pageSize).Page(page).ExcludeProjectId(excludeProjectId).ShowDeleted(showDeleted).ReturnTeams(returnTeams).ReturnProjectIds(returnProjectIds).ReturnLetters(returnLetters).OwnerCompanyFirst(ownerCompanyFirst).Onlyids(onlyids).IncludeTags(includeTags).IncludeObservers(includeObservers).IncludeCompanyDetails(includeCompanyDetails).IncludeClockIn(includeClockIn).GroupByCompany(groupByCompany).GetProjectRoles(getProjectRoles).GetCounts(getCounts).Fullprofile(fullprofile).CountOnly(countOnly).UserIds(userIds).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2ProjectsProjectIdPeopleUserIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETProjectsApiV2ProjectsProjectIdPeopleUserIds`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GETProjectsApiV2ProjectsProjectIdPeopleUserIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsProjectIdPeopleUserIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDateTime** | **string** |  | 
 **updatedAfterDate** | **string** |  | 
 **type_** | **string** |  | 
 **sortOrder** | **string** |  | 
 **sort** | **string** |  | 
 **searchTerm** | **string** |  | 
 **projectType** | **string** |  | 
 **inviteStatus** | **string** |  | 
 **firstLetter** | **string** |  | 
 **emailAddress** | **string** |  | 
 **dataSet** | **string** |  | 
 **projectId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **excludeProjectId** | **int32** |  | 
 **showDeleted** | **bool** |  | 
 **returnTeams** | **bool** |  | 
 **returnProjectIds** | **bool** |  | 
 **returnLetters** | **bool** |  | 
 **ownerCompanyFirst** | **bool** |  | 
 **onlyids** | **bool** |  | 
 **includeTags** | **bool** |  | 
 **includeObservers** | **bool** |  | 
 **includeCompanyDetails** | **bool** |  | 
 **includeClockIn** | **bool** |  | 
 **groupByCompany** | **bool** |  | 
 **getProjectRoles** | **bool** |  | 
 **getCounts** | **bool** |  | 
 **fullprofile** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **userIds** | **[]int32** |  | 
 **companyId** | **[]int32** |  | 

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


## GETProjectsApiV2ProjectsProjectIdTasksGantt

> GETProjectsApiV2ProjectsProjectIdTasksGantt(ctx).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).TagIds2(tagIds2).Startdate(startdate).StartDate(startDate).Sort(sort).Sort2(sort2).ResponsiblePartyIds(responsiblePartyIds).ResponsiblePartyId(responsiblePartyId).ProjectStatus(projectStatus).ProjectStatus2(projectStatus2).Include(include).FollowerIds(followerIds).FollowedByUserIds(followedByUserIds).Filter(filter).Filter2(filter2).Enddate(enddate).EndDate(endDate).DataSet(dataSet).DataSet2(dataSet2).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedBeforeDate2(completedBeforeDate2).CompletedAfterDate(completedAfterDate).CompletedAfterDate2(completedAfterDate2).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ProjectCategoryId2(projectCategoryId2).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).IncludeTaskId2(includeTaskId2).CompanyId(companyId).CompanyId2(companyId2).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).UseAllProjects2(useAllProjects2).StarredProjectsOnly(starredProjectsOnly).StarredProjectsOnly2(starredProjectsOnly2).ShowDeleted(showDeleted).ShowDeleted2(showDeleted2).ShowCompletedLists(showCompletedLists).ShowCompletedLists2(showCompletedLists2).OnlyStarredProjects(onlyStarredProjects).OnlyStarredProjects2(onlyStarredProjects2).OnlyArchivedProjects(onlyArchivedProjects).OnlyArchivedProjects2(onlyArchivedProjects2).NestSubTasks(nestSubTasks).NestSubTasks2(nestSubTasks2).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTags2(matchAllProjectTags2).IncludeUntaggedTasks(includeUntaggedTasks).IncludeUntaggedTasks2(includeUntaggedTasks2).IncludeToday(includeToday).IncludeToday2(includeToday2).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeTasksFromDeletedLists2(includeTasksFromDeletedLists2).IncludeReminders(includeReminders).IncludeReminders2(includeReminders2).IncludeLoggedTime(includeLoggedTime).IncludeLoggedTime2(includeLoggedTime2).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedTasks2(includeCompletedTasks2).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedSubtasks2(includeCompletedSubtasks2).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeCompletedPredecessors2(includeCompletedPredecessors2).IncludeArchivedProjects(includeArchivedProjects).IncludeArchivedProjects2(includeArchivedProjects2).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetSubTasks2(getSubTasks2).GetFiles(getFiles).GetFiles2(getFiles2).CountOnly(countOnly).CountOnly2(countOnly2).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectTagIds2(projectTagIds2).ProjectOwnerIds(projectOwnerIds).ProjectOwnerIds2(projectOwnerIds2).ProjectIds(projectIds).ProjectIds2(projectIds2).ProjectId(projectId).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealths2(projectHealths2).ProjectCompanyIds(projectCompanyIds).ProjectCompanyIds2(projectCompanyIds2).ProjectCategoryIds(projectCategoryIds).ProjectCategoryIds2(projectCategoryIds2).Execute()

Will return the total for all sub-tasks

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
    updatedAfterDate := "updatedAfterDate_example" // string |  (optional)
    today := "today_example" // string |  (optional)
    tagIds := "tagIds_example" // string |  (optional)
    tagIds2 := "tagIds_example" // string |  (optional)
    startdate := "startdate_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    sort2 := "sort_example" // string |  (optional)
    responsiblePartyIds := "responsiblePartyIds_example" // string |  (optional)
    responsiblePartyId := "responsiblePartyId_example" // string |  (optional)
    projectStatus := "projectStatus_example" // string |  (optional)
    projectStatus2 := "projectStatus_example" // string |  (optional)
    include := "include_example" // string |  (optional)
    followerIds := "followerIds_example" // string |  (optional)
    followedByUserIds := "followedByUserIds_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    filter2 := "filter_example" // string |  (optional)
    enddate := "enddate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    dataSet := "dataSet_example" // string |  (optional)
    dataSet2 := "dataSet_example" // string |  (optional)
    creatorIds := "creatorIds_example" // string |  (optional)
    createdBeforeDate := "createdBeforeDate_example" // string |  (optional)
    createdAfterDate := "createdAfterDate_example" // string |  (optional)
    completedBeforeDate := "completedBeforeDate_example" // string |  (optional)
    completedBeforeDate2 := "completedBeforeDate_example" // string |  (optional)
    completedAfterDate := "completedAfterDate_example" // string |  (optional)
    completedAfterDate2 := "completedAfterDate_example" // string |  (optional)
    callback := "callback_example" // string |  (optional)
    tasklistId := int32(56) // int32 |  (optional)
    taskId := int32(56) // int32 |  (optional)
    projectCategoryId := int32(56) // int32 |  (optional)
    projectCategoryId2 := int32(56) // int32 |  (optional)
    parentTaskId := int32(56) // int32 |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)
    includeTaskId := int32(56) // int32 |  (optional)
    includeTaskId2 := int32(56) // int32 |  (optional)
    companyId := int32(56) // int32 |  (optional)
    companyId2 := int32(56) // int32 |  (optional)
    useStartDatesForTodaysTasks := true // bool |  (optional)
    useAllProjects := true // bool |  (optional)
    useAllProjects2 := true // bool |  (optional)
    starredProjectsOnly := true // bool |  (optional)
    starredProjectsOnly2 := true // bool |  (optional)
    showDeleted := true // bool |  (optional)
    showDeleted2 := true // bool |  (optional)
    showCompletedLists := true // bool |  (optional)
    showCompletedLists2 := true // bool |  (optional)
    onlyStarredProjects := true // bool |  (optional)
    onlyStarredProjects2 := true // bool |  (optional)
    onlyArchivedProjects := true // bool |  (optional)
    onlyArchivedProjects2 := true // bool |  (optional)
    nestSubTasks := true // bool |  (optional)
    nestSubTasks2 := true // bool |  (optional)
    matchAllProjectTags := true // bool |  (optional)
    matchAllProjectTags2 := true // bool |  (optional)
    includeUntaggedTasks := true // bool |  (optional)
    includeUntaggedTasks2 := true // bool |  (optional)
    includeToday := true // bool |  (optional)
    includeToday2 := true // bool |  (optional)
    includeTasksWithoutDueDates := true // bool |  (optional)
    includeTasksFromDeletedLists := true // bool |  (optional)
    includeTasksFromDeletedLists2 := true // bool |  (optional)
    includeReminders := true // bool |  (optional)
    includeReminders2 := true // bool |  (optional)
    includeLoggedTime := true // bool |  (optional)
    includeLoggedTime2 := true // bool |  (optional)
    includeCompletedTasks := true // bool |  (optional)
    includeCompletedTasks2 := true // bool |  (optional)
    includeCompletedSubtasks := true // bool |  (optional)
    includeCompletedSubtasks2 := true // bool |  (optional)
    includeCompletedPredecessors := true // bool |  (optional)
    includeCompletedPredecessors2 := true // bool |  (optional)
    includeArchivedProjects := true // bool |  (optional)
    includeArchivedProjects2 := true // bool |  (optional)
    ignoreStartDates := true // bool |  (optional)
    getSubTasks := true // bool |  (optional)
    getSubTasks2 := true // bool |  (optional)
    getFiles := true // bool |  (optional)
    getFiles2 := true // bool |  (optional)
    countOnly := true // bool |  (optional)
    countOnly2 := true // bool |  (optional)
    allowTemplateTasks := true // bool |  (optional)
    projectTagIds := []int32{int32(123)} // []int32 |  (optional)
    projectTagIds2 := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds := []int32{int32(123)} // []int32 |  (optional)
    projectOwnerIds2 := []int32{int32(123)} // []int32 |  (optional)
    projectIds := []int32{int32(123)} // []int32 |  (optional)
    projectIds2 := []int32{int32(123)} // []int32 |  (optional)
    projectId := []int32{int32(123)} // []int32 | project *filters (optional)
    projectId2 := []int32{int32(123)} // []int32 | project *filters (optional)
    projectHealths := []int32{int32(123)} // []int32 |  (optional)
    projectHealths2 := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds := []int32{int32(123)} // []int32 |  (optional)
    projectCompanyIds2 := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds := []int32{int32(123)} // []int32 |  (optional)
    projectCategoryIds2 := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GETProjectsApiV2ProjectsProjectIdTasksGantt(context.Background()).UpdatedAfterDate(updatedAfterDate).Today(today).TagIds(tagIds).TagIds2(tagIds2).Startdate(startdate).StartDate(startDate).Sort(sort).Sort2(sort2).ResponsiblePartyIds(responsiblePartyIds).ResponsiblePartyId(responsiblePartyId).ProjectStatus(projectStatus).ProjectStatus2(projectStatus2).Include(include).FollowerIds(followerIds).FollowedByUserIds(followedByUserIds).Filter(filter).Filter2(filter2).Enddate(enddate).EndDate(endDate).DataSet(dataSet).DataSet2(dataSet2).CreatorIds(creatorIds).CreatedBeforeDate(createdBeforeDate).CreatedAfterDate(createdAfterDate).CompletedBeforeDate(completedBeforeDate).CompletedBeforeDate2(completedBeforeDate2).CompletedAfterDate(completedAfterDate).CompletedAfterDate2(completedAfterDate2).Callback(callback).TasklistId(tasklistId).TaskId(taskId).ProjectCategoryId(projectCategoryId).ProjectCategoryId2(projectCategoryId2).ParentTaskId(parentTaskId).PageSize(pageSize).Page(page).IncludeTaskId(includeTaskId).IncludeTaskId2(includeTaskId2).CompanyId(companyId).CompanyId2(companyId2).UseStartDatesForTodaysTasks(useStartDatesForTodaysTasks).UseAllProjects(useAllProjects).UseAllProjects2(useAllProjects2).StarredProjectsOnly(starredProjectsOnly).StarredProjectsOnly2(starredProjectsOnly2).ShowDeleted(showDeleted).ShowDeleted2(showDeleted2).ShowCompletedLists(showCompletedLists).ShowCompletedLists2(showCompletedLists2).OnlyStarredProjects(onlyStarredProjects).OnlyStarredProjects2(onlyStarredProjects2).OnlyArchivedProjects(onlyArchivedProjects).OnlyArchivedProjects2(onlyArchivedProjects2).NestSubTasks(nestSubTasks).NestSubTasks2(nestSubTasks2).MatchAllProjectTags(matchAllProjectTags).MatchAllProjectTags2(matchAllProjectTags2).IncludeUntaggedTasks(includeUntaggedTasks).IncludeUntaggedTasks2(includeUntaggedTasks2).IncludeToday(includeToday).IncludeToday2(includeToday2).IncludeTasksWithoutDueDates(includeTasksWithoutDueDates).IncludeTasksFromDeletedLists(includeTasksFromDeletedLists).IncludeTasksFromDeletedLists2(includeTasksFromDeletedLists2).IncludeReminders(includeReminders).IncludeReminders2(includeReminders2).IncludeLoggedTime(includeLoggedTime).IncludeLoggedTime2(includeLoggedTime2).IncludeCompletedTasks(includeCompletedTasks).IncludeCompletedTasks2(includeCompletedTasks2).IncludeCompletedSubtasks(includeCompletedSubtasks).IncludeCompletedSubtasks2(includeCompletedSubtasks2).IncludeCompletedPredecessors(includeCompletedPredecessors).IncludeCompletedPredecessors2(includeCompletedPredecessors2).IncludeArchivedProjects(includeArchivedProjects).IncludeArchivedProjects2(includeArchivedProjects2).IgnoreStartDates(ignoreStartDates).GetSubTasks(getSubTasks).GetSubTasks2(getSubTasks2).GetFiles(getFiles).GetFiles2(getFiles2).CountOnly(countOnly).CountOnly2(countOnly2).AllowTemplateTasks(allowTemplateTasks).ProjectTagIds(projectTagIds).ProjectTagIds2(projectTagIds2).ProjectOwnerIds(projectOwnerIds).ProjectOwnerIds2(projectOwnerIds2).ProjectIds(projectIds).ProjectIds2(projectIds2).ProjectId(projectId).ProjectId2(projectId2).ProjectHealths(projectHealths).ProjectHealths2(projectHealths2).ProjectCompanyIds(projectCompanyIds).ProjectCompanyIds2(projectCompanyIds2).ProjectCategoryIds(projectCategoryIds).ProjectCategoryIds2(projectCategoryIds2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GETProjectsApiV2ProjectsProjectIdTasksGantt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETProjectsApiV2ProjectsProjectIdTasksGanttRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAfterDate** | **string** |  | 
 **today** | **string** |  | 
 **tagIds** | **string** |  | 
 **tagIds2** | **string** |  | 
 **startdate** | **string** |  | 
 **startDate** | **string** |  | 
 **sort** | **string** |  | 
 **sort2** | **string** |  | 
 **responsiblePartyIds** | **string** |  | 
 **responsiblePartyId** | **string** |  | 
 **projectStatus** | **string** |  | 
 **projectStatus2** | **string** |  | 
 **include** | **string** |  | 
 **followerIds** | **string** |  | 
 **followedByUserIds** | **string** |  | 
 **filter** | **string** |  | 
 **filter2** | **string** |  | 
 **enddate** | **string** |  | 
 **endDate** | **string** |  | 
 **dataSet** | **string** |  | 
 **dataSet2** | **string** |  | 
 **creatorIds** | **string** |  | 
 **createdBeforeDate** | **string** |  | 
 **createdAfterDate** | **string** |  | 
 **completedBeforeDate** | **string** |  | 
 **completedBeforeDate2** | **string** |  | 
 **completedAfterDate** | **string** |  | 
 **completedAfterDate2** | **string** |  | 
 **callback** | **string** |  | 
 **tasklistId** | **int32** |  | 
 **taskId** | **int32** |  | 
 **projectCategoryId** | **int32** |  | 
 **projectCategoryId2** | **int32** |  | 
 **parentTaskId** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 
 **includeTaskId** | **int32** |  | 
 **includeTaskId2** | **int32** |  | 
 **companyId** | **int32** |  | 
 **companyId2** | **int32** |  | 
 **useStartDatesForTodaysTasks** | **bool** |  | 
 **useAllProjects** | **bool** |  | 
 **useAllProjects2** | **bool** |  | 
 **starredProjectsOnly** | **bool** |  | 
 **starredProjectsOnly2** | **bool** |  | 
 **showDeleted** | **bool** |  | 
 **showDeleted2** | **bool** |  | 
 **showCompletedLists** | **bool** |  | 
 **showCompletedLists2** | **bool** |  | 
 **onlyStarredProjects** | **bool** |  | 
 **onlyStarredProjects2** | **bool** |  | 
 **onlyArchivedProjects** | **bool** |  | 
 **onlyArchivedProjects2** | **bool** |  | 
 **nestSubTasks** | **bool** |  | 
 **nestSubTasks2** | **bool** |  | 
 **matchAllProjectTags** | **bool** |  | 
 **matchAllProjectTags2** | **bool** |  | 
 **includeUntaggedTasks** | **bool** |  | 
 **includeUntaggedTasks2** | **bool** |  | 
 **includeToday** | **bool** |  | 
 **includeToday2** | **bool** |  | 
 **includeTasksWithoutDueDates** | **bool** |  | 
 **includeTasksFromDeletedLists** | **bool** |  | 
 **includeTasksFromDeletedLists2** | **bool** |  | 
 **includeReminders** | **bool** |  | 
 **includeReminders2** | **bool** |  | 
 **includeLoggedTime** | **bool** |  | 
 **includeLoggedTime2** | **bool** |  | 
 **includeCompletedTasks** | **bool** |  | 
 **includeCompletedTasks2** | **bool** |  | 
 **includeCompletedSubtasks** | **bool** |  | 
 **includeCompletedSubtasks2** | **bool** |  | 
 **includeCompletedPredecessors** | **bool** |  | 
 **includeCompletedPredecessors2** | **bool** |  | 
 **includeArchivedProjects** | **bool** |  | 
 **includeArchivedProjects2** | **bool** |  | 
 **ignoreStartDates** | **bool** |  | 
 **getSubTasks** | **bool** |  | 
 **getSubTasks2** | **bool** |  | 
 **getFiles** | **bool** |  | 
 **getFiles2** | **bool** |  | 
 **countOnly** | **bool** |  | 
 **countOnly2** | **bool** |  | 
 **allowTemplateTasks** | **bool** |  | 
 **projectTagIds** | **[]int32** |  | 
 **projectTagIds2** | **[]int32** |  | 
 **projectOwnerIds** | **[]int32** |  | 
 **projectOwnerIds2** | **[]int32** |  | 
 **projectIds** | **[]int32** |  | 
 **projectIds2** | **[]int32** |  | 
 **projectId** | **[]int32** | project *filters | 
 **projectId2** | **[]int32** | project *filters | 
 **projectHealths** | **[]int32** |  | 
 **projectHealths2** | **[]int32** |  | 
 **projectCompanyIds** | **[]int32** |  | 
 **projectCompanyIds2** | **[]int32** |  | 
 **projectCategoryIds** | **[]int32** |  | 
 **projectCategoryIds2** | **[]int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

