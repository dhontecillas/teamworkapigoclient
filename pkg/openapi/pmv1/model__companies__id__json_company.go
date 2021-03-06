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

// CompaniesIdJsonCompany struct for CompaniesIdJsonCompany
type CompaniesIdJsonCompany struct {
	AddressOne *string `json:"address_one,omitempty"`
	AddressTwo *string `json:"address_two,omitempty"`
	Cid *string `json:"cid,omitempty"`
	City *string `json:"city,omitempty"`
	Countrycode *string `json:"countrycode,omitempty"`
	EmailOne *string `json:"email_one,omitempty"`
	EmailThree *string `json:"email_three,omitempty"`
	EmailTwo *string `json:"email_two,omitempty"`
	Fax *string `json:"fax,omitempty"`
	IndustryCatId *string `json:"industryCatId,omitempty"`
	LogoPendingFileRef *string `json:"logoPendingFileRef,omitempty"`
	Name *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	PrivateNotes *string `json:"privateNotes,omitempty"`
	Profile *string `json:"profile,omitempty"`
	RemoveLogo *bool `json:"removeLogo,omitempty"`
	State *string `json:"state,omitempty"`
	Tags *[]map[string]interface{} `json:"tags,omitempty"`
	Website *string `json:"website,omitempty"`
	Zip *string `json:"zip,omitempty"`
}

// NewCompaniesIdJsonCompany instantiates a new CompaniesIdJsonCompany object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompaniesIdJsonCompany() *CompaniesIdJsonCompany {
	this := CompaniesIdJsonCompany{}
	return &this
}

// NewCompaniesIdJsonCompanyWithDefaults instantiates a new CompaniesIdJsonCompany object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompaniesIdJsonCompanyWithDefaults() *CompaniesIdJsonCompany {
	this := CompaniesIdJsonCompany{}
	return &this
}

// GetAddressOne returns the AddressOne field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetAddressOne() string {
	if o == nil || o.AddressOne == nil {
		var ret string
		return ret
	}
	return *o.AddressOne
}

// GetAddressOneOk returns a tuple with the AddressOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetAddressOneOk() (*string, bool) {
	if o == nil || o.AddressOne == nil {
		return nil, false
	}
	return o.AddressOne, true
}

// HasAddressOne returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasAddressOne() bool {
	if o != nil && o.AddressOne != nil {
		return true
	}

	return false
}

// SetAddressOne gets a reference to the given string and assigns it to the AddressOne field.
func (o *CompaniesIdJsonCompany) SetAddressOne(v string) {
	o.AddressOne = &v
}

// GetAddressTwo returns the AddressTwo field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetAddressTwo() string {
	if o == nil || o.AddressTwo == nil {
		var ret string
		return ret
	}
	return *o.AddressTwo
}

// GetAddressTwoOk returns a tuple with the AddressTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetAddressTwoOk() (*string, bool) {
	if o == nil || o.AddressTwo == nil {
		return nil, false
	}
	return o.AddressTwo, true
}

// HasAddressTwo returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasAddressTwo() bool {
	if o != nil && o.AddressTwo != nil {
		return true
	}

	return false
}

// SetAddressTwo gets a reference to the given string and assigns it to the AddressTwo field.
func (o *CompaniesIdJsonCompany) SetAddressTwo(v string) {
	o.AddressTwo = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetCid() string {
	if o == nil || o.Cid == nil {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetCidOk() (*string, bool) {
	if o == nil || o.Cid == nil {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasCid() bool {
	if o != nil && o.Cid != nil {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *CompaniesIdJsonCompany) SetCid(v string) {
	o.Cid = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CompaniesIdJsonCompany) SetCity(v string) {
	o.City = &v
}

// GetCountrycode returns the Countrycode field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetCountrycode() string {
	if o == nil || o.Countrycode == nil {
		var ret string
		return ret
	}
	return *o.Countrycode
}

// GetCountrycodeOk returns a tuple with the Countrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetCountrycodeOk() (*string, bool) {
	if o == nil || o.Countrycode == nil {
		return nil, false
	}
	return o.Countrycode, true
}

// HasCountrycode returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasCountrycode() bool {
	if o != nil && o.Countrycode != nil {
		return true
	}

	return false
}

// SetCountrycode gets a reference to the given string and assigns it to the Countrycode field.
func (o *CompaniesIdJsonCompany) SetCountrycode(v string) {
	o.Countrycode = &v
}

// GetEmailOne returns the EmailOne field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetEmailOne() string {
	if o == nil || o.EmailOne == nil {
		var ret string
		return ret
	}
	return *o.EmailOne
}

// GetEmailOneOk returns a tuple with the EmailOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetEmailOneOk() (*string, bool) {
	if o == nil || o.EmailOne == nil {
		return nil, false
	}
	return o.EmailOne, true
}

// HasEmailOne returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasEmailOne() bool {
	if o != nil && o.EmailOne != nil {
		return true
	}

	return false
}

// SetEmailOne gets a reference to the given string and assigns it to the EmailOne field.
func (o *CompaniesIdJsonCompany) SetEmailOne(v string) {
	o.EmailOne = &v
}

// GetEmailThree returns the EmailThree field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetEmailThree() string {
	if o == nil || o.EmailThree == nil {
		var ret string
		return ret
	}
	return *o.EmailThree
}

// GetEmailThreeOk returns a tuple with the EmailThree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetEmailThreeOk() (*string, bool) {
	if o == nil || o.EmailThree == nil {
		return nil, false
	}
	return o.EmailThree, true
}

// HasEmailThree returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasEmailThree() bool {
	if o != nil && o.EmailThree != nil {
		return true
	}

	return false
}

// SetEmailThree gets a reference to the given string and assigns it to the EmailThree field.
func (o *CompaniesIdJsonCompany) SetEmailThree(v string) {
	o.EmailThree = &v
}

// GetEmailTwo returns the EmailTwo field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetEmailTwo() string {
	if o == nil || o.EmailTwo == nil {
		var ret string
		return ret
	}
	return *o.EmailTwo
}

// GetEmailTwoOk returns a tuple with the EmailTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetEmailTwoOk() (*string, bool) {
	if o == nil || o.EmailTwo == nil {
		return nil, false
	}
	return o.EmailTwo, true
}

// HasEmailTwo returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasEmailTwo() bool {
	if o != nil && o.EmailTwo != nil {
		return true
	}

	return false
}

// SetEmailTwo gets a reference to the given string and assigns it to the EmailTwo field.
func (o *CompaniesIdJsonCompany) SetEmailTwo(v string) {
	o.EmailTwo = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetFaxOk() (*string, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *CompaniesIdJsonCompany) SetFax(v string) {
	o.Fax = &v
}

// GetIndustryCatId returns the IndustryCatId field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetIndustryCatId() string {
	if o == nil || o.IndustryCatId == nil {
		var ret string
		return ret
	}
	return *o.IndustryCatId
}

// GetIndustryCatIdOk returns a tuple with the IndustryCatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetIndustryCatIdOk() (*string, bool) {
	if o == nil || o.IndustryCatId == nil {
		return nil, false
	}
	return o.IndustryCatId, true
}

// HasIndustryCatId returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasIndustryCatId() bool {
	if o != nil && o.IndustryCatId != nil {
		return true
	}

	return false
}

// SetIndustryCatId gets a reference to the given string and assigns it to the IndustryCatId field.
func (o *CompaniesIdJsonCompany) SetIndustryCatId(v string) {
	o.IndustryCatId = &v
}

// GetLogoPendingFileRef returns the LogoPendingFileRef field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetLogoPendingFileRef() string {
	if o == nil || o.LogoPendingFileRef == nil {
		var ret string
		return ret
	}
	return *o.LogoPendingFileRef
}

// GetLogoPendingFileRefOk returns a tuple with the LogoPendingFileRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetLogoPendingFileRefOk() (*string, bool) {
	if o == nil || o.LogoPendingFileRef == nil {
		return nil, false
	}
	return o.LogoPendingFileRef, true
}

// HasLogoPendingFileRef returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasLogoPendingFileRef() bool {
	if o != nil && o.LogoPendingFileRef != nil {
		return true
	}

	return false
}

// SetLogoPendingFileRef gets a reference to the given string and assigns it to the LogoPendingFileRef field.
func (o *CompaniesIdJsonCompany) SetLogoPendingFileRef(v string) {
	o.LogoPendingFileRef = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompaniesIdJsonCompany) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CompaniesIdJsonCompany) SetPhone(v string) {
	o.Phone = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetPrivateNotes() string {
	if o == nil || o.PrivateNotes == nil {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetPrivateNotesOk() (*string, bool) {
	if o == nil || o.PrivateNotes == nil {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasPrivateNotes() bool {
	if o != nil && o.PrivateNotes != nil {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *CompaniesIdJsonCompany) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *CompaniesIdJsonCompany) SetProfile(v string) {
	o.Profile = &v
}

// GetRemoveLogo returns the RemoveLogo field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetRemoveLogo() bool {
	if o == nil || o.RemoveLogo == nil {
		var ret bool
		return ret
	}
	return *o.RemoveLogo
}

// GetRemoveLogoOk returns a tuple with the RemoveLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetRemoveLogoOk() (*bool, bool) {
	if o == nil || o.RemoveLogo == nil {
		return nil, false
	}
	return o.RemoveLogo, true
}

// HasRemoveLogo returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasRemoveLogo() bool {
	if o != nil && o.RemoveLogo != nil {
		return true
	}

	return false
}

// SetRemoveLogo gets a reference to the given bool and assigns it to the RemoveLogo field.
func (o *CompaniesIdJsonCompany) SetRemoveLogo(v bool) {
	o.RemoveLogo = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CompaniesIdJsonCompany) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetTags() []map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetTagsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []map[string]interface{} and assigns it to the Tags field.
func (o *CompaniesIdJsonCompany) SetTags(v []map[string]interface{}) {
	o.Tags = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CompaniesIdJsonCompany) SetWebsite(v string) {
	o.Website = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CompaniesIdJsonCompany) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesIdJsonCompany) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CompaniesIdJsonCompany) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CompaniesIdJsonCompany) SetZip(v string) {
	o.Zip = &v
}

func (o CompaniesIdJsonCompany) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressOne != nil {
		toSerialize["address_one"] = o.AddressOne
	}
	if o.AddressTwo != nil {
		toSerialize["address_two"] = o.AddressTwo
	}
	if o.Cid != nil {
		toSerialize["cid"] = o.Cid
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Countrycode != nil {
		toSerialize["countrycode"] = o.Countrycode
	}
	if o.EmailOne != nil {
		toSerialize["email_one"] = o.EmailOne
	}
	if o.EmailThree != nil {
		toSerialize["email_three"] = o.EmailThree
	}
	if o.EmailTwo != nil {
		toSerialize["email_two"] = o.EmailTwo
	}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.IndustryCatId != nil {
		toSerialize["industryCatId"] = o.IndustryCatId
	}
	if o.LogoPendingFileRef != nil {
		toSerialize["logoPendingFileRef"] = o.LogoPendingFileRef
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.PrivateNotes != nil {
		toSerialize["privateNotes"] = o.PrivateNotes
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.RemoveLogo != nil {
		toSerialize["removeLogo"] = o.RemoveLogo
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableCompaniesIdJsonCompany struct {
	value *CompaniesIdJsonCompany
	isSet bool
}

func (v NullableCompaniesIdJsonCompany) Get() *CompaniesIdJsonCompany {
	return v.value
}

func (v *NullableCompaniesIdJsonCompany) Set(val *CompaniesIdJsonCompany) {
	v.value = val
	v.isSet = true
}

func (v NullableCompaniesIdJsonCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompaniesIdJsonCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompaniesIdJsonCompany(val *CompaniesIdJsonCompany) *NullableCompaniesIdJsonCompany {
	return &NullableCompaniesIdJsonCompany{value: val, isSet: true}
}

func (v NullableCompaniesIdJsonCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompaniesIdJsonCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


