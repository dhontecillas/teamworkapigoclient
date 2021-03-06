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

// CalendareventsJsonEvent struct for CalendareventsJsonEvent
type CalendareventsJsonEvent struct {
	AllDay *bool `json:"all-day,omitempty"`
	AttendeesCanEdit *bool `json:"attendees-can-edit,omitempty"`
	End *string `json:"end,omitempty"`
	Notify *bool `json:"notify,omitempty"`
	ProjectUsersCanEdit *bool `json:"project-users-can-edit,omitempty"`
	Repeat *CalendareventsJsonEventRepeat `json:"repeat,omitempty"`
	Start *string `json:"start,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewCalendareventsJsonEvent instantiates a new CalendareventsJsonEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendareventsJsonEvent() *CalendareventsJsonEvent {
	this := CalendareventsJsonEvent{}
	return &this
}

// NewCalendareventsJsonEventWithDefaults instantiates a new CalendareventsJsonEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendareventsJsonEventWithDefaults() *CalendareventsJsonEvent {
	this := CalendareventsJsonEvent{}
	return &this
}

// GetAllDay returns the AllDay field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetAllDay() bool {
	if o == nil || o.AllDay == nil {
		var ret bool
		return ret
	}
	return *o.AllDay
}

// GetAllDayOk returns a tuple with the AllDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetAllDayOk() (*bool, bool) {
	if o == nil || o.AllDay == nil {
		return nil, false
	}
	return o.AllDay, true
}

// HasAllDay returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasAllDay() bool {
	if o != nil && o.AllDay != nil {
		return true
	}

	return false
}

// SetAllDay gets a reference to the given bool and assigns it to the AllDay field.
func (o *CalendareventsJsonEvent) SetAllDay(v bool) {
	o.AllDay = &v
}

// GetAttendeesCanEdit returns the AttendeesCanEdit field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetAttendeesCanEdit() bool {
	if o == nil || o.AttendeesCanEdit == nil {
		var ret bool
		return ret
	}
	return *o.AttendeesCanEdit
}

// GetAttendeesCanEditOk returns a tuple with the AttendeesCanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetAttendeesCanEditOk() (*bool, bool) {
	if o == nil || o.AttendeesCanEdit == nil {
		return nil, false
	}
	return o.AttendeesCanEdit, true
}

// HasAttendeesCanEdit returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasAttendeesCanEdit() bool {
	if o != nil && o.AttendeesCanEdit != nil {
		return true
	}

	return false
}

// SetAttendeesCanEdit gets a reference to the given bool and assigns it to the AttendeesCanEdit field.
func (o *CalendareventsJsonEvent) SetAttendeesCanEdit(v bool) {
	o.AttendeesCanEdit = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *CalendareventsJsonEvent) SetEnd(v string) {
	o.End = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetNotify() bool {
	if o == nil || o.Notify == nil {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetNotifyOk() (*bool, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *CalendareventsJsonEvent) SetNotify(v bool) {
	o.Notify = &v
}

// GetProjectUsersCanEdit returns the ProjectUsersCanEdit field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetProjectUsersCanEdit() bool {
	if o == nil || o.ProjectUsersCanEdit == nil {
		var ret bool
		return ret
	}
	return *o.ProjectUsersCanEdit
}

// GetProjectUsersCanEditOk returns a tuple with the ProjectUsersCanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetProjectUsersCanEditOk() (*bool, bool) {
	if o == nil || o.ProjectUsersCanEdit == nil {
		return nil, false
	}
	return o.ProjectUsersCanEdit, true
}

// HasProjectUsersCanEdit returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasProjectUsersCanEdit() bool {
	if o != nil && o.ProjectUsersCanEdit != nil {
		return true
	}

	return false
}

// SetProjectUsersCanEdit gets a reference to the given bool and assigns it to the ProjectUsersCanEdit field.
func (o *CalendareventsJsonEvent) SetProjectUsersCanEdit(v bool) {
	o.ProjectUsersCanEdit = &v
}

// GetRepeat returns the Repeat field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetRepeat() CalendareventsJsonEventRepeat {
	if o == nil || o.Repeat == nil {
		var ret CalendareventsJsonEventRepeat
		return ret
	}
	return *o.Repeat
}

// GetRepeatOk returns a tuple with the Repeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetRepeatOk() (*CalendareventsJsonEventRepeat, bool) {
	if o == nil || o.Repeat == nil {
		return nil, false
	}
	return o.Repeat, true
}

// HasRepeat returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasRepeat() bool {
	if o != nil && o.Repeat != nil {
		return true
	}

	return false
}

// SetRepeat gets a reference to the given CalendareventsJsonEventRepeat and assigns it to the Repeat field.
func (o *CalendareventsJsonEvent) SetRepeat(v CalendareventsJsonEventRepeat) {
	o.Repeat = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *CalendareventsJsonEvent) SetStart(v string) {
	o.Start = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CalendareventsJsonEvent) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendareventsJsonEvent) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CalendareventsJsonEvent) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CalendareventsJsonEvent) SetTitle(v string) {
	o.Title = &v
}

func (o CalendareventsJsonEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllDay != nil {
		toSerialize["all-day"] = o.AllDay
	}
	if o.AttendeesCanEdit != nil {
		toSerialize["attendees-can-edit"] = o.AttendeesCanEdit
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.ProjectUsersCanEdit != nil {
		toSerialize["project-users-can-edit"] = o.ProjectUsersCanEdit
	}
	if o.Repeat != nil {
		toSerialize["repeat"] = o.Repeat
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableCalendareventsJsonEvent struct {
	value *CalendareventsJsonEvent
	isSet bool
}

func (v NullableCalendareventsJsonEvent) Get() *CalendareventsJsonEvent {
	return v.value
}

func (v *NullableCalendareventsJsonEvent) Set(val *CalendareventsJsonEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendareventsJsonEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendareventsJsonEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendareventsJsonEvent(val *CalendareventsJsonEvent) *NullableCalendareventsJsonEvent {
	return &NullableCalendareventsJsonEvent{value: val, isSet: true}
}

func (v NullableCalendareventsJsonEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendareventsJsonEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


