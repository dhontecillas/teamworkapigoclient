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

// ViewLinkItem LinkItem contains all the information returned from a link item.
type ViewLinkItem struct {
	Category *ViewRelationship `json:"category,omitempty"`
	CategoryId *int32 `json:"categoryId,omitempty"`
	Code *string `json:"code,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *int32 `json:"createdBy,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	DeletedBy *int32 `json:"deletedBy,omitempty"`
	Description *string `json:"description,omitempty"`
	ForceNewWindow *int32 `json:"forceNewWindow,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsPrivate *bool `json:"isPrivate,omitempty"`
	Project *ViewRelationship `json:"project,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *int32 `json:"updatedBy,omitempty"`
	Width *int32 `json:"width,omitempty"`
}

// NewViewLinkItem instantiates a new ViewLinkItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewLinkItem() *ViewLinkItem {
	this := ViewLinkItem{}
	return &this
}

// NewViewLinkItemWithDefaults instantiates a new ViewLinkItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewLinkItemWithDefaults() *ViewLinkItem {
	this := ViewLinkItem{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ViewLinkItem) GetCategory() ViewRelationship {
	if o == nil || o.Category == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetCategoryOk() (*ViewRelationship, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ViewLinkItem) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ViewRelationship and assigns it to the Category field.
func (o *ViewLinkItem) SetCategory(v ViewRelationship) {
	o.Category = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ViewLinkItem) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ViewLinkItem) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *ViewLinkItem) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ViewLinkItem) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ViewLinkItem) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ViewLinkItem) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ViewLinkItem) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ViewLinkItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ViewLinkItem) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ViewLinkItem) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetCreatedByOk() (*int32, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ViewLinkItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *ViewLinkItem) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ViewLinkItem) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ViewLinkItem) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *ViewLinkItem) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ViewLinkItem) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ViewLinkItem) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ViewLinkItem) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *ViewLinkItem) GetDeletedBy() int32 {
	if o == nil || o.DeletedBy == nil {
		var ret int32
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetDeletedByOk() (*int32, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *ViewLinkItem) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given int32 and assigns it to the DeletedBy field.
func (o *ViewLinkItem) SetDeletedBy(v int32) {
	o.DeletedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ViewLinkItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ViewLinkItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ViewLinkItem) SetDescription(v string) {
	o.Description = &v
}

// GetForceNewWindow returns the ForceNewWindow field value if set, zero value otherwise.
func (o *ViewLinkItem) GetForceNewWindow() int32 {
	if o == nil || o.ForceNewWindow == nil {
		var ret int32
		return ret
	}
	return *o.ForceNewWindow
}

// GetForceNewWindowOk returns a tuple with the ForceNewWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetForceNewWindowOk() (*int32, bool) {
	if o == nil || o.ForceNewWindow == nil {
		return nil, false
	}
	return o.ForceNewWindow, true
}

// HasForceNewWindow returns a boolean if a field has been set.
func (o *ViewLinkItem) HasForceNewWindow() bool {
	if o != nil && o.ForceNewWindow != nil {
		return true
	}

	return false
}

// SetForceNewWindow gets a reference to the given int32 and assigns it to the ForceNewWindow field.
func (o *ViewLinkItem) SetForceNewWindow(v int32) {
	o.ForceNewWindow = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ViewLinkItem) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ViewLinkItem) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ViewLinkItem) SetHeight(v int32) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewLinkItem) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewLinkItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewLinkItem) SetId(v int32) {
	o.Id = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *ViewLinkItem) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *ViewLinkItem) HasIsPrivate() bool {
	if o != nil && o.IsPrivate != nil {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *ViewLinkItem) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ViewLinkItem) GetProject() ViewRelationship {
	if o == nil || o.Project == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetProjectOk() (*ViewRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ViewLinkItem) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ViewRelationship and assigns it to the Project field.
func (o *ViewLinkItem) SetProject(v ViewRelationship) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ViewLinkItem) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ViewLinkItem) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ViewLinkItem) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ViewLinkItem) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ViewLinkItem) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ViewLinkItem) SetProvider(v string) {
	o.Provider = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewLinkItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewLinkItem) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewLinkItem) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ViewLinkItem) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ViewLinkItem) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ViewLinkItem) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ViewLinkItem) GetUpdatedBy() int32 {
	if o == nil || o.UpdatedBy == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetUpdatedByOk() (*int32, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ViewLinkItem) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int32 and assigns it to the UpdatedBy field.
func (o *ViewLinkItem) SetUpdatedBy(v int32) {
	o.UpdatedBy = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ViewLinkItem) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLinkItem) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ViewLinkItem) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ViewLinkItem) SetWidth(v int32) {
	o.Width = &v
}

func (o ViewLinkItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
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
	if o.ForceNewWindow != nil {
		toSerialize["forceNewWindow"] = o.ForceNewWindow
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsPrivate != nil {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
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
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

type NullableViewLinkItem struct {
	value *ViewLinkItem
	isSet bool
}

func (v NullableViewLinkItem) Get() *ViewLinkItem {
	return v.value
}

func (v *NullableViewLinkItem) Set(val *ViewLinkItem) {
	v.value = val
	v.isSet = true
}

func (v NullableViewLinkItem) IsSet() bool {
	return v.isSet
}

func (v *NullableViewLinkItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewLinkItem(val *ViewLinkItem) *NullableViewLinkItem {
	return &NullableViewLinkItem{value: val, isSet: true}
}

func (v NullableViewLinkItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewLinkItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


