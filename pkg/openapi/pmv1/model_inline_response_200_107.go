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

// InlineResponse200107 struct for InlineResponse200107
type InlineResponse200107 struct {
	STATUS *string `json:"STATUS,omitempty"`
	TimeEntries *[]InlineResponse200107TimeEntries `json:"time-entries,omitempty"`
}

// NewInlineResponse200107 instantiates a new InlineResponse200107 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200107() *InlineResponse200107 {
	this := InlineResponse200107{}
	return &this
}

// NewInlineResponse200107WithDefaults instantiates a new InlineResponse200107 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200107WithDefaults() *InlineResponse200107 {
	this := InlineResponse200107{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200107) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200107) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200107) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTimeEntries returns the TimeEntries field value if set, zero value otherwise.
func (o *InlineResponse200107) GetTimeEntries() []InlineResponse200107TimeEntries {
	if o == nil || o.TimeEntries == nil {
		var ret []InlineResponse200107TimeEntries
		return ret
	}
	return *o.TimeEntries
}

// GetTimeEntriesOk returns a tuple with the TimeEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetTimeEntriesOk() (*[]InlineResponse200107TimeEntries, bool) {
	if o == nil || o.TimeEntries == nil {
		return nil, false
	}
	return o.TimeEntries, true
}

// HasTimeEntries returns a boolean if a field has been set.
func (o *InlineResponse200107) HasTimeEntries() bool {
	if o != nil && o.TimeEntries != nil {
		return true
	}

	return false
}

// SetTimeEntries gets a reference to the given []InlineResponse200107TimeEntries and assigns it to the TimeEntries field.
func (o *InlineResponse200107) SetTimeEntries(v []InlineResponse200107TimeEntries) {
	o.TimeEntries = &v
}

func (o InlineResponse200107) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.TimeEntries != nil {
		toSerialize["time-entries"] = o.TimeEntries
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200107 struct {
	value *InlineResponse200107
	isSet bool
}

func (v NullableInlineResponse200107) Get() *InlineResponse200107 {
	return v.value
}

func (v *NullableInlineResponse200107) Set(val *InlineResponse200107) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200107) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200107) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200107(val *InlineResponse200107) *NullableInlineResponse200107 {
	return &NullableInlineResponse200107{value: val, isSet: true}
}

func (v NullableInlineResponse200107) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200107) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


