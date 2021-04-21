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

// ViewNotebookCategory NotebookCategory contains all the information returned from a notebook category.
type ViewNotebookCategory struct {
	Color *string `json:"color,omitempty"`
	ElementsCount *int32 `json:"elementsCount,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Parent *ViewRelationship `json:"parent,omitempty"`
	ParentId *int32 `json:"parentId,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewViewNotebookCategory instantiates a new ViewNotebookCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewNotebookCategory() *ViewNotebookCategory {
	this := ViewNotebookCategory{}
	return &this
}

// NewViewNotebookCategoryWithDefaults instantiates a new ViewNotebookCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewNotebookCategoryWithDefaults() *ViewNotebookCategory {
	this := ViewNotebookCategory{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ViewNotebookCategory) SetColor(v string) {
	o.Color = &v
}

// GetElementsCount returns the ElementsCount field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetElementsCount() int32 {
	if o == nil || o.ElementsCount == nil {
		var ret int32
		return ret
	}
	return *o.ElementsCount
}

// GetElementsCountOk returns a tuple with the ElementsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetElementsCountOk() (*int32, bool) {
	if o == nil || o.ElementsCount == nil {
		return nil, false
	}
	return o.ElementsCount, true
}

// HasElementsCount returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasElementsCount() bool {
	if o != nil && o.ElementsCount != nil {
		return true
	}

	return false
}

// SetElementsCount gets a reference to the given int32 and assigns it to the ElementsCount field.
func (o *ViewNotebookCategory) SetElementsCount(v int32) {
	o.ElementsCount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewNotebookCategory) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ViewNotebookCategory) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetParent() ViewRelationship {
	if o == nil || o.Parent == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetParentOk() (*ViewRelationship, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given ViewRelationship and assigns it to the Parent field.
func (o *ViewNotebookCategory) SetParent(v ViewRelationship) {
	o.Parent = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetParentId() int32 {
	if o == nil || o.ParentId == nil {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetParentIdOk() (*int32, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *ViewNotebookCategory) SetParentId(v int32) {
	o.ParentId = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewNotebookCategory) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewNotebookCategory) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewNotebookCategory) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewNotebookCategory) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewNotebookCategory) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o ViewNotebookCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.ElementsCount != nil {
		toSerialize["elementsCount"] = o.ElementsCount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableViewNotebookCategory struct {
	value *ViewNotebookCategory
	isSet bool
}

func (v NullableViewNotebookCategory) Get() *ViewNotebookCategory {
	return v.value
}

func (v *NullableViewNotebookCategory) Set(val *ViewNotebookCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableViewNotebookCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableViewNotebookCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewNotebookCategory(val *ViewNotebookCategory) *NullableViewNotebookCategory {
	return &NullableViewNotebookCategory{value: val, isSet: true}
}

func (v NullableViewNotebookCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewNotebookCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


