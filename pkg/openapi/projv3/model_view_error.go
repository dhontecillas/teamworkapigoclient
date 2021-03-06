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

// ViewError Error describes an API error.
type ViewError struct {
	// Code is an application-specific error code.
	Code *string `json:"code,omitempty"`
	// Detail is a human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// ID is a reference for this exact instance of the error.
	Id *string `json:"id,omitempty"`
	// Meta contains tags that are useful to detect specific issues.
	Meta *map[string]interface{} `json:"meta,omitempty"`
	// Title is a short, human-readable summary of the problem.
	Title *string `json:"title,omitempty"`
}

// NewViewError instantiates a new ViewError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewError() *ViewError {
	this := ViewError{}
	return &this
}

// NewViewErrorWithDefaults instantiates a new ViewError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewErrorWithDefaults() *ViewError {
	this := ViewError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ViewError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ViewError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ViewError) SetCode(v string) {
	o.Code = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ViewError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ViewError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ViewError) SetDetail(v string) {
	o.Detail = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ViewError) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ViewError) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewError) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ViewError) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ViewError) SetMeta(v map[string]interface{}) {
	o.Meta = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewError) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewError) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewError) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewError) SetTitle(v string) {
	o.Title = &v
}

func (o ViewError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableViewError struct {
	value *ViewError
	isSet bool
}

func (v NullableViewError) Get() *ViewError {
	return v.value
}

func (v *NullableViewError) Set(val *ViewError) {
	v.value = val
	v.isSet = true
}

func (v NullableViewError) IsSet() bool {
	return v.isSet
}

func (v *NullableViewError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewError(val *ViewError) *NullableViewError {
	return &NullableViewError{value: val, isSet: true}
}

func (v NullableViewError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


