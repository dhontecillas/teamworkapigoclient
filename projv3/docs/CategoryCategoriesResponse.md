# CategoryCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]ViewProjectCategory**](ViewProjectCategory.md) |  | [optional] 

## Methods

### NewCategoryCategoriesResponse

`func NewCategoryCategoriesResponse() *CategoryCategoriesResponse`

NewCategoryCategoriesResponse instantiates a new CategoryCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryCategoriesResponseWithDefaults

`func NewCategoryCategoriesResponseWithDefaults() *CategoryCategoriesResponse`

NewCategoryCategoriesResponseWithDefaults instantiates a new CategoryCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *CategoryCategoriesResponse) GetProjects() []ViewProjectCategory`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CategoryCategoriesResponse) GetProjectsOk() (*[]ViewProjectCategory, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CategoryCategoriesResponse) SetProjects(v []ViewProjectCategory)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CategoryCategoriesResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


