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

// InvoiceProjectMetricInvoicesResponse ProjectMetricInvoicesResponse contains information about a group of invoices.
type InvoiceProjectMetricInvoicesResponse struct {
	Data *InvoiceProjectMetricInvoicesResponseData `json:"data,omitempty"`
}

// NewInvoiceProjectMetricInvoicesResponse instantiates a new InvoiceProjectMetricInvoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceProjectMetricInvoicesResponse() *InvoiceProjectMetricInvoicesResponse {
	this := InvoiceProjectMetricInvoicesResponse{}
	return &this
}

// NewInvoiceProjectMetricInvoicesResponseWithDefaults instantiates a new InvoiceProjectMetricInvoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceProjectMetricInvoicesResponseWithDefaults() *InvoiceProjectMetricInvoicesResponse {
	this := InvoiceProjectMetricInvoicesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InvoiceProjectMetricInvoicesResponse) GetData() InvoiceProjectMetricInvoicesResponseData {
	if o == nil || o.Data == nil {
		var ret InvoiceProjectMetricInvoicesResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceProjectMetricInvoicesResponse) GetDataOk() (*InvoiceProjectMetricInvoicesResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvoiceProjectMetricInvoicesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given InvoiceProjectMetricInvoicesResponseData and assigns it to the Data field.
func (o *InvoiceProjectMetricInvoicesResponse) SetData(v InvoiceProjectMetricInvoicesResponseData) {
	o.Data = &v
}

func (o InvoiceProjectMetricInvoicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceProjectMetricInvoicesResponse struct {
	value *InvoiceProjectMetricInvoicesResponse
	isSet bool
}

func (v NullableInvoiceProjectMetricInvoicesResponse) Get() *InvoiceProjectMetricInvoicesResponse {
	return v.value
}

func (v *NullableInvoiceProjectMetricInvoicesResponse) Set(val *InvoiceProjectMetricInvoicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceProjectMetricInvoicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceProjectMetricInvoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceProjectMetricInvoicesResponse(val *InvoiceProjectMetricInvoicesResponse) *NullableInvoiceProjectMetricInvoicesResponse {
	return &NullableInvoiceProjectMetricInvoicesResponse{value: val, isSet: true}
}

func (v NullableInvoiceProjectMetricInvoicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceProjectMetricInvoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


