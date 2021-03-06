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

// ScheduleUserSchedulesResponse UserSchedulesResponse contains information about a group of schedules.
type ScheduleUserSchedulesResponse struct {
	Included *ScheduleUserSchedulesResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	UserSchedules *[]ViewUserSchedule `json:"userSchedules,omitempty"`
}

// NewScheduleUserSchedulesResponse instantiates a new ScheduleUserSchedulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleUserSchedulesResponse() *ScheduleUserSchedulesResponse {
	this := ScheduleUserSchedulesResponse{}
	return &this
}

// NewScheduleUserSchedulesResponseWithDefaults instantiates a new ScheduleUserSchedulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleUserSchedulesResponseWithDefaults() *ScheduleUserSchedulesResponse {
	this := ScheduleUserSchedulesResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ScheduleUserSchedulesResponse) GetIncluded() ScheduleUserSchedulesResponseIncluded {
	if o == nil || o.Included == nil {
		var ret ScheduleUserSchedulesResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUserSchedulesResponse) GetIncludedOk() (*ScheduleUserSchedulesResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ScheduleUserSchedulesResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given ScheduleUserSchedulesResponseIncluded and assigns it to the Included field.
func (o *ScheduleUserSchedulesResponse) SetIncluded(v ScheduleUserSchedulesResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScheduleUserSchedulesResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUserSchedulesResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScheduleUserSchedulesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *ScheduleUserSchedulesResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetUserSchedules returns the UserSchedules field value if set, zero value otherwise.
func (o *ScheduleUserSchedulesResponse) GetUserSchedules() []ViewUserSchedule {
	if o == nil || o.UserSchedules == nil {
		var ret []ViewUserSchedule
		return ret
	}
	return *o.UserSchedules
}

// GetUserSchedulesOk returns a tuple with the UserSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUserSchedulesResponse) GetUserSchedulesOk() (*[]ViewUserSchedule, bool) {
	if o == nil || o.UserSchedules == nil {
		return nil, false
	}
	return o.UserSchedules, true
}

// HasUserSchedules returns a boolean if a field has been set.
func (o *ScheduleUserSchedulesResponse) HasUserSchedules() bool {
	if o != nil && o.UserSchedules != nil {
		return true
	}

	return false
}

// SetUserSchedules gets a reference to the given []ViewUserSchedule and assigns it to the UserSchedules field.
func (o *ScheduleUserSchedulesResponse) SetUserSchedules(v []ViewUserSchedule) {
	o.UserSchedules = &v
}

func (o ScheduleUserSchedulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.UserSchedules != nil {
		toSerialize["userSchedules"] = o.UserSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleUserSchedulesResponse struct {
	value *ScheduleUserSchedulesResponse
	isSet bool
}

func (v NullableScheduleUserSchedulesResponse) Get() *ScheduleUserSchedulesResponse {
	return v.value
}

func (v *NullableScheduleUserSchedulesResponse) Set(val *ScheduleUserSchedulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleUserSchedulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleUserSchedulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleUserSchedulesResponse(val *ScheduleUserSchedulesResponse) *NullableScheduleUserSchedulesResponse {
	return &NullableScheduleUserSchedulesResponse{value: val, isSet: true}
}

func (v NullableScheduleUserSchedulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleUserSchedulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


