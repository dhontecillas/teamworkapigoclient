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

// ProjectEditCustomFieldProject EditCustomFieldProject contains all the information to update a project custom field value.
type ProjectEditCustomFieldProject struct {
	CustomfieldId *int32 `json:"customfieldId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

// NewProjectEditCustomFieldProject instantiates a new ProjectEditCustomFieldProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEditCustomFieldProject() *ProjectEditCustomFieldProject {
	this := ProjectEditCustomFieldProject{}
	return &this
}

// NewProjectEditCustomFieldProjectWithDefaults instantiates a new ProjectEditCustomFieldProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEditCustomFieldProjectWithDefaults() *ProjectEditCustomFieldProject {
	this := ProjectEditCustomFieldProject{}
	return &this
}

// GetCustomfieldId returns the CustomfieldId field value if set, zero value otherwise.
func (o *ProjectEditCustomFieldProject) GetCustomfieldId() int32 {
	if o == nil || o.CustomfieldId == nil {
		var ret int32
		return ret
	}
	return *o.CustomfieldId
}

// GetCustomfieldIdOk returns a tuple with the CustomfieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEditCustomFieldProject) GetCustomfieldIdOk() (*int32, bool) {
	if o == nil || o.CustomfieldId == nil {
		return nil, false
	}
	return o.CustomfieldId, true
}

// HasCustomfieldId returns a boolean if a field has been set.
func (o *ProjectEditCustomFieldProject) HasCustomfieldId() bool {
	if o != nil && o.CustomfieldId != nil {
		return true
	}

	return false
}

// SetCustomfieldId gets a reference to the given int32 and assigns it to the CustomfieldId field.
func (o *ProjectEditCustomFieldProject) SetCustomfieldId(v int32) {
	o.CustomfieldId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectEditCustomFieldProject) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEditCustomFieldProject) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectEditCustomFieldProject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectEditCustomFieldProject) SetId(v int32) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProjectEditCustomFieldProject) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEditCustomFieldProject) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectEditCustomFieldProject) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ProjectEditCustomFieldProject) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o ProjectEditCustomFieldProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomfieldId != nil {
		toSerialize["customfieldId"] = o.CustomfieldId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableProjectEditCustomFieldProject struct {
	value *ProjectEditCustomFieldProject
	isSet bool
}

func (v NullableProjectEditCustomFieldProject) Get() *ProjectEditCustomFieldProject {
	return v.value
}

func (v *NullableProjectEditCustomFieldProject) Set(val *ProjectEditCustomFieldProject) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEditCustomFieldProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEditCustomFieldProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEditCustomFieldProject(val *ProjectEditCustomFieldProject) *NullableProjectEditCustomFieldProject {
	return &NullableProjectEditCustomFieldProject{value: val, isSet: true}
}

func (v NullableProjectEditCustomFieldProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEditCustomFieldProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


