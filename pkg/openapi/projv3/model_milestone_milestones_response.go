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

// MilestoneMilestonesResponse MilestonesResponse contains information about a group of milestones.
type MilestoneMilestonesResponse struct {
	Included *MilestoneMilestonesResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	Milestones *[]ViewMilestone `json:"milestones,omitempty"`
}

// NewMilestoneMilestonesResponse instantiates a new MilestoneMilestonesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMilestoneMilestonesResponse() *MilestoneMilestonesResponse {
	this := MilestoneMilestonesResponse{}
	return &this
}

// NewMilestoneMilestonesResponseWithDefaults instantiates a new MilestoneMilestonesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMilestoneMilestonesResponseWithDefaults() *MilestoneMilestonesResponse {
	this := MilestoneMilestonesResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *MilestoneMilestonesResponse) GetIncluded() MilestoneMilestonesResponseIncluded {
	if o == nil || o.Included == nil {
		var ret MilestoneMilestonesResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneMilestonesResponse) GetIncludedOk() (*MilestoneMilestonesResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *MilestoneMilestonesResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given MilestoneMilestonesResponseIncluded and assigns it to the Included field.
func (o *MilestoneMilestonesResponse) SetIncluded(v MilestoneMilestonesResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MilestoneMilestonesResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneMilestonesResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MilestoneMilestonesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *MilestoneMilestonesResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *MilestoneMilestonesResponse) GetMilestones() []ViewMilestone {
	if o == nil || o.Milestones == nil {
		var ret []ViewMilestone
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneMilestonesResponse) GetMilestonesOk() (*[]ViewMilestone, bool) {
	if o == nil || o.Milestones == nil {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *MilestoneMilestonesResponse) HasMilestones() bool {
	if o != nil && o.Milestones != nil {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given []ViewMilestone and assigns it to the Milestones field.
func (o *MilestoneMilestonesResponse) SetMilestones(v []ViewMilestone) {
	o.Milestones = &v
}

func (o MilestoneMilestonesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Milestones != nil {
		toSerialize["milestones"] = o.Milestones
	}
	return json.Marshal(toSerialize)
}

type NullableMilestoneMilestonesResponse struct {
	value *MilestoneMilestonesResponse
	isSet bool
}

func (v NullableMilestoneMilestonesResponse) Get() *MilestoneMilestonesResponse {
	return v.value
}

func (v *NullableMilestoneMilestonesResponse) Set(val *MilestoneMilestonesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMilestoneMilestonesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMilestoneMilestonesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMilestoneMilestonesResponse(val *MilestoneMilestonesResponse) *NullableMilestoneMilestonesResponse {
	return &NullableMilestoneMilestonesResponse{value: val, isSet: true}
}

func (v NullableMilestoneMilestonesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMilestoneMilestonesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


