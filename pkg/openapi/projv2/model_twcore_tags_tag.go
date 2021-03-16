/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv2

import (
	"encoding/json"
)

// TwcoreTagsTag struct for TwcoreTagsTag
type TwcoreTagsTag struct {
	Color *string `json:"color,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewTwcoreTagsTag instantiates a new TwcoreTagsTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwcoreTagsTag() *TwcoreTagsTag {
	this := TwcoreTagsTag{}
	return &this
}

// NewTwcoreTagsTagWithDefaults instantiates a new TwcoreTagsTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwcoreTagsTagWithDefaults() *TwcoreTagsTag {
	this := TwcoreTagsTag{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TwcoreTagsTag) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwcoreTagsTag) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TwcoreTagsTag) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TwcoreTagsTag) SetColor(v string) {
	o.Color = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TwcoreTagsTag) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwcoreTagsTag) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TwcoreTagsTag) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TwcoreTagsTag) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TwcoreTagsTag) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwcoreTagsTag) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TwcoreTagsTag) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TwcoreTagsTag) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *TwcoreTagsTag) GetProjectId() int32 {
	if o == nil || o.ProjectId == nil {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwcoreTagsTag) GetProjectIdOk() (*int32, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TwcoreTagsTag) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *TwcoreTagsTag) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o TwcoreTagsTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableTwcoreTagsTag struct {
	value *TwcoreTagsTag
	isSet bool
}

func (v NullableTwcoreTagsTag) Get() *TwcoreTagsTag {
	return v.value
}

func (v *NullableTwcoreTagsTag) Set(val *TwcoreTagsTag) {
	v.value = val
	v.isSet = true
}

func (v NullableTwcoreTagsTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTwcoreTagsTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwcoreTagsTag(val *TwcoreTagsTag) *NullableTwcoreTagsTag {
	return &NullableTwcoreTagsTag{value: val, isSet: true}
}

func (v NullableTwcoreTagsTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwcoreTagsTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


