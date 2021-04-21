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

// InlineResponse20061 struct for InlineResponse20061
type InlineResponse20061 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Columns *[]InlineResponse20061Columns `json:"columns,omitempty"`
	People *InlineResponse20061People `json:"people,omitempty"`
}

// NewInlineResponse20061 instantiates a new InlineResponse20061 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20061() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// NewInlineResponse20061WithDefaults instantiates a new InlineResponse20061 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20061WithDefaults() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20061) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20061) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20061) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *InlineResponse20061) GetColumns() []InlineResponse20061Columns {
	if o == nil || o.Columns == nil {
		var ret []InlineResponse20061Columns
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetColumnsOk() (*[]InlineResponse20061Columns, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *InlineResponse20061) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []InlineResponse20061Columns and assigns it to the Columns field.
func (o *InlineResponse20061) SetColumns(v []InlineResponse20061Columns) {
	o.Columns = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *InlineResponse20061) GetPeople() InlineResponse20061People {
	if o == nil || o.People == nil {
		var ret InlineResponse20061People
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetPeopleOk() (*InlineResponse20061People, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *InlineResponse20061) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given InlineResponse20061People and assigns it to the People field.
func (o *InlineResponse20061) SetPeople(v InlineResponse20061People) {
	o.People = &v
}

func (o InlineResponse20061) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20061 struct {
	value *InlineResponse20061
	isSet bool
}

func (v NullableInlineResponse20061) Get() *InlineResponse20061 {
	return v.value
}

func (v *NullableInlineResponse20061) Set(val *InlineResponse20061) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20061) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20061) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20061(val *InlineResponse20061) *NullableInlineResponse20061 {
	return &NullableInlineResponse20061{value: val, isSet: true}
}

func (v NullableInlineResponse20061) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20061) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

