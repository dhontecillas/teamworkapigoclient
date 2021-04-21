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

// RiskRisk Risk represents a view of a risk.
type RiskRisk struct {
	CanEdit *bool `json:"canEdit,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *int32 `json:"createdBy,omitempty"`
	CreatedByUserId *int32 `json:"createdByUserId,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactCost *bool `json:"impactCost,omitempty"`
	ImpactPerformance *bool `json:"impactPerformance,omitempty"`
	ImpactSchedule *bool `json:"impactSchedule,omitempty"`
	ImpactValue *int32 `json:"impactValue,omitempty"`
	LastChangedByUserId *int32 `json:"lastChangedByUserId,omitempty"`
	LastChangedOn *string `json:"lastChangedOn,omitempty"`
	MitigationPlan *string `json:"mitigationPlan,omitempty"`
	Probability *string `json:"probability,omitempty"`
	ProbabilityValue *int32 `json:"probabilityValue,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	Result *int32 `json:"result,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *int32 `json:"updatedBy,omitempty"`
}

// NewRiskRisk instantiates a new RiskRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRisk() *RiskRisk {
	this := RiskRisk{}
	return &this
}

// NewRiskRiskWithDefaults instantiates a new RiskRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRiskWithDefaults() *RiskRisk {
	this := RiskRisk{}
	return &this
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *RiskRisk) GetCanEdit() bool {
	if o == nil || o.CanEdit == nil {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetCanEditOk() (*bool, bool) {
	if o == nil || o.CanEdit == nil {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *RiskRisk) HasCanEdit() bool {
	if o != nil && o.CanEdit != nil {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *RiskRisk) SetCanEdit(v bool) {
	o.CanEdit = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskRisk) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskRisk) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskRisk) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RiskRisk) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetCreatedByOk() (*int32, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RiskRisk) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *RiskRisk) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *RiskRisk) GetCreatedByUserId() int32 {
	if o == nil || o.CreatedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *RiskRisk) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *RiskRisk) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *RiskRisk) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *RiskRisk) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *RiskRisk) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *RiskRisk) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *RiskRisk) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *RiskRisk) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskRisk) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskRisk) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RiskRisk) SetId(v int32) {
	o.Id = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *RiskRisk) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *RiskRisk) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *RiskRisk) SetImpact(v string) {
	o.Impact = &v
}

// GetImpactCost returns the ImpactCost field value if set, zero value otherwise.
func (o *RiskRisk) GetImpactCost() bool {
	if o == nil || o.ImpactCost == nil {
		var ret bool
		return ret
	}
	return *o.ImpactCost
}

// GetImpactCostOk returns a tuple with the ImpactCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetImpactCostOk() (*bool, bool) {
	if o == nil || o.ImpactCost == nil {
		return nil, false
	}
	return o.ImpactCost, true
}

// HasImpactCost returns a boolean if a field has been set.
func (o *RiskRisk) HasImpactCost() bool {
	if o != nil && o.ImpactCost != nil {
		return true
	}

	return false
}

// SetImpactCost gets a reference to the given bool and assigns it to the ImpactCost field.
func (o *RiskRisk) SetImpactCost(v bool) {
	o.ImpactCost = &v
}

// GetImpactPerformance returns the ImpactPerformance field value if set, zero value otherwise.
func (o *RiskRisk) GetImpactPerformance() bool {
	if o == nil || o.ImpactPerformance == nil {
		var ret bool
		return ret
	}
	return *o.ImpactPerformance
}

// GetImpactPerformanceOk returns a tuple with the ImpactPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetImpactPerformanceOk() (*bool, bool) {
	if o == nil || o.ImpactPerformance == nil {
		return nil, false
	}
	return o.ImpactPerformance, true
}

// HasImpactPerformance returns a boolean if a field has been set.
func (o *RiskRisk) HasImpactPerformance() bool {
	if o != nil && o.ImpactPerformance != nil {
		return true
	}

	return false
}

// SetImpactPerformance gets a reference to the given bool and assigns it to the ImpactPerformance field.
func (o *RiskRisk) SetImpactPerformance(v bool) {
	o.ImpactPerformance = &v
}

// GetImpactSchedule returns the ImpactSchedule field value if set, zero value otherwise.
func (o *RiskRisk) GetImpactSchedule() bool {
	if o == nil || o.ImpactSchedule == nil {
		var ret bool
		return ret
	}
	return *o.ImpactSchedule
}

// GetImpactScheduleOk returns a tuple with the ImpactSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetImpactScheduleOk() (*bool, bool) {
	if o == nil || o.ImpactSchedule == nil {
		return nil, false
	}
	return o.ImpactSchedule, true
}

// HasImpactSchedule returns a boolean if a field has been set.
func (o *RiskRisk) HasImpactSchedule() bool {
	if o != nil && o.ImpactSchedule != nil {
		return true
	}

	return false
}

// SetImpactSchedule gets a reference to the given bool and assigns it to the ImpactSchedule field.
func (o *RiskRisk) SetImpactSchedule(v bool) {
	o.ImpactSchedule = &v
}

// GetImpactValue returns the ImpactValue field value if set, zero value otherwise.
func (o *RiskRisk) GetImpactValue() int32 {
	if o == nil || o.ImpactValue == nil {
		var ret int32
		return ret
	}
	return *o.ImpactValue
}

// GetImpactValueOk returns a tuple with the ImpactValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetImpactValueOk() (*int32, bool) {
	if o == nil || o.ImpactValue == nil {
		return nil, false
	}
	return o.ImpactValue, true
}

// HasImpactValue returns a boolean if a field has been set.
func (o *RiskRisk) HasImpactValue() bool {
	if o != nil && o.ImpactValue != nil {
		return true
	}

	return false
}

// SetImpactValue gets a reference to the given int32 and assigns it to the ImpactValue field.
func (o *RiskRisk) SetImpactValue(v int32) {
	o.ImpactValue = &v
}

// GetLastChangedByUserId returns the LastChangedByUserId field value if set, zero value otherwise.
func (o *RiskRisk) GetLastChangedByUserId() int32 {
	if o == nil || o.LastChangedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.LastChangedByUserId
}

// GetLastChangedByUserIdOk returns a tuple with the LastChangedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetLastChangedByUserIdOk() (*int32, bool) {
	if o == nil || o.LastChangedByUserId == nil {
		return nil, false
	}
	return o.LastChangedByUserId, true
}

// HasLastChangedByUserId returns a boolean if a field has been set.
func (o *RiskRisk) HasLastChangedByUserId() bool {
	if o != nil && o.LastChangedByUserId != nil {
		return true
	}

	return false
}

// SetLastChangedByUserId gets a reference to the given int32 and assigns it to the LastChangedByUserId field.
func (o *RiskRisk) SetLastChangedByUserId(v int32) {
	o.LastChangedByUserId = &v
}

// GetLastChangedOn returns the LastChangedOn field value if set, zero value otherwise.
func (o *RiskRisk) GetLastChangedOn() string {
	if o == nil || o.LastChangedOn == nil {
		var ret string
		return ret
	}
	return *o.LastChangedOn
}

// GetLastChangedOnOk returns a tuple with the LastChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetLastChangedOnOk() (*string, bool) {
	if o == nil || o.LastChangedOn == nil {
		return nil, false
	}
	return o.LastChangedOn, true
}

// HasLastChangedOn returns a boolean if a field has been set.
func (o *RiskRisk) HasLastChangedOn() bool {
	if o != nil && o.LastChangedOn != nil {
		return true
	}

	return false
}

// SetLastChangedOn gets a reference to the given string and assigns it to the LastChangedOn field.
func (o *RiskRisk) SetLastChangedOn(v string) {
	o.LastChangedOn = &v
}

// GetMitigationPlan returns the MitigationPlan field value if set, zero value otherwise.
func (o *RiskRisk) GetMitigationPlan() string {
	if o == nil || o.MitigationPlan == nil {
		var ret string
		return ret
	}
	return *o.MitigationPlan
}

// GetMitigationPlanOk returns a tuple with the MitigationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetMitigationPlanOk() (*string, bool) {
	if o == nil || o.MitigationPlan == nil {
		return nil, false
	}
	return o.MitigationPlan, true
}

// HasMitigationPlan returns a boolean if a field has been set.
func (o *RiskRisk) HasMitigationPlan() bool {
	if o != nil && o.MitigationPlan != nil {
		return true
	}

	return false
}

// SetMitigationPlan gets a reference to the given string and assigns it to the MitigationPlan field.
func (o *RiskRisk) SetMitigationPlan(v string) {
	o.MitigationPlan = &v
}

// GetProbability returns the Probability field value if set, zero value otherwise.
func (o *RiskRisk) GetProbability() string {
	if o == nil || o.Probability == nil {
		var ret string
		return ret
	}
	return *o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetProbabilityOk() (*string, bool) {
	if o == nil || o.Probability == nil {
		return nil, false
	}
	return o.Probability, true
}

// HasProbability returns a boolean if a field has been set.
func (o *RiskRisk) HasProbability() bool {
	if o != nil && o.Probability != nil {
		return true
	}

	return false
}

// SetProbability gets a reference to the given string and assigns it to the Probability field.
func (o *RiskRisk) SetProbability(v string) {
	o.Probability = &v
}

// GetProbabilityValue returns the ProbabilityValue field value if set, zero value otherwise.
func (o *RiskRisk) GetProbabilityValue() int32 {
	if o == nil || o.ProbabilityValue == nil {
		var ret int32
		return ret
	}
	return *o.ProbabilityValue
}

// GetProbabilityValueOk returns a tuple with the ProbabilityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetProbabilityValueOk() (*int32, bool) {
	if o == nil || o.ProbabilityValue == nil {
		return nil, false
	}
	return o.ProbabilityValue, true
}

// HasProbabilityValue returns a boolean if a field has been set.
func (o *RiskRisk) HasProbabilityValue() bool {
	if o != nil && o.ProbabilityValue != nil {
		return true
	}

	return false
}

// SetProbabilityValue gets a reference to the given int32 and assigns it to the ProbabilityValue field.
func (o *RiskRisk) SetProbabilityValue(v int32) {
	o.ProbabilityValue = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *RiskRisk) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *RiskRisk) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *RiskRisk) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RiskRisk) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RiskRisk) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *RiskRisk) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RiskRisk) GetResult() int32 {
	if o == nil || o.Result == nil {
		var ret int32
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetResultOk() (*int32, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RiskRisk) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given int32 and assigns it to the Result field.
func (o *RiskRisk) SetResult(v int32) {
	o.Result = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RiskRisk) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RiskRisk) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RiskRisk) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RiskRisk) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RiskRisk) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RiskRisk) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskRisk) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskRisk) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskRisk) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RiskRisk) GetUpdatedBy() int32 {
	if o == nil || o.UpdatedBy == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRisk) GetUpdatedByOk() (*int32, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RiskRisk) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *RiskRisk) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

func (o RiskRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanEdit != nil {
		toSerialize["canEdit"] = o.CanEdit
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.ImpactCost != nil {
		toSerialize["impactCost"] = o.ImpactCost
	}
	if o.ImpactPerformance != nil {
		toSerialize["impactPerformance"] = o.ImpactPerformance
	}
	if o.ImpactSchedule != nil {
		toSerialize["impactSchedule"] = o.ImpactSchedule
	}
	if o.ImpactValue != nil {
		toSerialize["impactValue"] = o.ImpactValue
	}
	if o.LastChangedByUserId != nil {
		toSerialize["lastChangedByUserId"] = o.LastChangedByUserId
	}
	if o.LastChangedOn != nil {
		toSerialize["lastChangedOn"] = o.LastChangedOn
	}
	if o.MitigationPlan != nil {
		toSerialize["mitigationPlan"] = o.MitigationPlan
	}
	if o.Probability != nil {
		toSerialize["probability"] = o.Probability
	}
	if o.ProbabilityValue != nil {
		toSerialize["probabilityValue"] = o.ProbabilityValue
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableRiskRisk struct {
	value *RiskRisk
	isSet bool
}

func (v NullableRiskRisk) Get() *RiskRisk {
	return v.value
}

func (v *NullableRiskRisk) Set(val *RiskRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRisk(val *RiskRisk) *NullableRiskRisk {
	return &NullableRiskRisk{value: val, isSet: true}
}

func (v NullableRiskRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

