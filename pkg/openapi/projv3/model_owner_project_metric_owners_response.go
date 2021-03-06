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

// OwnerProjectMetricOwnersResponse ProjectMetricOwnersResponse contains information about a group of owners.
type OwnerProjectMetricOwnersResponse struct {
	Data *[]OwnerProjectMetricOwner `json:"data,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
}

// NewOwnerProjectMetricOwnersResponse instantiates a new OwnerProjectMetricOwnersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerProjectMetricOwnersResponse() *OwnerProjectMetricOwnersResponse {
	this := OwnerProjectMetricOwnersResponse{}
	return &this
}

// NewOwnerProjectMetricOwnersResponseWithDefaults instantiates a new OwnerProjectMetricOwnersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerProjectMetricOwnersResponseWithDefaults() *OwnerProjectMetricOwnersResponse {
	this := OwnerProjectMetricOwnersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OwnerProjectMetricOwnersResponse) GetData() []OwnerProjectMetricOwner {
	if o == nil || o.Data == nil {
		var ret []OwnerProjectMetricOwner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerProjectMetricOwnersResponse) GetDataOk() (*[]OwnerProjectMetricOwner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OwnerProjectMetricOwnersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OwnerProjectMetricOwner and assigns it to the Data field.
func (o *OwnerProjectMetricOwnersResponse) SetData(v []OwnerProjectMetricOwner) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OwnerProjectMetricOwnersResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerProjectMetricOwnersResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OwnerProjectMetricOwnersResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *OwnerProjectMetricOwnersResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

func (o OwnerProjectMetricOwnersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableOwnerProjectMetricOwnersResponse struct {
	value *OwnerProjectMetricOwnersResponse
	isSet bool
}

func (v NullableOwnerProjectMetricOwnersResponse) Get() *OwnerProjectMetricOwnersResponse {
	return v.value
}

func (v *NullableOwnerProjectMetricOwnersResponse) Set(val *OwnerProjectMetricOwnersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerProjectMetricOwnersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerProjectMetricOwnersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerProjectMetricOwnersResponse(val *OwnerProjectMetricOwnersResponse) *NullableOwnerProjectMetricOwnersResponse {
	return &NullableOwnerProjectMetricOwnersResponse{value: val, isSet: true}
}

func (v NullableOwnerProjectMetricOwnersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerProjectMetricOwnersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


