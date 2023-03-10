# FilterGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "git"]
**Repository** | Pointer to **NullableString** |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFilterGit

`func NewFilterGit() *FilterGit`

NewFilterGit instantiates a new FilterGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterGitWithDefaults

`func NewFilterGitWithDefaults() *FilterGit`

NewFilterGitWithDefaults instantiates a new FilterGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FilterGit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilterGit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilterGit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FilterGit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRepository

`func (o *FilterGit) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FilterGit) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FilterGit) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *FilterGit) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *FilterGit) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *FilterGit) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetBranch

`func (o *FilterGit) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *FilterGit) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *FilterGit) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *FilterGit) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *FilterGit) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *FilterGit) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


