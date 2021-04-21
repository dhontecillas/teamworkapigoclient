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

// ViewPortfolioCard PortfolioCard contains all the information returned from a portfolio column.
type ViewPortfolioCard struct {
	Column *ViewRelationship `json:"column,omitempty"`
	ColumnId *int32 `json:"columnId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewViewPortfolioCard instantiates a new ViewPortfolioCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewPortfolioCard() *ViewPortfolioCard {
	this := ViewPortfolioCard{}
	return &this
}

// NewViewPortfolioCardWithDefaults instantiates a new ViewPortfolioCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewPortfolioCardWithDefaults() *ViewPortfolioCard {
	this := ViewPortfolioCard{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *ViewPortfolioCard) GetColumn() ViewRelationship {
	if o == nil || o.Column == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioCard) GetColumnOk() (*ViewRelationship, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *ViewPortfolioCard) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given ViewRelationship and assigns it to the Column field.
func (o *ViewPortfolioCard) SetColumn(v ViewRelationship) {
	o.Column = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *ViewPortfolioCard) GetColumnId() int32 {
	if o == nil || o.ColumnId == nil {
		var ret int32
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioCard) GetColumnIdOk() (*int32, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *ViewPortfolioCard) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given int32 and assigns it to the ColumnId field.
func (o *ViewPortfolioCard) SetColumnId(v int32) {
	o.ColumnId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewPortfolioCard) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioCard) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewPortfolioCard) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewPortfolioCard) SetId(v int32) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewPortfolioCard) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioCard) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewPortfolioCard) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewPortfolioCard) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewPortfolioCard) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewPortfolioCard) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewPortfolioCard) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewPortfolioCard) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o ViewPortfolioCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableViewPortfolioCard struct {
	value *ViewPortfolioCard
	isSet bool
}

func (v NullableViewPortfolioCard) Get() *ViewPortfolioCard {
	return v.value
}

func (v *NullableViewPortfolioCard) Set(val *ViewPortfolioCard) {
	v.value = val
	v.isSet = true
}

func (v NullableViewPortfolioCard) IsSet() bool {
	return v.isSet
}

func (v *NullableViewPortfolioCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewPortfolioCard(val *ViewPortfolioCard) *NullableViewPortfolioCard {
	return &NullableViewPortfolioCard{value: val, isSet: true}
}

func (v NullableViewPortfolioCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewPortfolioCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


