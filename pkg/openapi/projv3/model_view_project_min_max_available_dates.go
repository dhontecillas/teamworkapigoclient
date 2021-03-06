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

// ViewProjectMinMaxAvailableDates ProjectMinMaxAvailableDates shows suggested start and end dates for a project
type ViewProjectMinMaxAvailableDates struct {
	DeadlinesFound *bool `json:"deadlinesFound,omitempty"`
	MaxEndDate *string `json:"maxEndDate,omitempty"`
	MinStartDate *string `json:"minStartDate,omitempty"`
	SuggestedEndDate *string `json:"suggestedEndDate,omitempty"`
	SuggestedStartDate *string `json:"suggestedStartDate,omitempty"`
}

// NewViewProjectMinMaxAvailableDates instantiates a new ViewProjectMinMaxAvailableDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProjectMinMaxAvailableDates() *ViewProjectMinMaxAvailableDates {
	this := ViewProjectMinMaxAvailableDates{}
	return &this
}

// NewViewProjectMinMaxAvailableDatesWithDefaults instantiates a new ViewProjectMinMaxAvailableDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProjectMinMaxAvailableDatesWithDefaults() *ViewProjectMinMaxAvailableDates {
	this := ViewProjectMinMaxAvailableDates{}
	return &this
}

// GetDeadlinesFound returns the DeadlinesFound field value if set, zero value otherwise.
func (o *ViewProjectMinMaxAvailableDates) GetDeadlinesFound() bool {
	if o == nil || o.DeadlinesFound == nil {
		var ret bool
		return ret
	}
	return *o.DeadlinesFound
}

// GetDeadlinesFoundOk returns a tuple with the DeadlinesFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMinMaxAvailableDates) GetDeadlinesFoundOk() (*bool, bool) {
	if o == nil || o.DeadlinesFound == nil {
		return nil, false
	}
	return o.DeadlinesFound, true
}

// HasDeadlinesFound returns a boolean if a field has been set.
func (o *ViewProjectMinMaxAvailableDates) HasDeadlinesFound() bool {
	if o != nil && o.DeadlinesFound != nil {
		return true
	}

	return false
}

// SetDeadlinesFound gets a reference to the given bool and assigns it to the DeadlinesFound field.
func (o *ViewProjectMinMaxAvailableDates) SetDeadlinesFound(v bool) {
	o.DeadlinesFound = &v
}

// GetMaxEndDate returns the MaxEndDate field value if set, zero value otherwise.
func (o *ViewProjectMinMaxAvailableDates) GetMaxEndDate() string {
	if o == nil || o.MaxEndDate == nil {
		var ret string
		return ret
	}
	return *o.MaxEndDate
}

// GetMaxEndDateOk returns a tuple with the MaxEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMinMaxAvailableDates) GetMaxEndDateOk() (*string, bool) {
	if o == nil || o.MaxEndDate == nil {
		return nil, false
	}
	return o.MaxEndDate, true
}

// HasMaxEndDate returns a boolean if a field has been set.
func (o *ViewProjectMinMaxAvailableDates) HasMaxEndDate() bool {
	if o != nil && o.MaxEndDate != nil {
		return true
	}

	return false
}

// SetMaxEndDate gets a reference to the given string and assigns it to the MaxEndDate field.
func (o *ViewProjectMinMaxAvailableDates) SetMaxEndDate(v string) {
	o.MaxEndDate = &v
}

// GetMinStartDate returns the MinStartDate field value if set, zero value otherwise.
func (o *ViewProjectMinMaxAvailableDates) GetMinStartDate() string {
	if o == nil || o.MinStartDate == nil {
		var ret string
		return ret
	}
	return *o.MinStartDate
}

// GetMinStartDateOk returns a tuple with the MinStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMinMaxAvailableDates) GetMinStartDateOk() (*string, bool) {
	if o == nil || o.MinStartDate == nil {
		return nil, false
	}
	return o.MinStartDate, true
}

// HasMinStartDate returns a boolean if a field has been set.
func (o *ViewProjectMinMaxAvailableDates) HasMinStartDate() bool {
	if o != nil && o.MinStartDate != nil {
		return true
	}

	return false
}

// SetMinStartDate gets a reference to the given string and assigns it to the MinStartDate field.
func (o *ViewProjectMinMaxAvailableDates) SetMinStartDate(v string) {
	o.MinStartDate = &v
}

// GetSuggestedEndDate returns the SuggestedEndDate field value if set, zero value otherwise.
func (o *ViewProjectMinMaxAvailableDates) GetSuggestedEndDate() string {
	if o == nil || o.SuggestedEndDate == nil {
		var ret string
		return ret
	}
	return *o.SuggestedEndDate
}

// GetSuggestedEndDateOk returns a tuple with the SuggestedEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMinMaxAvailableDates) GetSuggestedEndDateOk() (*string, bool) {
	if o == nil || o.SuggestedEndDate == nil {
		return nil, false
	}
	return o.SuggestedEndDate, true
}

// HasSuggestedEndDate returns a boolean if a field has been set.
func (o *ViewProjectMinMaxAvailableDates) HasSuggestedEndDate() bool {
	if o != nil && o.SuggestedEndDate != nil {
		return true
	}

	return false
}

// SetSuggestedEndDate gets a reference to the given string and assigns it to the SuggestedEndDate field.
func (o *ViewProjectMinMaxAvailableDates) SetSuggestedEndDate(v string) {
	o.SuggestedEndDate = &v
}

// GetSuggestedStartDate returns the SuggestedStartDate field value if set, zero value otherwise.
func (o *ViewProjectMinMaxAvailableDates) GetSuggestedStartDate() string {
	if o == nil || o.SuggestedStartDate == nil {
		var ret string
		return ret
	}
	return *o.SuggestedStartDate
}

// GetSuggestedStartDateOk returns a tuple with the SuggestedStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMinMaxAvailableDates) GetSuggestedStartDateOk() (*string, bool) {
	if o == nil || o.SuggestedStartDate == nil {
		return nil, false
	}
	return o.SuggestedStartDate, true
}

// HasSuggestedStartDate returns a boolean if a field has been set.
func (o *ViewProjectMinMaxAvailableDates) HasSuggestedStartDate() bool {
	if o != nil && o.SuggestedStartDate != nil {
		return true
	}

	return false
}

// SetSuggestedStartDate gets a reference to the given string and assigns it to the SuggestedStartDate field.
func (o *ViewProjectMinMaxAvailableDates) SetSuggestedStartDate(v string) {
	o.SuggestedStartDate = &v
}

func (o ViewProjectMinMaxAvailableDates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeadlinesFound != nil {
		toSerialize["deadlinesFound"] = o.DeadlinesFound
	}
	if o.MaxEndDate != nil {
		toSerialize["maxEndDate"] = o.MaxEndDate
	}
	if o.MinStartDate != nil {
		toSerialize["minStartDate"] = o.MinStartDate
	}
	if o.SuggestedEndDate != nil {
		toSerialize["suggestedEndDate"] = o.SuggestedEndDate
	}
	if o.SuggestedStartDate != nil {
		toSerialize["suggestedStartDate"] = o.SuggestedStartDate
	}
	return json.Marshal(toSerialize)
}

type NullableViewProjectMinMaxAvailableDates struct {
	value *ViewProjectMinMaxAvailableDates
	isSet bool
}

func (v NullableViewProjectMinMaxAvailableDates) Get() *ViewProjectMinMaxAvailableDates {
	return v.value
}

func (v *NullableViewProjectMinMaxAvailableDates) Set(val *ViewProjectMinMaxAvailableDates) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProjectMinMaxAvailableDates) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProjectMinMaxAvailableDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProjectMinMaxAvailableDates(val *ViewProjectMinMaxAvailableDates) *NullableViewProjectMinMaxAvailableDates {
	return &NullableViewProjectMinMaxAvailableDates{value: val, isSet: true}
}

func (v NullableViewProjectMinMaxAvailableDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProjectMinMaxAvailableDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


