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

// InlineResponse20081 struct for InlineResponse20081
type InlineResponse20081 struct {
	People *[]InlineResponse20081People `json:"people,omitempty"`
}

// NewInlineResponse20081 instantiates a new InlineResponse20081 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20081() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// NewInlineResponse20081WithDefaults instantiates a new InlineResponse20081 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20081WithDefaults() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *InlineResponse20081) GetPeople() []InlineResponse20081People {
	if o == nil || o.People == nil {
		var ret []InlineResponse20081People
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetPeopleOk() (*[]InlineResponse20081People, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *InlineResponse20081) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []InlineResponse20081People and assigns it to the People field.
func (o *InlineResponse20081) SetPeople(v []InlineResponse20081People) {
	o.People = &v
}

func (o InlineResponse20081) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20081 struct {
	value *InlineResponse20081
	isSet bool
}

func (v NullableInlineResponse20081) Get() *InlineResponse20081 {
	return v.value
}

func (v *NullableInlineResponse20081) Set(val *InlineResponse20081) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20081) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20081) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20081(val *InlineResponse20081) *NullableInlineResponse20081 {
	return &NullableInlineResponse20081{value: val, isSet: true}
}

func (v NullableInlineResponse20081) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20081) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


