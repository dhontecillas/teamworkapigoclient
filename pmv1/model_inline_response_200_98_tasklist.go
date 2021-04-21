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

// InlineResponse20098Tasklist struct for InlineResponse20098Tasklist
type InlineResponse20098Tasklist struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TimeEstimates *InlineResponse20098TasklistTimeEstimates `json:"time-estimates,omitempty"`
	TimeTotals *InlineResponse20098TasklistTimeTotals `json:"time-totals,omitempty"`
}

// NewInlineResponse20098Tasklist instantiates a new InlineResponse20098Tasklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098Tasklist() *InlineResponse20098Tasklist {
	this := InlineResponse20098Tasklist{}
	return &this
}

// NewInlineResponse20098TasklistWithDefaults instantiates a new InlineResponse20098Tasklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098TasklistWithDefaults() *InlineResponse20098Tasklist {
	this := InlineResponse20098Tasklist{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20098Tasklist) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Tasklist) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20098Tasklist) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20098Tasklist) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20098Tasklist) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Tasklist) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20098Tasklist) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20098Tasklist) SetName(v string) {
	o.Name = &v
}

// GetTimeEstimates returns the TimeEstimates field value if set, zero value otherwise.
func (o *InlineResponse20098Tasklist) GetTimeEstimates() InlineResponse20098TasklistTimeEstimates {
	if o == nil || o.TimeEstimates == nil {
		var ret InlineResponse20098TasklistTimeEstimates
		return ret
	}
	return *o.TimeEstimates
}

// GetTimeEstimatesOk returns a tuple with the TimeEstimates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Tasklist) GetTimeEstimatesOk() (*InlineResponse20098TasklistTimeEstimates, bool) {
	if o == nil || o.TimeEstimates == nil {
		return nil, false
	}
	return o.TimeEstimates, true
}

// HasTimeEstimates returns a boolean if a field has been set.
func (o *InlineResponse20098Tasklist) HasTimeEstimates() bool {
	if o != nil && o.TimeEstimates != nil {
		return true
	}

	return false
}

// SetTimeEstimates gets a reference to the given InlineResponse20098TasklistTimeEstimates and assigns it to the TimeEstimates field.
func (o *InlineResponse20098Tasklist) SetTimeEstimates(v InlineResponse20098TasklistTimeEstimates) {
	o.TimeEstimates = &v
}

// GetTimeTotals returns the TimeTotals field value if set, zero value otherwise.
func (o *InlineResponse20098Tasklist) GetTimeTotals() InlineResponse20098TasklistTimeTotals {
	if o == nil || o.TimeTotals == nil {
		var ret InlineResponse20098TasklistTimeTotals
		return ret
	}
	return *o.TimeTotals
}

// GetTimeTotalsOk returns a tuple with the TimeTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Tasklist) GetTimeTotalsOk() (*InlineResponse20098TasklistTimeTotals, bool) {
	if o == nil || o.TimeTotals == nil {
		return nil, false
	}
	return o.TimeTotals, true
}

// HasTimeTotals returns a boolean if a field has been set.
func (o *InlineResponse20098Tasklist) HasTimeTotals() bool {
	if o != nil && o.TimeTotals != nil {
		return true
	}

	return false
}

// SetTimeTotals gets a reference to the given InlineResponse20098TasklistTimeTotals and assigns it to the TimeTotals field.
func (o *InlineResponse20098Tasklist) SetTimeTotals(v InlineResponse20098TasklistTimeTotals) {
	o.TimeTotals = &v
}

func (o InlineResponse20098Tasklist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeEstimates != nil {
		toSerialize["time-estimates"] = o.TimeEstimates
	}
	if o.TimeTotals != nil {
		toSerialize["time-totals"] = o.TimeTotals
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098Tasklist struct {
	value *InlineResponse20098Tasklist
	isSet bool
}

func (v NullableInlineResponse20098Tasklist) Get() *InlineResponse20098Tasklist {
	return v.value
}

func (v *NullableInlineResponse20098Tasklist) Set(val *InlineResponse20098Tasklist) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098Tasklist) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098Tasklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098Tasklist(val *InlineResponse20098Tasklist) *NullableInlineResponse20098Tasklist {
	return &NullableInlineResponse20098Tasklist{value: val, isSet: true}
}

func (v NullableInlineResponse20098Tasklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098Tasklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


