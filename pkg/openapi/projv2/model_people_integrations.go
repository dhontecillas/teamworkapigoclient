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

// PeopleIntegrations struct for PeopleIntegrations
type PeopleIntegrations struct {
	MicrosoftConnector *PeopleIntegrationEnabled `json:"microsoftConnector,omitempty"`
}

// NewPeopleIntegrations instantiates a new PeopleIntegrations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeopleIntegrations() *PeopleIntegrations {
	this := PeopleIntegrations{}
	return &this
}

// NewPeopleIntegrationsWithDefaults instantiates a new PeopleIntegrations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeopleIntegrationsWithDefaults() *PeopleIntegrations {
	this := PeopleIntegrations{}
	return &this
}

// GetMicrosoftConnector returns the MicrosoftConnector field value if set, zero value otherwise.
func (o *PeopleIntegrations) GetMicrosoftConnector() PeopleIntegrationEnabled {
	if o == nil || o.MicrosoftConnector == nil {
		var ret PeopleIntegrationEnabled
		return ret
	}
	return *o.MicrosoftConnector
}

// GetMicrosoftConnectorOk returns a tuple with the MicrosoftConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleIntegrations) GetMicrosoftConnectorOk() (*PeopleIntegrationEnabled, bool) {
	if o == nil || o.MicrosoftConnector == nil {
		return nil, false
	}
	return o.MicrosoftConnector, true
}

// HasMicrosoftConnector returns a boolean if a field has been set.
func (o *PeopleIntegrations) HasMicrosoftConnector() bool {
	if o != nil && o.MicrosoftConnector != nil {
		return true
	}

	return false
}

// SetMicrosoftConnector gets a reference to the given PeopleIntegrationEnabled and assigns it to the MicrosoftConnector field.
func (o *PeopleIntegrations) SetMicrosoftConnector(v PeopleIntegrationEnabled) {
	o.MicrosoftConnector = &v
}

func (o PeopleIntegrations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MicrosoftConnector != nil {
		toSerialize["microsoftConnector"] = o.MicrosoftConnector
	}
	return json.Marshal(toSerialize)
}

type NullablePeopleIntegrations struct {
	value *PeopleIntegrations
	isSet bool
}

func (v NullablePeopleIntegrations) Get() *PeopleIntegrations {
	return v.value
}

func (v *NullablePeopleIntegrations) Set(val *PeopleIntegrations) {
	v.value = val
	v.isSet = true
}

func (v NullablePeopleIntegrations) IsSet() bool {
	return v.isSet
}

func (v *NullablePeopleIntegrations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeopleIntegrations(val *PeopleIntegrations) *NullablePeopleIntegrations {
	return &NullablePeopleIntegrations{value: val, isSet: true}
}

func (v NullablePeopleIntegrations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeopleIntegrations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


