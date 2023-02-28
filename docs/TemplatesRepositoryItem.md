# TemplatesRepositoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Templates repository identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Templates repository name. | [optional] [readonly] 
**GitRepositoryUrl** | Pointer to **string** | Templates repository git repository URL. | [optional] [readonly] 
**GitRef** | Pointer to **string** | Templates repository git repository reference. | [optional] [readonly] 
**LastSyncSha** | Pointer to **NullableString** | Templates repository last synced git repository SHA. | [optional] [readonly] 
**Organization** | Pointer to **NullableString** | Organization identifier. | [optional] [readonly] 

## Methods

### NewTemplatesRepositoryItem

`func NewTemplatesRepositoryItem() *TemplatesRepositoryItem`

NewTemplatesRepositoryItem instantiates a new TemplatesRepositoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesRepositoryItemWithDefaults

`func NewTemplatesRepositoryItemWithDefaults() *TemplatesRepositoryItem`

NewTemplatesRepositoryItemWithDefaults instantiates a new TemplatesRepositoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplatesRepositoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatesRepositoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatesRepositoryItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatesRepositoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TemplatesRepositoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatesRepositoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatesRepositoryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplatesRepositoryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGitRepositoryUrl

`func (o *TemplatesRepositoryItem) GetGitRepositoryUrl() string`

GetGitRepositoryUrl returns the GitRepositoryUrl field if non-nil, zero value otherwise.

### GetGitRepositoryUrlOk

`func (o *TemplatesRepositoryItem) GetGitRepositoryUrlOk() (*string, bool)`

GetGitRepositoryUrlOk returns a tuple with the GitRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepositoryUrl

`func (o *TemplatesRepositoryItem) SetGitRepositoryUrl(v string)`

SetGitRepositoryUrl sets GitRepositoryUrl field to given value.

### HasGitRepositoryUrl

`func (o *TemplatesRepositoryItem) HasGitRepositoryUrl() bool`

HasGitRepositoryUrl returns a boolean if a field has been set.

### GetGitRef

`func (o *TemplatesRepositoryItem) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *TemplatesRepositoryItem) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *TemplatesRepositoryItem) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *TemplatesRepositoryItem) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetLastSyncSha

`func (o *TemplatesRepositoryItem) GetLastSyncSha() string`

GetLastSyncSha returns the LastSyncSha field if non-nil, zero value otherwise.

### GetLastSyncShaOk

`func (o *TemplatesRepositoryItem) GetLastSyncShaOk() (*string, bool)`

GetLastSyncShaOk returns a tuple with the LastSyncSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncSha

`func (o *TemplatesRepositoryItem) SetLastSyncSha(v string)`

SetLastSyncSha sets LastSyncSha field to given value.

### HasLastSyncSha

`func (o *TemplatesRepositoryItem) HasLastSyncSha() bool`

HasLastSyncSha returns a boolean if a field has been set.

### SetLastSyncShaNil

`func (o *TemplatesRepositoryItem) SetLastSyncShaNil(b bool)`

 SetLastSyncShaNil sets the value for LastSyncSha to be an explicit nil

### UnsetLastSyncSha
`func (o *TemplatesRepositoryItem) UnsetLastSyncSha()`

UnsetLastSyncSha ensures that no value is present for LastSyncSha, not even an explicit nil
### GetOrganization

`func (o *TemplatesRepositoryItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TemplatesRepositoryItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TemplatesRepositoryItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TemplatesRepositoryItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *TemplatesRepositoryItem) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *TemplatesRepositoryItem) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


