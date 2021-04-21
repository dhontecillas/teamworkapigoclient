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

// ViewLockdown Lockdown contains all the information returned from a lockdown.
type ViewLockdown struct {
	GrantAccessTo *ViewUserGroups `json:"grantAccessTo,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Item *ViewRelationship `json:"item,omitempty"`
	ItemID *int32 `json:"itemID,omitempty"`
	ItemType *string `json:"itemType,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	User *ViewRelationship `json:"user,omitempty"`
	UserID *int32 `json:"userID,omitempty"`
}

// NewViewLockdown instantiates a new ViewLockdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewLockdown() *ViewLockdown {
	this := ViewLockdown{}
	return &this
}

// NewViewLockdownWithDefaults instantiates a new ViewLockdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewLockdownWithDefaults() *ViewLockdown {
	this := ViewLockdown{}
	return &this
}

// GetGrantAccessTo returns the GrantAccessTo field value if set, zero value otherwise.
func (o *ViewLockdown) GetGrantAccessTo() ViewUserGroups {
	if o == nil || o.GrantAccessTo == nil {
		var ret ViewUserGroups
		return ret
	}
	return *o.GrantAccessTo
}

// GetGrantAccessToOk returns a tuple with the GrantAccessTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetGrantAccessToOk() (*ViewUserGroups, bool) {
	if o == nil || o.GrantAccessTo == nil {
		return nil, false
	}
	return o.GrantAccessTo, true
}

// HasGrantAccessTo returns a boolean if a field has been set.
func (o *ViewLockdown) HasGrantAccessTo() bool {
	if o != nil && o.GrantAccessTo != nil {
		return true
	}

	return false
}

// SetGrantAccessTo gets a reference to the given ViewUserGroups and assigns it to the GrantAccessTo field.
func (o *ViewLockdown) SetGrantAccessTo(v ViewUserGroups) {
	o.GrantAccessTo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewLockdown) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewLockdown) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewLockdown) SetId(v int32) {
	o.Id = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *ViewLockdown) GetItem() ViewRelationship {
	if o == nil || o.Item == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetItemOk() (*ViewRelationship, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *ViewLockdown) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given ViewRelationship and assigns it to the Item field.
func (o *ViewLockdown) SetItem(v ViewRelationship) {
	o.Item = &v
}

// GetItemID returns the ItemID field value if set, zero value otherwise.
func (o *ViewLockdown) GetItemID() int32 {
	if o == nil || o.ItemID == nil {
		var ret int32
		return ret
	}
	return *o.ItemID
}

// GetItemIDOk returns a tuple with the ItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetItemIDOk() (*int32, bool) {
	if o == nil || o.ItemID == nil {
		return nil, false
	}
	return o.ItemID, true
}

// HasItemID returns a boolean if a field has been set.
func (o *ViewLockdown) HasItemID() bool {
	if o != nil && o.ItemID != nil {
		return true
	}

	return false
}

// SetItemID gets a reference to the given int32 and assigns it to the ItemID field.
func (o *ViewLockdown) SetItemID(v int32) {
	o.ItemID = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *ViewLockdown) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *ViewLockdown) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *ViewLockdown) SetItemType(v string) {
	o.ItemType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ViewLockdown) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ViewLockdown) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ViewLockdown) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ViewLockdown) GetUser() ViewRelationship {
	if o == nil || o.User == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetUserOk() (*ViewRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ViewLockdown) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ViewRelationship and assigns it to the User field.
func (o *ViewLockdown) SetUser(v ViewRelationship) {
	o.User = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *ViewLockdown) GetUserID() int32 {
	if o == nil || o.UserID == nil {
		var ret int32
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLockdown) GetUserIDOk() (*int32, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *ViewLockdown) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given int32 and assigns it to the UserID field.
func (o *ViewLockdown) SetUserID(v int32) {
	o.UserID = &v
}

func (o ViewLockdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantAccessTo != nil {
		toSerialize["grantAccessTo"] = o.GrantAccessTo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	if o.ItemID != nil {
		toSerialize["itemID"] = o.ItemID
	}
	if o.ItemType != nil {
		toSerialize["itemType"] = o.ItemType
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	return json.Marshal(toSerialize)
}

type NullableViewLockdown struct {
	value *ViewLockdown
	isSet bool
}

func (v NullableViewLockdown) Get() *ViewLockdown {
	return v.value
}

func (v *NullableViewLockdown) Set(val *ViewLockdown) {
	v.value = val
	v.isSet = true
}

func (v NullableViewLockdown) IsSet() bool {
	return v.isSet
}

func (v *NullableViewLockdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewLockdown(val *ViewLockdown) *NullableViewLockdown {
	return &NullableViewLockdown{value: val, isSet: true}
}

func (v NullableViewLockdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewLockdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


