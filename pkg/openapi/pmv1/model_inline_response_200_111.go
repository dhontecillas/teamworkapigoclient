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

// InlineResponse200111 struct for InlineResponse200111
type InlineResponse200111 struct {
	STATUS *string `json:"STATUS,omitempty"`
	TimeTotals *InlineResponse20098TasklistTimeTotals `json:"time-totals,omitempty"`
}

// NewInlineResponse200111 instantiates a new InlineResponse200111 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200111() *InlineResponse200111 {
	this := InlineResponse200111{}
	return &this
}

// NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200111WithDefaults() *InlineResponse200111 {
	this := InlineResponse200111{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200111) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200111) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200111) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTimeTotals returns the TimeTotals field value if set, zero value otherwise.
func (o *InlineResponse200111) GetTimeTotals() InlineResponse20098TasklistTimeTotals {
	if o == nil || o.TimeTotals == nil {
		var ret InlineResponse20098TasklistTimeTotals
		return ret
	}
	return *o.TimeTotals
}

// GetTimeTotalsOk returns a tuple with the TimeTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200111) GetTimeTotalsOk() (*InlineResponse20098TasklistTimeTotals, bool) {
	if o == nil || o.TimeTotals == nil {
		return nil, false
	}
	return o.TimeTotals, true
}

// HasTimeTotals returns a boolean if a field has been set.
func (o *InlineResponse200111) HasTimeTotals() bool {
	if o != nil && o.TimeTotals != nil {
		return true
	}

	return false
}

// SetTimeTotals gets a reference to the given InlineResponse20098TasklistTimeTotals and assigns it to the TimeTotals field.
func (o *InlineResponse200111) SetTimeTotals(v InlineResponse20098TasklistTimeTotals) {
	o.TimeTotals = &v
}

func (o InlineResponse200111) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.TimeTotals != nil {
		toSerialize["time-totals"] = o.TimeTotals
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200111 struct {
	value *InlineResponse200111
	isSet bool
}

func (v NullableInlineResponse200111) Get() *InlineResponse200111 {
	return v.value
}

func (v *NullableInlineResponse200111) Set(val *InlineResponse200111) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200111) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200111) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200111(val *InlineResponse200111) *NullableInlineResponse200111 {
	return &NullableInlineResponse200111{value: val, isSet: true}
}

func (v NullableInlineResponse200111) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200111) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


