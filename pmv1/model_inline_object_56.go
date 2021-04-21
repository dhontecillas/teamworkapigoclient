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

// InlineObject56 struct for InlineObject56
type InlineObject56 struct {
	Column *EventtypesJsonEventtype `json:"column,omitempty"`
}

// NewInlineObject56 instantiates a new InlineObject56 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject56() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// NewInlineObject56WithDefaults instantiates a new InlineObject56 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject56WithDefaults() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *InlineObject56) GetColumn() EventtypesJsonEventtype {
	if o == nil || o.Column == nil {
		var ret EventtypesJsonEventtype
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject56) GetColumnOk() (*EventtypesJsonEventtype, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *InlineObject56) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given EventtypesJsonEventtype and assigns it to the Column field.
func (o *InlineObject56) SetColumn(v EventtypesJsonEventtype) {
	o.Column = &v
}

func (o InlineObject56) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject56 struct {
	value *InlineObject56
	isSet bool
}

func (v NullableInlineObject56) Get() *InlineObject56 {
	return v.value
}

func (v *NullableInlineObject56) Set(val *InlineObject56) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject56) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject56) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject56(val *InlineObject56) *NullableInlineObject56 {
	return &NullableInlineObject56{value: val, isSet: true}
}

func (v NullableInlineObject56) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject56) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

