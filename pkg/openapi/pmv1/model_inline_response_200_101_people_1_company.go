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

// InlineResponse200101People1Company struct for InlineResponse200101People1Company
type InlineResponse200101People1Company struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewInlineResponse200101People1Company instantiates a new InlineResponse200101People1Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200101People1Company() *InlineResponse200101People1Company {
	this := InlineResponse200101People1Company{}
	return &this
}

// NewInlineResponse200101People1CompanyWithDefaults instantiates a new InlineResponse200101People1Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200101People1CompanyWithDefaults() *InlineResponse200101People1Company {
	this := InlineResponse200101People1Company{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200101People1Company) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200101People1Company) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200101People1Company) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InlineResponse200101People1Company) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200101People1Company) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200101People1Company) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200101People1Company) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200101People1Company) SetName(v string) {
	o.Name = &v
}

func (o InlineResponse200101People1Company) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200101People1Company struct {
	value *InlineResponse200101People1Company
	isSet bool
}

func (v NullableInlineResponse200101People1Company) Get() *InlineResponse200101People1Company {
	return v.value
}

func (v *NullableInlineResponse200101People1Company) Set(val *InlineResponse200101People1Company) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200101People1Company) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200101People1Company) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200101People1Company(val *InlineResponse200101People1Company) *NullableInlineResponse200101People1Company {
	return &NullableInlineResponse200101People1Company{value: val, isSet: true}
}

func (v NullableInlineResponse200101People1Company) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200101People1Company) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


