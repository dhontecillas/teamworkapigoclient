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

// CustomfieldResponse Response contains information about a specific customfield.
type CustomfieldResponse struct {
	Count *int32 `json:"count,omitempty"`
	Customfield *ViewCustomField `json:"customfield,omitempty"`
}

// NewCustomfieldResponse instantiates a new CustomfieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomfieldResponse() *CustomfieldResponse {
	this := CustomfieldResponse{}
	return &this
}

// NewCustomfieldResponseWithDefaults instantiates a new CustomfieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomfieldResponseWithDefaults() *CustomfieldResponse {
	this := CustomfieldResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CustomfieldResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomfieldResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CustomfieldResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CustomfieldResponse) SetCount(v int32) {
	o.Count = &v
}

// GetCustomfield returns the Customfield field value if set, zero value otherwise.
func (o *CustomfieldResponse) GetCustomfield() ViewCustomField {
	if o == nil || o.Customfield == nil {
		var ret ViewCustomField
		return ret
	}
	return *o.Customfield
}

// GetCustomfieldOk returns a tuple with the Customfield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomfieldResponse) GetCustomfieldOk() (*ViewCustomField, bool) {
	if o == nil || o.Customfield == nil {
		return nil, false
	}
	return o.Customfield, true
}

// HasCustomfield returns a boolean if a field has been set.
func (o *CustomfieldResponse) HasCustomfield() bool {
	if o != nil && o.Customfield != nil {
		return true
	}

	return false
}

// SetCustomfield gets a reference to the given ViewCustomField and assigns it to the Customfield field.
func (o *CustomfieldResponse) SetCustomfield(v ViewCustomField) {
	o.Customfield = &v
}

func (o CustomfieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Customfield != nil {
		toSerialize["customfield"] = o.Customfield
	}
	return json.Marshal(toSerialize)
}

type NullableCustomfieldResponse struct {
	value *CustomfieldResponse
	isSet bool
}

func (v NullableCustomfieldResponse) Get() *CustomfieldResponse {
	return v.value
}

func (v *NullableCustomfieldResponse) Set(val *CustomfieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomfieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomfieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomfieldResponse(val *CustomfieldResponse) *NullableCustomfieldResponse {
	return &NullableCustomfieldResponse{value: val, isSet: true}
}

func (v NullableCustomfieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomfieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


