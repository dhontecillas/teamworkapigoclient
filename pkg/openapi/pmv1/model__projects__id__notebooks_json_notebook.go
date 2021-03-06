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

// ProjectsIdNotebooksJsonNotebook struct for ProjectsIdNotebooksJsonNotebook
type ProjectsIdNotebooksJsonNotebook struct {
	CategoryId *string `json:"category-id,omitempty"`
	CategoryName *string `json:"category-name,omitempty"`
	Content *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	// \"All\" - means notify all project users
	Notify *string `json:"notify,omitempty"`
	Private *bool `json:"private,omitempty"`
	Tags *string `json:"tags,omitempty"`
}

// NewProjectsIdNotebooksJsonNotebook instantiates a new ProjectsIdNotebooksJsonNotebook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsIdNotebooksJsonNotebook() *ProjectsIdNotebooksJsonNotebook {
	this := ProjectsIdNotebooksJsonNotebook{}
	return &this
}

// NewProjectsIdNotebooksJsonNotebookWithDefaults instantiates a new ProjectsIdNotebooksJsonNotebook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsIdNotebooksJsonNotebookWithDefaults() *ProjectsIdNotebooksJsonNotebook {
	this := ProjectsIdNotebooksJsonNotebook{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *ProjectsIdNotebooksJsonNotebook) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ProjectsIdNotebooksJsonNotebook) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ProjectsIdNotebooksJsonNotebook) SetContent(v string) {
	o.Content = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsIdNotebooksJsonNotebook) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectsIdNotebooksJsonNotebook) SetName(v string) {
	o.Name = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetNotify() string {
	if o == nil || o.Notify == nil {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetNotifyOk() (*string, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *ProjectsIdNotebooksJsonNotebook) SetNotify(v string) {
	o.Notify = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ProjectsIdNotebooksJsonNotebook) SetPrivate(v bool) {
	o.Private = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectsIdNotebooksJsonNotebook) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsIdNotebooksJsonNotebook) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectsIdNotebooksJsonNotebook) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProjectsIdNotebooksJsonNotebook) SetTags(v string) {
	o.Tags = &v
}

func (o ProjectsIdNotebooksJsonNotebook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["category-id"] = o.CategoryId
	}
	if o.CategoryName != nil {
		toSerialize["category-name"] = o.CategoryName
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsIdNotebooksJsonNotebook struct {
	value *ProjectsIdNotebooksJsonNotebook
	isSet bool
}

func (v NullableProjectsIdNotebooksJsonNotebook) Get() *ProjectsIdNotebooksJsonNotebook {
	return v.value
}

func (v *NullableProjectsIdNotebooksJsonNotebook) Set(val *ProjectsIdNotebooksJsonNotebook) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsIdNotebooksJsonNotebook) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsIdNotebooksJsonNotebook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsIdNotebooksJsonNotebook(val *ProjectsIdNotebooksJsonNotebook) *NullableProjectsIdNotebooksJsonNotebook {
	return &NullableProjectsIdNotebooksJsonNotebook{value: val, isSet: true}
}

func (v NullableProjectsIdNotebooksJsonNotebook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsIdNotebooksJsonNotebook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


