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

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	CardId *int32 `json:"cardId,omitempty"`
	ColumnId *int32 `json:"columnId,omitempty"`
	PositionAfterId *int32 `json:"positionAfterId,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *InlineObject2) GetCardId() int32 {
	if o == nil || o.CardId == nil {
		var ret int32
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetCardIdOk() (*int32, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *InlineObject2) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given int32 and assigns it to the CardId field.
func (o *InlineObject2) SetCardId(v int32) {
	o.CardId = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *InlineObject2) GetColumnId() int32 {
	if o == nil || o.ColumnId == nil {
		var ret int32
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetColumnIdOk() (*int32, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *InlineObject2) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given int32 and assigns it to the ColumnId field.
func (o *InlineObject2) SetColumnId(v int32) {
	o.ColumnId = &v
}

// GetPositionAfterId returns the PositionAfterId field value if set, zero value otherwise.
func (o *InlineObject2) GetPositionAfterId() int32 {
	if o == nil || o.PositionAfterId == nil {
		var ret int32
		return ret
	}
	return *o.PositionAfterId
}

// GetPositionAfterIdOk returns a tuple with the PositionAfterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetPositionAfterIdOk() (*int32, bool) {
	if o == nil || o.PositionAfterId == nil {
		return nil, false
	}
	return o.PositionAfterId, true
}

// HasPositionAfterId returns a boolean if a field has been set.
func (o *InlineObject2) HasPositionAfterId() bool {
	if o != nil && o.PositionAfterId != nil {
		return true
	}

	return false
}

// SetPositionAfterId gets a reference to the given int32 and assigns it to the PositionAfterId field.
func (o *InlineObject2) SetPositionAfterId(v int32) {
	o.PositionAfterId = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardId != nil {
		toSerialize["cardId"] = o.CardId
	}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.PositionAfterId != nil {
		toSerialize["positionAfterId"] = o.PositionAfterId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


