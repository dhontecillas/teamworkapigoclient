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

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewInlineResponse20018 instantiates a new InlineResponse20018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20018() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20018WithDefaults() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20018) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20018) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20018) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20018) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20018) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InlineResponse20018) SetId(v int32) {
	o.Id = &v
}

func (o InlineResponse20018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20018 struct {
	value *InlineResponse20018
	isSet bool
}

func (v NullableInlineResponse20018) Get() *InlineResponse20018 {
	return v.value
}

func (v *NullableInlineResponse20018) Set(val *InlineResponse20018) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20018) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20018(val *InlineResponse20018) *NullableInlineResponse20018 {
	return &NullableInlineResponse20018{value: val, isSet: true}
}

func (v NullableInlineResponse20018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


