# ComponentGitCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Service component name | [optional] [readonly] 
**Repository** | Pointer to **string** | Git repository | [optional] [readonly] 
**RefName** | Pointer to **string** | Git ref name | [optional] [readonly] 
**Path** | Pointer to **NullableString** | Git application path | [optional] [readonly] 
**RefSha** | Pointer to **NullableString** | Git ref sha | [optional] [readonly] 
**DeployedSha** | Pointer to **NullableString** | Git deployed sha | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 

## Methods

### NewComponentGitCollection

`func NewComponentGitCollection() *ComponentGitCollection`

NewComponentGitCollection instantiates a new ComponentGitCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentGitCollectionWithDefaults

`func NewComponentGitCollectionWithDefaults() *ComponentGitCollection`

NewComponentGitCollectionWithDefaults instantiates a new ComponentGitCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentGitCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentGitCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentGitCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentGitCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentGitCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentGitCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentGitCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentGitCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *ComponentGitCollection) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ComponentGitCollection) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ComponentGitCollection) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ComponentGitCollection) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRefName

`func (o *ComponentGitCollection) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ComponentGitCollection) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ComponentGitCollection) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ComponentGitCollection) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPath

`func (o *ComponentGitCollection) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ComponentGitCollection) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ComponentGitCollection) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ComponentGitCollection) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ComponentGitCollection) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ComponentGitCollection) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRefSha

`func (o *ComponentGitCollection) GetRefSha() string`

GetRefSha returns the RefSha field if non-nil, zero value otherwise.

### GetRefShaOk

`func (o *ComponentGitCollection) GetRefShaOk() (*string, bool)`

GetRefShaOk returns a tuple with the RefSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSha

`func (o *ComponentGitCollection) SetRefSha(v string)`

SetRefSha sets RefSha field to given value.

### HasRefSha

`func (o *ComponentGitCollection) HasRefSha() bool`

HasRefSha returns a boolean if a field has been set.

### SetRefShaNil

`func (o *ComponentGitCollection) SetRefShaNil(b bool)`

 SetRefShaNil sets the value for RefSha to be an explicit nil

### UnsetRefSha
`func (o *ComponentGitCollection) UnsetRefSha()`

UnsetRefSha ensures that no value is present for RefSha, not even an explicit nil
### GetDeployedSha

`func (o *ComponentGitCollection) GetDeployedSha() string`

GetDeployedSha returns the DeployedSha field if non-nil, zero value otherwise.

### GetDeployedShaOk

`func (o *ComponentGitCollection) GetDeployedShaOk() (*string, bool)`

GetDeployedShaOk returns a tuple with the DeployedSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSha

`func (o *ComponentGitCollection) SetDeployedSha(v string)`

SetDeployedSha sets DeployedSha field to given value.

### HasDeployedSha

`func (o *ComponentGitCollection) HasDeployedSha() bool`

HasDeployedSha returns a boolean if a field has been set.

### SetDeployedShaNil

`func (o *ComponentGitCollection) SetDeployedShaNil(b bool)`

 SetDeployedShaNil sets the value for DeployedSha to be an explicit nil

### UnsetDeployedSha
`func (o *ComponentGitCollection) UnsetDeployedSha()`

UnsetDeployedSha ensures that no value is present for DeployedSha, not even an explicit nil
### GetEnvironment

`func (o *ComponentGitCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ComponentGitCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ComponentGitCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ComponentGitCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


