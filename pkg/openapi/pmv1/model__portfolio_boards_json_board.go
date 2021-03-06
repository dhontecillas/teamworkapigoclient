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

// PortfolioBoardsJsonBoard struct for PortfolioBoardsJsonBoard
type PortfolioBoardsJsonBoard struct {
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewPortfolioBoardsJsonBoard instantiates a new PortfolioBoardsJsonBoard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioBoardsJsonBoard() *PortfolioBoardsJsonBoard {
	this := PortfolioBoardsJsonBoard{}
	return &this
}

// NewPortfolioBoardsJsonBoardWithDefaults instantiates a new PortfolioBoardsJsonBoard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioBoardsJsonBoardWithDefaults() *PortfolioBoardsJsonBoard {
	this := PortfolioBoardsJsonBoard{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PortfolioBoardsJsonBoard) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBoardsJsonBoard) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PortfolioBoardsJsonBoard) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PortfolioBoardsJsonBoard) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PortfolioBoardsJsonBoard) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBoardsJsonBoard) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PortfolioBoardsJsonBoard) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PortfolioBoardsJsonBoard) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortfolioBoardsJsonBoard) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioBoardsJsonBoard) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortfolioBoardsJsonBoard) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortfolioBoardsJsonBoard) SetName(v string) {
	o.Name = &v
}

func (o PortfolioBoardsJsonBoard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePortfolioBoardsJsonBoard struct {
	value *PortfolioBoardsJsonBoard
	isSet bool
}

func (v NullablePortfolioBoardsJsonBoard) Get() *PortfolioBoardsJsonBoard {
	return v.value
}

func (v *NullablePortfolioBoardsJsonBoard) Set(val *PortfolioBoardsJsonBoard) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioBoardsJsonBoard) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioBoardsJsonBoard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioBoardsJsonBoard(val *PortfolioBoardsJsonBoard) *NullablePortfolioBoardsJsonBoard {
	return &NullablePortfolioBoardsJsonBoard{value: val, isSet: true}
}

func (v NullablePortfolioBoardsJsonBoard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioBoardsJsonBoard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


