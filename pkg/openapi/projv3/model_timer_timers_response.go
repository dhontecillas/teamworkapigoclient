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

// TimerTimersResponse TimersResponse contains all the information returned when sending a GET request to the timers endpoint.
type TimerTimersResponse struct {
	Included *TimerResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	Timers *[]ViewTimer `json:"timers,omitempty"`
}

// NewTimerTimersResponse instantiates a new TimerTimersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerTimersResponse() *TimerTimersResponse {
	this := TimerTimersResponse{}
	return &this
}

// NewTimerTimersResponseWithDefaults instantiates a new TimerTimersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerTimersResponseWithDefaults() *TimerTimersResponse {
	this := TimerTimersResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TimerTimersResponse) GetIncluded() TimerResponseIncluded {
	if o == nil || o.Included == nil {
		var ret TimerResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimersResponse) GetIncludedOk() (*TimerResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TimerTimersResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given TimerResponseIncluded and assigns it to the Included field.
func (o *TimerTimersResponse) SetIncluded(v TimerResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TimerTimersResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimersResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TimerTimersResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *TimerTimersResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetTimers returns the Timers field value if set, zero value otherwise.
func (o *TimerTimersResponse) GetTimers() []ViewTimer {
	if o == nil || o.Timers == nil {
		var ret []ViewTimer
		return ret
	}
	return *o.Timers
}

// GetTimersOk returns a tuple with the Timers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimersResponse) GetTimersOk() (*[]ViewTimer, bool) {
	if o == nil || o.Timers == nil {
		return nil, false
	}
	return o.Timers, true
}

// HasTimers returns a boolean if a field has been set.
func (o *TimerTimersResponse) HasTimers() bool {
	if o != nil && o.Timers != nil {
		return true
	}

	return false
}

// SetTimers gets a reference to the given []ViewTimer and assigns it to the Timers field.
func (o *TimerTimersResponse) SetTimers(v []ViewTimer) {
	o.Timers = &v
}

func (o TimerTimersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Timers != nil {
		toSerialize["timers"] = o.Timers
	}
	return json.Marshal(toSerialize)
}

type NullableTimerTimersResponse struct {
	value *TimerTimersResponse
	isSet bool
}

func (v NullableTimerTimersResponse) Get() *TimerTimersResponse {
	return v.value
}

func (v *NullableTimerTimersResponse) Set(val *TimerTimersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerTimersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerTimersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerTimersResponse(val *TimerTimersResponse) *NullableTimerTimersResponse {
	return &NullableTimerTimersResponse{value: val, isSet: true}
}

func (v NullableTimerTimersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerTimersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


