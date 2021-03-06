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

// SummaryColumnEstimatedResponse ColumnEstimatedResponse contains estimated counters about the columns's state.
type SummaryColumnEstimatedResponse struct {
	Active *int32 `json:"active,omitempty"`
	Archived *int32 `json:"archived,omitempty"`
	Completed *int32 `json:"completed,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewSummaryColumnEstimatedResponse instantiates a new SummaryColumnEstimatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryColumnEstimatedResponse() *SummaryColumnEstimatedResponse {
	this := SummaryColumnEstimatedResponse{}
	return &this
}

// NewSummaryColumnEstimatedResponseWithDefaults instantiates a new SummaryColumnEstimatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryColumnEstimatedResponseWithDefaults() *SummaryColumnEstimatedResponse {
	this := SummaryColumnEstimatedResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SummaryColumnEstimatedResponse) GetActive() int32 {
	if o == nil || o.Active == nil {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnEstimatedResponse) GetActiveOk() (*int32, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SummaryColumnEstimatedResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *SummaryColumnEstimatedResponse) SetActive(v int32) {
	o.Active = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SummaryColumnEstimatedResponse) GetArchived() int32 {
	if o == nil || o.Archived == nil {
		var ret int32
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnEstimatedResponse) GetArchivedOk() (*int32, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SummaryColumnEstimatedResponse) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given int32 and assigns it to the Archived field.
func (o *SummaryColumnEstimatedResponse) SetArchived(v int32) {
	o.Archived = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SummaryColumnEstimatedResponse) GetCompleted() int32 {
	if o == nil || o.Completed == nil {
		var ret int32
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnEstimatedResponse) GetCompletedOk() (*int32, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SummaryColumnEstimatedResponse) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int32 and assigns it to the Completed field.
func (o *SummaryColumnEstimatedResponse) SetCompleted(v int32) {
	o.Completed = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SummaryColumnEstimatedResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryColumnEstimatedResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SummaryColumnEstimatedResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SummaryColumnEstimatedResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o SummaryColumnEstimatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryColumnEstimatedResponse struct {
	value *SummaryColumnEstimatedResponse
	isSet bool
}

func (v NullableSummaryColumnEstimatedResponse) Get() *SummaryColumnEstimatedResponse {
	return v.value
}

func (v *NullableSummaryColumnEstimatedResponse) Set(val *SummaryColumnEstimatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryColumnEstimatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryColumnEstimatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryColumnEstimatedResponse(val *SummaryColumnEstimatedResponse) *NullableSummaryColumnEstimatedResponse {
	return &NullableSummaryColumnEstimatedResponse{value: val, isSet: true}
}

func (v NullableSummaryColumnEstimatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryColumnEstimatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


