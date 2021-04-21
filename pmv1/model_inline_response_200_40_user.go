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

// InlineResponse20040User struct for InlineResponse20040User
type InlineResponse20040User struct {
	Billable *[]map[string]interface{} `json:"billable,omitempty"`
	Endepoch *string `json:"endepoch,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Id *string `json:"id,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Nonbillable *[]map[string]interface{} `json:"nonbillable,omitempty"`
	Startepoch *string `json:"startepoch,omitempty"`
}

// NewInlineResponse20040User instantiates a new InlineResponse20040User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20040User() *InlineResponse20040User {
	this := InlineResponse20040User{}
	return &this
}

// NewInlineResponse20040UserWithDefaults instantiates a new InlineResponse20040User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20040UserWithDefaults() *InlineResponse20040User {
	this := InlineResponse20040User{}
	return &this
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetBillable() []map[string]interface{} {
	if o == nil || o.Billable == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetBillableOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Billable == nil {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasBillable() bool {
	if o != nil && o.Billable != nil {
		return true
	}

	return false
}

// SetBillable gets a reference to the given []map[string]interface{} and assigns it to the Billable field.
func (o *InlineResponse20040User) SetBillable(v []map[string]interface{}) {
	o.Billable = &v
}

// GetEndepoch returns the Endepoch field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetEndepoch() string {
	if o == nil || o.Endepoch == nil {
		var ret string
		return ret
	}
	return *o.Endepoch
}

// GetEndepochOk returns a tuple with the Endepoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetEndepochOk() (*string, bool) {
	if o == nil || o.Endepoch == nil {
		return nil, false
	}
	return o.Endepoch, true
}

// HasEndepoch returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasEndepoch() bool {
	if o != nil && o.Endepoch != nil {
		return true
	}

	return false
}

// SetEndepoch gets a reference to the given string and assigns it to the Endepoch field.
func (o *InlineResponse20040User) SetEndepoch(v string) {
	o.Endepoch = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *InlineResponse20040User) SetFirstname(v string) {
	o.Firstname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20040User) SetId(v string) {
	o.Id = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *InlineResponse20040User) SetLastname(v string) {
	o.Lastname = &v
}

// GetNonbillable returns the Nonbillable field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetNonbillable() []map[string]interface{} {
	if o == nil || o.Nonbillable == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Nonbillable
}

// GetNonbillableOk returns a tuple with the Nonbillable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetNonbillableOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Nonbillable == nil {
		return nil, false
	}
	return o.Nonbillable, true
}

// HasNonbillable returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasNonbillable() bool {
	if o != nil && o.Nonbillable != nil {
		return true
	}

	return false
}

// SetNonbillable gets a reference to the given []map[string]interface{} and assigns it to the Nonbillable field.
func (o *InlineResponse20040User) SetNonbillable(v []map[string]interface{}) {
	o.Nonbillable = &v
}

// GetStartepoch returns the Startepoch field value if set, zero value otherwise.
func (o *InlineResponse20040User) GetStartepoch() string {
	if o == nil || o.Startepoch == nil {
		var ret string
		return ret
	}
	return *o.Startepoch
}

// GetStartepochOk returns a tuple with the Startepoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040User) GetStartepochOk() (*string, bool) {
	if o == nil || o.Startepoch == nil {
		return nil, false
	}
	return o.Startepoch, true
}

// HasStartepoch returns a boolean if a field has been set.
func (o *InlineResponse20040User) HasStartepoch() bool {
	if o != nil && o.Startepoch != nil {
		return true
	}

	return false
}

// SetStartepoch gets a reference to the given string and assigns it to the Startepoch field.
func (o *InlineResponse20040User) SetStartepoch(v string) {
	o.Startepoch = &v
}

func (o InlineResponse20040User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Billable != nil {
		toSerialize["billable"] = o.Billable
	}
	if o.Endepoch != nil {
		toSerialize["endepoch"] = o.Endepoch
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Nonbillable != nil {
		toSerialize["nonbillable"] = o.Nonbillable
	}
	if o.Startepoch != nil {
		toSerialize["startepoch"] = o.Startepoch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20040User struct {
	value *InlineResponse20040User
	isSet bool
}

func (v NullableInlineResponse20040User) Get() *InlineResponse20040User {
	return v.value
}

func (v *NullableInlineResponse20040User) Set(val *InlineResponse20040User) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20040User) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20040User) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20040User(val *InlineResponse20040User) *NullableInlineResponse20040User {
	return &NullableInlineResponse20040User{value: val, isSet: true}
}

func (v NullableInlineResponse20040User) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20040User) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


