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

// InlineResponse20039PersonIntegrations struct for InlineResponse20039PersonIntegrations
type InlineResponse20039PersonIntegrations struct {
	MicrosoftConnector *InlineResponse20039PersonIntegrationsMicrosoftConnector `json:"microsoftConnector,omitempty"`
}

// NewInlineResponse20039PersonIntegrations instantiates a new InlineResponse20039PersonIntegrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20039PersonIntegrations() *InlineResponse20039PersonIntegrations {
	this := InlineResponse20039PersonIntegrations{}
	return &this
}

// NewInlineResponse20039PersonIntegrationsWithDefaults instantiates a new InlineResponse20039PersonIntegrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20039PersonIntegrationsWithDefaults() *InlineResponse20039PersonIntegrations {
	this := InlineResponse20039PersonIntegrations{}
	return &this
}

// GetMicrosoftConnector returns the MicrosoftConnector field value if set, zero value otherwise.
func (o *InlineResponse20039PersonIntegrations) GetMicrosoftConnector() InlineResponse20039PersonIntegrationsMicrosoftConnector {
	if o == nil || o.MicrosoftConnector == nil {
		var ret InlineResponse20039PersonIntegrationsMicrosoftConnector
		return ret
	}
	return *o.MicrosoftConnector
}

// GetMicrosoftConnectorOk returns a tuple with the MicrosoftConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20039PersonIntegrations) GetMicrosoftConnectorOk() (*InlineResponse20039PersonIntegrationsMicrosoftConnector, bool) {
	if o == nil || o.MicrosoftConnector == nil {
		return nil, false
	}
	return o.MicrosoftConnector, true
}

// HasMicrosoftConnector returns a boolean if a field has been set.
func (o *InlineResponse20039PersonIntegrations) HasMicrosoftConnector() bool {
	if o != nil && o.MicrosoftConnector != nil {
		return true
	}

	return false
}

// SetMicrosoftConnector gets a reference to the given InlineResponse20039PersonIntegrationsMicrosoftConnector and assigns it to the MicrosoftConnector field.
func (o *InlineResponse20039PersonIntegrations) SetMicrosoftConnector(v InlineResponse20039PersonIntegrationsMicrosoftConnector) {
	o.MicrosoftConnector = &v
}

func (o InlineResponse20039PersonIntegrations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MicrosoftConnector != nil {
		toSerialize["microsoftConnector"] = o.MicrosoftConnector
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20039PersonIntegrations struct {
	value *InlineResponse20039PersonIntegrations
	isSet bool
}

func (v NullableInlineResponse20039PersonIntegrations) Get() *InlineResponse20039PersonIntegrations {
	return v.value
}

func (v *NullableInlineResponse20039PersonIntegrations) Set(val *InlineResponse20039PersonIntegrations) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20039PersonIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20039PersonIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20039PersonIntegrations(val *InlineResponse20039PersonIntegrations) *NullableInlineResponse20039PersonIntegrations {
	return &NullableInlineResponse20039PersonIntegrations{value: val, isSet: true}
}

func (v NullableInlineResponse20039PersonIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20039PersonIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


