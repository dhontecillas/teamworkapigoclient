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

// ProjectsIdRisksJsonRisk struct for ProjectsIdRisksJsonRisk
type ProjectsIdRisksJsonRisk struct {
	ImpactCost *string `json:"impactCost,omitempty"`
	ImpactPerformance *string `json:"impactPerformance,omitempty"`
	ImpactSchedule *string `json:"impactSchedule,omitempty"`
	ImpactValue *string `json:"impactValue,omitempty"`
	MitigationPlan *string `json:"mitigationPlan,omitempty"`
	ProbabilityValue *string `json:"probabilityValue,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewProjectsIdRisksJsonRisk instantiates a new ProjectsIdRisksJsonRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdRisksJsonRisk() *ProjectsIdRisksJsonRisk {
	this := ProjectsIdRisksJsonRisk{}
	return &this
}

// NewProjectsIdRisksJsonRiskWithDefaults instantiates a new ProjectsIdRisksJsonRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdRisksJsonRiskWithDefaults() *ProjectsIdRisksJsonRisk {
	this := ProjectsIdRisksJsonRisk{}
	return &this
}

// GetImpactCost returns the ImpactCost field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetImpactCost() string {
	if o == nil || o.ImpactCost == nil {
		var ret string
		return ret
	}
	return *o.ImpactCost
}

// GetImpactCostOk returns a tuple with the ImpactCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetImpactCostOk() (*string, bool) {
	if o == nil || o.ImpactCost == nil {
		return nil, false
	}
	return o.ImpactCost, true
}

// HasImpactCost returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasImpactCost() bool {
	if o != nil && o.ImpactCost != nil {
		return true
	}

	return false
}

// SetImpactCost gets a reference to the given string and assigns it to the ImpactCost field.
func (o *ProjectsIdRisksJsonRisk) SetImpactCost(v string) {
	o.ImpactCost = &v
}

// GetImpactPerformance returns the ImpactPerformance field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetImpactPerformance() string {
	if o == nil || o.ImpactPerformance == nil {
		var ret string
		return ret
	}
	return *o.ImpactPerformance
}

// GetImpactPerformanceOk returns a tuple with the ImpactPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetImpactPerformanceOk() (*string, bool) {
	if o == nil || o.ImpactPerformance == nil {
		return nil, false
	}
	return o.ImpactPerformance, true
}

// HasImpactPerformance returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasImpactPerformance() bool {
	if o != nil && o.ImpactPerformance != nil {
		return true
	}

	return false
}

// SetImpactPerformance gets a reference to the given string and assigns it to the ImpactPerformance field.
func (o *ProjectsIdRisksJsonRisk) SetImpactPerformance(v string) {
	o.ImpactPerformance = &v
}

// GetImpactSchedule returns the ImpactSchedule field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetImpactSchedule() string {
	if o == nil || o.ImpactSchedule == nil {
		var ret string
		return ret
	}
	return *o.ImpactSchedule
}

// GetImpactScheduleOk returns a tuple with the ImpactSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetImpactScheduleOk() (*string, bool) {
	if o == nil || o.ImpactSchedule == nil {
		return nil, false
	}
	return o.ImpactSchedule, true
}

// HasImpactSchedule returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasImpactSchedule() bool {
	if o != nil && o.ImpactSchedule != nil {
		return true
	}

	return false
}

// SetImpactSchedule gets a reference to the given string and assigns it to the ImpactSchedule field.
func (o *ProjectsIdRisksJsonRisk) SetImpactSchedule(v string) {
	o.ImpactSchedule = &v
}

// GetImpactValue returns the ImpactValue field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetImpactValue() string {
	if o == nil || o.ImpactValue == nil {
		var ret string
		return ret
	}
	return *o.ImpactValue
}

// GetImpactValueOk returns a tuple with the ImpactValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetImpactValueOk() (*string, bool) {
	if o == nil || o.ImpactValue == nil {
		return nil, false
	}
	return o.ImpactValue, true
}

// HasImpactValue returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasImpactValue() bool {
	if o != nil && o.ImpactValue != nil {
		return true
	}

	return false
}

// SetImpactValue gets a reference to the given string and assigns it to the ImpactValue field.
func (o *ProjectsIdRisksJsonRisk) SetImpactValue(v string) {
	o.ImpactValue = &v
}

// GetMitigationPlan returns the MitigationPlan field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetMitigationPlan() string {
	if o == nil || o.MitigationPlan == nil {
		var ret string
		return ret
	}
	return *o.MitigationPlan
}

// GetMitigationPlanOk returns a tuple with the MitigationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetMitigationPlanOk() (*string, bool) {
	if o == nil || o.MitigationPlan == nil {
		return nil, false
	}
	return o.MitigationPlan, true
}

// HasMitigationPlan returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasMitigationPlan() bool {
	if o != nil && o.MitigationPlan != nil {
		return true
	}

	return false
}

// SetMitigationPlan gets a reference to the given string and assigns it to the MitigationPlan field.
func (o *ProjectsIdRisksJsonRisk) SetMitigationPlan(v string) {
	o.MitigationPlan = &v
}

// GetProbabilityValue returns the ProbabilityValue field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetProbabilityValue() string {
	if o == nil || o.ProbabilityValue == nil {
		var ret string
		return ret
	}
	return *o.ProbabilityValue
}

// GetProbabilityValueOk returns a tuple with the ProbabilityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetProbabilityValueOk() (*string, bool) {
	if o == nil || o.ProbabilityValue == nil {
		return nil, false
	}
	return o.ProbabilityValue, true
}

// HasProbabilityValue returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasProbabilityValue() bool {
	if o != nil && o.ProbabilityValue != nil {
		return true
	}

	return false
}

// SetProbabilityValue gets a reference to the given string and assigns it to the ProbabilityValue field.
func (o *ProjectsIdRisksJsonRisk) SetProbabilityValue(v string) {
	o.ProbabilityValue = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProjectsIdRisksJsonRisk) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectsIdRisksJsonRisk) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdRisksJsonRisk) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectsIdRisksJsonRisk) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProjectsIdRisksJsonRisk) SetStatus(v string) {
	o.Status = &v
}

func (o ProjectsIdRisksJsonRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.MitigationPlan != nil {
		toSerialize["mitigationPlan"] = o.MitigationPlan
	}
	if o.ProbabilityValue != nil {
		toSerialize["probabilityValue"] = o.ProbabilityValue
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdRisksJsonRisk struct {
	value *ProjectsIdRisksJsonRisk
	isSet bool
}

func (v NullableProjectsIdRisksJsonRisk) Get() *ProjectsIdRisksJsonRisk {
	return v.value
}

func (v *NullableProjectsIdRisksJsonRisk) Set(val *ProjectsIdRisksJsonRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdRisksJsonRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdRisksJsonRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdRisksJsonRisk(val *ProjectsIdRisksJsonRisk) *NullableProjectsIdRisksJsonRisk {
	return &NullableProjectsIdRisksJsonRisk{value: val, isSet: true}
}

func (v NullableProjectsIdRisksJsonRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdRisksJsonRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


