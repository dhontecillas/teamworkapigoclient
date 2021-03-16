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

// TwdataaccountAnnouncementResponse struct for TwdataaccountAnnouncementResponse
type TwdataaccountAnnouncementResponse struct {
	CalloutImage *string `json:"calloutImage,omitempty"`
	Code *string `json:"code,omitempty"`
	Date *string `json:"date,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsRead *bool `json:"isRead,omitempty"`
	Link *string `json:"link,omitempty"`
	LinkText *string `json:"linkText,omitempty"`
	Text *string `json:"text,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewTwdataaccountAnnouncementResponse instantiates a new TwdataaccountAnnouncementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwdataaccountAnnouncementResponse() *TwdataaccountAnnouncementResponse {
	this := TwdataaccountAnnouncementResponse{}
	return &this
}

// NewTwdataaccountAnnouncementResponseWithDefaults instantiates a new TwdataaccountAnnouncementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwdataaccountAnnouncementResponseWithDefaults() *TwdataaccountAnnouncementResponse {
	this := TwdataaccountAnnouncementResponse{}
	return &this
}

// GetCalloutImage returns the CalloutImage field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetCalloutImage() string {
	if o == nil || o.CalloutImage == nil {
		var ret string
		return ret
	}
	return *o.CalloutImage
}

// GetCalloutImageOk returns a tuple with the CalloutImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetCalloutImageOk() (*string, bool) {
	if o == nil || o.CalloutImage == nil {
		return nil, false
	}
	return o.CalloutImage, true
}

// HasCalloutImage returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasCalloutImage() bool {
	if o != nil && o.CalloutImage != nil {
		return true
	}

	return false
}

// SetCalloutImage gets a reference to the given string and assigns it to the CalloutImage field.
func (o *TwdataaccountAnnouncementResponse) SetCalloutImage(v string) {
	o.CalloutImage = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TwdataaccountAnnouncementResponse) SetCode(v string) {
	o.Code = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *TwdataaccountAnnouncementResponse) SetDate(v string) {
	o.Date = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetHtml() string {
	if o == nil || o.Html == nil {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetHtmlOk() (*string, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *TwdataaccountAnnouncementResponse) SetHtml(v string) {
	o.Html = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TwdataaccountAnnouncementResponse) SetId(v int32) {
	o.Id = &v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetIsRead() bool {
	if o == nil || o.IsRead == nil {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetIsReadOk() (*bool, bool) {
	if o == nil || o.IsRead == nil {
		return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasIsRead() bool {
	if o != nil && o.IsRead != nil {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *TwdataaccountAnnouncementResponse) SetIsRead(v bool) {
	o.IsRead = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *TwdataaccountAnnouncementResponse) SetLink(v string) {
	o.Link = &v
}

// GetLinkText returns the LinkText field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetLinkText() string {
	if o == nil || o.LinkText == nil {
		var ret string
		return ret
	}
	return *o.LinkText
}

// GetLinkTextOk returns a tuple with the LinkText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetLinkTextOk() (*string, bool) {
	if o == nil || o.LinkText == nil {
		return nil, false
	}
	return o.LinkText, true
}

// HasLinkText returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasLinkText() bool {
	if o != nil && o.LinkText != nil {
		return true
	}

	return false
}

// SetLinkText gets a reference to the given string and assigns it to the LinkText field.
func (o *TwdataaccountAnnouncementResponse) SetLinkText(v string) {
	o.LinkText = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TwdataaccountAnnouncementResponse) SetText(v string) {
	o.Text = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TwdataaccountAnnouncementResponse) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TwdataaccountAnnouncementResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwdataaccountAnnouncementResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TwdataaccountAnnouncementResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TwdataaccountAnnouncementResponse) SetType(v string) {
	o.Type = &v
}

func (o TwdataaccountAnnouncementResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalloutImage != nil {
		toSerialize["calloutImage"] = o.CalloutImage
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsRead != nil {
		toSerialize["isRead"] = o.IsRead
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.LinkText != nil {
		toSerialize["linkText"] = o.LinkText
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTwdataaccountAnnouncementResponse struct {
	value *TwdataaccountAnnouncementResponse
	isSet bool
}

func (v NullableTwdataaccountAnnouncementResponse) Get() *TwdataaccountAnnouncementResponse {
	return v.value
}

func (v *NullableTwdataaccountAnnouncementResponse) Set(val *TwdataaccountAnnouncementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTwdataaccountAnnouncementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTwdataaccountAnnouncementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwdataaccountAnnouncementResponse(val *TwdataaccountAnnouncementResponse) *NullableTwdataaccountAnnouncementResponse {
	return &NullableTwdataaccountAnnouncementResponse{value: val, isSet: true}
}

func (v NullableTwdataaccountAnnouncementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwdataaccountAnnouncementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


