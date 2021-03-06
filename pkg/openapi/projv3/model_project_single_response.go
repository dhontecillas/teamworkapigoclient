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

// ProjectSingleResponse SingleResponse contains information about a project.
type ProjectSingleResponse struct {
	Included *ProjectProjectsResponseIncluded `json:"included,omitempty"`
	Project *ViewProject `json:"project,omitempty"`
}

// NewProjectSingleResponse instantiates a new ProjectSingleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSingleResponse() *ProjectSingleResponse {
	this := ProjectSingleResponse{}
	return &this
}

// NewProjectSingleResponseWithDefaults instantiates a new ProjectSingleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSingleResponseWithDefaults() *ProjectSingleResponse {
	this := ProjectSingleResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ProjectSingleResponse) GetIncluded() ProjectProjectsResponseIncluded {
	if o == nil || o.Included == nil {
		var ret ProjectProjectsResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSingleResponse) GetIncludedOk() (*ProjectProjectsResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ProjectSingleResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given ProjectProjectsResponseIncluded and assigns it to the Included field.
func (o *ProjectSingleResponse) SetIncluded(v ProjectProjectsResponseIncluded) {
	o.Included = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectSingleResponse) GetProject() ViewProject {
	if o == nil || o.Project == nil {
		var ret ViewProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSingleResponse) GetProjectOk() (*ViewProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectSingleResponse) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewProject and assigns it to the Project field.
func (o *ProjectSingleResponse) SetProject(v ViewProject) {
	o.Project = &v
}

func (o ProjectSingleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProjectSingleResponse struct {
	value *ProjectSingleResponse
	isSet bool
}

func (v NullableProjectSingleResponse) Get() *ProjectSingleResponse {
	return v.value
}

func (v *NullableProjectSingleResponse) Set(val *ProjectSingleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSingleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSingleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSingleResponse(val *ProjectSingleResponse) *NullableProjectSingleResponse {
	return &NullableProjectSingleResponse{value: val, isSet: true}
}

func (v NullableProjectSingleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSingleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


