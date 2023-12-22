# ProjectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Environment labels. | [optional] [readonly] 
**Name** | Pointer to **string** | Project name. | [optional] [readonly] 
**TotalEnvironments** | Pointer to **int32** | Environment identifier. | [optional] [readonly] 
**BuildSettings** | Pointer to [**NullableBuildSettingsItem**](BuildSettingsItem.md) |  | [optional] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewProjectItem

`func NewProjectItem() *ProjectItem`

NewProjectItem instantiates a new ProjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectItemWithDefaults

`func NewProjectItemWithDefaults() *ProjectItem`

NewProjectItemWithDefaults instantiates a new ProjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *ProjectItem) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ProjectItem) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ProjectItem) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ProjectItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *ProjectItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalEnvironments

`func (o *ProjectItem) GetTotalEnvironments() int32`

GetTotalEnvironments returns the TotalEnvironments field if non-nil, zero value otherwise.

### GetTotalEnvironmentsOk

`func (o *ProjectItem) GetTotalEnvironmentsOk() (*int32, bool)`

GetTotalEnvironmentsOk returns a tuple with the TotalEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnvironments

`func (o *ProjectItem) SetTotalEnvironments(v int32)`

SetTotalEnvironments sets TotalEnvironments field to given value.

### HasTotalEnvironments

`func (o *ProjectItem) HasTotalEnvironments() bool`

HasTotalEnvironments returns a boolean if a field has been set.

### GetBuildSettings

`func (o *ProjectItem) GetBuildSettings() BuildSettingsItem`

GetBuildSettings returns the BuildSettings field if non-nil, zero value otherwise.

### GetBuildSettingsOk

`func (o *ProjectItem) GetBuildSettingsOk() (*BuildSettingsItem, bool)`

GetBuildSettingsOk returns a tuple with the BuildSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSettings

`func (o *ProjectItem) SetBuildSettings(v BuildSettingsItem)`

SetBuildSettings sets BuildSettings field to given value.

### HasBuildSettings

`func (o *ProjectItem) HasBuildSettings() bool`

HasBuildSettings returns a boolean if a field has been set.

### SetBuildSettingsNil

`func (o *ProjectItem) SetBuildSettingsNil(b bool)`

 SetBuildSettingsNil sets the value for BuildSettings to be an explicit nil

### UnsetBuildSettings
`func (o *ProjectItem) UnsetBuildSettings()`

UnsetBuildSettings ensures that no value is present for BuildSettings, not even an explicit nil
### GetOrganization

`func (o *ProjectItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


