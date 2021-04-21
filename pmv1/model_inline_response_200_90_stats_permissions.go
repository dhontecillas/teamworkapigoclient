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

// InlineResponse20090StatsPermissions struct for InlineResponse20090StatsPermissions
type InlineResponse20090StatsPermissions struct {
	CanAddMessages *bool `json:"canAddMessages,omitempty"`
	CanAddMilestones *bool `json:"canAddMilestones,omitempty"`
	CanAddTasks *bool `json:"canAddTasks,omitempty"`
	CanLogTime *bool `json:"canLogTime,omitempty"`
}

// NewInlineResponse20090StatsPermissions instantiates a new InlineResponse20090StatsPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090StatsPermissions() *InlineResponse20090StatsPermissions {
	this := InlineResponse20090StatsPermissions{}
	return &this
}

// NewInlineResponse20090StatsPermissionsWithDefaults instantiates a new InlineResponse20090StatsPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090StatsPermissionsWithDefaults() *InlineResponse20090StatsPermissions {
	this := InlineResponse20090StatsPermissions{}
	return &this
}

// GetCanAddMessages returns the CanAddMessages field value if set, zero value otherwise.
func (o *InlineResponse20090StatsPermissions) GetCanAddMessages() bool {
	if o == nil || o.CanAddMessages == nil {
		var ret bool
		return ret
	}
	return *o.CanAddMessages
}

// GetCanAddMessagesOk returns a tuple with the CanAddMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090StatsPermissions) GetCanAddMessagesOk() (*bool, bool) {
	if o == nil || o.CanAddMessages == nil {
		return nil, false
	}
	return o.CanAddMessages, true
}

// HasCanAddMessages returns a boolean if a field has been set.
func (o *InlineResponse20090StatsPermissions) HasCanAddMessages() bool {
	if o != nil && o.CanAddMessages != nil {
		return true
	}

	return false
}

// SetCanAddMessages gets a reference to the given bool and assigns it to the CanAddMessages field.
func (o *InlineResponse20090StatsPermissions) SetCanAddMessages(v bool) {
	o.CanAddMessages = &v
}

// GetCanAddMilestones returns the CanAddMilestones field value if set, zero value otherwise.
func (o *InlineResponse20090StatsPermissions) GetCanAddMilestones() bool {
	if o == nil || o.CanAddMilestones == nil {
		var ret bool
		return ret
	}
	return *o.CanAddMilestones
}

// GetCanAddMilestonesOk returns a tuple with the CanAddMilestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090StatsPermissions) GetCanAddMilestonesOk() (*bool, bool) {
	if o == nil || o.CanAddMilestones == nil {
		return nil, false
	}
	return o.CanAddMilestones, true
}

// HasCanAddMilestones returns a boolean if a field has been set.
func (o *InlineResponse20090StatsPermissions) HasCanAddMilestones() bool {
	if o != nil && o.CanAddMilestones != nil {
		return true
	}

	return false
}

// SetCanAddMilestones gets a reference to the given bool and assigns it to the CanAddMilestones field.
func (o *InlineResponse20090StatsPermissions) SetCanAddMilestones(v bool) {
	o.CanAddMilestones = &v
}

// GetCanAddTasks returns the CanAddTasks field value if set, zero value otherwise.
func (o *InlineResponse20090StatsPermissions) GetCanAddTasks() bool {
	if o == nil || o.CanAddTasks == nil {
		var ret bool
		return ret
	}
	return *o.CanAddTasks
}

// GetCanAddTasksOk returns a tuple with the CanAddTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090StatsPermissions) GetCanAddTasksOk() (*bool, bool) {
	if o == nil || o.CanAddTasks == nil {
		return nil, false
	}
	return o.CanAddTasks, true
}

// HasCanAddTasks returns a boolean if a field has been set.
func (o *InlineResponse20090StatsPermissions) HasCanAddTasks() bool {
	if o != nil && o.CanAddTasks != nil {
		return true
	}

	return false
}

// SetCanAddTasks gets a reference to the given bool and assigns it to the CanAddTasks field.
func (o *InlineResponse20090StatsPermissions) SetCanAddTasks(v bool) {
	o.CanAddTasks = &v
}

// GetCanLogTime returns the CanLogTime field value if set, zero value otherwise.
func (o *InlineResponse20090StatsPermissions) GetCanLogTime() bool {
	if o == nil || o.CanLogTime == nil {
		var ret bool
		return ret
	}
	return *o.CanLogTime
}

// GetCanLogTimeOk returns a tuple with the CanLogTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090StatsPermissions) GetCanLogTimeOk() (*bool, bool) {
	if o == nil || o.CanLogTime == nil {
		return nil, false
	}
	return o.CanLogTime, true
}

// HasCanLogTime returns a boolean if a field has been set.
func (o *InlineResponse20090StatsPermissions) HasCanLogTime() bool {
	if o != nil && o.CanLogTime != nil {
		return true
	}

	return false
}

// SetCanLogTime gets a reference to the given bool and assigns it to the CanLogTime field.
func (o *InlineResponse20090StatsPermissions) SetCanLogTime(v bool) {
	o.CanLogTime = &v
}

func (o InlineResponse20090StatsPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanAddMessages != nil {
		toSerialize["canAddMessages"] = o.CanAddMessages
	}
	if o.CanAddMilestones != nil {
		toSerialize["canAddMilestones"] = o.CanAddMilestones
	}
	if o.CanAddTasks != nil {
		toSerialize["canAddTasks"] = o.CanAddTasks
	}
	if o.CanLogTime != nil {
		toSerialize["canLogTime"] = o.CanLogTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090StatsPermissions struct {
	value *InlineResponse20090StatsPermissions
	isSet bool
}

func (v NullableInlineResponse20090StatsPermissions) Get() *InlineResponse20090StatsPermissions {
	return v.value
}

func (v *NullableInlineResponse20090StatsPermissions) Set(val *InlineResponse20090StatsPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090StatsPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090StatsPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090StatsPermissions(val *InlineResponse20090StatsPermissions) *NullableInlineResponse20090StatsPermissions {
	return &NullableInlineResponse20090StatsPermissions{value: val, isSet: true}
}

func (v NullableInlineResponse20090StatsPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090StatsPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

