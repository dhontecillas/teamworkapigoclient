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

// ColumnResponseIncluded struct for ColumnResponseIncluded
type ColumnResponseIncluded struct {
	Tasklists *map[string]ViewTasklist `json:"tasklists,omitempty"`
}

// NewColumnResponseIncluded instantiates a new ColumnResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnResponseIncluded() *ColumnResponseIncluded {
	this := ColumnResponseIncluded{}
	return &this
}

// NewColumnResponseIncludedWithDefaults instantiates a new ColumnResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnResponseIncludedWithDefaults() *ColumnResponseIncluded {
	this := ColumnResponseIncluded{}
	return &this
}

// GetTasklists returns the Tasklists field value if set, zero value otherwise.
func (o *ColumnResponseIncluded) GetTasklists() map[string]ViewTasklist {
	if o == nil || o.Tasklists == nil {
		var ret map[string]ViewTasklist
		return ret
	}
	return *o.Tasklists
}

// GetTasklistsOk returns a tuple with the Tasklists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool) {
	if o == nil || o.Tasklists == nil {
		return nil, false
	}
	return o.Tasklists, true
}

// HasTasklists returns a boolean if a field has been set.
func (o *ColumnResponseIncluded) HasTasklists() bool {
	if o != nil && o.Tasklists != nil {
		return true
	}

	return false
}

// SetTasklists gets a reference to the given map[string]ViewTasklist and assigns it to the Tasklists field.
func (o *ColumnResponseIncluded) SetTasklists(v map[string]ViewTasklist) {
	o.Tasklists = &v
}

func (o ColumnResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tasklists != nil {
		toSerialize["tasklists"] = o.Tasklists
	}
	return json.Marshal(toSerialize)
}

type NullableColumnResponseIncluded struct {
	value *ColumnResponseIncluded
	isSet bool
}

func (v NullableColumnResponseIncluded) Get() *ColumnResponseIncluded {
	return v.value
}

func (v *NullableColumnResponseIncluded) Set(val *ColumnResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnResponseIncluded(val *ColumnResponseIncluded) *NullableColumnResponseIncluded {
	return &NullableColumnResponseIncluded{value: val, isSet: true}
}

func (v NullableColumnResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


