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

// PlannerWorkloadPlannersResponseIncluded struct for PlannerWorkloadPlannersResponseIncluded
type PlannerWorkloadPlannersResponseIncluded struct {
	CalendarEvents *map[string]ViewCalendarEvent `json:"calendarEvents,omitempty"`
	Companies *map[string]ViewCompany `json:"companies,omitempty"`
	Milestones *map[string]ViewMilestone `json:"milestones,omitempty"`
	Tasklists *map[string]ViewTasklist `json:"tasklists,omitempty"`
	Tasks *map[string]ViewTask `json:"tasks,omitempty"`
	Timelogs *map[string]ViewTimelog `json:"timelogs,omitempty"`
	Users *map[string]ViewUser `json:"users,omitempty"`
	WorkingHourEntries *map[string]ViewWorkingHourEntry `json:"workingHourEntries,omitempty"`
	WorkingHours *map[string]ViewWorkingHour `json:"workingHours,omitempty"`
}

// NewPlannerWorkloadPlannersResponseIncluded instantiates a new PlannerWorkloadPlannersResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannerWorkloadPlannersResponseIncluded() *PlannerWorkloadPlannersResponseIncluded {
	this := PlannerWorkloadPlannersResponseIncluded{}
	return &this
}

// NewPlannerWorkloadPlannersResponseIncludedWithDefaults instantiates a new PlannerWorkloadPlannersResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannerWorkloadPlannersResponseIncludedWithDefaults() *PlannerWorkloadPlannersResponseIncluded {
	this := PlannerWorkloadPlannersResponseIncluded{}
	return &this
}

// GetCalendarEvents returns the CalendarEvents field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetCalendarEvents() map[string]ViewCalendarEvent {
	if o == nil || o.CalendarEvents == nil {
		var ret map[string]ViewCalendarEvent
		return ret
	}
	return *o.CalendarEvents
}

// GetCalendarEventsOk returns a tuple with the CalendarEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetCalendarEventsOk() (*map[string]ViewCalendarEvent, bool) {
	if o == nil || o.CalendarEvents == nil {
		return nil, false
	}
	return o.CalendarEvents, true
}

// HasCalendarEvents returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasCalendarEvents() bool {
	if o != nil && o.CalendarEvents != nil {
		return true
	}

	return false
}

// SetCalendarEvents gets a reference to the given map[string]ViewCalendarEvent and assigns it to the CalendarEvents field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetCalendarEvents(v map[string]ViewCalendarEvent) {
	o.CalendarEvents = &v
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetCompanies() map[string]ViewCompany {
	if o == nil || o.Companies == nil {
		var ret map[string]ViewCompany
		return ret
	}
	return *o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetCompaniesOk() (*map[string]ViewCompany, bool) {
	if o == nil || o.Companies == nil {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasCompanies() bool {
	if o != nil && o.Companies != nil {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given map[string]ViewCompany and assigns it to the Companies field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetCompanies(v map[string]ViewCompany) {
	o.Companies = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetMilestones() map[string]ViewMilestone {
	if o == nil || o.Milestones == nil {
		var ret map[string]ViewMilestone
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetMilestonesOk() (*map[string]ViewMilestone, bool) {
	if o == nil || o.Milestones == nil {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasMilestones() bool {
	if o != nil && o.Milestones != nil {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given map[string]ViewMilestone and assigns it to the Milestones field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetMilestones(v map[string]ViewMilestone) {
	o.Milestones = &v
}

// GetTasklists returns the Tasklists field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTasklists() map[string]ViewTasklist {
	if o == nil || o.Tasklists == nil {
		var ret map[string]ViewTasklist
		return ret
	}
	return *o.Tasklists
}

// GetTasklistsOk returns a tuple with the Tasklists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTasklistsOk() (*map[string]ViewTasklist, bool) {
	if o == nil || o.Tasklists == nil {
		return nil, false
	}
	return o.Tasklists, true
}

// HasTasklists returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasTasklists() bool {
	if o != nil && o.Tasklists != nil {
		return true
	}

	return false
}

// SetTasklists gets a reference to the given map[string]ViewTasklist and assigns it to the Tasklists field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetTasklists(v map[string]ViewTasklist) {
	o.Tasklists = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTasks() map[string]ViewTask {
	if o == nil || o.Tasks == nil {
		var ret map[string]ViewTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given map[string]ViewTask and assigns it to the Tasks field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetTasks(v map[string]ViewTask) {
	o.Tasks = &v
}

// GetTimelogs returns the Timelogs field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTimelogs() map[string]ViewTimelog {
	if o == nil || o.Timelogs == nil {
		var ret map[string]ViewTimelog
		return ret
	}
	return *o.Timelogs
}

// GetTimelogsOk returns a tuple with the Timelogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetTimelogsOk() (*map[string]ViewTimelog, bool) {
	if o == nil || o.Timelogs == nil {
		return nil, false
	}
	return o.Timelogs, true
}

// HasTimelogs returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasTimelogs() bool {
	if o != nil && o.Timelogs != nil {
		return true
	}

	return false
}

// SetTimelogs gets a reference to the given map[string]ViewTimelog and assigns it to the Timelogs field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetTimelogs(v map[string]ViewTimelog) {
	o.Timelogs = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetUsers() map[string]ViewUser {
	if o == nil || o.Users == nil {
		var ret map[string]ViewUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]ViewUser and assigns it to the Users field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetUsers(v map[string]ViewUser) {
	o.Users = &v
}

// GetWorkingHourEntries returns the WorkingHourEntries field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHourEntries() map[string]ViewWorkingHourEntry {
	if o == nil || o.WorkingHourEntries == nil {
		var ret map[string]ViewWorkingHourEntry
		return ret
	}
	return *o.WorkingHourEntries
}

// GetWorkingHourEntriesOk returns a tuple with the WorkingHourEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHourEntriesOk() (*map[string]ViewWorkingHourEntry, bool) {
	if o == nil || o.WorkingHourEntries == nil {
		return nil, false
	}
	return o.WorkingHourEntries, true
}

// HasWorkingHourEntries returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasWorkingHourEntries() bool {
	if o != nil && o.WorkingHourEntries != nil {
		return true
	}

	return false
}

// SetWorkingHourEntries gets a reference to the given map[string]ViewWorkingHourEntry and assigns it to the WorkingHourEntries field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetWorkingHourEntries(v map[string]ViewWorkingHourEntry) {
	o.WorkingHourEntries = &v
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHours() map[string]ViewWorkingHour {
	if o == nil || o.WorkingHours == nil {
		var ret map[string]ViewWorkingHour
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) GetWorkingHoursOk() (*map[string]ViewWorkingHour, bool) {
	if o == nil || o.WorkingHours == nil {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *PlannerWorkloadPlannersResponseIncluded) HasWorkingHours() bool {
	if o != nil && o.WorkingHours != nil {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given map[string]ViewWorkingHour and assigns it to the WorkingHours field.
func (o *PlannerWorkloadPlannersResponseIncluded) SetWorkingHours(v map[string]ViewWorkingHour) {
	o.WorkingHours = &v
}

func (o PlannerWorkloadPlannersResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalendarEvents != nil {
		toSerialize["calendarEvents"] = o.CalendarEvents
	}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	if o.Milestones != nil {
		toSerialize["milestones"] = o.Milestones
	}
	if o.Tasklists != nil {
		toSerialize["tasklists"] = o.Tasklists
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Timelogs != nil {
		toSerialize["timelogs"] = o.Timelogs
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.WorkingHourEntries != nil {
		toSerialize["workingHourEntries"] = o.WorkingHourEntries
	}
	if o.WorkingHours != nil {
		toSerialize["workingHours"] = o.WorkingHours
	}
	return json.Marshal(toSerialize)
}

type NullablePlannerWorkloadPlannersResponseIncluded struct {
	value *PlannerWorkloadPlannersResponseIncluded
	isSet bool
}

func (v NullablePlannerWorkloadPlannersResponseIncluded) Get() *PlannerWorkloadPlannersResponseIncluded {
	return v.value
}

func (v *NullablePlannerWorkloadPlannersResponseIncluded) Set(val *PlannerWorkloadPlannersResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannerWorkloadPlannersResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannerWorkloadPlannersResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannerWorkloadPlannersResponseIncluded(val *PlannerWorkloadPlannersResponseIncluded) *NullablePlannerWorkloadPlannersResponseIncluded {
	return &NullablePlannerWorkloadPlannersResponseIncluded{value: val, isSet: true}
}

func (v NullablePlannerWorkloadPlannersResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannerWorkloadPlannersResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


