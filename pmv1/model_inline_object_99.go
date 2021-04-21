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

// InlineObject99 struct for InlineObject99
type InlineObject99 struct {
	Content *string `json:"content,omitempty"`
	CreatorId *int32 `json:"creator-id,omitempty"`
	Notify *bool `json:"notify,omitempty"`
	Private *bool `json:"private,omitempty"`
	TasklistId *int32 `json:"tasklistId,omitempty"`
	TodoItem *ProjectsProjidTasksQuickaddJsonTodoItem `json:"todo-item,omitempty"`
}

// NewInlineObject99 instantiates a new InlineObject99 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject99() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// NewInlineObject99WithDefaults instantiates a new InlineObject99 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject99WithDefaults() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *InlineObject99) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *InlineObject99) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *InlineObject99) SetContent(v string) {
	o.Content = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *InlineObject99) GetCreatorId() int32 {
	if o == nil || o.CreatorId == nil {
		var ret int32
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetCreatorIdOk() (*int32, bool) {
	if o == nil || o.CreatorId == nil {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *InlineObject99) HasCreatorId() bool {
	if o != nil && o.CreatorId != nil {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given int32 and assigns it to the CreatorId field.
func (o *InlineObject99) SetCreatorId(v int32) {
	o.CreatorId = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *InlineObject99) GetNotify() bool {
	if o == nil || o.Notify == nil {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetNotifyOk() (*bool, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *InlineObject99) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *InlineObject99) SetNotify(v bool) {
	o.Notify = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *InlineObject99) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *InlineObject99) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *InlineObject99) SetPrivate(v bool) {
	o.Private = &v
}

// GetTasklistId returns the TasklistId field value if set, zero value otherwise.
func (o *InlineObject99) GetTasklistId() int32 {
	if o == nil || o.TasklistId == nil {
		var ret int32
		return ret
	}
	return *o.TasklistId
}

// GetTasklistIdOk returns a tuple with the TasklistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetTasklistIdOk() (*int32, bool) {
	if o == nil || o.TasklistId == nil {
		return nil, false
	}
	return o.TasklistId, true
}

// HasTasklistId returns a boolean if a field has been set.
func (o *InlineObject99) HasTasklistId() bool {
	if o != nil && o.TasklistId != nil {
		return true
	}

	return false
}

// SetTasklistId gets a reference to the given int32 and assigns it to the TasklistId field.
func (o *InlineObject99) SetTasklistId(v int32) {
	o.TasklistId = &v
}

// GetTodoItem returns the TodoItem field value if set, zero value otherwise.
func (o *InlineObject99) GetTodoItem() ProjectsProjidTasksQuickaddJsonTodoItem {
	if o == nil || o.TodoItem == nil {
		var ret ProjectsProjidTasksQuickaddJsonTodoItem
		return ret
	}
	return *o.TodoItem
}

// GetTodoItemOk returns a tuple with the TodoItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetTodoItemOk() (*ProjectsProjidTasksQuickaddJsonTodoItem, bool) {
	if o == nil || o.TodoItem == nil {
		return nil, false
	}
	return o.TodoItem, true
}

// HasTodoItem returns a boolean if a field has been set.
func (o *InlineObject99) HasTodoItem() bool {
	if o != nil && o.TodoItem != nil {
		return true
	}

	return false
}

// SetTodoItem gets a reference to the given ProjectsProjidTasksQuickaddJsonTodoItem and assigns it to the TodoItem field.
func (o *InlineObject99) SetTodoItem(v ProjectsProjidTasksQuickaddJsonTodoItem) {
	o.TodoItem = &v
}

func (o InlineObject99) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreatorId != nil {
		toSerialize["creator-id"] = o.CreatorId
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.TasklistId != nil {
		toSerialize["tasklistId"] = o.TasklistId
	}
	if o.TodoItem != nil {
		toSerialize["todo-item"] = o.TodoItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject99 struct {
	value *InlineObject99
	isSet bool
}

func (v NullableInlineObject99) Get() *InlineObject99 {
	return v.value
}

func (v *NullableInlineObject99) Set(val *InlineObject99) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject99) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject99) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject99(val *InlineObject99) *NullableInlineObject99 {
	return &NullableInlineObject99{value: val, isSet: true}
}

func (v NullableInlineObject99) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject99) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

