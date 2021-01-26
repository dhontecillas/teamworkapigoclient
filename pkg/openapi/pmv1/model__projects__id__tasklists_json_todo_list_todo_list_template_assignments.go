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

// ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments struct for ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments
type ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments struct {
	Role1 *[]int32 `json:"Role1,omitempty"`
	Role2 *[]int32 `json:"Role2,omitempty"`
}

// NewProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments instantiates a new ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments() *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments {
	this := ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments{}
	return &this
}

// NewProjectsIdTasklistsJsonTodoListTodoListTemplateAssignmentsWithDefaults instantiates a new ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdTasklistsJsonTodoListTodoListTemplateAssignmentsWithDefaults() *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments {
	this := ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments{}
	return &this
}

// GetRole1 returns the Role1 field value if set, zero value otherwise.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) GetRole1() []int32 {
	if o == nil || o.Role1 == nil {
		var ret []int32
		return ret
	}
	return *o.Role1
}

// GetRole1Ok returns a tuple with the Role1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) GetRole1Ok() (*[]int32, bool) {
	if o == nil || o.Role1 == nil {
		return nil, false
	}
	return o.Role1, true
}

// HasRole1 returns a boolean if a field has been set.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) HasRole1() bool {
	if o != nil && o.Role1 != nil {
		return true
	}

	return false
}

// SetRole1 gets a reference to the given []int32 and assigns it to the Role1 field.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) SetRole1(v []int32) {
	o.Role1 = &v
}

// GetRole2 returns the Role2 field value if set, zero value otherwise.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) GetRole2() []int32 {
	if o == nil || o.Role2 == nil {
		var ret []int32
		return ret
	}
	return *o.Role2
}

// GetRole2Ok returns a tuple with the Role2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) GetRole2Ok() (*[]int32, bool) {
	if o == nil || o.Role2 == nil {
		return nil, false
	}
	return o.Role2, true
}

// HasRole2 returns a boolean if a field has been set.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) HasRole2() bool {
	if o != nil && o.Role2 != nil {
		return true
	}

	return false
}

// SetRole2 gets a reference to the given []int32 and assigns it to the Role2 field.
func (o *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) SetRole2(v []int32) {
	o.Role2 = &v
}

func (o ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role1 != nil {
		toSerialize["Role1"] = o.Role1
	}
	if o.Role2 != nil {
		toSerialize["Role2"] = o.Role2
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments struct {
	value *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments
	isSet bool
}

func (v NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) Get() *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments {
	return v.value
}

func (v *NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) Set(val *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments(val *ProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) *NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments {
	return &NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments{value: val, isSet: true}
}

func (v NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdTasklistsJsonTodoListTodoListTemplateAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


