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

// InlineResponse20074Roles struct for InlineResponse20074Roles
type InlineResponse20074Roles struct {
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Users *[]InlineResponse20074Users `json:"users,omitempty"`
}

// NewInlineResponse20074Roles instantiates a new InlineResponse20074Roles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074Roles() *InlineResponse20074Roles {
	this := InlineResponse20074Roles{}
	return &this
}

// NewInlineResponse20074RolesWithDefaults instantiates a new InlineResponse20074Roles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074RolesWithDefaults() *InlineResponse20074Roles {
	this := InlineResponse20074Roles{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20074Roles) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074Roles) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20074Roles) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20074Roles) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20074Roles) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074Roles) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20074Roles) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20074Roles) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20074Roles) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074Roles) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20074Roles) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20074Roles) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineResponse20074Roles) GetUsers() []InlineResponse20074Users {
	if o == nil || o.Users == nil {
		var ret []InlineResponse20074Users
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074Roles) GetUsersOk() (*[]InlineResponse20074Users, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineResponse20074Roles) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []InlineResponse20074Users and assigns it to the Users field.
func (o *InlineResponse20074Roles) SetUsers(v []InlineResponse20074Users) {
	o.Users = &v
}

func (o InlineResponse20074Roles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074Roles struct {
	value *InlineResponse20074Roles
	isSet bool
}

func (v NullableInlineResponse20074Roles) Get() *InlineResponse20074Roles {
	return v.value
}

func (v *NullableInlineResponse20074Roles) Set(val *InlineResponse20074Roles) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074Roles) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074Roles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074Roles(val *InlineResponse20074Roles) *NullableInlineResponse20074Roles {
	return &NullableInlineResponse20074Roles{value: val, isSet: true}
}

func (v NullableInlineResponse20074Roles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074Roles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


