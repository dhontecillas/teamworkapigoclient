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

// FileProjectFile ProjectFile contains all the information returned from a file.
type FileProjectFile struct {
	CategoryId *int32 `json:"categoryId,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Notify *PayloadNotify `json:"notify,omitempty"`
	NotifyCurrentUser *bool `json:"notifyCurrentUser,omitempty"`
	Privacy *PayloadUserGroups `json:"privacy,omitempty"`
	Private *bool `json:"private,omitempty"`
	TagIds *PayloadNullableInt64Slice `json:"tagIds,omitempty"`
	VersionId *int32 `json:"versionId,omitempty"`
}

// NewFileProjectFile instantiates a new FileProjectFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileProjectFile() *FileProjectFile {
	this := FileProjectFile{}
	return &this
}

// NewFileProjectFileWithDefaults instantiates a new FileProjectFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileProjectFileWithDefaults() *FileProjectFile {
	this := FileProjectFile{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *FileProjectFile) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *FileProjectFile) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *FileProjectFile) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileProjectFile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileProjectFile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileProjectFile) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileProjectFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileProjectFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileProjectFile) SetName(v string) {
	o.Name = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *FileProjectFile) GetNotify() PayloadNotify {
	if o == nil || o.Notify == nil {
		var ret PayloadNotify
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetNotifyOk() (*PayloadNotify, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *FileProjectFile) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given PayloadNotify and assigns it to the Notify field.
func (o *FileProjectFile) SetNotify(v PayloadNotify) {
	o.Notify = &v
}

// GetNotifyCurrentUser returns the NotifyCurrentUser field value if set, zero value otherwise.
func (o *FileProjectFile) GetNotifyCurrentUser() bool {
	if o == nil || o.NotifyCurrentUser == nil {
		var ret bool
		return ret
	}
	return *o.NotifyCurrentUser
}

// GetNotifyCurrentUserOk returns a tuple with the NotifyCurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetNotifyCurrentUserOk() (*bool, bool) {
	if o == nil || o.NotifyCurrentUser == nil {
		return nil, false
	}
	return o.NotifyCurrentUser, true
}

// HasNotifyCurrentUser returns a boolean if a field has been set.
func (o *FileProjectFile) HasNotifyCurrentUser() bool {
	if o != nil && o.NotifyCurrentUser != nil {
		return true
	}

	return false
}

// SetNotifyCurrentUser gets a reference to the given bool and assigns it to the NotifyCurrentUser field.
func (o *FileProjectFile) SetNotifyCurrentUser(v bool) {
	o.NotifyCurrentUser = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *FileProjectFile) GetPrivacy() PayloadUserGroups {
	if o == nil || o.Privacy == nil {
		var ret PayloadUserGroups
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetPrivacyOk() (*PayloadUserGroups, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *FileProjectFile) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given PayloadUserGroups and assigns it to the Privacy field.
func (o *FileProjectFile) SetPrivacy(v PayloadUserGroups) {
	o.Privacy = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FileProjectFile) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FileProjectFile) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *FileProjectFile) SetPrivate(v bool) {
	o.Private = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *FileProjectFile) GetTagIds() PayloadNullableInt64Slice {
	if o == nil || o.TagIds == nil {
		var ret PayloadNullableInt64Slice
		return ret
	}
	return *o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetTagIdsOk() (*PayloadNullableInt64Slice, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *FileProjectFile) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given PayloadNullableInt64Slice and assigns it to the TagIds field.
func (o *FileProjectFile) SetTagIds(v PayloadNullableInt64Slice) {
	o.TagIds = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *FileProjectFile) GetVersionId() int32 {
	if o == nil || o.VersionId == nil {
		var ret int32
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFile) GetVersionIdOk() (*int32, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *FileProjectFile) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given int32 and assigns it to the VersionId field.
func (o *FileProjectFile) SetVersionId(v int32) {
	o.VersionId = &v
}

func (o FileProjectFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
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
	if o.NotifyCurrentUser != nil {
		toSerialize["notifyCurrentUser"] = o.NotifyCurrentUser
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.TagIds != nil {
		toSerialize["tagIds"] = o.TagIds
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableFileProjectFile struct {
	value *FileProjectFile
	isSet bool
}

func (v NullableFileProjectFile) Get() *FileProjectFile {
	return v.value
}

func (v *NullableFileProjectFile) Set(val *FileProjectFile) {
	v.value = val
	v.isSet = true
}

func (v NullableFileProjectFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFileProjectFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileProjectFile(val *FileProjectFile) *NullableFileProjectFile {
	return &NullableFileProjectFile{value: val, isSet: true}
}

func (v NullableFileProjectFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileProjectFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


