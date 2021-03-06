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

// ViewFilePermissions FilePermissions determine the permissions associated with a file.
type ViewFilePermissions struct {
	CanAddFiles *bool `json:"canAddFiles,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanEdit *bool `json:"canEdit,omitempty"`
}

// NewViewFilePermissions instantiates a new ViewFilePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewFilePermissions() *ViewFilePermissions {
	this := ViewFilePermissions{}
	return &this
}

// NewViewFilePermissionsWithDefaults instantiates a new ViewFilePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewFilePermissionsWithDefaults() *ViewFilePermissions {
	this := ViewFilePermissions{}
	return &this
}

// GetCanAddFiles returns the CanAddFiles field value if set, zero value otherwise.
func (o *ViewFilePermissions) GetCanAddFiles() bool {
	if o == nil || o.CanAddFiles == nil {
		var ret bool
		return ret
	}
	return *o.CanAddFiles
}

// GetCanAddFilesOk returns a tuple with the CanAddFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFilePermissions) GetCanAddFilesOk() (*bool, bool) {
	if o == nil || o.CanAddFiles == nil {
		return nil, false
	}
	return o.CanAddFiles, true
}

// HasCanAddFiles returns a boolean if a field has been set.
func (o *ViewFilePermissions) HasCanAddFiles() bool {
	if o != nil && o.CanAddFiles != nil {
		return true
	}

	return false
}

// SetCanAddFiles gets a reference to the given bool and assigns it to the CanAddFiles field.
func (o *ViewFilePermissions) SetCanAddFiles(v bool) {
	o.CanAddFiles = &v
}

// GetCanDelete returns the CanDelete field value if set, zero value otherwise.
func (o *ViewFilePermissions) GetCanDelete() bool {
	if o == nil || o.CanDelete == nil {
		var ret bool
		return ret
	}
	return *o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFilePermissions) GetCanDeleteOk() (*bool, bool) {
	if o == nil || o.CanDelete == nil {
		return nil, false
	}
	return o.CanDelete, true
}

// HasCanDelete returns a boolean if a field has been set.
func (o *ViewFilePermissions) HasCanDelete() bool {
	if o != nil && o.CanDelete != nil {
		return true
	}

	return false
}

// SetCanDelete gets a reference to the given bool and assigns it to the CanDelete field.
func (o *ViewFilePermissions) SetCanDelete(v bool) {
	o.CanDelete = &v
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *ViewFilePermissions) GetCanEdit() bool {
	if o == nil || o.CanEdit == nil {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFilePermissions) GetCanEditOk() (*bool, bool) {
	if o == nil || o.CanEdit == nil {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *ViewFilePermissions) HasCanEdit() bool {
	if o != nil && o.CanEdit != nil {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *ViewFilePermissions) SetCanEdit(v bool) {
	o.CanEdit = &v
}

func (o ViewFilePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanAddFiles != nil {
		toSerialize["canAddFiles"] = o.CanAddFiles
	}
	if o.CanDelete != nil {
		toSerialize["canDelete"] = o.CanDelete
	}
	if o.CanEdit != nil {
		toSerialize["canEdit"] = o.CanEdit
	}
	return json.Marshal(toSerialize)
}

type NullableViewFilePermissions struct {
	value *ViewFilePermissions
	isSet bool
}

func (v NullableViewFilePermissions) Get() *ViewFilePermissions {
	return v.value
}

func (v *NullableViewFilePermissions) Set(val *ViewFilePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableViewFilePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableViewFilePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewFilePermissions(val *ViewFilePermissions) *NullableViewFilePermissions {
	return &NullableViewFilePermissions{value: val, isSet: true}
}

func (v NullableViewFilePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewFilePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


