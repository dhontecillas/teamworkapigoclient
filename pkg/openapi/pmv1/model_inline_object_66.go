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

// InlineObject66 struct for InlineObject66
type InlineObject66 struct {
	Notebook *ProjectsIdNotebooksJsonNotebook `json:"notebook,omitempty"`
}

// NewInlineObject66 instantiates a new InlineObject66 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject66() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// NewInlineObject66WithDefaults instantiates a new InlineObject66 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject66WithDefaults() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// GetNotebook returns the Notebook field value if set, zero value otherwise.
func (o *InlineObject66) GetNotebook() ProjectsIdNotebooksJsonNotebook {
	if o == nil || o.Notebook == nil {
		var ret ProjectsIdNotebooksJsonNotebook
		return ret
	}
	return *o.Notebook
}

// GetNotebookOk returns a tuple with the Notebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetNotebookOk() (*ProjectsIdNotebooksJsonNotebook, bool) {
	if o == nil || o.Notebook == nil {
		return nil, false
	}
	return o.Notebook, true
}

// HasNotebook returns a boolean if a field has been set.
func (o *InlineObject66) HasNotebook() bool {
	if o != nil && o.Notebook != nil {
		return true
	}

	return false
}

// SetNotebook gets a reference to the given ProjectsIdNotebooksJsonNotebook and assigns it to the Notebook field.
func (o *InlineObject66) SetNotebook(v ProjectsIdNotebooksJsonNotebook) {
	o.Notebook = &v
}

func (o InlineObject66) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Notebook != nil {
		toSerialize["notebook"] = o.Notebook
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject66 struct {
	value *InlineObject66
	isSet bool
}

func (v NullableInlineObject66) Get() *InlineObject66 {
	return v.value
}

func (v *NullableInlineObject66) Set(val *InlineObject66) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject66) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject66) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject66(val *InlineObject66) *NullableInlineObject66 {
	return &NullableInlineObject66{value: val, isSet: true}
}

func (v NullableInlineObject66) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject66) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


