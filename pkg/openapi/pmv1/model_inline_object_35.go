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

// InlineObject35 struct for InlineObject35
type InlineObject35 struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewInlineObject35 instantiates a new InlineObject35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject35() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// NewInlineObject35WithDefaults instantiates a new InlineObject35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject35WithDefaults() *InlineObject35 {
	this := InlineObject35{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineObject35) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject35) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineObject35) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *InlineObject35) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o InlineObject35) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject35 struct {
	value *InlineObject35
	isSet bool
}

func (v NullableInlineObject35) Get() *InlineObject35 {
	return v.value
}

func (v *NullableInlineObject35) Set(val *InlineObject35) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject35) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject35(val *InlineObject35) *NullableInlineObject35 {
	return &NullableInlineObject35{value: val, isSet: true}
}

func (v NullableInlineObject35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


