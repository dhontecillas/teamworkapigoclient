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

// InlineResponse20063EmailaddressAddresses struct for InlineResponse20063EmailaddressAddresses
type InlineResponse20063EmailaddressAddresses struct {
	Links *string `json:"links,omitempty"`
	Notebooks *string `json:"notebooks,omitempty"`
	Tasks *string `json:"tasks,omitempty"`
}

// NewInlineResponse20063EmailaddressAddresses instantiates a new InlineResponse20063EmailaddressAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063EmailaddressAddresses() *InlineResponse20063EmailaddressAddresses {
	this := InlineResponse20063EmailaddressAddresses{}
	return &this
}

// NewInlineResponse20063EmailaddressAddressesWithDefaults instantiates a new InlineResponse20063EmailaddressAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063EmailaddressAddressesWithDefaults() *InlineResponse20063EmailaddressAddresses {
	this := InlineResponse20063EmailaddressAddresses{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InlineResponse20063EmailaddressAddresses) GetLinks() string {
	if o == nil || o.Links == nil {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063EmailaddressAddresses) GetLinksOk() (*string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InlineResponse20063EmailaddressAddresses) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *InlineResponse20063EmailaddressAddresses) SetLinks(v string) {
	o.Links = &v
}

// GetNotebooks returns the Notebooks field value if set, zero value otherwise.
func (o *InlineResponse20063EmailaddressAddresses) GetNotebooks() string {
	if o == nil || o.Notebooks == nil {
		var ret string
		return ret
	}
	return *o.Notebooks
}

// GetNotebooksOk returns a tuple with the Notebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063EmailaddressAddresses) GetNotebooksOk() (*string, bool) {
	if o == nil || o.Notebooks == nil {
		return nil, false
	}
	return o.Notebooks, true
}

// HasNotebooks returns a boolean if a field has been set.
func (o *InlineResponse20063EmailaddressAddresses) HasNotebooks() bool {
	if o != nil && o.Notebooks != nil {
		return true
	}

	return false
}

// SetNotebooks gets a reference to the given string and assigns it to the Notebooks field.
func (o *InlineResponse20063EmailaddressAddresses) SetNotebooks(v string) {
	o.Notebooks = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InlineResponse20063EmailaddressAddresses) GetTasks() string {
	if o == nil || o.Tasks == nil {
		var ret string
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063EmailaddressAddresses) GetTasksOk() (*string, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InlineResponse20063EmailaddressAddresses) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given string and assigns it to the Tasks field.
func (o *InlineResponse20063EmailaddressAddresses) SetTasks(v string) {
	o.Tasks = &v
}

func (o InlineResponse20063EmailaddressAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Notebooks != nil {
		toSerialize["notebooks"] = o.Notebooks
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20063EmailaddressAddresses struct {
	value *InlineResponse20063EmailaddressAddresses
	isSet bool
}

func (v NullableInlineResponse20063EmailaddressAddresses) Get() *InlineResponse20063EmailaddressAddresses {
	return v.value
}

func (v *NullableInlineResponse20063EmailaddressAddresses) Set(val *InlineResponse20063EmailaddressAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063EmailaddressAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063EmailaddressAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063EmailaddressAddresses(val *InlineResponse20063EmailaddressAddresses) *NullableInlineResponse20063EmailaddressAddresses {
	return &NullableInlineResponse20063EmailaddressAddresses{value: val, isSet: true}
}

func (v NullableInlineResponse20063EmailaddressAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063EmailaddressAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


