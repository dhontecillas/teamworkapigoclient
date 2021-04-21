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

// BudgetRequest Request contains information of a budget to be created or updated.
type BudgetRequest struct {
	Budget *BudgetProjectBudget `json:"budget,omitempty"`
}

// NewBudgetRequest instantiates a new BudgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRequest() *BudgetRequest {
	this := BudgetRequest{}
	return &this
}

// NewBudgetRequestWithDefaults instantiates a new BudgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRequestWithDefaults() *BudgetRequest {
	this := BudgetRequest{}
	return &this
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *BudgetRequest) GetBudget() BudgetProjectBudget {
	if o == nil || o.Budget == nil {
		var ret BudgetProjectBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRequest) GetBudgetOk() (*BudgetProjectBudget, bool) {
	if o == nil || o.Budget == nil {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *BudgetRequest) HasBudget() bool {
	if o != nil && o.Budget != nil {
		return true
	}

	return false
}

// SetBudget gets a reference to the given BudgetProjectBudget and assigns it to the Budget field.
func (o *BudgetRequest) SetBudget(v BudgetProjectBudget) {
	o.Budget = &v
}

func (o BudgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Budget != nil {
		toSerialize["budget"] = o.Budget
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetRequest struct {
	value *BudgetRequest
	isSet bool
}

func (v NullableBudgetRequest) Get() *BudgetRequest {
	return v.value
}

func (v *NullableBudgetRequest) Set(val *BudgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRequest(val *BudgetRequest) *NullableBudgetRequest {
	return &NullableBudgetRequest{value: val, isSet: true}
}

func (v NullableBudgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


