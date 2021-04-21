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

// SummaryUnreadResponse UnreadResponse contains counters for unread objects.
type SummaryUnreadResponse struct {
	Comments *int32 `json:"comments,omitempty"`
	Messages *int32 `json:"messages,omitempty"`
}

// NewSummaryUnreadResponse instantiates a new SummaryUnreadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryUnreadResponse() *SummaryUnreadResponse {
	this := SummaryUnreadResponse{}
	return &this
}

// NewSummaryUnreadResponseWithDefaults instantiates a new SummaryUnreadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryUnreadResponseWithDefaults() *SummaryUnreadResponse {
	this := SummaryUnreadResponse{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *SummaryUnreadResponse) GetComments() int32 {
	if o == nil || o.Comments == nil {
		var ret int32
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryUnreadResponse) GetCommentsOk() (*int32, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *SummaryUnreadResponse) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given int32 and assigns it to the Comments field.
func (o *SummaryUnreadResponse) SetComments(v int32) {
	o.Comments = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SummaryUnreadResponse) GetMessages() int32 {
	if o == nil || o.Messages == nil {
		var ret int32
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryUnreadResponse) GetMessagesOk() (*int32, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SummaryUnreadResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given int32 and assigns it to the Messages field.
func (o *SummaryUnreadResponse) SetMessages(v int32) {
	o.Messages = &v
}

func (o SummaryUnreadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableSummaryUnreadResponse struct {
	value *SummaryUnreadResponse
	isSet bool
}

func (v NullableSummaryUnreadResponse) Get() *SummaryUnreadResponse {
	return v.value
}

func (v *NullableSummaryUnreadResponse) Set(val *SummaryUnreadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryUnreadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryUnreadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryUnreadResponse(val *SummaryUnreadResponse) *NullableSummaryUnreadResponse {
	return &NullableSummaryUnreadResponse{value: val, isSet: true}
}

func (v NullableSummaryUnreadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryUnreadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


