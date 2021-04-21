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

// SummaryMilestoneResponse MilestoneResponse contains groups of counter for milestones.
type SummaryMilestoneResponse struct {
	Everyone *SummaryMilestoneCountsResponse `json:"everyone,omitempty"`
	Mine *SummaryMilestoneCountsResponse `json:"mine,omitempty"`
}

// NewSummaryMilestoneResponse instantiates a new SummaryMilestoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryMilestoneResponse() *SummaryMilestoneResponse {
	this := SummaryMilestoneResponse{}
	return &this
}

// NewSummaryMilestoneResponseWithDefaults instantiates a new SummaryMilestoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryMilestoneResponseWithDefaults() *SummaryMilestoneResponse {
	this := SummaryMilestoneResponse{}
	return &this
}

// GetEveryone returns the Everyone field value if set, zero value otherwise.
func (o *SummaryMilestoneResponse) GetEveryone() SummaryMilestoneCountsResponse {
	if o == nil || o.Everyone == nil {
		var ret SummaryMilestoneCountsResponse
		return ret
	}
	return *o.Everyone
}

// GetEveryoneOk returns a tuple with the Everyone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryMilestoneResponse) GetEveryoneOk() (*SummaryMilestoneCountsResponse, bool) {
	if o == nil || o.Everyone == nil {
		return nil, false
	}
	return o.Everyone, true
}

// HasEveryone returns a boolean if a field has been set.
func (o *SummaryMilestoneResponse) HasEveryone() bool {
	if o != nil && o.Everyone != nil {
		return true
	}

	return false
}

// SetEveryone gets a reference to the given SummaryMilestoneCountsResponse and assigns it to the Everyone field.
func (o *SummaryMilestoneResponse) SetEveryone(v SummaryMilestoneCountsResponse) {
	o.Everyone = &v
}

// GetMine returns the Mine field value if set, zero value otherwise.
func (o *SummaryMilestoneResponse) GetMine() SummaryMilestoneCountsResponse {
	if o == nil || o.Mine == nil {
		var ret SummaryMilestoneCountsResponse
		return ret
	}
	return *o.Mine
}

// GetMineOk returns a tuple with the Mine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryMilestoneResponse) GetMineOk() (*SummaryMilestoneCountsResponse, bool) {
	if o == nil || o.Mine == nil {
		return nil, false
	}
	return o.Mine, true
}

// HasMine returns a boolean if a field has been set.
func (o *SummaryMilestoneResponse) HasMine() bool {
	if o != nil && o.Mine != nil {
		return true
	}

	return false
}

// SetMine gets a reference to the given SummaryMilestoneCountsResponse and assigns it to the Mine field.
func (o *SummaryMilestoneResponse) SetMine(v SummaryMilestoneCountsResponse) {
	o.Mine = &v
}

func (o SummaryMilestoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Everyone != nil {
		toSerialize["everyone"] = o.Everyone
	}
	if o.Mine != nil {
		toSerialize["mine"] = o.Mine
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryMilestoneResponse struct {
	value *SummaryMilestoneResponse
	isSet bool
}

func (v NullableSummaryMilestoneResponse) Get() *SummaryMilestoneResponse {
	return v.value
}

func (v *NullableSummaryMilestoneResponse) Set(val *SummaryMilestoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryMilestoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryMilestoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryMilestoneResponse(val *SummaryMilestoneResponse) *NullableSummaryMilestoneResponse {
	return &NullableSummaryMilestoneResponse{value: val, isSet: true}
}

func (v NullableSummaryMilestoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryMilestoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


