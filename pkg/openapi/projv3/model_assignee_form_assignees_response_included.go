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

// AssigneeFormAssigneesResponseIncluded struct for AssigneeFormAssigneesResponseIncluded
type AssigneeFormAssigneesResponseIncluded struct {
	Companies *map[string]ViewCompany `json:"companies,omitempty"`
	Projects *map[string]ViewProject `json:"projects,omitempty"`
	Teams *map[string]ViewTeam `json:"teams,omitempty"`
	Users *map[string]ViewUser `json:"users,omitempty"`
}

// NewAssigneeFormAssigneesResponseIncluded instantiates a new AssigneeFormAssigneesResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssigneeFormAssigneesResponseIncluded() *AssigneeFormAssigneesResponseIncluded {
	this := AssigneeFormAssigneesResponseIncluded{}
	return &this
}

// NewAssigneeFormAssigneesResponseIncludedWithDefaults instantiates a new AssigneeFormAssigneesResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssigneeFormAssigneesResponseIncludedWithDefaults() *AssigneeFormAssigneesResponseIncluded {
	this := AssigneeFormAssigneesResponseIncluded{}
	return &this
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *AssigneeFormAssigneesResponseIncluded) GetCompanies() map[string]ViewCompany {
	if o == nil || o.Companies == nil {
		var ret map[string]ViewCompany
		return ret
	}
	return *o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeFormAssigneesResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool) {
	if o == nil || o.Companies == nil {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *AssigneeFormAssigneesResponseIncluded) HasCompanies() bool {
	if o != nil && o.Companies != nil {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given map[string]ViewCompany and assigns it to the Companies field.
func (o *AssigneeFormAssigneesResponseIncluded) SetCompanies(v map[string]ViewCompany) {
	o.Companies = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *AssigneeFormAssigneesResponseIncluded) GetProjects() map[string]ViewProject {
	if o == nil || o.Projects == nil {
		var ret map[string]ViewProject
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeFormAssigneesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *AssigneeFormAssigneesResponseIncluded) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given map[string]ViewProject and assigns it to the Projects field.
func (o *AssigneeFormAssigneesResponseIncluded) SetProjects(v map[string]ViewProject) {
	o.Projects = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *AssigneeFormAssigneesResponseIncluded) GetTeams() map[string]ViewTeam {
	if o == nil || o.Teams == nil {
		var ret map[string]ViewTeam
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeFormAssigneesResponseIncluded) GetTeamsOk() (*map[string]ViewTeam, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *AssigneeFormAssigneesResponseIncluded) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given map[string]ViewTeam and assigns it to the Teams field.
func (o *AssigneeFormAssigneesResponseIncluded) SetTeams(v map[string]ViewTeam) {
	o.Teams = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AssigneeFormAssigneesResponseIncluded) GetUsers() map[string]ViewUser {
	if o == nil || o.Users == nil {
		var ret map[string]ViewUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeFormAssigneesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AssigneeFormAssigneesResponseIncluded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]ViewUser and assigns it to the Users field.
func (o *AssigneeFormAssigneesResponseIncluded) SetUsers(v map[string]ViewUser) {
	o.Users = &v
}

func (o AssigneeFormAssigneesResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableAssigneeFormAssigneesResponseIncluded struct {
	value *AssigneeFormAssigneesResponseIncluded
	isSet bool
}

func (v NullableAssigneeFormAssigneesResponseIncluded) Get() *AssigneeFormAssigneesResponseIncluded {
	return v.value
}

func (v *NullableAssigneeFormAssigneesResponseIncluded) Set(val *AssigneeFormAssigneesResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableAssigneeFormAssigneesResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableAssigneeFormAssigneesResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssigneeFormAssigneesResponseIncluded(val *AssigneeFormAssigneesResponseIncluded) *NullableAssigneeFormAssigneesResponseIncluded {
	return &NullableAssigneeFormAssigneesResponseIncluded{value: val, isSet: true}
}

func (v NullableAssigneeFormAssigneesResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssigneeFormAssigneesResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


