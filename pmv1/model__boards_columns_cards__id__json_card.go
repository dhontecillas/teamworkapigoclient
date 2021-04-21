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

// BoardsColumnsCardsIdJsonCard struct for BoardsColumnsCardsIdJsonCard
type BoardsColumnsCardsIdJsonCard struct {
	Content *string `json:"content,omitempty"`
	DueDate *string `json:"due-date,omitempty"`
	EstimatedMinutes *int32 `json:"estimated-minutes,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Progress *string `json:"progress,omitempty"`
	ResponsiblePartyId *string `json:"responsible-party-id,omitempty"`
	StartDate *string `json:"start-date,omitempty"`
	Tags *string `json:"tags,omitempty"`
	TaskListId *string `json:"taskListId,omitempty"`
}

// NewBoardsColumnsCardsIdJsonCard instantiates a new BoardsColumnsCardsIdJsonCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardsColumnsCardsIdJsonCard() *BoardsColumnsCardsIdJsonCard {
	this := BoardsColumnsCardsIdJsonCard{}
	return &this
}

// NewBoardsColumnsCardsIdJsonCardWithDefaults instantiates a new BoardsColumnsCardsIdJsonCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardsColumnsCardsIdJsonCardWithDefaults() *BoardsColumnsCardsIdJsonCard {
	this := BoardsColumnsCardsIdJsonCard{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *BoardsColumnsCardsIdJsonCard) SetContent(v string) {
	o.Content = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *BoardsColumnsCardsIdJsonCard) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEstimatedMinutes returns the EstimatedMinutes field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetEstimatedMinutes() int32 {
	if o == nil || o.EstimatedMinutes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMinutes
}

// GetEstimatedMinutesOk returns a tuple with the EstimatedMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetEstimatedMinutesOk() (*int32, bool) {
	if o == nil || o.EstimatedMinutes == nil {
		return nil, false
	}
	return o.EstimatedMinutes, true
}

// HasEstimatedMinutes returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasEstimatedMinutes() bool {
	if o != nil && o.EstimatedMinutes != nil {
		return true
	}

	return false
}

// SetEstimatedMinutes gets a reference to the given int32 and assigns it to the EstimatedMinutes field.
func (o *BoardsColumnsCardsIdJsonCard) SetEstimatedMinutes(v int32) {
	o.EstimatedMinutes = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *BoardsColumnsCardsIdJsonCard) SetPriority(v string) {
	o.Priority = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *BoardsColumnsCardsIdJsonCard) SetProgress(v string) {
	o.Progress = &v
}

// GetResponsiblePartyId returns the ResponsiblePartyId field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetResponsiblePartyId() string {
	if o == nil || o.ResponsiblePartyId == nil {
		var ret string
		return ret
	}
	return *o.ResponsiblePartyId
}

// GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetResponsiblePartyIdOk() (*string, bool) {
	if o == nil || o.ResponsiblePartyId == nil {
		return nil, false
	}
	return o.ResponsiblePartyId, true
}

// HasResponsiblePartyId returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasResponsiblePartyId() bool {
	if o != nil && o.ResponsiblePartyId != nil {
		return true
	}

	return false
}

// SetResponsiblePartyId gets a reference to the given string and assigns it to the ResponsiblePartyId field.
func (o *BoardsColumnsCardsIdJsonCard) SetResponsiblePartyId(v string) {
	o.ResponsiblePartyId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *BoardsColumnsCardsIdJsonCard) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *BoardsColumnsCardsIdJsonCard) SetTags(v string) {
	o.Tags = &v
}

// GetTaskListId returns the TaskListId field value if set, zero value otherwise.
func (o *BoardsColumnsCardsIdJsonCard) GetTaskListId() string {
	if o == nil || o.TaskListId == nil {
		var ret string
		return ret
	}
	return *o.TaskListId
}

// GetTaskListIdOk returns a tuple with the TaskListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardsColumnsCardsIdJsonCard) GetTaskListIdOk() (*string, bool) {
	if o == nil || o.TaskListId == nil {
		return nil, false
	}
	return o.TaskListId, true
}

// HasTaskListId returns a boolean if a field has been set.
func (o *BoardsColumnsCardsIdJsonCard) HasTaskListId() bool {
	if o != nil && o.TaskListId != nil {
		return true
	}

	return false
}

// SetTaskListId gets a reference to the given string and assigns it to the TaskListId field.
func (o *BoardsColumnsCardsIdJsonCard) SetTaskListId(v string) {
	o.TaskListId = &v
}

func (o BoardsColumnsCardsIdJsonCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.DueDate != nil {
		toSerialize["due-date"] = o.DueDate
	}
	if o.EstimatedMinutes != nil {
		toSerialize["estimated-minutes"] = o.EstimatedMinutes
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
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
	if o.TaskListId != nil {
		toSerialize["taskListId"] = o.TaskListId
	}
	return json.Marshal(toSerialize)
}

type NullableBoardsColumnsCardsIdJsonCard struct {
	value *BoardsColumnsCardsIdJsonCard
	isSet bool
}

func (v NullableBoardsColumnsCardsIdJsonCard) Get() *BoardsColumnsCardsIdJsonCard {
	return v.value
}

func (v *NullableBoardsColumnsCardsIdJsonCard) Set(val *BoardsColumnsCardsIdJsonCard) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardsColumnsCardsIdJsonCard) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardsColumnsCardsIdJsonCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardsColumnsCardsIdJsonCard(val *BoardsColumnsCardsIdJsonCard) *NullableBoardsColumnsCardsIdJsonCard {
	return &NullableBoardsColumnsCardsIdJsonCard{value: val, isSet: true}
}

func (v NullableBoardsColumnsCardsIdJsonCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardsColumnsCardsIdJsonCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

