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

// InlineResponse20023 struct for InlineResponse20023
type InlineResponse20023 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Invoice *InlineResponse20023Invoice `json:"invoice,omitempty"`
}

// NewInlineResponse20023 instantiates a new InlineResponse20023 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023WithDefaults() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20023) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20023) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20023) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *InlineResponse20023) GetInvoice() InlineResponse20023Invoice {
	if o == nil || o.Invoice == nil {
		var ret InlineResponse20023Invoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetInvoiceOk() (*InlineResponse20023Invoice, bool) {
	if o == nil || o.Invoice == nil {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *InlineResponse20023) HasInvoice() bool {
	if o != nil && o.Invoice != nil {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given InlineResponse20023Invoice and assigns it to the Invoice field.
func (o *InlineResponse20023) SetInvoice(v InlineResponse20023Invoice) {
	o.Invoice = &v
}

func (o InlineResponse20023) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Invoice != nil {
		toSerialize["invoice"] = o.Invoice
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023 struct {
	value *InlineResponse20023
	isSet bool
}

func (v NullableInlineResponse20023) Get() *InlineResponse20023 {
	return v.value
}

func (v *NullableInlineResponse20023) Set(val *InlineResponse20023) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023(val *InlineResponse20023) *NullableInlineResponse20023 {
	return &NullableInlineResponse20023{value: val, isSet: true}
}

func (v NullableInlineResponse20023) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


