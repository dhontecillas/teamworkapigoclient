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

// PeopleResponse struct for PeopleResponse
type PeopleResponse struct {
	Companies *[]PeopleCompanyResponse `json:"companies,omitempty"`
	Letters *[]string `json:"letters,omitempty"`
	People *[]PeoplePersonOfPeople `json:"people,omitempty"`
}

// NewPeopleResponse instantiates a new PeopleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeopleResponse() *PeopleResponse {
	this := PeopleResponse{}
	return &this
}

// NewPeopleResponseWithDefaults instantiates a new PeopleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeopleResponseWithDefaults() *PeopleResponse {
	this := PeopleResponse{}
	return &this
}

// GetCompanies returns the Companies field value if set, zero value otherwise.
func (o *PeopleResponse) GetCompanies() []PeopleCompanyResponse {
	if o == nil || o.Companies == nil {
		var ret []PeopleCompanyResponse
		return ret
	}
	return *o.Companies
}

// GetCompaniesOk returns a tuple with the Companies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetCompaniesOk() (*[]PeopleCompanyResponse, bool) {
	if o == nil || o.Companies == nil {
		return nil, false
	}
	return o.Companies, true
}

// HasCompanies returns a boolean if a field has been set.
func (o *PeopleResponse) HasCompanies() bool {
	if o != nil && o.Companies != nil {
		return true
	}

	return false
}

// SetCompanies gets a reference to the given []PeopleCompanyResponse and assigns it to the Companies field.
func (o *PeopleResponse) SetCompanies(v []PeopleCompanyResponse) {
	o.Companies = &v
}

// GetLetters returns the Letters field value if set, zero value otherwise.
func (o *PeopleResponse) GetLetters() []string {
	if o == nil || o.Letters == nil {
		var ret []string
		return ret
	}
	return *o.Letters
}

// GetLettersOk returns a tuple with the Letters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetLettersOk() (*[]string, bool) {
	if o == nil || o.Letters == nil {
		return nil, false
	}
	return o.Letters, true
}

// HasLetters returns a boolean if a field has been set.
func (o *PeopleResponse) HasLetters() bool {
	if o != nil && o.Letters != nil {
		return true
	}

	return false
}

// SetLetters gets a reference to the given []string and assigns it to the Letters field.
func (o *PeopleResponse) SetLetters(v []string) {
	o.Letters = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PeopleResponse) GetPeople() []PeoplePersonOfPeople {
	if o == nil || o.People == nil {
		var ret []PeoplePersonOfPeople
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetPeopleOk() (*[]PeoplePersonOfPeople, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PeopleResponse) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []PeoplePersonOfPeople and assigns it to the People field.
func (o *PeopleResponse) SetPeople(v []PeoplePersonOfPeople) {
	o.People = &v
}

func (o PeopleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Companies != nil {
		toSerialize["companies"] = o.Companies
	}
	if o.Letters != nil {
		toSerialize["letters"] = o.Letters
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	return json.Marshal(toSerialize)
}

type NullablePeopleResponse struct {
	value *PeopleResponse
	isSet bool
}

func (v NullablePeopleResponse) Get() *PeopleResponse {
	return v.value
}

func (v *NullablePeopleResponse) Set(val *PeopleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePeopleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePeopleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeopleResponse(val *PeopleResponse) *NullablePeopleResponse {
	return &NullablePeopleResponse{value: val, isSet: true}
}

func (v NullablePeopleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeopleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


