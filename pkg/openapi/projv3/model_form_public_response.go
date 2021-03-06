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

// FormPublicResponse PublicResponse contains information about a specific public form.
type FormPublicResponse struct {
	Form *ViewPublicForm `json:"form,omitempty"`
}

// NewFormPublicResponse instantiates a new FormPublicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormPublicResponse() *FormPublicResponse {
	this := FormPublicResponse{}
	return &this
}

// NewFormPublicResponseWithDefaults instantiates a new FormPublicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormPublicResponseWithDefaults() *FormPublicResponse {
	this := FormPublicResponse{}
	return &this
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *FormPublicResponse) GetForm() ViewPublicForm {
	if o == nil || o.Form == nil {
		var ret ViewPublicForm
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormPublicResponse) GetFormOk() (*ViewPublicForm, bool) {
	if o == nil || o.Form == nil {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *FormPublicResponse) HasForm() bool {
	if o != nil && o.Form != nil {
		return true
	}

	return false
}

// SetForm gets a reference to the given ViewPublicForm and assigns it to the Form field.
func (o *FormPublicResponse) SetForm(v ViewPublicForm) {
	o.Form = &v
}

func (o FormPublicResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	return json.Marshal(toSerialize)
}

type NullableFormPublicResponse struct {
	value *FormPublicResponse
	isSet bool
}

func (v NullableFormPublicResponse) Get() *FormPublicResponse {
	return v.value
}

func (v *NullableFormPublicResponse) Set(val *FormPublicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFormPublicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFormPublicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormPublicResponse(val *FormPublicResponse) *NullableFormPublicResponse {
	return &NullableFormPublicResponse{value: val, isSet: true}
}

func (v NullableFormPublicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormPublicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


