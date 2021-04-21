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

// TaskTasksResponse TasksResponse contains information about a group of tasks.
type TaskTasksResponse struct {
	Included *TaskResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	Tasks *[]ViewTask `json:"tasks,omitempty"`
}

// NewTaskTasksResponse instantiates a new TaskTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTasksResponse() *TaskTasksResponse {
	this := TaskTasksResponse{}
	return &this
}

// NewTaskTasksResponseWithDefaults instantiates a new TaskTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTasksResponseWithDefaults() *TaskTasksResponse {
	this := TaskTasksResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TaskTasksResponse) GetIncluded() TaskResponseIncluded {
	if o == nil || o.Included == nil {
		var ret TaskResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTasksResponse) GetIncludedOk() (*TaskResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TaskTasksResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given TaskResponseIncluded and assigns it to the Included field.
func (o *TaskTasksResponse) SetIncluded(v TaskResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TaskTasksResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTasksResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TaskTasksResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *TaskTasksResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *TaskTasksResponse) GetTasks() []ViewTask {
	if o == nil || o.Tasks == nil {
		var ret []ViewTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTasksResponse) GetTasksOk() (*[]ViewTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *TaskTasksResponse) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ViewTask and assigns it to the Tasks field.
func (o *TaskTasksResponse) SetTasks(v []ViewTask) {
	o.Tasks = &v
}

func (o TaskTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableTaskTasksResponse struct {
	value *TaskTasksResponse
	isSet bool
}

func (v NullableTaskTasksResponse) Get() *TaskTasksResponse {
	return v.value
}

func (v *NullableTaskTasksResponse) Set(val *TaskTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTasksResponse(val *TaskTasksResponse) *NullableTaskTasksResponse {
	return &NullableTaskTasksResponse{value: val, isSet: true}
}

func (v NullableTaskTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


