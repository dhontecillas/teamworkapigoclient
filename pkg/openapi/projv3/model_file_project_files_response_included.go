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

// FileProjectFilesResponseIncluded struct for FileProjectFilesResponseIncluded
type FileProjectFilesResponseIncluded struct {
	Comments *map[string]ViewComment `json:"comments,omitempty"`
	FileCategories *map[string]ViewProjectFileCategory `json:"fileCategories,omitempty"`
	Messages *map[string]ViewProjectMessage `json:"messages,omitempty"`
	Permissions *map[string]ViewFilePermissions `json:"permissions,omitempty"`
	Projects *map[string]ViewProject `json:"projects,omitempty"`
	Tags *map[string]ViewTag `json:"tags,omitempty"`
	Tasks *map[string]ViewTask `json:"tasks,omitempty"`
	Users *map[string]ViewUser `json:"users,omitempty"`
}

// NewFileProjectFilesResponseIncluded instantiates a new FileProjectFilesResponseIncluded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileProjectFilesResponseIncluded() *FileProjectFilesResponseIncluded {
	this := FileProjectFilesResponseIncluded{}
	return &this
}

// NewFileProjectFilesResponseIncludedWithDefaults instantiates a new FileProjectFilesResponseIncluded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileProjectFilesResponseIncludedWithDefaults() *FileProjectFilesResponseIncluded {
	this := FileProjectFilesResponseIncluded{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetComments() map[string]ViewComment {
	if o == nil || o.Comments == nil {
		var ret map[string]ViewComment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetCommentsOk() (*map[string]ViewComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given map[string]ViewComment and assigns it to the Comments field.
func (o *FileProjectFilesResponseIncluded) SetComments(v map[string]ViewComment) {
	o.Comments = &v
}

// GetFileCategories returns the FileCategories field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetFileCategories() map[string]ViewProjectFileCategory {
	if o == nil || o.FileCategories == nil {
		var ret map[string]ViewProjectFileCategory
		return ret
	}
	return *o.FileCategories
}

// GetFileCategoriesOk returns a tuple with the FileCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetFileCategoriesOk() (*map[string]ViewProjectFileCategory, bool) {
	if o == nil || o.FileCategories == nil {
		return nil, false
	}
	return o.FileCategories, true
}

// HasFileCategories returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasFileCategories() bool {
	if o != nil && o.FileCategories != nil {
		return true
	}

	return false
}

// SetFileCategories gets a reference to the given map[string]ViewProjectFileCategory and assigns it to the FileCategories field.
func (o *FileProjectFilesResponseIncluded) SetFileCategories(v map[string]ViewProjectFileCategory) {
	o.FileCategories = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetMessages() map[string]ViewProjectMessage {
	if o == nil || o.Messages == nil {
		var ret map[string]ViewProjectMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetMessagesOk() (*map[string]ViewProjectMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given map[string]ViewProjectMessage and assigns it to the Messages field.
func (o *FileProjectFilesResponseIncluded) SetMessages(v map[string]ViewProjectMessage) {
	o.Messages = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetPermissions() map[string]ViewFilePermissions {
	if o == nil || o.Permissions == nil {
		var ret map[string]ViewFilePermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetPermissionsOk() (*map[string]ViewFilePermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]ViewFilePermissions and assigns it to the Permissions field.
func (o *FileProjectFilesResponseIncluded) SetPermissions(v map[string]ViewFilePermissions) {
	o.Permissions = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetProjects() map[string]ViewProject {
	if o == nil || o.Projects == nil {
		var ret map[string]ViewProject
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetProjectsOk() (*map[string]ViewProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given map[string]ViewProject and assigns it to the Projects field.
func (o *FileProjectFilesResponseIncluded) SetProjects(v map[string]ViewProject) {
	o.Projects = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetTags() map[string]ViewTag {
	if o == nil || o.Tags == nil {
		var ret map[string]ViewTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetTagsOk() (*map[string]ViewTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]ViewTag and assigns it to the Tags field.
func (o *FileProjectFilesResponseIncluded) SetTags(v map[string]ViewTag) {
	o.Tags = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetTasks() map[string]ViewTask {
	if o == nil || o.Tasks == nil {
		var ret map[string]ViewTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetTasksOk() (*map[string]ViewTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given map[string]ViewTask and assigns it to the Tasks field.
func (o *FileProjectFilesResponseIncluded) SetTasks(v map[string]ViewTask) {
	o.Tasks = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *FileProjectFilesResponseIncluded) GetUsers() map[string]ViewUser {
	if o == nil || o.Users == nil {
		var ret map[string]ViewUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileProjectFilesResponseIncluded) GetUsersOk() (*map[string]ViewUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *FileProjectFilesResponseIncluded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given map[string]ViewUser and assigns it to the Users field.
func (o *FileProjectFilesResponseIncluded) SetUsers(v map[string]ViewUser) {
	o.Users = &v
}

func (o FileProjectFilesResponseIncluded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.FileCategories != nil {
		toSerialize["fileCategories"] = o.FileCategories
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableFileProjectFilesResponseIncluded struct {
	value *FileProjectFilesResponseIncluded
	isSet bool
}

func (v NullableFileProjectFilesResponseIncluded) Get() *FileProjectFilesResponseIncluded {
	return v.value
}

func (v *NullableFileProjectFilesResponseIncluded) Set(val *FileProjectFilesResponseIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableFileProjectFilesResponseIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableFileProjectFilesResponseIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileProjectFilesResponseIncluded(val *FileProjectFilesResponseIncluded) *NullableFileProjectFilesResponseIncluded {
	return &NullableFileProjectFilesResponseIncluded{value: val, isSet: true}
}

func (v NullableFileProjectFilesResponseIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileProjectFilesResponseIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


