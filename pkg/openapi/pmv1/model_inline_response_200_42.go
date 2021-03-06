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

// InlineResponse20042 struct for InlineResponse20042
type InlineResponse20042 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Boards *[]InlineResponse20042Boards `json:"boards,omitempty"`
}

// NewInlineResponse20042 instantiates a new InlineResponse20042 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20042() *InlineResponse20042 {
	this := InlineResponse20042{}
	return &this
}

// NewInlineResponse20042WithDefaults instantiates a new InlineResponse20042 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20042WithDefaults() *InlineResponse20042 {
	this := InlineResponse20042{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20042) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20042) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20042) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetBoards returns the Boards field value if set, zero value otherwise.
func (o *InlineResponse20042) GetBoards() []InlineResponse20042Boards {
	if o == nil || o.Boards == nil {
		var ret []InlineResponse20042Boards
		return ret
	}
	return *o.Boards
}

// GetBoardsOk returns a tuple with the Boards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetBoardsOk() (*[]InlineResponse20042Boards, bool) {
	if o == nil || o.Boards == nil {
		return nil, false
	}
	return o.Boards, true
}

// HasBoards returns a boolean if a field has been set.
func (o *InlineResponse20042) HasBoards() bool {
	if o != nil && o.Boards != nil {
		return true
	}

	return false
}

// SetBoards gets a reference to the given []InlineResponse20042Boards and assigns it to the Boards field.
func (o *InlineResponse20042) SetBoards(v []InlineResponse20042Boards) {
	o.Boards = &v
}

func (o InlineResponse20042) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Boards != nil {
		toSerialize["boards"] = o.Boards
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20042 struct {
	value *InlineResponse20042
	isSet bool
}

func (v NullableInlineResponse20042) Get() *InlineResponse20042 {
	return v.value
}

func (v *NullableInlineResponse20042) Set(val *InlineResponse20042) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20042) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20042) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20042(val *InlineResponse20042) *NullableInlineResponse20042 {
	return &NullableInlineResponse20042{value: val, isSet: true}
}

func (v NullableInlineResponse20042) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20042) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


