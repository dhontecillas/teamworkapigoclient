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

// NotebookVersionsResponse VersionsResponse contains information about a group of notebook versions
type NotebookVersionsResponse struct {
	Included *NotebookVersionResponseIncluded `json:"included,omitempty"`
	Versions *[]ViewNotebookVersion `json:"versions,omitempty"`
}

// NewNotebookVersionsResponse instantiates a new NotebookVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookVersionsResponse() *NotebookVersionsResponse {
	this := NotebookVersionsResponse{}
	return &this
}

// NewNotebookVersionsResponseWithDefaults instantiates a new NotebookVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookVersionsResponseWithDefaults() *NotebookVersionsResponse {
	this := NotebookVersionsResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *NotebookVersionsResponse) GetIncluded() NotebookVersionResponseIncluded {
	if o == nil || o.Included == nil {
		var ret NotebookVersionResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookVersionsResponse) GetIncludedOk() (*NotebookVersionResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *NotebookVersionsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given NotebookVersionResponseIncluded and assigns it to the Included field.
func (o *NotebookVersionsResponse) SetIncluded(v NotebookVersionResponseIncluded) {
	o.Included = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *NotebookVersionsResponse) GetVersions() []ViewNotebookVersion {
	if o == nil || o.Versions == nil {
		var ret []ViewNotebookVersion
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookVersionsResponse) GetVersionsOk() (*[]ViewNotebookVersion, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *NotebookVersionsResponse) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []ViewNotebookVersion and assigns it to the Versions field.
func (o *NotebookVersionsResponse) SetVersions(v []ViewNotebookVersion) {
	o.Versions = &v
}

func (o NotebookVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableNotebookVersionsResponse struct {
	value *NotebookVersionsResponse
	isSet bool
}

func (v NullableNotebookVersionsResponse) Get() *NotebookVersionsResponse {
	return v.value
}

func (v *NullableNotebookVersionsResponse) Set(val *NotebookVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookVersionsResponse(val *NotebookVersionsResponse) *NullableNotebookVersionsResponse {
	return &NullableNotebookVersionsResponse{value: val, isSet: true}
}

func (v NullableNotebookVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


