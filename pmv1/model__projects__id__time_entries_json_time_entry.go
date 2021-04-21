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

// ProjectsIdTimeEntriesJsonTimeEntry struct for ProjectsIdTimeEntriesJsonTimeEntry
type ProjectsIdTimeEntriesJsonTimeEntry struct {
	Date string `json:"date"`
	Description *string `json:"description,omitempty"`
	Hours string `json:"hours"`
	Isbillable *string `json:"isbillable,omitempty"`
	Minutes string `json:"minutes"`
	PersonId *string `json:"person-id,omitempty"`
	Tags *string `json:"tags,omitempty"`
	Time string `json:"time"`
}

// NewProjectsIdTimeEntriesJsonTimeEntry instantiates a new ProjectsIdTimeEntriesJsonTimeEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdTimeEntriesJsonTimeEntry(date string, hours string, minutes string, time string) *ProjectsIdTimeEntriesJsonTimeEntry {
	this := ProjectsIdTimeEntriesJsonTimeEntry{}
	this.Date = date
	this.Hours = hours
	this.Minutes = minutes
	this.Time = time
	return &this
}

// NewProjectsIdTimeEntriesJsonTimeEntryWithDefaults instantiates a new ProjectsIdTimeEntriesJsonTimeEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdTimeEntriesJsonTimeEntryWithDefaults() *ProjectsIdTimeEntriesJsonTimeEntry {
	this := ProjectsIdTimeEntriesJsonTimeEntry{}
	return &this
}

// GetDate returns the Date field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetDate(v string) {
	o.Date = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetDescription(v string) {
	o.Description = &v
}

// GetHours returns the Hours field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetHours() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetHoursOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetHours(v string) {
	o.Hours = v
}

// GetIsbillable returns the Isbillable field value if set, zero value otherwise.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetIsbillable() string {
	if o == nil || o.Isbillable == nil {
		var ret string
		return ret
	}
	return *o.Isbillable
}

// GetIsbillableOk returns a tuple with the Isbillable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetIsbillableOk() (*string, bool) {
	if o == nil || o.Isbillable == nil {
		return nil, false
	}
	return o.Isbillable, true
}

// HasIsbillable returns a boolean if a field has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) HasIsbillable() bool {
	if o != nil && o.Isbillable != nil {
		return true
	}

	return false
}

// SetIsbillable gets a reference to the given string and assigns it to the Isbillable field.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetIsbillable(v string) {
	o.Isbillable = &v
}

// GetMinutes returns the Minutes field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetMinutes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetMinutesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetMinutes(v string) {
	o.Minutes = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetPersonId() string {
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetPersonId(v string) {
	o.PersonId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetTags(v string) {
	o.Tags = &v
}

// GetTime returns the Time field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ProjectsIdTimeEntriesJsonTimeEntry) GetTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ProjectsIdTimeEntriesJsonTimeEntry) SetTime(v string) {
	o.Time = v
}

func (o ProjectsIdTimeEntriesJsonTimeEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["hours"] = o.Hours
	}
	if o.Isbillable != nil {
		toSerialize["isbillable"] = o.Isbillable
	}
	if true {
		toSerialize["minutes"] = o.Minutes
	}
	if o.PersonId != nil {
		toSerialize["person-id"] = o.PersonId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdTimeEntriesJsonTimeEntry struct {
	value *ProjectsIdTimeEntriesJsonTimeEntry
	isSet bool
}

func (v NullableProjectsIdTimeEntriesJsonTimeEntry) Get() *ProjectsIdTimeEntriesJsonTimeEntry {
	return v.value
}

func (v *NullableProjectsIdTimeEntriesJsonTimeEntry) Set(val *ProjectsIdTimeEntriesJsonTimeEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdTimeEntriesJsonTimeEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdTimeEntriesJsonTimeEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdTimeEntriesJsonTimeEntry(val *ProjectsIdTimeEntriesJsonTimeEntry) *NullableProjectsIdTimeEntriesJsonTimeEntry {
	return &NullableProjectsIdTimeEntriesJsonTimeEntry{value: val, isSet: true}
}

func (v NullableProjectsIdTimeEntriesJsonTimeEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdTimeEntriesJsonTimeEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


