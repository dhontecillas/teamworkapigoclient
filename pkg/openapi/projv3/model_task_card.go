/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"encoding/json"
)

// TaskCard Card stores information about the card created with the task.
type TaskCard struct {
	ColumnId *int32 `json:"columnId,omitempty"`
}

// NewTaskCard instantiates a new TaskCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCard() *TaskCard {
	this := TaskCard{}
	return &this
}

// NewTaskCardWithDefaults instantiates a new TaskCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCardWithDefaults() *TaskCard {
	this := TaskCard{}
	return &this
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *TaskCard) GetColumnId() int32 {
	if o == nil || o.ColumnId == nil {
		var ret int32
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCard) GetColumnIdOk() (*int32, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *TaskCard) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given int32 and assigns it to the ColumnId field.
func (o *TaskCard) SetColumnId(v int32) {
	o.ColumnId = &v
}

func (o TaskCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	return json.Marshal(toSerialize)
}

type NullableTaskCard struct {
	value *TaskCard
	isSet bool
}

func (v NullableTaskCard) Get() *TaskCard {
	return v.value
}

func (v *NullableTaskCard) Set(val *TaskCard) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCard) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCard(val *TaskCard) *NullableTaskCard {
	return &NullableTaskCard{value: val, isSet: true}
}

func (v NullableTaskCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


