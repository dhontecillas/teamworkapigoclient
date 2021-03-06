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

// ProjectProjectsResponse ProjectsResponse contains information about a group of projects.
type ProjectProjectsResponse struct {
	Included *ProjectProjectsResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	Projects *[]ViewProject `json:"projects,omitempty"`
}

// NewProjectProjectsResponse instantiates a new ProjectProjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectsResponse() *ProjectProjectsResponse {
	this := ProjectProjectsResponse{}
	return &this
}

// NewProjectProjectsResponseWithDefaults instantiates a new ProjectProjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectsResponseWithDefaults() *ProjectProjectsResponse {
	this := ProjectProjectsResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ProjectProjectsResponse) GetIncluded() ProjectProjectsResponseIncluded {
	if o == nil || o.Included == nil {
		var ret ProjectProjectsResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectsResponse) GetIncludedOk() (*ProjectProjectsResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ProjectProjectsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given ProjectProjectsResponseIncluded and assigns it to the Included field.
func (o *ProjectProjectsResponse) SetIncluded(v ProjectProjectsResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProjectProjectsResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectsResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProjectProjectsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *ProjectProjectsResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ProjectProjectsResponse) GetProjects() []ViewProject {
	if o == nil || o.Projects == nil {
		var ret []ViewProject
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectsResponse) GetProjectsOk() (*[]ViewProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ProjectProjectsResponse) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ViewProject and assigns it to the Projects field.
func (o *ProjectProjectsResponse) SetProjects(v []ViewProject) {
	o.Projects = &v
}

func (o ProjectProjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableProjectProjectsResponse struct {
	value *ProjectProjectsResponse
	isSet bool
}

func (v NullableProjectProjectsResponse) Get() *ProjectProjectsResponse {
	return v.value
}

func (v *NullableProjectProjectsResponse) Set(val *ProjectProjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectsResponse(val *ProjectProjectsResponse) *NullableProjectProjectsResponse {
	return &NullableProjectProjectsResponse{value: val, isSet: true}
}

func (v NullableProjectProjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


