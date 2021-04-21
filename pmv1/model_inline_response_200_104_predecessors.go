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

// InlineResponse200104Predecessors struct for InlineResponse200104Predecessors
type InlineResponse200104Predecessors struct {
	DependentCant *string `json:"dependentCant,omitempty"`
	Hardness *string `json:"hardness,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PredecessorMust *string `json:"predecessorMust,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	ProjectType *string `json:"projectType,omitempty"`
	ResponsiblePartyId *string `json:"responsible-party-id,omitempty"`
	ResponsiblePartyNames *string `json:"responsible-party-names,omitempty"`
	ResponsiblePartySummary *string `json:"responsible-party-summary,omitempty"`
}

// NewInlineResponse200104Predecessors instantiates a new InlineResponse200104Predecessors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200104Predecessors() *InlineResponse200104Predecessors {
	this := InlineResponse200104Predecessors{}
	return &this
}

// NewInlineResponse200104PredecessorsWithDefaults instantiates a new InlineResponse200104Predecessors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200104PredecessorsWithDefaults() *InlineResponse200104Predecessors {
	this := InlineResponse200104Predecessors{}
	return &this
}

// GetDependentCant returns the DependentCant field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetDependentCant() string {
	if o == nil || o.DependentCant == nil {
		var ret string
		return ret
	}
	return *o.DependentCant
}

// GetDependentCantOk returns a tuple with the DependentCant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetDependentCantOk() (*string, bool) {
	if o == nil || o.DependentCant == nil {
		return nil, false
	}
	return o.DependentCant, true
}

// HasDependentCant returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasDependentCant() bool {
	if o != nil && o.DependentCant != nil {
		return true
	}

	return false
}

// SetDependentCant gets a reference to the given string and assigns it to the DependentCant field.
func (o *InlineResponse200104Predecessors) SetDependentCant(v string) {
	o.DependentCant = &v
}

// GetHardness returns the Hardness field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetHardness() string {
	if o == nil || o.Hardness == nil {
		var ret string
		return ret
	}
	return *o.Hardness
}

// GetHardnessOk returns a tuple with the Hardness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetHardnessOk() (*string, bool) {
	if o == nil || o.Hardness == nil {
		return nil, false
	}
	return o.Hardness, true
}

// HasHardness returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasHardness() bool {
	if o != nil && o.Hardness != nil {
		return true
	}

	return false
}

// SetHardness gets a reference to the given string and assigns it to the Hardness field.
func (o *InlineResponse200104Predecessors) SetHardness(v string) {
	o.Hardness = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200104Predecessors) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200104Predecessors) SetName(v string) {
	o.Name = &v
}

// GetPredecessorMust returns the PredecessorMust field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetPredecessorMust() string {
	if o == nil || o.PredecessorMust == nil {
		var ret string
		return ret
	}
	return *o.PredecessorMust
}

// GetPredecessorMustOk returns a tuple with the PredecessorMust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetPredecessorMustOk() (*string, bool) {
	if o == nil || o.PredecessorMust == nil {
		return nil, false
	}
	return o.PredecessorMust, true
}

// HasPredecessorMust returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasPredecessorMust() bool {
	if o != nil && o.PredecessorMust != nil {
		return true
	}

	return false
}

// SetPredecessorMust gets a reference to the given string and assigns it to the PredecessorMust field.
func (o *InlineResponse200104Predecessors) SetPredecessorMust(v string) {
	o.PredecessorMust = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse200104Predecessors) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetProjectType() string {
	if o == nil || o.ProjectType == nil {
		var ret string
		return ret
	}
	return *o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetProjectTypeOk() (*string, bool) {
	if o == nil || o.ProjectType == nil {
		return nil, false
	}
	return o.ProjectType, true
}

// HasProjectType returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasProjectType() bool {
	if o != nil && o.ProjectType != nil {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given string and assigns it to the ProjectType field.
func (o *InlineResponse200104Predecessors) SetProjectType(v string) {
	o.ProjectType = &v
}

// GetResponsiblePartyId returns the ResponsiblePartyId field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetResponsiblePartyId() string {
	if o == nil || o.ResponsiblePartyId == nil {
		var ret string
		return ret
	}
	return *o.ResponsiblePartyId
}

// GetResponsiblePartyIdOk returns a tuple with the ResponsiblePartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetResponsiblePartyIdOk() (*string, bool) {
	if o == nil || o.ResponsiblePartyId == nil {
		return nil, false
	}
	return o.ResponsiblePartyId, true
}

// HasResponsiblePartyId returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasResponsiblePartyId() bool {
	if o != nil && o.ResponsiblePartyId != nil {
		return true
	}

	return false
}

// SetResponsiblePartyId gets a reference to the given string and assigns it to the ResponsiblePartyId field.
func (o *InlineResponse200104Predecessors) SetResponsiblePartyId(v string) {
	o.ResponsiblePartyId = &v
}

// GetResponsiblePartyNames returns the ResponsiblePartyNames field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetResponsiblePartyNames() string {
	if o == nil || o.ResponsiblePartyNames == nil {
		var ret string
		return ret
	}
	return *o.ResponsiblePartyNames
}

// GetResponsiblePartyNamesOk returns a tuple with the ResponsiblePartyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetResponsiblePartyNamesOk() (*string, bool) {
	if o == nil || o.ResponsiblePartyNames == nil {
		return nil, false
	}
	return o.ResponsiblePartyNames, true
}

// HasResponsiblePartyNames returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasResponsiblePartyNames() bool {
	if o != nil && o.ResponsiblePartyNames != nil {
		return true
	}

	return false
}

// SetResponsiblePartyNames gets a reference to the given string and assigns it to the ResponsiblePartyNames field.
func (o *InlineResponse200104Predecessors) SetResponsiblePartyNames(v string) {
	o.ResponsiblePartyNames = &v
}

// GetResponsiblePartySummary returns the ResponsiblePartySummary field value if set, zero value otherwise.
func (o *InlineResponse200104Predecessors) GetResponsiblePartySummary() string {
	if o == nil || o.ResponsiblePartySummary == nil {
		var ret string
		return ret
	}
	return *o.ResponsiblePartySummary
}

// GetResponsiblePartySummaryOk returns a tuple with the ResponsiblePartySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104Predecessors) GetResponsiblePartySummaryOk() (*string, bool) {
	if o == nil || o.ResponsiblePartySummary == nil {
		return nil, false
	}
	return o.ResponsiblePartySummary, true
}

// HasResponsiblePartySummary returns a boolean if a field has been set.
func (o *InlineResponse200104Predecessors) HasResponsiblePartySummary() bool {
	if o != nil && o.ResponsiblePartySummary != nil {
		return true
	}

	return false
}

// SetResponsiblePartySummary gets a reference to the given string and assigns it to the ResponsiblePartySummary field.
func (o *InlineResponse200104Predecessors) SetResponsiblePartySummary(v string) {
	o.ResponsiblePartySummary = &v
}

func (o InlineResponse200104Predecessors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DependentCant != nil {
		toSerialize["dependentCant"] = o.DependentCant
	}
	if o.Hardness != nil {
		toSerialize["hardness"] = o.Hardness
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PredecessorMust != nil {
		toSerialize["predecessorMust"] = o.PredecessorMust
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectType != nil {
		toSerialize["projectType"] = o.ProjectType
	}
	if o.ResponsiblePartyId != nil {
		toSerialize["responsible-party-id"] = o.ResponsiblePartyId
	}
	if o.ResponsiblePartyNames != nil {
		toSerialize["responsible-party-names"] = o.ResponsiblePartyNames
	}
	if o.ResponsiblePartySummary != nil {
		toSerialize["responsible-party-summary"] = o.ResponsiblePartySummary
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200104Predecessors struct {
	value *InlineResponse200104Predecessors
	isSet bool
}

func (v NullableInlineResponse200104Predecessors) Get() *InlineResponse200104Predecessors {
	return v.value
}

func (v *NullableInlineResponse200104Predecessors) Set(val *InlineResponse200104Predecessors) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200104Predecessors) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200104Predecessors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200104Predecessors(val *InlineResponse200104Predecessors) *NullableInlineResponse200104Predecessors {
	return &NullableInlineResponse200104Predecessors{value: val, isSet: true}
}

func (v NullableInlineResponse200104Predecessors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200104Predecessors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

