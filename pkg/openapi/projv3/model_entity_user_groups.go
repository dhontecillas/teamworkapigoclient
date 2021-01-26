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

// EntityUserGroups UserGroups are common lists for storing users, companies and teams ids together
type EntityUserGroups struct {
	CompanyIDs *[]int32 `json:"CompanyIDs,omitempty"`
	TeamIDs *[]int32 `json:"TeamIDs,omitempty"`
	UserIDs *[]int32 `json:"UserIDs,omitempty"`
}

// NewEntityUserGroups instantiates a new EntityUserGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityUserGroups() *EntityUserGroups {
	this := EntityUserGroups{}
	return &this
}

// NewEntityUserGroupsWithDefaults instantiates a new EntityUserGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityUserGroupsWithDefaults() *EntityUserGroups {
	this := EntityUserGroups{}
	return &this
}

// GetCompanyIDs returns the CompanyIDs field value if set, zero value otherwise.
func (o *EntityUserGroups) GetCompanyIDs() []int32 {
	if o == nil || o.CompanyIDs == nil {
		var ret []int32
		return ret
	}
	return *o.CompanyIDs
}

// GetCompanyIDsOk returns a tuple with the CompanyIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityUserGroups) GetCompanyIDsOk() (*[]int32, bool) {
	if o == nil || o.CompanyIDs == nil {
		return nil, false
	}
	return o.CompanyIDs, true
}

// HasCompanyIDs returns a boolean if a field has been set.
func (o *EntityUserGroups) HasCompanyIDs() bool {
	if o != nil && o.CompanyIDs != nil {
		return true
	}

	return false
}

// SetCompanyIDs gets a reference to the given []int32 and assigns it to the CompanyIDs field.
func (o *EntityUserGroups) SetCompanyIDs(v []int32) {
	o.CompanyIDs = &v
}

// GetTeamIDs returns the TeamIDs field value if set, zero value otherwise.
func (o *EntityUserGroups) GetTeamIDs() []int32 {
	if o == nil || o.TeamIDs == nil {
		var ret []int32
		return ret
	}
	return *o.TeamIDs
}

// GetTeamIDsOk returns a tuple with the TeamIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityUserGroups) GetTeamIDsOk() (*[]int32, bool) {
	if o == nil || o.TeamIDs == nil {
		return nil, false
	}
	return o.TeamIDs, true
}

// HasTeamIDs returns a boolean if a field has been set.
func (o *EntityUserGroups) HasTeamIDs() bool {
	if o != nil && o.TeamIDs != nil {
		return true
	}

	return false
}

// SetTeamIDs gets a reference to the given []int32 and assigns it to the TeamIDs field.
func (o *EntityUserGroups) SetTeamIDs(v []int32) {
	o.TeamIDs = &v
}

// GetUserIDs returns the UserIDs field value if set, zero value otherwise.
func (o *EntityUserGroups) GetUserIDs() []int32 {
	if o == nil || o.UserIDs == nil {
		var ret []int32
		return ret
	}
	return *o.UserIDs
}

// GetUserIDsOk returns a tuple with the UserIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityUserGroups) GetUserIDsOk() (*[]int32, bool) {
	if o == nil || o.UserIDs == nil {
		return nil, false
	}
	return o.UserIDs, true
}

// HasUserIDs returns a boolean if a field has been set.
func (o *EntityUserGroups) HasUserIDs() bool {
	if o != nil && o.UserIDs != nil {
		return true
	}

	return false
}

// SetUserIDs gets a reference to the given []int32 and assigns it to the UserIDs field.
func (o *EntityUserGroups) SetUserIDs(v []int32) {
	o.UserIDs = &v
}

func (o EntityUserGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyIDs != nil {
		toSerialize["CompanyIDs"] = o.CompanyIDs
	}
	if o.TeamIDs != nil {
		toSerialize["TeamIDs"] = o.TeamIDs
	}
	if o.UserIDs != nil {
		toSerialize["UserIDs"] = o.UserIDs
	}
	return json.Marshal(toSerialize)
}

type NullableEntityUserGroups struct {
	value *EntityUserGroups
	isSet bool
}

func (v NullableEntityUserGroups) Get() *EntityUserGroups {
	return v.value
}

func (v *NullableEntityUserGroups) Set(val *EntityUserGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityUserGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityUserGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityUserGroups(val *EntityUserGroups) *NullableEntityUserGroups {
	return &NullableEntityUserGroups{value: val, isSet: true}
}

func (v NullableEntityUserGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityUserGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


