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

// InlineResponse2009Reactions struct for InlineResponse2009Reactions
type InlineResponse2009Reactions struct {
	AvatarURL *string `json:"avatarURL,omitempty"`
	CompanyId *string `json:"companyId,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	PostedOn *string `json:"postedOn,omitempty"`
	ReactionType *string `json:"reactionType,omitempty"`
}

// NewInlineResponse2009Reactions instantiates a new InlineResponse2009Reactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009Reactions() *InlineResponse2009Reactions {
	this := InlineResponse2009Reactions{}
	return &this
}

// NewInlineResponse2009ReactionsWithDefaults instantiates a new InlineResponse2009Reactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009ReactionsWithDefaults() *InlineResponse2009Reactions {
	this := InlineResponse2009Reactions{}
	return &this
}

// GetAvatarURL returns the AvatarURL field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetAvatarURL() string {
	if o == nil || o.AvatarURL == nil {
		var ret string
		return ret
	}
	return *o.AvatarURL
}

// GetAvatarURLOk returns a tuple with the AvatarURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetAvatarURLOk() (*string, bool) {
	if o == nil || o.AvatarURL == nil {
		return nil, false
	}
	return o.AvatarURL, true
}

// HasAvatarURL returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasAvatarURL() bool {
	if o != nil && o.AvatarURL != nil {
		return true
	}

	return false
}

// SetAvatarURL gets a reference to the given string and assigns it to the AvatarURL field.
func (o *InlineResponse2009Reactions) SetAvatarURL(v string) {
	o.AvatarURL = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *InlineResponse2009Reactions) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *InlineResponse2009Reactions) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *InlineResponse2009Reactions) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *InlineResponse2009Reactions) SetFullName(v string) {
	o.FullName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2009Reactions) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *InlineResponse2009Reactions) SetLastName(v string) {
	o.LastName = &v
}

// GetPostedOn returns the PostedOn field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetPostedOn() string {
	if o == nil || o.PostedOn == nil {
		var ret string
		return ret
	}
	return *o.PostedOn
}

// GetPostedOnOk returns a tuple with the PostedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetPostedOnOk() (*string, bool) {
	if o == nil || o.PostedOn == nil {
		return nil, false
	}
	return o.PostedOn, true
}

// HasPostedOn returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasPostedOn() bool {
	if o != nil && o.PostedOn != nil {
		return true
	}

	return false
}

// SetPostedOn gets a reference to the given string and assigns it to the PostedOn field.
func (o *InlineResponse2009Reactions) SetPostedOn(v string) {
	o.PostedOn = &v
}

// GetReactionType returns the ReactionType field value if set, zero value otherwise.
func (o *InlineResponse2009Reactions) GetReactionType() string {
	if o == nil || o.ReactionType == nil {
		var ret string
		return ret
	}
	return *o.ReactionType
}

// GetReactionTypeOk returns a tuple with the ReactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Reactions) GetReactionTypeOk() (*string, bool) {
	if o == nil || o.ReactionType == nil {
		return nil, false
	}
	return o.ReactionType, true
}

// HasReactionType returns a boolean if a field has been set.
func (o *InlineResponse2009Reactions) HasReactionType() bool {
	if o != nil && o.ReactionType != nil {
		return true
	}

	return false
}

// SetReactionType gets a reference to the given string and assigns it to the ReactionType field.
func (o *InlineResponse2009Reactions) SetReactionType(v string) {
	o.ReactionType = &v
}

func (o InlineResponse2009Reactions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarURL != nil {
		toSerialize["avatarURL"] = o.AvatarURL
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CompanyName != nil {
		toSerialize["companyName"] = o.CompanyName
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.FullName != nil {
		toSerialize["fullName"] = o.FullName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.PostedOn != nil {
		toSerialize["postedOn"] = o.PostedOn
	}
	if o.ReactionType != nil {
		toSerialize["reactionType"] = o.ReactionType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009Reactions struct {
	value *InlineResponse2009Reactions
	isSet bool
}

func (v NullableInlineResponse2009Reactions) Get() *InlineResponse2009Reactions {
	return v.value
}

func (v *NullableInlineResponse2009Reactions) Set(val *InlineResponse2009Reactions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009Reactions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009Reactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009Reactions(val *InlineResponse2009Reactions) *NullableInlineResponse2009Reactions {
	return &NullableInlineResponse2009Reactions{value: val, isSet: true}
}

func (v NullableInlineResponse2009Reactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009Reactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


