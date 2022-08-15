# OrganizationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Organization identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Organization name as defined in Bunnyshell UI. | [optional] [readonly] 
**Timezone** | Pointer to **string** | Organization timezone as defined in Bunnyshell UI. | [optional] [readonly] 
**TotalUsers** | Pointer to **int32** | Organization total users. | [optional] [readonly] 
**TotalProjects** | Pointer to **int32** | Organization total projects. | [optional] [readonly] 
**AvailableRegistries** | Pointer to **int32** | Organization total available registry integrations. | [optional] [readonly] 
**AvailableClusters** | Pointer to **int32** | Organization total available cluster integrations. | [optional] [readonly] 
**AvailableGitIntegration** | Pointer to **int32** | Organization total available git integrations. | [optional] [readonly] 

## Methods

### NewOrganizationItem

`func NewOrganizationItem() *OrganizationItem`

NewOrganizationItem instantiates a new OrganizationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationItemWithDefaults

`func NewOrganizationItemWithDefaults() *OrganizationItem`

NewOrganizationItemWithDefaults instantiates a new OrganizationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimezone

`func (o *OrganizationItem) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OrganizationItem) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OrganizationItem) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *OrganizationItem) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTotalUsers

`func (o *OrganizationItem) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *OrganizationItem) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *OrganizationItem) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.

### HasTotalUsers

`func (o *OrganizationItem) HasTotalUsers() bool`

HasTotalUsers returns a boolean if a field has been set.

### GetTotalProjects

`func (o *OrganizationItem) GetTotalProjects() int32`

GetTotalProjects returns the TotalProjects field if non-nil, zero value otherwise.

### GetTotalProjectsOk

`func (o *OrganizationItem) GetTotalProjectsOk() (*int32, bool)`

GetTotalProjectsOk returns a tuple with the TotalProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProjects

`func (o *OrganizationItem) SetTotalProjects(v int32)`

SetTotalProjects sets TotalProjects field to given value.

### HasTotalProjects

`func (o *OrganizationItem) HasTotalProjects() bool`

HasTotalProjects returns a boolean if a field has been set.

### GetAvailableRegistries

`func (o *OrganizationItem) GetAvailableRegistries() int32`

GetAvailableRegistries returns the AvailableRegistries field if non-nil, zero value otherwise.

### GetAvailableRegistriesOk

`func (o *OrganizationItem) GetAvailableRegistriesOk() (*int32, bool)`

GetAvailableRegistriesOk returns a tuple with the AvailableRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRegistries

`func (o *OrganizationItem) SetAvailableRegistries(v int32)`

SetAvailableRegistries sets AvailableRegistries field to given value.

### HasAvailableRegistries

`func (o *OrganizationItem) HasAvailableRegistries() bool`

HasAvailableRegistries returns a boolean if a field has been set.

### GetAvailableClusters

`func (o *OrganizationItem) GetAvailableClusters() int32`

GetAvailableClusters returns the AvailableClusters field if non-nil, zero value otherwise.

### GetAvailableClustersOk

`func (o *OrganizationItem) GetAvailableClustersOk() (*int32, bool)`

GetAvailableClustersOk returns a tuple with the AvailableClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableClusters

`func (o *OrganizationItem) SetAvailableClusters(v int32)`

SetAvailableClusters sets AvailableClusters field to given value.

### HasAvailableClusters

`func (o *OrganizationItem) HasAvailableClusters() bool`

HasAvailableClusters returns a boolean if a field has been set.

### GetAvailableGitIntegration

`func (o *OrganizationItem) GetAvailableGitIntegration() int32`

GetAvailableGitIntegration returns the AvailableGitIntegration field if non-nil, zero value otherwise.

### GetAvailableGitIntegrationOk

`func (o *OrganizationItem) GetAvailableGitIntegrationOk() (*int32, bool)`

GetAvailableGitIntegrationOk returns a tuple with the AvailableGitIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableGitIntegration

`func (o *OrganizationItem) SetAvailableGitIntegration(v int32)`

SetAvailableGitIntegration sets AvailableGitIntegration field to given value.

### HasAvailableGitIntegration

`func (o *OrganizationItem) HasAvailableGitIntegration() bool`

HasAvailableGitIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


