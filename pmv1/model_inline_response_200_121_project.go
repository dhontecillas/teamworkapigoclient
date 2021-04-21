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

// InlineResponse200121Project struct for InlineResponse200121Project
type InlineResponse200121Project struct {
	ActivePages *InlineResponse20060ProjectActivePages `json:"active-pages,omitempty"`
	Announcement *string `json:"announcement,omitempty"`
	AnnouncementHTML *string `json:"announcementHTML,omitempty"`
	BoardData *InlineResponse200121ProjectBoardData `json:"boardData,omitempty"`
	Category *InlineResponse2002Column `json:"category,omitempty"`
	Company *InlineResponse2002People12345Company `json:"company,omitempty"`
	CreatedOn *string `json:"created-on,omitempty"`
	DefaultPrivacy *string `json:"defaultPrivacy,omitempty"`
	Description *string `json:"description,omitempty"`
	DirectFileUploadsEnabled *bool `json:"directFileUploadsEnabled,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	FilesAutoNewVersion *bool `json:"filesAutoNewVersion,omitempty"`
	HarvestTimersEnabled *bool `json:"harvest-timers-enabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Integrations *InlineResponse20060ProjectIntegrations `json:"integrations,omitempty"`
	LastChangedOn *string `json:"last-changed-on,omitempty"`
	Logo *string `json:"logo,omitempty"`
	LogoFromCompany *bool `json:"logoFromCompany,omitempty"`
	Name *string `json:"name,omitempty"`
	Notifyeveryone *bool `json:"notifyeveryone,omitempty"`
	OverviewStartPage *string `json:"overview-start-page,omitempty"`
	Owner *InlineResponse20045CardOwner `json:"owner,omitempty"`
	PrivacyEnabled *bool `json:"privacyEnabled,omitempty"`
	ReplyByEmailEnabled *bool `json:"replyByEmailEnabled,omitempty"`
	ShowAnnouncement *bool `json:"show-announcement,omitempty"`
	SkipWeekends *bool `json:"skipWeekends,omitempty"`
	Starred *bool `json:"starred,omitempty"`
	StartPage *string `json:"start-page,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	Status *string `json:"status,omitempty"`
	SubStatus *string `json:"subStatus,omitempty"`
	Tags *[]map[string]interface{} `json:"tags,omitempty"`
	TasksStartPage *string `json:"tasks-start-page,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse200121Project instantiates a new InlineResponse200121Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121Project() *InlineResponse200121Project {
	this := InlineResponse200121Project{}
	return &this
}

// NewInlineResponse200121ProjectWithDefaults instantiates a new InlineResponse200121Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121ProjectWithDefaults() *InlineResponse200121Project {
	this := InlineResponse200121Project{}
	return &this
}

// GetActivePages returns the ActivePages field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetActivePages() InlineResponse20060ProjectActivePages {
	if o == nil || o.ActivePages == nil {
		var ret InlineResponse20060ProjectActivePages
		return ret
	}
	return *o.ActivePages
}

// GetActivePagesOk returns a tuple with the ActivePages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetActivePagesOk() (*InlineResponse20060ProjectActivePages, bool) {
	if o == nil || o.ActivePages == nil {
		return nil, false
	}
	return o.ActivePages, true
}

// HasActivePages returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasActivePages() bool {
	if o != nil && o.ActivePages != nil {
		return true
	}

	return false
}

// SetActivePages gets a reference to the given InlineResponse20060ProjectActivePages and assigns it to the ActivePages field.
func (o *InlineResponse200121Project) SetActivePages(v InlineResponse20060ProjectActivePages) {
	o.ActivePages = &v
}

// GetAnnouncement returns the Announcement field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetAnnouncement() string {
	if o == nil || o.Announcement == nil {
		var ret string
		return ret
	}
	return *o.Announcement
}

// GetAnnouncementOk returns a tuple with the Announcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetAnnouncementOk() (*string, bool) {
	if o == nil || o.Announcement == nil {
		return nil, false
	}
	return o.Announcement, true
}

// HasAnnouncement returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasAnnouncement() bool {
	if o != nil && o.Announcement != nil {
		return true
	}

	return false
}

// SetAnnouncement gets a reference to the given string and assigns it to the Announcement field.
func (o *InlineResponse200121Project) SetAnnouncement(v string) {
	o.Announcement = &v
}

// GetAnnouncementHTML returns the AnnouncementHTML field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetAnnouncementHTML() string {
	if o == nil || o.AnnouncementHTML == nil {
		var ret string
		return ret
	}
	return *o.AnnouncementHTML
}

// GetAnnouncementHTMLOk returns a tuple with the AnnouncementHTML field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetAnnouncementHTMLOk() (*string, bool) {
	if o == nil || o.AnnouncementHTML == nil {
		return nil, false
	}
	return o.AnnouncementHTML, true
}

// HasAnnouncementHTML returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasAnnouncementHTML() bool {
	if o != nil && o.AnnouncementHTML != nil {
		return true
	}

	return false
}

// SetAnnouncementHTML gets a reference to the given string and assigns it to the AnnouncementHTML field.
func (o *InlineResponse200121Project) SetAnnouncementHTML(v string) {
	o.AnnouncementHTML = &v
}

// GetBoardData returns the BoardData field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetBoardData() InlineResponse200121ProjectBoardData {
	if o == nil || o.BoardData == nil {
		var ret InlineResponse200121ProjectBoardData
		return ret
	}
	return *o.BoardData
}

// GetBoardDataOk returns a tuple with the BoardData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetBoardDataOk() (*InlineResponse200121ProjectBoardData, bool) {
	if o == nil || o.BoardData == nil {
		return nil, false
	}
	return o.BoardData, true
}

// HasBoardData returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasBoardData() bool {
	if o != nil && o.BoardData != nil {
		return true
	}

	return false
}

// SetBoardData gets a reference to the given InlineResponse200121ProjectBoardData and assigns it to the BoardData field.
func (o *InlineResponse200121Project) SetBoardData(v InlineResponse200121ProjectBoardData) {
	o.BoardData = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetCategory() InlineResponse2002Column {
	if o == nil || o.Category == nil {
		var ret InlineResponse2002Column
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetCategoryOk() (*InlineResponse2002Column, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given InlineResponse2002Column and assigns it to the Category field.
func (o *InlineResponse200121Project) SetCategory(v InlineResponse2002Column) {
	o.Category = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetCompany() InlineResponse2002People12345Company {
	if o == nil || o.Company == nil {
		var ret InlineResponse2002People12345Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetCompanyOk() (*InlineResponse2002People12345Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given InlineResponse2002People12345Company and assigns it to the Company field.
func (o *InlineResponse200121Project) SetCompany(v InlineResponse2002People12345Company) {
	o.Company = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *InlineResponse200121Project) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetDefaultPrivacy returns the DefaultPrivacy field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetDefaultPrivacy() string {
	if o == nil || o.DefaultPrivacy == nil {
		var ret string
		return ret
	}
	return *o.DefaultPrivacy
}

// GetDefaultPrivacyOk returns a tuple with the DefaultPrivacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetDefaultPrivacyOk() (*string, bool) {
	if o == nil || o.DefaultPrivacy == nil {
		return nil, false
	}
	return o.DefaultPrivacy, true
}

// HasDefaultPrivacy returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasDefaultPrivacy() bool {
	if o != nil && o.DefaultPrivacy != nil {
		return true
	}

	return false
}

// SetDefaultPrivacy gets a reference to the given string and assigns it to the DefaultPrivacy field.
func (o *InlineResponse200121Project) SetDefaultPrivacy(v string) {
	o.DefaultPrivacy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200121Project) SetDescription(v string) {
	o.Description = &v
}

// GetDirectFileUploadsEnabled returns the DirectFileUploadsEnabled field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetDirectFileUploadsEnabled() bool {
	if o == nil || o.DirectFileUploadsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DirectFileUploadsEnabled
}

// GetDirectFileUploadsEnabledOk returns a tuple with the DirectFileUploadsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetDirectFileUploadsEnabledOk() (*bool, bool) {
	if o == nil || o.DirectFileUploadsEnabled == nil {
		return nil, false
	}
	return o.DirectFileUploadsEnabled, true
}

// HasDirectFileUploadsEnabled returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasDirectFileUploadsEnabled() bool {
	if o != nil && o.DirectFileUploadsEnabled != nil {
		return true
	}

	return false
}

// SetDirectFileUploadsEnabled gets a reference to the given bool and assigns it to the DirectFileUploadsEnabled field.
func (o *InlineResponse200121Project) SetDirectFileUploadsEnabled(v bool) {
	o.DirectFileUploadsEnabled = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *InlineResponse200121Project) SetEndDate(v string) {
	o.EndDate = &v
}

// GetFilesAutoNewVersion returns the FilesAutoNewVersion field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetFilesAutoNewVersion() bool {
	if o == nil || o.FilesAutoNewVersion == nil {
		var ret bool
		return ret
	}
	return *o.FilesAutoNewVersion
}

// GetFilesAutoNewVersionOk returns a tuple with the FilesAutoNewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetFilesAutoNewVersionOk() (*bool, bool) {
	if o == nil || o.FilesAutoNewVersion == nil {
		return nil, false
	}
	return o.FilesAutoNewVersion, true
}

// HasFilesAutoNewVersion returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasFilesAutoNewVersion() bool {
	if o != nil && o.FilesAutoNewVersion != nil {
		return true
	}

	return false
}

// SetFilesAutoNewVersion gets a reference to the given bool and assigns it to the FilesAutoNewVersion field.
func (o *InlineResponse200121Project) SetFilesAutoNewVersion(v bool) {
	o.FilesAutoNewVersion = &v
}

// GetHarvestTimersEnabled returns the HarvestTimersEnabled field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetHarvestTimersEnabled() bool {
	if o == nil || o.HarvestTimersEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HarvestTimersEnabled
}

// GetHarvestTimersEnabledOk returns a tuple with the HarvestTimersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetHarvestTimersEnabledOk() (*bool, bool) {
	if o == nil || o.HarvestTimersEnabled == nil {
		return nil, false
	}
	return o.HarvestTimersEnabled, true
}

// HasHarvestTimersEnabled returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasHarvestTimersEnabled() bool {
	if o != nil && o.HarvestTimersEnabled != nil {
		return true
	}

	return false
}

// SetHarvestTimersEnabled gets a reference to the given bool and assigns it to the HarvestTimersEnabled field.
func (o *InlineResponse200121Project) SetHarvestTimersEnabled(v bool) {
	o.HarvestTimersEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200121Project) SetId(v string) {
	o.Id = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetIntegrations() InlineResponse20060ProjectIntegrations {
	if o == nil || o.Integrations == nil {
		var ret InlineResponse20060ProjectIntegrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetIntegrationsOk() (*InlineResponse20060ProjectIntegrations, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasIntegrations() bool {
	if o != nil && o.Integrations != nil {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given InlineResponse20060ProjectIntegrations and assigns it to the Integrations field.
func (o *InlineResponse200121Project) SetIntegrations(v InlineResponse20060ProjectIntegrations) {
	o.Integrations = &v
}

// GetLastChangedOn returns the LastChangedOn field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetLastChangedOn() string {
	if o == nil || o.LastChangedOn == nil {
		var ret string
		return ret
	}
	return *o.LastChangedOn
}

// GetLastChangedOnOk returns a tuple with the LastChangedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetLastChangedOnOk() (*string, bool) {
	if o == nil || o.LastChangedOn == nil {
		return nil, false
	}
	return o.LastChangedOn, true
}

// HasLastChangedOn returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasLastChangedOn() bool {
	if o != nil && o.LastChangedOn != nil {
		return true
	}

	return false
}

// SetLastChangedOn gets a reference to the given string and assigns it to the LastChangedOn field.
func (o *InlineResponse200121Project) SetLastChangedOn(v string) {
	o.LastChangedOn = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *InlineResponse200121Project) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoFromCompany returns the LogoFromCompany field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetLogoFromCompany() bool {
	if o == nil || o.LogoFromCompany == nil {
		var ret bool
		return ret
	}
	return *o.LogoFromCompany
}

// GetLogoFromCompanyOk returns a tuple with the LogoFromCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetLogoFromCompanyOk() (*bool, bool) {
	if o == nil || o.LogoFromCompany == nil {
		return nil, false
	}
	return o.LogoFromCompany, true
}

// HasLogoFromCompany returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasLogoFromCompany() bool {
	if o != nil && o.LogoFromCompany != nil {
		return true
	}

	return false
}

// SetLogoFromCompany gets a reference to the given bool and assigns it to the LogoFromCompany field.
func (o *InlineResponse200121Project) SetLogoFromCompany(v bool) {
	o.LogoFromCompany = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200121Project) SetName(v string) {
	o.Name = &v
}

// GetNotifyeveryone returns the Notifyeveryone field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetNotifyeveryone() bool {
	if o == nil || o.Notifyeveryone == nil {
		var ret bool
		return ret
	}
	return *o.Notifyeveryone
}

// GetNotifyeveryoneOk returns a tuple with the Notifyeveryone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetNotifyeveryoneOk() (*bool, bool) {
	if o == nil || o.Notifyeveryone == nil {
		return nil, false
	}
	return o.Notifyeveryone, true
}

// HasNotifyeveryone returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasNotifyeveryone() bool {
	if o != nil && o.Notifyeveryone != nil {
		return true
	}

	return false
}

// SetNotifyeveryone gets a reference to the given bool and assigns it to the Notifyeveryone field.
func (o *InlineResponse200121Project) SetNotifyeveryone(v bool) {
	o.Notifyeveryone = &v
}

// GetOverviewStartPage returns the OverviewStartPage field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetOverviewStartPage() string {
	if o == nil || o.OverviewStartPage == nil {
		var ret string
		return ret
	}
	return *o.OverviewStartPage
}

// GetOverviewStartPageOk returns a tuple with the OverviewStartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetOverviewStartPageOk() (*string, bool) {
	if o == nil || o.OverviewStartPage == nil {
		return nil, false
	}
	return o.OverviewStartPage, true
}

// HasOverviewStartPage returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasOverviewStartPage() bool {
	if o != nil && o.OverviewStartPage != nil {
		return true
	}

	return false
}

// SetOverviewStartPage gets a reference to the given string and assigns it to the OverviewStartPage field.
func (o *InlineResponse200121Project) SetOverviewStartPage(v string) {
	o.OverviewStartPage = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetOwner() InlineResponse20045CardOwner {
	if o == nil || o.Owner == nil {
		var ret InlineResponse20045CardOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetOwnerOk() (*InlineResponse20045CardOwner, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given InlineResponse20045CardOwner and assigns it to the Owner field.
func (o *InlineResponse200121Project) SetOwner(v InlineResponse20045CardOwner) {
	o.Owner = &v
}

// GetPrivacyEnabled returns the PrivacyEnabled field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetPrivacyEnabled() bool {
	if o == nil || o.PrivacyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PrivacyEnabled
}

// GetPrivacyEnabledOk returns a tuple with the PrivacyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetPrivacyEnabledOk() (*bool, bool) {
	if o == nil || o.PrivacyEnabled == nil {
		return nil, false
	}
	return o.PrivacyEnabled, true
}

// HasPrivacyEnabled returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasPrivacyEnabled() bool {
	if o != nil && o.PrivacyEnabled != nil {
		return true
	}

	return false
}

// SetPrivacyEnabled gets a reference to the given bool and assigns it to the PrivacyEnabled field.
func (o *InlineResponse200121Project) SetPrivacyEnabled(v bool) {
	o.PrivacyEnabled = &v
}

// GetReplyByEmailEnabled returns the ReplyByEmailEnabled field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetReplyByEmailEnabled() bool {
	if o == nil || o.ReplyByEmailEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplyByEmailEnabled
}

// GetReplyByEmailEnabledOk returns a tuple with the ReplyByEmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetReplyByEmailEnabledOk() (*bool, bool) {
	if o == nil || o.ReplyByEmailEnabled == nil {
		return nil, false
	}
	return o.ReplyByEmailEnabled, true
}

// HasReplyByEmailEnabled returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasReplyByEmailEnabled() bool {
	if o != nil && o.ReplyByEmailEnabled != nil {
		return true
	}

	return false
}

// SetReplyByEmailEnabled gets a reference to the given bool and assigns it to the ReplyByEmailEnabled field.
func (o *InlineResponse200121Project) SetReplyByEmailEnabled(v bool) {
	o.ReplyByEmailEnabled = &v
}

// GetShowAnnouncement returns the ShowAnnouncement field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetShowAnnouncement() bool {
	if o == nil || o.ShowAnnouncement == nil {
		var ret bool
		return ret
	}
	return *o.ShowAnnouncement
}

// GetShowAnnouncementOk returns a tuple with the ShowAnnouncement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetShowAnnouncementOk() (*bool, bool) {
	if o == nil || o.ShowAnnouncement == nil {
		return nil, false
	}
	return o.ShowAnnouncement, true
}

// HasShowAnnouncement returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasShowAnnouncement() bool {
	if o != nil && o.ShowAnnouncement != nil {
		return true
	}

	return false
}

// SetShowAnnouncement gets a reference to the given bool and assigns it to the ShowAnnouncement field.
func (o *InlineResponse200121Project) SetShowAnnouncement(v bool) {
	o.ShowAnnouncement = &v
}

// GetSkipWeekends returns the SkipWeekends field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetSkipWeekends() bool {
	if o == nil || o.SkipWeekends == nil {
		var ret bool
		return ret
	}
	return *o.SkipWeekends
}

// GetSkipWeekendsOk returns a tuple with the SkipWeekends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetSkipWeekendsOk() (*bool, bool) {
	if o == nil || o.SkipWeekends == nil {
		return nil, false
	}
	return o.SkipWeekends, true
}

// HasSkipWeekends returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasSkipWeekends() bool {
	if o != nil && o.SkipWeekends != nil {
		return true
	}

	return false
}

// SetSkipWeekends gets a reference to the given bool and assigns it to the SkipWeekends field.
func (o *InlineResponse200121Project) SetSkipWeekends(v bool) {
	o.SkipWeekends = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetStarred() bool {
	if o == nil || o.Starred == nil {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetStarredOk() (*bool, bool) {
	if o == nil || o.Starred == nil {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasStarred() bool {
	if o != nil && o.Starred != nil {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *InlineResponse200121Project) SetStarred(v bool) {
	o.Starred = &v
}

// GetStartPage returns the StartPage field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetStartPage() string {
	if o == nil || o.StartPage == nil {
		var ret string
		return ret
	}
	return *o.StartPage
}

// GetStartPageOk returns a tuple with the StartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetStartPageOk() (*string, bool) {
	if o == nil || o.StartPage == nil {
		return nil, false
	}
	return o.StartPage, true
}

// HasStartPage returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasStartPage() bool {
	if o != nil && o.StartPage != nil {
		return true
	}

	return false
}

// SetStartPage gets a reference to the given string and assigns it to the StartPage field.
func (o *InlineResponse200121Project) SetStartPage(v string) {
	o.StartPage = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *InlineResponse200121Project) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200121Project) SetStatus(v string) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetSubStatus() string {
	if o == nil || o.SubStatus == nil {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetSubStatusOk() (*string, bool) {
	if o == nil || o.SubStatus == nil {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasSubStatus() bool {
	if o != nil && o.SubStatus != nil {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *InlineResponse200121Project) SetSubStatus(v string) {
	o.SubStatus = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetTags() []map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetTagsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []map[string]interface{} and assigns it to the Tags field.
func (o *InlineResponse200121Project) SetTags(v []map[string]interface{}) {
	o.Tags = &v
}

// GetTasksStartPage returns the TasksStartPage field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetTasksStartPage() string {
	if o == nil || o.TasksStartPage == nil {
		var ret string
		return ret
	}
	return *o.TasksStartPage
}

// GetTasksStartPageOk returns a tuple with the TasksStartPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetTasksStartPageOk() (*string, bool) {
	if o == nil || o.TasksStartPage == nil {
		return nil, false
	}
	return o.TasksStartPage, true
}

// HasTasksStartPage returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasTasksStartPage() bool {
	if o != nil && o.TasksStartPage != nil {
		return true
	}

	return false
}

// SetTasksStartPage gets a reference to the given string and assigns it to the TasksStartPage field.
func (o *InlineResponse200121Project) SetTasksStartPage(v string) {
	o.TasksStartPage = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200121Project) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Project) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200121Project) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200121Project) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse200121Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivePages != nil {
		toSerialize["active-pages"] = o.ActivePages
	}
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DirectFileUploadsEnabled != nil {
		toSerialize["directFileUploadsEnabled"] = o.DirectFileUploadsEnabled
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
	if o.LastChangedOn != nil {
		toSerialize["last-changed-on"] = o.LastChangedOn
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.LogoFromCompany != nil {
		toSerialize["logoFromCompany"] = o.LogoFromCompany
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
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
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
	if o.SkipWeekends != nil {
		toSerialize["skipWeekends"] = o.SkipWeekends
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200121Project struct {
	value *InlineResponse200121Project
	isSet bool
}

func (v NullableInlineResponse200121Project) Get() *InlineResponse200121Project {
	return v.value
}

func (v *NullableInlineResponse200121Project) Set(val *InlineResponse200121Project) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121Project) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121Project) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121Project(val *InlineResponse200121Project) *NullableInlineResponse200121Project {
	return &NullableInlineResponse200121Project{value: val, isSet: true}
}

func (v NullableInlineResponse200121Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121Project) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


