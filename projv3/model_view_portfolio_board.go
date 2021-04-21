/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"encoding/json"
)

// ViewPortfolioBoard PortfolioBoard contains all the information returned from a portfolio board.
type ViewPortfolioBoard struct {
	Color *string `json:"color,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewViewPortfolioBoard instantiates a new ViewPortfolioBoard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewPortfolioBoard() *ViewPortfolioBoard {
	this := ViewPortfolioBoard{}
	return &this
}

// NewViewPortfolioBoardWithDefaults instantiates a new ViewPortfolioBoard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewPortfolioBoardWithDefaults() *ViewPortfolioBoard {
	this := ViewPortfolioBoard{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ViewPortfolioBoard) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioBoard) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ViewPortfolioBoard) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ViewPortfolioBoard) SetColor(v string) {
	o.Color = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewPortfolioBoard) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioBoard) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewPortfolioBoard) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewPortfolioBoard) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViewPortfolioBoard) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioBoard) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViewPortfolioBoard) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViewPortfolioBoard) SetName(v string) {
	o.Name = &v
}

func (o ViewPortfolioBoard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableViewPortfolioBoard struct {
	value *ViewPortfolioBoard
	isSet bool
}

func (v NullableViewPortfolioBoard) Get() *ViewPortfolioBoard {
	return v.value
}

func (v *NullableViewPortfolioBoard) Set(val *ViewPortfolioBoard) {
	v.value = val
	v.isSet = true
}

func (v NullableViewPortfolioBoard) IsSet() bool {
	return v.isSet
}

func (v *NullableViewPortfolioBoard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewPortfolioBoard(val *ViewPortfolioBoard) *NullableViewPortfolioBoard {
	return &NullableViewPortfolioBoard{value: val, isSet: true}
}

func (v NullableViewPortfolioBoard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewPortfolioBoard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

