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

// SummaryColumnResponse ColumnResponse contains counters for columns.
type SummaryColumnResponse struct {
	Count *int32 `json:"count,omitempty"`
	Data *[]SummaryColumnDataResponse `json:"data,omitempty"`
}

// NewSummaryColumnResponse instantiates a new SummaryColumnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryColumnResponse() *SummaryColumnResponse {
	this := SummaryColumnResponse{}
	return &this
}

// NewSummaryColumnResponseWithDefaults instantiates a new SummaryColumnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryColumnResponseWithDefaults() *SummaryColumnResponse {
	this := SummaryColumnResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SummaryColumnResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SummaryColumnResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SummaryColumnResponse) SetCount(v int32) {
	o.Count = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SummaryColumnResponse) GetData() []SummaryColumnDataResponse {
	if o == nil || o.Data == nil {
		var ret []SummaryColumnDataResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnResponse) GetDataOk() (*[]SummaryColumnDataResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SummaryColumnResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SummaryColumnDataResponse and assigns it to the Data field.
func (o *SummaryColumnResponse) SetData(v []SummaryColumnDataResponse) {
	o.Data = &v
}

func (o SummaryColumnResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryColumnResponse struct {
	value *SummaryColumnResponse
	isSet bool
}

func (v NullableSummaryColumnResponse) Get() *SummaryColumnResponse {
	return v.value
}

func (v *NullableSummaryColumnResponse) Set(val *SummaryColumnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryColumnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryColumnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryColumnResponse(val *SummaryColumnResponse) *NullableSummaryColumnResponse {
	return &NullableSummaryColumnResponse{value: val, isSet: true}
}

func (v NullableSummaryColumnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryColumnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


