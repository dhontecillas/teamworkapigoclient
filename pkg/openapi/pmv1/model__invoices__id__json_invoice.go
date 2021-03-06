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

// InvoicesIdJsonInvoice struct for InvoicesIdJsonInvoice
type InvoicesIdJsonInvoice struct {
	CurrencyCode *string `json:"currency-code,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayDate *string `json:"display-date,omitempty"`
	FixedCost *string `json:"fixed-cost,omitempty"`
	Number *string `json:"number,omitempty"`
	PoNumber *string `json:"po-number,omitempty"`
}

// NewInvoicesIdJsonInvoice instantiates a new InvoicesIdJsonInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesIdJsonInvoice() *InvoicesIdJsonInvoice {
	this := InvoicesIdJsonInvoice{}
	return &this
}

// NewInvoicesIdJsonInvoiceWithDefaults instantiates a new InvoicesIdJsonInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesIdJsonInvoiceWithDefaults() *InvoicesIdJsonInvoice {
	this := InvoicesIdJsonInvoice{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *InvoicesIdJsonInvoice) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InvoicesIdJsonInvoice) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDate returns the DisplayDate field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetDisplayDate() string {
	if o == nil || o.DisplayDate == nil {
		var ret string
		return ret
	}
	return *o.DisplayDate
}

// GetDisplayDateOk returns a tuple with the DisplayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetDisplayDateOk() (*string, bool) {
	if o == nil || o.DisplayDate == nil {
		return nil, false
	}
	return o.DisplayDate, true
}

// HasDisplayDate returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasDisplayDate() bool {
	if o != nil && o.DisplayDate != nil {
		return true
	}

	return false
}

// SetDisplayDate gets a reference to the given string and assigns it to the DisplayDate field.
func (o *InvoicesIdJsonInvoice) SetDisplayDate(v string) {
	o.DisplayDate = &v
}

// GetFixedCost returns the FixedCost field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetFixedCost() string {
	if o == nil || o.FixedCost == nil {
		var ret string
		return ret
	}
	return *o.FixedCost
}

// GetFixedCostOk returns a tuple with the FixedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetFixedCostOk() (*string, bool) {
	if o == nil || o.FixedCost == nil {
		return nil, false
	}
	return o.FixedCost, true
}

// HasFixedCost returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasFixedCost() bool {
	if o != nil && o.FixedCost != nil {
		return true
	}

	return false
}

// SetFixedCost gets a reference to the given string and assigns it to the FixedCost field.
func (o *InvoicesIdJsonInvoice) SetFixedCost(v string) {
	o.FixedCost = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *InvoicesIdJsonInvoice) SetNumber(v string) {
	o.Number = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *InvoicesIdJsonInvoice) GetPoNumber() string {
	if o == nil || o.PoNumber == nil {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesIdJsonInvoice) GetPoNumberOk() (*string, bool) {
	if o == nil || o.PoNumber == nil {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *InvoicesIdJsonInvoice) HasPoNumber() bool {
	if o != nil && o.PoNumber != nil {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *InvoicesIdJsonInvoice) SetPoNumber(v string) {
	o.PoNumber = &v
}

func (o InvoicesIdJsonInvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode != nil {
		toSerialize["currency-code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayDate != nil {
		toSerialize["display-date"] = o.DisplayDate
	}
	if o.FixedCost != nil {
		toSerialize["fixed-cost"] = o.FixedCost
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.PoNumber != nil {
		toSerialize["po-number"] = o.PoNumber
	}
	return json.Marshal(toSerialize)
}

type NullableInvoicesIdJsonInvoice struct {
	value *InvoicesIdJsonInvoice
	isSet bool
}

func (v NullableInvoicesIdJsonInvoice) Get() *InvoicesIdJsonInvoice {
	return v.value
}

func (v *NullableInvoicesIdJsonInvoice) Set(val *InvoicesIdJsonInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesIdJsonInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesIdJsonInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesIdJsonInvoice(val *InvoicesIdJsonInvoice) *NullableInvoicesIdJsonInvoice {
	return &NullableInvoicesIdJsonInvoice{value: val, isSet: true}
}

func (v NullableInvoicesIdJsonInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesIdJsonInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


