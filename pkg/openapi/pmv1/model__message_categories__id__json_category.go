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

// MessageCategoriesIdJsonCategory struct for MessageCategoriesIdJsonCategory
type MessageCategoriesIdJsonCategory struct {
	Name string `json:"name"`
	// If you pass in a parent id for a category, you will be creating a sub category. If you don't include this parameter, you will be not be creating a sub category.
	ParentId *string `json:"parent-id,omitempty"`
}

// NewMessageCategoriesIdJsonCategory instantiates a new MessageCategoriesIdJsonCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCategoriesIdJsonCategory(name string) *MessageCategoriesIdJsonCategory {
	this := MessageCategoriesIdJsonCategory{}
	this.Name = name
	return &this
}

// NewMessageCategoriesIdJsonCategoryWithDefaults instantiates a new MessageCategoriesIdJsonCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCategoriesIdJsonCategoryWithDefaults() *MessageCategoriesIdJsonCategory {
	this := MessageCategoriesIdJsonCategory{}
	return &this
}

// GetName returns the Name field value
func (o *MessageCategoriesIdJsonCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageCategoriesIdJsonCategory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageCategoriesIdJsonCategory) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *MessageCategoriesIdJsonCategory) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageCategoriesIdJsonCategory) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *MessageCategoriesIdJsonCategory) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *MessageCategoriesIdJsonCategory) SetParentId(v string) {
	o.ParentId = &v
}

func (o MessageCategoriesIdJsonCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ParentId != nil {
		toSerialize["parent-id"] = o.ParentId
	}
	return json.Marshal(toSerialize)
}

type NullableMessageCategoriesIdJsonCategory struct {
	value *MessageCategoriesIdJsonCategory
	isSet bool
}

func (v NullableMessageCategoriesIdJsonCategory) Get() *MessageCategoriesIdJsonCategory {
	return v.value
}

func (v *NullableMessageCategoriesIdJsonCategory) Set(val *MessageCategoriesIdJsonCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCategoriesIdJsonCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCategoriesIdJsonCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCategoriesIdJsonCategory(val *MessageCategoriesIdJsonCategory) *NullableMessageCategoriesIdJsonCategory {
	return &NullableMessageCategoriesIdJsonCategory{value: val, isSet: true}
}

func (v NullableMessageCategoriesIdJsonCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCategoriesIdJsonCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


