# InlineResponse20045CardUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedDate** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**InlineResponse20045CardOwner**](InlineResponse20045CardOwner.md) |  | [optional] 

## Methods

### NewInlineResponse20045CardUpdate

`func NewInlineResponse20045CardUpdate() *InlineResponse20045CardUpdate`

NewInlineResponse20045CardUpdate instantiates a new InlineResponse20045CardUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045CardUpdateWithDefaults

`func NewInlineResponse20045CardUpdateWithDefaults() *InlineResponse20045CardUpdate`

NewInlineResponse20045CardUpdateWithDefaults instantiates a new InlineResponse20045CardUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *InlineResponse20045CardUpdate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InlineResponse20045CardUpdate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InlineResponse20045CardUpdate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *InlineResponse20045CardUpdate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20045CardUpdate) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20045CardUpdate) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20045CardUpdate) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20045CardUpdate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse20045CardUpdate) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse20045CardUpdate) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse20045CardUpdate) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse20045CardUpdate) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedDate

`func (o *InlineResponse20045CardUpdate) GetDeletedDate() string`

GetDeletedDate returns the DeletedDate field if non-nil, zero value otherwise.

### GetDeletedDateOk

`func (o *InlineResponse20045CardUpdate) GetDeletedDateOk() (*string, bool)`

GetDeletedDateOk returns a tuple with the DeletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDate

`func (o *InlineResponse20045CardUpdate) SetDeletedDate(v string)`

SetDeletedDate sets DeletedDate field to given value.

### HasDeletedDate

`func (o *InlineResponse20045CardUpdate) HasDeletedDate() bool`

HasDeletedDate returns a boolean if a field has been set.

### GetHealth

`func (o *InlineResponse20045CardUpdate) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *InlineResponse20045CardUpdate) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *InlineResponse20045CardUpdate) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *InlineResponse20045CardUpdate) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20045CardUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20045CardUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20045CardUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20045CardUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *InlineResponse20045CardUpdate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse20045CardUpdate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse20045CardUpdate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse20045CardUpdate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetText

`func (o *InlineResponse20045CardUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InlineResponse20045CardUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InlineResponse20045CardUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *InlineResponse20045CardUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse20045CardUpdate) GetUser() InlineResponse20045CardOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse20045CardUpdate) GetUserOk() (*InlineResponse20045CardOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse20045CardUpdate) SetUser(v InlineResponse20045CardOwner)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse20045CardUpdate) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


