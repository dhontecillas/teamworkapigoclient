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

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	Person *PeopleJsonPerson `json:"person,omitempty"`
}

// NewInlineObject37 instantiates a new InlineObject37 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject37() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// NewInlineObject37WithDefaults instantiates a new InlineObject37 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject37WithDefaults() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *InlineObject37) GetPerson() PeopleJsonPerson {
	if o == nil || o.Person == nil {
		var ret PeopleJsonPerson
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetPersonOk() (*PeopleJsonPerson, bool) {
	if o == nil || o.Person == nil {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *InlineObject37) HasPerson() bool {
	if o != nil && o.Person != nil {
		return true
	}

	return false
}

// SetPerson gets a reference to the given PeopleJsonPerson and assigns it to the Person field.
func (o *InlineObject37) SetPerson(v PeopleJsonPerson) {
	o.Person = &v
}

func (o InlineObject37) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Person != nil {
		toSerialize["person"] = o.Person
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject37 struct {
	value *InlineObject37
	isSet bool
}

func (v NullableInlineObject37) Get() *InlineObject37 {
	return v.value
}

func (v *NullableInlineObject37) Set(val *InlineObject37) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject37) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject37) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject37(val *InlineObject37) *NullableInlineObject37 {
	return &NullableInlineObject37{value: val, isSet: true}
}

func (v NullableInlineObject37) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject37) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


