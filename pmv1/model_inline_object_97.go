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

// InlineObject97 struct for InlineObject97
type InlineObject97 struct {
	TaskEstimatedMinutes *int32 `json:"taskEstimatedMinutes,omitempty"`
	TaskId *int32 `json:"taskId,omitempty"`
}

// NewInlineObject97 instantiates a new InlineObject97 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject97() *InlineObject97 {
	this := InlineObject97{}
	return &this
}

// NewInlineObject97WithDefaults instantiates a new InlineObject97 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject97WithDefaults() *InlineObject97 {
	this := InlineObject97{}
	return &this
}

// GetTaskEstimatedMinutes returns the TaskEstimatedMinutes field value if set, zero value otherwise.
func (o *InlineObject97) GetTaskEstimatedMinutes() int32 {
	if o == nil || o.TaskEstimatedMinutes == nil {
		var ret int32
		return ret
	}
	return *o.TaskEstimatedMinutes
}

// GetTaskEstimatedMinutesOk returns a tuple with the TaskEstimatedMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetTaskEstimatedMinutesOk() (*int32, bool) {
	if o == nil || o.TaskEstimatedMinutes == nil {
		return nil, false
	}
	return o.TaskEstimatedMinutes, true
}

// HasTaskEstimatedMinutes returns a boolean if a field has been set.
func (o *InlineObject97) HasTaskEstimatedMinutes() bool {
	if o != nil && o.TaskEstimatedMinutes != nil {
		return true
	}

	return false
}

// SetTaskEstimatedMinutes gets a reference to the given int32 and assigns it to the TaskEstimatedMinutes field.
func (o *InlineObject97) SetTaskEstimatedMinutes(v int32) {
	o.TaskEstimatedMinutes = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *InlineObject97) GetTaskId() int32 {
	if o == nil || o.TaskId == nil {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetTaskIdOk() (*int32, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *InlineObject97) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *InlineObject97) SetTaskId(v int32) {
	o.TaskId = &v
}

func (o InlineObject97) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskEstimatedMinutes != nil {
		toSerialize["taskEstimatedMinutes"] = o.TaskEstimatedMinutes
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject97 struct {
	value *InlineObject97
	isSet bool
}

func (v NullableInlineObject97) Get() *InlineObject97 {
	return v.value
}

func (v *NullableInlineObject97) Set(val *InlineObject97) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject97) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject97) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject97(val *InlineObject97) *NullableInlineObject97 {
	return &NullableInlineObject97{value: val, isSet: true}
}

func (v NullableInlineObject97) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject97) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


