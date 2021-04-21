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

// InlineResponse200121ProjectBoardData struct for InlineResponse200121ProjectBoardData
type InlineResponse200121ProjectBoardData struct {
	Board *InlineResponse2002Column `json:"board,omitempty"`
	Card *CalendareventsIdJsonEventType `json:"card,omitempty"`
	Column *InlineResponse2002Column `json:"column,omitempty"`
}

// NewInlineResponse200121ProjectBoardData instantiates a new InlineResponse200121ProjectBoardData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121ProjectBoardData() *InlineResponse200121ProjectBoardData {
	this := InlineResponse200121ProjectBoardData{}
	return &this
}

// NewInlineResponse200121ProjectBoardDataWithDefaults instantiates a new InlineResponse200121ProjectBoardData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121ProjectBoardDataWithDefaults() *InlineResponse200121ProjectBoardData {
	this := InlineResponse200121ProjectBoardData{}
	return &this
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *InlineResponse200121ProjectBoardData) GetBoard() InlineResponse2002Column {
	if o == nil || o.Board == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121ProjectBoardData) GetBoardOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Board == nil {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *InlineResponse200121ProjectBoardData) HasBoard() bool {
	if o != nil && o.Board != nil {
		return true
	}

	return false
}

// SetBoard gets a reference to the given InlineResponse2002Column and assigns it to the Board field.
func (o *InlineResponse200121ProjectBoardData) SetBoard(v InlineResponse2002Column) {
	o.Board = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *InlineResponse200121ProjectBoardData) GetCard() CalendareventsIdJsonEventType {
	if o == nil || o.Card == nil {
		var ret CalendareventsIdJsonEventType
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121ProjectBoardData) GetCardOk() (*CalendareventsIdJsonEventType, bool) {
	if o == nil || o.Card == nil {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *InlineResponse200121ProjectBoardData) HasCard() bool {
	if o != nil && o.Card != nil {
		return true
	}

	return false
}

// SetCard gets a reference to the given CalendareventsIdJsonEventType and assigns it to the Card field.
func (o *InlineResponse200121ProjectBoardData) SetCard(v CalendareventsIdJsonEventType) {
	o.Card = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *InlineResponse200121ProjectBoardData) GetColumn() InlineResponse2002Column {
	if o == nil || o.Column == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121ProjectBoardData) GetColumnOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *InlineResponse200121ProjectBoardData) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given InlineResponse2002Column and assigns it to the Column field.
func (o *InlineResponse200121ProjectBoardData) SetColumn(v InlineResponse2002Column) {
	o.Column = &v
}

func (o InlineResponse200121ProjectBoardData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Board != nil {
		toSerialize["board"] = o.Board
	}
	if o.Card != nil {
		toSerialize["card"] = o.Card
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200121ProjectBoardData struct {
	value *InlineResponse200121ProjectBoardData
	isSet bool
}

func (v NullableInlineResponse200121ProjectBoardData) Get() *InlineResponse200121ProjectBoardData {
	return v.value
}

func (v *NullableInlineResponse200121ProjectBoardData) Set(val *InlineResponse200121ProjectBoardData) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121ProjectBoardData) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121ProjectBoardData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121ProjectBoardData(val *InlineResponse200121ProjectBoardData) *NullableInlineResponse200121ProjectBoardData {
	return &NullableInlineResponse200121ProjectBoardData{value: val, isSet: true}
}

func (v NullableInlineResponse200121ProjectBoardData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121ProjectBoardData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

