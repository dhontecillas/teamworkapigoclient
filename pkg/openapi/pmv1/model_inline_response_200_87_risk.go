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

// InlineResponse20087Risk struct for InlineResponse20087Risk
type InlineResponse20087Risk struct {
	CreatedByUserFirstName *string `json:"createdByUserFirstName,omitempty"`
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	CreatedByUserLastName *string `json:"createdByUserLastName,omitempty"`
	CreatedOn *string `json:"createdOn,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Id *string `json:"id,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactCost *string `json:"impactCost,omitempty"`
	ImpactPerformance *string `json:"impactPerformance,omitempty"`
	ImpactSchedule *string `json:"impactSchedule,omitempty"`
	ImpactValue *string `json:"impactValue,omitempty"`
	LastChangedByUserFirstName *string `json:"lastChangedByUserFirstName,omitempty"`
	LastChangedByUserId *string `json:"lastChangedByUserId,omitempty"`
	LastChangedByUserLastName *string `json:"lastChangedByUserLastName,omitempty"`
	LastChangedOn *string `json:"lastChangedOn,omitempty"`
	MitigationPlan *string `json:"mitigationPlan,omitempty"`
	Probability *string `json:"probability,omitempty"`
	ProbabilityValue *string `json:"probabilityValue,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	Result *string `json:"result,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse20087Risk instantiates a new InlineResponse20087Risk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087Risk() *InlineResponse20087Risk {
	this := InlineResponse20087Risk{}
	return &this
}

// NewInlineResponse20087RiskWithDefaults instantiates a new InlineResponse20087Risk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087RiskWithDefaults() *InlineResponse20087Risk {
	this := InlineResponse20087Risk{}
	return &this
}

// GetCreatedByUserFirstName returns the CreatedByUserFirstName field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetCreatedByUserFirstName() string {
	if o == nil || o.CreatedByUserFirstName == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserFirstName
}

// GetCreatedByUserFirstNameOk returns a tuple with the CreatedByUserFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetCreatedByUserFirstNameOk() (*string, bool) {
	if o == nil || o.CreatedByUserFirstName == nil {
		return nil, false
	}
	return o.CreatedByUserFirstName, true
}

// HasCreatedByUserFirstName returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasCreatedByUserFirstName() bool {
	if o != nil && o.CreatedByUserFirstName != nil {
		return true
	}

	return false
}

// SetCreatedByUserFirstName gets a reference to the given string and assigns it to the CreatedByUserFirstName field.
func (o *InlineResponse20087Risk) SetCreatedByUserFirstName(v string) {
	o.CreatedByUserFirstName = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetCreatedByUserId() string {
	if o == nil || o.CreatedByUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *InlineResponse20087Risk) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetCreatedByUserLastName returns the CreatedByUserLastName field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetCreatedByUserLastName() string {
	if o == nil || o.CreatedByUserLastName == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUserLastName
}

// GetCreatedByUserLastNameOk returns a tuple with the CreatedByUserLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetCreatedByUserLastNameOk() (*string, bool) {
	if o == nil || o.CreatedByUserLastName == nil {
		return nil, false
	}
	return o.CreatedByUserLastName, true
}

// HasCreatedByUserLastName returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasCreatedByUserLastName() bool {
	if o != nil && o.CreatedByUserLastName != nil {
		return true
	}

	return false
}

// SetCreatedByUserLastName gets a reference to the given string and assigns it to the CreatedByUserLastName field.
func (o *InlineResponse20087Risk) SetCreatedByUserLastName(v string) {
	o.CreatedByUserLastName = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *InlineResponse20087Risk) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *InlineResponse20087Risk) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20087Risk) SetId(v string) {
	o.Id = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *InlineResponse20087Risk) SetImpact(v string) {
	o.Impact = &v
}

// GetImpactCost returns the ImpactCost field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetImpactCost() string {
	if o == nil || o.ImpactCost == nil {
		var ret string
		return ret
	}
	return *o.ImpactCost
}

// GetImpactCostOk returns a tuple with the ImpactCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetImpactCostOk() (*string, bool) {
	if o == nil || o.ImpactCost == nil {
		return nil, false
	}
	return o.ImpactCost, true
}

// HasImpactCost returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasImpactCost() bool {
	if o != nil && o.ImpactCost != nil {
		return true
	}

	return false
}

// SetImpactCost gets a reference to the given string and assigns it to the ImpactCost field.
func (o *InlineResponse20087Risk) SetImpactCost(v string) {
	o.ImpactCost = &v
}

// GetImpactPerformance returns the ImpactPerformance field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetImpactPerformance() string {
	if o == nil || o.ImpactPerformance == nil {
		var ret string
		return ret
	}
	return *o.ImpactPerformance
}

// GetImpactPerformanceOk returns a tuple with the ImpactPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetImpactPerformanceOk() (*string, bool) {
	if o == nil || o.ImpactPerformance == nil {
		return nil, false
	}
	return o.ImpactPerformance, true
}

// HasImpactPerformance returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasImpactPerformance() bool {
	if o != nil && o.ImpactPerformance != nil {
		return true
	}

	return false
}

// SetImpactPerformance gets a reference to the given string and assigns it to the ImpactPerformance field.
func (o *InlineResponse20087Risk) SetImpactPerformance(v string) {
	o.ImpactPerformance = &v
}

// GetImpactSchedule returns the ImpactSchedule field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetImpactSchedule() string {
	if o == nil || o.ImpactSchedule == nil {
		var ret string
		return ret
	}
	return *o.ImpactSchedule
}

// GetImpactScheduleOk returns a tuple with the ImpactSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetImpactScheduleOk() (*string, bool) {
	if o == nil || o.ImpactSchedule == nil {
		return nil, false
	}
	return o.ImpactSchedule, true
}

// HasImpactSchedule returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasImpactSchedule() bool {
	if o != nil && o.ImpactSchedule != nil {
		return true
	}

	return false
}

// SetImpactSchedule gets a reference to the given string and assigns it to the ImpactSchedule field.
func (o *InlineResponse20087Risk) SetImpactSchedule(v string) {
	o.ImpactSchedule = &v
}

// GetImpactValue returns the ImpactValue field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetImpactValue() string {
	if o == nil || o.ImpactValue == nil {
		var ret string
		return ret
	}
	return *o.ImpactValue
}

// GetImpactValueOk returns a tuple with the ImpactValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetImpactValueOk() (*string, bool) {
	if o == nil || o.ImpactValue == nil {
		return nil, false
	}
	return o.ImpactValue, true
}

// HasImpactValue returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasImpactValue() bool {
	if o != nil && o.ImpactValue != nil {
		return true
	}

	return false
}

// SetImpactValue gets a reference to the given string and assigns it to the ImpactValue field.
func (o *InlineResponse20087Risk) SetImpactValue(v string) {
	o.ImpactValue = &v
}

// GetLastChangedByUserFirstName returns the LastChangedByUserFirstName field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetLastChangedByUserFirstName() string {
	if o == nil || o.LastChangedByUserFirstName == nil {
		var ret string
		return ret
	}
	return *o.LastChangedByUserFirstName
}

// GetLastChangedByUserFirstNameOk returns a tuple with the LastChangedByUserFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetLastChangedByUserFirstNameOk() (*string, bool) {
	if o == nil || o.LastChangedByUserFirstName == nil {
		return nil, false
	}
	return o.LastChangedByUserFirstName, true
}

// HasLastChangedByUserFirstName returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasLastChangedByUserFirstName() bool {
	if o != nil && o.LastChangedByUserFirstName != nil {
		return true
	}

	return false
}

// SetLastChangedByUserFirstName gets a reference to the given string and assigns it to the LastChangedByUserFirstName field.
func (o *InlineResponse20087Risk) SetLastChangedByUserFirstName(v string) {
	o.LastChangedByUserFirstName = &v
}

// GetLastChangedByUserId returns the LastChangedByUserId field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetLastChangedByUserId() string {
	if o == nil || o.LastChangedByUserId == nil {
		var ret string
		return ret
	}
	return *o.LastChangedByUserId
}

// GetLastChangedByUserIdOk returns a tuple with the LastChangedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetLastChangedByUserIdOk() (*string, bool) {
	if o == nil || o.LastChangedByUserId == nil {
		return nil, false
	}
	return o.LastChangedByUserId, true
}

// HasLastChangedByUserId returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasLastChangedByUserId() bool {
	if o != nil && o.LastChangedByUserId != nil {
		return true
	}

	return false
}

// SetLastChangedByUserId gets a reference to the given string and assigns it to the LastChangedByUserId field.
func (o *InlineResponse20087Risk) SetLastChangedByUserId(v string) {
	o.LastChangedByUserId = &v
}

// GetLastChangedByUserLastName returns the LastChangedByUserLastName field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetLastChangedByUserLastName() string {
	if o == nil || o.LastChangedByUserLastName == nil {
		var ret string
		return ret
	}
	return *o.LastChangedByUserLastName
}

// GetLastChangedByUserLastNameOk returns a tuple with the LastChangedByUserLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetLastChangedByUserLastNameOk() (*string, bool) {
	if o == nil || o.LastChangedByUserLastName == nil {
		return nil, false
	}
	return o.LastChangedByUserLastName, true
}

// HasLastChangedByUserLastName returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasLastChangedByUserLastName() bool {
	if o != nil && o.LastChangedByUserLastName != nil {
		return true
	}

	return false
}

// SetLastChangedByUserLastName gets a reference to the given string and assigns it to the LastChangedByUserLastName field.
func (o *InlineResponse20087Risk) SetLastChangedByUserLastName(v string) {
	o.LastChangedByUserLastName = &v
}

// GetLastChangedOn returns the LastChangedOn field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetLastChangedOn() string {
	if o == nil || o.LastChangedOn == nil {
		var ret string
		return ret
	}
	return *o.LastChangedOn
}

// GetLastChangedOnOk returns a tuple with the LastChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetLastChangedOnOk() (*string, bool) {
	if o == nil || o.LastChangedOn == nil {
		return nil, false
	}
	return o.LastChangedOn, true
}

// HasLastChangedOn returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasLastChangedOn() bool {
	if o != nil && o.LastChangedOn != nil {
		return true
	}

	return false
}

// SetLastChangedOn gets a reference to the given string and assigns it to the LastChangedOn field.
func (o *InlineResponse20087Risk) SetLastChangedOn(v string) {
	o.LastChangedOn = &v
}

// GetMitigationPlan returns the MitigationPlan field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetMitigationPlan() string {
	if o == nil || o.MitigationPlan == nil {
		var ret string
		return ret
	}
	return *o.MitigationPlan
}

// GetMitigationPlanOk returns a tuple with the MitigationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetMitigationPlanOk() (*string, bool) {
	if o == nil || o.MitigationPlan == nil {
		return nil, false
	}
	return o.MitigationPlan, true
}

// HasMitigationPlan returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasMitigationPlan() bool {
	if o != nil && o.MitigationPlan != nil {
		return true
	}

	return false
}

// SetMitigationPlan gets a reference to the given string and assigns it to the MitigationPlan field.
func (o *InlineResponse20087Risk) SetMitigationPlan(v string) {
	o.MitigationPlan = &v
}

// GetProbability returns the Probability field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetProbability() string {
	if o == nil || o.Probability == nil {
		var ret string
		return ret
	}
	return *o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetProbabilityOk() (*string, bool) {
	if o == nil || o.Probability == nil {
		return nil, false
	}
	return o.Probability, true
}

// HasProbability returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasProbability() bool {
	if o != nil && o.Probability != nil {
		return true
	}

	return false
}

// SetProbability gets a reference to the given string and assigns it to the Probability field.
func (o *InlineResponse20087Risk) SetProbability(v string) {
	o.Probability = &v
}

// GetProbabilityValue returns the ProbabilityValue field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetProbabilityValue() string {
	if o == nil || o.ProbabilityValue == nil {
		var ret string
		return ret
	}
	return *o.ProbabilityValue
}

// GetProbabilityValueOk returns a tuple with the ProbabilityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetProbabilityValueOk() (*string, bool) {
	if o == nil || o.ProbabilityValue == nil {
		return nil, false
	}
	return o.ProbabilityValue, true
}

// HasProbabilityValue returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasProbabilityValue() bool {
	if o != nil && o.ProbabilityValue != nil {
		return true
	}

	return false
}

// SetProbabilityValue gets a reference to the given string and assigns it to the ProbabilityValue field.
func (o *InlineResponse20087Risk) SetProbabilityValue(v string) {
	o.ProbabilityValue = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse20087Risk) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InlineResponse20087Risk) SetResult(v string) {
	o.Result = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InlineResponse20087Risk) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20087Risk) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Risk) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20087Risk) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20087Risk) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse20087Risk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedByUserFirstName != nil {
		toSerialize["createdByUserFirstName"] = o.CreatedByUserFirstName
	}
	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if o.CreatedByUserLastName != nil {
		toSerialize["createdByUserLastName"] = o.CreatedByUserLastName
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
	if o.LastChangedByUserFirstName != nil {
		toSerialize["lastChangedByUserFirstName"] = o.LastChangedByUserFirstName
	}
	if o.LastChangedByUserId != nil {
		toSerialize["lastChangedByUserId"] = o.LastChangedByUserId
	}
	if o.LastChangedByUserLastName != nil {
		toSerialize["lastChangedByUserLastName"] = o.LastChangedByUserLastName
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20087Risk struct {
	value *InlineResponse20087Risk
	isSet bool
}

func (v NullableInlineResponse20087Risk) Get() *InlineResponse20087Risk {
	return v.value
}

func (v *NullableInlineResponse20087Risk) Set(val *InlineResponse20087Risk) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087Risk) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087Risk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087Risk(val *InlineResponse20087Risk) *NullableInlineResponse20087Risk {
	return &NullableInlineResponse20087Risk{value: val, isSet: true}
}

func (v NullableInlineResponse20087Risk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087Risk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


