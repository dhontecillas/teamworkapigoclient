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

// InlineResponse2002People12345 struct for InlineResponse2002People12345
type InlineResponse2002People12345 struct {
	AvatarUrl *string `json:"avatarUrl,omitempty"`
	Company *InlineResponse2002People12345Company `json:"company,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"lastName,omitempty"`
}

// NewInlineResponse2002People12345 instantiates a new InlineResponse2002People12345 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002People12345() *InlineResponse2002People12345 {
	this := InlineResponse2002People12345{}
	return &this
}

// NewInlineResponse2002People12345WithDefaults instantiates a new InlineResponse2002People12345 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002People12345WithDefaults() *InlineResponse2002People12345 {
	this := InlineResponse2002People12345{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *InlineResponse2002People12345) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002People12345) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *InlineResponse2002People12345) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *InlineResponse2002People12345) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse2002People12345) GetCompany() InlineResponse2002People12345Company {
	if o == nil || o.Company == nil {
		var ret InlineResponse2002People12345Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002People12345) GetCompanyOk() (*InlineResponse2002People12345Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse2002People12345) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given InlineResponse2002People12345Company and assigns it to the Company field.
func (o *InlineResponse2002People12345) SetCompany(v InlineResponse2002People12345Company) {
	o.Company = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *InlineResponse2002People12345) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002People12345) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *InlineResponse2002People12345) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *InlineResponse2002People12345) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2002People12345) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002People12345) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2002People12345) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2002People12345) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InlineResponse2002People12345) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002People12345) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InlineResponse2002People12345) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InlineResponse2002People12345) SetLastName(v string) {
	o.LastName = &v
}

func (o InlineResponse2002People12345) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarUrl != nil {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002People12345 struct {
	value *InlineResponse2002People12345
	isSet bool
}

func (v NullableInlineResponse2002People12345) Get() *InlineResponse2002People12345 {
	return v.value
}

func (v *NullableInlineResponse2002People12345) Set(val *InlineResponse2002People12345) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002People12345) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002People12345) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002People12345(val *InlineResponse2002People12345) *NullableInlineResponse2002People12345 {
	return &NullableInlineResponse2002People12345{value: val, isSet: true}
}

func (v NullableInlineResponse2002People12345) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002People12345) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


