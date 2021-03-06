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

// DashboardUserDashboardsResponseIncluded struct for DashboardUserDashboardsResponseIncluded
type DashboardUserDashboardsResponseIncluded struct {
	DashboardPanelSettings *map[string]ViewUserDashboardPanelSetting `json:"dashboardPanelSettings,omitempty"`
	DashboardPanels *map[string]ViewUserDashboardPanel `json:"dashboardPanels,omitempty"`
	DashboardSettings *map[string]ViewUserDashboardSetting `json:"dashboardSettings,omitempty"`
	Projects *map[string]ViewProject `json:"projects,omitempty"`
	Users *map[string]ViewUser `json:"users,omitempty"`
}

// NewDashboardUserDashboardsResponseIncluded instantiates a new DashboardUserDashboardsResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardUserDashboardsResponseIncluded() *DashboardUserDashboardsResponseIncluded {
	this := DashboardUserDashboardsResponseIncluded{}
	return &this
}

// NewDashboardUserDashboardsResponseIncludedWithDefaults instantiates a new DashboardUserDashboardsResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardUserDashboardsResponseIncludedWithDefaults() *DashboardUserDashboardsResponseIncluded {
	this := DashboardUserDashboardsResponseIncluded{}
	return &this
}

// GetDashboardPanelSettings returns the DashboardPanelSettings field value if set, zero value otherwise.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardPanelSettings() map[string]ViewUserDashboardPanelSetting {
	if o == nil || o.DashboardPanelSettings == nil {
		var ret map[string]ViewUserDashboardPanelSetting
		return ret
	}
	return *o.DashboardPanelSettings
}

// GetDashboardPanelSettingsOk returns a tuple with the DashboardPanelSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardPanelSettingsOk() (*map[string]ViewUserDashboardPanelSetting, bool) {
	if o == nil || o.DashboardPanelSettings == nil {
		return nil, false
	}
	return o.DashboardPanelSettings, true
}

// HasDashboardPanelSettings returns a boolean if a field has been set.
func (o *DashboardUserDashboardsResponseIncluded) HasDashboardPanelSettings() bool {
	if o != nil && o.DashboardPanelSettings != nil {
		return true
	}

	return false
}

// SetDashboardPanelSettings gets a reference to the given map[string]ViewUserDashboardPanelSetting and assigns it to the DashboardPanelSettings field.
func (o *DashboardUserDashboardsResponseIncluded) SetDashboardPanelSettings(v map[string]ViewUserDashboardPanelSetting) {
	o.DashboardPanelSettings = &v
}

// GetDashboardPanels returns the DashboardPanels field value if set, zero value otherwise.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardPanels() map[string]ViewUserDashboardPanel {
	if o == nil || o.DashboardPanels == nil {
		var ret map[string]ViewUserDashboardPanel
		return ret
	}
	return *o.DashboardPanels
}

// GetDashboardPanelsOk returns a tuple with the DashboardPanels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardPanelsOk() (*map[string]ViewUserDashboardPanel, bool) {
	if o == nil || o.DashboardPanels == nil {
		return nil, false
	}
	return o.DashboardPanels, true
}

// HasDashboardPanels returns a boolean if a field has been set.
func (o *DashboardUserDashboardsResponseIncluded) HasDashboardPanels() bool {
	if o != nil && o.DashboardPanels != nil {
		return true
	}

	return false
}

// SetDashboardPanels gets a reference to the given map[string]ViewUserDashboardPanel and assigns it to the DashboardPanels field.
func (o *DashboardUserDashboardsResponseIncluded) SetDashboardPanels(v map[string]ViewUserDashboardPanel) {
	o.DashboardPanels = &v
}

// GetDashboardSettings returns the DashboardSettings field value if set, zero value otherwise.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardSettings() map[string]ViewUserDashboardSetting {
	if o == nil || o.DashboardSettings == nil {
		var ret map[string]ViewUserDashboardSetting
		return ret
	}
	return *o.DashboardSettings
}

// GetDashboardSettingsOk returns a tuple with the DashboardSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserDashboardsResponseIncluded) GetDashboardSettingsOk() (*map[string]ViewUserDashboardSetting, bool) {
	if o == nil || o.DashboardSettings == nil {
		return nil, false
	}
	return o.DashboardSettings, true
}

// HasDashboardSettings returns a boolean if a field has been set.
func (o *DashboardUserDashboardsResponseIncluded) HasDashboardSettings() bool {
	if o != nil && o.DashboardSettings != nil {
		return true
	}

	return false
}

// SetDashboardSettings gets a reference to the given map[string]ViewUserDashboardSetting and assigns it to the DashboardSettings field.
func (o *DashboardUserDashboardsResponseIncluded) SetDashboardSettings(v map[string]ViewUserDashboardSetting) {
	o.DashboardSettings = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *DashboardUserDashboardsResponseIncluded) GetProjects() map[string]ViewProject {
	if o == nil || o.Projects == nil {
		var ret map[string]ViewProject
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserDashboardsResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *DashboardUserDashboardsResponseIncluded) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given map[string]ViewProject and assigns it to the Projects field.
func (o *DashboardUserDashboardsResponseIncluded) SetProjects(v map[string]ViewProject) {
	o.Projects = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *DashboardUserDashboardsResponseIncluded) GetUsers() map[string]ViewUser {
	if o == nil || o.Users == nil {
		var ret map[string]ViewUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserDashboardsResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *DashboardUserDashboardsResponseIncluded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]ViewUser and assigns it to the Users field.
func (o *DashboardUserDashboardsResponseIncluded) SetUsers(v map[string]ViewUser) {
	o.Users = &v
}

func (o DashboardUserDashboardsResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DashboardPanelSettings != nil {
		toSerialize["dashboardPanelSettings"] = o.DashboardPanelSettings
	}
	if o.DashboardPanels != nil {
		toSerialize["dashboardPanels"] = o.DashboardPanels
	}
	if o.DashboardSettings != nil {
		toSerialize["dashboardSettings"] = o.DashboardSettings
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardUserDashboardsResponseIncluded struct {
	value *DashboardUserDashboardsResponseIncluded
	isSet bool
}

func (v NullableDashboardUserDashboardsResponseIncluded) Get() *DashboardUserDashboardsResponseIncluded {
	return v.value
}

func (v *NullableDashboardUserDashboardsResponseIncluded) Set(val *DashboardUserDashboardsResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUserDashboardsResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUserDashboardsResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUserDashboardsResponseIncluded(val *DashboardUserDashboardsResponseIncluded) *NullableDashboardUserDashboardsResponseIncluded {
	return &NullableDashboardUserDashboardsResponseIncluded{value: val, isSet: true}
}

func (v NullableDashboardUserDashboardsResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUserDashboardsResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


