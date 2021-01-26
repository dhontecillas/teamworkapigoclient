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

// ViewUser User contains all the information returned from an user.
type ViewUser struct {
	AvatarUrl *string `json:"avatarUrl,omitempty"`
	Company *ViewRelationship `json:"company,omitempty"`
	CompanyId *int32 `json:"companyId,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	LengthOfDay *float32 `json:"lengthOfDay,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	WorkingHour *ViewRelationship `json:"workingHour,omitempty"`
	WorkingHoursId *int32 `json:"workingHoursId,omitempty"`
}

// NewViewUser instantiates a new ViewUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewUser() *ViewUser {
	this := ViewUser{}
	return &this
}

// NewViewUserWithDefaults instantiates a new ViewUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewUserWithDefaults() *ViewUser {
	this := ViewUser{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *ViewUser) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *ViewUser) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *ViewUser) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ViewUser) GetCompany() ViewRelationship {
	if o == nil || o.Company == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetCompanyOk() (*ViewRelationship, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ViewUser) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given ViewRelationship and assigns it to the Company field.
func (o *ViewUser) SetCompany(v ViewRelationship) {
	o.Company = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ViewUser) GetCompanyId() int32 {
	if o == nil || o.CompanyId == nil {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetCompanyIdOk() (*int32, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ViewUser) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *ViewUser) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ViewUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ViewUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ViewUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ViewUser) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ViewUser) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ViewUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewUser) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewUser) SetId(v int32) {
	o.Id = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *ViewUser) GetIsAdmin() bool {
	if o == nil || o.IsAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetIsAdminOk() (*bool, bool) {
	if o == nil || o.IsAdmin == nil {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *ViewUser) HasIsAdmin() bool {
	if o != nil && o.IsAdmin != nil {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *ViewUser) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ViewUser) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ViewUser) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ViewUser) SetLastName(v string) {
	o.LastName = &v
}

// GetLengthOfDay returns the LengthOfDay field value if set, zero value otherwise.
func (o *ViewUser) GetLengthOfDay() float32 {
	if o == nil || o.LengthOfDay == nil {
		var ret float32
		return ret
	}
	return *o.LengthOfDay
}

// GetLengthOfDayOk returns a tuple with the LengthOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetLengthOfDayOk() (*float32, bool) {
	if o == nil || o.LengthOfDay == nil {
		return nil, false
	}
	return o.LengthOfDay, true
}

// HasLengthOfDay returns a boolean if a field has been set.
func (o *ViewUser) HasLengthOfDay() bool {
	if o != nil && o.LengthOfDay != nil {
		return true
	}

	return false
}

// SetLengthOfDay gets a reference to the given float32 and assigns it to the LengthOfDay field.
func (o *ViewUser) SetLengthOfDay(v float32) {
	o.LengthOfDay = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewUser) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewUser) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewUser) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ViewUser) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ViewUser) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ViewUser) SetType(v string) {
	o.Type = &v
}

// GetWorkingHour returns the WorkingHour field value if set, zero value otherwise.
func (o *ViewUser) GetWorkingHour() ViewRelationship {
	if o == nil || o.WorkingHour == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.WorkingHour
}

// GetWorkingHourOk returns a tuple with the WorkingHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetWorkingHourOk() (*ViewRelationship, bool) {
	if o == nil || o.WorkingHour == nil {
		return nil, false
	}
	return o.WorkingHour, true
}

// HasWorkingHour returns a boolean if a field has been set.
func (o *ViewUser) HasWorkingHour() bool {
	if o != nil && o.WorkingHour != nil {
		return true
	}

	return false
}

// SetWorkingHour gets a reference to the given ViewRelationship and assigns it to the WorkingHour field.
func (o *ViewUser) SetWorkingHour(v ViewRelationship) {
	o.WorkingHour = &v
}

// GetWorkingHoursId returns the WorkingHoursId field value if set, zero value otherwise.
func (o *ViewUser) GetWorkingHoursId() int32 {
	if o == nil || o.WorkingHoursId == nil {
		var ret int32
		return ret
	}
	return *o.WorkingHoursId
}

// GetWorkingHoursIdOk returns a tuple with the WorkingHoursId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUser) GetWorkingHoursIdOk() (*int32, bool) {
	if o == nil || o.WorkingHoursId == nil {
		return nil, false
	}
	return o.WorkingHoursId, true
}

// HasWorkingHoursId returns a boolean if a field has been set.
func (o *ViewUser) HasWorkingHoursId() bool {
	if o != nil && o.WorkingHoursId != nil {
		return true
	}

	return false
}

// SetWorkingHoursId gets a reference to the given int32 and assigns it to the WorkingHoursId field.
func (o *ViewUser) SetWorkingHoursId(v int32) {
	o.WorkingHoursId = &v
}

func (o ViewUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarUrl != nil {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAdmin != nil {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.LengthOfDay != nil {
		toSerialize["lengthOfDay"] = o.LengthOfDay
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WorkingHour != nil {
		toSerialize["workingHour"] = o.WorkingHour
	}
	if o.WorkingHoursId != nil {
		toSerialize["workingHoursId"] = o.WorkingHoursId
	}
	return json.Marshal(toSerialize)
}

type NullableViewUser struct {
	value *ViewUser
	isSet bool
}

func (v NullableViewUser) Get() *ViewUser {
	return v.value
}

func (v *NullableViewUser) Set(val *ViewUser) {
	v.value = val
	v.isSet = true
}

func (v NullableViewUser) IsSet() bool {
	return v.isSet
}

func (v *NullableViewUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewUser(val *ViewUser) *NullableViewUser {
	return &NullableViewUser{value: val, isSet: true}
}

func (v NullableViewUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


