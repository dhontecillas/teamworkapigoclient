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

// InlineResponse20098TasklistTimeEstimates struct for InlineResponse20098TasklistTimeEstimates
type InlineResponse20098TasklistTimeEstimates struct {
	ActiveHoursEstimated *string `json:"active-hours-estimated,omitempty"`
	ActiveMinsEstimated *string `json:"active-mins-estimated,omitempty"`
	CompletedHoursEstimated *string `json:"completed-hours-estimated,omitempty"`
	CompletedMinsEstimated *string `json:"completed-mins-estimated,omitempty"`
	TotalHoursEstimated *string `json:"total-hours-estimated,omitempty"`
	TotalMinsEstimated *string `json:"total-mins-estimated,omitempty"`
}

// NewInlineResponse20098TasklistTimeEstimates instantiates a new InlineResponse20098TasklistTimeEstimates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098TasklistTimeEstimates() *InlineResponse20098TasklistTimeEstimates {
	this := InlineResponse20098TasklistTimeEstimates{}
	return &this
}

// NewInlineResponse20098TasklistTimeEstimatesWithDefaults instantiates a new InlineResponse20098TasklistTimeEstimates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098TasklistTimeEstimatesWithDefaults() *InlineResponse20098TasklistTimeEstimates {
	this := InlineResponse20098TasklistTimeEstimates{}
	return &this
}

// GetActiveHoursEstimated returns the ActiveHoursEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetActiveHoursEstimated() string {
	if o == nil || o.ActiveHoursEstimated == nil {
		var ret string
		return ret
	}
	return *o.ActiveHoursEstimated
}

// GetActiveHoursEstimatedOk returns a tuple with the ActiveHoursEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetActiveHoursEstimatedOk() (*string, bool) {
	if o == nil || o.ActiveHoursEstimated == nil {
		return nil, false
	}
	return o.ActiveHoursEstimated, true
}

// HasActiveHoursEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasActiveHoursEstimated() bool {
	if o != nil && o.ActiveHoursEstimated != nil {
		return true
	}

	return false
}

// SetActiveHoursEstimated gets a reference to the given string and assigns it to the ActiveHoursEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetActiveHoursEstimated(v string) {
	o.ActiveHoursEstimated = &v
}

// GetActiveMinsEstimated returns the ActiveMinsEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetActiveMinsEstimated() string {
	if o == nil || o.ActiveMinsEstimated == nil {
		var ret string
		return ret
	}
	return *o.ActiveMinsEstimated
}

// GetActiveMinsEstimatedOk returns a tuple with the ActiveMinsEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetActiveMinsEstimatedOk() (*string, bool) {
	if o == nil || o.ActiveMinsEstimated == nil {
		return nil, false
	}
	return o.ActiveMinsEstimated, true
}

// HasActiveMinsEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasActiveMinsEstimated() bool {
	if o != nil && o.ActiveMinsEstimated != nil {
		return true
	}

	return false
}

// SetActiveMinsEstimated gets a reference to the given string and assigns it to the ActiveMinsEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetActiveMinsEstimated(v string) {
	o.ActiveMinsEstimated = &v
}

// GetCompletedHoursEstimated returns the CompletedHoursEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetCompletedHoursEstimated() string {
	if o == nil || o.CompletedHoursEstimated == nil {
		var ret string
		return ret
	}
	return *o.CompletedHoursEstimated
}

// GetCompletedHoursEstimatedOk returns a tuple with the CompletedHoursEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetCompletedHoursEstimatedOk() (*string, bool) {
	if o == nil || o.CompletedHoursEstimated == nil {
		return nil, false
	}
	return o.CompletedHoursEstimated, true
}

// HasCompletedHoursEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasCompletedHoursEstimated() bool {
	if o != nil && o.CompletedHoursEstimated != nil {
		return true
	}

	return false
}

// SetCompletedHoursEstimated gets a reference to the given string and assigns it to the CompletedHoursEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetCompletedHoursEstimated(v string) {
	o.CompletedHoursEstimated = &v
}

// GetCompletedMinsEstimated returns the CompletedMinsEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetCompletedMinsEstimated() string {
	if o == nil || o.CompletedMinsEstimated == nil {
		var ret string
		return ret
	}
	return *o.CompletedMinsEstimated
}

// GetCompletedMinsEstimatedOk returns a tuple with the CompletedMinsEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetCompletedMinsEstimatedOk() (*string, bool) {
	if o == nil || o.CompletedMinsEstimated == nil {
		return nil, false
	}
	return o.CompletedMinsEstimated, true
}

// HasCompletedMinsEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasCompletedMinsEstimated() bool {
	if o != nil && o.CompletedMinsEstimated != nil {
		return true
	}

	return false
}

// SetCompletedMinsEstimated gets a reference to the given string and assigns it to the CompletedMinsEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetCompletedMinsEstimated(v string) {
	o.CompletedMinsEstimated = &v
}

// GetTotalHoursEstimated returns the TotalHoursEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetTotalHoursEstimated() string {
	if o == nil || o.TotalHoursEstimated == nil {
		var ret string
		return ret
	}
	return *o.TotalHoursEstimated
}

// GetTotalHoursEstimatedOk returns a tuple with the TotalHoursEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetTotalHoursEstimatedOk() (*string, bool) {
	if o == nil || o.TotalHoursEstimated == nil {
		return nil, false
	}
	return o.TotalHoursEstimated, true
}

// HasTotalHoursEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasTotalHoursEstimated() bool {
	if o != nil && o.TotalHoursEstimated != nil {
		return true
	}

	return false
}

// SetTotalHoursEstimated gets a reference to the given string and assigns it to the TotalHoursEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetTotalHoursEstimated(v string) {
	o.TotalHoursEstimated = &v
}

// GetTotalMinsEstimated returns the TotalMinsEstimated field value if set, zero value otherwise.
func (o *InlineResponse20098TasklistTimeEstimates) GetTotalMinsEstimated() string {
	if o == nil || o.TotalMinsEstimated == nil {
		var ret string
		return ret
	}
	return *o.TotalMinsEstimated
}

// GetTotalMinsEstimatedOk returns a tuple with the TotalMinsEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098TasklistTimeEstimates) GetTotalMinsEstimatedOk() (*string, bool) {
	if o == nil || o.TotalMinsEstimated == nil {
		return nil, false
	}
	return o.TotalMinsEstimated, true
}

// HasTotalMinsEstimated returns a boolean if a field has been set.
func (o *InlineResponse20098TasklistTimeEstimates) HasTotalMinsEstimated() bool {
	if o != nil && o.TotalMinsEstimated != nil {
		return true
	}

	return false
}

// SetTotalMinsEstimated gets a reference to the given string and assigns it to the TotalMinsEstimated field.
func (o *InlineResponse20098TasklistTimeEstimates) SetTotalMinsEstimated(v string) {
	o.TotalMinsEstimated = &v
}

func (o InlineResponse20098TasklistTimeEstimates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveHoursEstimated != nil {
		toSerialize["active-hours-estimated"] = o.ActiveHoursEstimated
	}
	if o.ActiveMinsEstimated != nil {
		toSerialize["active-mins-estimated"] = o.ActiveMinsEstimated
	}
	if o.CompletedHoursEstimated != nil {
		toSerialize["completed-hours-estimated"] = o.CompletedHoursEstimated
	}
	if o.CompletedMinsEstimated != nil {
		toSerialize["completed-mins-estimated"] = o.CompletedMinsEstimated
	}
	if o.TotalHoursEstimated != nil {
		toSerialize["total-hours-estimated"] = o.TotalHoursEstimated
	}
	if o.TotalMinsEstimated != nil {
		toSerialize["total-mins-estimated"] = o.TotalMinsEstimated
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098TasklistTimeEstimates struct {
	value *InlineResponse20098TasklistTimeEstimates
	isSet bool
}

func (v NullableInlineResponse20098TasklistTimeEstimates) Get() *InlineResponse20098TasklistTimeEstimates {
	return v.value
}

func (v *NullableInlineResponse20098TasklistTimeEstimates) Set(val *InlineResponse20098TasklistTimeEstimates) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098TasklistTimeEstimates) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098TasklistTimeEstimates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098TasklistTimeEstimates(val *InlineResponse20098TasklistTimeEstimates) *NullableInlineResponse20098TasklistTimeEstimates {
	return &NullableInlineResponse20098TasklistTimeEstimates{value: val, isSet: true}
}

func (v NullableInlineResponse20098TasklistTimeEstimates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098TasklistTimeEstimates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


