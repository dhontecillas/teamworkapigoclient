# TaskRequestTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EveryoneMustDo** | Pointer to **bool** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**Notify** | Pointer to **bool** |  | [optional] 
**ParseInlineTags** | Pointer to **bool** |  | [optional] 
**PositionAfterTaskId** | Pointer to **int32** |  | [optional] 
**ShiftProjectDates** | Pointer to **bool** |  | [optional] 
**UseDefaults** | Pointer to **bool** |  | [optional] 
**UseNotifyViaTWIM** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskRequestTaskOptions

`func NewTaskRequestTaskOptions() *TaskRequestTaskOptions`

NewTaskRequestTaskOptions instantiates a new TaskRequestTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRequestTaskOptionsWithDefaults

`func NewTaskRequestTaskOptionsWithDefaults() *TaskRequestTaskOptions`

NewTaskRequestTaskOptionsWithDefaults instantiates a new TaskRequestTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEveryoneMustDo

`func (o *TaskRequestTaskOptions) GetEveryoneMustDo() bool`

GetEveryoneMustDo returns the EveryoneMustDo field if non-nil, zero value otherwise.

### GetEveryoneMustDoOk

`func (o *TaskRequestTaskOptions) GetEveryoneMustDoOk() (*bool, bool)`

GetEveryoneMustDoOk returns a tuple with the EveryoneMustDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneMustDo

`func (o *TaskRequestTaskOptions) SetEveryoneMustDo(v bool)`

SetEveryoneMustDo sets EveryoneMustDo field to given value.

### HasEveryoneMustDo

`func (o *TaskRequestTaskOptions) HasEveryoneMustDo() bool`

HasEveryoneMustDo returns a boolean if a field has been set.

### GetIsTemplate

`func (o *TaskRequestTaskOptions) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *TaskRequestTaskOptions) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *TaskRequestTaskOptions) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *TaskRequestTaskOptions) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetNotify

`func (o *TaskRequestTaskOptions) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *TaskRequestTaskOptions) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *TaskRequestTaskOptions) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *TaskRequestTaskOptions) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetParseInlineTags

`func (o *TaskRequestTaskOptions) GetParseInlineTags() bool`

GetParseInlineTags returns the ParseInlineTags field if non-nil, zero value otherwise.

### GetParseInlineTagsOk

`func (o *TaskRequestTaskOptions) GetParseInlineTagsOk() (*bool, bool)`

GetParseInlineTagsOk returns a tuple with the ParseInlineTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseInlineTags

`func (o *TaskRequestTaskOptions) SetParseInlineTags(v bool)`

SetParseInlineTags sets ParseInlineTags field to given value.

### HasParseInlineTags

`func (o *TaskRequestTaskOptions) HasParseInlineTags() bool`

HasParseInlineTags returns a boolean if a field has been set.

### GetPositionAfterTaskId

`func (o *TaskRequestTaskOptions) GetPositionAfterTaskId() int32`

GetPositionAfterTaskId returns the PositionAfterTaskId field if non-nil, zero value otherwise.

### GetPositionAfterTaskIdOk

`func (o *TaskRequestTaskOptions) GetPositionAfterTaskIdOk() (*int32, bool)`

GetPositionAfterTaskIdOk returns a tuple with the PositionAfterTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAfterTaskId

`func (o *TaskRequestTaskOptions) SetPositionAfterTaskId(v int32)`

SetPositionAfterTaskId sets PositionAfterTaskId field to given value.

### HasPositionAfterTaskId

`func (o *TaskRequestTaskOptions) HasPositionAfterTaskId() bool`

HasPositionAfterTaskId returns a boolean if a field has been set.

### GetShiftProjectDates

`func (o *TaskRequestTaskOptions) GetShiftProjectDates() bool`

GetShiftProjectDates returns the ShiftProjectDates field if non-nil, zero value otherwise.

### GetShiftProjectDatesOk

`func (o *TaskRequestTaskOptions) GetShiftProjectDatesOk() (*bool, bool)`

GetShiftProjectDatesOk returns a tuple with the ShiftProjectDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftProjectDates

`func (o *TaskRequestTaskOptions) SetShiftProjectDates(v bool)`

SetShiftProjectDates sets ShiftProjectDates field to given value.

### HasShiftProjectDates

`func (o *TaskRequestTaskOptions) HasShiftProjectDates() bool`

HasShiftProjectDates returns a boolean if a field has been set.

### GetUseDefaults

`func (o *TaskRequestTaskOptions) GetUseDefaults() bool`

GetUseDefaults returns the UseDefaults field if non-nil, zero value otherwise.

### GetUseDefaultsOk

`func (o *TaskRequestTaskOptions) GetUseDefaultsOk() (*bool, bool)`

GetUseDefaultsOk returns a tuple with the UseDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaults

`func (o *TaskRequestTaskOptions) SetUseDefaults(v bool)`

SetUseDefaults sets UseDefaults field to given value.

### HasUseDefaults

`func (o *TaskRequestTaskOptions) HasUseDefaults() bool`

HasUseDefaults returns a boolean if a field has been set.

### GetUseNotifyViaTWIM

`func (o *TaskRequestTaskOptions) GetUseNotifyViaTWIM() bool`

GetUseNotifyViaTWIM returns the UseNotifyViaTWIM field if non-nil, zero value otherwise.

### GetUseNotifyViaTWIMOk

`func (o *TaskRequestTaskOptions) GetUseNotifyViaTWIMOk() (*bool, bool)`

GetUseNotifyViaTWIMOk returns a tuple with the UseNotifyViaTWIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotifyViaTWIM

`func (o *TaskRequestTaskOptions) SetUseNotifyViaTWIM(v bool)`

SetUseNotifyViaTWIM sets UseNotifyViaTWIM field to given value.

### HasUseNotifyViaTWIM

`func (o *TaskRequestTaskOptions) HasUseNotifyViaTWIM() bool`

HasUseNotifyViaTWIM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


