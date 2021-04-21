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

// AppAppsResponseIncluded struct for AppAppsResponseIncluded
type AppAppsResponseIncluded struct {
	Appsettings *map[string]ViewAppSetting `json:"appsettings,omitempty"`
}

// NewAppAppsResponseIncluded instantiates a new AppAppsResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAppsResponseIncluded() *AppAppsResponseIncluded {
	this := AppAppsResponseIncluded{}
	return &this
}

// NewAppAppsResponseIncludedWithDefaults instantiates a new AppAppsResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAppsResponseIncludedWithDefaults() *AppAppsResponseIncluded {
	this := AppAppsResponseIncluded{}
	return &this
}

// GetAppsettings returns the Appsettings field value if set, zero value otherwise.
func (o *AppAppsResponseIncluded) GetAppsettings() map[string]ViewAppSetting {
	if o == nil || o.Appsettings == nil {
		var ret map[string]ViewAppSetting
		return ret
	}
	return *o.Appsettings
}

// GetAppsettingsOk returns a tuple with the Appsettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAppsResponseIncluded) GetAppsettingsOk() (*map[string]ViewAppSetting, bool) {
	if o == nil || o.Appsettings == nil {
		return nil, false
	}
	return o.Appsettings, true
}

// HasAppsettings returns a boolean if a field has been set.
func (o *AppAppsResponseIncluded) HasAppsettings() bool {
	if o != nil && o.Appsettings != nil {
		return true
	}

	return false
}

// SetAppsettings gets a reference to the given map[string]ViewAppSetting and assigns it to the Appsettings field.
func (o *AppAppsResponseIncluded) SetAppsettings(v map[string]ViewAppSetting) {
	o.Appsettings = &v
}

func (o AppAppsResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appsettings != nil {
		toSerialize["appsettings"] = o.Appsettings
	}
	return json.Marshal(toSerialize)
}

type NullableAppAppsResponseIncluded struct {
	value *AppAppsResponseIncluded
	isSet bool
}

func (v NullableAppAppsResponseIncluded) Get() *AppAppsResponseIncluded {
	return v.value
}

func (v *NullableAppAppsResponseIncluded) Set(val *AppAppsResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAppsResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAppsResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAppsResponseIncluded(val *AppAppsResponseIncluded) *NullableAppAppsResponseIncluded {
	return &NullableAppAppsResponseIncluded{value: val, isSet: true}
}

func (v NullableAppAppsResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAppsResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


