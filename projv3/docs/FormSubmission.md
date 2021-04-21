# FormSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewFormSubmission

`func NewFormSubmission() *FormSubmission`

NewFormSubmission instantiates a new FormSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmissionWithDefaults

`func NewFormSubmissionWithDefaults() *FormSubmission`

NewFormSubmissionWithDefaults instantiates a new FormSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FormSubmission) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FormSubmission) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FormSubmission) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FormSubmission) HasData() bool`

HasData returns a boolean if a field has been set.

### GetVersion

`func (o *FormSubmission) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormSubmission) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormSubmission) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FormSubmission) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


