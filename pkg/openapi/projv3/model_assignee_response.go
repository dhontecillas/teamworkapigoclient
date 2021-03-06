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

// AssigneeResponse Response contains information about a specific assignee.
type AssigneeResponse struct {
	Assignees *[]ViewFormAssignee `json:"assignees,omitempty"`
	Errors *[]ErrorsBulkError `json:"errors,omitempty"`
	Included *AssigneeFormAssigneesResponseIncluded `json:"included,omitempty"`
}

// NewAssigneeResponse instantiates a new AssigneeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssigneeResponse() *AssigneeResponse {
	this := AssigneeResponse{}
	return &this
}

// NewAssigneeResponseWithDefaults instantiates a new AssigneeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssigneeResponseWithDefaults() *AssigneeResponse {
	this := AssigneeResponse{}
	return &this
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *AssigneeResponse) GetAssignees() []ViewFormAssignee {
	if o == nil || o.Assignees == nil {
		var ret []ViewFormAssignee
		return ret
	}
	return *o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponse) GetAssigneesOk() (*[]ViewFormAssignee, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *AssigneeResponse) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []ViewFormAssignee and assigns it to the Assignees field.
func (o *AssigneeResponse) SetAssignees(v []ViewFormAssignee) {
	o.Assignees = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AssigneeResponse) GetErrors() []ErrorsBulkError {
	if o == nil || o.Errors == nil {
		var ret []ErrorsBulkError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponse) GetErrorsOk() (*[]ErrorsBulkError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AssigneeResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorsBulkError and assigns it to the Errors field.
func (o *AssigneeResponse) SetErrors(v []ErrorsBulkError) {
	o.Errors = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AssigneeResponse) GetIncluded() AssigneeFormAssigneesResponseIncluded {
	if o == nil || o.Included == nil {
		var ret AssigneeFormAssigneesResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponse) GetIncludedOk() (*AssigneeFormAssigneesResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AssigneeResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given AssigneeFormAssigneesResponseIncluded and assigns it to the Included field.
func (o *AssigneeResponse) SetIncluded(v AssigneeFormAssigneesResponseIncluded) {
	o.Included = &v
}

func (o AssigneeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableAssigneeResponse struct {
	value *AssigneeResponse
	isSet bool
}

func (v NullableAssigneeResponse) Get() *AssigneeResponse {
	return v.value
}

func (v *NullableAssigneeResponse) Set(val *AssigneeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssigneeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssigneeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssigneeResponse(val *AssigneeResponse) *NullableAssigneeResponse {
	return &NullableAssigneeResponse{value: val, isSet: true}
}

func (v NullableAssigneeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssigneeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


