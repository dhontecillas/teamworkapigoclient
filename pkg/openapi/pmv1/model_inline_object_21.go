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

// InlineObject21 struct for InlineObject21
type InlineObject21 struct {
	Invoice InvoicesIdJsonInvoice `json:"invoice"`
}

// NewInlineObject21 instantiates a new InlineObject21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject21(invoice InvoicesIdJsonInvoice) *InlineObject21 {
	this := InlineObject21{}
	this.Invoice = invoice
	return &this
}

// NewInlineObject21WithDefaults instantiates a new InlineObject21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject21WithDefaults() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// GetInvoice returns the Invoice field value
func (o *InlineObject21) GetInvoice() InvoicesIdJsonInvoice {
	if o == nil {
		var ret InvoicesIdJsonInvoice
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetInvoiceOk() (*InvoicesIdJsonInvoice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *InlineObject21) SetInvoice(v InvoicesIdJsonInvoice) {
	o.Invoice = v
}

func (o InlineObject21) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invoice"] = o.Invoice
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject21 struct {
	value *InlineObject21
	isSet bool
}

func (v NullableInlineObject21) Get() *InlineObject21 {
	return v.value
}

func (v *NullableInlineObject21) Set(val *InlineObject21) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject21) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject21(val *InlineObject21) *NullableInlineObject21 {
	return &NullableInlineObject21{value: val, isSet: true}
}

func (v NullableInlineObject21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


