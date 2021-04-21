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

// CommentCommentsResponse CommentsResponse contains information about a group of comments.
type CommentCommentsResponse struct {
	Comments *[]CommentFullComment `json:"comments,omitempty"`
	Included *CommentCommentsResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
}

// NewCommentCommentsResponse instantiates a new CommentCommentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentCommentsResponse() *CommentCommentsResponse {
	this := CommentCommentsResponse{}
	return &this
}

// NewCommentCommentsResponseWithDefaults instantiates a new CommentCommentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentCommentsResponseWithDefaults() *CommentCommentsResponse {
	this := CommentCommentsResponse{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CommentCommentsResponse) GetComments() []CommentFullComment {
	if o == nil || o.Comments == nil {
		var ret []CommentFullComment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentCommentsResponse) GetCommentsOk() (*[]CommentFullComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CommentCommentsResponse) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []CommentFullComment and assigns it to the Comments field.
func (o *CommentCommentsResponse) SetComments(v []CommentFullComment) {
	o.Comments = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CommentCommentsResponse) GetIncluded() CommentCommentsResponseIncluded {
	if o == nil || o.Included == nil {
		var ret CommentCommentsResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentCommentsResponse) GetIncludedOk() (*CommentCommentsResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CommentCommentsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given CommentCommentsResponseIncluded and assigns it to the Included field.
func (o *CommentCommentsResponse) SetIncluded(v CommentCommentsResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommentCommentsResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentCommentsResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommentCommentsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *CommentCommentsResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

func (o CommentCommentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableCommentCommentsResponse struct {
	value *CommentCommentsResponse
	isSet bool
}

func (v NullableCommentCommentsResponse) Get() *CommentCommentsResponse {
	return v.value
}

func (v *NullableCommentCommentsResponse) Set(val *CommentCommentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentCommentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentCommentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentCommentsResponse(val *CommentCommentsResponse) *NullableCommentCommentsResponse {
	return &NullableCommentCommentsResponse{value: val, isSet: true}
}

func (v NullableCommentCommentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentCommentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


