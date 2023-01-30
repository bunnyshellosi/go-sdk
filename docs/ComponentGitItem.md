# ComponentGitItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Service component name | [optional] [readonly] 
**Repository** | Pointer to **string** | Git repository | [optional] [readonly] 
**RefName** | Pointer to **string** | Git ref name | [optional] [readonly] 
**RefSha** | Pointer to **NullableString** | Git ref sha | [optional] [readonly] 
**DeployedSha** | Pointer to **NullableString** | Git deployed sha | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 

## Methods

### NewComponentGitItem

`func NewComponentGitItem() *ComponentGitItem`

NewComponentGitItem instantiates a new ComponentGitItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentGitItemWithDefaults

`func NewComponentGitItemWithDefaults() *ComponentGitItem`

NewComponentGitItemWithDefaults instantiates a new ComponentGitItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentGitItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentGitItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentGitItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentGitItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentGitItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentGitItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentGitItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentGitItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *ComponentGitItem) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ComponentGitItem) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ComponentGitItem) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ComponentGitItem) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRefName

`func (o *ComponentGitItem) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ComponentGitItem) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ComponentGitItem) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ComponentGitItem) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRefSha

`func (o *ComponentGitItem) GetRefSha() string`

GetRefSha returns the RefSha field if non-nil, zero value otherwise.

### GetRefShaOk

`func (o *ComponentGitItem) GetRefShaOk() (*string, bool)`

GetRefShaOk returns a tuple with the RefSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSha

`func (o *ComponentGitItem) SetRefSha(v string)`

SetRefSha sets RefSha field to given value.

### HasRefSha

`func (o *ComponentGitItem) HasRefSha() bool`

HasRefSha returns a boolean if a field has been set.

### SetRefShaNil

`func (o *ComponentGitItem) SetRefShaNil(b bool)`

 SetRefShaNil sets the value for RefSha to be an explicit nil

### UnsetRefSha
`func (o *ComponentGitItem) UnsetRefSha()`

UnsetRefSha ensures that no value is present for RefSha, not even an explicit nil
### GetDeployedSha

`func (o *ComponentGitItem) GetDeployedSha() string`

GetDeployedSha returns the DeployedSha field if non-nil, zero value otherwise.

### GetDeployedShaOk

`func (o *ComponentGitItem) GetDeployedShaOk() (*string, bool)`

GetDeployedShaOk returns a tuple with the DeployedSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSha

`func (o *ComponentGitItem) SetDeployedSha(v string)`

SetDeployedSha sets DeployedSha field to given value.

### HasDeployedSha

`func (o *ComponentGitItem) HasDeployedSha() bool`

HasDeployedSha returns a boolean if a field has been set.

### SetDeployedShaNil

`func (o *ComponentGitItem) SetDeployedShaNil(b bool)`

 SetDeployedShaNil sets the value for DeployedSha to be an explicit nil

### UnsetDeployedSha
`func (o *ComponentGitItem) UnsetDeployedSha()`

UnsetDeployedSha ensures that no value is present for DeployedSha, not even an explicit nil
### GetEnvironment

`func (o *ComponentGitItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ComponentGitItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ComponentGitItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ComponentGitItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


