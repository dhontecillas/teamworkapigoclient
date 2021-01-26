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

// ProjectsTemplateTemplateIDJsonProject struct for ProjectsTemplateTemplateIDJsonProject
type ProjectsTemplateTemplateIDJsonProject struct {
	CategoryId *int32 `json:"category-id,omitempty"`
	CompanyId *int32 `json:"companyId,omitempty"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
	People *string `json:"people,omitempty"`
	ProjectOwnerId *string `json:"projectOwnerId,omitempty"`
	Tags *string `json:"tags,omitempty"`
	TemplateDateTargetDefault string `json:"templateDateTargetDefault"`
	UseBilling *bool `json:"use-billing,omitempty"`
	UseComments *bool `json:"use-comments,omitempty"`
	UseFiles *bool `json:"use-files,omitempty"`
	UseLinks *bool `json:"use-links,omitempty"`
	UseMessages *bool `json:"use-messages,omitempty"`
	UseMilestones *bool `json:"use-milestones,omitempty"`
	UseNotebook *bool `json:"use-notebook,omitempty"`
	UseRiskregister *bool `json:"use-riskregister,omitempty"`
	UseTasks *bool `json:"use-tasks,omitempty"`
	UseTime *bool `json:"use-time,omitempty"`
}

// NewProjectsTemplateTemplateIDJsonProject instantiates a new ProjectsTemplateTemplateIDJsonProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsTemplateTemplateIDJsonProject(name string, templateDateTargetDefault string, ) *ProjectsTemplateTemplateIDJsonProject {
	this := ProjectsTemplateTemplateIDJsonProject{}
	this.Name = name
	this.TemplateDateTargetDefault = templateDateTargetDefault
	return &this
}

// NewProjectsTemplateTemplateIDJsonProjectWithDefaults instantiates a new ProjectsTemplateTemplateIDJsonProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsTemplateTemplateIDJsonProjectWithDefaults() *ProjectsTemplateTemplateIDJsonProject {
	this := ProjectsTemplateTemplateIDJsonProject{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetCompanyId() int32 {
	if o == nil || o.CompanyId == nil {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetCompanyIdOk() (*int32, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *ProjectsTemplateTemplateIDJsonProject) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectsTemplateTemplateIDJsonProject) SetName(v string) {
	o.Name = v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetPeople() string {
	if o == nil || o.People == nil {
		var ret string
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetPeopleOk() (*string, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given string and assigns it to the People field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetPeople(v string) {
	o.People = &v
}

// GetProjectOwnerId returns the ProjectOwnerId field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetProjectOwnerId() string {
	if o == nil || o.ProjectOwnerId == nil {
		var ret string
		return ret
	}
	return *o.ProjectOwnerId
}

// GetProjectOwnerIdOk returns a tuple with the ProjectOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetProjectOwnerIdOk() (*string, bool) {
	if o == nil || o.ProjectOwnerId == nil {
		return nil, false
	}
	return o.ProjectOwnerId, true
}

// HasProjectOwnerId returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasProjectOwnerId() bool {
	if o != nil && o.ProjectOwnerId != nil {
		return true
	}

	return false
}

// SetProjectOwnerId gets a reference to the given string and assigns it to the ProjectOwnerId field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetProjectOwnerId(v string) {
	o.ProjectOwnerId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetTags(v string) {
	o.Tags = &v
}

// GetTemplateDateTargetDefault returns the TemplateDateTargetDefault field value
func (o *ProjectsTemplateTemplateIDJsonProject) GetTemplateDateTargetDefault() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TemplateDateTargetDefault
}

// GetTemplateDateTargetDefaultOk returns a tuple with the TemplateDateTargetDefault field value
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetTemplateDateTargetDefaultOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateDateTargetDefault, true
}

// SetTemplateDateTargetDefault sets field value
func (o *ProjectsTemplateTemplateIDJsonProject) SetTemplateDateTargetDefault(v string) {
	o.TemplateDateTargetDefault = v
}

// GetUseBilling returns the UseBilling field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseBilling() bool {
	if o == nil || o.UseBilling == nil {
		var ret bool
		return ret
	}
	return *o.UseBilling
}

// GetUseBillingOk returns a tuple with the UseBilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseBillingOk() (*bool, bool) {
	if o == nil || o.UseBilling == nil {
		return nil, false
	}
	return o.UseBilling, true
}

// HasUseBilling returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseBilling() bool {
	if o != nil && o.UseBilling != nil {
		return true
	}

	return false
}

// SetUseBilling gets a reference to the given bool and assigns it to the UseBilling field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseBilling(v bool) {
	o.UseBilling = &v
}

// GetUseComments returns the UseComments field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseComments() bool {
	if o == nil || o.UseComments == nil {
		var ret bool
		return ret
	}
	return *o.UseComments
}

// GetUseCommentsOk returns a tuple with the UseComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseCommentsOk() (*bool, bool) {
	if o == nil || o.UseComments == nil {
		return nil, false
	}
	return o.UseComments, true
}

// HasUseComments returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseComments() bool {
	if o != nil && o.UseComments != nil {
		return true
	}

	return false
}

// SetUseComments gets a reference to the given bool and assigns it to the UseComments field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseComments(v bool) {
	o.UseComments = &v
}

// GetUseFiles returns the UseFiles field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseFiles() bool {
	if o == nil || o.UseFiles == nil {
		var ret bool
		return ret
	}
	return *o.UseFiles
}

// GetUseFilesOk returns a tuple with the UseFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseFilesOk() (*bool, bool) {
	if o == nil || o.UseFiles == nil {
		return nil, false
	}
	return o.UseFiles, true
}

// HasUseFiles returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseFiles() bool {
	if o != nil && o.UseFiles != nil {
		return true
	}

	return false
}

// SetUseFiles gets a reference to the given bool and assigns it to the UseFiles field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseFiles(v bool) {
	o.UseFiles = &v
}

// GetUseLinks returns the UseLinks field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseLinks() bool {
	if o == nil || o.UseLinks == nil {
		var ret bool
		return ret
	}
	return *o.UseLinks
}

// GetUseLinksOk returns a tuple with the UseLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseLinksOk() (*bool, bool) {
	if o == nil || o.UseLinks == nil {
		return nil, false
	}
	return o.UseLinks, true
}

// HasUseLinks returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseLinks() bool {
	if o != nil && o.UseLinks != nil {
		return true
	}

	return false
}

// SetUseLinks gets a reference to the given bool and assigns it to the UseLinks field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseLinks(v bool) {
	o.UseLinks = &v
}

// GetUseMessages returns the UseMessages field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseMessages() bool {
	if o == nil || o.UseMessages == nil {
		var ret bool
		return ret
	}
	return *o.UseMessages
}

// GetUseMessagesOk returns a tuple with the UseMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseMessagesOk() (*bool, bool) {
	if o == nil || o.UseMessages == nil {
		return nil, false
	}
	return o.UseMessages, true
}

// HasUseMessages returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseMessages() bool {
	if o != nil && o.UseMessages != nil {
		return true
	}

	return false
}

// SetUseMessages gets a reference to the given bool and assigns it to the UseMessages field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseMessages(v bool) {
	o.UseMessages = &v
}

// GetUseMilestones returns the UseMilestones field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseMilestones() bool {
	if o == nil || o.UseMilestones == nil {
		var ret bool
		return ret
	}
	return *o.UseMilestones
}

// GetUseMilestonesOk returns a tuple with the UseMilestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseMilestonesOk() (*bool, bool) {
	if o == nil || o.UseMilestones == nil {
		return nil, false
	}
	return o.UseMilestones, true
}

// HasUseMilestones returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseMilestones() bool {
	if o != nil && o.UseMilestones != nil {
		return true
	}

	return false
}

// SetUseMilestones gets a reference to the given bool and assigns it to the UseMilestones field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseMilestones(v bool) {
	o.UseMilestones = &v
}

// GetUseNotebook returns the UseNotebook field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseNotebook() bool {
	if o == nil || o.UseNotebook == nil {
		var ret bool
		return ret
	}
	return *o.UseNotebook
}

// GetUseNotebookOk returns a tuple with the UseNotebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseNotebookOk() (*bool, bool) {
	if o == nil || o.UseNotebook == nil {
		return nil, false
	}
	return o.UseNotebook, true
}

// HasUseNotebook returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseNotebook() bool {
	if o != nil && o.UseNotebook != nil {
		return true
	}

	return false
}

// SetUseNotebook gets a reference to the given bool and assigns it to the UseNotebook field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseNotebook(v bool) {
	o.UseNotebook = &v
}

// GetUseRiskregister returns the UseRiskregister field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseRiskregister() bool {
	if o == nil || o.UseRiskregister == nil {
		var ret bool
		return ret
	}
	return *o.UseRiskregister
}

// GetUseRiskregisterOk returns a tuple with the UseRiskregister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseRiskregisterOk() (*bool, bool) {
	if o == nil || o.UseRiskregister == nil {
		return nil, false
	}
	return o.UseRiskregister, true
}

// HasUseRiskregister returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseRiskregister() bool {
	if o != nil && o.UseRiskregister != nil {
		return true
	}

	return false
}

// SetUseRiskregister gets a reference to the given bool and assigns it to the UseRiskregister field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseRiskregister(v bool) {
	o.UseRiskregister = &v
}

// GetUseTasks returns the UseTasks field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseTasks() bool {
	if o == nil || o.UseTasks == nil {
		var ret bool
		return ret
	}
	return *o.UseTasks
}

// GetUseTasksOk returns a tuple with the UseTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseTasksOk() (*bool, bool) {
	if o == nil || o.UseTasks == nil {
		return nil, false
	}
	return o.UseTasks, true
}

// HasUseTasks returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseTasks() bool {
	if o != nil && o.UseTasks != nil {
		return true
	}

	return false
}

// SetUseTasks gets a reference to the given bool and assigns it to the UseTasks field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseTasks(v bool) {
	o.UseTasks = &v
}

// GetUseTime returns the UseTime field value if set, zero value otherwise.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseTime() bool {
	if o == nil || o.UseTime == nil {
		var ret bool
		return ret
	}
	return *o.UseTime
}

// GetUseTimeOk returns a tuple with the UseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) GetUseTimeOk() (*bool, bool) {
	if o == nil || o.UseTime == nil {
		return nil, false
	}
	return o.UseTime, true
}

// HasUseTime returns a boolean if a field has been set.
func (o *ProjectsTemplateTemplateIDJsonProject) HasUseTime() bool {
	if o != nil && o.UseTime != nil {
		return true
	}

	return false
}

// SetUseTime gets a reference to the given bool and assigns it to the UseTime field.
func (o *ProjectsTemplateTemplateIDJsonProject) SetUseTime(v bool) {
	o.UseTime = &v
}

func (o ProjectsTemplateTemplateIDJsonProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryId != nil {
		toSerialize["category-id"] = o.CategoryId
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	if o.ProjectOwnerId != nil {
		toSerialize["projectOwnerId"] = o.ProjectOwnerId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["templateDateTargetDefault"] = o.TemplateDateTargetDefault
	}
	if o.UseBilling != nil {
		toSerialize["use-billing"] = o.UseBilling
	}
	if o.UseComments != nil {
		toSerialize["use-comments"] = o.UseComments
	}
	if o.UseFiles != nil {
		toSerialize["use-files"] = o.UseFiles
	}
	if o.UseLinks != nil {
		toSerialize["use-links"] = o.UseLinks
	}
	if o.UseMessages != nil {
		toSerialize["use-messages"] = o.UseMessages
	}
	if o.UseMilestones != nil {
		toSerialize["use-milestones"] = o.UseMilestones
	}
	if o.UseNotebook != nil {
		toSerialize["use-notebook"] = o.UseNotebook
	}
	if o.UseRiskregister != nil {
		toSerialize["use-riskregister"] = o.UseRiskregister
	}
	if o.UseTasks != nil {
		toSerialize["use-tasks"] = o.UseTasks
	}
	if o.UseTime != nil {
		toSerialize["use-time"] = o.UseTime
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsTemplateTemplateIDJsonProject struct {
	value *ProjectsTemplateTemplateIDJsonProject
	isSet bool
}

func (v NullableProjectsTemplateTemplateIDJsonProject) Get() *ProjectsTemplateTemplateIDJsonProject {
	return v.value
}

func (v *NullableProjectsTemplateTemplateIDJsonProject) Set(val *ProjectsTemplateTemplateIDJsonProject) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsTemplateTemplateIDJsonProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsTemplateTemplateIDJsonProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsTemplateTemplateIDJsonProject(val *ProjectsTemplateTemplateIDJsonProject) *NullableProjectsTemplateTemplateIDJsonProject {
	return &NullableProjectsTemplateTemplateIDJsonProject{value: val, isSet: true}
}

func (v NullableProjectsTemplateTemplateIDJsonProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsTemplateTemplateIDJsonProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


