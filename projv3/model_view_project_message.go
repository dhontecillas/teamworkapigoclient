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

// ViewProjectMessage ProjectMessage contains all the information returned from a project message.
type ViewProjectMessage struct {
	Id *int32 `json:"id,omitempty"`
	Post *ViewRelationship `json:"post,omitempty"`
	PostId *int32 `json:"postId,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewViewProjectMessage instantiates a new ViewProjectMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProjectMessage() *ViewProjectMessage {
	this := ViewProjectMessage{}
	return &this
}

// NewViewProjectMessageWithDefaults instantiates a new ViewProjectMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProjectMessageWithDefaults() *ViewProjectMessage {
	this := ViewProjectMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ViewProjectMessage) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMessage) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ViewProjectMessage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ViewProjectMessage) SetId(v int32) {
	o.Id = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *ViewProjectMessage) GetPost() ViewRelationship {
	if o == nil || o.Post == nil {
		var ret ViewRelationship
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMessage) GetPostOk() (*ViewRelationship, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *ViewProjectMessage) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given ViewRelationship and assigns it to the Post field.
func (o *ViewProjectMessage) SetPost(v ViewRelationship) {
	o.Post = &v
}

// GetPostId returns the PostId field value if set, zero value otherwise.
func (o *ViewProjectMessage) GetPostId() int32 {
	if o == nil || o.PostId == nil {
		var ret int32
		return ret
	}
	return *o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMessage) GetPostIdOk() (*int32, bool) {
	if o == nil || o.PostId == nil {
		return nil, false
	}
	return o.PostId, true
}

// HasPostId returns a boolean if a field has been set.
func (o *ViewProjectMessage) HasPostId() bool {
	if o != nil && o.PostId != nil {
		return true
	}

	return false
}

// SetPostId gets a reference to the given int32 and assigns it to the PostId field.
func (o *ViewProjectMessage) SetPostId(v int32) {
	o.PostId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ViewProjectMessage) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProjectMessage) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ViewProjectMessage) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ViewProjectMessage) SetTitle(v string) {
	o.Title = &v
}

func (o ViewProjectMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	if o.PostId != nil {
		toSerialize["postId"] = o.PostId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableViewProjectMessage struct {
	value *ViewProjectMessage
	isSet bool
}

func (v NullableViewProjectMessage) Get() *ViewProjectMessage {
	return v.value
}

func (v *NullableViewProjectMessage) Set(val *ViewProjectMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProjectMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProjectMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProjectMessage(val *ViewProjectMessage) *NullableViewProjectMessage {
	return &NullableViewProjectMessage{value: val, isSet: true}
}

func (v NullableViewProjectMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProjectMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


