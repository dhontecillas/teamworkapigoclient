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

// InlineResponse20076 struct for InlineResponse20076
type InlineResponse20076 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Tasklists *[]InlineResponse20076Tasklists `json:"tasklists,omitempty"`
}

// NewInlineResponse20076 instantiates a new InlineResponse20076 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20076() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20076WithDefaults() *InlineResponse20076 {
	this := InlineResponse20076{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20076) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20076) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20076) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTasklists returns the Tasklists field value if set, zero value otherwise.
func (o *InlineResponse20076) GetTasklists() []InlineResponse20076Tasklists {
	if o == nil || o.Tasklists == nil {
		var ret []InlineResponse20076Tasklists
		return ret
	}
	return *o.Tasklists
}

// GetTasklistsOk returns a tuple with the Tasklists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076) GetTasklistsOk() (*[]InlineResponse20076Tasklists, bool) {
	if o == nil || o.Tasklists == nil {
		return nil, false
	}
	return o.Tasklists, true
}

// HasTasklists returns a boolean if a field has been set.
func (o *InlineResponse20076) HasTasklists() bool {
	if o != nil && o.Tasklists != nil {
		return true
	}

	return false
}

// SetTasklists gets a reference to the given []InlineResponse20076Tasklists and assigns it to the Tasklists field.
func (o *InlineResponse20076) SetTasklists(v []InlineResponse20076Tasklists) {
	o.Tasklists = &v
}

func (o InlineResponse20076) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Tasklists != nil {
		toSerialize["tasklists"] = o.Tasklists
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20076 struct {
	value *InlineResponse20076
	isSet bool
}

func (v NullableInlineResponse20076) Get() *InlineResponse20076 {
	return v.value
}

func (v *NullableInlineResponse20076) Set(val *InlineResponse20076) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20076) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20076) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20076(val *InlineResponse20076) *NullableInlineResponse20076 {
	return &NullableInlineResponse20076{value: val, isSet: true}
}

func (v NullableInlineResponse20076) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20076) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


