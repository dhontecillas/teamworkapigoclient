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

// InlineResponse20053IntegrationsOnedrivebusiness struct for InlineResponse20053IntegrationsOnedrivebusiness
type InlineResponse20053IntegrationsOnedrivebusiness struct {
	Account *string `json:"account,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Folder *string `json:"folder,omitempty"`
	Foldername *string `json:"foldername,omitempty"`
}

// NewInlineResponse20053IntegrationsOnedrivebusiness instantiates a new InlineResponse20053IntegrationsOnedrivebusiness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053IntegrationsOnedrivebusiness() *InlineResponse20053IntegrationsOnedrivebusiness {
	this := InlineResponse20053IntegrationsOnedrivebusiness{}
	return &this
}

// NewInlineResponse20053IntegrationsOnedrivebusinessWithDefaults instantiates a new InlineResponse20053IntegrationsOnedrivebusiness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053IntegrationsOnedrivebusinessWithDefaults() *InlineResponse20053IntegrationsOnedrivebusiness {
	this := InlineResponse20053IntegrationsOnedrivebusiness{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) SetAccount(v string) {
	o.Account = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) SetFolder(v string) {
	o.Folder = &v
}

// GetFoldername returns the Foldername field value if set, zero value otherwise.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetFoldername() string {
	if o == nil || o.Foldername == nil {
		var ret string
		return ret
	}
	return *o.Foldername
}

// GetFoldernameOk returns a tuple with the Foldername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) GetFoldernameOk() (*string, bool) {
	if o == nil || o.Foldername == nil {
		return nil, false
	}
	return o.Foldername, true
}

// HasFoldername returns a boolean if a field has been set.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) HasFoldername() bool {
	if o != nil && o.Foldername != nil {
		return true
	}

	return false
}

// SetFoldername gets a reference to the given string and assigns it to the Foldername field.
func (o *InlineResponse20053IntegrationsOnedrivebusiness) SetFoldername(v string) {
	o.Foldername = &v
}

func (o InlineResponse20053IntegrationsOnedrivebusiness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.Foldername != nil {
		toSerialize["foldername"] = o.Foldername
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20053IntegrationsOnedrivebusiness struct {
	value *InlineResponse20053IntegrationsOnedrivebusiness
	isSet bool
}

func (v NullableInlineResponse20053IntegrationsOnedrivebusiness) Get() *InlineResponse20053IntegrationsOnedrivebusiness {
	return v.value
}

func (v *NullableInlineResponse20053IntegrationsOnedrivebusiness) Set(val *InlineResponse20053IntegrationsOnedrivebusiness) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053IntegrationsOnedrivebusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053IntegrationsOnedrivebusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053IntegrationsOnedrivebusiness(val *InlineResponse20053IntegrationsOnedrivebusiness) *NullableInlineResponse20053IntegrationsOnedrivebusiness {
	return &NullableInlineResponse20053IntegrationsOnedrivebusiness{value: val, isSet: true}
}

func (v NullableInlineResponse20053IntegrationsOnedrivebusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053IntegrationsOnedrivebusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


