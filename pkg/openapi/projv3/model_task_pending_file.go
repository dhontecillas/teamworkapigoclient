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

// TaskPendingFile PendingFile stores information about a file uploaded on-the-fly.
type TaskPendingFile struct {
	CategoryId *int32 `json:"categoryId,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

// NewTaskPendingFile instantiates a new TaskPendingFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPendingFile() *TaskPendingFile {
	this := TaskPendingFile{}
	return &this
}

// NewTaskPendingFileWithDefaults instantiates a new TaskPendingFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPendingFileWithDefaults() *TaskPendingFile {
	this := TaskPendingFile{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *TaskPendingFile) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPendingFile) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *TaskPendingFile) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *TaskPendingFile) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TaskPendingFile) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPendingFile) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TaskPendingFile) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TaskPendingFile) SetReference(v string) {
	o.Reference = &v
}

func (o TaskPendingFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableTaskPendingFile struct {
	value *TaskPendingFile
	isSet bool
}

func (v NullableTaskPendingFile) Get() *TaskPendingFile {
	return v.value
}

func (v *NullableTaskPendingFile) Set(val *TaskPendingFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPendingFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPendingFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPendingFile(val *TaskPendingFile) *NullableTaskPendingFile {
	return &NullableTaskPendingFile{value: val, isSet: true}
}

func (v NullableTaskPendingFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPendingFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


