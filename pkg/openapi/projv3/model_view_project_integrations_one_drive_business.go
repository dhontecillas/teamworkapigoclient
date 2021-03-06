/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"encoding/json"
)

// ViewProjectIntegrationsOneDriveBusiness struct for ViewProjectIntegrationsOneDriveBusiness
type ViewProjectIntegrationsOneDriveBusiness struct {
	Account *string `json:"account,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Folder *string `json:"folder,omitempty"`
	FolderName *string `json:"folderName,omitempty"`
}

// NewViewProjectIntegrationsOneDriveBusiness instantiates a new ViewProjectIntegrationsOneDriveBusiness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProjectIntegrationsOneDriveBusiness() *ViewProjectIntegrationsOneDriveBusiness {
	this := ViewProjectIntegrationsOneDriveBusiness{}
	return &this
}

// NewViewProjectIntegrationsOneDriveBusinessWithDefaults instantiates a new ViewProjectIntegrationsOneDriveBusiness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProjectIntegrationsOneDriveBusinessWithDefaults() *ViewProjectIntegrationsOneDriveBusiness {
	this := ViewProjectIntegrationsOneDriveBusiness{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ViewProjectIntegrationsOneDriveBusiness) SetAccount(v string) {
	o.Account = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ViewProjectIntegrationsOneDriveBusiness) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *ViewProjectIntegrationsOneDriveBusiness) SetFolder(v string) {
	o.Folder = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetFolderName() string {
	if o == nil || o.FolderName == nil {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) GetFolderNameOk() (*string, bool) {
	if o == nil || o.FolderName == nil {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *ViewProjectIntegrationsOneDriveBusiness) HasFolderName() bool {
	if o != nil && o.FolderName != nil {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *ViewProjectIntegrationsOneDriveBusiness) SetFolderName(v string) {
	o.FolderName = &v
}

func (o ViewProjectIntegrationsOneDriveBusiness) MarshalJSON() ([]byte, error) {
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

type NullableViewProjectIntegrationsOneDriveBusiness struct {
	value *ViewProjectIntegrationsOneDriveBusiness
	isSet bool
}

func (v NullableViewProjectIntegrationsOneDriveBusiness) Get() *ViewProjectIntegrationsOneDriveBusiness {
	return v.value
}

func (v *NullableViewProjectIntegrationsOneDriveBusiness) Set(val *ViewProjectIntegrationsOneDriveBusiness) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProjectIntegrationsOneDriveBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProjectIntegrationsOneDriveBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProjectIntegrationsOneDriveBusiness(val *ViewProjectIntegrationsOneDriveBusiness) *NullableViewProjectIntegrationsOneDriveBusiness {
	return &NullableViewProjectIntegrationsOneDriveBusiness{value: val, isSet: true}
}

func (v NullableViewProjectIntegrationsOneDriveBusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProjectIntegrationsOneDriveBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


