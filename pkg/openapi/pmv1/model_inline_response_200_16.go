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

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	STATUS *string `json:"STATUS,omitempty"`
	CurrencyCodes *[]InlineResponse20016CurrencyCodes `json:"currency-codes,omitempty"`
}

// NewInlineResponse20016 instantiates a new InlineResponse20016 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016WithDefaults() *InlineResponse20016 {
	this := InlineResponse20016{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20016) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20016) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20016) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetCurrencyCodes returns the CurrencyCodes field value if set, zero value otherwise.
func (o *InlineResponse20016) GetCurrencyCodes() []InlineResponse20016CurrencyCodes {
	if o == nil || o.CurrencyCodes == nil {
		var ret []InlineResponse20016CurrencyCodes
		return ret
	}
	return *o.CurrencyCodes
}

// GetCurrencyCodesOk returns a tuple with the CurrencyCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016) GetCurrencyCodesOk() (*[]InlineResponse20016CurrencyCodes, bool) {
	if o == nil || o.CurrencyCodes == nil {
		return nil, false
	}
	return o.CurrencyCodes, true
}

// HasCurrencyCodes returns a boolean if a field has been set.
func (o *InlineResponse20016) HasCurrencyCodes() bool {
	if o != nil && o.CurrencyCodes != nil {
		return true
	}

	return false
}

// SetCurrencyCodes gets a reference to the given []InlineResponse20016CurrencyCodes and assigns it to the CurrencyCodes field.
func (o *InlineResponse20016) SetCurrencyCodes(v []InlineResponse20016CurrencyCodes) {
	o.CurrencyCodes = &v
}

func (o InlineResponse20016) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.CurrencyCodes != nil {
		toSerialize["currency-codes"] = o.CurrencyCodes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016 struct {
	value *InlineResponse20016
	isSet bool
}

func (v NullableInlineResponse20016) Get() *InlineResponse20016 {
	return v.value
}

func (v *NullableInlineResponse20016) Set(val *InlineResponse20016) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016(val *InlineResponse20016) *NullableInlineResponse20016 {
	return &NullableInlineResponse20016{value: val, isSet: true}
}

func (v NullableInlineResponse20016) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


