/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"encoding/json"
)

// ViewUserUtilization UserUtilization contains all the information returned from a utilization.
type ViewUserUtilization struct {
	Data *[]ViewUserUtilizationData `json:"data,omitempty"`
	User *ViewRelationship `json:"user,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
}

// NewViewUserUtilization instantiates a new ViewUserUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewUserUtilization() *ViewUserUtilization {
	this := ViewUserUtilization{}
	return &this
}

// NewViewUserUtilizationWithDefaults instantiates a new ViewUserUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewUserUtilizationWithDefaults() *ViewUserUtilization {
	this := ViewUserUtilization{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ViewUserUtilization) GetData() []ViewUserUtilizationData {
	if o == nil || o.Data == nil {
		var ret []ViewUserUtilizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserUtilization) GetDataOk() (*[]ViewUserUtilizationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ViewUserUtilization) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ViewUserUtilizationData and assigns it to the Data field.
func (o *ViewUserUtilization) SetData(v []ViewUserUtilizationData) {
	o.Data = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ViewUserUtilization) GetUser() ViewRelationship {
	if o == nil || o.User == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserUtilization) GetUserOk() (*ViewRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ViewUserUtilization) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ViewRelationship and assigns it to the User field.
func (o *ViewUserUtilization) SetUser(v ViewRelationship) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ViewUserUtilization) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserUtilization) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ViewUserUtilization) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ViewUserUtilization) SetUserId(v int32) {
	o.UserId = &v
}

func (o ViewUserUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableViewUserUtilization struct {
	value *ViewUserUtilization
	isSet bool
}

func (v NullableViewUserUtilization) Get() *ViewUserUtilization {
	return v.value
}

func (v *NullableViewUserUtilization) Set(val *ViewUserUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableViewUserUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableViewUserUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewUserUtilization(val *ViewUserUtilization) *NullableViewUserUtilization {
	return &NullableViewUserUtilization{value: val, isSet: true}
}

func (v NullableViewUserUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewUserUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


