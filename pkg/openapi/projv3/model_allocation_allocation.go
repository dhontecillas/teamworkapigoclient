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

// AllocationAllocation Allocation contains all the information returned from a allocation.
type AllocationAllocation struct {
	AssignedUserID *int32 `json:"assignedUserID,omitempty"`
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	DistributeType *string `json:"distributeType,omitempty"`
	// in minutes
	Duration *int32 `json:"duration,omitempty"`
	// Date unmarshals represents a Unified API Spec date format.
	EndedAt *map[string]interface{} `json:"endedAt,omitempty"`
	HoursPerDay *int32 `json:"hoursPerDay,omitempty"`
	IgnoreCollisions *bool `json:"ignoreCollisions,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	RecurringRule *PayloadNullableRRule `json:"recurringRule,omitempty"`
	// Date unmarshals represents a Unified API Spec date format.
	StartedAt *map[string]interface{} `json:"startedAt,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewAllocationAllocation instantiates a new AllocationAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationAllocation() *AllocationAllocation {
	this := AllocationAllocation{}
	return &this
}

// NewAllocationAllocationWithDefaults instantiates a new AllocationAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationAllocationWithDefaults() *AllocationAllocation {
	this := AllocationAllocation{}
	return &this
}

// GetAssignedUserID returns the AssignedUserID field value if set, zero value otherwise.
func (o *AllocationAllocation) GetAssignedUserID() int32 {
	if o == nil || o.AssignedUserID == nil {
		var ret int32
		return ret
	}
	return *o.AssignedUserID
}

// GetAssignedUserIDOk returns a tuple with the AssignedUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetAssignedUserIDOk() (*int32, bool) {
	if o == nil || o.AssignedUserID == nil {
		return nil, false
	}
	return o.AssignedUserID, true
}

// HasAssignedUserID returns a boolean if a field has been set.
func (o *AllocationAllocation) HasAssignedUserID() bool {
	if o != nil && o.AssignedUserID != nil {
		return true
	}

	return false
}

// SetAssignedUserID gets a reference to the given int32 and assigns it to the AssignedUserID field.
func (o *AllocationAllocation) SetAssignedUserID(v int32) {
	o.AssignedUserID = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *AllocationAllocation) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *AllocationAllocation) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *AllocationAllocation) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllocationAllocation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllocationAllocation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllocationAllocation) SetDescription(v string) {
	o.Description = &v
}

// GetDistributeType returns the DistributeType field value if set, zero value otherwise.
func (o *AllocationAllocation) GetDistributeType() string {
	if o == nil || o.DistributeType == nil {
		var ret string
		return ret
	}
	return *o.DistributeType
}

// GetDistributeTypeOk returns a tuple with the DistributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetDistributeTypeOk() (*string, bool) {
	if o == nil || o.DistributeType == nil {
		return nil, false
	}
	return o.DistributeType, true
}

// HasDistributeType returns a boolean if a field has been set.
func (o *AllocationAllocation) HasDistributeType() bool {
	if o != nil && o.DistributeType != nil {
		return true
	}

	return false
}

// SetDistributeType gets a reference to the given string and assigns it to the DistributeType field.
func (o *AllocationAllocation) SetDistributeType(v string) {
	o.DistributeType = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AllocationAllocation) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AllocationAllocation) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AllocationAllocation) SetDuration(v int32) {
	o.Duration = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *AllocationAllocation) GetEndedAt() map[string]interface{} {
	if o == nil || o.EndedAt == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetEndedAtOk() (*map[string]interface{}, bool) {
	if o == nil || o.EndedAt == nil {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *AllocationAllocation) HasEndedAt() bool {
	if o != nil && o.EndedAt != nil {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given map[string]interface{} and assigns it to the EndedAt field.
func (o *AllocationAllocation) SetEndedAt(v map[string]interface{}) {
	o.EndedAt = &v
}

// GetHoursPerDay returns the HoursPerDay field value if set, zero value otherwise.
func (o *AllocationAllocation) GetHoursPerDay() int32 {
	if o == nil || o.HoursPerDay == nil {
		var ret int32
		return ret
	}
	return *o.HoursPerDay
}

// GetHoursPerDayOk returns a tuple with the HoursPerDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetHoursPerDayOk() (*int32, bool) {
	if o == nil || o.HoursPerDay == nil {
		return nil, false
	}
	return o.HoursPerDay, true
}

// HasHoursPerDay returns a boolean if a field has been set.
func (o *AllocationAllocation) HasHoursPerDay() bool {
	if o != nil && o.HoursPerDay != nil {
		return true
	}

	return false
}

// SetHoursPerDay gets a reference to the given int32 and assigns it to the HoursPerDay field.
func (o *AllocationAllocation) SetHoursPerDay(v int32) {
	o.HoursPerDay = &v
}

// GetIgnoreCollisions returns the IgnoreCollisions field value if set, zero value otherwise.
func (o *AllocationAllocation) GetIgnoreCollisions() bool {
	if o == nil || o.IgnoreCollisions == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreCollisions
}

// GetIgnoreCollisionsOk returns a tuple with the IgnoreCollisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetIgnoreCollisionsOk() (*bool, bool) {
	if o == nil || o.IgnoreCollisions == nil {
		return nil, false
	}
	return o.IgnoreCollisions, true
}

// HasIgnoreCollisions returns a boolean if a field has been set.
func (o *AllocationAllocation) HasIgnoreCollisions() bool {
	if o != nil && o.IgnoreCollisions != nil {
		return true
	}

	return false
}

// SetIgnoreCollisions gets a reference to the given bool and assigns it to the IgnoreCollisions field.
func (o *AllocationAllocation) SetIgnoreCollisions(v bool) {
	o.IgnoreCollisions = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AllocationAllocation) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AllocationAllocation) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *AllocationAllocation) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetRecurringRule returns the RecurringRule field value if set, zero value otherwise.
func (o *AllocationAllocation) GetRecurringRule() PayloadNullableRRule {
	if o == nil || o.RecurringRule == nil {
		var ret PayloadNullableRRule
		return ret
	}
	return *o.RecurringRule
}

// GetRecurringRuleOk returns a tuple with the RecurringRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetRecurringRuleOk() (*PayloadNullableRRule, bool) {
	if o == nil || o.RecurringRule == nil {
		return nil, false
	}
	return o.RecurringRule, true
}

// HasRecurringRule returns a boolean if a field has been set.
func (o *AllocationAllocation) HasRecurringRule() bool {
	if o != nil && o.RecurringRule != nil {
		return true
	}

	return false
}

// SetRecurringRule gets a reference to the given PayloadNullableRRule and assigns it to the RecurringRule field.
func (o *AllocationAllocation) SetRecurringRule(v PayloadNullableRRule) {
	o.RecurringRule = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *AllocationAllocation) GetStartedAt() map[string]interface{} {
	if o == nil || o.StartedAt == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetStartedAtOk() (*map[string]interface{}, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *AllocationAllocation) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given map[string]interface{} and assigns it to the StartedAt field.
func (o *AllocationAllocation) SetStartedAt(v map[string]interface{}) {
	o.StartedAt = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AllocationAllocation) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationAllocation) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AllocationAllocation) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AllocationAllocation) SetTitle(v string) {
	o.Title = &v
}

func (o AllocationAllocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedUserID != nil {
		toSerialize["assignedUserID"] = o.AssignedUserID
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DistributeType != nil {
		toSerialize["distributeType"] = o.DistributeType
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.EndedAt != nil {
		toSerialize["endedAt"] = o.EndedAt
	}
	if o.HoursPerDay != nil {
		toSerialize["hoursPerDay"] = o.HoursPerDay
	}
	if o.IgnoreCollisions != nil {
		toSerialize["ignoreCollisions"] = o.IgnoreCollisions
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.RecurringRule != nil {
		toSerialize["recurringRule"] = o.RecurringRule
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableAllocationAllocation struct {
	value *AllocationAllocation
	isSet bool
}

func (v NullableAllocationAllocation) Get() *AllocationAllocation {
	return v.value
}

func (v *NullableAllocationAllocation) Set(val *AllocationAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationAllocation(val *AllocationAllocation) *NullableAllocationAllocation {
	return &NullableAllocationAllocation{value: val, isSet: true}
}

func (v NullableAllocationAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


