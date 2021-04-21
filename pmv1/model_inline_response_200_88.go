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

// InlineResponse20088 struct for InlineResponse20088
type InlineResponse20088 struct {
	Role *InlineResponse20088Role `json:"role,omitempty"`
}

// NewInlineResponse20088 instantiates a new InlineResponse20088 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20088() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// NewInlineResponse20088WithDefaults instantiates a new InlineResponse20088 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20088WithDefaults() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InlineResponse20088) GetRole() InlineResponse20088Role {
	if o == nil || o.Role == nil {
		var ret InlineResponse20088Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetRoleOk() (*InlineResponse20088Role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InlineResponse20088) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given InlineResponse20088Role and assigns it to the Role field.
func (o *InlineResponse20088) SetRole(v InlineResponse20088Role) {
	o.Role = &v
}

func (o InlineResponse20088) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20088 struct {
	value *InlineResponse20088
	isSet bool
}

func (v NullableInlineResponse20088) Get() *InlineResponse20088 {
	return v.value
}

func (v *NullableInlineResponse20088) Set(val *InlineResponse20088) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20088) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20088) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20088(val *InlineResponse20088) *NullableInlineResponse20088 {
	return &NullableInlineResponse20088{value: val, isSet: true}
}

func (v NullableInlineResponse20088) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20088) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

