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

// CalendareventsIdJsonEventReminders struct for CalendareventsIdJsonEventReminders
type CalendareventsIdJsonEventReminders struct {
	Before *string `json:"before,omitempty"`
	Period *string `json:"period,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCalendareventsIdJsonEventReminders instantiates a new CalendareventsIdJsonEventReminders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendareventsIdJsonEventReminders() *CalendareventsIdJsonEventReminders {
	this := CalendareventsIdJsonEventReminders{}
	return &this
}

// NewCalendareventsIdJsonEventRemindersWithDefaults instantiates a new CalendareventsIdJsonEventReminders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendareventsIdJsonEventRemindersWithDefaults() *CalendareventsIdJsonEventReminders {
	this := CalendareventsIdJsonEventReminders{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *CalendareventsIdJsonEventReminders) GetBefore() string {
	if o == nil || o.Before == nil {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsIdJsonEventReminders) GetBeforeOk() (*string, bool) {
	if o == nil || o.Before == nil {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *CalendareventsIdJsonEventReminders) HasBefore() bool {
	if o != nil && o.Before != nil {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *CalendareventsIdJsonEventReminders) SetBefore(v string) {
	o.Before = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CalendareventsIdJsonEventReminders) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsIdJsonEventReminders) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CalendareventsIdJsonEventReminders) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *CalendareventsIdJsonEventReminders) SetPeriod(v string) {
	o.Period = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CalendareventsIdJsonEventReminders) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsIdJsonEventReminders) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CalendareventsIdJsonEventReminders) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CalendareventsIdJsonEventReminders) SetType(v string) {
	o.Type = &v
}

func (o CalendareventsIdJsonEventReminders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Before != nil {
		toSerialize["before"] = o.Before
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCalendareventsIdJsonEventReminders struct {
	value *CalendareventsIdJsonEventReminders
	isSet bool
}

func (v NullableCalendareventsIdJsonEventReminders) Get() *CalendareventsIdJsonEventReminders {
	return v.value
}

func (v *NullableCalendareventsIdJsonEventReminders) Set(val *CalendareventsIdJsonEventReminders) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendareventsIdJsonEventReminders) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendareventsIdJsonEventReminders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendareventsIdJsonEventReminders(val *CalendareventsIdJsonEventReminders) *NullableCalendareventsIdJsonEventReminders {
	return &NullableCalendareventsIdJsonEventReminders{value: val, isSet: true}
}

func (v NullableCalendareventsIdJsonEventReminders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendareventsIdJsonEventReminders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

