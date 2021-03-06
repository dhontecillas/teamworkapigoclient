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

// InlineObject102 struct for InlineObject102
type InlineObject102 struct {
	TodoItem *TasksTaskIdJsonTodoItem `json:"todo-item,omitempty"`
}

// NewInlineObject102 instantiates a new InlineObject102 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject102() *InlineObject102 {
	this := InlineObject102{}
	return &this
}

// NewInlineObject102WithDefaults instantiates a new InlineObject102 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject102WithDefaults() *InlineObject102 {
	this := InlineObject102{}
	return &this
}

// GetTodoItem returns the TodoItem field value if set, zero value otherwise.
func (o *InlineObject102) GetTodoItem() TasksTaskIdJsonTodoItem {
	if o == nil || o.TodoItem == nil {
		var ret TasksTaskIdJsonTodoItem
		return ret
	}
	return *o.TodoItem
}

// GetTodoItemOk returns a tuple with the TodoItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject102) GetTodoItemOk() (*TasksTaskIdJsonTodoItem, bool) {
	if o == nil || o.TodoItem == nil {
		return nil, false
	}
	return o.TodoItem, true
}

// HasTodoItem returns a boolean if a field has been set.
func (o *InlineObject102) HasTodoItem() bool {
	if o != nil && o.TodoItem != nil {
		return true
	}

	return false
}

// SetTodoItem gets a reference to the given TasksTaskIdJsonTodoItem and assigns it to the TodoItem field.
func (o *InlineObject102) SetTodoItem(v TasksTaskIdJsonTodoItem) {
	o.TodoItem = &v
}

func (o InlineObject102) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TodoItem != nil {
		toSerialize["todo-item"] = o.TodoItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject102 struct {
	value *InlineObject102
	isSet bool
}

func (v NullableInlineObject102) Get() *InlineObject102 {
	return v.value
}

func (v *NullableInlineObject102) Set(val *InlineObject102) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject102) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject102) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject102(val *InlineObject102) *NullableInlineObject102 {
	return &NullableInlineObject102{value: val, isSet: true}
}

func (v NullableInlineObject102) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject102) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


