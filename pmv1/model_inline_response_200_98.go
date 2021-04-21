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

// InlineResponse20098 struct for InlineResponse20098
type InlineResponse20098 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Projects *[]InlineResponse20098Projects `json:"projects,omitempty"`
}

// NewInlineResponse20098 instantiates a new InlineResponse20098 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// NewInlineResponse20098WithDefaults instantiates a new InlineResponse20098 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098WithDefaults() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineResponse20098) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineResponse20098) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineResponse20098) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *InlineResponse20098) GetProjects() []InlineResponse20098Projects {
	if o == nil || o.Projects == nil {
		var ret []InlineResponse20098Projects
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetProjectsOk() (*[]InlineResponse20098Projects, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *InlineResponse20098) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []InlineResponse20098Projects and assigns it to the Projects field.
func (o *InlineResponse20098) SetProjects(v []InlineResponse20098Projects) {
	o.Projects = &v
}

func (o InlineResponse20098) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098 struct {
	value *InlineResponse20098
	isSet bool
}

func (v NullableInlineResponse20098) Get() *InlineResponse20098 {
	return v.value
}

func (v *NullableInlineResponse20098) Set(val *InlineResponse20098) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098(val *InlineResponse20098) *NullableInlineResponse20098 {
	return &NullableInlineResponse20098{value: val, isSet: true}
}

func (v NullableInlineResponse20098) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

