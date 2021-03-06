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

// InlineObject90 struct for InlineObject90
type InlineObject90 struct {
	TodoList *TasklistsIdJsonTodoList `json:"todo-list,omitempty"`
}

// NewInlineObject90 instantiates a new InlineObject90 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject90() *InlineObject90 {
	this := InlineObject90{}
	return &this
}

// NewInlineObject90WithDefaults instantiates a new InlineObject90 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject90WithDefaults() *InlineObject90 {
	this := InlineObject90{}
	return &this
}

// GetTodoList returns the TodoList field value if set, zero value otherwise.
func (o *InlineObject90) GetTodoList() TasklistsIdJsonTodoList {
	if o == nil || o.TodoList == nil {
		var ret TasklistsIdJsonTodoList
		return ret
	}
	return *o.TodoList
}

// GetTodoListOk returns a tuple with the TodoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject90) GetTodoListOk() (*TasklistsIdJsonTodoList, bool) {
	if o == nil || o.TodoList == nil {
		return nil, false
	}
	return o.TodoList, true
}

// HasTodoList returns a boolean if a field has been set.
func (o *InlineObject90) HasTodoList() bool {
	if o != nil && o.TodoList != nil {
		return true
	}

	return false
}

// SetTodoList gets a reference to the given TasklistsIdJsonTodoList and assigns it to the TodoList field.
func (o *InlineObject90) SetTodoList(v TasklistsIdJsonTodoList) {
	o.TodoList = &v
}

func (o InlineObject90) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TodoList != nil {
		toSerialize["todo-list"] = o.TodoList
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject90 struct {
	value *InlineObject90
	isSet bool
}

func (v NullableInlineObject90) Get() *InlineObject90 {
	return v.value
}

func (v *NullableInlineObject90) Set(val *InlineObject90) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject90) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject90) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject90(val *InlineObject90) *NullableInlineObject90 {
	return &NullableInlineObject90{value: val, isSet: true}
}

func (v NullableInlineObject90) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject90) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


