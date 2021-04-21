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

// InlineResponse20026 struct for InlineResponse20026
type InlineResponse20026 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Link *InlineResponse20026Link `json:"link,omitempty"`
}

// NewInlineResponse20026 instantiates a new InlineResponse20026 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026WithDefaults() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20026) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20026) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20026) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *InlineResponse20026) GetLink() InlineResponse20026Link {
	if o == nil || o.Link == nil {
		var ret InlineResponse20026Link
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetLinkOk() (*InlineResponse20026Link, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *InlineResponse20026) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given InlineResponse20026Link and assigns it to the Link field.
func (o *InlineResponse20026) SetLink(v InlineResponse20026Link) {
	o.Link = &v
}

func (o InlineResponse20026) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026 struct {
	value *InlineResponse20026
	isSet bool
}

func (v NullableInlineResponse20026) Get() *InlineResponse20026 {
	return v.value
}

func (v *NullableInlineResponse20026) Set(val *InlineResponse20026) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026(val *InlineResponse20026) *NullableInlineResponse20026 {
	return &NullableInlineResponse20026{value: val, isSet: true}
}

func (v NullableInlineResponse20026) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


