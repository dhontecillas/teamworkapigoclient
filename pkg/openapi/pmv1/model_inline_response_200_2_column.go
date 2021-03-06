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

// InlineResponse2002Column struct for InlineResponse2002Column
type InlineResponse2002Column struct {
	Color *string `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse2002Column instantiates a new InlineResponse2002Column object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Column() *InlineResponse2002Column {
	this := InlineResponse2002Column{}
	return &this
}

// NewInlineResponse2002ColumnWithDefaults instantiates a new InlineResponse2002Column object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ColumnWithDefaults() *InlineResponse2002Column {
	this := InlineResponse2002Column{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *InlineResponse2002Column) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Column) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *InlineResponse2002Column) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *InlineResponse2002Column) SetColor(v string) {
	o.Color = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2002Column) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Column) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2002Column) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2002Column) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2002Column) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Column) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2002Column) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2002Column) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse2002Column) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2002Column struct {
	value *InlineResponse2002Column
	isSet bool
}

func (v NullableInlineResponse2002Column) Get() *InlineResponse2002Column {
	return v.value
}

func (v *NullableInlineResponse2002Column) Set(val *InlineResponse2002Column) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Column) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Column) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Column(val *InlineResponse2002Column) *NullableInlineResponse2002Column {
	return &NullableInlineResponse2002Column{value: val, isSet: true}
}

func (v NullableInlineResponse2002Column) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Column) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


