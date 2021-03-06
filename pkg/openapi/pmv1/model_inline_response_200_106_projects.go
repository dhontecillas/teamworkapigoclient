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

// InlineResponse200106Projects struct for InlineResponse200106Projects
type InlineResponse200106Projects struct {
	Company *InlineResponse2002People12345Company `json:"company,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Tasklist *InlineResponse200106Tasklist `json:"tasklist,omitempty"`
}

// NewInlineResponse200106Projects instantiates a new InlineResponse200106Projects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200106Projects() *InlineResponse200106Projects {
	this := InlineResponse200106Projects{}
	return &this
}

// NewInlineResponse200106ProjectsWithDefaults instantiates a new InlineResponse200106Projects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200106ProjectsWithDefaults() *InlineResponse200106Projects {
	this := InlineResponse200106Projects{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse200106Projects) GetCompany() InlineResponse2002People12345Company {
	if o == nil || o.Company == nil {
		var ret InlineResponse2002People12345Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106Projects) GetCompanyOk() (*InlineResponse2002People12345Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse200106Projects) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given InlineResponse2002People12345Company and assigns it to the Company field.
func (o *InlineResponse200106Projects) SetCompany(v InlineResponse2002People12345Company) {
	o.Company = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200106Projects) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106Projects) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200106Projects) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200106Projects) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200106Projects) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106Projects) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200106Projects) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200106Projects) SetName(v string) {
	o.Name = &v
}

// GetTasklist returns the Tasklist field value if set, zero value otherwise.
func (o *InlineResponse200106Projects) GetTasklist() InlineResponse200106Tasklist {
	if o == nil || o.Tasklist == nil {
		var ret InlineResponse200106Tasklist
		return ret
	}
	return *o.Tasklist
}

// GetTasklistOk returns a tuple with the Tasklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106Projects) GetTasklistOk() (*InlineResponse200106Tasklist, bool) {
	if o == nil || o.Tasklist == nil {
		return nil, false
	}
	return o.Tasklist, true
}

// HasTasklist returns a boolean if a field has been set.
func (o *InlineResponse200106Projects) HasTasklist() bool {
	if o != nil && o.Tasklist != nil {
		return true
	}

	return false
}

// SetTasklist gets a reference to the given InlineResponse200106Tasklist and assigns it to the Tasklist field.
func (o *InlineResponse200106Projects) SetTasklist(v InlineResponse200106Tasklist) {
	o.Tasklist = &v
}

func (o InlineResponse200106Projects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tasklist != nil {
		toSerialize["tasklist"] = o.Tasklist
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200106Projects struct {
	value *InlineResponse200106Projects
	isSet bool
}

func (v NullableInlineResponse200106Projects) Get() *InlineResponse200106Projects {
	return v.value
}

func (v *NullableInlineResponse200106Projects) Set(val *InlineResponse200106Projects) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200106Projects) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200106Projects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200106Projects(val *InlineResponse200106Projects) *NullableInlineResponse200106Projects {
	return &NullableInlineResponse200106Projects{value: val, isSet: true}
}

func (v NullableInlineResponse200106Projects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200106Projects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


