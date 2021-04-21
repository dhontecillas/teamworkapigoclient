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

// LateResponse Response contains the count of late tasks. Following this format to satisfy the Numerics integration.
type LateResponse struct {
	Data *LateTaskMetricLate `json:"data,omitempty"`
}

// NewLateResponse instantiates a new LateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLateResponse() *LateResponse {
	this := LateResponse{}
	return &this
}

// NewLateResponseWithDefaults instantiates a new LateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLateResponseWithDefaults() *LateResponse {
	this := LateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LateResponse) GetData() LateTaskMetricLate {
	if o == nil || o.Data == nil {
		var ret LateTaskMetricLate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LateResponse) GetDataOk() (*LateTaskMetricLate, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LateResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LateTaskMetricLate and assigns it to the Data field.
func (o *LateResponse) SetData(v LateTaskMetricLate) {
	o.Data = &v
}

func (o LateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLateResponse struct {
	value *LateResponse
	isSet bool
}

func (v NullableLateResponse) Get() *LateResponse {
	return v.value
}

func (v *NullableLateResponse) Set(val *LateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLateResponse(val *LateResponse) *NullableLateResponse {
	return &NullableLateResponse{value: val, isSet: true}
}

func (v NullableLateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


