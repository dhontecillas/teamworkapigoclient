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

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	Category ProjectsIdMessageCategoriesJsonCategory `json:"category"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63(category ProjectsIdMessageCategoriesJsonCategory) *InlineObject63 {
	this := InlineObject63{}
	this.Category = category
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetCategory returns the Category field value
func (o *InlineObject63) GetCategory() ProjectsIdMessageCategoriesJsonCategory {
	if o == nil {
		var ret ProjectsIdMessageCategoriesJsonCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetCategoryOk() (*ProjectsIdMessageCategoriesJsonCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *InlineObject63) SetCategory(v ProjectsIdMessageCategoriesJsonCategory) {
	o.Category = v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

