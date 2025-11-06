# TemplateSettingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchFallbackOnDelete** | Pointer to **NullableBool** | When a gitBranch is deleted, edit the environment with the branch from Template. | [optional] 
**BranchFallbackOnDeleteAutoDeploy** | Pointer to **NullableBool** | When gitBranch is edited with the one from Template, also auto-deploy. | [optional] 

## Methods

### NewTemplateSettingsItem

`func NewTemplateSettingsItem() *TemplateSettingsItem`

NewTemplateSettingsItem instantiates a new TemplateSettingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSettingsItemWithDefaults

`func NewTemplateSettingsItemWithDefaults() *TemplateSettingsItem`

NewTemplateSettingsItemWithDefaults instantiates a new TemplateSettingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchFallbackOnDelete

`func (o *TemplateSettingsItem) GetBranchFallbackOnDelete() bool`

GetBranchFallbackOnDelete returns the BranchFallbackOnDelete field if non-nil, zero value otherwise.

### GetBranchFallbackOnDeleteOk

`func (o *TemplateSettingsItem) GetBranchFallbackOnDeleteOk() (*bool, bool)`

GetBranchFallbackOnDeleteOk returns a tuple with the BranchFallbackOnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchFallbackOnDelete

`func (o *TemplateSettingsItem) SetBranchFallbackOnDelete(v bool)`

SetBranchFallbackOnDelete sets BranchFallbackOnDelete field to given value.

### HasBranchFallbackOnDelete

`func (o *TemplateSettingsItem) HasBranchFallbackOnDelete() bool`

HasBranchFallbackOnDelete returns a boolean if a field has been set.

### SetBranchFallbackOnDeleteNil

`func (o *TemplateSettingsItem) SetBranchFallbackOnDeleteNil(b bool)`

 SetBranchFallbackOnDeleteNil sets the value for BranchFallbackOnDelete to be an explicit nil

### UnsetBranchFallbackOnDelete
`func (o *TemplateSettingsItem) UnsetBranchFallbackOnDelete()`

UnsetBranchFallbackOnDelete ensures that no value is present for BranchFallbackOnDelete, not even an explicit nil
### GetBranchFallbackOnDeleteAutoDeploy

`func (o *TemplateSettingsItem) GetBranchFallbackOnDeleteAutoDeploy() bool`

GetBranchFallbackOnDeleteAutoDeploy returns the BranchFallbackOnDeleteAutoDeploy field if non-nil, zero value otherwise.

### GetBranchFallbackOnDeleteAutoDeployOk

`func (o *TemplateSettingsItem) GetBranchFallbackOnDeleteAutoDeployOk() (*bool, bool)`

GetBranchFallbackOnDeleteAutoDeployOk returns a tuple with the BranchFallbackOnDeleteAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchFallbackOnDeleteAutoDeploy

`func (o *TemplateSettingsItem) SetBranchFallbackOnDeleteAutoDeploy(v bool)`

SetBranchFallbackOnDeleteAutoDeploy sets BranchFallbackOnDeleteAutoDeploy field to given value.

### HasBranchFallbackOnDeleteAutoDeploy

`func (o *TemplateSettingsItem) HasBranchFallbackOnDeleteAutoDeploy() bool`

HasBranchFallbackOnDeleteAutoDeploy returns a boolean if a field has been set.

### SetBranchFallbackOnDeleteAutoDeployNil

`func (o *TemplateSettingsItem) SetBranchFallbackOnDeleteAutoDeployNil(b bool)`

 SetBranchFallbackOnDeleteAutoDeployNil sets the value for BranchFallbackOnDeleteAutoDeploy to be an explicit nil

### UnsetBranchFallbackOnDeleteAutoDeploy
`func (o *TemplateSettingsItem) UnsetBranchFallbackOnDeleteAutoDeploy()`

UnsetBranchFallbackOnDeleteAutoDeploy ensures that no value is present for BranchFallbackOnDeleteAutoDeploy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


