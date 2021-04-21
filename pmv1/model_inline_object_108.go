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

// InlineObject108 struct for InlineObject108
type InlineObject108 struct {
	STATUS *string `json:"STATUS,omitempty"`
	Webhook *WebhooksIdJsonWebhook `json:"webhook,omitempty"`
}

// NewInlineObject108 instantiates a new InlineObject108 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject108() *InlineObject108 {
	this := InlineObject108{}
	return &this
}

// NewInlineObject108WithDefaults instantiates a new InlineObject108 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject108WithDefaults() *InlineObject108 {
	this := InlineObject108{}
	return &this
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *InlineObject108) GetSTATUS() string {
	if o == nil || o.STATUS == nil {
		var ret string
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetSTATUSOk() (*string, bool) {
	if o == nil || o.STATUS == nil {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *InlineObject108) HasSTATUS() bool {
	if o != nil && o.STATUS != nil {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given string and assigns it to the STATUS field.
func (o *InlineObject108) SetSTATUS(v string) {
	o.STATUS = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *InlineObject108) GetWebhook() WebhooksIdJsonWebhook {
	if o == nil || o.Webhook == nil {
		var ret WebhooksIdJsonWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetWebhookOk() (*WebhooksIdJsonWebhook, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *InlineObject108) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhooksIdJsonWebhook and assigns it to the Webhook field.
func (o *InlineObject108) SetWebhook(v WebhooksIdJsonWebhook) {
	o.Webhook = &v
}

func (o InlineObject108) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.STATUS != nil {
		toSerialize["STATUS"] = o.STATUS
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject108 struct {
	value *InlineObject108
	isSet bool
}

func (v NullableInlineObject108) Get() *InlineObject108 {
	return v.value
}

func (v *NullableInlineObject108) Set(val *InlineObject108) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject108) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject108) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject108(val *InlineObject108) *NullableInlineObject108 {
	return &NullableInlineObject108{value: val, isSet: true}
}

func (v NullableInlineObject108) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject108) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

