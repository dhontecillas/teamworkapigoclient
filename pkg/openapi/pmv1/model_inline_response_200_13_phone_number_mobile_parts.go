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

// InlineResponse20013PhoneNumberMobileParts struct for InlineResponse20013PhoneNumberMobileParts
type InlineResponse20013PhoneNumberMobileParts struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

// NewInlineResponse20013PhoneNumberMobileParts instantiates a new InlineResponse20013PhoneNumberMobileParts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013PhoneNumberMobileParts() *InlineResponse20013PhoneNumberMobileParts {
	this := InlineResponse20013PhoneNumberMobileParts{}
	return &this
}

// NewInlineResponse20013PhoneNumberMobilePartsWithDefaults instantiates a new InlineResponse20013PhoneNumberMobileParts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013PhoneNumberMobilePartsWithDefaults() *InlineResponse20013PhoneNumberMobileParts {
	this := InlineResponse20013PhoneNumberMobileParts{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *InlineResponse20013PhoneNumberMobileParts) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *InlineResponse20013PhoneNumberMobileParts) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *InlineResponse20013PhoneNumberMobileParts) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *InlineResponse20013PhoneNumberMobileParts) SetPhone(v string) {
	o.Phone = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineResponse20013PhoneNumberMobileParts) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineResponse20013PhoneNumberMobileParts) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineResponse20013PhoneNumberMobileParts) SetPrefix(v string) {
	o.Prefix = &v
}

func (o InlineResponse20013PhoneNumberMobileParts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20013PhoneNumberMobileParts struct {
	value *InlineResponse20013PhoneNumberMobileParts
	isSet bool
}

func (v NullableInlineResponse20013PhoneNumberMobileParts) Get() *InlineResponse20013PhoneNumberMobileParts {
	return v.value
}

func (v *NullableInlineResponse20013PhoneNumberMobileParts) Set(val *InlineResponse20013PhoneNumberMobileParts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013PhoneNumberMobileParts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013PhoneNumberMobileParts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013PhoneNumberMobileParts(val *InlineResponse20013PhoneNumberMobileParts) *NullableInlineResponse20013PhoneNumberMobileParts {
	return &NullableInlineResponse20013PhoneNumberMobileParts{value: val, isSet: true}
}

func (v NullableInlineResponse20013PhoneNumberMobileParts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013PhoneNumberMobileParts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


