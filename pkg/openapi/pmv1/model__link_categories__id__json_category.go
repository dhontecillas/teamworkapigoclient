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

// LinkCategoriesIdJsonCategory struct for LinkCategoriesIdJsonCategory
type LinkCategoriesIdJsonCategory struct {
	Name string `json:"name"`
	ParentId *string `json:"parent-id,omitempty"`
}

// NewLinkCategoriesIdJsonCategory instantiates a new LinkCategoriesIdJsonCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkCategoriesIdJsonCategory(name string, ) *LinkCategoriesIdJsonCategory {
	this := LinkCategoriesIdJsonCategory{}
	this.Name = name
	return &this
}

// NewLinkCategoriesIdJsonCategoryWithDefaults instantiates a new LinkCategoriesIdJsonCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkCategoriesIdJsonCategoryWithDefaults() *LinkCategoriesIdJsonCategory {
	this := LinkCategoriesIdJsonCategory{}
	return &this
}

// GetName returns the Name field value
func (o *LinkCategoriesIdJsonCategory) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkCategoriesIdJsonCategory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkCategoriesIdJsonCategory) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *LinkCategoriesIdJsonCategory) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkCategoriesIdJsonCategory) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *LinkCategoriesIdJsonCategory) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *LinkCategoriesIdJsonCategory) SetParentId(v string) {
	o.ParentId = &v
}

func (o LinkCategoriesIdJsonCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ParentId != nil {
		toSerialize["parent-id"] = o.ParentId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkCategoriesIdJsonCategory struct {
	value *LinkCategoriesIdJsonCategory
	isSet bool
}

func (v NullableLinkCategoriesIdJsonCategory) Get() *LinkCategoriesIdJsonCategory {
	return v.value
}

func (v *NullableLinkCategoriesIdJsonCategory) Set(val *LinkCategoriesIdJsonCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkCategoriesIdJsonCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkCategoriesIdJsonCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkCategoriesIdJsonCategory(val *LinkCategoriesIdJsonCategory) *NullableLinkCategoriesIdJsonCategory {
	return &NullableLinkCategoriesIdJsonCategory{value: val, isSet: true}
}

func (v NullableLinkCategoriesIdJsonCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkCategoriesIdJsonCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


