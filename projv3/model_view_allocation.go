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

// ViewAllocation Allocation contains all the information returned from a allocation.
type ViewAllocation struct {
	AssignedUser *ViewRelationship `json:"assignedUser,omitempty"`
	AssignedUserID *int32 `json:"assignedUserID,omitempty"`
	// in minutes
	AvailableDuration *int32 `json:"availableDuration,omitempty"`
	Color *string `json:"color,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *int32 `json:"createdBy,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	DeletedBy *int32 `json:"deletedBy,omitempty"`
	Description *string `json:"description,omitempty"`
	DistributeType *string `json:"distributeType,omitempty"`
	// in minutes
	Duration *int32 `json:"duration,omitempty"`
	// Date represents a Unified API Spec date format.
	EndedAt *map[string]interface{} `json:"endedAt,omitempty"`
	HoursPerDay *float32 `json:"hoursPerDay,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Installation *ViewRelationship `json:"installation,omitempty"`
	InstallationId *int32 `json:"installationId,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	RecurringRule *string `json:"recurringRule,omitempty"`
	// Date represents a Unified API Spec date format.
	StartedAt *map[string]interface{} `json:"startedAt,omitempty"`
	Status *string `json:"status,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *int32 `json:"updatedBy,omitempty"`
}

// NewViewAllocation instantiates a new ViewAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewAllocation() *ViewAllocation {
	this := ViewAllocation{}
	return &this
}

// NewViewAllocationWithDefaults instantiates a new ViewAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewAllocationWithDefaults() *ViewAllocation {
	this := ViewAllocation{}
	return &this
}

// GetAssignedUser returns the AssignedUser field value if set, zero value otherwise.
func (o *ViewAllocation) GetAssignedUser() ViewRelationship {
	if o == nil || o.AssignedUser == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.AssignedUser
}

// GetAssignedUserOk returns a tuple with the AssignedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetAssignedUserOk() (*ViewRelationship, bool) {
	if o == nil || o.AssignedUser == nil {
		return nil, false
	}
	return o.AssignedUser, true
}

// HasAssignedUser returns a boolean if a field has been set.
func (o *ViewAllocation) HasAssignedUser() bool {
	if o != nil && o.AssignedUser != nil {
		return true
	}

	return false
}

// SetAssignedUser gets a reference to the given ViewRelationship and assigns it to the AssignedUser field.
func (o *ViewAllocation) SetAssignedUser(v ViewRelationship) {
	o.AssignedUser = &v
}

// GetAssignedUserID returns the AssignedUserID field value if set, zero value otherwise.
func (o *ViewAllocation) GetAssignedUserID() int32 {
	if o == nil || o.AssignedUserID == nil {
		var ret int32
		return ret
	}
	return *o.AssignedUserID
}

// GetAssignedUserIDOk returns a tuple with the AssignedUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetAssignedUserIDOk() (*int32, bool) {
	if o == nil || o.AssignedUserID == nil {
		return nil, false
	}
	return o.AssignedUserID, true
}

// HasAssignedUserID returns a boolean if a field has been set.
func (o *ViewAllocation) HasAssignedUserID() bool {
	if o != nil && o.AssignedUserID != nil {
		return true
	}

	return false
}

// SetAssignedUserID gets a reference to the given int32 and assigns it to the AssignedUserID field.
func (o *ViewAllocation) SetAssignedUserID(v int32) {
	o.AssignedUserID = &v
}

// GetAvailableDuration returns the AvailableDuration field value if set, zero value otherwise.
func (o *ViewAllocation) GetAvailableDuration() int32 {
	if o == nil || o.AvailableDuration == nil {
		var ret int32
		return ret
	}
	return *o.AvailableDuration
}

// GetAvailableDurationOk returns a tuple with the AvailableDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetAvailableDurationOk() (*int32, bool) {
	if o == nil || o.AvailableDuration == nil {
		return nil, false
	}
	return o.AvailableDuration, true
}

// HasAvailableDuration returns a boolean if a field has been set.
func (o *ViewAllocation) HasAvailableDuration() bool {
	if o != nil && o.AvailableDuration != nil {
		return true
	}

	return false
}

// SetAvailableDuration gets a reference to the given int32 and assigns it to the AvailableDuration field.
func (o *ViewAllocation) SetAvailableDuration(v int32) {
	o.AvailableDuration = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ViewAllocation) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ViewAllocation) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ViewAllocation) SetColor(v string) {
	o.Color = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ViewAllocation) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ViewAllocation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ViewAllocation) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ViewAllocation) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetCreatedByOk() (*int32, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ViewAllocation) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *ViewAllocation) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ViewAllocation) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ViewAllocation) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ViewAllocation) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *ViewAllocation) GetDeletedBy() int32 {
	if o == nil || o.DeletedBy == nil {
		var ret int32
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetDeletedByOk() (*int32, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *ViewAllocation) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given int32 and assigns it to the DeletedBy field.
func (o *ViewAllocation) SetDeletedBy(v int32) {
	o.DeletedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ViewAllocation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ViewAllocation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ViewAllocation) SetDescription(v string) {
	o.Description = &v
}

// GetDistributeType returns the DistributeType field value if set, zero value otherwise.
func (o *ViewAllocation) GetDistributeType() string {
	if o == nil || o.DistributeType == nil {
		var ret string
		return ret
	}
	return *o.DistributeType
}

// GetDistributeTypeOk returns a tuple with the DistributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetDistributeTypeOk() (*string, bool) {
	if o == nil || o.DistributeType == nil {
		return nil, false
	}
	return o.DistributeType, true
}

// HasDistributeType returns a boolean if a field has been set.
func (o *ViewAllocation) HasDistributeType() bool {
	if o != nil && o.DistributeType != nil {
		return true
	}

	return false
}

// SetDistributeType gets a reference to the given string and assigns it to the DistributeType field.
func (o *ViewAllocation) SetDistributeType(v string) {
	o.DistributeType = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ViewAllocation) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ViewAllocation) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *ViewAllocation) SetDuration(v int32) {
	o.Duration = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *ViewAllocation) GetEndedAt() map[string]interface{} {
	if o == nil || o.EndedAt == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetEndedAtOk() (*map[string]interface{}, bool) {
	if o == nil || o.EndedAt == nil {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *ViewAllocation) HasEndedAt() bool {
	if o != nil && o.EndedAt != nil {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given map[string]interface{} and assigns it to the EndedAt field.
func (o *ViewAllocation) SetEndedAt(v map[string]interface{}) {
	o.EndedAt = &v
}

// GetHoursPerDay returns the HoursPerDay field value if set, zero value otherwise.
func (o *ViewAllocation) GetHoursPerDay() float32 {
	if o == nil || o.HoursPerDay == nil {
		var ret float32
		return ret
	}
	return *o.HoursPerDay
}

// GetHoursPerDayOk returns a tuple with the HoursPerDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetHoursPerDayOk() (*float32, bool) {
	if o == nil || o.HoursPerDay == nil {
		return nil, false
	}
	return o.HoursPerDay, true
}

// HasHoursPerDay returns a boolean if a field has been set.
func (o *ViewAllocation) HasHoursPerDay() bool {
	if o != nil && o.HoursPerDay != nil {
		return true
	}

	return false
}

// SetHoursPerDay gets a reference to the given float32 and assigns it to the HoursPerDay field.
func (o *ViewAllocation) SetHoursPerDay(v float32) {
	o.HoursPerDay = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewAllocation) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewAllocation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewAllocation) SetId(v int32) {
	o.Id = &v
}

// GetInstallation returns the Installation field value if set, zero value otherwise.
func (o *ViewAllocation) GetInstallation() ViewRelationship {
	if o == nil || o.Installation == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Installation
}

// GetInstallationOk returns a tuple with the Installation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetInstallationOk() (*ViewRelationship, bool) {
	if o == nil || o.Installation == nil {
		return nil, false
	}
	return o.Installation, true
}

// HasInstallation returns a boolean if a field has been set.
func (o *ViewAllocation) HasInstallation() bool {
	if o != nil && o.Installation != nil {
		return true
	}

	return false
}

// SetInstallation gets a reference to the given ViewRelationship and assigns it to the Installation field.
func (o *ViewAllocation) SetInstallation(v ViewRelationship) {
	o.Installation = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *ViewAllocation) GetInstallationId() int32 {
	if o == nil || o.InstallationId == nil {
		var ret int32
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetInstallationIdOk() (*int32, bool) {
	if o == nil || o.InstallationId == nil {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *ViewAllocation) HasInstallationId() bool {
	if o != nil && o.InstallationId != nil {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given int32 and assigns it to the InstallationId field.
func (o *ViewAllocation) SetInstallationId(v int32) {
	o.InstallationId = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewAllocation) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewAllocation) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewAllocation) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewAllocation) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewAllocation) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewAllocation) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetRecurringRule returns the RecurringRule field value if set, zero value otherwise.
func (o *ViewAllocation) GetRecurringRule() string {
	if o == nil || o.RecurringRule == nil {
		var ret string
		return ret
	}
	return *o.RecurringRule
}

// GetRecurringRuleOk returns a tuple with the RecurringRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetRecurringRuleOk() (*string, bool) {
	if o == nil || o.RecurringRule == nil {
		return nil, false
	}
	return o.RecurringRule, true
}

// HasRecurringRule returns a boolean if a field has been set.
func (o *ViewAllocation) HasRecurringRule() bool {
	if o != nil && o.RecurringRule != nil {
		return true
	}

	return false
}

// SetRecurringRule gets a reference to the given string and assigns it to the RecurringRule field.
func (o *ViewAllocation) SetRecurringRule(v string) {
	o.RecurringRule = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *ViewAllocation) GetStartedAt() map[string]interface{} {
	if o == nil || o.StartedAt == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetStartedAtOk() (*map[string]interface{}, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ViewAllocation) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given map[string]interface{} and assigns it to the StartedAt field.
func (o *ViewAllocation) SetStartedAt(v map[string]interface{}) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ViewAllocation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ViewAllocation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ViewAllocation) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewAllocation) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewAllocation) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewAllocation) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ViewAllocation) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ViewAllocation) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ViewAllocation) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ViewAllocation) GetUpdatedBy() int32 {
	if o == nil || o.UpdatedBy == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllocation) GetUpdatedByOk() (*int32, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ViewAllocation) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *ViewAllocation) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

func (o ViewAllocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedUser != nil {
		toSerialize["assignedUser"] = o.AssignedUser
	}
	if o.AssignedUserID != nil {
		toSerialize["assignedUserID"] = o.AssignedUserID
	}
	if o.AvailableDuration != nil {
		toSerialize["availableDuration"] = o.AvailableDuration
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deletedBy"] = o.DeletedBy
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Installation != nil {
		toSerialize["installation"] = o.Installation
	}
	if o.InstallationId != nil {
		toSerialize["installationId"] = o.InstallationId
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableViewAllocation struct {
	value *ViewAllocation
	isSet bool
}

func (v NullableViewAllocation) Get() *ViewAllocation {
	return v.value
}

func (v *NullableViewAllocation) Set(val *ViewAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableViewAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableViewAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewAllocation(val *ViewAllocation) *NullableViewAllocation {
	return &NullableViewAllocation{value: val, isSet: true}
}

func (v NullableViewAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

