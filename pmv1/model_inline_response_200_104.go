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

// InlineResponse200104 struct for InlineResponse200104
type InlineResponse200104 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Predecessors *[]InlineResponse200104Predecessors `json:"predecessors,omitempty"`
}

// NewInlineResponse200104 instantiates a new InlineResponse200104 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200104() *InlineResponse200104 {
	this := InlineResponse200104{}
	return &this
}

// NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200104WithDefaults() *InlineResponse200104 {
	this := InlineResponse200104{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200104) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200104) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200104) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetPredecessors returns the Predecessors field value if set, zero value otherwise.
func (o *InlineResponse200104) GetPredecessors() []InlineResponse200104Predecessors {
	if o == nil || o.Predecessors == nil {
		var ret []InlineResponse200104Predecessors
		return ret
	}
	return *o.Predecessors
}

// GetPredecessorsOk returns a tuple with the Predecessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104) GetPredecessorsOk() (*[]InlineResponse200104Predecessors, bool) {
	if o == nil || o.Predecessors == nil {
		return nil, false
	}
	return o.Predecessors, true
}

// HasPredecessors returns a boolean if a field has been set.
func (o *InlineResponse200104) HasPredecessors() bool {
	if o != nil && o.Predecessors != nil {
		return true
	}

	return false
}

// SetPredecessors gets a reference to the given []InlineResponse200104Predecessors and assigns it to the Predecessors field.
func (o *InlineResponse200104) SetPredecessors(v []InlineResponse200104Predecessors) {
	o.Predecessors = &v
}

func (o InlineResponse200104) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Predecessors != nil {
		toSerialize["predecessors"] = o.Predecessors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200104 struct {
	value *InlineResponse200104
	isSet bool
}

func (v NullableInlineResponse200104) Get() *InlineResponse200104 {
	return v.value
}

func (v *NullableInlineResponse200104) Set(val *InlineResponse200104) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200104) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200104) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200104(val *InlineResponse200104) *NullableInlineResponse200104 {
	return &NullableInlineResponse200104{value: val, isSet: true}
}

func (v NullableInlineResponse200104) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200104) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

