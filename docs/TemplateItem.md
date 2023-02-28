# TemplateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Template identifier. | [optional] [readonly] 
**Key** | Pointer to **string** | Template key. | [optional] [readonly] 
**Name** | Pointer to **string** | Template name. | [optional] [readonly] 
**GitSha** | Pointer to **string** | Template git repository SHA. | [optional] [readonly] 
**ShortDescription** | Pointer to **NullableString** | Template short description. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | Template tags. | [optional] [readonly] 
**Organization** | Pointer to **NullableString** | Organization identifier. | [optional] [readonly] 
**TemplatesRepository** | Pointer to **NullableString** | Templates repository identifier. | [optional] [readonly] 

## Methods

### NewTemplateItem

`func NewTemplateItem() *TemplateItem`

NewTemplateItem instantiates a new TemplateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateItemWithDefaults

`func NewTemplateItemWithDefaults() *TemplateItem`

NewTemplateItemWithDefaults instantiates a new TemplateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *TemplateItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TemplateItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TemplateItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TemplateItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TemplateItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGitSha

`func (o *TemplateItem) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *TemplateItem) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *TemplateItem) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *TemplateItem) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetShortDescription

`func (o *TemplateItem) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *TemplateItem) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *TemplateItem) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *TemplateItem) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *TemplateItem) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *TemplateItem) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetTags

`func (o *TemplateItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOrganization

`func (o *TemplateItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TemplateItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TemplateItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TemplateItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *TemplateItem) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *TemplateItem) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTemplatesRepository

`func (o *TemplateItem) GetTemplatesRepository() string`

GetTemplatesRepository returns the TemplatesRepository field if non-nil, zero value otherwise.

### GetTemplatesRepositoryOk

`func (o *TemplateItem) GetTemplatesRepositoryOk() (*string, bool)`

GetTemplatesRepositoryOk returns a tuple with the TemplatesRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesRepository

`func (o *TemplateItem) SetTemplatesRepository(v string)`

SetTemplatesRepository sets TemplatesRepository field to given value.

### HasTemplatesRepository

`func (o *TemplateItem) HasTemplatesRepository() bool`

HasTemplatesRepository returns a boolean if a field has been set.

### SetTemplatesRepositoryNil

`func (o *TemplateItem) SetTemplatesRepositoryNil(b bool)`

 SetTemplatesRepositoryNil sets the value for TemplatesRepository to be an explicit nil

### UnsetTemplatesRepository
`func (o *TemplateItem) UnsetTemplatesRepository()`

UnsetTemplatesRepository ensures that no value is present for TemplatesRepository, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


