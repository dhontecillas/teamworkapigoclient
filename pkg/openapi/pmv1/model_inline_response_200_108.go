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

// InlineResponse200108 struct for InlineResponse200108
type InlineResponse200108 struct {
	STATUS *string `json:"STATUS,omitempty"`
	TodoItems *[]InlineResponse200108TodoItems `json:"todo-items,omitempty"`
}

// NewInlineResponse200108 instantiates a new InlineResponse200108 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200108() *InlineResponse200108 {
	this := InlineResponse200108{}
	return &this
}

// NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200108WithDefaults() *InlineResponse200108 {
	this := InlineResponse200108{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse200108) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse200108) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse200108) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetTodoItems returns the TodoItems field value if set, zero value otherwise.
func (o *InlineResponse200108) GetTodoItems() []InlineResponse200108TodoItems {
	if o == nil || o.TodoItems == nil {
		var ret []InlineResponse200108TodoItems
		return ret
	}
	return *o.TodoItems
}

// GetTodoItemsOk returns a tuple with the TodoItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108) GetTodoItemsOk() (*[]InlineResponse200108TodoItems, bool) {
	if o == nil || o.TodoItems == nil {
		return nil, false
	}
	return o.TodoItems, true
}

// HasTodoItems returns a boolean if a field has been set.
func (o *InlineResponse200108) HasTodoItems() bool {
	if o != nil && o.TodoItems != nil {
		return true
	}

	return false
}

// SetTodoItems gets a reference to the given []InlineResponse200108TodoItems and assigns it to the TodoItems field.
func (o *InlineResponse200108) SetTodoItems(v []InlineResponse200108TodoItems) {
	o.TodoItems = &v
}

func (o InlineResponse200108) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.TodoItems != nil {
		toSerialize["todo-items"] = o.TodoItems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200108 struct {
	value *InlineResponse200108
	isSet bool
}

func (v NullableInlineResponse200108) Get() *InlineResponse200108 {
	return v.value
}

func (v *NullableInlineResponse200108) Set(val *InlineResponse200108) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200108) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200108) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200108(val *InlineResponse200108) *NullableInlineResponse200108 {
	return &NullableInlineResponse200108{value: val, isSet: true}
}

func (v NullableInlineResponse200108) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200108) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


