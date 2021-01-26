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

// ProjectsIdInvoicesJsonInvoice struct for ProjectsIdInvoicesJsonInvoice
type ProjectsIdInvoicesJsonInvoice struct {
	CurrencyCode *string `json:"currency-code,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayDate string `json:"display-date"`
	FixedCost *string `json:"fixed-cost,omitempty"`
	Number string `json:"number"`
	PoNumber *string `json:"po-number,omitempty"`
}

// NewProjectsIdInvoicesJsonInvoice instantiates a new ProjectsIdInvoicesJsonInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdInvoicesJsonInvoice(displayDate string, number string, ) *ProjectsIdInvoicesJsonInvoice {
	this := ProjectsIdInvoicesJsonInvoice{}
	this.DisplayDate = displayDate
	this.Number = number
	return &this
}

// NewProjectsIdInvoicesJsonInvoiceWithDefaults instantiates a new ProjectsIdInvoicesJsonInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdInvoicesJsonInvoiceWithDefaults() *ProjectsIdInvoicesJsonInvoice {
	this := ProjectsIdInvoicesJsonInvoice{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ProjectsIdInvoicesJsonInvoice) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ProjectsIdInvoicesJsonInvoice) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ProjectsIdInvoicesJsonInvoice) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsIdInvoicesJsonInvoice) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsIdInvoicesJsonInvoice) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsIdInvoicesJsonInvoice) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDate returns the DisplayDate field value
func (o *ProjectsIdInvoicesJsonInvoice) GetDisplayDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayDate
}

// GetDisplayDateOk returns a tuple with the DisplayDate field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetDisplayDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayDate, true
}

// SetDisplayDate sets field value
func (o *ProjectsIdInvoicesJsonInvoice) SetDisplayDate(v string) {
	o.DisplayDate = v
}

// GetFixedCost returns the FixedCost field value if set, zero value otherwise.
func (o *ProjectsIdInvoicesJsonInvoice) GetFixedCost() string {
	if o == nil || o.FixedCost == nil {
		var ret string
		return ret
	}
	return *o.FixedCost
}

// GetFixedCostOk returns a tuple with the FixedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetFixedCostOk() (*string, bool) {
	if o == nil || o.FixedCost == nil {
		return nil, false
	}
	return o.FixedCost, true
}

// HasFixedCost returns a boolean if a field has been set.
func (o *ProjectsIdInvoicesJsonInvoice) HasFixedCost() bool {
	if o != nil && o.FixedCost != nil {
		return true
	}

	return false
}

// SetFixedCost gets a reference to the given string and assigns it to the FixedCost field.
func (o *ProjectsIdInvoicesJsonInvoice) SetFixedCost(v string) {
	o.FixedCost = &v
}

// GetNumber returns the Number field value
func (o *ProjectsIdInvoicesJsonInvoice) GetNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *ProjectsIdInvoicesJsonInvoice) SetNumber(v string) {
	o.Number = v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *ProjectsIdInvoicesJsonInvoice) GetPoNumber() string {
	if o == nil || o.PoNumber == nil {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdInvoicesJsonInvoice) GetPoNumberOk() (*string, bool) {
	if o == nil || o.PoNumber == nil {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *ProjectsIdInvoicesJsonInvoice) HasPoNumber() bool {
	if o != nil && o.PoNumber != nil {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *ProjectsIdInvoicesJsonInvoice) SetPoNumber(v string) {
	o.PoNumber = &v
}

func (o ProjectsIdInvoicesJsonInvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode != nil {
		toSerialize["currency-code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["display-date"] = o.DisplayDate
	}
	if o.FixedCost != nil {
		toSerialize["fixed-cost"] = o.FixedCost
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if o.PoNumber != nil {
		toSerialize["po-number"] = o.PoNumber
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdInvoicesJsonInvoice struct {
	value *ProjectsIdInvoicesJsonInvoice
	isSet bool
}

func (v NullableProjectsIdInvoicesJsonInvoice) Get() *ProjectsIdInvoicesJsonInvoice {
	return v.value
}

func (v *NullableProjectsIdInvoicesJsonInvoice) Set(val *ProjectsIdInvoicesJsonInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdInvoicesJsonInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdInvoicesJsonInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdInvoicesJsonInvoice(val *ProjectsIdInvoicesJsonInvoice) *NullableProjectsIdInvoicesJsonInvoice {
	return &NullableProjectsIdInvoicesJsonInvoice{value: val, isSet: true}
}

func (v NullableProjectsIdInvoicesJsonInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdInvoicesJsonInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


