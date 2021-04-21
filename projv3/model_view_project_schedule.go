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

// ViewProjectSchedule ProjectSchedule contains all the information returned from a project schedule.
type ViewProjectSchedule struct {
	AllocatedTotalMinutes *int32 `json:"allocatedTotalMinutes,omitempty"`
	BudgetIds *[]int32 `json:"budgetIds,omitempty"`
	BudgetMinutes *int32 `json:"budgetMinutes,omitempty"`
	Budgets *[]ViewRelationship `json:"budgets,omitempty"`
	LoggedTotalMinutes *int32 `json:"loggedTotalMinutes,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	Schedule *map[string]ViewProjectScheduleEntry `json:"schedule,omitempty"`
	UnavailableTotalMinutes *int32 `json:"unavailableTotalMinutes,omitempty"`
	UserIds *[]int32 `json:"userIds,omitempty"`
	Users *[]ViewRelationship `json:"users,omitempty"`
}

// NewViewProjectSchedule instantiates a new ViewProjectSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProjectSchedule() *ViewProjectSchedule {
	this := ViewProjectSchedule{}
	return &this
}

// NewViewProjectScheduleWithDefaults instantiates a new ViewProjectSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProjectScheduleWithDefaults() *ViewProjectSchedule {
	this := ViewProjectSchedule{}
	return &this
}

// GetAllocatedTotalMinutes returns the AllocatedTotalMinutes field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetAllocatedTotalMinutes() int32 {
	if o == nil || o.AllocatedTotalMinutes == nil {
		var ret int32
		return ret
	}
	return *o.AllocatedTotalMinutes
}

// GetAllocatedTotalMinutesOk returns a tuple with the AllocatedTotalMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetAllocatedTotalMinutesOk() (*int32, bool) {
	if o == nil || o.AllocatedTotalMinutes == nil {
		return nil, false
	}
	return o.AllocatedTotalMinutes, true
}

// HasAllocatedTotalMinutes returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasAllocatedTotalMinutes() bool {
	if o != nil && o.AllocatedTotalMinutes != nil {
		return true
	}

	return false
}

// SetAllocatedTotalMinutes gets a reference to the given int32 and assigns it to the AllocatedTotalMinutes field.
func (o *ViewProjectSchedule) SetAllocatedTotalMinutes(v int32) {
	o.AllocatedTotalMinutes = &v
}

// GetBudgetIds returns the BudgetIds field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetBudgetIds() []int32 {
	if o == nil || o.BudgetIds == nil {
		var ret []int32
		return ret
	}
	return *o.BudgetIds
}

// GetBudgetIdsOk returns a tuple with the BudgetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetBudgetIdsOk() (*[]int32, bool) {
	if o == nil || o.BudgetIds == nil {
		return nil, false
	}
	return o.BudgetIds, true
}

// HasBudgetIds returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasBudgetIds() bool {
	if o != nil && o.BudgetIds != nil {
		return true
	}

	return false
}

// SetBudgetIds gets a reference to the given []int32 and assigns it to the BudgetIds field.
func (o *ViewProjectSchedule) SetBudgetIds(v []int32) {
	o.BudgetIds = &v
}

// GetBudgetMinutes returns the BudgetMinutes field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetBudgetMinutes() int32 {
	if o == nil || o.BudgetMinutes == nil {
		var ret int32
		return ret
	}
	return *o.BudgetMinutes
}

// GetBudgetMinutesOk returns a tuple with the BudgetMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetBudgetMinutesOk() (*int32, bool) {
	if o == nil || o.BudgetMinutes == nil {
		return nil, false
	}
	return o.BudgetMinutes, true
}

// HasBudgetMinutes returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasBudgetMinutes() bool {
	if o != nil && o.BudgetMinutes != nil {
		return true
	}

	return false
}

// SetBudgetMinutes gets a reference to the given int32 and assigns it to the BudgetMinutes field.
func (o *ViewProjectSchedule) SetBudgetMinutes(v int32) {
	o.BudgetMinutes = &v
}

// GetBudgets returns the Budgets field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetBudgets() []ViewRelationship {
	if o == nil || o.Budgets == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Budgets
}

// GetBudgetsOk returns a tuple with the Budgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetBudgetsOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Budgets == nil {
		return nil, false
	}
	return o.Budgets, true
}

// HasBudgets returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasBudgets() bool {
	if o != nil && o.Budgets != nil {
		return true
	}

	return false
}

// SetBudgets gets a reference to the given []ViewRelationship and assigns it to the Budgets field.
func (o *ViewProjectSchedule) SetBudgets(v []ViewRelationship) {
	o.Budgets = &v
}

// GetLoggedTotalMinutes returns the LoggedTotalMinutes field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetLoggedTotalMinutes() int32 {
	if o == nil || o.LoggedTotalMinutes == nil {
		var ret int32
		return ret
	}
	return *o.LoggedTotalMinutes
}

// GetLoggedTotalMinutesOk returns a tuple with the LoggedTotalMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetLoggedTotalMinutesOk() (*int32, bool) {
	if o == nil || o.LoggedTotalMinutes == nil {
		return nil, false
	}
	return o.LoggedTotalMinutes, true
}

// HasLoggedTotalMinutes returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasLoggedTotalMinutes() bool {
	if o != nil && o.LoggedTotalMinutes != nil {
		return true
	}

	return false
}

// SetLoggedTotalMinutes gets a reference to the given int32 and assigns it to the LoggedTotalMinutes field.
func (o *ViewProjectSchedule) SetLoggedTotalMinutes(v int32) {
	o.LoggedTotalMinutes = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewProjectSchedule) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewProjectSchedule) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetSchedule() map[string]ViewProjectScheduleEntry {
	if o == nil || o.Schedule == nil {
		var ret map[string]ViewProjectScheduleEntry
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetScheduleOk() (*map[string]ViewProjectScheduleEntry, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given map[string]ViewProjectScheduleEntry and assigns it to the Schedule field.
func (o *ViewProjectSchedule) SetSchedule(v map[string]ViewProjectScheduleEntry) {
	o.Schedule = &v
}

// GetUnavailableTotalMinutes returns the UnavailableTotalMinutes field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetUnavailableTotalMinutes() int32 {
	if o == nil || o.UnavailableTotalMinutes == nil {
		var ret int32
		return ret
	}
	return *o.UnavailableTotalMinutes
}

// GetUnavailableTotalMinutesOk returns a tuple with the UnavailableTotalMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetUnavailableTotalMinutesOk() (*int32, bool) {
	if o == nil || o.UnavailableTotalMinutes == nil {
		return nil, false
	}
	return o.UnavailableTotalMinutes, true
}

// HasUnavailableTotalMinutes returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasUnavailableTotalMinutes() bool {
	if o != nil && o.UnavailableTotalMinutes != nil {
		return true
	}

	return false
}

// SetUnavailableTotalMinutes gets a reference to the given int32 and assigns it to the UnavailableTotalMinutes field.
func (o *ViewProjectSchedule) SetUnavailableTotalMinutes(v int32) {
	o.UnavailableTotalMinutes = &v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetUserIds() []int32 {
	if o == nil || o.UserIds == nil {
		var ret []int32
		return ret
	}
	return *o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetUserIdsOk() (*[]int32, bool) {
	if o == nil || o.UserIds == nil {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasUserIds() bool {
	if o != nil && o.UserIds != nil {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []int32 and assigns it to the UserIds field.
func (o *ViewProjectSchedule) SetUserIds(v []int32) {
	o.UserIds = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ViewProjectSchedule) GetUsers() []ViewRelationship {
	if o == nil || o.Users == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectSchedule) GetUsersOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ViewProjectSchedule) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ViewRelationship and assigns it to the Users field.
func (o *ViewProjectSchedule) SetUsers(v []ViewRelationship) {
	o.Users = &v
}

func (o ViewProjectSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllocatedTotalMinutes != nil {
		toSerialize["allocatedTotalMinutes"] = o.AllocatedTotalMinutes
	}
	if o.BudgetIds != nil {
		toSerialize["budgetIds"] = o.BudgetIds
	}
	if o.BudgetMinutes != nil {
		toSerialize["budgetMinutes"] = o.BudgetMinutes
	}
	if o.Budgets != nil {
		toSerialize["budgets"] = o.Budgets
	}
	if o.LoggedTotalMinutes != nil {
		toSerialize["loggedTotalMinutes"] = o.LoggedTotalMinutes
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.UnavailableTotalMinutes != nil {
		toSerialize["unavailableTotalMinutes"] = o.UnavailableTotalMinutes
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableViewProjectSchedule struct {
	value *ViewProjectSchedule
	isSet bool
}

func (v NullableViewProjectSchedule) Get() *ViewProjectSchedule {
	return v.value
}

func (v *NullableViewProjectSchedule) Set(val *ViewProjectSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProjectSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProjectSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProjectSchedule(val *ViewProjectSchedule) *NullableViewProjectSchedule {
	return &NullableViewProjectSchedule{value: val, isSet: true}
}

func (v NullableViewProjectSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProjectSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

