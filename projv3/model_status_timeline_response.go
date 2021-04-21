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

// StatusTimelineResponse TimelineResponse contains information about a group of statuses.
type StatusTimelineResponse struct {
	Included *StatusTimelineResponseIncluded `json:"included,omitempty"`
	Meta *ViewMeta `json:"meta,omitempty"`
	Statuses *[]ViewStatus `json:"statuses,omitempty"`
}

// NewStatusTimelineResponse instantiates a new StatusTimelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusTimelineResponse() *StatusTimelineResponse {
	this := StatusTimelineResponse{}
	return &this
}

// NewStatusTimelineResponseWithDefaults instantiates a new StatusTimelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusTimelineResponseWithDefaults() *StatusTimelineResponse {
	this := StatusTimelineResponse{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *StatusTimelineResponse) GetIncluded() StatusTimelineResponseIncluded {
	if o == nil || o.Included == nil {
		var ret StatusTimelineResponseIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusTimelineResponse) GetIncludedOk() (*StatusTimelineResponseIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *StatusTimelineResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given StatusTimelineResponseIncluded and assigns it to the Included field.
func (o *StatusTimelineResponse) SetIncluded(v StatusTimelineResponseIncluded) {
	o.Included = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StatusTimelineResponse) GetMeta() ViewMeta {
	if o == nil || o.Meta == nil {
		var ret ViewMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusTimelineResponse) GetMetaOk() (*ViewMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StatusTimelineResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ViewMeta and assigns it to the Meta field.
func (o *StatusTimelineResponse) SetMeta(v ViewMeta) {
	o.Meta = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *StatusTimelineResponse) GetStatuses() []ViewStatus {
	if o == nil || o.Statuses == nil {
		var ret []ViewStatus
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusTimelineResponse) GetStatusesOk() (*[]ViewStatus, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *StatusTimelineResponse) HasStatuses() bool {
	if o != nil && o.Statuses != nil {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []ViewStatus and assigns it to the Statuses field.
func (o *StatusTimelineResponse) SetStatuses(v []ViewStatus) {
	o.Statuses = &v
}

func (o StatusTimelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	return json.Marshal(toSerialize)
}

type NullableStatusTimelineResponse struct {
	value *StatusTimelineResponse
	isSet bool
}

func (v NullableStatusTimelineResponse) Get() *StatusTimelineResponse {
	return v.value
}

func (v *NullableStatusTimelineResponse) Set(val *StatusTimelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusTimelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusTimelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusTimelineResponse(val *StatusTimelineResponse) *NullableStatusTimelineResponse {
	return &NullableStatusTimelineResponse{value: val, isSet: true}
}

func (v NullableStatusTimelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusTimelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

