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

// InlineResponse20075StatsBilling struct for InlineResponse20075StatsBilling
type InlineResponse20075StatsBilling struct {
	Active *string `json:"active,omitempty"`
	Completed *string `json:"completed,omitempty"`
}

// NewInlineResponse20075StatsBilling instantiates a new InlineResponse20075StatsBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075StatsBilling() *InlineResponse20075StatsBilling {
	this := InlineResponse20075StatsBilling{}
	return &this
}

// NewInlineResponse20075StatsBillingWithDefaults instantiates a new InlineResponse20075StatsBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075StatsBillingWithDefaults() *InlineResponse20075StatsBilling {
	this := InlineResponse20075StatsBilling{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse20075StatsBilling) GetActive() string {
	if o == nil || o.Active == nil {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsBilling) GetActiveOk() (*string, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse20075StatsBilling) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *InlineResponse20075StatsBilling) SetActive(v string) {
	o.Active = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *InlineResponse20075StatsBilling) GetCompleted() string {
	if o == nil || o.Completed == nil {
		var ret string
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075StatsBilling) GetCompletedOk() (*string, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *InlineResponse20075StatsBilling) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given string and assigns it to the Completed field.
func (o *InlineResponse20075StatsBilling) SetCompleted(v string) {
	o.Completed = &v
}

func (o InlineResponse20075StatsBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20075StatsBilling struct {
	value *InlineResponse20075StatsBilling
	isSet bool
}

func (v NullableInlineResponse20075StatsBilling) Get() *InlineResponse20075StatsBilling {
	return v.value
}

func (v *NullableInlineResponse20075StatsBilling) Set(val *InlineResponse20075StatsBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075StatsBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075StatsBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075StatsBilling(val *InlineResponse20075StatsBilling) *NullableInlineResponse20075StatsBilling {
	return &NullableInlineResponse20075StatsBilling{value: val, isSet: true}
}

func (v NullableInlineResponse20075StatsBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075StatsBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


