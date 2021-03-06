/*
 * Teamwork Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmv1

import (
	"encoding/json"
)

// TasksIdJsonTodoItem1 struct for TasksIdJsonTodoItem1
type TasksIdJsonTodoItem1 struct {
	// Existing files can be attached to the task by specifying a comma-separated list of File IDs.
	Attachments *[]string `json:"attachments,omitempty"`
	// id's of people which will be followers. Change id means if notify=true they will get notified if any changes occur on that task.
	ChangeFollowerIds *string `json:"changeFollowerIds,omitempty"`
	// Used to put a new Task directly in to a Column in the Boards view
	ColumnId *string `json:"columnId,omitempty"`
	// id's of people which will be followers if someone creates a comment. Comment id means if notify=true they will get notified if any comments are created on that task.
	CommentFollowerIds *string `json:"commentFollowerIds,omitempty"`
	// The name of the task you are adding.
	Content string `json:"content"`
	// Tasks can be assigned a description.
	Description *string `json:"description,omitempty"`
	// Tasks can be assigned a date for when they are due to be finished. The format should be YYYYMMDD.
	DueDate *string `json:"due-date,omitempty"`
	// Set an estimated number of minutes for a task to be completed by setting this parameter.
	EstimatedMinutes *string `json:"estimated-minutes,omitempty"`
	// Used in relation to private. If you set a Task as private (1), you can optionally add one or more User IDs to make the task visible to
	GrantAccessTo *string `json:"grant-access-to,omitempty"`
	// This parameter can be set to true to notify people assigned to this task by email.
	Notify *bool `json:"notify,omitempty"`
	// Set this to the ID of a parent task if you wish to make your task a subtask.
	ParentTaskId *string `json:"parentTaskId,omitempty"`
	// New files can be attached using this option (see the uploading files section for more info). This is a comma-separated list of PendingFileRef's.
	PendingFileAttachments *string `json:"pendingFileAttachments,omitempty"`
	// A task can be placed after another task by setting this parameter to a task ID. Additionally, some other options are available: -2 - Ignore -1 - Place task at the top of the list 0 - Place task at the bottom of the list
	PositionAfterTask *int32 `json:"positionAfterTask,omitempty"`
	// An array of predecessor objects can be passed to specify which tasks need to have a specified state before this task can be marked as completed. Each predecessor object should contain two keys: id: The ID of the task that predecesses this one. type: The state the task needs to be, can be \"complete\" or \"start\".
	Predecessors *[]TasksIdJsonTodoItem1Predecessors `json:"predecessors,omitempty"`
	// Tasks can be assigned the following priorities: - not set - low - medium - high
	Priority *string `json:"priority,omitempty"`
	// Set to 1 to make the task Private. Setting a 0 will set the Task back to normal
	Private *int32 `json:"private,omitempty"`
	// Set the progress from 0 to 90
	Progress *string `json:"progress,omitempty"`
	// This can be used to assign the new task to a person or group of people. For example: -1 would assign the task to everyone 32 would assign the task to user 32. 32,55 would assign the task to users 32 and 55 etc.
	ResponsiblePartyId *int32 `json:"responsible-party-id,omitempty"`
	// Tasks can be assigned a date to start on. The format should be YYYYMMDD.
	StartDate *string `json:"start-date,omitempty"`
	// A comma separated list of tags for the task
	Tags *string `json:"tags,omitempty"`
}

// NewTasksIdJsonTodoItem1 instantiates a new TasksIdJsonTodoItem1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksIdJsonTodoItem1(content string) *TasksIdJsonTodoItem1 {
	this := TasksIdJsonTodoItem1{}
	this.Content = content
	return &this
}

// NewTasksIdJsonTodoItem1WithDefaults instantiates a new TasksIdJsonTodoItem1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksIdJsonTodoItem1WithDefaults() *TasksIdJsonTodoItem1 {
	this := TasksIdJsonTodoItem1{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetAttachments() []string {
	if o == nil || o.Attachments == nil {
		var ret []string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetAttachmentsOk() (*[]string, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *TasksIdJsonTodoItem1) SetAttachments(v []string) {
	o.Attachments = &v
}

// GetChangeFollowerIds returns the ChangeFollowerIds field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetChangeFollowerIds() string {
	if o == nil || o.ChangeFollowerIds == nil {
		var ret string
		return ret
	}
	return *o.ChangeFollowerIds
}

// GetChangeFollowerIdsOk returns a tuple with the ChangeFollowerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetChangeFollowerIdsOk() (*string, bool) {
	if o == nil || o.ChangeFollowerIds == nil {
		return nil, false
	}
	return o.ChangeFollowerIds, true
}

// HasChangeFollowerIds returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasChangeFollowerIds() bool {
	if o != nil && o.ChangeFollowerIds != nil {
		return true
	}

	return false
}

// SetChangeFollowerIds gets a reference to the given string and assigns it to the ChangeFollowerIds field.
func (o *TasksIdJsonTodoItem1) SetChangeFollowerIds(v string) {
	o.ChangeFollowerIds = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetColumnId() string {
	if o == nil || o.ColumnId == nil {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetColumnIdOk() (*string, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *TasksIdJsonTodoItem1) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetCommentFollowerIds returns the CommentFollowerIds field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetCommentFollowerIds() string {
	if o == nil || o.CommentFollowerIds == nil {
		var ret string
		return ret
	}
	return *o.CommentFollowerIds
}

// GetCommentFollowerIdsOk returns a tuple with the CommentFollowerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetCommentFollowerIdsOk() (*string, bool) {
	if o == nil || o.CommentFollowerIds == nil {
		return nil, false
	}
	return o.CommentFollowerIds, true
}

// HasCommentFollowerIds returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasCommentFollowerIds() bool {
	if o != nil && o.CommentFollowerIds != nil {
		return true
	}

	return false
}

// SetCommentFollowerIds gets a reference to the given string and assigns it to the CommentFollowerIds field.
func (o *TasksIdJsonTodoItem1) SetCommentFollowerIds(v string) {
	o.CommentFollowerIds = &v
}

// GetContent returns the Content field value
func (o *TasksIdJsonTodoItem1) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *TasksIdJsonTodoItem1) SetContent(v string) {
	o.Content = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TasksIdJsonTodoItem1) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *TasksIdJsonTodoItem1) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEstimatedMinutes returns the EstimatedMinutes field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetEstimatedMinutes() string {
	if o == nil || o.EstimatedMinutes == nil {
		var ret string
		return ret
	}
	return *o.EstimatedMinutes
}

// GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetEstimatedMinutesOk() (*string, bool) {
	if o == nil || o.EstimatedMinutes == nil {
		return nil, false
	}
	return o.EstimatedMinutes, true
}

// HasEstimatedMinutes returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasEstimatedMinutes() bool {
	if o != nil && o.EstimatedMinutes != nil {
		return true
	}

	return false
}

// SetEstimatedMinutes gets a reference to the given string and assigns it to the EstimatedMinutes field.
func (o *TasksIdJsonTodoItem1) SetEstimatedMinutes(v string) {
	o.EstimatedMinutes = &v
}

// GetGrantAccessTo returns the GrantAccessTo field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetGrantAccessTo() string {
	if o == nil || o.GrantAccessTo == nil {
		var ret string
		return ret
	}
	return *o.GrantAccessTo
}

// GetGrantAccessToOk returns a tuple with the GrantAccessTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetGrantAccessToOk() (*string, bool) {
	if o == nil || o.GrantAccessTo == nil {
		return nil, false
	}
	return o.GrantAccessTo, true
}

// HasGrantAccessTo returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasGrantAccessTo() bool {
	if o != nil && o.GrantAccessTo != nil {
		return true
	}

	return false
}

// SetGrantAccessTo gets a reference to the given string and assigns it to the GrantAccessTo field.
func (o *TasksIdJsonTodoItem1) SetGrantAccessTo(v string) {
	o.GrantAccessTo = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetNotify() bool {
	if o == nil || o.Notify == nil {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetNotifyOk() (*bool, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *TasksIdJsonTodoItem1) SetNotify(v bool) {
	o.Notify = &v
}

// GetParentTaskId returns the ParentTaskId field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetParentTaskId() string {
	if o == nil || o.ParentTaskId == nil {
		var ret string
		return ret
	}
	return *o.ParentTaskId
}

// GetParentTaskIdOk returns a tuple with the ParentTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetParentTaskIdOk() (*string, bool) {
	if o == nil || o.ParentTaskId == nil {
		return nil, false
	}
	return o.ParentTaskId, true
}

// HasParentTaskId returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasParentTaskId() bool {
	if o != nil && o.ParentTaskId != nil {
		return true
	}

	return false
}

// SetParentTaskId gets a reference to the given string and assigns it to the ParentTaskId field.
func (o *TasksIdJsonTodoItem1) SetParentTaskId(v string) {
	o.ParentTaskId = &v
}

// GetPendingFileAttachments returns the PendingFileAttachments field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetPendingFileAttachments() string {
	if o == nil || o.PendingFileAttachments == nil {
		var ret string
		return ret
	}
	return *o.PendingFileAttachments
}

// GetPendingFileAttachmentsOk returns a tuple with the PendingFileAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetPendingFileAttachmentsOk() (*string, bool) {
	if o == nil || o.PendingFileAttachments == nil {
		return nil, false
	}
	return o.PendingFileAttachments, true
}

// HasPendingFileAttachments returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasPendingFileAttachments() bool {
	if o != nil && o.PendingFileAttachments != nil {
		return true
	}

	return false
}

// SetPendingFileAttachments gets a reference to the given string and assigns it to the PendingFileAttachments field.
func (o *TasksIdJsonTodoItem1) SetPendingFileAttachments(v string) {
	o.PendingFileAttachments = &v
}

// GetPositionAfterTask returns the PositionAfterTask field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetPositionAfterTask() int32 {
	if o == nil || o.PositionAfterTask == nil {
		var ret int32
		return ret
	}
	return *o.PositionAfterTask
}

// GetPositionAfterTaskOk returns a tuple with the PositionAfterTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetPositionAfterTaskOk() (*int32, bool) {
	if o == nil || o.PositionAfterTask == nil {
		return nil, false
	}
	return o.PositionAfterTask, true
}

// HasPositionAfterTask returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasPositionAfterTask() bool {
	if o != nil && o.PositionAfterTask != nil {
		return true
	}

	return false
}

// SetPositionAfterTask gets a reference to the given int32 and assigns it to the PositionAfterTask field.
func (o *TasksIdJsonTodoItem1) SetPositionAfterTask(v int32) {
	o.PositionAfterTask = &v
}

// GetPredecessors returns the Predecessors field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetPredecessors() []TasksIdJsonTodoItem1Predecessors {
	if o == nil || o.Predecessors == nil {
		var ret []TasksIdJsonTodoItem1Predecessors
		return ret
	}
	return *o.Predecessors
}

// GetPredecessorsOk returns a tuple with the Predecessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetPredecessorsOk() (*[]TasksIdJsonTodoItem1Predecessors, bool) {
	if o == nil || o.Predecessors == nil {
		return nil, false
	}
	return o.Predecessors, true
}

// HasPredecessors returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasPredecessors() bool {
	if o != nil && o.Predecessors != nil {
		return true
	}

	return false
}

// SetPredecessors gets a reference to the given []TasksIdJsonTodoItem1Predecessors and assigns it to the Predecessors field.
func (o *TasksIdJsonTodoItem1) SetPredecessors(v []TasksIdJsonTodoItem1Predecessors) {
	o.Predecessors = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *TasksIdJsonTodoItem1) SetPriority(v string) {
	o.Priority = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetPrivate() int32 {
	if o == nil || o.Private == nil {
		var ret int32
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetPrivateOk() (*int32, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given int32 and assigns it to the Private field.
func (o *TasksIdJsonTodoItem1) SetPrivate(v int32) {
	o.Private = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *TasksIdJsonTodoItem1) SetProgress(v string) {
	o.Progress = &v
}

// GetResponsiblePartyId returns the ResponsiblePartyId field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetResponsiblePartyId() int32 {
	if o == nil || o.ResponsiblePartyId == nil {
		var ret int32
		return ret
	}
	return *o.ResponsiblePartyId
}

// GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetResponsiblePartyIdOk() (*int32, bool) {
	if o == nil || o.ResponsiblePartyId == nil {
		return nil, false
	}
	return o.ResponsiblePartyId, true
}

// HasResponsiblePartyId returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasResponsiblePartyId() bool {
	if o != nil && o.ResponsiblePartyId != nil {
		return true
	}

	return false
}

// SetResponsiblePartyId gets a reference to the given int32 and assigns it to the ResponsiblePartyId field.
func (o *TasksIdJsonTodoItem1) SetResponsiblePartyId(v int32) {
	o.ResponsiblePartyId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TasksIdJsonTodoItem1) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TasksIdJsonTodoItem1) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksIdJsonTodoItem1) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TasksIdJsonTodoItem1) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *TasksIdJsonTodoItem1) SetTags(v string) {
	o.Tags = &v
}

func (o TasksIdJsonTodoItem1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.ChangeFollowerIds != nil {
		toSerialize["changeFollowerIds"] = o.ChangeFollowerIds
	}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.CommentFollowerIds != nil {
		toSerialize["commentFollowerIds"] = o.CommentFollowerIds
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DueDate != nil {
		toSerialize["due-date"] = o.DueDate
	}
	if o.EstimatedMinutes != nil {
		toSerialize["estimated-minutes"] = o.EstimatedMinutes
	}
	if o.GrantAccessTo != nil {
		toSerialize["grant-access-to"] = o.GrantAccessTo
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.ParentTaskId != nil {
		toSerialize["parentTaskId"] = o.ParentTaskId
	}
	if o.PendingFileAttachments != nil {
		toSerialize["pendingFileAttachments"] = o.PendingFileAttachments
	}
	if o.PositionAfterTask != nil {
		toSerialize["positionAfterTask"] = o.PositionAfterTask
	}
	if o.Predecessors != nil {
		toSerialize["predecessors"] = o.Predecessors
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.ResponsiblePartyId != nil {
		toSerialize["responsible-party-id"] = o.ResponsiblePartyId
	}
	if o.StartDate != nil {
		toSerialize["start-date"] = o.StartDate
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTasksIdJsonTodoItem1 struct {
	value *TasksIdJsonTodoItem1
	isSet bool
}

func (v NullableTasksIdJsonTodoItem1) Get() *TasksIdJsonTodoItem1 {
	return v.value
}

func (v *NullableTasksIdJsonTodoItem1) Set(val *TasksIdJsonTodoItem1) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksIdJsonTodoItem1) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksIdJsonTodoItem1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksIdJsonTodoItem1(val *TasksIdJsonTodoItem1) *NullableTasksIdJsonTodoItem1 {
	return &NullableTasksIdJsonTodoItem1{value: val, isSet: true}
}

func (v NullableTasksIdJsonTodoItem1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksIdJsonTodoItem1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


