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

// InlineResponse20072Rates struct for InlineResponse20072Rates
type InlineResponse20072Rates struct {
	AccountDefault *string `json:"account-default,omitempty"`
	ProjectDefault *string `json:"project-default,omitempty"`
	Users *InlineResponse20072RatesUsers `json:"users,omitempty"`
}

// NewInlineResponse20072Rates instantiates a new InlineResponse20072Rates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20072Rates() *InlineResponse20072Rates {
	this := InlineResponse20072Rates{}
	return &this
}

// NewInlineResponse20072RatesWithDefaults instantiates a new InlineResponse20072Rates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20072RatesWithDefaults() *InlineResponse20072Rates {
	this := InlineResponse20072Rates{}
	return &this
}

// GetAccountDefault returns the AccountDefault field value if set, zero value otherwise.
func (o *InlineResponse20072Rates) GetAccountDefault() string {
	if o == nil || o.AccountDefault == nil {
		var ret string
		return ret
	}
	return *o.AccountDefault
}

// GetAccountDefaultOk returns a tuple with the AccountDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Rates) GetAccountDefaultOk() (*string, bool) {
	if o == nil || o.AccountDefault == nil {
		return nil, false
	}
	return o.AccountDefault, true
}

// HasAccountDefault returns a boolean if a field has been set.
func (o *InlineResponse20072Rates) HasAccountDefault() bool {
	if o != nil && o.AccountDefault != nil {
		return true
	}

	return false
}

// SetAccountDefault gets a reference to the given string and assigns it to the AccountDefault field.
func (o *InlineResponse20072Rates) SetAccountDefault(v string) {
	o.AccountDefault = &v
}

// GetProjectDefault returns the ProjectDefault field value if set, zero value otherwise.
func (o *InlineResponse20072Rates) GetProjectDefault() string {
	if o == nil || o.ProjectDefault == nil {
		var ret string
		return ret
	}
	return *o.ProjectDefault
}

// GetProjectDefaultOk returns a tuple with the ProjectDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Rates) GetProjectDefaultOk() (*string, bool) {
	if o == nil || o.ProjectDefault == nil {
		return nil, false
	}
	return o.ProjectDefault, true
}

// HasProjectDefault returns a boolean if a field has been set.
func (o *InlineResponse20072Rates) HasProjectDefault() bool {
	if o != nil && o.ProjectDefault != nil {
		return true
	}

	return false
}

// SetProjectDefault gets a reference to the given string and assigns it to the ProjectDefault field.
func (o *InlineResponse20072Rates) SetProjectDefault(v string) {
	o.ProjectDefault = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineResponse20072Rates) GetUsers() InlineResponse20072RatesUsers {
	if o == nil || o.Users == nil {
		var ret InlineResponse20072RatesUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Rates) GetUsersOk() (*InlineResponse20072RatesUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineResponse20072Rates) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given InlineResponse20072RatesUsers and assigns it to the Users field.
func (o *InlineResponse20072Rates) SetUsers(v InlineResponse20072RatesUsers) {
	o.Users = &v
}

func (o InlineResponse20072Rates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountDefault != nil {
		toSerialize["account-default"] = o.AccountDefault
	}
	if o.ProjectDefault != nil {
		toSerialize["project-default"] = o.ProjectDefault
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20072Rates struct {
	value *InlineResponse20072Rates
	isSet bool
}

func (v NullableInlineResponse20072Rates) Get() *InlineResponse20072Rates {
	return v.value
}

func (v *NullableInlineResponse20072Rates) Set(val *InlineResponse20072Rates) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20072Rates) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20072Rates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20072Rates(val *InlineResponse20072Rates) *NullableInlineResponse20072Rates {
	return &NullableInlineResponse20072Rates{value: val, isSet: true}
}

func (v NullableInlineResponse20072Rates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20072Rates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


