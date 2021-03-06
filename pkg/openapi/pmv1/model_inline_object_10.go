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

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	Company CompaniesJsonCompany `json:"company"`
}

// NewInlineObject10 instantiates a new InlineObject10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10(company CompaniesJsonCompany) *InlineObject10 {
	this := InlineObject10{}
	this.Company = company
	return &this
}

// NewInlineObject10WithDefaults instantiates a new InlineObject10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10WithDefaults() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// GetCompany returns the Company field value
func (o *InlineObject10) GetCompany() CompaniesJsonCompany {
	if o == nil {
		var ret CompaniesJsonCompany
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetCompanyOk() (*CompaniesJsonCompany, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *InlineObject10) SetCompany(v CompaniesJsonCompany) {
	o.Company = v
}

func (o InlineObject10) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company"] = o.Company
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject10 struct {
	value *InlineObject10
	isSet bool
}

func (v NullableInlineObject10) Get() *InlineObject10 {
	return v.value
}

func (v *NullableInlineObject10) Set(val *InlineObject10) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10(val *InlineObject10) *NullableInlineObject10 {
	return &NullableInlineObject10{value: val, isSet: true}
}

func (v NullableInlineObject10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


