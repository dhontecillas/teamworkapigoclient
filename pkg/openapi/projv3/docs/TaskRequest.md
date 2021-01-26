# TaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomfieldTask** | Pointer to [**TaskCustomFieldTask**](task.CustomFieldTask.md) |  | [optional] 

## Methods

### NewTaskRequest

`func NewTaskRequest() *TaskRequest`

NewTaskRequest instantiates a new TaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRequestWithDefaults

`func NewTaskRequestWithDefaults() *TaskRequest`

NewTaskRequestWithDefaults instantiates a new TaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomfieldTask

`func (o *TaskRequest) GetCustomfieldTask() TaskCustomFieldTask`

GetCustomfieldTask returns the CustomfieldTask field if non-nil, zero value otherwise.

### GetCustomfieldTaskOk

`func (o *TaskRequest) GetCustomfieldTaskOk() (*TaskCustomFieldTask, bool)`

GetCustomfieldTaskOk returns a tuple with the CustomfieldTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomfieldTask

`func (o *TaskRequest) SetCustomfieldTask(v TaskCustomFieldTask)`

SetCustomfieldTask sets CustomfieldTask field to given value.

### HasCustomfieldTask

`func (o *TaskRequest) HasCustomfieldTask() bool`

HasCustomfieldTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


