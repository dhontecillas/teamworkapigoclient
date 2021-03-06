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

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Cards *[]InlineResponse2002Cards `json:"cards,omitempty"`
	Column *InlineResponse2002Column `json:"column,omitempty"`
	People *InlineResponse2002People `json:"people,omitempty"`
}

// NewInlineResponse2002 instantiates a new InlineResponse2002 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002WithDefaults() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse2002) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse2002) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse2002) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *InlineResponse2002) GetCards() []InlineResponse2002Cards {
	if o == nil || o.Cards == nil {
		var ret []InlineResponse2002Cards
		return ret
	}
	return *o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetCardsOk() (*[]InlineResponse2002Cards, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *InlineResponse2002) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given []InlineResponse2002Cards and assigns it to the Cards field.
func (o *InlineResponse2002) SetCards(v []InlineResponse2002Cards) {
	o.Cards = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *InlineResponse2002) GetColumn() InlineResponse2002Column {
	if o == nil || o.Column == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetColumnOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *InlineResponse2002) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given InlineResponse2002Column and assigns it to the Column field.
func (o *InlineResponse2002) SetColumn(v InlineResponse2002Column) {
	o.Column = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *InlineResponse2002) GetPeople() InlineResponse2002People {
	if o == nil || o.People == nil {
		var ret InlineResponse2002People
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetPeopleOk() (*InlineResponse2002People, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *InlineResponse2002) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given InlineResponse2002People and assigns it to the People field.
func (o *InlineResponse2002) SetPeople(v InlineResponse2002People) {
	o.People = &v
}

func (o InlineResponse2002) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002 struct {
	value *InlineResponse2002
	isSet bool
}

func (v NullableInlineResponse2002) Get() *InlineResponse2002 {
	return v.value
}

func (v *NullableInlineResponse2002) Set(val *InlineResponse2002) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002(val *InlineResponse2002) *NullableInlineResponse2002 {
	return &NullableInlineResponse2002{value: val, isSet: true}
}

func (v NullableInlineResponse2002) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


