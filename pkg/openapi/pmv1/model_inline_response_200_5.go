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

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Event *InlineResponse2005Event `json:"event,omitempty"`
}

// NewInlineResponse2005 instantiates a new InlineResponse2005 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005WithDefaults() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse2005) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse2005) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse2005) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *InlineResponse2005) GetEvent() InlineResponse2005Event {
	if o == nil || o.Event == nil {
		var ret InlineResponse2005Event
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetEventOk() (*InlineResponse2005Event, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *InlineResponse2005) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given InlineResponse2005Event and assigns it to the Event field.
func (o *InlineResponse2005) SetEvent(v InlineResponse2005Event) {
	o.Event = &v
}

func (o InlineResponse2005) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005 struct {
	value *InlineResponse2005
	isSet bool
}

func (v NullableInlineResponse2005) Get() *InlineResponse2005 {
	return v.value
}

func (v *NullableInlineResponse2005) Set(val *InlineResponse2005) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005(val *InlineResponse2005) *NullableInlineResponse2005 {
	return &NullableInlineResponse2005{value: val, isSet: true}
}

func (v NullableInlineResponse2005) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


