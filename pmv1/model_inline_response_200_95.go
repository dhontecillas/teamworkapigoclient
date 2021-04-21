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

// InlineResponse20095 struct for InlineResponse20095
type InlineResponse20095 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Tasklists *[]InlineResponse20095Tasklists `json:"tasklists,omitempty"`
}

// NewInlineResponse20095 instantiates a new InlineResponse20095 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// NewInlineResponse20095WithDefaults instantiates a new InlineResponse20095 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095WithDefaults() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20095) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20095) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20095) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTasklists returns the Tasklists field value if set, zero value otherwise.
func (o *InlineResponse20095) GetTasklists() []InlineResponse20095Tasklists {
	if o == nil || o.Tasklists == nil {
		var ret []InlineResponse20095Tasklists
		return ret
	}
	return *o.Tasklists
}

// GetTasklistsOk returns a tuple with the Tasklists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetTasklistsOk() (*[]InlineResponse20095Tasklists, bool) {
	if o == nil || o.Tasklists == nil {
		return nil, false
	}
	return o.Tasklists, true
}

// HasTasklists returns a boolean if a field has been set.
func (o *InlineResponse20095) HasTasklists() bool {
	if o != nil && o.Tasklists != nil {
		return true
	}

	return false
}

// SetTasklists gets a reference to the given []InlineResponse20095Tasklists and assigns it to the Tasklists field.
func (o *InlineResponse20095) SetTasklists(v []InlineResponse20095Tasklists) {
	o.Tasklists = &v
}

func (o InlineResponse20095) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Tasklists != nil {
		toSerialize["tasklists"] = o.Tasklists
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095 struct {
	value *InlineResponse20095
	isSet bool
}

func (v NullableInlineResponse20095) Get() *InlineResponse20095 {
	return v.value
}

func (v *NullableInlineResponse20095) Set(val *InlineResponse20095) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095(val *InlineResponse20095) *NullableInlineResponse20095 {
	return &NullableInlineResponse20095{value: val, isSet: true}
}

func (v NullableInlineResponse20095) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


