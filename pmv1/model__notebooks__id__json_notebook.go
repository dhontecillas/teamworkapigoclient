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

// NotebooksIdJsonNotebook struct for NotebooksIdJsonNotebook
type NotebooksIdJsonNotebook struct {
	CategoryId *string `json:"category-id,omitempty"`
	CategoryName *string `json:"category-name,omitempty"`
	Content *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Notify *string `json:"notify,omitempty"`
	Private *bool `json:"private,omitempty"`
	Tags *string `json:"tags,omitempty"`
}

// NewNotebooksIdJsonNotebook instantiates a new NotebooksIdJsonNotebook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebooksIdJsonNotebook() *NotebooksIdJsonNotebook {
	this := NotebooksIdJsonNotebook{}
	return &this
}

// NewNotebooksIdJsonNotebookWithDefaults instantiates a new NotebooksIdJsonNotebook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebooksIdJsonNotebookWithDefaults() *NotebooksIdJsonNotebook {
	this := NotebooksIdJsonNotebook{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *NotebooksIdJsonNotebook) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *NotebooksIdJsonNotebook) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *NotebooksIdJsonNotebook) SetContent(v string) {
	o.Content = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotebooksIdJsonNotebook) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotebooksIdJsonNotebook) SetName(v string) {
	o.Name = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetNotify() string {
	if o == nil || o.Notify == nil {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetNotifyOk() (*string, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *NotebooksIdJsonNotebook) SetNotify(v string) {
	o.Notify = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *NotebooksIdJsonNotebook) SetPrivate(v bool) {
	o.Private = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NotebooksIdJsonNotebook) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebooksIdJsonNotebook) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NotebooksIdJsonNotebook) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *NotebooksIdJsonNotebook) SetTags(v string) {
	o.Tags = &v
}

func (o NotebooksIdJsonNotebook) MarshalJSON() ([]byte, error) {
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

type NullableNotebooksIdJsonNotebook struct {
	value *NotebooksIdJsonNotebook
	isSet bool
}

func (v NullableNotebooksIdJsonNotebook) Get() *NotebooksIdJsonNotebook {
	return v.value
}

func (v *NullableNotebooksIdJsonNotebook) Set(val *NotebooksIdJsonNotebook) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebooksIdJsonNotebook) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebooksIdJsonNotebook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebooksIdJsonNotebook(val *NotebooksIdJsonNotebook) *NullableNotebooksIdJsonNotebook {
	return &NullableNotebooksIdJsonNotebook{value: val, isSet: true}
}

func (v NullableNotebooksIdJsonNotebook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebooksIdJsonNotebook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


