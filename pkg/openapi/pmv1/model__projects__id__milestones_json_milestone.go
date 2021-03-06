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

// ProjectsIdMilestonesJsonMilestone struct for ProjectsIdMilestonesJsonMilestone
type ProjectsIdMilestonesJsonMilestone struct {
	Deadline *string `json:"deadline,omitempty"`
	Description *string `json:"description,omitempty"`
	Notify *bool `json:"notify,omitempty"`
	Reminder *bool `json:"reminder,omitempty"`
	ResponsiblePartyIds *string `json:"responsible-party-ids,omitempty"`
	Tags *string `json:"tags,omitempty"`
	TaskListIds *[]string `json:"taskListIds,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewProjectsIdMilestonesJsonMilestone instantiates a new ProjectsIdMilestonesJsonMilestone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdMilestonesJsonMilestone() *ProjectsIdMilestonesJsonMilestone {
	this := ProjectsIdMilestonesJsonMilestone{}
	return &this
}

// NewProjectsIdMilestonesJsonMilestoneWithDefaults instantiates a new ProjectsIdMilestonesJsonMilestone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdMilestonesJsonMilestoneWithDefaults() *ProjectsIdMilestonesJsonMilestone {
	this := ProjectsIdMilestonesJsonMilestone{}
	return &this
}

// GetDeadline returns the Deadline field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetDeadline() string {
	if o == nil || o.Deadline == nil {
		var ret string
		return ret
	}
	return *o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetDeadlineOk() (*string, bool) {
	if o == nil || o.Deadline == nil {
		return nil, false
	}
	return o.Deadline, true
}

// HasDeadline returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasDeadline() bool {
	if o != nil && o.Deadline != nil {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given string and assigns it to the Deadline field.
func (o *ProjectsIdMilestonesJsonMilestone) SetDeadline(v string) {
	o.Deadline = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsIdMilestonesJsonMilestone) SetDescription(v string) {
	o.Description = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetNotify() bool {
	if o == nil || o.Notify == nil {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetNotifyOk() (*bool, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *ProjectsIdMilestonesJsonMilestone) SetNotify(v bool) {
	o.Notify = &v
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetReminder() bool {
	if o == nil || o.Reminder == nil {
		var ret bool
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetReminderOk() (*bool, bool) {
	if o == nil || o.Reminder == nil {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasReminder() bool {
	if o != nil && o.Reminder != nil {
		return true
	}

	return false
}

// SetReminder gets a reference to the given bool and assigns it to the Reminder field.
func (o *ProjectsIdMilestonesJsonMilestone) SetReminder(v bool) {
	o.Reminder = &v
}

// GetResponsiblePartyIds returns the ResponsiblePartyIds field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetResponsiblePartyIds() string {
	if o == nil || o.ResponsiblePartyIds == nil {
		var ret string
		return ret
	}
	return *o.ResponsiblePartyIds
}

// GetResponsiblePartyIdsOk returns a tuple with the ResponsiblePartyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetResponsiblePartyIdsOk() (*string, bool) {
	if o == nil || o.ResponsiblePartyIds == nil {
		return nil, false
	}
	return o.ResponsiblePartyIds, true
}

// HasResponsiblePartyIds returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasResponsiblePartyIds() bool {
	if o != nil && o.ResponsiblePartyIds != nil {
		return true
	}

	return false
}

// SetResponsiblePartyIds gets a reference to the given string and assigns it to the ResponsiblePartyIds field.
func (o *ProjectsIdMilestonesJsonMilestone) SetResponsiblePartyIds(v string) {
	o.ResponsiblePartyIds = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProjectsIdMilestonesJsonMilestone) SetTags(v string) {
	o.Tags = &v
}

// GetTaskListIds returns the TaskListIds field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetTaskListIds() []string {
	if o == nil || o.TaskListIds == nil {
		var ret []string
		return ret
	}
	return *o.TaskListIds
}

// GetTaskListIdsOk returns a tuple with the TaskListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetTaskListIdsOk() (*[]string, bool) {
	if o == nil || o.TaskListIds == nil {
		return nil, false
	}
	return o.TaskListIds, true
}

// HasTaskListIds returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasTaskListIds() bool {
	if o != nil && o.TaskListIds != nil {
		return true
	}

	return false
}

// SetTaskListIds gets a reference to the given []string and assigns it to the TaskListIds field.
func (o *ProjectsIdMilestonesJsonMilestone) SetTaskListIds(v []string) {
	o.TaskListIds = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProjectsIdMilestonesJsonMilestone) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdMilestonesJsonMilestone) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProjectsIdMilestonesJsonMilestone) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProjectsIdMilestonesJsonMilestone) SetTitle(v string) {
	o.Title = &v
}

func (o ProjectsIdMilestonesJsonMilestone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deadline != nil {
		toSerialize["deadline"] = o.Deadline
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.Reminder != nil {
		toSerialize["reminder"] = o.Reminder
	}
	if o.ResponsiblePartyIds != nil {
		toSerialize["responsible-party-ids"] = o.ResponsiblePartyIds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TaskListIds != nil {
		toSerialize["taskListIds"] = o.TaskListIds
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdMilestonesJsonMilestone struct {
	value *ProjectsIdMilestonesJsonMilestone
	isSet bool
}

func (v NullableProjectsIdMilestonesJsonMilestone) Get() *ProjectsIdMilestonesJsonMilestone {
	return v.value
}

func (v *NullableProjectsIdMilestonesJsonMilestone) Set(val *ProjectsIdMilestonesJsonMilestone) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdMilestonesJsonMilestone) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdMilestonesJsonMilestone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdMilestonesJsonMilestone(val *ProjectsIdMilestonesJsonMilestone) *NullableProjectsIdMilestonesJsonMilestone {
	return &NullableProjectsIdMilestonesJsonMilestone{value: val, isSet: true}
}

func (v NullableProjectsIdMilestonesJsonMilestone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdMilestonesJsonMilestone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


