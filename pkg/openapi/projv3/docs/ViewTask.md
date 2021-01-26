# ViewTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeCompanies** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**AssigneeCompanyIDs** | Pointer to **[]int32** |  | [optional] 
**AssigneeTeamIDs** | Pointer to **[]int32** |  | [optional] 
**AssigneeTeams** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**AssigneeUserIDs** | Pointer to **[]int32** |  | [optional] 
**AssigneeUsers** | Pointer to [**[]ViewRelationship**](ViewRelationship.md) |  | [optional] 
**CreatedBy** | Pointer to **int32** |  | [optional] 
**CreatedByUserID** | Pointer to **int32** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EstimateMinutes** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsPrivate** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentTask** | Pointer to [**ViewRelationship**](view.Relationship.md) |  | [optional] 
**ParentTaskId** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tasklist** | Pointer to [**ViewRelationship**](view.Relationship.md) |  | [optional] 
**TasklistId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserPermissions** | Pointer to [**ViewTaskPermissions**](view.TaskPermissions.md) |  | [optional] 

## Methods

### NewViewTask

`func NewViewTask() *ViewTask`

NewViewTask instantiates a new ViewTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTaskWithDefaults

`func NewViewTaskWithDefaults() *ViewTask`

NewViewTaskWithDefaults instantiates a new ViewTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeCompanies

`func (o *ViewTask) GetAssigneeCompanies() []ViewRelationship`

GetAssigneeCompanies returns the AssigneeCompanies field if non-nil, zero value otherwise.

### GetAssigneeCompaniesOk

`func (o *ViewTask) GetAssigneeCompaniesOk() (*[]ViewRelationship, bool)`

GetAssigneeCompaniesOk returns a tuple with the AssigneeCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanies

`func (o *ViewTask) SetAssigneeCompanies(v []ViewRelationship)`

SetAssigneeCompanies sets AssigneeCompanies field to given value.

### HasAssigneeCompanies

`func (o *ViewTask) HasAssigneeCompanies() bool`

HasAssigneeCompanies returns a boolean if a field has been set.

### GetAssigneeCompanyIDs

`func (o *ViewTask) GetAssigneeCompanyIDs() []int32`

GetAssigneeCompanyIDs returns the AssigneeCompanyIDs field if non-nil, zero value otherwise.

### GetAssigneeCompanyIDsOk

`func (o *ViewTask) GetAssigneeCompanyIDsOk() (*[]int32, bool)`

GetAssigneeCompanyIDsOk returns a tuple with the AssigneeCompanyIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanyIDs

`func (o *ViewTask) SetAssigneeCompanyIDs(v []int32)`

SetAssigneeCompanyIDs sets AssigneeCompanyIDs field to given value.

### HasAssigneeCompanyIDs

`func (o *ViewTask) HasAssigneeCompanyIDs() bool`

HasAssigneeCompanyIDs returns a boolean if a field has been set.

### GetAssigneeTeamIDs

`func (o *ViewTask) GetAssigneeTeamIDs() []int32`

GetAssigneeTeamIDs returns the AssigneeTeamIDs field if non-nil, zero value otherwise.

### GetAssigneeTeamIDsOk

`func (o *ViewTask) GetAssigneeTeamIDsOk() (*[]int32, bool)`

GetAssigneeTeamIDsOk returns a tuple with the AssigneeTeamIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeTeamIDs

`func (o *ViewTask) SetAssigneeTeamIDs(v []int32)`

SetAssigneeTeamIDs sets AssigneeTeamIDs field to given value.

### HasAssigneeTeamIDs

`func (o *ViewTask) HasAssigneeTeamIDs() bool`

HasAssigneeTeamIDs returns a boolean if a field has been set.

### GetAssigneeTeams

`func (o *ViewTask) GetAssigneeTeams() []ViewRelationship`

GetAssigneeTeams returns the AssigneeTeams field if non-nil, zero value otherwise.

### GetAssigneeTeamsOk

`func (o *ViewTask) GetAssigneeTeamsOk() (*[]ViewRelationship, bool)`

GetAssigneeTeamsOk returns a tuple with the AssigneeTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeTeams

`func (o *ViewTask) SetAssigneeTeams(v []ViewRelationship)`

SetAssigneeTeams sets AssigneeTeams field to given value.

### HasAssigneeTeams

`func (o *ViewTask) HasAssigneeTeams() bool`

HasAssigneeTeams returns a boolean if a field has been set.

### GetAssigneeUserIDs

`func (o *ViewTask) GetAssigneeUserIDs() []int32`

GetAssigneeUserIDs returns the AssigneeUserIDs field if non-nil, zero value otherwise.

### GetAssigneeUserIDsOk

`func (o *ViewTask) GetAssigneeUserIDsOk() (*[]int32, bool)`

GetAssigneeUserIDsOk returns a tuple with the AssigneeUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUserIDs

`func (o *ViewTask) SetAssigneeUserIDs(v []int32)`

SetAssigneeUserIDs sets AssigneeUserIDs field to given value.

### HasAssigneeUserIDs

`func (o *ViewTask) HasAssigneeUserIDs() bool`

HasAssigneeUserIDs returns a boolean if a field has been set.

### GetAssigneeUsers

`func (o *ViewTask) GetAssigneeUsers() []ViewRelationship`

GetAssigneeUsers returns the AssigneeUsers field if non-nil, zero value otherwise.

### GetAssigneeUsersOk

`func (o *ViewTask) GetAssigneeUsersOk() (*[]ViewRelationship, bool)`

GetAssigneeUsersOk returns a tuple with the AssigneeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUsers

`func (o *ViewTask) SetAssigneeUsers(v []ViewRelationship)`

SetAssigneeUsers sets AssigneeUsers field to given value.

### HasAssigneeUsers

`func (o *ViewTask) HasAssigneeUsers() bool`

HasAssigneeUsers returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ViewTask) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ViewTask) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ViewTask) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ViewTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserID

`func (o *ViewTask) GetCreatedByUserID() int32`

GetCreatedByUserID returns the CreatedByUserID field if non-nil, zero value otherwise.

### GetCreatedByUserIDOk

`func (o *ViewTask) GetCreatedByUserIDOk() (*int32, bool)`

GetCreatedByUserIDOk returns a tuple with the CreatedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserID

`func (o *ViewTask) SetCreatedByUserID(v int32)`

SetCreatedByUserID sets CreatedByUserID field to given value.

### HasCreatedByUserID

`func (o *ViewTask) HasCreatedByUserID() bool`

HasCreatedByUserID returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ViewTask) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ViewTask) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ViewTask) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ViewTask) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDueDate

`func (o *ViewTask) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ViewTask) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ViewTask) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ViewTask) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEstimateMinutes

`func (o *ViewTask) GetEstimateMinutes() int32`

GetEstimateMinutes returns the EstimateMinutes field if non-nil, zero value otherwise.

### GetEstimateMinutesOk

`func (o *ViewTask) GetEstimateMinutesOk() (*int32, bool)`

GetEstimateMinutesOk returns a tuple with the EstimateMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateMinutes

`func (o *ViewTask) SetEstimateMinutes(v int32)`

SetEstimateMinutes sets EstimateMinutes field to given value.

### HasEstimateMinutes

`func (o *ViewTask) HasEstimateMinutes() bool`

HasEstimateMinutes returns a boolean if a field has been set.

### GetId

`func (o *ViewTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ViewTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ViewTask) GetIsPrivate() int32`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ViewTask) GetIsPrivateOk() (*int32, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ViewTask) SetIsPrivate(v int32)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ViewTask) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetName

`func (o *ViewTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTask

`func (o *ViewTask) GetParentTask() ViewRelationship`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *ViewTask) GetParentTaskOk() (*ViewRelationship, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *ViewTask) SetParentTask(v ViewRelationship)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *ViewTask) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### GetParentTaskId

`func (o *ViewTask) GetParentTaskId() int32`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *ViewTask) GetParentTaskIdOk() (*int32, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *ViewTask) SetParentTaskId(v int32)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *ViewTask) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetStartDate

`func (o *ViewTask) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ViewTask) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ViewTask) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ViewTask) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *ViewTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ViewTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasklist

`func (o *ViewTask) GetTasklist() ViewRelationship`

GetTasklist returns the Tasklist field if non-nil, zero value otherwise.

### GetTasklistOk

`func (o *ViewTask) GetTasklistOk() (*ViewRelationship, bool)`

GetTasklistOk returns a tuple with the Tasklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklist

`func (o *ViewTask) SetTasklist(v ViewRelationship)`

SetTasklist sets Tasklist field to given value.

### HasTasklist

`func (o *ViewTask) HasTasklist() bool`

HasTasklist returns a boolean if a field has been set.

### GetTasklistId

`func (o *ViewTask) GetTasklistId() int32`

GetTasklistId returns the TasklistId field if non-nil, zero value otherwise.

### GetTasklistIdOk

`func (o *ViewTask) GetTasklistIdOk() (*int32, bool)`

GetTasklistIdOk returns a tuple with the TasklistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasklistId

`func (o *ViewTask) SetTasklistId(v int32)`

SetTasklistId sets TasklistId field to given value.

### HasTasklistId

`func (o *ViewTask) HasTasklistId() bool`

HasTasklistId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ViewTask) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ViewTask) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ViewTask) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ViewTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserPermissions

`func (o *ViewTask) GetUserPermissions() ViewTaskPermissions`

GetUserPermissions returns the UserPermissions field if non-nil, zero value otherwise.

### GetUserPermissionsOk

`func (o *ViewTask) GetUserPermissionsOk() (*ViewTaskPermissions, bool)`

GetUserPermissionsOk returns a tuple with the UserPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPermissions

`func (o *ViewTask) SetUserPermissions(v ViewTaskPermissions)`

SetUserPermissions sets UserPermissions field to given value.

### HasUserPermissions

`func (o *ViewTask) HasUserPermissions() bool`

HasUserPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


