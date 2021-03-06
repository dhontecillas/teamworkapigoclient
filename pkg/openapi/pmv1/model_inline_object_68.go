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

// InlineObject68 struct for InlineObject68
type InlineObject68 struct {
	Add *ProjectsIdPeopleJsonAdd `json:"add,omitempty"`
	Remove *ProjectsIdPeopleJsonAdd `json:"remove,omitempty"`
}

// NewInlineObject68 instantiates a new InlineObject68 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject68() *InlineObject68 {
	this := InlineObject68{}
	return &this
}

// NewInlineObject68WithDefaults instantiates a new InlineObject68 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject68WithDefaults() *InlineObject68 {
	this := InlineObject68{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *InlineObject68) GetAdd() ProjectsIdPeopleJsonAdd {
	if o == nil || o.Add == nil {
		var ret ProjectsIdPeopleJsonAdd
		return ret
	}
	return *o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject68) GetAddOk() (*ProjectsIdPeopleJsonAdd, bool) {
	if o == nil || o.Add == nil {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *InlineObject68) HasAdd() bool {
	if o != nil && o.Add != nil {
		return true
	}

	return false
}

// SetAdd gets a reference to the given ProjectsIdPeopleJsonAdd and assigns it to the Add field.
func (o *InlineObject68) SetAdd(v ProjectsIdPeopleJsonAdd) {
	o.Add = &v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *InlineObject68) GetRemove() ProjectsIdPeopleJsonAdd {
	if o == nil || o.Remove == nil {
		var ret ProjectsIdPeopleJsonAdd
		return ret
	}
	return *o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject68) GetRemoveOk() (*ProjectsIdPeopleJsonAdd, bool) {
	if o == nil || o.Remove == nil {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *InlineObject68) HasRemove() bool {
	if o != nil && o.Remove != nil {
		return true
	}

	return false
}

// SetRemove gets a reference to the given ProjectsIdPeopleJsonAdd and assigns it to the Remove field.
func (o *InlineObject68) SetRemove(v ProjectsIdPeopleJsonAdd) {
	o.Remove = &v
}

func (o InlineObject68) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Add != nil {
		toSerialize["add"] = o.Add
	}
	if o.Remove != nil {
		toSerialize["remove"] = o.Remove
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject68 struct {
	value *InlineObject68
	isSet bool
}

func (v NullableInlineObject68) Get() *InlineObject68 {
	return v.value
}

func (v *NullableInlineObject68) Set(val *InlineObject68) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject68) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject68) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject68(val *InlineObject68) *NullableInlineObject68 {
	return &NullableInlineObject68{value: val, isSet: true}
}

func (v NullableInlineObject68) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject68) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


