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

// InlineObject22 struct for InlineObject22
type InlineObject22 struct {
	Lineitems *InvoicesIdLineitemsJsonLineitems `json:"lineitems,omitempty"`
}

// NewInlineObject22 instantiates a new InlineObject22 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject22() *InlineObject22 {
	this := InlineObject22{}
	return &this
}

// NewInlineObject22WithDefaults instantiates a new InlineObject22 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject22WithDefaults() *InlineObject22 {
	this := InlineObject22{}
	return &this
}

// GetLineitems returns the Lineitems field value if set, zero value otherwise.
func (o *InlineObject22) GetLineitems() InvoicesIdLineitemsJsonLineitems {
	if o == nil || o.Lineitems == nil {
		var ret InvoicesIdLineitemsJsonLineitems
		return ret
	}
	return *o.Lineitems
}

// GetLineitemsOk returns a tuple with the Lineitems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject22) GetLineitemsOk() (*InvoicesIdLineitemsJsonLineitems, bool) {
	if o == nil || o.Lineitems == nil {
		return nil, false
	}
	return o.Lineitems, true
}

// HasLineitems returns a boolean if a field has been set.
func (o *InlineObject22) HasLineitems() bool {
	if o != nil && o.Lineitems != nil {
		return true
	}

	return false
}

// SetLineitems gets a reference to the given InvoicesIdLineitemsJsonLineitems and assigns it to the Lineitems field.
func (o *InlineObject22) SetLineitems(v InvoicesIdLineitemsJsonLineitems) {
	o.Lineitems = &v
}

func (o InlineObject22) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lineitems != nil {
		toSerialize["lineitems"] = o.Lineitems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject22 struct {
	value *InlineObject22
	isSet bool
}

func (v NullableInlineObject22) Get() *InlineObject22 {
	return v.value
}

func (v *NullableInlineObject22) Set(val *InlineObject22) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject22) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject22) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject22(val *InlineObject22) *NullableInlineObject22 {
	return &NullableInlineObject22{value: val, isSet: true}
}

func (v NullableInlineObject22) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject22) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


