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

// InlineResponse200122Activity struct for InlineResponse200122Activity
type InlineResponse200122Activity struct {
	Activitytype *string `json:"activitytype,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate *string `json:"due-date,omitempty"`
	Extradescription *string `json:"extradescription,omitempty"`
	ForUserAvatarUrl *string `json:"for-user-avatar-url,omitempty"`
	Foruserid *string `json:"foruserid,omitempty"`
	Forusername *string `json:"forusername,omitempty"`
	FromUserAvatarUrl *string `json:"from-user-avatar-url,omitempty"`
	Fromusername *string `json:"fromusername,omitempty"`
	Id *string `json:"id,omitempty"`
	Isprivate *string `json:"isprivate,omitempty"`
	Itemid *string `json:"itemid,omitempty"`
	Itemlink *string `json:"itemlink,omitempty"`
	Link *string `json:"link,omitempty"`
	ProjectId *string `json:"project-id,omitempty"`
	Publicinfo *string `json:"publicinfo,omitempty"`
	Type *string `json:"type,omitempty"`
	Userid *string `json:"userid,omitempty"`
}

// NewInlineResponse200122Activity instantiates a new InlineResponse200122Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200122Activity() *InlineResponse200122Activity {
	this := InlineResponse200122Activity{}
	return &this
}

// NewInlineResponse200122ActivityWithDefaults instantiates a new InlineResponse200122Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200122ActivityWithDefaults() *InlineResponse200122Activity {
	this := InlineResponse200122Activity{}
	return &this
}

// GetActivitytype returns the Activitytype field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetActivitytype() string {
	if o == nil || o.Activitytype == nil {
		var ret string
		return ret
	}
	return *o.Activitytype
}

// GetActivitytypeOk returns a tuple with the Activitytype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetActivitytypeOk() (*string, bool) {
	if o == nil || o.Activitytype == nil {
		return nil, false
	}
	return o.Activitytype, true
}

// HasActivitytype returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasActivitytype() bool {
	if o != nil && o.Activitytype != nil {
		return true
	}

	return false
}

// SetActivitytype gets a reference to the given string and assigns it to the Activitytype field.
func (o *InlineResponse200122Activity) SetActivitytype(v string) {
	o.Activitytype = &v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetDatetime() string {
	if o == nil || o.Datetime == nil {
		var ret string
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetDatetimeOk() (*string, bool) {
	if o == nil || o.Datetime == nil {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasDatetime() bool {
	if o != nil && o.Datetime != nil {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given string and assigns it to the Datetime field.
func (o *InlineResponse200122Activity) SetDatetime(v string) {
	o.Datetime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200122Activity) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InlineResponse200122Activity) SetDueDate(v string) {
	o.DueDate = &v
}

// GetExtradescription returns the Extradescription field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetExtradescription() string {
	if o == nil || o.Extradescription == nil {
		var ret string
		return ret
	}
	return *o.Extradescription
}

// GetExtradescriptionOk returns a tuple with the Extradescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetExtradescriptionOk() (*string, bool) {
	if o == nil || o.Extradescription == nil {
		return nil, false
	}
	return o.Extradescription, true
}

// HasExtradescription returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasExtradescription() bool {
	if o != nil && o.Extradescription != nil {
		return true
	}

	return false
}

// SetExtradescription gets a reference to the given string and assigns it to the Extradescription field.
func (o *InlineResponse200122Activity) SetExtradescription(v string) {
	o.Extradescription = &v
}

// GetForUserAvatarUrl returns the ForUserAvatarUrl field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetForUserAvatarUrl() string {
	if o == nil || o.ForUserAvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.ForUserAvatarUrl
}

// GetForUserAvatarUrlOk returns a tuple with the ForUserAvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetForUserAvatarUrlOk() (*string, bool) {
	if o == nil || o.ForUserAvatarUrl == nil {
		return nil, false
	}
	return o.ForUserAvatarUrl, true
}

// HasForUserAvatarUrl returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasForUserAvatarUrl() bool {
	if o != nil && o.ForUserAvatarUrl != nil {
		return true
	}

	return false
}

// SetForUserAvatarUrl gets a reference to the given string and assigns it to the ForUserAvatarUrl field.
func (o *InlineResponse200122Activity) SetForUserAvatarUrl(v string) {
	o.ForUserAvatarUrl = &v
}

// GetForuserid returns the Foruserid field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetForuserid() string {
	if o == nil || o.Foruserid == nil {
		var ret string
		return ret
	}
	return *o.Foruserid
}

// GetForuseridOk returns a tuple with the Foruserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetForuseridOk() (*string, bool) {
	if o == nil || o.Foruserid == nil {
		return nil, false
	}
	return o.Foruserid, true
}

// HasForuserid returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasForuserid() bool {
	if o != nil && o.Foruserid != nil {
		return true
	}

	return false
}

// SetForuserid gets a reference to the given string and assigns it to the Foruserid field.
func (o *InlineResponse200122Activity) SetForuserid(v string) {
	o.Foruserid = &v
}

// GetForusername returns the Forusername field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetForusername() string {
	if o == nil || o.Forusername == nil {
		var ret string
		return ret
	}
	return *o.Forusername
}

// GetForusernameOk returns a tuple with the Forusername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetForusernameOk() (*string, bool) {
	if o == nil || o.Forusername == nil {
		return nil, false
	}
	return o.Forusername, true
}

// HasForusername returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasForusername() bool {
	if o != nil && o.Forusername != nil {
		return true
	}

	return false
}

// SetForusername gets a reference to the given string and assigns it to the Forusername field.
func (o *InlineResponse200122Activity) SetForusername(v string) {
	o.Forusername = &v
}

// GetFromUserAvatarUrl returns the FromUserAvatarUrl field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetFromUserAvatarUrl() string {
	if o == nil || o.FromUserAvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.FromUserAvatarUrl
}

// GetFromUserAvatarUrlOk returns a tuple with the FromUserAvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetFromUserAvatarUrlOk() (*string, bool) {
	if o == nil || o.FromUserAvatarUrl == nil {
		return nil, false
	}
	return o.FromUserAvatarUrl, true
}

// HasFromUserAvatarUrl returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasFromUserAvatarUrl() bool {
	if o != nil && o.FromUserAvatarUrl != nil {
		return true
	}

	return false
}

// SetFromUserAvatarUrl gets a reference to the given string and assigns it to the FromUserAvatarUrl field.
func (o *InlineResponse200122Activity) SetFromUserAvatarUrl(v string) {
	o.FromUserAvatarUrl = &v
}

// GetFromusername returns the Fromusername field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetFromusername() string {
	if o == nil || o.Fromusername == nil {
		var ret string
		return ret
	}
	return *o.Fromusername
}

// GetFromusernameOk returns a tuple with the Fromusername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetFromusernameOk() (*string, bool) {
	if o == nil || o.Fromusername == nil {
		return nil, false
	}
	return o.Fromusername, true
}

// HasFromusername returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasFromusername() bool {
	if o != nil && o.Fromusername != nil {
		return true
	}

	return false
}

// SetFromusername gets a reference to the given string and assigns it to the Fromusername field.
func (o *InlineResponse200122Activity) SetFromusername(v string) {
	o.Fromusername = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200122Activity) SetId(v string) {
	o.Id = &v
}

// GetIsprivate returns the Isprivate field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetIsprivate() string {
	if o == nil || o.Isprivate == nil {
		var ret string
		return ret
	}
	return *o.Isprivate
}

// GetIsprivateOk returns a tuple with the Isprivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetIsprivateOk() (*string, bool) {
	if o == nil || o.Isprivate == nil {
		return nil, false
	}
	return o.Isprivate, true
}

// HasIsprivate returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasIsprivate() bool {
	if o != nil && o.Isprivate != nil {
		return true
	}

	return false
}

// SetIsprivate gets a reference to the given string and assigns it to the Isprivate field.
func (o *InlineResponse200122Activity) SetIsprivate(v string) {
	o.Isprivate = &v
}

// GetItemid returns the Itemid field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetItemid() string {
	if o == nil || o.Itemid == nil {
		var ret string
		return ret
	}
	return *o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetItemidOk() (*string, bool) {
	if o == nil || o.Itemid == nil {
		return nil, false
	}
	return o.Itemid, true
}

// HasItemid returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasItemid() bool {
	if o != nil && o.Itemid != nil {
		return true
	}

	return false
}

// SetItemid gets a reference to the given string and assigns it to the Itemid field.
func (o *InlineResponse200122Activity) SetItemid(v string) {
	o.Itemid = &v
}

// GetItemlink returns the Itemlink field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetItemlink() string {
	if o == nil || o.Itemlink == nil {
		var ret string
		return ret
	}
	return *o.Itemlink
}

// GetItemlinkOk returns a tuple with the Itemlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetItemlinkOk() (*string, bool) {
	if o == nil || o.Itemlink == nil {
		return nil, false
	}
	return o.Itemlink, true
}

// HasItemlink returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasItemlink() bool {
	if o != nil && o.Itemlink != nil {
		return true
	}

	return false
}

// SetItemlink gets a reference to the given string and assigns it to the Itemlink field.
func (o *InlineResponse200122Activity) SetItemlink(v string) {
	o.Itemlink = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *InlineResponse200122Activity) SetLink(v string) {
	o.Link = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InlineResponse200122Activity) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPublicinfo returns the Publicinfo field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetPublicinfo() string {
	if o == nil || o.Publicinfo == nil {
		var ret string
		return ret
	}
	return *o.Publicinfo
}

// GetPublicinfoOk returns a tuple with the Publicinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetPublicinfoOk() (*string, bool) {
	if o == nil || o.Publicinfo == nil {
		return nil, false
	}
	return o.Publicinfo, true
}

// HasPublicinfo returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasPublicinfo() bool {
	if o != nil && o.Publicinfo != nil {
		return true
	}

	return false
}

// SetPublicinfo gets a reference to the given string and assigns it to the Publicinfo field.
func (o *InlineResponse200122Activity) SetPublicinfo(v string) {
	o.Publicinfo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200122Activity) SetType(v string) {
	o.Type = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *InlineResponse200122Activity) GetUserid() string {
	if o == nil || o.Userid == nil {
		var ret string
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200122Activity) GetUseridOk() (*string, bool) {
	if o == nil || o.Userid == nil {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *InlineResponse200122Activity) HasUserid() bool {
	if o != nil && o.Userid != nil {
		return true
	}

	return false
}

// SetUserid gets a reference to the given string and assigns it to the Userid field.
func (o *InlineResponse200122Activity) SetUserid(v string) {
	o.Userid = &v
}

func (o InlineResponse200122Activity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activitytype != nil {
		toSerialize["activitytype"] = o.Activitytype
	}
	if o.Datetime != nil {
		toSerialize["datetime"] = o.Datetime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DueDate != nil {
		toSerialize["due-date"] = o.DueDate
	}
	if o.Extradescription != nil {
		toSerialize["extradescription"] = o.Extradescription
	}
	if o.ForUserAvatarUrl != nil {
		toSerialize["for-user-avatar-url"] = o.ForUserAvatarUrl
	}
	if o.Foruserid != nil {
		toSerialize["foruserid"] = o.Foruserid
	}
	if o.Forusername != nil {
		toSerialize["forusername"] = o.Forusername
	}
	if o.FromUserAvatarUrl != nil {
		toSerialize["from-user-avatar-url"] = o.FromUserAvatarUrl
	}
	if o.Fromusername != nil {
		toSerialize["fromusername"] = o.Fromusername
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Isprivate != nil {
		toSerialize["isprivate"] = o.Isprivate
	}
	if o.Itemid != nil {
		toSerialize["itemid"] = o.Itemid
	}
	if o.Itemlink != nil {
		toSerialize["itemlink"] = o.Itemlink
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.ProjectId != nil {
		toSerialize["project-id"] = o.ProjectId
	}
	if o.Publicinfo != nil {
		toSerialize["publicinfo"] = o.Publicinfo
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Userid != nil {
		toSerialize["userid"] = o.Userid
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200122Activity struct {
	value *InlineResponse200122Activity
	isSet bool
}

func (v NullableInlineResponse200122Activity) Get() *InlineResponse200122Activity {
	return v.value
}

func (v *NullableInlineResponse200122Activity) Set(val *InlineResponse200122Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200122Activity) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200122Activity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200122Activity(val *InlineResponse200122Activity) *NullableInlineResponse200122Activity {
	return &NullableInlineResponse200122Activity{value: val, isSet: true}
}

func (v NullableInlineResponse200122Activity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200122Activity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


