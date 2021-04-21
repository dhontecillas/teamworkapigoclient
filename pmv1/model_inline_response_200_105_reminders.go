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

// InlineResponse200105Reminders struct for InlineResponse200105Reminders
type InlineResponse200105Reminders struct {
	CreatedByUserFirstname *string `json:"created-by-user-firstname,omitempty"`
	CreatedByUserId *string `json:"created-by-user-id,omitempty"`
	CreatedByUserLastname *string `json:"created-by-user-lastname,omitempty"`
	CreatedDateTime *string `json:"created-date-time,omitempty"`
	DateTimeUtc *string `json:"date-time-utc,omitempty"`
	Id *string `json:"id,omitempty"`
	Note *string `json:"note,omitempty"`
	TaskId *string `json:"task-id,omitempty"`
	Type *string `json:"type,omitempty"`
	UserFirstname *string `json:"user-firstname,omitempty"`
	UserId *string `json:"user-id,omitempty"`
	UserLastname *string `json:"user-lastname,omitempty"`
	WasSent *string `json:"was-sent,omitempty"`
}

// NewInlineResponse200105Reminders instantiates a new InlineResponse200105Reminders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200105Reminders() *InlineResponse200105Reminders {
	this := InlineResponse200105Reminders{}
	return &this
}

// NewInlineResponse200105RemindersWithDefaults instantiates a new InlineResponse200105Reminders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200105RemindersWithDefaults() *InlineResponse200105Reminders {
	this := InlineResponse200105Reminders{}
	return &this
}

// GetCreatedByUserFirstname returns the CreatedByUserFirstname field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetCreatedByUserFirstname() string {
	if o == nil || o.CreatedByUserFirstname == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserFirstname
}

// GetCreatedByUserFirstnameOk returns a tuple with the CreatedByUserFirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetCreatedByUserFirstnameOk() (*string, bool) {
	if o == nil || o.CreatedByUserFirstname == nil {
		return nil, false
	}
	return o.CreatedByUserFirstname, true
}

// HasCreatedByUserFirstname returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasCreatedByUserFirstname() bool {
	if o != nil && o.CreatedByUserFirstname != nil {
		return true
	}

	return false
}

// SetCreatedByUserFirstname gets a reference to the given string and assigns it to the CreatedByUserFirstname field.
func (o *InlineResponse200105Reminders) SetCreatedByUserFirstname(v string) {
	o.CreatedByUserFirstname = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetCreatedByUserId() string {
	if o == nil || o.CreatedByUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *InlineResponse200105Reminders) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetCreatedByUserLastname returns the CreatedByUserLastname field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetCreatedByUserLastname() string {
	if o == nil || o.CreatedByUserLastname == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserLastname
}

// GetCreatedByUserLastnameOk returns a tuple with the CreatedByUserLastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetCreatedByUserLastnameOk() (*string, bool) {
	if o == nil || o.CreatedByUserLastname == nil {
		return nil, false
	}
	return o.CreatedByUserLastname, true
}

// HasCreatedByUserLastname returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasCreatedByUserLastname() bool {
	if o != nil && o.CreatedByUserLastname != nil {
		return true
	}

	return false
}

// SetCreatedByUserLastname gets a reference to the given string and assigns it to the CreatedByUserLastname field.
func (o *InlineResponse200105Reminders) SetCreatedByUserLastname(v string) {
	o.CreatedByUserLastname = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetCreatedDateTime() string {
	if o == nil || o.CreatedDateTime == nil {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *InlineResponse200105Reminders) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetDateTimeUtc returns the DateTimeUtc field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetDateTimeUtc() string {
	if o == nil || o.DateTimeUtc == nil {
		var ret string
		return ret
	}
	return *o.DateTimeUtc
}

// GetDateTimeUtcOk returns a tuple with the DateTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetDateTimeUtcOk() (*string, bool) {
	if o == nil || o.DateTimeUtc == nil {
		return nil, false
	}
	return o.DateTimeUtc, true
}

// HasDateTimeUtc returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasDateTimeUtc() bool {
	if o != nil && o.DateTimeUtc != nil {
		return true
	}

	return false
}

// SetDateTimeUtc gets a reference to the given string and assigns it to the DateTimeUtc field.
func (o *InlineResponse200105Reminders) SetDateTimeUtc(v string) {
	o.DateTimeUtc = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200105Reminders) SetId(v string) {
	o.Id = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *InlineResponse200105Reminders) SetNote(v string) {
	o.Note = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *InlineResponse200105Reminders) SetTaskId(v string) {
	o.TaskId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200105Reminders) SetType(v string) {
	o.Type = &v
}

// GetUserFirstname returns the UserFirstname field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetUserFirstname() string {
	if o == nil || o.UserFirstname == nil {
		var ret string
		return ret
	}
	return *o.UserFirstname
}

// GetUserFirstnameOk returns a tuple with the UserFirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetUserFirstnameOk() (*string, bool) {
	if o == nil || o.UserFirstname == nil {
		return nil, false
	}
	return o.UserFirstname, true
}

// HasUserFirstname returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasUserFirstname() bool {
	if o != nil && o.UserFirstname != nil {
		return true
	}

	return false
}

// SetUserFirstname gets a reference to the given string and assigns it to the UserFirstname field.
func (o *InlineResponse200105Reminders) SetUserFirstname(v string) {
	o.UserFirstname = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *InlineResponse200105Reminders) SetUserId(v string) {
	o.UserId = &v
}

// GetUserLastname returns the UserLastname field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetUserLastname() string {
	if o == nil || o.UserLastname == nil {
		var ret string
		return ret
	}
	return *o.UserLastname
}

// GetUserLastnameOk returns a tuple with the UserLastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetUserLastnameOk() (*string, bool) {
	if o == nil || o.UserLastname == nil {
		return nil, false
	}
	return o.UserLastname, true
}

// HasUserLastname returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasUserLastname() bool {
	if o != nil && o.UserLastname != nil {
		return true
	}

	return false
}

// SetUserLastname gets a reference to the given string and assigns it to the UserLastname field.
func (o *InlineResponse200105Reminders) SetUserLastname(v string) {
	o.UserLastname = &v
}

// GetWasSent returns the WasSent field value if set, zero value otherwise.
func (o *InlineResponse200105Reminders) GetWasSent() string {
	if o == nil || o.WasSent == nil {
		var ret string
		return ret
	}
	return *o.WasSent
}

// GetWasSentOk returns a tuple with the WasSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105Reminders) GetWasSentOk() (*string, bool) {
	if o == nil || o.WasSent == nil {
		return nil, false
	}
	return o.WasSent, true
}

// HasWasSent returns a boolean if a field has been set.
func (o *InlineResponse200105Reminders) HasWasSent() bool {
	if o != nil && o.WasSent != nil {
		return true
	}

	return false
}

// SetWasSent gets a reference to the given string and assigns it to the WasSent field.
func (o *InlineResponse200105Reminders) SetWasSent(v string) {
	o.WasSent = &v
}

func (o InlineResponse200105Reminders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedByUserFirstname != nil {
		toSerialize["created-by-user-firstname"] = o.CreatedByUserFirstname
	}
	if o.CreatedByUserId != nil {
		toSerialize["created-by-user-id"] = o.CreatedByUserId
	}
	if o.CreatedByUserLastname != nil {
		toSerialize["created-by-user-lastname"] = o.CreatedByUserLastname
	}
	if o.CreatedDateTime != nil {
		toSerialize["created-date-time"] = o.CreatedDateTime
	}
	if o.DateTimeUtc != nil {
		toSerialize["date-time-utc"] = o.DateTimeUtc
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.TaskId != nil {
		toSerialize["task-id"] = o.TaskId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UserFirstname != nil {
		toSerialize["user-firstname"] = o.UserFirstname
	}
	if o.UserId != nil {
		toSerialize["user-id"] = o.UserId
	}
	if o.UserLastname != nil {
		toSerialize["user-lastname"] = o.UserLastname
	}
	if o.WasSent != nil {
		toSerialize["was-sent"] = o.WasSent
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200105Reminders struct {
	value *InlineResponse200105Reminders
	isSet bool
}

func (v NullableInlineResponse200105Reminders) Get() *InlineResponse200105Reminders {
	return v.value
}

func (v *NullableInlineResponse200105Reminders) Set(val *InlineResponse200105Reminders) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200105Reminders) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200105Reminders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200105Reminders(val *InlineResponse200105Reminders) *NullableInlineResponse200105Reminders {
	return &NullableInlineResponse200105Reminders{value: val, isSet: true}
}

func (v NullableInlineResponse200105Reminders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200105Reminders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


