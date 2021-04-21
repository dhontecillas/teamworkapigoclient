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

// InlineResponse20027 struct for InlineResponse20027
type InlineResponse20027 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Person *InlineResponse20027Person `json:"person,omitempty"`
}

// NewInlineResponse20027 instantiates a new InlineResponse20027 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027WithDefaults() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20027) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20027) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20027) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *InlineResponse20027) GetPerson() InlineResponse20027Person {
	if o == nil || o.Person == nil {
		var ret InlineResponse20027Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetPersonOk() (*InlineResponse20027Person, bool) {
	if o == nil || o.Person == nil {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *InlineResponse20027) HasPerson() bool {
	if o != nil && o.Person != nil {
		return true
	}

	return false
}

// SetPerson gets a reference to the given InlineResponse20027Person and assigns it to the Person field.
func (o *InlineResponse20027) SetPerson(v InlineResponse20027Person) {
	o.Person = &v
}

func (o InlineResponse20027) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Person != nil {
		toSerialize["person"] = o.Person
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027 struct {
	value *InlineResponse20027
	isSet bool
}

func (v NullableInlineResponse20027) Get() *InlineResponse20027 {
	return v.value
}

func (v *NullableInlineResponse20027) Set(val *InlineResponse20027) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027(val *InlineResponse20027) *NullableInlineResponse20027 {
	return &NullableInlineResponse20027{value: val, isSet: true}
}

func (v NullableInlineResponse20027) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


