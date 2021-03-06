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

// InlineResponse20089SearchResult struct for InlineResponse20089SearchResult
type InlineResponse20089SearchResult struct {
	MoreAvailable *bool `json:"moreAvailable,omitempty"`
	Tasks *[]InlineResponse20089SearchResultTasks `json:"tasks,omitempty"`
}

// NewInlineResponse20089SearchResult instantiates a new InlineResponse20089SearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089SearchResult() *InlineResponse20089SearchResult {
	this := InlineResponse20089SearchResult{}
	return &this
}

// NewInlineResponse20089SearchResultWithDefaults instantiates a new InlineResponse20089SearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089SearchResultWithDefaults() *InlineResponse20089SearchResult {
	this := InlineResponse20089SearchResult{}
	return &this
}

// GetMoreAvailable returns the MoreAvailable field value if set, zero value otherwise.
func (o *InlineResponse20089SearchResult) GetMoreAvailable() bool {
	if o == nil || o.MoreAvailable == nil {
		var ret bool
		return ret
	}
	return *o.MoreAvailable
}

// GetMoreAvailableOk returns a tuple with the MoreAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089SearchResult) GetMoreAvailableOk() (*bool, bool) {
	if o == nil || o.MoreAvailable == nil {
		return nil, false
	}
	return o.MoreAvailable, true
}

// HasMoreAvailable returns a boolean if a field has been set.
func (o *InlineResponse20089SearchResult) HasMoreAvailable() bool {
	if o != nil && o.MoreAvailable != nil {
		return true
	}

	return false
}

// SetMoreAvailable gets a reference to the given bool and assigns it to the MoreAvailable field.
func (o *InlineResponse20089SearchResult) SetMoreAvailable(v bool) {
	o.MoreAvailable = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InlineResponse20089SearchResult) GetTasks() []InlineResponse20089SearchResultTasks {
	if o == nil || o.Tasks == nil {
		var ret []InlineResponse20089SearchResultTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089SearchResult) GetTasksOk() (*[]InlineResponse20089SearchResultTasks, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InlineResponse20089SearchResult) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []InlineResponse20089SearchResultTasks and assigns it to the Tasks field.
func (o *InlineResponse20089SearchResult) SetTasks(v []InlineResponse20089SearchResultTasks) {
	o.Tasks = &v
}

func (o InlineResponse20089SearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoreAvailable != nil {
		toSerialize["moreAvailable"] = o.MoreAvailable
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089SearchResult struct {
	value *InlineResponse20089SearchResult
	isSet bool
}

func (v NullableInlineResponse20089SearchResult) Get() *InlineResponse20089SearchResult {
	return v.value
}

func (v *NullableInlineResponse20089SearchResult) Set(val *InlineResponse20089SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089SearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089SearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089SearchResult(val *InlineResponse20089SearchResult) *NullableInlineResponse20089SearchResult {
	return &NullableInlineResponse20089SearchResult{value: val, isSet: true}
}

func (v NullableInlineResponse20089SearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089SearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


