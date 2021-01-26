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

// InlineResponse200120 struct for InlineResponse200120
type InlineResponse200120 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Workload *[]InlineResponse200120Workload `json:"workload,omitempty"`
}

// NewInlineResponse200120 instantiates a new InlineResponse200120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200120() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// NewInlineResponse200120WithDefaults instantiates a new InlineResponse200120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200120WithDefaults() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200120) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200120) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200120) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetWorkload returns the Workload field value if set, zero value otherwise.
func (o *InlineResponse200120) GetWorkload() []InlineResponse200120Workload {
	if o == nil || o.Workload == nil {
		var ret []InlineResponse200120Workload
		return ret
	}
	return *o.Workload
}

// GetWorkloadOk returns a tuple with the Workload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetWorkloadOk() (*[]InlineResponse200120Workload, bool) {
	if o == nil || o.Workload == nil {
		return nil, false
	}
	return o.Workload, true
}

// HasWorkload returns a boolean if a field has been set.
func (o *InlineResponse200120) HasWorkload() bool {
	if o != nil && o.Workload != nil {
		return true
	}

	return false
}

// SetWorkload gets a reference to the given []InlineResponse200120Workload and assigns it to the Workload field.
func (o *InlineResponse200120) SetWorkload(v []InlineResponse200120Workload) {
	o.Workload = &v
}

func (o InlineResponse200120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Workload != nil {
		toSerialize["workload"] = o.Workload
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200120 struct {
	value *InlineResponse200120
	isSet bool
}

func (v NullableInlineResponse200120) Get() *InlineResponse200120 {
	return v.value
}

func (v *NullableInlineResponse200120) Set(val *InlineResponse200120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200120(val *InlineResponse200120) *NullableInlineResponse200120 {
	return &NullableInlineResponse200120{value: val, isSet: true}
}

func (v NullableInlineResponse200120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


