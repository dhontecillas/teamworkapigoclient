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

// InlineObject54 struct for InlineObject54
type InlineObject54 struct {
	Get *string `json:"get,omitempty"`
	ReactionType *string `json:"reactionType,omitempty"`
}

// NewInlineObject54 instantiates a new InlineObject54 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject54() *InlineObject54 {
	this := InlineObject54{}
	return &this
}

// NewInlineObject54WithDefaults instantiates a new InlineObject54 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject54WithDefaults() *InlineObject54 {
	this := InlineObject54{}
	return &this
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *InlineObject54) GetGet() string {
	if o == nil || o.Get == nil {
		var ret string
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetGetOk() (*string, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *InlineObject54) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given string and assigns it to the Get field.
func (o *InlineObject54) SetGet(v string) {
	o.Get = &v
}

// GetReactionType returns the ReactionType field value if set, zero value otherwise.
func (o *InlineObject54) GetReactionType() string {
	if o == nil || o.ReactionType == nil {
		var ret string
		return ret
	}
	return *o.ReactionType
}

// GetReactionTypeOk returns a tuple with the ReactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetReactionTypeOk() (*string, bool) {
	if o == nil || o.ReactionType == nil {
		return nil, false
	}
	return o.ReactionType, true
}

// HasReactionType returns a boolean if a field has been set.
func (o *InlineObject54) HasReactionType() bool {
	if o != nil && o.ReactionType != nil {
		return true
	}

	return false
}

// SetReactionType gets a reference to the given string and assigns it to the ReactionType field.
func (o *InlineObject54) SetReactionType(v string) {
	o.ReactionType = &v
}

func (o InlineObject54) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.ReactionType != nil {
		toSerialize["reactionType"] = o.ReactionType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject54 struct {
	value *InlineObject54
	isSet bool
}

func (v NullableInlineObject54) Get() *InlineObject54 {
	return v.value
}

func (v *NullableInlineObject54) Set(val *InlineObject54) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject54) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject54) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject54(val *InlineObject54) *NullableInlineObject54 {
	return &NullableInlineObject54{value: val, isSet: true}
}

func (v NullableInlineObject54) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject54) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


