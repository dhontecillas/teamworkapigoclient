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

// FormFormsResponse FormsResponse contains information about a group of forms.
type FormFormsResponse struct {
	Forms *[]ViewForm `json:"forms,omitempty"`
	Included *CalendarEventsResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
}

// NewFormFormsResponse instantiates a new FormFormsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFormsResponse() *FormFormsResponse {
	this := FormFormsResponse{}
	return &this
}

// NewFormFormsResponseWithDefaults instantiates a new FormFormsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFormsResponseWithDefaults() *FormFormsResponse {
	this := FormFormsResponse{}
	return &this
}

// GetForms returns the Forms field value if set, zero value otherwise.
func (o *FormFormsResponse) GetForms() []ViewForm {
	if o == nil || o.Forms == nil {
		var ret []ViewForm
		return ret
	}
	return *o.Forms
}

// GetFormsOk returns a tuple with the Forms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFormsResponse) GetFormsOk() (*[]ViewForm, bool) {
	if o == nil || o.Forms == nil {
		return nil, false
	}
	return o.Forms, true
}

// HasForms returns a boolean if a field has been set.
func (o *FormFormsResponse) HasForms() bool {
	if o != nil && o.Forms != nil {
		return true
	}

	return false
}

// SetForms gets a reference to the given []ViewForm and assigns it to the Forms field.
func (o *FormFormsResponse) SetForms(v []ViewForm) {
	o.Forms = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *FormFormsResponse) GetIncluded() CalendarEventsResponseIncluded {
	if o == nil || o.Included == nil {
		var ret CalendarEventsResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFormsResponse) GetIncludedOk() (*CalendarEventsResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *FormFormsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given CalendarEventsResponseIncluded and assigns it to the Included field.
func (o *FormFormsResponse) SetIncluded(v CalendarEventsResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FormFormsResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFormsResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FormFormsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *FormFormsResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

func (o FormFormsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Forms != nil {
		toSerialize["forms"] = o.Forms
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableFormFormsResponse struct {
	value *FormFormsResponse
	isSet bool
}

func (v NullableFormFormsResponse) Get() *FormFormsResponse {
	return v.value
}

func (v *NullableFormFormsResponse) Set(val *FormFormsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFormsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFormsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFormsResponse(val *FormFormsResponse) *NullableFormFormsResponse {
	return &NullableFormFormsResponse{value: val, isSet: true}
}

func (v NullableFormFormsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFormsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


