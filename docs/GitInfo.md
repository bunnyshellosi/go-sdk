# GitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **NullableString** |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGitInfo

`func NewGitInfo() *GitInfo`

NewGitInfo instantiates a new GitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitInfoWithDefaults

`func NewGitInfoWithDefaults() *GitInfo`

NewGitInfoWithDefaults instantiates a new GitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *GitInfo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitInfo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitInfo) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GitInfo) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *GitInfo) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *GitInfo) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetBranch

`func (o *GitInfo) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitInfo) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitInfo) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitInfo) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *GitInfo) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *GitInfo) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


