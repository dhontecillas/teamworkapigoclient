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

// FileversionFileversion Fileversion contains all the information returned from a fileversion.
type FileversionFileversion struct {
	CategoryId *int32 `json:"categoryId,omitempty"`
	Description *string `json:"description,omitempty"`
	FileData *string `json:"fileData,omitempty"`
	FileRef *string `json:"fileRef,omitempty"`
	FileSize *int32 `json:"fileSize,omitempty"`
	FileSource *string `json:"fileSource,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Notify *PayloadNotify `json:"notify,omitempty"`
	NotifyCurrentUser *bool `json:"notifyCurrentUser,omitempty"`
	PendingFileRef *string `json:"pendingFileRef,omitempty"`
	Privacy *PayloadUserGroups `json:"privacy,omitempty"`
	Private *bool `json:"private,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewFileversionFileversion instantiates a new FileversionFileversion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileversionFileversion() *FileversionFileversion {
	this := FileversionFileversion{}
	return &this
}

// NewFileversionFileversionWithDefaults instantiates a new FileversionFileversion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileversionFileversionWithDefaults() *FileversionFileversion {
	this := FileversionFileversion{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *FileversionFileversion) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *FileversionFileversion) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *FileversionFileversion) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FileversionFileversion) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FileversionFileversion) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FileversionFileversion) SetDescription(v string) {
	o.Description = &v
}

// GetFileData returns the FileData field value if set, zero value otherwise.
func (o *FileversionFileversion) GetFileData() string {
	if o == nil || o.FileData == nil {
		var ret string
		return ret
	}
	return *o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetFileDataOk() (*string, bool) {
	if o == nil || o.FileData == nil {
		return nil, false
	}
	return o.FileData, true
}

// HasFileData returns a boolean if a field has been set.
func (o *FileversionFileversion) HasFileData() bool {
	if o != nil && o.FileData != nil {
		return true
	}

	return false
}

// SetFileData gets a reference to the given string and assigns it to the FileData field.
func (o *FileversionFileversion) SetFileData(v string) {
	o.FileData = &v
}

// GetFileRef returns the FileRef field value if set, zero value otherwise.
func (o *FileversionFileversion) GetFileRef() string {
	if o == nil || o.FileRef == nil {
		var ret string
		return ret
	}
	return *o.FileRef
}

// GetFileRefOk returns a tuple with the FileRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetFileRefOk() (*string, bool) {
	if o == nil || o.FileRef == nil {
		return nil, false
	}
	return o.FileRef, true
}

// HasFileRef returns a boolean if a field has been set.
func (o *FileversionFileversion) HasFileRef() bool {
	if o != nil && o.FileRef != nil {
		return true
	}

	return false
}

// SetFileRef gets a reference to the given string and assigns it to the FileRef field.
func (o *FileversionFileversion) SetFileRef(v string) {
	o.FileRef = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *FileversionFileversion) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *FileversionFileversion) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *FileversionFileversion) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileSource returns the FileSource field value if set, zero value otherwise.
func (o *FileversionFileversion) GetFileSource() string {
	if o == nil || o.FileSource == nil {
		var ret string
		return ret
	}
	return *o.FileSource
}

// GetFileSourceOk returns a tuple with the FileSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetFileSourceOk() (*string, bool) {
	if o == nil || o.FileSource == nil {
		return nil, false
	}
	return o.FileSource, true
}

// HasFileSource returns a boolean if a field has been set.
func (o *FileversionFileversion) HasFileSource() bool {
	if o != nil && o.FileSource != nil {
		return true
	}

	return false
}

// SetFileSource gets a reference to the given string and assigns it to the FileSource field.
func (o *FileversionFileversion) SetFileSource(v string) {
	o.FileSource = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileversionFileversion) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileversionFileversion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileversionFileversion) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileversionFileversion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileversionFileversion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileversionFileversion) SetName(v string) {
	o.Name = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *FileversionFileversion) GetNotify() PayloadNotify {
	if o == nil || o.Notify == nil {
		var ret PayloadNotify
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetNotifyOk() (*PayloadNotify, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *FileversionFileversion) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given PayloadNotify and assigns it to the Notify field.
func (o *FileversionFileversion) SetNotify(v PayloadNotify) {
	o.Notify = &v
}

// GetNotifyCurrentUser returns the NotifyCurrentUser field value if set, zero value otherwise.
func (o *FileversionFileversion) GetNotifyCurrentUser() bool {
	if o == nil || o.NotifyCurrentUser == nil {
		var ret bool
		return ret
	}
	return *o.NotifyCurrentUser
}

// GetNotifyCurrentUserOk returns a tuple with the NotifyCurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetNotifyCurrentUserOk() (*bool, bool) {
	if o == nil || o.NotifyCurrentUser == nil {
		return nil, false
	}
	return o.NotifyCurrentUser, true
}

// HasNotifyCurrentUser returns a boolean if a field has been set.
func (o *FileversionFileversion) HasNotifyCurrentUser() bool {
	if o != nil && o.NotifyCurrentUser != nil {
		return true
	}

	return false
}

// SetNotifyCurrentUser gets a reference to the given bool and assigns it to the NotifyCurrentUser field.
func (o *FileversionFileversion) SetNotifyCurrentUser(v bool) {
	o.NotifyCurrentUser = &v
}

// GetPendingFileRef returns the PendingFileRef field value if set, zero value otherwise.
func (o *FileversionFileversion) GetPendingFileRef() string {
	if o == nil || o.PendingFileRef == nil {
		var ret string
		return ret
	}
	return *o.PendingFileRef
}

// GetPendingFileRefOk returns a tuple with the PendingFileRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetPendingFileRefOk() (*string, bool) {
	if o == nil || o.PendingFileRef == nil {
		return nil, false
	}
	return o.PendingFileRef, true
}

// HasPendingFileRef returns a boolean if a field has been set.
func (o *FileversionFileversion) HasPendingFileRef() bool {
	if o != nil && o.PendingFileRef != nil {
		return true
	}

	return false
}

// SetPendingFileRef gets a reference to the given string and assigns it to the PendingFileRef field.
func (o *FileversionFileversion) SetPendingFileRef(v string) {
	o.PendingFileRef = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *FileversionFileversion) GetPrivacy() PayloadUserGroups {
	if o == nil || o.Privacy == nil {
		var ret PayloadUserGroups
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetPrivacyOk() (*PayloadUserGroups, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *FileversionFileversion) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given PayloadUserGroups and assigns it to the Privacy field.
func (o *FileversionFileversion) SetPrivacy(v PayloadUserGroups) {
	o.Privacy = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FileversionFileversion) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FileversionFileversion) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *FileversionFileversion) SetPrivate(v bool) {
	o.Private = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *FileversionFileversion) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileversionFileversion) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *FileversionFileversion) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *FileversionFileversion) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o FileversionFileversion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FileData != nil {
		toSerialize["fileData"] = o.FileData
	}
	if o.FileRef != nil {
		toSerialize["fileRef"] = o.FileRef
	}
	if o.FileSize != nil {
		toSerialize["fileSize"] = o.FileSize
	}
	if o.FileSource != nil {
		toSerialize["fileSource"] = o.FileSource
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.PendingFileRef != nil {
		toSerialize["pendingFileRef"] = o.PendingFileRef
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableFileversionFileversion struct {
	value *FileversionFileversion
	isSet bool
}

func (v NullableFileversionFileversion) Get() *FileversionFileversion {
	return v.value
}

func (v *NullableFileversionFileversion) Set(val *FileversionFileversion) {
	v.value = val
	v.isSet = true
}

func (v NullableFileversionFileversion) IsSet() bool {
	return v.isSet
}

func (v *NullableFileversionFileversion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileversionFileversion(val *FileversionFileversion) *NullableFileversionFileversion {
	return &NullableFileversionFileversion{value: val, isSet: true}
}

func (v NullableFileversionFileversion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileversionFileversion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


