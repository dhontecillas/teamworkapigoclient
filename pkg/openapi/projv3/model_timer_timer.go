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

// TimerTimer Timer contains all the information returned from a timer.
type TimerTimer struct {
	Description *string `json:"description,omitempty"`
	IsBillable *bool `json:"isBillable,omitempty"`
	IsRunning *bool `json:"isRunning,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	Seconds *int32 `json:"seconds,omitempty"`
	StopRunningTimers *bool `json:"stopRunningTimers,omitempty"`
	TaskId *int32 `json:"taskId,omitempty"`
}

// NewTimerTimer instantiates a new TimerTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerTimer() *TimerTimer {
	this := TimerTimer{}
	return &this
}

// NewTimerTimerWithDefaults instantiates a new TimerTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerTimerWithDefaults() *TimerTimer {
	this := TimerTimer{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TimerTimer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TimerTimer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TimerTimer) SetDescription(v string) {
	o.Description = &v
}

// GetIsBillable returns the IsBillable field value if set, zero value otherwise.
func (o *TimerTimer) GetIsBillable() bool {
	if o == nil || o.IsBillable == nil {
		var ret bool
		return ret
	}
	return *o.IsBillable
}

// GetIsBillableOk returns a tuple with the IsBillable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetIsBillableOk() (*bool, bool) {
	if o == nil || o.IsBillable == nil {
		return nil, false
	}
	return o.IsBillable, true
}

// HasIsBillable returns a boolean if a field has been set.
func (o *TimerTimer) HasIsBillable() bool {
	if o != nil && o.IsBillable != nil {
		return true
	}

	return false
}

// SetIsBillable gets a reference to the given bool and assigns it to the IsBillable field.
func (o *TimerTimer) SetIsBillable(v bool) {
	o.IsBillable = &v
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *TimerTimer) GetIsRunning() bool {
	if o == nil || o.IsRunning == nil {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetIsRunningOk() (*bool, bool) {
	if o == nil || o.IsRunning == nil {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *TimerTimer) HasIsRunning() bool {
	if o != nil && o.IsRunning != nil {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *TimerTimer) SetIsRunning(v bool) {
	o.IsRunning = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *TimerTimer) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TimerTimer) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *TimerTimer) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *TimerTimer) GetSeconds() int32 {
	if o == nil || o.Seconds == nil {
		var ret int32
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetSecondsOk() (*int32, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *TimerTimer) HasSeconds() bool {
	if o != nil && o.Seconds != nil {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int32 and assigns it to the Seconds field.
func (o *TimerTimer) SetSeconds(v int32) {
	o.Seconds = &v
}

// GetStopRunningTimers returns the StopRunningTimers field value if set, zero value otherwise.
func (o *TimerTimer) GetStopRunningTimers() bool {
	if o == nil || o.StopRunningTimers == nil {
		var ret bool
		return ret
	}
	return *o.StopRunningTimers
}

// GetStopRunningTimersOk returns a tuple with the StopRunningTimers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetStopRunningTimersOk() (*bool, bool) {
	if o == nil || o.StopRunningTimers == nil {
		return nil, false
	}
	return o.StopRunningTimers, true
}

// HasStopRunningTimers returns a boolean if a field has been set.
func (o *TimerTimer) HasStopRunningTimers() bool {
	if o != nil && o.StopRunningTimers != nil {
		return true
	}

	return false
}

// SetStopRunningTimers gets a reference to the given bool and assigns it to the StopRunningTimers field.
func (o *TimerTimer) SetStopRunningTimers(v bool) {
	o.StopRunningTimers = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TimerTimer) GetTaskId() int32 {
	if o == nil || o.TaskId == nil {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerTimer) GetTaskIdOk() (*int32, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TimerTimer) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *TimerTimer) SetTaskId(v int32) {
	o.TaskId = &v
}

func (o TimerTimer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsBillable != nil {
		toSerialize["isBillable"] = o.IsBillable
	}
	if o.IsRunning != nil {
		toSerialize["isRunning"] = o.IsRunning
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}
	if o.StopRunningTimers != nil {
		toSerialize["stopRunningTimers"] = o.StopRunningTimers
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	return json.Marshal(toSerialize)
}

type NullableTimerTimer struct {
	value *TimerTimer
	isSet bool
}

func (v NullableTimerTimer) Get() *TimerTimer {
	return v.value
}

func (v *NullableTimerTimer) Set(val *TimerTimer) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerTimer(val *TimerTimer) *NullableTimerTimer {
	return &NullableTimerTimer{value: val, isSet: true}
}

func (v NullableTimerTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


