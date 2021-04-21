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

// InlineObject98 struct for InlineObject98
type InlineObject98 struct {
	Task *TasksIdFilesJsonTask `json:"task,omitempty"`
}

// NewInlineObject98 instantiates a new InlineObject98 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject98() *InlineObject98 {
	this := InlineObject98{}
	return &this
}

// NewInlineObject98WithDefaults instantiates a new InlineObject98 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject98WithDefaults() *InlineObject98 {
	this := InlineObject98{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *InlineObject98) GetTask() TasksIdFilesJsonTask {
	if o == nil || o.Task == nil {
		var ret TasksIdFilesJsonTask
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject98) GetTaskOk() (*TasksIdFilesJsonTask, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *InlineObject98) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given TasksIdFilesJsonTask and assigns it to the Task field.
func (o *InlineObject98) SetTask(v TasksIdFilesJsonTask) {
	o.Task = &v
}

func (o InlineObject98) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject98 struct {
	value *InlineObject98
	isSet bool
}

func (v NullableInlineObject98) Get() *InlineObject98 {
	return v.value
}

func (v *NullableInlineObject98) Set(val *InlineObject98) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject98) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject98) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject98(val *InlineObject98) *NullableInlineObject98 {
	return &NullableInlineObject98{value: val, isSet: true}
}

func (v NullableInlineObject98) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject98) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


