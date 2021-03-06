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

// BudgetProjectBudgetsUsagesResponse ProjectBudgetsUsagesResponse contains information about a group of budgets.
type BudgetProjectBudgetsUsagesResponse struct {
	BudgetUsages *[]BudgetProjectBudgetUsage `json:"budgetUsages,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
}

// NewBudgetProjectBudgetsUsagesResponse instantiates a new BudgetProjectBudgetsUsagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetProjectBudgetsUsagesResponse() *BudgetProjectBudgetsUsagesResponse {
	this := BudgetProjectBudgetsUsagesResponse{}
	return &this
}

// NewBudgetProjectBudgetsUsagesResponseWithDefaults instantiates a new BudgetProjectBudgetsUsagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetProjectBudgetsUsagesResponseWithDefaults() *BudgetProjectBudgetsUsagesResponse {
	this := BudgetProjectBudgetsUsagesResponse{}
	return &this
}

// GetBudgetUsages returns the BudgetUsages field value if set, zero value otherwise.
func (o *BudgetProjectBudgetsUsagesResponse) GetBudgetUsages() []BudgetProjectBudgetUsage {
	if o == nil || o.BudgetUsages == nil {
		var ret []BudgetProjectBudgetUsage
		return ret
	}
	return *o.BudgetUsages
}

// GetBudgetUsagesOk returns a tuple with the BudgetUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetsUsagesResponse) GetBudgetUsagesOk() (*[]BudgetProjectBudgetUsage, bool) {
	if o == nil || o.BudgetUsages == nil {
		return nil, false
	}
	return o.BudgetUsages, true
}

// HasBudgetUsages returns a boolean if a field has been set.
func (o *BudgetProjectBudgetsUsagesResponse) HasBudgetUsages() bool {
	if o != nil && o.BudgetUsages != nil {
		return true
	}

	return false
}

// SetBudgetUsages gets a reference to the given []BudgetProjectBudgetUsage and assigns it to the BudgetUsages field.
func (o *BudgetProjectBudgetsUsagesResponse) SetBudgetUsages(v []BudgetProjectBudgetUsage) {
	o.BudgetUsages = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BudgetProjectBudgetsUsagesResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetProjectBudgetsUsagesResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BudgetProjectBudgetsUsagesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *BudgetProjectBudgetsUsagesResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

func (o BudgetProjectBudgetsUsagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BudgetUsages != nil {
		toSerialize["budgetUsages"] = o.BudgetUsages
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetProjectBudgetsUsagesResponse struct {
	value *BudgetProjectBudgetsUsagesResponse
	isSet bool
}

func (v NullableBudgetProjectBudgetsUsagesResponse) Get() *BudgetProjectBudgetsUsagesResponse {
	return v.value
}

func (v *NullableBudgetProjectBudgetsUsagesResponse) Set(val *BudgetProjectBudgetsUsagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetProjectBudgetsUsagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetProjectBudgetsUsagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetProjectBudgetsUsagesResponse(val *BudgetProjectBudgetsUsagesResponse) *NullableBudgetProjectBudgetsUsagesResponse {
	return &NullableBudgetProjectBudgetsUsagesResponse{value: val, isSet: true}
}

func (v NullableBudgetProjectBudgetsUsagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetProjectBudgetsUsagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


