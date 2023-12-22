# ProjectVariableCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Project variable identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Project variable name. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | Project variable value. | [optional] [readonly] 
**Secret** | Pointer to **bool** | Project variable marked as secret. | [optional] [readonly] 
**Project** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewProjectVariableCollection

`func NewProjectVariableCollection() *ProjectVariableCollection`

NewProjectVariableCollection instantiates a new ProjectVariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectVariableCollectionWithDefaults

`func NewProjectVariableCollectionWithDefaults() *ProjectVariableCollection`

NewProjectVariableCollectionWithDefaults instantiates a new ProjectVariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectVariableCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectVariableCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectVariableCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectVariableCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectVariableCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectVariableCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectVariableCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectVariableCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ProjectVariableCollection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectVariableCollection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectVariableCollection) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProjectVariableCollection) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProjectVariableCollection) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProjectVariableCollection) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *ProjectVariableCollection) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProjectVariableCollection) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProjectVariableCollection) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProjectVariableCollection) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetProject

`func (o *ProjectVariableCollection) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectVariableCollection) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectVariableCollection) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectVariableCollection) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectVariableCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectVariableCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectVariableCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectVariableCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


