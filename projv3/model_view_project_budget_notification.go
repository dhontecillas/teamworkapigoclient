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

// ViewProjectBudgetNotification ProjectBudgetNotification contains all the information returned from a notification.
type ViewProjectBudgetNotification struct {
	CapacityThreshold *float32 `json:"capacityThreshold,omitempty"`
	Companies *[]ViewRelationship `json:"companies,omitempty"`
	CompanyIds *[]int32 `json:"companyIds,omitempty"`
	Id *int32 `json:"id,omitempty"`
	NotificationMedium *string `json:"notificationMedium,omitempty"`
	TeamIds *[]int32 `json:"teamIds,omitempty"`
	Teams *[]ViewRelationship `json:"teams,omitempty"`
	UserIds *[]int32 `json:"userIds,omitempty"`
	Users *[]ViewRelationship `json:"users,omitempty"`
}

// NewViewProjectBudgetNotification instantiates a new ViewProjectBudgetNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProjectBudgetNotification() *ViewProjectBudgetNotification {
	this := ViewProjectBudgetNotification{}
	return &this
}

// NewViewProjectBudgetNotificationWithDefaults instantiates a new ViewProjectBudgetNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProjectBudgetNotificationWithDefaults() *ViewProjectBudgetNotification {
	this := ViewProjectBudgetNotification{}
	return &this
}

// GetCapacityThreshold returns the CapacityThreshold field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetCapacityThreshold() float32 {
	if o == nil || o.CapacityThreshold == nil {
		var ret float32
		return ret
	}
	return *o.CapacityThreshold
}

// GetCapacityThresholdOk returns a tuple with the CapacityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetCapacityThresholdOk() (*float32, bool) {
	if o == nil || o.CapacityThreshold == nil {
		return nil, false
	}
	return o.CapacityThreshold, true
}

// HasCapacityThreshold returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasCapacityThreshold() bool {
	if o != nil && o.CapacityThreshold != nil {
		return true
	}

	return false
}

// SetCapacityThreshold gets a reference to the given float32 and assigns it to the CapacityThreshold field.
func (o *ViewProjectBudgetNotification) SetCapacityThreshold(v float32) {
	o.CapacityThreshold = &v
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetCompanies() []ViewRelationship {
	if o == nil || o.Companies == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetCompaniesOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Companies == nil {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasCompanies() bool {
	if o != nil && o.Companies != nil {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given []ViewRelationship and assigns it to the Companies field.
func (o *ViewProjectBudgetNotification) SetCompanies(v []ViewRelationship) {
	o.Companies = &v
}

// GetCompanyIds returns the CompanyIds field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetCompanyIds() []int32 {
	if o == nil || o.CompanyIds == nil {
		var ret []int32
		return ret
	}
	return *o.CompanyIds
}

// GetCompanyIdsOk returns a tuple with the CompanyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetCompanyIdsOk() (*[]int32, bool) {
	if o == nil || o.CompanyIds == nil {
		return nil, false
	}
	return o.CompanyIds, true
}

// HasCompanyIds returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasCompanyIds() bool {
	if o != nil && o.CompanyIds != nil {
		return true
	}

	return false
}

// SetCompanyIds gets a reference to the given []int32 and assigns it to the CompanyIds field.
func (o *ViewProjectBudgetNotification) SetCompanyIds(v []int32) {
	o.CompanyIds = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewProjectBudgetNotification) SetId(v int32) {
	o.Id = &v
}

// GetNotificationMedium returns the NotificationMedium field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetNotificationMedium() string {
	if o == nil || o.NotificationMedium == nil {
		var ret string
		return ret
	}
	return *o.NotificationMedium
}

// GetNotificationMediumOk returns a tuple with the NotificationMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetNotificationMediumOk() (*string, bool) {
	if o == nil || o.NotificationMedium == nil {
		return nil, false
	}
	return o.NotificationMedium, true
}

// HasNotificationMedium returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasNotificationMedium() bool {
	if o != nil && o.NotificationMedium != nil {
		return true
	}

	return false
}

// SetNotificationMedium gets a reference to the given string and assigns it to the NotificationMedium field.
func (o *ViewProjectBudgetNotification) SetNotificationMedium(v string) {
	o.NotificationMedium = &v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetTeamIds() []int32 {
	if o == nil || o.TeamIds == nil {
		var ret []int32
		return ret
	}
	return *o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetTeamIdsOk() (*[]int32, bool) {
	if o == nil || o.TeamIds == nil {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasTeamIds() bool {
	if o != nil && o.TeamIds != nil {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []int32 and assigns it to the TeamIds field.
func (o *ViewProjectBudgetNotification) SetTeamIds(v []int32) {
	o.TeamIds = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetTeams() []ViewRelationship {
	if o == nil || o.Teams == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetTeamsOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []ViewRelationship and assigns it to the Teams field.
func (o *ViewProjectBudgetNotification) SetTeams(v []ViewRelationship) {
	o.Teams = &v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetUserIds() []int32 {
	if o == nil || o.UserIds == nil {
		var ret []int32
		return ret
	}
	return *o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetUserIdsOk() (*[]int32, bool) {
	if o == nil || o.UserIds == nil {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasUserIds() bool {
	if o != nil && o.UserIds != nil {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []int32 and assigns it to the UserIds field.
func (o *ViewProjectBudgetNotification) SetUserIds(v []int32) {
	o.UserIds = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ViewProjectBudgetNotification) GetUsers() []ViewRelationship {
	if o == nil || o.Users == nil {
		var ret []ViewRelationship
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectBudgetNotification) GetUsersOk() (*[]ViewRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ViewProjectBudgetNotification) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ViewRelationship and assigns it to the Users field.
func (o *ViewProjectBudgetNotification) SetUsers(v []ViewRelationship) {
	o.Users = &v
}

func (o ViewProjectBudgetNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityThreshold != nil {
		toSerialize["capacityThreshold"] = o.CapacityThreshold
	}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	if o.CompanyIds != nil {
		toSerialize["companyIds"] = o.CompanyIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NotificationMedium != nil {
		toSerialize["notificationMedium"] = o.NotificationMedium
	}
	if o.TeamIds != nil {
		toSerialize["teamIds"] = o.TeamIds
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableViewProjectBudgetNotification struct {
	value *ViewProjectBudgetNotification
	isSet bool
}

func (v NullableViewProjectBudgetNotification) Get() *ViewProjectBudgetNotification {
	return v.value
}

func (v *NullableViewProjectBudgetNotification) Set(val *ViewProjectBudgetNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProjectBudgetNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProjectBudgetNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProjectBudgetNotification(val *ViewProjectBudgetNotification) *NullableViewProjectBudgetNotification {
	return &NullableViewProjectBudgetNotification{value: val, isSet: true}
}

func (v NullableViewProjectBudgetNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProjectBudgetNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


