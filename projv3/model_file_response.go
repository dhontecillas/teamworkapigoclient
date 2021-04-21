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

// FileResponse Response contains information about a specific file.
type FileResponse struct {
	File *ViewProjectFile `json:"file,omitempty"`
	Included *FileProjectFilesResponseIncluded `json:"included,omitempty"`
}

// NewFileResponse instantiates a new FileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileResponse() *FileResponse {
	this := FileResponse{}
	return &this
}

// NewFileResponseWithDefaults instantiates a new FileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileResponseWithDefaults() *FileResponse {
	this := FileResponse{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileResponse) GetFile() ViewProjectFile {
	if o == nil || o.File == nil {
		var ret ViewProjectFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResponse) GetFileOk() (*ViewProjectFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileResponse) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given ViewProjectFile and assigns it to the File field.
func (o *FileResponse) SetFile(v ViewProjectFile) {
	o.File = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *FileResponse) GetIncluded() FileProjectFilesResponseIncluded {
	if o == nil || o.Included == nil {
		var ret FileProjectFilesResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResponse) GetIncludedOk() (*FileProjectFilesResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *FileResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given FileProjectFilesResponseIncluded and assigns it to the Included field.
func (o *FileResponse) SetIncluded(v FileProjectFilesResponseIncluded) {
	o.Included = &v
}

func (o FileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableFileResponse struct {
	value *FileResponse
	isSet bool
}

func (v NullableFileResponse) Get() *FileResponse {
	return v.value
}

func (v *NullableFileResponse) Set(val *FileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileResponse(val *FileResponse) *NullableFileResponse {
	return &NullableFileResponse{value: val, isSet: true}
}

func (v NullableFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

