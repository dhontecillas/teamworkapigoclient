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

// FileRequest Request contains information of a file to be created or updated.
type FileRequest struct {
	File *FileProjectFile `json:"file,omitempty"`
}

// NewFileRequest instantiates a new FileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileRequest() *FileRequest {
	this := FileRequest{}
	return &this
}

// NewFileRequestWithDefaults instantiates a new FileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileRequestWithDefaults() *FileRequest {
	this := FileRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileRequest) GetFile() FileProjectFile {
	if o == nil || o.File == nil {
		var ret FileProjectFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetFileOk() (*FileProjectFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileRequest) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given FileProjectFile and assigns it to the File field.
func (o *FileRequest) SetFile(v FileProjectFile) {
	o.File = &v
}

func (o FileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableFileRequest struct {
	value *FileRequest
	isSet bool
}

func (v NullableFileRequest) Get() *FileRequest {
	return v.value
}

func (v *NullableFileRequest) Set(val *FileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileRequest(val *FileRequest) *NullableFileRequest {
	return &NullableFileRequest{value: val, isSet: true}
}

func (v NullableFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


