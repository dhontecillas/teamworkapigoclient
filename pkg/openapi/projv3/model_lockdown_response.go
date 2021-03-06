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

// LockdownResponse Response contains information about a specific lockdown.
type LockdownResponse struct {
	Included *LockdownResponseIncluded `json:"included,omitempty"`
	Lockdown *ViewLockdown `json:"lockdown,omitempty"`
}

// NewLockdownResponse instantiates a new LockdownResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockdownResponse() *LockdownResponse {
	this := LockdownResponse{}
	return &this
}

// NewLockdownResponseWithDefaults instantiates a new LockdownResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockdownResponseWithDefaults() *LockdownResponse {
	this := LockdownResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *LockdownResponse) GetIncluded() LockdownResponseIncluded {
	if o == nil || o.Included == nil {
		var ret LockdownResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockdownResponse) GetIncludedOk() (*LockdownResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *LockdownResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given LockdownResponseIncluded and assigns it to the Included field.
func (o *LockdownResponse) SetIncluded(v LockdownResponseIncluded) {
	o.Included = &v
}

// GetLockdown returns the Lockdown field value if set, zero value otherwise.
func (o *LockdownResponse) GetLockdown() ViewLockdown {
	if o == nil || o.Lockdown == nil {
		var ret ViewLockdown
		return ret
	}
	return *o.Lockdown
}

// GetLockdownOk returns a tuple with the Lockdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockdownResponse) GetLockdownOk() (*ViewLockdown, bool) {
	if o == nil || o.Lockdown == nil {
		return nil, false
	}
	return o.Lockdown, true
}

// HasLockdown returns a boolean if a field has been set.
func (o *LockdownResponse) HasLockdown() bool {
	if o != nil && o.Lockdown != nil {
		return true
	}

	return false
}

// SetLockdown gets a reference to the given ViewLockdown and assigns it to the Lockdown field.
func (o *LockdownResponse) SetLockdown(v ViewLockdown) {
	o.Lockdown = &v
}

func (o LockdownResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Lockdown != nil {
		toSerialize["lockdown"] = o.Lockdown
	}
	return json.Marshal(toSerialize)
}

type NullableLockdownResponse struct {
	value *LockdownResponse
	isSet bool
}

func (v NullableLockdownResponse) Get() *LockdownResponse {
	return v.value
}

func (v *NullableLockdownResponse) Set(val *LockdownResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLockdownResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLockdownResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockdownResponse(val *LockdownResponse) *NullableLockdownResponse {
	return &NullableLockdownResponse{value: val, isSet: true}
}

func (v NullableLockdownResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockdownResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


