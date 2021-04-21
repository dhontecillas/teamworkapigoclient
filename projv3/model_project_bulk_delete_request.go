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

// ProjectBulkDeleteRequest BulkDeleteRequest contains the ids of the customfields that should be removed.
type ProjectBulkDeleteRequest struct {
	CustomfieldProjectIds *[]int32 `json:"customfieldProjectIds,omitempty"`
}

// NewProjectBulkDeleteRequest instantiates a new ProjectBulkDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBulkDeleteRequest() *ProjectBulkDeleteRequest {
	this := ProjectBulkDeleteRequest{}
	return &this
}

// NewProjectBulkDeleteRequestWithDefaults instantiates a new ProjectBulkDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBulkDeleteRequestWithDefaults() *ProjectBulkDeleteRequest {
	this := ProjectBulkDeleteRequest{}
	return &this
}

// GetCustomfieldProjectIds returns the CustomfieldProjectIds field value if set, zero value otherwise.
func (o *ProjectBulkDeleteRequest) GetCustomfieldProjectIds() []int32 {
	if o == nil || o.CustomfieldProjectIds == nil {
		var ret []int32
		return ret
	}
	return *o.CustomfieldProjectIds
}

// GetCustomfieldProjectIdsOk returns a tuple with the CustomfieldProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBulkDeleteRequest) GetCustomfieldProjectIdsOk() (*[]int32, bool) {
	if o == nil || o.CustomfieldProjectIds == nil {
		return nil, false
	}
	return o.CustomfieldProjectIds, true
}

// HasCustomfieldProjectIds returns a boolean if a field has been set.
func (o *ProjectBulkDeleteRequest) HasCustomfieldProjectIds() bool {
	if o != nil && o.CustomfieldProjectIds != nil {
		return true
	}

	return false
}

// SetCustomfieldProjectIds gets a reference to the given []int32 and assigns it to the CustomfieldProjectIds field.
func (o *ProjectBulkDeleteRequest) SetCustomfieldProjectIds(v []int32) {
	o.CustomfieldProjectIds = &v
}

func (o ProjectBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomfieldProjectIds != nil {
		toSerialize["customfieldProjectIds"] = o.CustomfieldProjectIds
	}
	return json.Marshal(toSerialize)
}

type NullableProjectBulkDeleteRequest struct {
	value *ProjectBulkDeleteRequest
	isSet bool
}

func (v NullableProjectBulkDeleteRequest) Get() *ProjectBulkDeleteRequest {
	return v.value
}

func (v *NullableProjectBulkDeleteRequest) Set(val *ProjectBulkDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBulkDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBulkDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBulkDeleteRequest(val *ProjectBulkDeleteRequest) *NullableProjectBulkDeleteRequest {
	return &NullableProjectBulkDeleteRequest{value: val, isSet: true}
}

func (v NullableProjectBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBulkDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


