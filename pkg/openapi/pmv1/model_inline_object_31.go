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

// InlineObject31 struct for InlineObject31
type InlineObject31 struct {
	Milestone *MilestonesIdJsonMilestone `json:"milestone,omitempty"`
	MoveUpcomingMilestones *bool `json:"move-upcoming-milestones,omitempty"`
	MoveUpcomingMilestonesOffWeekends *bool `json:"move-upcoming-milestones-off-weekends,omitempty"`
}

// NewInlineObject31 instantiates a new InlineObject31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject31() *InlineObject31 {
	this := InlineObject31{}
	return &this
}

// NewInlineObject31WithDefaults instantiates a new InlineObject31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject31WithDefaults() *InlineObject31 {
	this := InlineObject31{}
	return &this
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *InlineObject31) GetMilestone() MilestonesIdJsonMilestone {
	if o == nil || o.Milestone == nil {
		var ret MilestonesIdJsonMilestone
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetMilestoneOk() (*MilestonesIdJsonMilestone, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *InlineObject31) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given MilestonesIdJsonMilestone and assigns it to the Milestone field.
func (o *InlineObject31) SetMilestone(v MilestonesIdJsonMilestone) {
	o.Milestone = &v
}

// GetMoveUpcomingMilestones returns the MoveUpcomingMilestones field value if set, zero value otherwise.
func (o *InlineObject31) GetMoveUpcomingMilestones() bool {
	if o == nil || o.MoveUpcomingMilestones == nil {
		var ret bool
		return ret
	}
	return *o.MoveUpcomingMilestones
}

// GetMoveUpcomingMilestonesOk returns a tuple with the MoveUpcomingMilestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetMoveUpcomingMilestonesOk() (*bool, bool) {
	if o == nil || o.MoveUpcomingMilestones == nil {
		return nil, false
	}
	return o.MoveUpcomingMilestones, true
}

// HasMoveUpcomingMilestones returns a boolean if a field has been set.
func (o *InlineObject31) HasMoveUpcomingMilestones() bool {
	if o != nil && o.MoveUpcomingMilestones != nil {
		return true
	}

	return false
}

// SetMoveUpcomingMilestones gets a reference to the given bool and assigns it to the MoveUpcomingMilestones field.
func (o *InlineObject31) SetMoveUpcomingMilestones(v bool) {
	o.MoveUpcomingMilestones = &v
}

// GetMoveUpcomingMilestonesOffWeekends returns the MoveUpcomingMilestonesOffWeekends field value if set, zero value otherwise.
func (o *InlineObject31) GetMoveUpcomingMilestonesOffWeekends() bool {
	if o == nil || o.MoveUpcomingMilestonesOffWeekends == nil {
		var ret bool
		return ret
	}
	return *o.MoveUpcomingMilestonesOffWeekends
}

// GetMoveUpcomingMilestonesOffWeekendsOk returns a tuple with the MoveUpcomingMilestonesOffWeekends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetMoveUpcomingMilestonesOffWeekendsOk() (*bool, bool) {
	if o == nil || o.MoveUpcomingMilestonesOffWeekends == nil {
		return nil, false
	}
	return o.MoveUpcomingMilestonesOffWeekends, true
}

// HasMoveUpcomingMilestonesOffWeekends returns a boolean if a field has been set.
func (o *InlineObject31) HasMoveUpcomingMilestonesOffWeekends() bool {
	if o != nil && o.MoveUpcomingMilestonesOffWeekends != nil {
		return true
	}

	return false
}

// SetMoveUpcomingMilestonesOffWeekends gets a reference to the given bool and assigns it to the MoveUpcomingMilestonesOffWeekends field.
func (o *InlineObject31) SetMoveUpcomingMilestonesOffWeekends(v bool) {
	o.MoveUpcomingMilestonesOffWeekends = &v
}

func (o InlineObject31) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Milestone != nil {
		toSerialize["milestone"] = o.Milestone
	}
	if o.MoveUpcomingMilestones != nil {
		toSerialize["move-upcoming-milestones"] = o.MoveUpcomingMilestones
	}
	if o.MoveUpcomingMilestonesOffWeekends != nil {
		toSerialize["move-upcoming-milestones-off-weekends"] = o.MoveUpcomingMilestonesOffWeekends
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject31 struct {
	value *InlineObject31
	isSet bool
}

func (v NullableInlineObject31) Get() *InlineObject31 {
	return v.value
}

func (v *NullableInlineObject31) Set(val *InlineObject31) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject31) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject31(val *InlineObject31) *NullableInlineObject31 {
	return &NullableInlineObject31{value: val, isSet: true}
}

func (v NullableInlineObject31) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


