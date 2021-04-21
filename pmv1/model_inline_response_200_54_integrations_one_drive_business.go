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

// InlineResponse20054IntegrationsOneDriveBusiness struct for InlineResponse20054IntegrationsOneDriveBusiness
type InlineResponse20054IntegrationsOneDriveBusiness struct {
	Account *string `json:"account,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Folder *string `json:"folder,omitempty"`
	FolderName *string `json:"folderName,omitempty"`
}

// NewInlineResponse20054IntegrationsOneDriveBusiness instantiates a new InlineResponse20054IntegrationsOneDriveBusiness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054IntegrationsOneDriveBusiness() *InlineResponse20054IntegrationsOneDriveBusiness {
	this := InlineResponse20054IntegrationsOneDriveBusiness{}
	return &this
}

// NewInlineResponse20054IntegrationsOneDriveBusinessWithDefaults instantiates a new InlineResponse20054IntegrationsOneDriveBusiness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054IntegrationsOneDriveBusinessWithDefaults() *InlineResponse20054IntegrationsOneDriveBusiness {
	this := InlineResponse20054IntegrationsOneDriveBusiness{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) SetAccount(v string) {
	o.Account = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) SetFolder(v string) {
	o.Folder = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetFolderName() string {
	if o == nil || o.FolderName == nil {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) GetFolderNameOk() (*string, bool) {
	if o == nil || o.FolderName == nil {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) HasFolderName() bool {
	if o != nil && o.FolderName != nil {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *InlineResponse20054IntegrationsOneDriveBusiness) SetFolderName(v string) {
	o.FolderName = &v
}

func (o InlineResponse20054IntegrationsOneDriveBusiness) MarshalJSON() ([]byte, error) {
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
	if o.FolderName != nil {
		toSerialize["folderName"] = o.FolderName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054IntegrationsOneDriveBusiness struct {
	value *InlineResponse20054IntegrationsOneDriveBusiness
	isSet bool
}

func (v NullableInlineResponse20054IntegrationsOneDriveBusiness) Get() *InlineResponse20054IntegrationsOneDriveBusiness {
	return v.value
}

func (v *NullableInlineResponse20054IntegrationsOneDriveBusiness) Set(val *InlineResponse20054IntegrationsOneDriveBusiness) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054IntegrationsOneDriveBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054IntegrationsOneDriveBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054IntegrationsOneDriveBusiness(val *InlineResponse20054IntegrationsOneDriveBusiness) *NullableInlineResponse20054IntegrationsOneDriveBusiness {
	return &NullableInlineResponse20054IntegrationsOneDriveBusiness{value: val, isSet: true}
}

func (v NullableInlineResponse20054IntegrationsOneDriveBusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054IntegrationsOneDriveBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

