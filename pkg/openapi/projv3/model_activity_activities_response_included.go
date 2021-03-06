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

// ActivityActivitiesResponseIncluded struct for ActivityActivitiesResponseIncluded
type ActivityActivitiesResponseIncluded struct {
	Companies *map[string]ViewCompany `json:"companies,omitempty"`
	Projects *map[string]ViewProject `json:"projects,omitempty"`
	Users *map[string]ViewUser `json:"users,omitempty"`
}

// NewActivityActivitiesResponseIncluded instantiates a new ActivityActivitiesResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityActivitiesResponseIncluded() *ActivityActivitiesResponseIncluded {
	this := ActivityActivitiesResponseIncluded{}
	return &this
}

// NewActivityActivitiesResponseIncludedWithDefaults instantiates a new ActivityActivitiesResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityActivitiesResponseIncludedWithDefaults() *ActivityActivitiesResponseIncluded {
	this := ActivityActivitiesResponseIncluded{}
	return &this
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *ActivityActivitiesResponseIncluded) GetCompanies() map[string]ViewCompany {
	if o == nil || o.Companies == nil {
		var ret map[string]ViewCompany
		return ret
	}
	return *o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivitiesResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool) {
	if o == nil || o.Companies == nil {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *ActivityActivitiesResponseIncluded) HasCompanies() bool {
	if o != nil && o.Companies != nil {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given map[string]ViewCompany and assigns it to the Companies field.
func (o *ActivityActivitiesResponseIncluded) SetCompanies(v map[string]ViewCompany) {
	o.Companies = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ActivityActivitiesResponseIncluded) GetProjects() map[string]ViewProject {
	if o == nil || o.Projects == nil {
		var ret map[string]ViewProject
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivitiesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ActivityActivitiesResponseIncluded) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given map[string]ViewProject and assigns it to the Projects field.
func (o *ActivityActivitiesResponseIncluded) SetProjects(v map[string]ViewProject) {
	o.Projects = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ActivityActivitiesResponseIncluded) GetUsers() map[string]ViewUser {
	if o == nil || o.Users == nil {
		var ret map[string]ViewUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivitiesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ActivityActivitiesResponseIncluded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]ViewUser and assigns it to the Users field.
func (o *ActivityActivitiesResponseIncluded) SetUsers(v map[string]ViewUser) {
	o.Users = &v
}

func (o ActivityActivitiesResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableActivityActivitiesResponseIncluded struct {
	value *ActivityActivitiesResponseIncluded
	isSet bool
}

func (v NullableActivityActivitiesResponseIncluded) Get() *ActivityActivitiesResponseIncluded {
	return v.value
}

func (v *NullableActivityActivitiesResponseIncluded) Set(val *ActivityActivitiesResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityActivitiesResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityActivitiesResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityActivitiesResponseIncluded(val *ActivityActivitiesResponseIncluded) *NullableActivityActivitiesResponseIncluded {
	return &NullableActivityActivitiesResponseIncluded{value: val, isSet: true}
}

func (v NullableActivityActivitiesResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityActivitiesResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


