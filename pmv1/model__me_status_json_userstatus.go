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

// MeStatusJsonUserstatus struct for MeStatusJsonUserstatus
type MeStatusJsonUserstatus struct {
	Notify *string `json:"notify,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewMeStatusJsonUserstatus instantiates a new MeStatusJsonUserstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeStatusJsonUserstatus() *MeStatusJsonUserstatus {
	this := MeStatusJsonUserstatus{}
	return &this
}

// NewMeStatusJsonUserstatusWithDefaults instantiates a new MeStatusJsonUserstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeStatusJsonUserstatusWithDefaults() *MeStatusJsonUserstatus {
	this := MeStatusJsonUserstatus{}
	return &this
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *MeStatusJsonUserstatus) GetNotify() string {
	if o == nil || o.Notify == nil {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeStatusJsonUserstatus) GetNotifyOk() (*string, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *MeStatusJsonUserstatus) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *MeStatusJsonUserstatus) SetNotify(v string) {
	o.Notify = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MeStatusJsonUserstatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeStatusJsonUserstatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MeStatusJsonUserstatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MeStatusJsonUserstatus) SetStatus(v string) {
	o.Status = &v
}

func (o MeStatusJsonUserstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMeStatusJsonUserstatus struct {
	value *MeStatusJsonUserstatus
	isSet bool
}

func (v NullableMeStatusJsonUserstatus) Get() *MeStatusJsonUserstatus {
	return v.value
}

func (v *NullableMeStatusJsonUserstatus) Set(val *MeStatusJsonUserstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMeStatusJsonUserstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMeStatusJsonUserstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeStatusJsonUserstatus(val *MeStatusJsonUserstatus) *NullableMeStatusJsonUserstatus {
	return &NullableMeStatusJsonUserstatus{value: val, isSet: true}
}

func (v NullableMeStatusJsonUserstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeStatusJsonUserstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


