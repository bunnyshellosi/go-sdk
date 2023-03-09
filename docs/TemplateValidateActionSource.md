# TemplateValidateActionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "string"]
**Url** | **string** |  | 
**Ref** | **string** |  | 
**DirPath** | **string** |  | 
**OrganizationId** | **string** |  | 
**ValidateComponents** | Pointer to **bool** |  | [optional] 
**BunnyshellYaml** | **string** |  | 
**TemplateYaml** | **string** |  | 
**ValidateForOrganizationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTemplateValidateActionSource

`func NewTemplateValidateActionSource(url string, ref string, dirPath string, organizationId string, bunnyshellYaml string, templateYaml string, ) *TemplateValidateActionSource`

NewTemplateValidateActionSource instantiates a new TemplateValidateActionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateValidateActionSourceWithDefaults

`func NewTemplateValidateActionSourceWithDefaults() *TemplateValidateActionSource`

NewTemplateValidateActionSourceWithDefaults instantiates a new TemplateValidateActionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateValidateActionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateValidateActionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateValidateActionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateValidateActionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *TemplateValidateActionSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TemplateValidateActionSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TemplateValidateActionSource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRef

`func (o *TemplateValidateActionSource) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *TemplateValidateActionSource) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *TemplateValidateActionSource) SetRef(v string)`

SetRef sets Ref field to given value.


### GetDirPath

`func (o *TemplateValidateActionSource) GetDirPath() string`

GetDirPath returns the DirPath field if non-nil, zero value otherwise.

### GetDirPathOk

`func (o *TemplateValidateActionSource) GetDirPathOk() (*string, bool)`

GetDirPathOk returns a tuple with the DirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirPath

`func (o *TemplateValidateActionSource) SetDirPath(v string)`

SetDirPath sets DirPath field to given value.


### GetOrganizationId

`func (o *TemplateValidateActionSource) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TemplateValidateActionSource) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TemplateValidateActionSource) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetValidateComponents

`func (o *TemplateValidateActionSource) GetValidateComponents() bool`

GetValidateComponents returns the ValidateComponents field if non-nil, zero value otherwise.

### GetValidateComponentsOk

`func (o *TemplateValidateActionSource) GetValidateComponentsOk() (*bool, bool)`

GetValidateComponentsOk returns a tuple with the ValidateComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateComponents

`func (o *TemplateValidateActionSource) SetValidateComponents(v bool)`

SetValidateComponents sets ValidateComponents field to given value.

### HasValidateComponents

`func (o *TemplateValidateActionSource) HasValidateComponents() bool`

HasValidateComponents returns a boolean if a field has been set.

### GetBunnyshellYaml

`func (o *TemplateValidateActionSource) GetBunnyshellYaml() string`

GetBunnyshellYaml returns the BunnyshellYaml field if non-nil, zero value otherwise.

### GetBunnyshellYamlOk

`func (o *TemplateValidateActionSource) GetBunnyshellYamlOk() (*string, bool)`

GetBunnyshellYamlOk returns a tuple with the BunnyshellYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBunnyshellYaml

`func (o *TemplateValidateActionSource) SetBunnyshellYaml(v string)`

SetBunnyshellYaml sets BunnyshellYaml field to given value.


### GetTemplateYaml

`func (o *TemplateValidateActionSource) GetTemplateYaml() string`

GetTemplateYaml returns the TemplateYaml field if non-nil, zero value otherwise.

### GetTemplateYamlOk

`func (o *TemplateValidateActionSource) GetTemplateYamlOk() (*string, bool)`

GetTemplateYamlOk returns a tuple with the TemplateYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateYaml

`func (o *TemplateValidateActionSource) SetTemplateYaml(v string)`

SetTemplateYaml sets TemplateYaml field to given value.


### GetValidateForOrganizationId

`func (o *TemplateValidateActionSource) GetValidateForOrganizationId() string`

GetValidateForOrganizationId returns the ValidateForOrganizationId field if non-nil, zero value otherwise.

### GetValidateForOrganizationIdOk

`func (o *TemplateValidateActionSource) GetValidateForOrganizationIdOk() (*string, bool)`

GetValidateForOrganizationIdOk returns a tuple with the ValidateForOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateForOrganizationId

`func (o *TemplateValidateActionSource) SetValidateForOrganizationId(v string)`

SetValidateForOrganizationId sets ValidateForOrganizationId field to given value.

### HasValidateForOrganizationId

`func (o *TemplateValidateActionSource) HasValidateForOrganizationId() bool`

HasValidateForOrganizationId returns a boolean if a field has been set.

### SetValidateForOrganizationIdNil

`func (o *TemplateValidateActionSource) SetValidateForOrganizationIdNil(b bool)`

 SetValidateForOrganizationIdNil sets the value for ValidateForOrganizationId to be an explicit nil

### UnsetValidateForOrganizationId
`func (o *TemplateValidateActionSource) UnsetValidateForOrganizationId()`

UnsetValidateForOrganizationId ensures that no value is present for ValidateForOrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


