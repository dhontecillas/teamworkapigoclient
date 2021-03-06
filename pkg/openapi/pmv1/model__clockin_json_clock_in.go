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

// ClockinJsonClockIn struct for ClockinJsonClockIn
type ClockinJsonClockIn struct {
	ClockInDatetime *string `json:"clockInDatetime,omitempty"`
	ClockInInfo *string `json:"clockInInfo,omitempty"`
	ClockOutDatetime *string `json:"clockOutDatetime,omitempty"`
	ClockOutInfo *string `json:"clockOutInfo,omitempty"`
	UserId int32 `json:"userId"`
}

// NewClockinJsonClockIn instantiates a new ClockinJsonClockIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClockinJsonClockIn(userId int32) *ClockinJsonClockIn {
	this := ClockinJsonClockIn{}
	this.UserId = userId
	return &this
}

// NewClockinJsonClockInWithDefaults instantiates a new ClockinJsonClockIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClockinJsonClockInWithDefaults() *ClockinJsonClockIn {
	this := ClockinJsonClockIn{}
	return &this
}

// GetClockInDatetime returns the ClockInDatetime field value if set, zero value otherwise.
func (o *ClockinJsonClockIn) GetClockInDatetime() string {
	if o == nil || o.ClockInDatetime == nil {
		var ret string
		return ret
	}
	return *o.ClockInDatetime
}

// GetClockInDatetimeOk returns a tuple with the ClockInDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockinJsonClockIn) GetClockInDatetimeOk() (*string, bool) {
	if o == nil || o.ClockInDatetime == nil {
		return nil, false
	}
	return o.ClockInDatetime, true
}

// HasClockInDatetime returns a boolean if a field has been set.
func (o *ClockinJsonClockIn) HasClockInDatetime() bool {
	if o != nil && o.ClockInDatetime != nil {
		return true
	}

	return false
}

// SetClockInDatetime gets a reference to the given string and assigns it to the ClockInDatetime field.
func (o *ClockinJsonClockIn) SetClockInDatetime(v string) {
	o.ClockInDatetime = &v
}

// GetClockInInfo returns the ClockInInfo field value if set, zero value otherwise.
func (o *ClockinJsonClockIn) GetClockInInfo() string {
	if o == nil || o.ClockInInfo == nil {
		var ret string
		return ret
	}
	return *o.ClockInInfo
}

// GetClockInInfoOk returns a tuple with the ClockInInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockinJsonClockIn) GetClockInInfoOk() (*string, bool) {
	if o == nil || o.ClockInInfo == nil {
		return nil, false
	}
	return o.ClockInInfo, true
}

// HasClockInInfo returns a boolean if a field has been set.
func (o *ClockinJsonClockIn) HasClockInInfo() bool {
	if o != nil && o.ClockInInfo != nil {
		return true
	}

	return false
}

// SetClockInInfo gets a reference to the given string and assigns it to the ClockInInfo field.
func (o *ClockinJsonClockIn) SetClockInInfo(v string) {
	o.ClockInInfo = &v
}

// GetClockOutDatetime returns the ClockOutDatetime field value if set, zero value otherwise.
func (o *ClockinJsonClockIn) GetClockOutDatetime() string {
	if o == nil || o.ClockOutDatetime == nil {
		var ret string
		return ret
	}
	return *o.ClockOutDatetime
}

// GetClockOutDatetimeOk returns a tuple with the ClockOutDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockinJsonClockIn) GetClockOutDatetimeOk() (*string, bool) {
	if o == nil || o.ClockOutDatetime == nil {
		return nil, false
	}
	return o.ClockOutDatetime, true
}

// HasClockOutDatetime returns a boolean if a field has been set.
func (o *ClockinJsonClockIn) HasClockOutDatetime() bool {
	if o != nil && o.ClockOutDatetime != nil {
		return true
	}

	return false
}

// SetClockOutDatetime gets a reference to the given string and assigns it to the ClockOutDatetime field.
func (o *ClockinJsonClockIn) SetClockOutDatetime(v string) {
	o.ClockOutDatetime = &v
}

// GetClockOutInfo returns the ClockOutInfo field value if set, zero value otherwise.
func (o *ClockinJsonClockIn) GetClockOutInfo() string {
	if o == nil || o.ClockOutInfo == nil {
		var ret string
		return ret
	}
	return *o.ClockOutInfo
}

// GetClockOutInfoOk returns a tuple with the ClockOutInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockinJsonClockIn) GetClockOutInfoOk() (*string, bool) {
	if o == nil || o.ClockOutInfo == nil {
		return nil, false
	}
	return o.ClockOutInfo, true
}

// HasClockOutInfo returns a boolean if a field has been set.
func (o *ClockinJsonClockIn) HasClockOutInfo() bool {
	if o != nil && o.ClockOutInfo != nil {
		return true
	}

	return false
}

// SetClockOutInfo gets a reference to the given string and assigns it to the ClockOutInfo field.
func (o *ClockinJsonClockIn) SetClockOutInfo(v string) {
	o.ClockOutInfo = &v
}

// GetUserId returns the UserId field value
func (o *ClockinJsonClockIn) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ClockinJsonClockIn) GetUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ClockinJsonClockIn) SetUserId(v int32) {
	o.UserId = v
}

func (o ClockinJsonClockIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClockInDatetime != nil {
		toSerialize["clockInDatetime"] = o.ClockInDatetime
	}
	if o.ClockInInfo != nil {
		toSerialize["clockInInfo"] = o.ClockInInfo
	}
	if o.ClockOutDatetime != nil {
		toSerialize["clockOutDatetime"] = o.ClockOutDatetime
	}
	if o.ClockOutInfo != nil {
		toSerialize["clockOutInfo"] = o.ClockOutInfo
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableClockinJsonClockIn struct {
	value *ClockinJsonClockIn
	isSet bool
}

func (v NullableClockinJsonClockIn) Get() *ClockinJsonClockIn {
	return v.value
}

func (v *NullableClockinJsonClockIn) Set(val *ClockinJsonClockIn) {
	v.value = val
	v.isSet = true
}

func (v NullableClockinJsonClockIn) IsSet() bool {
	return v.isSet
}

func (v *NullableClockinJsonClockIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClockinJsonClockIn(val *ClockinJsonClockIn) *NullableClockinJsonClockIn {
	return &NullableClockinJsonClockIn{value: val, isSet: true}
}

func (v NullableClockinJsonClockIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClockinJsonClockIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


