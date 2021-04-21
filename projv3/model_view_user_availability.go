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

// ViewUserAvailability UserAvailability contains all the information returned from a availability.
type ViewUserAvailability struct {
	Dates *map[string]ViewUserAvailabilityDate `json:"dates,omitempty"`
	User *ViewRelationship `json:"user,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
}

// NewViewUserAvailability instantiates a new ViewUserAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewUserAvailability() *ViewUserAvailability {
	this := ViewUserAvailability{}
	return &this
}

// NewViewUserAvailabilityWithDefaults instantiates a new ViewUserAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewUserAvailabilityWithDefaults() *ViewUserAvailability {
	this := ViewUserAvailability{}
	return &this
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *ViewUserAvailability) GetDates() map[string]ViewUserAvailabilityDate {
	if o == nil || o.Dates == nil {
		var ret map[string]ViewUserAvailabilityDate
		return ret
	}
	return *o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserAvailability) GetDatesOk() (*map[string]ViewUserAvailabilityDate, bool) {
	if o == nil || o.Dates == nil {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *ViewUserAvailability) HasDates() bool {
	if o != nil && o.Dates != nil {
		return true
	}

	return false
}

// SetDates gets a reference to the given map[string]ViewUserAvailabilityDate and assigns it to the Dates field.
func (o *ViewUserAvailability) SetDates(v map[string]ViewUserAvailabilityDate) {
	o.Dates = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ViewUserAvailability) GetUser() ViewRelationship {
	if o == nil || o.User == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserAvailability) GetUserOk() (*ViewRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ViewUserAvailability) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ViewRelationship and assigns it to the User field.
func (o *ViewUserAvailability) SetUser(v ViewRelationship) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ViewUserAvailability) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserAvailability) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ViewUserAvailability) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ViewUserAvailability) SetUserId(v int32) {
	o.UserId = &v
}

func (o ViewUserAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dates != nil {
		toSerialize["dates"] = o.Dates
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableViewUserAvailability struct {
	value *ViewUserAvailability
	isSet bool
}

func (v NullableViewUserAvailability) Get() *ViewUserAvailability {
	return v.value
}

func (v *NullableViewUserAvailability) Set(val *ViewUserAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableViewUserAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableViewUserAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewUserAvailability(val *ViewUserAvailability) *NullableViewUserAvailability {
	return &NullableViewUserAvailability{value: val, isSet: true}
}

func (v NullableViewUserAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewUserAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

