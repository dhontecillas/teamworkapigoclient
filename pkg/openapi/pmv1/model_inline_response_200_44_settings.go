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

// InlineResponse20044Settings struct for InlineResponse20044Settings
type InlineResponse20044Settings struct {
	Category *bool `json:"category,omitempty"`
	Company *bool `json:"company,omitempty"`
	Dates *bool `json:"dates,omitempty"`
	Description *bool `json:"description,omitempty"`
	Name *bool `json:"name,omitempty"`
	Starred *bool `json:"starred,omitempty"`
	Tags *bool `json:"tags,omitempty"`
}

// NewInlineResponse20044Settings instantiates a new InlineResponse20044Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044Settings() *InlineResponse20044Settings {
	this := InlineResponse20044Settings{}
	return &this
}

// NewInlineResponse20044SettingsWithDefaults instantiates a new InlineResponse20044Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044SettingsWithDefaults() *InlineResponse20044Settings {
	this := InlineResponse20044Settings{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetCategory() bool {
	if o == nil || o.Category == nil {
		var ret bool
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetCategoryOk() (*bool, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given bool and assigns it to the Category field.
func (o *InlineResponse20044Settings) SetCategory(v bool) {
	o.Category = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetCompany() bool {
	if o == nil || o.Company == nil {
		var ret bool
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetCompanyOk() (*bool, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given bool and assigns it to the Company field.
func (o *InlineResponse20044Settings) SetCompany(v bool) {
	o.Company = &v
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetDates() bool {
	if o == nil || o.Dates == nil {
		var ret bool
		return ret
	}
	return *o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetDatesOk() (*bool, bool) {
	if o == nil || o.Dates == nil {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasDates() bool {
	if o != nil && o.Dates != nil {
		return true
	}

	return false
}

// SetDates gets a reference to the given bool and assigns it to the Dates field.
func (o *InlineResponse20044Settings) SetDates(v bool) {
	o.Dates = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetDescription() bool {
	if o == nil || o.Description == nil {
		var ret bool
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetDescriptionOk() (*bool, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given bool and assigns it to the Description field.
func (o *InlineResponse20044Settings) SetDescription(v bool) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetName() bool {
	if o == nil || o.Name == nil {
		var ret bool
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetNameOk() (*bool, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given bool and assigns it to the Name field.
func (o *InlineResponse20044Settings) SetName(v bool) {
	o.Name = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetStarred() bool {
	if o == nil || o.Starred == nil {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetStarredOk() (*bool, bool) {
	if o == nil || o.Starred == nil {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasStarred() bool {
	if o != nil && o.Starred != nil {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *InlineResponse20044Settings) SetStarred(v bool) {
	o.Starred = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20044Settings) GetTags() bool {
	if o == nil || o.Tags == nil {
		var ret bool
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044Settings) GetTagsOk() (*bool, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20044Settings) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given bool and assigns it to the Tags field.
func (o *InlineResponse20044Settings) SetTags(v bool) {
	o.Tags = &v
}

func (o InlineResponse20044Settings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Dates != nil {
		toSerialize["dates"] = o.Dates
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Starred != nil {
		toSerialize["starred"] = o.Starred
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20044Settings struct {
	value *InlineResponse20044Settings
	isSet bool
}

func (v NullableInlineResponse20044Settings) Get() *InlineResponse20044Settings {
	return v.value
}

func (v *NullableInlineResponse20044Settings) Set(val *InlineResponse20044Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044Settings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044Settings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044Settings(val *InlineResponse20044Settings) *NullableInlineResponse20044Settings {
	return &NullableInlineResponse20044Settings{value: val, isSet: true}
}

func (v NullableInlineResponse20044Settings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044Settings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


