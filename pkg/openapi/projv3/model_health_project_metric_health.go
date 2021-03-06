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

// HealthProjectMetricHealth ProjectMetricHealth stores a specific health counter.
type HealthProjectMetricHealth struct {
	Name *string `json:"name,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewHealthProjectMetricHealth instantiates a new HealthProjectMetricHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthProjectMetricHealth() *HealthProjectMetricHealth {
	this := HealthProjectMetricHealth{}
	return &this
}

// NewHealthProjectMetricHealthWithDefaults instantiates a new HealthProjectMetricHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthProjectMetricHealthWithDefaults() *HealthProjectMetricHealth {
	this := HealthProjectMetricHealth{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HealthProjectMetricHealth) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProjectMetricHealth) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HealthProjectMetricHealth) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HealthProjectMetricHealth) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HealthProjectMetricHealth) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthProjectMetricHealth) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HealthProjectMetricHealth) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *HealthProjectMetricHealth) SetValue(v int32) {
	o.Value = &v
}

func (o HealthProjectMetricHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHealthProjectMetricHealth struct {
	value *HealthProjectMetricHealth
	isSet bool
}

func (v NullableHealthProjectMetricHealth) Get() *HealthProjectMetricHealth {
	return v.value
}

func (v *NullableHealthProjectMetricHealth) Set(val *HealthProjectMetricHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthProjectMetricHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthProjectMetricHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthProjectMetricHealth(val *HealthProjectMetricHealth) *NullableHealthProjectMetricHealth {
	return &NullableHealthProjectMetricHealth{value: val, isSet: true}
}

func (v NullableHealthProjectMetricHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthProjectMetricHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


