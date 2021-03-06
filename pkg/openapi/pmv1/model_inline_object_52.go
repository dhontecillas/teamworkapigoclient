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

// InlineObject52 struct for InlineObject52
type InlineObject52 struct {
	Project ProjectsTemplateTemplateIDJsonProject `json:"project"`
}

// NewInlineObject52 instantiates a new InlineObject52 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject52(project ProjectsTemplateTemplateIDJsonProject) *InlineObject52 {
	this := InlineObject52{}
	this.Project = project
	return &this
}

// NewInlineObject52WithDefaults instantiates a new InlineObject52 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject52WithDefaults() *InlineObject52 {
	this := InlineObject52{}
	return &this
}

// GetProject returns the Project field value
func (o *InlineObject52) GetProject() ProjectsTemplateTemplateIDJsonProject {
	if o == nil {
		var ret ProjectsTemplateTemplateIDJsonProject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetProjectOk() (*ProjectsTemplateTemplateIDJsonProject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *InlineObject52) SetProject(v ProjectsTemplateTemplateIDJsonProject) {
	o.Project = v
}

func (o InlineObject52) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject52 struct {
	value *InlineObject52
	isSet bool
}

func (v NullableInlineObject52) Get() *InlineObject52 {
	return v.value
}

func (v *NullableInlineObject52) Set(val *InlineObject52) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject52) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject52) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject52(val *InlineObject52) *NullableInlineObject52 {
	return &NullableInlineObject52{value: val, isSet: true}
}

func (v NullableInlineObject52) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject52) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


