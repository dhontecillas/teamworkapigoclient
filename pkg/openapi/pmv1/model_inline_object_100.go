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

// InlineObject100 struct for InlineObject100
type InlineObject100 struct {
	Reminder *TasksIdRemindersJsonReminder `json:"reminder,omitempty"`
}

// NewInlineObject100 instantiates a new InlineObject100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject100() *InlineObject100 {
	this := InlineObject100{}
	return &this
}

// NewInlineObject100WithDefaults instantiates a new InlineObject100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject100WithDefaults() *InlineObject100 {
	this := InlineObject100{}
	return &this
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *InlineObject100) GetReminder() TasksIdRemindersJsonReminder {
	if o == nil || o.Reminder == nil {
		var ret TasksIdRemindersJsonReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetReminderOk() (*TasksIdRemindersJsonReminder, bool) {
	if o == nil || o.Reminder == nil {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *InlineObject100) HasReminder() bool {
	if o != nil && o.Reminder != nil {
		return true
	}

	return false
}

// SetReminder gets a reference to the given TasksIdRemindersJsonReminder and assigns it to the Reminder field.
func (o *InlineObject100) SetReminder(v TasksIdRemindersJsonReminder) {
	o.Reminder = &v
}

func (o InlineObject100) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reminder != nil {
		toSerialize["reminder"] = o.Reminder
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject100 struct {
	value *InlineObject100
	isSet bool
}

func (v NullableInlineObject100) Get() *InlineObject100 {
	return v.value
}

func (v *NullableInlineObject100) Set(val *InlineObject100) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject100) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject100(val *InlineObject100) *NullableInlineObject100 {
	return &NullableInlineObject100{value: val, isSet: true}
}

func (v NullableInlineObject100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


