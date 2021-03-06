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

// InlineResponse20074 struct for InlineResponse20074
type InlineResponse20074 struct {
	Roles *[]InlineResponse20074Roles `json:"roles,omitempty"`
}

// NewInlineResponse20074 instantiates a new InlineResponse20074 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074() *InlineResponse20074 {
	this := InlineResponse20074{}
	return &this
}

// NewInlineResponse20074WithDefaults instantiates a new InlineResponse20074 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074WithDefaults() *InlineResponse20074 {
	this := InlineResponse20074{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InlineResponse20074) GetRoles() []InlineResponse20074Roles {
	if o == nil || o.Roles == nil {
		var ret []InlineResponse20074Roles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetRolesOk() (*[]InlineResponse20074Roles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InlineResponse20074) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []InlineResponse20074Roles and assigns it to the Roles field.
func (o *InlineResponse20074) SetRoles(v []InlineResponse20074Roles) {
	o.Roles = &v
}

func (o InlineResponse20074) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074 struct {
	value *InlineResponse20074
	isSet bool
}

func (v NullableInlineResponse20074) Get() *InlineResponse20074 {
	return v.value
}

func (v *NullableInlineResponse20074) Set(val *InlineResponse20074) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074(val *InlineResponse20074) *NullableInlineResponse20074 {
	return &NullableInlineResponse20074{value: val, isSet: true}
}

func (v NullableInlineResponse20074) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


