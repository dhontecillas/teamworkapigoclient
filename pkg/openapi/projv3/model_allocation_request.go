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

// AllocationRequest Request contains information of a allocation to be created or updated.
type AllocationRequest struct {
	Allocation *AllocationAllocation `json:"allocation,omitempty"`
}

// NewAllocationRequest instantiates a new AllocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationRequest() *AllocationRequest {
	this := AllocationRequest{}
	return &this
}

// NewAllocationRequestWithDefaults instantiates a new AllocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationRequestWithDefaults() *AllocationRequest {
	this := AllocationRequest{}
	return &this
}

// GetAllocation returns the Allocation field value if set, zero value otherwise.
func (o *AllocationRequest) GetAllocation() AllocationAllocation {
	if o == nil || o.Allocation == nil {
		var ret AllocationAllocation
		return ret
	}
	return *o.Allocation
}

// GetAllocationOk returns a tuple with the Allocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationRequest) GetAllocationOk() (*AllocationAllocation, bool) {
	if o == nil || o.Allocation == nil {
		return nil, false
	}
	return o.Allocation, true
}

// HasAllocation returns a boolean if a field has been set.
func (o *AllocationRequest) HasAllocation() bool {
	if o != nil && o.Allocation != nil {
		return true
	}

	return false
}

// SetAllocation gets a reference to the given AllocationAllocation and assigns it to the Allocation field.
func (o *AllocationRequest) SetAllocation(v AllocationAllocation) {
	o.Allocation = &v
}

func (o AllocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allocation != nil {
		toSerialize["allocation"] = o.Allocation
	}
	return json.Marshal(toSerialize)
}

type NullableAllocationRequest struct {
	value *AllocationRequest
	isSet bool
}

func (v NullableAllocationRequest) Get() *AllocationRequest {
	return v.value
}

func (v *NullableAllocationRequest) Set(val *AllocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationRequest(val *AllocationRequest) *NullableAllocationRequest {
	return &NullableAllocationRequest{value: val, isSet: true}
}

func (v NullableAllocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


