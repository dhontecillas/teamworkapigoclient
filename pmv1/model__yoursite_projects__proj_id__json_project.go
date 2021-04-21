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

// YoursiteProjectsProjIdJsonProject struct for YoursiteProjectsProjIdJsonProject
type YoursiteProjectsProjIdJsonProject struct {
	ProjectOwnerId *string `json:"projectOwnerId,omitempty"`
}

// NewYoursiteProjectsProjIdJsonProject instantiates a new YoursiteProjectsProjIdJsonProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYoursiteProjectsProjIdJsonProject() *YoursiteProjectsProjIdJsonProject {
	this := YoursiteProjectsProjIdJsonProject{}
	return &this
}

// NewYoursiteProjectsProjIdJsonProjectWithDefaults instantiates a new YoursiteProjectsProjIdJsonProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYoursiteProjectsProjIdJsonProjectWithDefaults() *YoursiteProjectsProjIdJsonProject {
	this := YoursiteProjectsProjIdJsonProject{}
	return &this
}

// GetProjectOwnerId returns the ProjectOwnerId field value if set, zero value otherwise.
func (o *YoursiteProjectsProjIdJsonProject) GetProjectOwnerId() string {
	if o == nil || o.ProjectOwnerId == nil {
		var ret string
		return ret
	}
	return *o.ProjectOwnerId
}

// GetProjectOwnerIdOk returns a tuple with the ProjectOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YoursiteProjectsProjIdJsonProject) GetProjectOwnerIdOk() (*string, bool) {
	if o == nil || o.ProjectOwnerId == nil {
		return nil, false
	}
	return o.ProjectOwnerId, true
}

// HasProjectOwnerId returns a boolean if a field has been set.
func (o *YoursiteProjectsProjIdJsonProject) HasProjectOwnerId() bool {
	if o != nil && o.ProjectOwnerId != nil {
		return true
	}

	return false
}

// SetProjectOwnerId gets a reference to the given string and assigns it to the ProjectOwnerId field.
func (o *YoursiteProjectsProjIdJsonProject) SetProjectOwnerId(v string) {
	o.ProjectOwnerId = &v
}

func (o YoursiteProjectsProjIdJsonProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectOwnerId != nil {
		toSerialize["projectOwnerId"] = o.ProjectOwnerId
	}
	return json.Marshal(toSerialize)
}

type NullableYoursiteProjectsProjIdJsonProject struct {
	value *YoursiteProjectsProjIdJsonProject
	isSet bool
}

func (v NullableYoursiteProjectsProjIdJsonProject) Get() *YoursiteProjectsProjIdJsonProject {
	return v.value
}

func (v *NullableYoursiteProjectsProjIdJsonProject) Set(val *YoursiteProjectsProjIdJsonProject) {
	v.value = val
	v.isSet = true
}

func (v NullableYoursiteProjectsProjIdJsonProject) IsSet() bool {
	return v.isSet
}

func (v *NullableYoursiteProjectsProjIdJsonProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYoursiteProjectsProjIdJsonProject(val *YoursiteProjectsProjIdJsonProject) *NullableYoursiteProjectsProjIdJsonProject {
	return &NullableYoursiteProjectsProjIdJsonProject{value: val, isSet: true}
}

func (v NullableYoursiteProjectsProjIdJsonProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYoursiteProjectsProjIdJsonProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


