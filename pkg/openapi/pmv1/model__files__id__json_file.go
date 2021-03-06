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

// FilesIdJsonFile struct for FilesIdJsonFile
type FilesIdJsonFile struct {
	CategoryId *string `json:"category-id,omitempty"`
	CategoryName *string `json:"category-name,omitempty"`
	CommentsCount *string `json:"comments-count,omitempty"`
	Description *string `json:"description,omitempty"`
	DownloadURL *string `json:"download-URL,omitempty"`
	FileSource *string `json:"file-source,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	Private *string `json:"private,omitempty"`
	ProjectId *string `json:"project-id,omitempty"`
	ProjectName *string `json:"project-name,omitempty"`
	Size *string `json:"size,omitempty"`
	UploadedByUserFirstName *string `json:"uploaded-by-user-first-name,omitempty"`
	UploadedByUserLastName *string `json:"uploaded-by-user-last-name,omitempty"`
	UploadedByUserId *string `json:"uploaded-by-userId,omitempty"`
	UploadedDate *string `json:"uploaded-date,omitempty"`
	Version *string `json:"version,omitempty"`
	VersionId *string `json:"version-id,omitempty"`
}

// NewFilesIdJsonFile instantiates a new FilesIdJsonFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesIdJsonFile() *FilesIdJsonFile {
	this := FilesIdJsonFile{}
	return &this
}

// NewFilesIdJsonFileWithDefaults instantiates a new FilesIdJsonFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesIdJsonFileWithDefaults() *FilesIdJsonFile {
	this := FilesIdJsonFile{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *FilesIdJsonFile) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *FilesIdJsonFile) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetCommentsCount returns the CommentsCount field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetCommentsCount() string {
	if o == nil || o.CommentsCount == nil {
		var ret string
		return ret
	}
	return *o.CommentsCount
}

// GetCommentsCountOk returns a tuple with the CommentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetCommentsCountOk() (*string, bool) {
	if o == nil || o.CommentsCount == nil {
		return nil, false
	}
	return o.CommentsCount, true
}

// HasCommentsCount returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasCommentsCount() bool {
	if o != nil && o.CommentsCount != nil {
		return true
	}

	return false
}

// SetCommentsCount gets a reference to the given string and assigns it to the CommentsCount field.
func (o *FilesIdJsonFile) SetCommentsCount(v string) {
	o.CommentsCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FilesIdJsonFile) SetDescription(v string) {
	o.Description = &v
}

// GetDownloadURL returns the DownloadURL field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetDownloadURL() string {
	if o == nil || o.DownloadURL == nil {
		var ret string
		return ret
	}
	return *o.DownloadURL
}

// GetDownloadURLOk returns a tuple with the DownloadURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetDownloadURLOk() (*string, bool) {
	if o == nil || o.DownloadURL == nil {
		return nil, false
	}
	return o.DownloadURL, true
}

// HasDownloadURL returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasDownloadURL() bool {
	if o != nil && o.DownloadURL != nil {
		return true
	}

	return false
}

// SetDownloadURL gets a reference to the given string and assigns it to the DownloadURL field.
func (o *FilesIdJsonFile) SetDownloadURL(v string) {
	o.DownloadURL = &v
}

// GetFileSource returns the FileSource field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetFileSource() string {
	if o == nil || o.FileSource == nil {
		var ret string
		return ret
	}
	return *o.FileSource
}

// GetFileSourceOk returns a tuple with the FileSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetFileSourceOk() (*string, bool) {
	if o == nil || o.FileSource == nil {
		return nil, false
	}
	return o.FileSource, true
}

// HasFileSource returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasFileSource() bool {
	if o != nil && o.FileSource != nil {
		return true
	}

	return false
}

// SetFileSource gets a reference to the given string and assigns it to the FileSource field.
func (o *FilesIdJsonFile) SetFileSource(v string) {
	o.FileSource = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FilesIdJsonFile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilesIdJsonFile) SetName(v string) {
	o.Name = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetOriginalName() string {
	if o == nil || o.OriginalName == nil {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetOriginalNameOk() (*string, bool) {
	if o == nil || o.OriginalName == nil {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasOriginalName() bool {
	if o != nil && o.OriginalName != nil {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *FilesIdJsonFile) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetPrivate() string {
	if o == nil || o.Private == nil {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetPrivateOk() (*string, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *FilesIdJsonFile) SetPrivate(v string) {
	o.Private = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *FilesIdJsonFile) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *FilesIdJsonFile) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *FilesIdJsonFile) SetSize(v string) {
	o.Size = &v
}

// GetUploadedByUserFirstName returns the UploadedByUserFirstName field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetUploadedByUserFirstName() string {
	if o == nil || o.UploadedByUserFirstName == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserFirstName
}

// GetUploadedByUserFirstNameOk returns a tuple with the UploadedByUserFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetUploadedByUserFirstNameOk() (*string, bool) {
	if o == nil || o.UploadedByUserFirstName == nil {
		return nil, false
	}
	return o.UploadedByUserFirstName, true
}

// HasUploadedByUserFirstName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasUploadedByUserFirstName() bool {
	if o != nil && o.UploadedByUserFirstName != nil {
		return true
	}

	return false
}

// SetUploadedByUserFirstName gets a reference to the given string and assigns it to the UploadedByUserFirstName field.
func (o *FilesIdJsonFile) SetUploadedByUserFirstName(v string) {
	o.UploadedByUserFirstName = &v
}

// GetUploadedByUserLastName returns the UploadedByUserLastName field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetUploadedByUserLastName() string {
	if o == nil || o.UploadedByUserLastName == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserLastName
}

// GetUploadedByUserLastNameOk returns a tuple with the UploadedByUserLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetUploadedByUserLastNameOk() (*string, bool) {
	if o == nil || o.UploadedByUserLastName == nil {
		return nil, false
	}
	return o.UploadedByUserLastName, true
}

// HasUploadedByUserLastName returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasUploadedByUserLastName() bool {
	if o != nil && o.UploadedByUserLastName != nil {
		return true
	}

	return false
}

// SetUploadedByUserLastName gets a reference to the given string and assigns it to the UploadedByUserLastName field.
func (o *FilesIdJsonFile) SetUploadedByUserLastName(v string) {
	o.UploadedByUserLastName = &v
}

// GetUploadedByUserId returns the UploadedByUserId field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetUploadedByUserId() string {
	if o == nil || o.UploadedByUserId == nil {
		var ret string
		return ret
	}
	return *o.UploadedByUserId
}

// GetUploadedByUserIdOk returns a tuple with the UploadedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetUploadedByUserIdOk() (*string, bool) {
	if o == nil || o.UploadedByUserId == nil {
		return nil, false
	}
	return o.UploadedByUserId, true
}

// HasUploadedByUserId returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasUploadedByUserId() bool {
	if o != nil && o.UploadedByUserId != nil {
		return true
	}

	return false
}

// SetUploadedByUserId gets a reference to the given string and assigns it to the UploadedByUserId field.
func (o *FilesIdJsonFile) SetUploadedByUserId(v string) {
	o.UploadedByUserId = &v
}

// GetUploadedDate returns the UploadedDate field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetUploadedDate() string {
	if o == nil || o.UploadedDate == nil {
		var ret string
		return ret
	}
	return *o.UploadedDate
}

// GetUploadedDateOk returns a tuple with the UploadedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetUploadedDateOk() (*string, bool) {
	if o == nil || o.UploadedDate == nil {
		return nil, false
	}
	return o.UploadedDate, true
}

// HasUploadedDate returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasUploadedDate() bool {
	if o != nil && o.UploadedDate != nil {
		return true
	}

	return false
}

// SetUploadedDate gets a reference to the given string and assigns it to the UploadedDate field.
func (o *FilesIdJsonFile) SetUploadedDate(v string) {
	o.UploadedDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FilesIdJsonFile) SetVersion(v string) {
	o.Version = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *FilesIdJsonFile) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesIdJsonFile) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *FilesIdJsonFile) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *FilesIdJsonFile) SetVersionId(v string) {
	o.VersionId = &v
}

func (o FilesIdJsonFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["category-id"] = o.CategoryId
	}
	if o.CategoryName != nil {
		toSerialize["category-name"] = o.CategoryName
	}
	if o.CommentsCount != nil {
		toSerialize["comments-count"] = o.CommentsCount
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DownloadURL != nil {
		toSerialize["download-URL"] = o.DownloadURL
	}
	if o.FileSource != nil {
		toSerialize["file-source"] = o.FileSource
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OriginalName != nil {
		toSerialize["originalName"] = o.OriginalName
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.ProjectName != nil {
		toSerialize["project-name"] = o.ProjectName
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.UploadedByUserFirstName != nil {
		toSerialize["uploaded-by-user-first-name"] = o.UploadedByUserFirstName
	}
	if o.UploadedByUserLastName != nil {
		toSerialize["uploaded-by-user-last-name"] = o.UploadedByUserLastName
	}
	if o.UploadedByUserId != nil {
		toSerialize["uploaded-by-userId"] = o.UploadedByUserId
	}
	if o.UploadedDate != nil {
		toSerialize["uploaded-date"] = o.UploadedDate
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionId != nil {
		toSerialize["version-id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableFilesIdJsonFile struct {
	value *FilesIdJsonFile
	isSet bool
}

func (v NullableFilesIdJsonFile) Get() *FilesIdJsonFile {
	return v.value
}

func (v *NullableFilesIdJsonFile) Set(val *FilesIdJsonFile) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesIdJsonFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesIdJsonFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesIdJsonFile(val *FilesIdJsonFile) *NullableFilesIdJsonFile {
	return &NullableFilesIdJsonFile{value: val, isSet: true}
}

func (v NullableFilesIdJsonFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesIdJsonFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


