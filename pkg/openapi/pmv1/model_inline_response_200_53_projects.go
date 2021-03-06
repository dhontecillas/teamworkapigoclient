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

// InlineResponse20053Projects struct for InlineResponse20053Projects
type InlineResponse20053Projects struct {
	Announcement *string `json:"announcement,omitempty"`
	AnnouncementHTML *string `json:"announcementHTML,omitempty"`
	BoardData *map[string]interface{} `json:"boardData,omitempty"`
	Category *InlineResponse2002Column `json:"category,omitempty"`
	Company *InlineResponse20014Company `json:"company,omitempty"`
	CreatedOn *string `json:"created-on,omitempty"`
	DefaultPrivacy *string `json:"defaultPrivacy,omitempty"`
	Defaults *InlineResponse20053Defaults `json:"defaults,omitempty"`
	Description *string `json:"description,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	FilesAutoNewVersion *bool `json:"filesAutoNewVersion,omitempty"`
	HarvestTimersEnabled *bool `json:"harvest-timers-enabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Integrations *InlineResponse20053Integrations `json:"integrations,omitempty"`
	IsProjectAdmin *bool `json:"isProjectAdmin,omitempty"`
	LastChangedOn *string `json:"last-changed-on,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	Notifyeveryone *bool `json:"notifyeveryone,omitempty"`
	OverviewStartPage *string `json:"overview-start-page,omitempty"`
	PrivacyEnabled *bool `json:"privacyEnabled,omitempty"`
	ReplyByEmailEnabled *bool `json:"replyByEmailEnabled,omitempty"`
	ShowAnnouncement *bool `json:"show-announcement,omitempty"`
	Starred *bool `json:"starred,omitempty"`
	StartPage *string `json:"start-page,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	Status *string `json:"status,omitempty"`
	SubStatus *string `json:"subStatus,omitempty"`
	Tags *[]map[string]interface{} `json:"tags,omitempty"`
	TasksStartPage *string `json:"tasks-start-page,omitempty"`
}

// NewInlineResponse20053Projects instantiates a new InlineResponse20053Projects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053Projects() *InlineResponse20053Projects {
	this := InlineResponse20053Projects{}
	return &this
}

// NewInlineResponse20053ProjectsWithDefaults instantiates a new InlineResponse20053Projects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053ProjectsWithDefaults() *InlineResponse20053Projects {
	this := InlineResponse20053Projects{}
	return &this
}

// GetAnnouncement returns the Announcement field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetAnnouncement() string {
	if o == nil || o.Announcement == nil {
		var ret string
		return ret
	}
	return *o.Announcement
}

// GetAnnouncementOk returns a tuple with the Announcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetAnnouncementOk() (*string, bool) {
	if o == nil || o.Announcement == nil {
		return nil, false
	}
	return o.Announcement, true
}

// HasAnnouncement returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasAnnouncement() bool {
	if o != nil && o.Announcement != nil {
		return true
	}

	return false
}

// SetAnnouncement gets a reference to the given string and assigns it to the Announcement field.
func (o *InlineResponse20053Projects) SetAnnouncement(v string) {
	o.Announcement = &v
}

// GetAnnouncementHTML returns the AnnouncementHTML field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetAnnouncementHTML() string {
	if o == nil || o.AnnouncementHTML == nil {
		var ret string
		return ret
	}
	return *o.AnnouncementHTML
}

// GetAnnouncementHTMLOk returns a tuple with the AnnouncementHTML field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetAnnouncementHTMLOk() (*string, bool) {
	if o == nil || o.AnnouncementHTML == nil {
		return nil, false
	}
	return o.AnnouncementHTML, true
}

// HasAnnouncementHTML returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasAnnouncementHTML() bool {
	if o != nil && o.AnnouncementHTML != nil {
		return true
	}

	return false
}

// SetAnnouncementHTML gets a reference to the given string and assigns it to the AnnouncementHTML field.
func (o *InlineResponse20053Projects) SetAnnouncementHTML(v string) {
	o.AnnouncementHTML = &v
}

// GetBoardData returns the BoardData field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetBoardData() map[string]interface{} {
	if o == nil || o.BoardData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.BoardData
}

// GetBoardDataOk returns a tuple with the BoardData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetBoardDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.BoardData == nil {
		return nil, false
	}
	return o.BoardData, true
}

// HasBoardData returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasBoardData() bool {
	if o != nil && o.BoardData != nil {
		return true
	}

	return false
}

// SetBoardData gets a reference to the given map[string]interface{} and assigns it to the BoardData field.
func (o *InlineResponse20053Projects) SetBoardData(v map[string]interface{}) {
	o.BoardData = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetCategory() InlineResponse2002Column {
	if o == nil || o.Category == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetCategoryOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given InlineResponse2002Column and assigns it to the Category field.
func (o *InlineResponse20053Projects) SetCategory(v InlineResponse2002Column) {
	o.Category = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetCompany() InlineResponse20014Company {
	if o == nil || o.Company == nil {
		var ret InlineResponse20014Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetCompanyOk() (*InlineResponse20014Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given InlineResponse20014Company and assigns it to the Company field.
func (o *InlineResponse20053Projects) SetCompany(v InlineResponse20014Company) {
	o.Company = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *InlineResponse20053Projects) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetDefaultPrivacy returns the DefaultPrivacy field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetDefaultPrivacy() string {
	if o == nil || o.DefaultPrivacy == nil {
		var ret string
		return ret
	}
	return *o.DefaultPrivacy
}

// GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetDefaultPrivacyOk() (*string, bool) {
	if o == nil || o.DefaultPrivacy == nil {
		return nil, false
	}
	return o.DefaultPrivacy, true
}

// HasDefaultPrivacy returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasDefaultPrivacy() bool {
	if o != nil && o.DefaultPrivacy != nil {
		return true
	}

	return false
}

// SetDefaultPrivacy gets a reference to the given string and assigns it to the DefaultPrivacy field.
func (o *InlineResponse20053Projects) SetDefaultPrivacy(v string) {
	o.DefaultPrivacy = &v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetDefaults() InlineResponse20053Defaults {
	if o == nil || o.Defaults == nil {
		var ret InlineResponse20053Defaults
		return ret
	}
	return *o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetDefaultsOk() (*InlineResponse20053Defaults, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasDefaults() bool {
	if o != nil && o.Defaults != nil {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given InlineResponse20053Defaults and assigns it to the Defaults field.
func (o *InlineResponse20053Projects) SetDefaults(v InlineResponse20053Defaults) {
	o.Defaults = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20053Projects) SetDescription(v string) {
	o.Description = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *InlineResponse20053Projects) SetEndDate(v string) {
	o.EndDate = &v
}

// GetFilesAutoNewVersion returns the FilesAutoNewVersion field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetFilesAutoNewVersion() bool {
	if o == nil || o.FilesAutoNewVersion == nil {
		var ret bool
		return ret
	}
	return *o.FilesAutoNewVersion
}

// GetFilesAutoNewVersionOk returns a tuple with the FilesAutoNewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetFilesAutoNewVersionOk() (*bool, bool) {
	if o == nil || o.FilesAutoNewVersion == nil {
		return nil, false
	}
	return o.FilesAutoNewVersion, true
}

// HasFilesAutoNewVersion returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasFilesAutoNewVersion() bool {
	if o != nil && o.FilesAutoNewVersion != nil {
		return true
	}

	return false
}

// SetFilesAutoNewVersion gets a reference to the given bool and assigns it to the FilesAutoNewVersion field.
func (o *InlineResponse20053Projects) SetFilesAutoNewVersion(v bool) {
	o.FilesAutoNewVersion = &v
}

// GetHarvestTimersEnabled returns the HarvestTimersEnabled field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetHarvestTimersEnabled() bool {
	if o == nil || o.HarvestTimersEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HarvestTimersEnabled
}

// GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetHarvestTimersEnabledOk() (*bool, bool) {
	if o == nil || o.HarvestTimersEnabled == nil {
		return nil, false
	}
	return o.HarvestTimersEnabled, true
}

// HasHarvestTimersEnabled returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasHarvestTimersEnabled() bool {
	if o != nil && o.HarvestTimersEnabled != nil {
		return true
	}

	return false
}

// SetHarvestTimersEnabled gets a reference to the given bool and assigns it to the HarvestTimersEnabled field.
func (o *InlineResponse20053Projects) SetHarvestTimersEnabled(v bool) {
	o.HarvestTimersEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20053Projects) SetId(v string) {
	o.Id = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetIntegrations() InlineResponse20053Integrations {
	if o == nil || o.Integrations == nil {
		var ret InlineResponse20053Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetIntegrationsOk() (*InlineResponse20053Integrations, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasIntegrations() bool {
	if o != nil && o.Integrations != nil {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given InlineResponse20053Integrations and assigns it to the Integrations field.
func (o *InlineResponse20053Projects) SetIntegrations(v InlineResponse20053Integrations) {
	o.Integrations = &v
}

// GetIsProjectAdmin returns the IsProjectAdmin field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetIsProjectAdmin() bool {
	if o == nil || o.IsProjectAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsProjectAdmin
}

// GetIsProjectAdminOk returns a tuple with the IsProjectAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetIsProjectAdminOk() (*bool, bool) {
	if o == nil || o.IsProjectAdmin == nil {
		return nil, false
	}
	return o.IsProjectAdmin, true
}

// HasIsProjectAdmin returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasIsProjectAdmin() bool {
	if o != nil && o.IsProjectAdmin != nil {
		return true
	}

	return false
}

// SetIsProjectAdmin gets a reference to the given bool and assigns it to the IsProjectAdmin field.
func (o *InlineResponse20053Projects) SetIsProjectAdmin(v bool) {
	o.IsProjectAdmin = &v
}

// GetLastChangedOn returns the LastChangedOn field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetLastChangedOn() string {
	if o == nil || o.LastChangedOn == nil {
		var ret string
		return ret
	}
	return *o.LastChangedOn
}

// GetLastChangedOnOk returns a tuple with the LastChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetLastChangedOnOk() (*string, bool) {
	if o == nil || o.LastChangedOn == nil {
		return nil, false
	}
	return o.LastChangedOn, true
}

// HasLastChangedOn returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasLastChangedOn() bool {
	if o != nil && o.LastChangedOn != nil {
		return true
	}

	return false
}

// SetLastChangedOn gets a reference to the given string and assigns it to the LastChangedOn field.
func (o *InlineResponse20053Projects) SetLastChangedOn(v string) {
	o.LastChangedOn = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *InlineResponse20053Projects) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20053Projects) SetName(v string) {
	o.Name = &v
}

// GetNotifyeveryone returns the Notifyeveryone field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetNotifyeveryone() bool {
	if o == nil || o.Notifyeveryone == nil {
		var ret bool
		return ret
	}
	return *o.Notifyeveryone
}

// GetNotifyeveryoneOk returns a tuple with the Notifyeveryone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetNotifyeveryoneOk() (*bool, bool) {
	if o == nil || o.Notifyeveryone == nil {
		return nil, false
	}
	return o.Notifyeveryone, true
}

// HasNotifyeveryone returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasNotifyeveryone() bool {
	if o != nil && o.Notifyeveryone != nil {
		return true
	}

	return false
}

// SetNotifyeveryone gets a reference to the given bool and assigns it to the Notifyeveryone field.
func (o *InlineResponse20053Projects) SetNotifyeveryone(v bool) {
	o.Notifyeveryone = &v
}

// GetOverviewStartPage returns the OverviewStartPage field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetOverviewStartPage() string {
	if o == nil || o.OverviewStartPage == nil {
		var ret string
		return ret
	}
	return *o.OverviewStartPage
}

// GetOverviewStartPageOk returns a tuple with the OverviewStartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetOverviewStartPageOk() (*string, bool) {
	if o == nil || o.OverviewStartPage == nil {
		return nil, false
	}
	return o.OverviewStartPage, true
}

// HasOverviewStartPage returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasOverviewStartPage() bool {
	if o != nil && o.OverviewStartPage != nil {
		return true
	}

	return false
}

// SetOverviewStartPage gets a reference to the given string and assigns it to the OverviewStartPage field.
func (o *InlineResponse20053Projects) SetOverviewStartPage(v string) {
	o.OverviewStartPage = &v
}

// GetPrivacyEnabled returns the PrivacyEnabled field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetPrivacyEnabled() bool {
	if o == nil || o.PrivacyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PrivacyEnabled
}

// GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetPrivacyEnabledOk() (*bool, bool) {
	if o == nil || o.PrivacyEnabled == nil {
		return nil, false
	}
	return o.PrivacyEnabled, true
}

// HasPrivacyEnabled returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasPrivacyEnabled() bool {
	if o != nil && o.PrivacyEnabled != nil {
		return true
	}

	return false
}

// SetPrivacyEnabled gets a reference to the given bool and assigns it to the PrivacyEnabled field.
func (o *InlineResponse20053Projects) SetPrivacyEnabled(v bool) {
	o.PrivacyEnabled = &v
}

// GetReplyByEmailEnabled returns the ReplyByEmailEnabled field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetReplyByEmailEnabled() bool {
	if o == nil || o.ReplyByEmailEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplyByEmailEnabled
}

// GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetReplyByEmailEnabledOk() (*bool, bool) {
	if o == nil || o.ReplyByEmailEnabled == nil {
		return nil, false
	}
	return o.ReplyByEmailEnabled, true
}

// HasReplyByEmailEnabled returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasReplyByEmailEnabled() bool {
	if o != nil && o.ReplyByEmailEnabled != nil {
		return true
	}

	return false
}

// SetReplyByEmailEnabled gets a reference to the given bool and assigns it to the ReplyByEmailEnabled field.
func (o *InlineResponse20053Projects) SetReplyByEmailEnabled(v bool) {
	o.ReplyByEmailEnabled = &v
}

// GetShowAnnouncement returns the ShowAnnouncement field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetShowAnnouncement() bool {
	if o == nil || o.ShowAnnouncement == nil {
		var ret bool
		return ret
	}
	return *o.ShowAnnouncement
}

// GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetShowAnnouncementOk() (*bool, bool) {
	if o == nil || o.ShowAnnouncement == nil {
		return nil, false
	}
	return o.ShowAnnouncement, true
}

// HasShowAnnouncement returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasShowAnnouncement() bool {
	if o != nil && o.ShowAnnouncement != nil {
		return true
	}

	return false
}

// SetShowAnnouncement gets a reference to the given bool and assigns it to the ShowAnnouncement field.
func (o *InlineResponse20053Projects) SetShowAnnouncement(v bool) {
	o.ShowAnnouncement = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetStarred() bool {
	if o == nil || o.Starred == nil {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetStarredOk() (*bool, bool) {
	if o == nil || o.Starred == nil {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasStarred() bool {
	if o != nil && o.Starred != nil {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *InlineResponse20053Projects) SetStarred(v bool) {
	o.Starred = &v
}

// GetStartPage returns the StartPage field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetStartPage() string {
	if o == nil || o.StartPage == nil {
		var ret string
		return ret
	}
	return *o.StartPage
}

// GetStartPageOk returns a tuple with the StartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetStartPageOk() (*string, bool) {
	if o == nil || o.StartPage == nil {
		return nil, false
	}
	return o.StartPage, true
}

// HasStartPage returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasStartPage() bool {
	if o != nil && o.StartPage != nil {
		return true
	}

	return false
}

// SetStartPage gets a reference to the given string and assigns it to the StartPage field.
func (o *InlineResponse20053Projects) SetStartPage(v string) {
	o.StartPage = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *InlineResponse20053Projects) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20053Projects) SetStatus(v string) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetSubStatus() string {
	if o == nil || o.SubStatus == nil {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetSubStatusOk() (*string, bool) {
	if o == nil || o.SubStatus == nil {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasSubStatus() bool {
	if o != nil && o.SubStatus != nil {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *InlineResponse20053Projects) SetSubStatus(v string) {
	o.SubStatus = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetTags() []map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetTagsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []map[string]interface{} and assigns it to the Tags field.
func (o *InlineResponse20053Projects) SetTags(v []map[string]interface{}) {
	o.Tags = &v
}

// GetTasksStartPage returns the TasksStartPage field value if set, zero value otherwise.
func (o *InlineResponse20053Projects) GetTasksStartPage() string {
	if o == nil || o.TasksStartPage == nil {
		var ret string
		return ret
	}
	return *o.TasksStartPage
}

// GetTasksStartPageOk returns a tuple with the TasksStartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Projects) GetTasksStartPageOk() (*string, bool) {
	if o == nil || o.TasksStartPage == nil {
		return nil, false
	}
	return o.TasksStartPage, true
}

// HasTasksStartPage returns a boolean if a field has been set.
func (o *InlineResponse20053Projects) HasTasksStartPage() bool {
	if o != nil && o.TasksStartPage != nil {
		return true
	}

	return false
}

// SetTasksStartPage gets a reference to the given string and assigns it to the TasksStartPage field.
func (o *InlineResponse20053Projects) SetTasksStartPage(v string) {
	o.TasksStartPage = &v
}

func (o InlineResponse20053Projects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Announcement != nil {
		toSerialize["announcement"] = o.Announcement
	}
	if o.AnnouncementHTML != nil {
		toSerialize["announcementHTML"] = o.AnnouncementHTML
	}
	if o.BoardData != nil {
		toSerialize["boardData"] = o.BoardData
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.CreatedOn != nil {
		toSerialize["created-on"] = o.CreatedOn
	}
	if o.DefaultPrivacy != nil {
		toSerialize["defaultPrivacy"] = o.DefaultPrivacy
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.FilesAutoNewVersion != nil {
		toSerialize["filesAutoNewVersion"] = o.FilesAutoNewVersion
	}
	if o.HarvestTimersEnabled != nil {
		toSerialize["harvest-timers-enabled"] = o.HarvestTimersEnabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.IsProjectAdmin != nil {
		toSerialize["isProjectAdmin"] = o.IsProjectAdmin
	}
	if o.LastChangedOn != nil {
		toSerialize["last-changed-on"] = o.LastChangedOn
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notifyeveryone != nil {
		toSerialize["notifyeveryone"] = o.Notifyeveryone
	}
	if o.OverviewStartPage != nil {
		toSerialize["overview-start-page"] = o.OverviewStartPage
	}
	if o.PrivacyEnabled != nil {
		toSerialize["privacyEnabled"] = o.PrivacyEnabled
	}
	if o.ReplyByEmailEnabled != nil {
		toSerialize["replyByEmailEnabled"] = o.ReplyByEmailEnabled
	}
	if o.ShowAnnouncement != nil {
		toSerialize["show-announcement"] = o.ShowAnnouncement
	}
	if o.Starred != nil {
		toSerialize["starred"] = o.Starred
	}
	if o.StartPage != nil {
		toSerialize["start-page"] = o.StartPage
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SubStatus != nil {
		toSerialize["subStatus"] = o.SubStatus
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TasksStartPage != nil {
		toSerialize["tasks-start-page"] = o.TasksStartPage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20053Projects struct {
	value *InlineResponse20053Projects
	isSet bool
}

func (v NullableInlineResponse20053Projects) Get() *InlineResponse20053Projects {
	return v.value
}

func (v *NullableInlineResponse20053Projects) Set(val *InlineResponse20053Projects) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053Projects) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053Projects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053Projects(val *InlineResponse20053Projects) *NullableInlineResponse20053Projects {
	return &NullableInlineResponse20053Projects{value: val, isSet: true}
}

func (v NullableInlineResponse20053Projects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053Projects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


