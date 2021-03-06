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

// InlineObject78 struct for InlineObject78
type InlineObject78 struct {
	NotifyIds *string `json:"notifyIds,omitempty"`
	Update *ProjectsProjIdUpdateJsonUpdate `json:"update,omitempty"`
}

// NewInlineObject78 instantiates a new InlineObject78 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject78() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// NewInlineObject78WithDefaults instantiates a new InlineObject78 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject78WithDefaults() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// GetNotifyIds returns the NotifyIds field value if set, zero value otherwise.
func (o *InlineObject78) GetNotifyIds() string {
	if o == nil || o.NotifyIds == nil {
		var ret string
		return ret
	}
	return *o.NotifyIds
}

// GetNotifyIdsOk returns a tuple with the NotifyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject78) GetNotifyIdsOk() (*string, bool) {
	if o == nil || o.NotifyIds == nil {
		return nil, false
	}
	return o.NotifyIds, true
}

// HasNotifyIds returns a boolean if a field has been set.
func (o *InlineObject78) HasNotifyIds() bool {
	if o != nil && o.NotifyIds != nil {
		return true
	}

	return false
}

// SetNotifyIds gets a reference to the given string and assigns it to the NotifyIds field.
func (o *InlineObject78) SetNotifyIds(v string) {
	o.NotifyIds = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *InlineObject78) GetUpdate() ProjectsProjIdUpdateJsonUpdate {
	if o == nil || o.Update == nil {
		var ret ProjectsProjIdUpdateJsonUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject78) GetUpdateOk() (*ProjectsProjIdUpdateJsonUpdate, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *InlineObject78) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given ProjectsProjIdUpdateJsonUpdate and assigns it to the Update field.
func (o *InlineObject78) SetUpdate(v ProjectsProjIdUpdateJsonUpdate) {
	o.Update = &v
}

func (o InlineObject78) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotifyIds != nil {
		toSerialize["notifyIds"] = o.NotifyIds
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject78 struct {
	value *InlineObject78
	isSet bool
}

func (v NullableInlineObject78) Get() *InlineObject78 {
	return v.value
}

func (v *NullableInlineObject78) Set(val *InlineObject78) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject78) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject78) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject78(val *InlineObject78) *NullableInlineObject78 {
	return &NullableInlineObject78{value: val, isSet: true}
}

func (v NullableInlineObject78) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject78) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


