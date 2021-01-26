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

// TagBulkDeleteRequest BulkDeleteRequest contains the ids of the tags that should be removed.
type TagBulkDeleteRequest struct {
	TagIds *[]int32 `json:"tagIds,omitempty"`
	TagNames *string `json:"tagNames,omitempty"`
}

// NewTagBulkDeleteRequest instantiates a new TagBulkDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagBulkDeleteRequest() *TagBulkDeleteRequest {
	this := TagBulkDeleteRequest{}
	return &this
}

// NewTagBulkDeleteRequestWithDefaults instantiates a new TagBulkDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagBulkDeleteRequestWithDefaults() *TagBulkDeleteRequest {
	this := TagBulkDeleteRequest{}
	return &this
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *TagBulkDeleteRequest) GetTagIds() []int32 {
	if o == nil || o.TagIds == nil {
		var ret []int32
		return ret
	}
	return *o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBulkDeleteRequest) GetTagIdsOk() (*[]int32, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *TagBulkDeleteRequest) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int32 and assigns it to the TagIds field.
func (o *TagBulkDeleteRequest) SetTagIds(v []int32) {
	o.TagIds = &v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise.
func (o *TagBulkDeleteRequest) GetTagNames() string {
	if o == nil || o.TagNames == nil {
		var ret string
		return ret
	}
	return *o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBulkDeleteRequest) GetTagNamesOk() (*string, bool) {
	if o == nil || o.TagNames == nil {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *TagBulkDeleteRequest) HasTagNames() bool {
	if o != nil && o.TagNames != nil {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given string and assigns it to the TagNames field.
func (o *TagBulkDeleteRequest) SetTagNames(v string) {
	o.TagNames = &v
}

func (o TagBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TagIds != nil {
		toSerialize["tagIds"] = o.TagIds
	}
	if o.TagNames != nil {
		toSerialize["tagNames"] = o.TagNames
	}
	return json.Marshal(toSerialize)
}

type NullableTagBulkDeleteRequest struct {
	value *TagBulkDeleteRequest
	isSet bool
}

func (v NullableTagBulkDeleteRequest) Get() *TagBulkDeleteRequest {
	return v.value
}

func (v *NullableTagBulkDeleteRequest) Set(val *TagBulkDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagBulkDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagBulkDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagBulkDeleteRequest(val *TagBulkDeleteRequest) *NullableTagBulkDeleteRequest {
	return &NullableTagBulkDeleteRequest{value: val, isSet: true}
}

func (v NullableTagBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagBulkDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


