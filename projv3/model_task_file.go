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

// TaskFile File stores information about a uploaded file.
type TaskFile struct {
	CategoryId *int32 `json:"categoryId,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewTaskFile instantiates a new TaskFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskFile() *TaskFile {
	this := TaskFile{}
	return &this
}

// NewTaskFileWithDefaults instantiates a new TaskFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskFileWithDefaults() *TaskFile {
	this := TaskFile{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *TaskFile) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFile) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *TaskFile) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *TaskFile) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskFile) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFile) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaskFile) SetId(v int32) {
	o.Id = &v
}

func (o TaskFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTaskFile struct {
	value *TaskFile
	isSet bool
}

func (v NullableTaskFile) Get() *TaskFile {
	return v.value
}

func (v *NullableTaskFile) Set(val *TaskFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskFile(val *TaskFile) *NullableTaskFile {
	return &NullableTaskFile{value: val, isSet: true}
}

func (v NullableTaskFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


