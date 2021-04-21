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

// InlineResponse20044Columns struct for InlineResponse20044Columns
type InlineResponse20044Columns struct {
	CanEdit *bool `json:"canEdit,omitempty"`
	Color *string `json:"color,omitempty"`
	DateCreated *string `json:"dateCreated,omitempty"`
	DateUpdated *string `json:"dateUpdated,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DeletedDate *string `json:"deletedDate,omitempty"`
	DisplayOrder *string `json:"displayOrder,omitempty"`
	HasTriggers *bool `json:"hasTriggers,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Settings *InlineResponse20044Settings `json:"settings,omitempty"`
	Sort *string `json:"sort,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty"`
}

// NewInlineResponse20044Columns instantiates a new InlineResponse20044Columns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044Columns() *InlineResponse20044Columns {
	this := InlineResponse20044Columns{}
	return &this
}

// NewInlineResponse20044ColumnsWithDefaults instantiates a new InlineResponse20044Columns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044ColumnsWithDefaults() *InlineResponse20044Columns {
	this := InlineResponse20044Columns{}
	return &this
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetCanEdit() bool {
	if o == nil || o.CanEdit == nil {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetCanEditOk() (*bool, bool) {
	if o == nil || o.CanEdit == nil {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasCanEdit() bool {
	if o != nil && o.CanEdit != nil {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *InlineResponse20044Columns) SetCanEdit(v bool) {
	o.CanEdit = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *InlineResponse20044Columns) SetColor(v string) {
	o.Color = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *InlineResponse20044Columns) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetDateUpdated() string {
	if o == nil || o.DateUpdated == nil {
		var ret string
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetDateUpdatedOk() (*string, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given string and assigns it to the DateUpdated field.
func (o *InlineResponse20044Columns) SetDateUpdated(v string) {
	o.DateUpdated = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *InlineResponse20044Columns) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDeletedDate returns the DeletedDate field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetDeletedDate() string {
	if o == nil || o.DeletedDate == nil {
		var ret string
		return ret
	}
	return *o.DeletedDate
}

// GetDeletedDateOk returns a tuple with the DeletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetDeletedDateOk() (*string, bool) {
	if o == nil || o.DeletedDate == nil {
		return nil, false
	}
	return o.DeletedDate, true
}

// HasDeletedDate returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasDeletedDate() bool {
	if o != nil && o.DeletedDate != nil {
		return true
	}

	return false
}

// SetDeletedDate gets a reference to the given string and assigns it to the DeletedDate field.
func (o *InlineResponse20044Columns) SetDeletedDate(v string) {
	o.DeletedDate = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetDisplayOrder() string {
	if o == nil || o.DisplayOrder == nil {
		var ret string
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetDisplayOrderOk() (*string, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given string and assigns it to the DisplayOrder field.
func (o *InlineResponse20044Columns) SetDisplayOrder(v string) {
	o.DisplayOrder = &v
}

// GetHasTriggers returns the HasTriggers field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetHasTriggers() bool {
	if o == nil || o.HasTriggers == nil {
		var ret bool
		return ret
	}
	return *o.HasTriggers
}

// GetHasTriggersOk returns a tuple with the HasTriggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetHasTriggersOk() (*bool, bool) {
	if o == nil || o.HasTriggers == nil {
		return nil, false
	}
	return o.HasTriggers, true
}

// HasHasTriggers returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasHasTriggers() bool {
	if o != nil && o.HasTriggers != nil {
		return true
	}

	return false
}

// SetHasTriggers gets a reference to the given bool and assigns it to the HasTriggers field.
func (o *InlineResponse20044Columns) SetHasTriggers(v bool) {
	o.HasTriggers = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20044Columns) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20044Columns) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetSettings() InlineResponse20044Settings {
	if o == nil || o.Settings == nil {
		var ret InlineResponse20044Settings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetSettingsOk() (*InlineResponse20044Settings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given InlineResponse20044Settings and assigns it to the Settings field.
func (o *InlineResponse20044Columns) SetSettings(v InlineResponse20044Settings) {
	o.Settings = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *InlineResponse20044Columns) SetSort(v string) {
	o.Sort = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *InlineResponse20044Columns) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Columns) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *InlineResponse20044Columns) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *InlineResponse20044Columns) SetSortOrder(v string) {
	o.SortOrder = &v
}

func (o InlineResponse20044Columns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanEdit != nil {
		toSerialize["canEdit"] = o.CanEdit
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.DateUpdated != nil {
		toSerialize["dateUpdated"] = o.DateUpdated
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.DeletedDate != nil {
		toSerialize["deletedDate"] = o.DeletedDate
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.HasTriggers != nil {
		toSerialize["hasTriggers"] = o.HasTriggers
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044Columns struct {
	value *InlineResponse20044Columns
	isSet bool
}

func (v NullableInlineResponse20044Columns) Get() *InlineResponse20044Columns {
	return v.value
}

func (v *NullableInlineResponse20044Columns) Set(val *InlineResponse20044Columns) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044Columns) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044Columns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044Columns(val *InlineResponse20044Columns) *NullableInlineResponse20044Columns {
	return &NullableInlineResponse20044Columns{value: val, isSet: true}
}

func (v NullableInlineResponse20044Columns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044Columns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


