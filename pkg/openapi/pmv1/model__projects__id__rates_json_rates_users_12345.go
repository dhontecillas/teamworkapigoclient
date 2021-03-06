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

// ProjectsIdRatesJsonRatesUsers12345 struct for ProjectsIdRatesJsonRatesUsers12345
type ProjectsIdRatesJsonRatesUsers12345 struct {
	Rate *int32 `json:"rate,omitempty"`
}

// NewProjectsIdRatesJsonRatesUsers12345 instantiates a new ProjectsIdRatesJsonRatesUsers12345 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdRatesJsonRatesUsers12345() *ProjectsIdRatesJsonRatesUsers12345 {
	this := ProjectsIdRatesJsonRatesUsers12345{}
	return &this
}

// NewProjectsIdRatesJsonRatesUsers12345WithDefaults instantiates a new ProjectsIdRatesJsonRatesUsers12345 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdRatesJsonRatesUsers12345WithDefaults() *ProjectsIdRatesJsonRatesUsers12345 {
	this := ProjectsIdRatesJsonRatesUsers12345{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ProjectsIdRatesJsonRatesUsers12345) GetRate() int32 {
	if o == nil || o.Rate == nil {
		var ret int32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRatesJsonRatesUsers12345) GetRateOk() (*int32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ProjectsIdRatesJsonRatesUsers12345) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given int32 and assigns it to the Rate field.
func (o *ProjectsIdRatesJsonRatesUsers12345) SetRate(v int32) {
	o.Rate = &v
}

func (o ProjectsIdRatesJsonRatesUsers12345) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdRatesJsonRatesUsers12345 struct {
	value *ProjectsIdRatesJsonRatesUsers12345
	isSet bool
}

func (v NullableProjectsIdRatesJsonRatesUsers12345) Get() *ProjectsIdRatesJsonRatesUsers12345 {
	return v.value
}

func (v *NullableProjectsIdRatesJsonRatesUsers12345) Set(val *ProjectsIdRatesJsonRatesUsers12345) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdRatesJsonRatesUsers12345) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdRatesJsonRatesUsers12345) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdRatesJsonRatesUsers12345(val *ProjectsIdRatesJsonRatesUsers12345) *NullableProjectsIdRatesJsonRatesUsers12345 {
	return &NullableProjectsIdRatesJsonRatesUsers12345{value: val, isSet: true}
}

func (v NullableProjectsIdRatesJsonRatesUsers12345) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdRatesJsonRatesUsers12345) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


