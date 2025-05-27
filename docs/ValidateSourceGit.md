# ValidateSourceGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "git"]
**Url** | **string** |  | 
**Ref** | **string** |  | 
**DirPath** | **string** |  | 
**OrganizationId** | **string** |  | 
**ValidateAllowExtraFields** | Pointer to **bool** |  | [optional] 
**ValidateComponents** | Pointer to **bool** |  | [optional] 

## Methods

### NewValidateSourceGit

`func NewValidateSourceGit(url string, ref string, dirPath string, organizationId string, ) *ValidateSourceGit`

NewValidateSourceGit instantiates a new ValidateSourceGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateSourceGitWithDefaults

`func NewValidateSourceGitWithDefaults() *ValidateSourceGit`

NewValidateSourceGitWithDefaults instantiates a new ValidateSourceGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ValidateSourceGit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidateSourceGit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidateSourceGit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidateSourceGit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ValidateSourceGit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ValidateSourceGit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ValidateSourceGit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRef

`func (o *ValidateSourceGit) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ValidateSourceGit) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ValidateSourceGit) SetRef(v string)`

SetRef sets Ref field to given value.


### GetDirPath

`func (o *ValidateSourceGit) GetDirPath() string`

GetDirPath returns the DirPath field if non-nil, zero value otherwise.

### GetDirPathOk

`func (o *ValidateSourceGit) GetDirPathOk() (*string, bool)`

GetDirPathOk returns a tuple with the DirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirPath

`func (o *ValidateSourceGit) SetDirPath(v string)`

SetDirPath sets DirPath field to given value.


### GetOrganizationId

`func (o *ValidateSourceGit) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ValidateSourceGit) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ValidateSourceGit) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetValidateAllowExtraFields

`func (o *ValidateSourceGit) GetValidateAllowExtraFields() bool`

GetValidateAllowExtraFields returns the ValidateAllowExtraFields field if non-nil, zero value otherwise.

### GetValidateAllowExtraFieldsOk

`func (o *ValidateSourceGit) GetValidateAllowExtraFieldsOk() (*bool, bool)`

GetValidateAllowExtraFieldsOk returns a tuple with the ValidateAllowExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAllowExtraFields

`func (o *ValidateSourceGit) SetValidateAllowExtraFields(v bool)`

SetValidateAllowExtraFields sets ValidateAllowExtraFields field to given value.

### HasValidateAllowExtraFields

`func (o *ValidateSourceGit) HasValidateAllowExtraFields() bool`

HasValidateAllowExtraFields returns a boolean if a field has been set.

### GetValidateComponents

`func (o *ValidateSourceGit) GetValidateComponents() bool`

GetValidateComponents returns the ValidateComponents field if non-nil, zero value otherwise.

### GetValidateComponentsOk

`func (o *ValidateSourceGit) GetValidateComponentsOk() (*bool, bool)`

GetValidateComponentsOk returns a tuple with the ValidateComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateComponents

`func (o *ValidateSourceGit) SetValidateComponents(v bool)`

SetValidateComponents sets ValidateComponents field to given value.

### HasValidateComponents

`func (o *ValidateSourceGit) HasValidateComponents() bool`

HasValidateComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


