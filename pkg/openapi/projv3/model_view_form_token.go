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

// ViewFormToken FormToken contains all the information returned from a token.
type ViewFormToken struct {
	CanonicalURL *string `json:"canonicalURL,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *ViewRelationship `json:"createdBy,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	DeletedBy *ViewRelationship `json:"deletedBy,omitempty"`
	Expires *bool `json:"expires,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	Form *ViewRelationship `json:"form,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Installation *ViewRelationship `json:"installation,omitempty"`
	State *string `json:"state,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *ViewRelationship `json:"updatedBy,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewViewFormToken instantiates a new ViewFormToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewFormToken() *ViewFormToken {
	this := ViewFormToken{}
	return &this
}

// NewViewFormTokenWithDefaults instantiates a new ViewFormToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewFormTokenWithDefaults() *ViewFormToken {
	this := ViewFormToken{}
	return &this
}

// GetCanonicalURL returns the CanonicalURL field value if set, zero value otherwise.
func (o *ViewFormToken) GetCanonicalURL() string {
	if o == nil || o.CanonicalURL == nil {
		var ret string
		return ret
	}
	return *o.CanonicalURL
}

// GetCanonicalURLOk returns a tuple with the CanonicalURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetCanonicalURLOk() (*string, bool) {
	if o == nil || o.CanonicalURL == nil {
		return nil, false
	}
	return o.CanonicalURL, true
}

// HasCanonicalURL returns a boolean if a field has been set.
func (o *ViewFormToken) HasCanonicalURL() bool {
	if o != nil && o.CanonicalURL != nil {
		return true
	}

	return false
}

// SetCanonicalURL gets a reference to the given string and assigns it to the CanonicalURL field.
func (o *ViewFormToken) SetCanonicalURL(v string) {
	o.CanonicalURL = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ViewFormToken) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ViewFormToken) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ViewFormToken) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ViewFormToken) GetCreatedBy() ViewRelationship {
	if o == nil || o.CreatedBy == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetCreatedByOk() (*ViewRelationship, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ViewFormToken) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ViewRelationship and assigns it to the CreatedBy field.
func (o *ViewFormToken) SetCreatedBy(v ViewRelationship) {
	o.CreatedBy = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ViewFormToken) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ViewFormToken) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ViewFormToken) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *ViewFormToken) GetDeletedBy() ViewRelationship {
	if o == nil || o.DeletedBy == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetDeletedByOk() (*ViewRelationship, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *ViewFormToken) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given ViewRelationship and assigns it to the DeletedBy field.
func (o *ViewFormToken) SetDeletedBy(v ViewRelationship) {
	o.DeletedBy = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *ViewFormToken) GetExpires() bool {
	if o == nil || o.Expires == nil {
		var ret bool
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetExpiresOk() (*bool, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *ViewFormToken) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given bool and assigns it to the Expires field.
func (o *ViewFormToken) SetExpires(v bool) {
	o.Expires = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ViewFormToken) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ViewFormToken) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ViewFormToken) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *ViewFormToken) GetForm() ViewRelationship {
	if o == nil || o.Form == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetFormOk() (*ViewRelationship, bool) {
	if o == nil || o.Form == nil {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *ViewFormToken) HasForm() bool {
	if o != nil && o.Form != nil {
		return true
	}

	return false
}

// SetForm gets a reference to the given ViewRelationship and assigns it to the Form field.
func (o *ViewFormToken) SetForm(v ViewRelationship) {
	o.Form = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewFormToken) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewFormToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewFormToken) SetId(v int32) {
	o.Id = &v
}

// GetInstallation returns the Installation field value if set, zero value otherwise.
func (o *ViewFormToken) GetInstallation() ViewRelationship {
	if o == nil || o.Installation == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Installation
}

// GetInstallationOk returns a tuple with the Installation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetInstallationOk() (*ViewRelationship, bool) {
	if o == nil || o.Installation == nil {
		return nil, false
	}
	return o.Installation, true
}

// HasInstallation returns a boolean if a field has been set.
func (o *ViewFormToken) HasInstallation() bool {
	if o != nil && o.Installation != nil {
		return true
	}

	return false
}

// SetInstallation gets a reference to the given ViewRelationship and assigns it to the Installation field.
func (o *ViewFormToken) SetInstallation(v ViewRelationship) {
	o.Installation = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ViewFormToken) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ViewFormToken) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ViewFormToken) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ViewFormToken) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ViewFormToken) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ViewFormToken) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ViewFormToken) GetUpdatedBy() ViewRelationship {
	if o == nil || o.UpdatedBy == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetUpdatedByOk() (*ViewRelationship, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ViewFormToken) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given ViewRelationship and assigns it to the UpdatedBy field.
func (o *ViewFormToken) SetUpdatedBy(v ViewRelationship) {
	o.UpdatedBy = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ViewFormToken) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewFormToken) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ViewFormToken) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ViewFormToken) SetValue(v string) {
	o.Value = &v
}

func (o ViewFormToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanonicalURL != nil {
		toSerialize["canonicalURL"] = o.CanonicalURL
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deletedBy"] = o.DeletedBy
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Installation != nil {
		toSerialize["installation"] = o.Installation
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableViewFormToken struct {
	value *ViewFormToken
	isSet bool
}

func (v NullableViewFormToken) Get() *ViewFormToken {
	return v.value
}

func (v *NullableViewFormToken) Set(val *ViewFormToken) {
	v.value = val
	v.isSet = true
}

func (v NullableViewFormToken) IsSet() bool {
	return v.isSet
}

func (v *NullableViewFormToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewFormToken(val *ViewFormToken) *NullableViewFormToken {
	return &NullableViewFormToken{value: val, isSet: true}
}

func (v NullableViewFormToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewFormToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


