# InlineResponse20097TodoItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DLM** | Pointer to **int32** |  | [optional] 
**AttachmentsCount** | Pointer to **int32** |  | [optional] 
**CanComplete** | Pointer to **bool** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] 
**CanLogTime** | Pointer to **bool** |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**CompanyId** | Pointer to **int32** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatorAvatarUrl** | Pointer to **string** |  | [optional] 
**CreatorFirstname** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **int32** |  | [optional] 
**CreatorLastname** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** | If a task due date is defined, it will be returned in this field matching due date base. If there is no task due date set, but a milestone is assoicated, it will return the milestone due date. | [optional] 
**DueDateBase** | Pointer to **string** | If the task has a defined due date, this will be set. Otherwise, it will be empty.  | [optional] 
**EstimatedMinutes** | Pointer to **int32** |  | [optional] 
**HarvestEnabled** | Pointer to **bool** |  | [optional] 
**HasDependencies** | Pointer to **int32** | It is a flag on the task level to indicate whether another task is dependant on it or not. For example, if you had a task which had 2 predecessors. They must be completed before the parent task can be. If you made an api call for one of those predecessor tasks, the has-dependencies flag would be set to 1 as it has task/tasks dependant on it. | [optional] 
**HasReminders** | Pointer to **bool** |  | [optional] 
**HasUnreadComments** | Pointer to **bool** |  | [optional] 
**HasTickets** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LockdownId** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Predecessors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **int32** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**ResponsiblePartyFirstname** | Pointer to **string** |  | [optional] 
**ResponsiblePartyId** | Pointer to **string** |  | [optional] 
**ResponsiblePartyIds** | Pointer to **string** |  | [optional] 
**ResponsiblePartyLastname** | Pointer to **string** |  | [optional] 
**ResponsiblePartyNames** | Pointer to **string** |  | [optional] 
**ResponsiblePartySummary** | Pointer to **string** |  | [optional] 
**ResponsiblePartyType** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]InlineResponse20054Category**](InlineResponse20054Category.md) |  | [optional] 
**TasklistIsTemplate** | Pointer to **bool** |  | [optional] 
**TasklistLockdownId** | Pointer to **string** |  | [optional] 
**TasklistPrivate** | Pointer to **bool** |  | [optional] 
**TimeIsLogged** | Pointer to **string** |  | [optional] 
**TodoListId** | Pointer to **int32** |  | [optional] 
**TodoListName** | Pointer to **string** |  | [optional] 
**UserFollowingChanges** | Pointer to **bool** |  | [optional] 
**UserFollowingComments** | Pointer to **bool** |  | [optional] 
**ViewEstimatedTime** | Pointer to **bool** |  | [optional] 

## Methods

### NewInlineResponse20097TodoItems

`func NewInlineResponse20097TodoItems() *InlineResponse20097TodoItems`

NewInlineResponse20097TodoItems instantiates a new InlineResponse20097TodoItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20097TodoItemsWithDefaults

`func NewInlineResponse20097TodoItemsWithDefaults() *InlineResponse20097TodoItems`

NewInlineResponse20097TodoItemsWithDefaults instantiates a new InlineResponse20097TodoItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDLM

`func (o *InlineResponse20097TodoItems) GetDLM() int32`

GetDLM returns the DLM field if non-nil, zero value otherwise.

### GetDLMOk

`func (o *InlineResponse20097TodoItems) GetDLMOk() (*int32, bool)`

GetDLMOk returns a tuple with the DLM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLM

`func (o *InlineResponse20097TodoItems) SetDLM(v int32)`

SetDLM sets DLM field to given value.

### HasDLM

`func (o *InlineResponse20097TodoItems) HasDLM() bool`

HasDLM returns a boolean if a field has been set.

### GetAttachmentsCount

`func (o *InlineResponse20097TodoItems) GetAttachmentsCount() int32`

GetAttachmentsCount returns the AttachmentsCount field if non-nil, zero value otherwise.

### GetAttachmentsCountOk

`func (o *InlineResponse20097TodoItems) GetAttachmentsCountOk() (*int32, bool)`

GetAttachmentsCountOk returns a tuple with the AttachmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsCount

`func (o *InlineResponse20097TodoItems) SetAttachmentsCount(v int32)`

SetAttachmentsCount sets AttachmentsCount field to given value.

### HasAttachmentsCount

`func (o *InlineResponse20097TodoItems) HasAttachmentsCount() bool`

HasAttachmentsCount returns a boolean if a field has been set.

### GetCanComplete

`func (o *InlineResponse20097TodoItems) GetCanComplete() bool`

GetCanComplete returns the CanComplete field if non-nil, zero value otherwise.

### GetCanCompleteOk

`func (o *InlineResponse20097TodoItems) GetCanCompleteOk() (*bool, bool)`

GetCanCompleteOk returns a tuple with the CanComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanComplete

`func (o *InlineResponse20097TodoItems) SetCanComplete(v bool)`

SetCanComplete sets CanComplete field to given value.

### HasCanComplete

`func (o *InlineResponse20097TodoItems) HasCanComplete() bool`

HasCanComplete returns a boolean if a field has been set.

### GetCanEdit

`func (o *InlineResponse20097TodoItems) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *InlineResponse20097TodoItems) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *InlineResponse20097TodoItems) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *InlineResponse20097TodoItems) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCanLogTime

`func (o *InlineResponse20097TodoItems) GetCanLogTime() bool`

GetCanLogTime returns the CanLogTime field if non-nil, zero value otherwise.

### GetCanLogTimeOk

`func (o *InlineResponse20097TodoItems) GetCanLogTimeOk() (*bool, bool)`

GetCanLogTimeOk returns a tuple with the CanLogTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLogTime

`func (o *InlineResponse20097TodoItems) SetCanLogTime(v bool)`

SetCanLogTime sets CanLogTime field to given value.

### HasCanLogTime

`func (o *InlineResponse20097TodoItems) HasCanLogTime() bool`

HasCanLogTime returns a boolean if a field has been set.

### GetCommentsCount

`func (o *InlineResponse20097TodoItems) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *InlineResponse20097TodoItems) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *InlineResponse20097TodoItems) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *InlineResponse20097TodoItems) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse20097TodoItems) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20097TodoItems) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20097TodoItems) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20097TodoItems) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20097TodoItems) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20097TodoItems) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20097TodoItems) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20097TodoItems) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse20097TodoItems) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse20097TodoItems) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse20097TodoItems) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse20097TodoItems) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetContent

`func (o *InlineResponse20097TodoItems) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InlineResponse20097TodoItems) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InlineResponse20097TodoItems) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *InlineResponse20097TodoItems) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedOn

`func (o *InlineResponse20097TodoItems) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *InlineResponse20097TodoItems) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *InlineResponse20097TodoItems) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *InlineResponse20097TodoItems) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatorAvatarUrl

`func (o *InlineResponse20097TodoItems) GetCreatorAvatarUrl() string`

GetCreatorAvatarUrl returns the CreatorAvatarUrl field if non-nil, zero value otherwise.

### GetCreatorAvatarUrlOk

`func (o *InlineResponse20097TodoItems) GetCreatorAvatarUrlOk() (*string, bool)`

GetCreatorAvatarUrlOk returns a tuple with the CreatorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorAvatarUrl

`func (o *InlineResponse20097TodoItems) SetCreatorAvatarUrl(v string)`

SetCreatorAvatarUrl sets CreatorAvatarUrl field to given value.

### HasCreatorAvatarUrl

`func (o *InlineResponse20097TodoItems) HasCreatorAvatarUrl() bool`

HasCreatorAvatarUrl returns a boolean if a field has been set.

### GetCreatorFirstname

`func (o *InlineResponse20097TodoItems) GetCreatorFirstname() string`

GetCreatorFirstname returns the CreatorFirstname field if non-nil, zero value otherwise.

### GetCreatorFirstnameOk

`func (o *InlineResponse20097TodoItems) GetCreatorFirstnameOk() (*string, bool)`

GetCreatorFirstnameOk returns a tuple with the CreatorFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFirstname

`func (o *InlineResponse20097TodoItems) SetCreatorFirstname(v string)`

SetCreatorFirstname sets CreatorFirstname field to given value.

### HasCreatorFirstname

`func (o *InlineResponse20097TodoItems) HasCreatorFirstname() bool`

HasCreatorFirstname returns a boolean if a field has been set.

### GetCreatorId

`func (o *InlineResponse20097TodoItems) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *InlineResponse20097TodoItems) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *InlineResponse20097TodoItems) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *InlineResponse20097TodoItems) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorLastname

`func (o *InlineResponse20097TodoItems) GetCreatorLastname() string`

GetCreatorLastname returns the CreatorLastname field if non-nil, zero value otherwise.

### GetCreatorLastnameOk

`func (o *InlineResponse20097TodoItems) GetCreatorLastnameOk() (*string, bool)`

GetCreatorLastnameOk returns a tuple with the CreatorLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorLastname

`func (o *InlineResponse20097TodoItems) SetCreatorLastname(v string)`

SetCreatorLastname sets CreatorLastname field to given value.

### HasCreatorLastname

`func (o *InlineResponse20097TodoItems) HasCreatorLastname() bool`

HasCreatorLastname returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20097TodoItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20097TodoItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20097TodoItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20097TodoItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *InlineResponse20097TodoItems) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InlineResponse20097TodoItems) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InlineResponse20097TodoItems) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InlineResponse20097TodoItems) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetDueDateBase

`func (o *InlineResponse20097TodoItems) GetDueDateBase() string`

GetDueDateBase returns the DueDateBase field if non-nil, zero value otherwise.

### GetDueDateBaseOk

`func (o *InlineResponse20097TodoItems) GetDueDateBaseOk() (*string, bool)`

GetDueDateBaseOk returns a tuple with the DueDateBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateBase

`func (o *InlineResponse20097TodoItems) SetDueDateBase(v string)`

SetDueDateBase sets DueDateBase field to given value.

### HasDueDateBase

`func (o *InlineResponse20097TodoItems) HasDueDateBase() bool`

HasDueDateBase returns a boolean if a field has been set.

### GetEstimatedMinutes

`func (o *InlineResponse20097TodoItems) GetEstimatedMinutes() int32`

GetEstimatedMinutes returns the EstimatedMinutes field if non-nil, zero value otherwise.

### GetEstimatedMinutesOk

`func (o *InlineResponse20097TodoItems) GetEstimatedMinutesOk() (*int32, bool)`

GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutes

`func (o *InlineResponse20097TodoItems) SetEstimatedMinutes(v int32)`

SetEstimatedMinutes sets EstimatedMinutes field to given value.

### HasEstimatedMinutes

`func (o *InlineResponse20097TodoItems) HasEstimatedMinutes() bool`

HasEstimatedMinutes returns a boolean if a field has been set.

### GetHarvestEnabled

`func (o *InlineResponse20097TodoItems) GetHarvestEnabled() bool`

GetHarvestEnabled returns the HarvestEnabled field if non-nil, zero value otherwise.

### GetHarvestEnabledOk

`func (o *InlineResponse20097TodoItems) GetHarvestEnabledOk() (*bool, bool)`

GetHarvestEnabledOk returns a tuple with the HarvestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarvestEnabled

`func (o *InlineResponse20097TodoItems) SetHarvestEnabled(v bool)`

SetHarvestEnabled sets HarvestEnabled field to given value.

### HasHarvestEnabled

`func (o *InlineResponse20097TodoItems) HasHarvestEnabled() bool`

HasHarvestEnabled returns a boolean if a field has been set.

### GetHasDependencies

`func (o *InlineResponse20097TodoItems) GetHasDependencies() int32`

GetHasDependencies returns the HasDependencies field if non-nil, zero value otherwise.

### GetHasDependenciesOk

`func (o *InlineResponse20097TodoItems) GetHasDependenciesOk() (*int32, bool)`

GetHasDependenciesOk returns a tuple with the HasDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDependencies

`func (o *InlineResponse20097TodoItems) SetHasDependencies(v int32)`

SetHasDependencies sets HasDependencies field to given value.

### HasHasDependencies

`func (o *InlineResponse20097TodoItems) HasHasDependencies() bool`

HasHasDependencies returns a boolean if a field has been set.

### GetHasReminders

`func (o *InlineResponse20097TodoItems) GetHasReminders() bool`

GetHasReminders returns the HasReminders field if non-nil, zero value otherwise.

### GetHasRemindersOk

`func (o *InlineResponse20097TodoItems) GetHasRemindersOk() (*bool, bool)`

GetHasRemindersOk returns a tuple with the HasReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReminders

`func (o *InlineResponse20097TodoItems) SetHasReminders(v bool)`

SetHasReminders sets HasReminders field to given value.

### HasHasReminders

`func (o *InlineResponse20097TodoItems) HasHasReminders() bool`

HasHasReminders returns a boolean if a field has been set.

### GetHasUnreadComments

`func (o *InlineResponse20097TodoItems) GetHasUnreadComments() bool`

GetHasUnreadComments returns the HasUnreadComments field if non-nil, zero value otherwise.

### GetHasUnreadCommentsOk

`func (o *InlineResponse20097TodoItems) GetHasUnreadCommentsOk() (*bool, bool)`

GetHasUnreadCommentsOk returns a tuple with the HasUnreadComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnreadComments

`func (o *InlineResponse20097TodoItems) SetHasUnreadComments(v bool)`

SetHasUnreadComments sets HasUnreadComments field to given value.

### HasHasUnreadComments

`func (o *InlineResponse20097TodoItems) HasHasUnreadComments() bool`

HasHasUnreadComments returns a boolean if a field has been set.

### GetHasTickets

`func (o *InlineResponse20097TodoItems) GetHasTickets() bool`

GetHasTickets returns the HasTickets field if non-nil, zero value otherwise.

### GetHasTicketsOk

`func (o *InlineResponse20097TodoItems) GetHasTicketsOk() (*bool, bool)`

GetHasTicketsOk returns a tuple with the HasTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTickets

`func (o *InlineResponse20097TodoItems) SetHasTickets(v bool)`

SetHasTickets sets HasTickets field to given value.

### HasHasTickets

`func (o *InlineResponse20097TodoItems) HasHasTickets() bool`

HasHasTickets returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20097TodoItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20097TodoItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20097TodoItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20097TodoItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20097TodoItems) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20097TodoItems) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20097TodoItems) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20097TodoItems) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLockdownId

`func (o *InlineResponse20097TodoItems) GetLockdownId() string`

GetLockdownId returns the LockdownId field if non-nil, zero value otherwise.

### GetLockdownIdOk

`func (o *InlineResponse20097TodoItems) GetLockdownIdOk() (*string, bool)`

GetLockdownIdOk returns a tuple with the LockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdownId

`func (o *InlineResponse20097TodoItems) SetLockdownId(v string)`

SetLockdownId sets LockdownId field to given value.

### HasLockdownId

`func (o *InlineResponse20097TodoItems) HasLockdownId() bool`

HasLockdownId returns a boolean if a field has been set.

### GetOrder

`func (o *InlineResponse20097TodoItems) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InlineResponse20097TodoItems) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InlineResponse20097TodoItems) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InlineResponse20097TodoItems) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentTaskId

`func (o *InlineResponse20097TodoItems) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *InlineResponse20097TodoItems) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *InlineResponse20097TodoItems) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *InlineResponse20097TodoItems) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetPosition

`func (o *InlineResponse20097TodoItems) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InlineResponse20097TodoItems) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InlineResponse20097TodoItems) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InlineResponse20097TodoItems) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPredecessors

`func (o *InlineResponse20097TodoItems) GetPredecessors() []map[string]interface{}`

GetPredecessors returns the Predecessors field if non-nil, zero value otherwise.

### GetPredecessorsOk

`func (o *InlineResponse20097TodoItems) GetPredecessorsOk() (*[]map[string]interface{}, bool)`

GetPredecessorsOk returns a tuple with the Predecessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessors

`func (o *InlineResponse20097TodoItems) SetPredecessors(v []map[string]interface{})`

SetPredecessors sets Predecessors field to given value.

### HasPredecessors

`func (o *InlineResponse20097TodoItems) HasPredecessors() bool`

HasPredecessors returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse20097TodoItems) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse20097TodoItems) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse20097TodoItems) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse20097TodoItems) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivate

`func (o *InlineResponse20097TodoItems) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *InlineResponse20097TodoItems) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *InlineResponse20097TodoItems) SetPrivate(v int32)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *InlineResponse20097TodoItems) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProgress

`func (o *InlineResponse20097TodoItems) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InlineResponse20097TodoItems) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InlineResponse20097TodoItems) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InlineResponse20097TodoItems) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse20097TodoItems) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse20097TodoItems) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse20097TodoItems) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse20097TodoItems) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *InlineResponse20097TodoItems) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InlineResponse20097TodoItems) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InlineResponse20097TodoItems) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InlineResponse20097TodoItems) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetResponsiblePartyFirstname

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyFirstname() string`

GetResponsiblePartyFirstname returns the ResponsiblePartyFirstname field if non-nil, zero value otherwise.

### GetResponsiblePartyFirstnameOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyFirstnameOk() (*string, bool)`

GetResponsiblePartyFirstnameOk returns a tuple with the ResponsiblePartyFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyFirstname

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyFirstname(v string)`

SetResponsiblePartyFirstname sets ResponsiblePartyFirstname field to given value.

### HasResponsiblePartyFirstname

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyFirstname() bool`

HasResponsiblePartyFirstname returns a boolean if a field has been set.

### GetResponsiblePartyId

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyId() string`

GetResponsiblePartyId returns the ResponsiblePartyId field if non-nil, zero value otherwise.

### GetResponsiblePartyIdOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyIdOk() (*string, bool)`

GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyId

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyId(v string)`

SetResponsiblePartyId sets ResponsiblePartyId field to given value.

### HasResponsiblePartyId

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyId() bool`

HasResponsiblePartyId returns a boolean if a field has been set.

### GetResponsiblePartyIds

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyIds() string`

GetResponsiblePartyIds returns the ResponsiblePartyIds field if non-nil, zero value otherwise.

### GetResponsiblePartyIdsOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyIdsOk() (*string, bool)`

GetResponsiblePartyIdsOk returns a tuple with the ResponsiblePartyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyIds

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyIds(v string)`

SetResponsiblePartyIds sets ResponsiblePartyIds field to given value.

### HasResponsiblePartyIds

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyIds() bool`

HasResponsiblePartyIds returns a boolean if a field has been set.

### GetResponsiblePartyLastname

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyLastname() string`

GetResponsiblePartyLastname returns the ResponsiblePartyLastname field if non-nil, zero value otherwise.

### GetResponsiblePartyLastnameOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyLastnameOk() (*string, bool)`

GetResponsiblePartyLastnameOk returns a tuple with the ResponsiblePartyLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyLastname

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyLastname(v string)`

SetResponsiblePartyLastname sets ResponsiblePartyLastname field to given value.

### HasResponsiblePartyLastname

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyLastname() bool`

HasResponsiblePartyLastname returns a boolean if a field has been set.

### GetResponsiblePartyNames

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyNames() string`

GetResponsiblePartyNames returns the ResponsiblePartyNames field if non-nil, zero value otherwise.

### GetResponsiblePartyNamesOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyNamesOk() (*string, bool)`

GetResponsiblePartyNamesOk returns a tuple with the ResponsiblePartyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyNames

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyNames(v string)`

SetResponsiblePartyNames sets ResponsiblePartyNames field to given value.

### HasResponsiblePartyNames

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyNames() bool`

HasResponsiblePartyNames returns a boolean if a field has been set.

### GetResponsiblePartySummary

`func (o *InlineResponse20097TodoItems) GetResponsiblePartySummary() string`

GetResponsiblePartySummary returns the ResponsiblePartySummary field if non-nil, zero value otherwise.

### GetResponsiblePartySummaryOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartySummaryOk() (*string, bool)`

GetResponsiblePartySummaryOk returns a tuple with the ResponsiblePartySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartySummary

`func (o *InlineResponse20097TodoItems) SetResponsiblePartySummary(v string)`

SetResponsiblePartySummary sets ResponsiblePartySummary field to given value.

### HasResponsiblePartySummary

`func (o *InlineResponse20097TodoItems) HasResponsiblePartySummary() bool`

HasResponsiblePartySummary returns a boolean if a field has been set.

### GetResponsiblePartyType

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyType() string`

GetResponsiblePartyType returns the ResponsiblePartyType field if non-nil, zero value otherwise.

### GetResponsiblePartyTypeOk

`func (o *InlineResponse20097TodoItems) GetResponsiblePartyTypeOk() (*string, bool)`

GetResponsiblePartyTypeOk returns a tuple with the ResponsiblePartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyType

`func (o *InlineResponse20097TodoItems) SetResponsiblePartyType(v string)`

SetResponsiblePartyType sets ResponsiblePartyType field to given value.

### HasResponsiblePartyType

`func (o *InlineResponse20097TodoItems) HasResponsiblePartyType() bool`

HasResponsiblePartyType returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse20097TodoItems) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse20097TodoItems) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse20097TodoItems) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse20097TodoItems) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20097TodoItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20097TodoItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20097TodoItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20097TodoItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20097TodoItems) GetTags() []InlineResponse20054Category`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20097TodoItems) GetTagsOk() (*[]InlineResponse20054Category, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20097TodoItems) SetTags(v []InlineResponse20054Category)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20097TodoItems) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasklistIsTemplate

`func (o *InlineResponse20097TodoItems) GetTasklistIsTemplate() bool`

GetTasklistIsTemplate returns the TasklistIsTemplate field if non-nil, zero value otherwise.

### GetTasklistIsTemplateOk

`func (o *InlineResponse20097TodoItems) GetTasklistIsTemplateOk() (*bool, bool)`

GetTasklistIsTemplateOk returns a tuple with the TasklistIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistIsTemplate

`func (o *InlineResponse20097TodoItems) SetTasklistIsTemplate(v bool)`

SetTasklistIsTemplate sets TasklistIsTemplate field to given value.

### HasTasklistIsTemplate

`func (o *InlineResponse20097TodoItems) HasTasklistIsTemplate() bool`

HasTasklistIsTemplate returns a boolean if a field has been set.

### GetTasklistLockdownId

`func (o *InlineResponse20097TodoItems) GetTasklistLockdownId() string`

GetTasklistLockdownId returns the TasklistLockdownId field if non-nil, zero value otherwise.

### GetTasklistLockdownIdOk

`func (o *InlineResponse20097TodoItems) GetTasklistLockdownIdOk() (*string, bool)`

GetTasklistLockdownIdOk returns a tuple with the TasklistLockdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistLockdownId

`func (o *InlineResponse20097TodoItems) SetTasklistLockdownId(v string)`

SetTasklistLockdownId sets TasklistLockdownId field to given value.

### HasTasklistLockdownId

`func (o *InlineResponse20097TodoItems) HasTasklistLockdownId() bool`

HasTasklistLockdownId returns a boolean if a field has been set.

### GetTasklistPrivate

`func (o *InlineResponse20097TodoItems) GetTasklistPrivate() bool`

GetTasklistPrivate returns the TasklistPrivate field if non-nil, zero value otherwise.

### GetTasklistPrivateOk

`func (o *InlineResponse20097TodoItems) GetTasklistPrivateOk() (*bool, bool)`

GetTasklistPrivateOk returns a tuple with the TasklistPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistPrivate

`func (o *InlineResponse20097TodoItems) SetTasklistPrivate(v bool)`

SetTasklistPrivate sets TasklistPrivate field to given value.

### HasTasklistPrivate

`func (o *InlineResponse20097TodoItems) HasTasklistPrivate() bool`

HasTasklistPrivate returns a boolean if a field has been set.

### GetTimeIsLogged

`func (o *InlineResponse20097TodoItems) GetTimeIsLogged() string`

GetTimeIsLogged returns the TimeIsLogged field if non-nil, zero value otherwise.

### GetTimeIsLoggedOk

`func (o *InlineResponse20097TodoItems) GetTimeIsLoggedOk() (*string, bool)`

GetTimeIsLoggedOk returns a tuple with the TimeIsLogged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIsLogged

`func (o *InlineResponse20097TodoItems) SetTimeIsLogged(v string)`

SetTimeIsLogged sets TimeIsLogged field to given value.

### HasTimeIsLogged

`func (o *InlineResponse20097TodoItems) HasTimeIsLogged() bool`

HasTimeIsLogged returns a boolean if a field has been set.

### GetTodoListId

`func (o *InlineResponse20097TodoItems) GetTodoListId() int32`

GetTodoListId returns the TodoListId field if non-nil, zero value otherwise.

### GetTodoListIdOk

`func (o *InlineResponse20097TodoItems) GetTodoListIdOk() (*int32, bool)`

GetTodoListIdOk returns a tuple with the TodoListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListId

`func (o *InlineResponse20097TodoItems) SetTodoListId(v int32)`

SetTodoListId sets TodoListId field to given value.

### HasTodoListId

`func (o *InlineResponse20097TodoItems) HasTodoListId() bool`

HasTodoListId returns a boolean if a field has been set.

### GetTodoListName

`func (o *InlineResponse20097TodoItems) GetTodoListName() string`

GetTodoListName returns the TodoListName field if non-nil, zero value otherwise.

### GetTodoListNameOk

`func (o *InlineResponse20097TodoItems) GetTodoListNameOk() (*string, bool)`

GetTodoListNameOk returns a tuple with the TodoListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoListName

`func (o *InlineResponse20097TodoItems) SetTodoListName(v string)`

SetTodoListName sets TodoListName field to given value.

### HasTodoListName

`func (o *InlineResponse20097TodoItems) HasTodoListName() bool`

HasTodoListName returns a boolean if a field has been set.

### GetUserFollowingChanges

`func (o *InlineResponse20097TodoItems) GetUserFollowingChanges() bool`

GetUserFollowingChanges returns the UserFollowingChanges field if non-nil, zero value otherwise.

### GetUserFollowingChangesOk

`func (o *InlineResponse20097TodoItems) GetUserFollowingChangesOk() (*bool, bool)`

GetUserFollowingChangesOk returns a tuple with the UserFollowingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingChanges

`func (o *InlineResponse20097TodoItems) SetUserFollowingChanges(v bool)`

SetUserFollowingChanges sets UserFollowingChanges field to given value.

### HasUserFollowingChanges

`func (o *InlineResponse20097TodoItems) HasUserFollowingChanges() bool`

HasUserFollowingChanges returns a boolean if a field has been set.

### GetUserFollowingComments

`func (o *InlineResponse20097TodoItems) GetUserFollowingComments() bool`

GetUserFollowingComments returns the UserFollowingComments field if non-nil, zero value otherwise.

### GetUserFollowingCommentsOk

`func (o *InlineResponse20097TodoItems) GetUserFollowingCommentsOk() (*bool, bool)`

GetUserFollowingCommentsOk returns a tuple with the UserFollowingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFollowingComments

`func (o *InlineResponse20097TodoItems) SetUserFollowingComments(v bool)`

SetUserFollowingComments sets UserFollowingComments field to given value.

### HasUserFollowingComments

`func (o *InlineResponse20097TodoItems) HasUserFollowingComments() bool`

HasUserFollowingComments returns a boolean if a field has been set.

### GetViewEstimatedTime

`func (o *InlineResponse20097TodoItems) GetViewEstimatedTime() bool`

GetViewEstimatedTime returns the ViewEstimatedTime field if non-nil, zero value otherwise.

### GetViewEstimatedTimeOk

`func (o *InlineResponse20097TodoItems) GetViewEstimatedTimeOk() (*bool, bool)`

GetViewEstimatedTimeOk returns a tuple with the ViewEstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEstimatedTime

`func (o *InlineResponse20097TodoItems) SetViewEstimatedTime(v bool)`

SetViewEstimatedTime sets ViewEstimatedTime field to given value.

### HasViewEstimatedTime

`func (o *InlineResponse20097TodoItems) HasViewEstimatedTime() bool`

HasViewEstimatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


