# InlineResponse20084People

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**InlineResponse20013Address**](InlineResponse20013Address.md) |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressCountry** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**AddressState** | Pointer to **string** |  | [optional] 
**AddressZip** | Pointer to **string** |  | [optional] 
**Administrator** | Pointer to **bool** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailAlt1** | Pointer to **string** |  | [optional] 
**EmailAlt2** | Pointer to **string** |  | [optional] 
**EmailAlt3** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HasAccessToNewProjects** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImHandle** | Pointer to **string** |  | [optional] 
**ImService** | Pointer to **string** |  | [optional] 
**InOwnerCompany** | Pointer to **bool** |  | [optional] 
**LastActive** | Pointer to **string** |  | [optional] 
**LastChangedOn** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LengthOfDay** | Pointer to **string** |  | [optional] 
**LoginCount** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**OpenId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**InlineResponse20013Permissions**](InlineResponse20013Permissions.md) |  | [optional] 
**PhoneNumberFax** | Pointer to **string** |  | [optional] 
**PhoneNumberHome** | Pointer to **string** |  | [optional] 
**PhoneNumberMobile** | Pointer to **string** |  | [optional] 
**PhoneNumberMobileParts** | Pointer to [**InlineResponse20013PhoneNumberMobileParts**](InlineResponse20013PhoneNumberMobileParts.md) |  | [optional] 
**PhoneNumberOffice** | Pointer to **string** |  | [optional] 
**PhoneNumberOfficeExt** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ProfileText** | Pointer to **string** |  | [optional] 
**SiteOwner** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TextFormat** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**UseShorthandDurations** | Pointer to **bool** |  | [optional] 
**UserInvited** | Pointer to **string** |  | [optional] 
**UserInvitedDate** | Pointer to **string** |  | [optional] 
**UserInvitedStatus** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserUUID** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20084People

`func NewInlineResponse20084People() *InlineResponse20084People`

NewInlineResponse20084People instantiates a new InlineResponse20084People object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20084PeopleWithDefaults

`func NewInlineResponse20084PeopleWithDefaults() *InlineResponse20084People`

NewInlineResponse20084PeopleWithDefaults instantiates a new InlineResponse20084People object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InlineResponse20084People) GetAddress() InlineResponse20013Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20084People) GetAddressOk() (*InlineResponse20013Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20084People) SetAddress(v InlineResponse20013Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse20084People) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressCity

`func (o *InlineResponse20084People) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *InlineResponse20084People) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *InlineResponse20084People) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *InlineResponse20084People) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *InlineResponse20084People) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *InlineResponse20084People) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *InlineResponse20084People) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *InlineResponse20084People) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressLine1

`func (o *InlineResponse20084People) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *InlineResponse20084People) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *InlineResponse20084People) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *InlineResponse20084People) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *InlineResponse20084People) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *InlineResponse20084People) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *InlineResponse20084People) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *InlineResponse20084People) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressState

`func (o *InlineResponse20084People) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *InlineResponse20084People) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *InlineResponse20084People) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *InlineResponse20084People) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetAddressZip

`func (o *InlineResponse20084People) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *InlineResponse20084People) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *InlineResponse20084People) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *InlineResponse20084People) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### GetAdministrator

`func (o *InlineResponse20084People) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *InlineResponse20084People) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *InlineResponse20084People) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *InlineResponse20084People) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *InlineResponse20084People) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *InlineResponse20084People) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *InlineResponse20084People) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *InlineResponse20084People) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCompanyId

`func (o *InlineResponse20084People) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InlineResponse20084People) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InlineResponse20084People) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InlineResponse20084People) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *InlineResponse20084People) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineResponse20084People) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineResponse20084People) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineResponse20084People) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20084People) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20084People) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20084People) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20084People) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *InlineResponse20084People) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InlineResponse20084People) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InlineResponse20084People) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *InlineResponse20084People) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InlineResponse20084People) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InlineResponse20084People) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InlineResponse20084People) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InlineResponse20084People) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAlt1

`func (o *InlineResponse20084People) GetEmailAlt1() string`

GetEmailAlt1 returns the EmailAlt1 field if non-nil, zero value otherwise.

### GetEmailAlt1Ok

`func (o *InlineResponse20084People) GetEmailAlt1Ok() (*string, bool)`

GetEmailAlt1Ok returns a tuple with the EmailAlt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt1

`func (o *InlineResponse20084People) SetEmailAlt1(v string)`

SetEmailAlt1 sets EmailAlt1 field to given value.

### HasEmailAlt1

`func (o *InlineResponse20084People) HasEmailAlt1() bool`

HasEmailAlt1 returns a boolean if a field has been set.

### GetEmailAlt2

`func (o *InlineResponse20084People) GetEmailAlt2() string`

GetEmailAlt2 returns the EmailAlt2 field if non-nil, zero value otherwise.

### GetEmailAlt2Ok

`func (o *InlineResponse20084People) GetEmailAlt2Ok() (*string, bool)`

GetEmailAlt2Ok returns a tuple with the EmailAlt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt2

`func (o *InlineResponse20084People) SetEmailAlt2(v string)`

SetEmailAlt2 sets EmailAlt2 field to given value.

### HasEmailAlt2

`func (o *InlineResponse20084People) HasEmailAlt2() bool`

HasEmailAlt2 returns a boolean if a field has been set.

### GetEmailAlt3

`func (o *InlineResponse20084People) GetEmailAlt3() string`

GetEmailAlt3 returns the EmailAlt3 field if non-nil, zero value otherwise.

### GetEmailAlt3Ok

`func (o *InlineResponse20084People) GetEmailAlt3Ok() (*string, bool)`

GetEmailAlt3Ok returns a tuple with the EmailAlt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlt3

`func (o *InlineResponse20084People) SetEmailAlt3(v string)`

SetEmailAlt3 sets EmailAlt3 field to given value.

### HasEmailAlt3

`func (o *InlineResponse20084People) HasEmailAlt3() bool`

HasEmailAlt3 returns a boolean if a field has been set.

### GetFirstName

`func (o *InlineResponse20084People) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineResponse20084People) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineResponse20084People) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InlineResponse20084People) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHasAccessToNewProjects

`func (o *InlineResponse20084People) GetHasAccessToNewProjects() bool`

GetHasAccessToNewProjects returns the HasAccessToNewProjects field if non-nil, zero value otherwise.

### GetHasAccessToNewProjectsOk

`func (o *InlineResponse20084People) GetHasAccessToNewProjectsOk() (*bool, bool)`

GetHasAccessToNewProjectsOk returns a tuple with the HasAccessToNewProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewProjects

`func (o *InlineResponse20084People) SetHasAccessToNewProjects(v bool)`

SetHasAccessToNewProjects sets HasAccessToNewProjects field to given value.

### HasHasAccessToNewProjects

`func (o *InlineResponse20084People) HasHasAccessToNewProjects() bool`

HasHasAccessToNewProjects returns a boolean if a field has been set.

### GetId

`func (o *InlineResponse20084People) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20084People) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20084People) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20084People) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImHandle

`func (o *InlineResponse20084People) GetImHandle() string`

GetImHandle returns the ImHandle field if non-nil, zero value otherwise.

### GetImHandleOk

`func (o *InlineResponse20084People) GetImHandleOk() (*string, bool)`

GetImHandleOk returns a tuple with the ImHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImHandle

`func (o *InlineResponse20084People) SetImHandle(v string)`

SetImHandle sets ImHandle field to given value.

### HasImHandle

`func (o *InlineResponse20084People) HasImHandle() bool`

HasImHandle returns a boolean if a field has been set.

### GetImService

`func (o *InlineResponse20084People) GetImService() string`

GetImService returns the ImService field if non-nil, zero value otherwise.

### GetImServiceOk

`func (o *InlineResponse20084People) GetImServiceOk() (*string, bool)`

GetImServiceOk returns a tuple with the ImService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImService

`func (o *InlineResponse20084People) SetImService(v string)`

SetImService sets ImService field to given value.

### HasImService

`func (o *InlineResponse20084People) HasImService() bool`

HasImService returns a boolean if a field has been set.

### GetInOwnerCompany

`func (o *InlineResponse20084People) GetInOwnerCompany() bool`

GetInOwnerCompany returns the InOwnerCompany field if non-nil, zero value otherwise.

### GetInOwnerCompanyOk

`func (o *InlineResponse20084People) GetInOwnerCompanyOk() (*bool, bool)`

GetInOwnerCompanyOk returns a tuple with the InOwnerCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOwnerCompany

`func (o *InlineResponse20084People) SetInOwnerCompany(v bool)`

SetInOwnerCompany sets InOwnerCompany field to given value.

### HasInOwnerCompany

`func (o *InlineResponse20084People) HasInOwnerCompany() bool`

HasInOwnerCompany returns a boolean if a field has been set.

### GetLastActive

`func (o *InlineResponse20084People) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *InlineResponse20084People) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *InlineResponse20084People) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *InlineResponse20084People) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetLastChangedOn

`func (o *InlineResponse20084People) GetLastChangedOn() string`

GetLastChangedOn returns the LastChangedOn field if non-nil, zero value otherwise.

### GetLastChangedOnOk

`func (o *InlineResponse20084People) GetLastChangedOnOk() (*string, bool)`

GetLastChangedOnOk returns a tuple with the LastChangedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangedOn

`func (o *InlineResponse20084People) SetLastChangedOn(v string)`

SetLastChangedOn sets LastChangedOn field to given value.

### HasLastChangedOn

`func (o *InlineResponse20084People) HasLastChangedOn() bool`

HasLastChangedOn returns a boolean if a field has been set.

### GetLastLogin

`func (o *InlineResponse20084People) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *InlineResponse20084People) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *InlineResponse20084People) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *InlineResponse20084People) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *InlineResponse20084People) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineResponse20084People) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineResponse20084People) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InlineResponse20084People) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLengthOfDay

`func (o *InlineResponse20084People) GetLengthOfDay() string`

GetLengthOfDay returns the LengthOfDay field if non-nil, zero value otherwise.

### GetLengthOfDayOk

`func (o *InlineResponse20084People) GetLengthOfDayOk() (*string, bool)`

GetLengthOfDayOk returns a tuple with the LengthOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthOfDay

`func (o *InlineResponse20084People) SetLengthOfDay(v string)`

SetLengthOfDay sets LengthOfDay field to given value.

### HasLengthOfDay

`func (o *InlineResponse20084People) HasLengthOfDay() bool`

HasLengthOfDay returns a boolean if a field has been set.

### GetLoginCount

`func (o *InlineResponse20084People) GetLoginCount() string`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *InlineResponse20084People) GetLoginCountOk() (*string, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *InlineResponse20084People) SetLoginCount(v string)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *InlineResponse20084People) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20084People) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20084People) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20084People) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20084People) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOpenId

`func (o *InlineResponse20084People) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *InlineResponse20084People) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *InlineResponse20084People) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *InlineResponse20084People) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse20084People) GetPermissions() InlineResponse20013Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse20084People) GetPermissionsOk() (*InlineResponse20013Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse20084People) SetPermissions(v InlineResponse20013Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse20084People) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhoneNumberFax

`func (o *InlineResponse20084People) GetPhoneNumberFax() string`

GetPhoneNumberFax returns the PhoneNumberFax field if non-nil, zero value otherwise.

### GetPhoneNumberFaxOk

`func (o *InlineResponse20084People) GetPhoneNumberFaxOk() (*string, bool)`

GetPhoneNumberFaxOk returns a tuple with the PhoneNumberFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberFax

`func (o *InlineResponse20084People) SetPhoneNumberFax(v string)`

SetPhoneNumberFax sets PhoneNumberFax field to given value.

### HasPhoneNumberFax

`func (o *InlineResponse20084People) HasPhoneNumberFax() bool`

HasPhoneNumberFax returns a boolean if a field has been set.

### GetPhoneNumberHome

`func (o *InlineResponse20084People) GetPhoneNumberHome() string`

GetPhoneNumberHome returns the PhoneNumberHome field if non-nil, zero value otherwise.

### GetPhoneNumberHomeOk

`func (o *InlineResponse20084People) GetPhoneNumberHomeOk() (*string, bool)`

GetPhoneNumberHomeOk returns a tuple with the PhoneNumberHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberHome

`func (o *InlineResponse20084People) SetPhoneNumberHome(v string)`

SetPhoneNumberHome sets PhoneNumberHome field to given value.

### HasPhoneNumberHome

`func (o *InlineResponse20084People) HasPhoneNumberHome() bool`

HasPhoneNumberHome returns a boolean if a field has been set.

### GetPhoneNumberMobile

`func (o *InlineResponse20084People) GetPhoneNumberMobile() string`

GetPhoneNumberMobile returns the PhoneNumberMobile field if non-nil, zero value otherwise.

### GetPhoneNumberMobileOk

`func (o *InlineResponse20084People) GetPhoneNumberMobileOk() (*string, bool)`

GetPhoneNumberMobileOk returns a tuple with the PhoneNumberMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobile

`func (o *InlineResponse20084People) SetPhoneNumberMobile(v string)`

SetPhoneNumberMobile sets PhoneNumberMobile field to given value.

### HasPhoneNumberMobile

`func (o *InlineResponse20084People) HasPhoneNumberMobile() bool`

HasPhoneNumberMobile returns a boolean if a field has been set.

### GetPhoneNumberMobileParts

`func (o *InlineResponse20084People) GetPhoneNumberMobileParts() InlineResponse20013PhoneNumberMobileParts`

GetPhoneNumberMobileParts returns the PhoneNumberMobileParts field if non-nil, zero value otherwise.

### GetPhoneNumberMobilePartsOk

`func (o *InlineResponse20084People) GetPhoneNumberMobilePartsOk() (*InlineResponse20013PhoneNumberMobileParts, bool)`

GetPhoneNumberMobilePartsOk returns a tuple with the PhoneNumberMobileParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberMobileParts

`func (o *InlineResponse20084People) SetPhoneNumberMobileParts(v InlineResponse20013PhoneNumberMobileParts)`

SetPhoneNumberMobileParts sets PhoneNumberMobileParts field to given value.

### HasPhoneNumberMobileParts

`func (o *InlineResponse20084People) HasPhoneNumberMobileParts() bool`

HasPhoneNumberMobileParts returns a boolean if a field has been set.

### GetPhoneNumberOffice

`func (o *InlineResponse20084People) GetPhoneNumberOffice() string`

GetPhoneNumberOffice returns the PhoneNumberOffice field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeOk

`func (o *InlineResponse20084People) GetPhoneNumberOfficeOk() (*string, bool)`

GetPhoneNumberOfficeOk returns a tuple with the PhoneNumberOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOffice

`func (o *InlineResponse20084People) SetPhoneNumberOffice(v string)`

SetPhoneNumberOffice sets PhoneNumberOffice field to given value.

### HasPhoneNumberOffice

`func (o *InlineResponse20084People) HasPhoneNumberOffice() bool`

HasPhoneNumberOffice returns a boolean if a field has been set.

### GetPhoneNumberOfficeExt

`func (o *InlineResponse20084People) GetPhoneNumberOfficeExt() string`

GetPhoneNumberOfficeExt returns the PhoneNumberOfficeExt field if non-nil, zero value otherwise.

### GetPhoneNumberOfficeExtOk

`func (o *InlineResponse20084People) GetPhoneNumberOfficeExtOk() (*string, bool)`

GetPhoneNumberOfficeExtOk returns a tuple with the PhoneNumberOfficeExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberOfficeExt

`func (o *InlineResponse20084People) SetPhoneNumberOfficeExt(v string)`

SetPhoneNumberOfficeExt sets PhoneNumberOfficeExt field to given value.

### HasPhoneNumberOfficeExt

`func (o *InlineResponse20084People) HasPhoneNumberOfficeExt() bool`

HasPhoneNumberOfficeExt returns a boolean if a field has been set.

### GetPid

`func (o *InlineResponse20084People) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *InlineResponse20084People) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *InlineResponse20084People) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *InlineResponse20084People) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse20084People) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse20084People) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse20084People) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse20084People) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProfileText

`func (o *InlineResponse20084People) GetProfileText() string`

GetProfileText returns the ProfileText field if non-nil, zero value otherwise.

### GetProfileTextOk

`func (o *InlineResponse20084People) GetProfileTextOk() (*string, bool)`

GetProfileTextOk returns a tuple with the ProfileText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileText

`func (o *InlineResponse20084People) SetProfileText(v string)`

SetProfileText sets ProfileText field to given value.

### HasProfileText

`func (o *InlineResponse20084People) HasProfileText() bool`

HasProfileText returns a boolean if a field has been set.

### GetSiteOwner

`func (o *InlineResponse20084People) GetSiteOwner() bool`

GetSiteOwner returns the SiteOwner field if non-nil, zero value otherwise.

### GetSiteOwnerOk

`func (o *InlineResponse20084People) GetSiteOwnerOk() (*bool, bool)`

GetSiteOwnerOk returns a tuple with the SiteOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwner

`func (o *InlineResponse20084People) SetSiteOwner(v bool)`

SetSiteOwner sets SiteOwner field to given value.

### HasSiteOwner

`func (o *InlineResponse20084People) HasSiteOwner() bool`

HasSiteOwner returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20084People) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20084People) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20084People) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20084People) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTextFormat

`func (o *InlineResponse20084People) GetTextFormat() string`

GetTextFormat returns the TextFormat field if non-nil, zero value otherwise.

### GetTextFormatOk

`func (o *InlineResponse20084People) GetTextFormatOk() (*string, bool)`

GetTextFormatOk returns a tuple with the TextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFormat

`func (o *InlineResponse20084People) SetTextFormat(v string)`

SetTextFormat sets TextFormat field to given value.

### HasTextFormat

`func (o *InlineResponse20084People) HasTextFormat() bool`

HasTextFormat returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse20084People) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse20084People) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse20084People) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse20084People) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwitter

`func (o *InlineResponse20084People) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *InlineResponse20084People) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *InlineResponse20084People) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *InlineResponse20084People) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetUseShorthandDurations

`func (o *InlineResponse20084People) GetUseShorthandDurations() bool`

GetUseShorthandDurations returns the UseShorthandDurations field if non-nil, zero value otherwise.

### GetUseShorthandDurationsOk

`func (o *InlineResponse20084People) GetUseShorthandDurationsOk() (*bool, bool)`

GetUseShorthandDurationsOk returns a tuple with the UseShorthandDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShorthandDurations

`func (o *InlineResponse20084People) SetUseShorthandDurations(v bool)`

SetUseShorthandDurations sets UseShorthandDurations field to given value.

### HasUseShorthandDurations

`func (o *InlineResponse20084People) HasUseShorthandDurations() bool`

HasUseShorthandDurations returns a boolean if a field has been set.

### GetUserInvited

`func (o *InlineResponse20084People) GetUserInvited() string`

GetUserInvited returns the UserInvited field if non-nil, zero value otherwise.

### GetUserInvitedOk

`func (o *InlineResponse20084People) GetUserInvitedOk() (*string, bool)`

GetUserInvitedOk returns a tuple with the UserInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvited

`func (o *InlineResponse20084People) SetUserInvited(v string)`

SetUserInvited sets UserInvited field to given value.

### HasUserInvited

`func (o *InlineResponse20084People) HasUserInvited() bool`

HasUserInvited returns a boolean if a field has been set.

### GetUserInvitedDate

`func (o *InlineResponse20084People) GetUserInvitedDate() string`

GetUserInvitedDate returns the UserInvitedDate field if non-nil, zero value otherwise.

### GetUserInvitedDateOk

`func (o *InlineResponse20084People) GetUserInvitedDateOk() (*string, bool)`

GetUserInvitedDateOk returns a tuple with the UserInvitedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedDate

`func (o *InlineResponse20084People) SetUserInvitedDate(v string)`

SetUserInvitedDate sets UserInvitedDate field to given value.

### HasUserInvitedDate

`func (o *InlineResponse20084People) HasUserInvitedDate() bool`

HasUserInvitedDate returns a boolean if a field has been set.

### GetUserInvitedStatus

`func (o *InlineResponse20084People) GetUserInvitedStatus() string`

GetUserInvitedStatus returns the UserInvitedStatus field if non-nil, zero value otherwise.

### GetUserInvitedStatusOk

`func (o *InlineResponse20084People) GetUserInvitedStatusOk() (*string, bool)`

GetUserInvitedStatusOk returns a tuple with the UserInvitedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitedStatus

`func (o *InlineResponse20084People) SetUserInvitedStatus(v string)`

SetUserInvitedStatus sets UserInvitedStatus field to given value.

### HasUserInvitedStatus

`func (o *InlineResponse20084People) HasUserInvitedStatus() bool`

HasUserInvitedStatus returns a boolean if a field has been set.

### GetUserName

`func (o *InlineResponse20084People) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *InlineResponse20084People) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *InlineResponse20084People) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *InlineResponse20084People) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *InlineResponse20084People) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *InlineResponse20084People) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *InlineResponse20084People) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *InlineResponse20084People) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUUID

`func (o *InlineResponse20084People) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *InlineResponse20084People) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *InlineResponse20084People) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *InlineResponse20084People) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


