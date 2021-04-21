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

// InlineResponse20075Stats struct for InlineResponse20075Stats
type InlineResponse20075Stats struct {
	Billing *InlineResponse20075StatsBilling `json:"billing,omitempty"`
	Columns *InlineResponse20075StatsColumns `json:"columns,omitempty"`
	Events *InlineResponse20075StatsEvents `json:"events,omitempty"`
	Milestones *InlineResponse20075StatsMilestones `json:"milestones,omitempty"`
	Tasks *InlineResponse20075StatsTasks `json:"tasks,omitempty"`
}

// NewInlineResponse20075Stats instantiates a new InlineResponse20075Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075Stats() *InlineResponse20075Stats {
	this := InlineResponse20075Stats{}
	return &this
}

// NewInlineResponse20075StatsWithDefaults instantiates a new InlineResponse20075Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075StatsWithDefaults() *InlineResponse20075Stats {
	this := InlineResponse20075Stats{}
	return &this
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *InlineResponse20075Stats) GetBilling() InlineResponse20075StatsBilling {
	if o == nil || o.Billing == nil {
		var ret InlineResponse20075StatsBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Stats) GetBillingOk() (*InlineResponse20075StatsBilling, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *InlineResponse20075Stats) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given InlineResponse20075StatsBilling and assigns it to the Billing field.
func (o *InlineResponse20075Stats) SetBilling(v InlineResponse20075StatsBilling) {
	o.Billing = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *InlineResponse20075Stats) GetColumns() InlineResponse20075StatsColumns {
	if o == nil || o.Columns == nil {
		var ret InlineResponse20075StatsColumns
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Stats) GetColumnsOk() (*InlineResponse20075StatsColumns, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *InlineResponse20075Stats) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given InlineResponse20075StatsColumns and assigns it to the Columns field.
func (o *InlineResponse20075Stats) SetColumns(v InlineResponse20075StatsColumns) {
	o.Columns = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *InlineResponse20075Stats) GetEvents() InlineResponse20075StatsEvents {
	if o == nil || o.Events == nil {
		var ret InlineResponse20075StatsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Stats) GetEventsOk() (*InlineResponse20075StatsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *InlineResponse20075Stats) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given InlineResponse20075StatsEvents and assigns it to the Events field.
func (o *InlineResponse20075Stats) SetEvents(v InlineResponse20075StatsEvents) {
	o.Events = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *InlineResponse20075Stats) GetMilestones() InlineResponse20075StatsMilestones {
	if o == nil || o.Milestones == nil {
		var ret InlineResponse20075StatsMilestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Stats) GetMilestonesOk() (*InlineResponse20075StatsMilestones, bool) {
	if o == nil || o.Milestones == nil {
		return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *InlineResponse20075Stats) HasMilestones() bool {
	if o != nil && o.Milestones != nil {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given InlineResponse20075StatsMilestones and assigns it to the Milestones field.
func (o *InlineResponse20075Stats) SetMilestones(v InlineResponse20075StatsMilestones) {
	o.Milestones = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InlineResponse20075Stats) GetTasks() InlineResponse20075StatsTasks {
	if o == nil || o.Tasks == nil {
		var ret InlineResponse20075StatsTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075Stats) GetTasksOk() (*InlineResponse20075StatsTasks, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InlineResponse20075Stats) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given InlineResponse20075StatsTasks and assigns it to the Tasks field.
func (o *InlineResponse20075Stats) SetTasks(v InlineResponse20075StatsTasks) {
	o.Tasks = &v
}

func (o InlineResponse20075Stats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Milestones != nil {
		toSerialize["milestones"] = o.Milestones
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20075Stats struct {
	value *InlineResponse20075Stats
	isSet bool
}

func (v NullableInlineResponse20075Stats) Get() *InlineResponse20075Stats {
	return v.value
}

func (v *NullableInlineResponse20075Stats) Set(val *InlineResponse20075Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075Stats) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075Stats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075Stats(val *InlineResponse20075Stats) *NullableInlineResponse20075Stats {
	return &NullableInlineResponse20075Stats{value: val, isSet: true}
}

func (v NullableInlineResponse20075Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075Stats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

