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

// InlineResponse20029 struct for InlineResponse20029
type InlineResponse20029 struct {
	Userstatus *MeStatusJsonUserstatus `json:"userstatus,omitempty"`
}

// NewInlineResponse20029 instantiates a new InlineResponse20029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029WithDefaults() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// GetUserstatus returns the Userstatus field value if set, zero value otherwise.
func (o *InlineResponse20029) GetUserstatus() MeStatusJsonUserstatus {
	if o == nil || o.Userstatus == nil {
		var ret MeStatusJsonUserstatus
		return ret
	}
	return *o.Userstatus
}

// GetUserstatusOk returns a tuple with the Userstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetUserstatusOk() (*MeStatusJsonUserstatus, bool) {
	if o == nil || o.Userstatus == nil {
		return nil, false
	}
	return o.Userstatus, true
}

// HasUserstatus returns a boolean if a field has been set.
func (o *InlineResponse20029) HasUserstatus() bool {
	if o != nil && o.Userstatus != nil {
		return true
	}

	return false
}

// SetUserstatus gets a reference to the given MeStatusJsonUserstatus and assigns it to the Userstatus field.
func (o *InlineResponse20029) SetUserstatus(v MeStatusJsonUserstatus) {
	o.Userstatus = &v
}

func (o InlineResponse20029) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Userstatus != nil {
		toSerialize["userstatus"] = o.Userstatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029 struct {
	value *InlineResponse20029
	isSet bool
}

func (v NullableInlineResponse20029) Get() *InlineResponse20029 {
	return v.value
}

func (v *NullableInlineResponse20029) Set(val *InlineResponse20029) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029(val *InlineResponse20029) *NullableInlineResponse20029 {
	return &NullableInlineResponse20029{value: val, isSet: true}
}

func (v NullableInlineResponse20029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


