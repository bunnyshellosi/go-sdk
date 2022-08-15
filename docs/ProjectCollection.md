# ProjectCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Project name. | [optional] [readonly] 
**TotalEnvironments** | Pointer to **int32** | Environment identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewProjectCollection

`func NewProjectCollection() *ProjectCollection`

NewProjectCollection instantiates a new ProjectCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCollectionWithDefaults

`func NewProjectCollectionWithDefaults() *ProjectCollection`

NewProjectCollectionWithDefaults instantiates a new ProjectCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalEnvironments

`func (o *ProjectCollection) GetTotalEnvironments() int32`

GetTotalEnvironments returns the TotalEnvironments field if non-nil, zero value otherwise.

### GetTotalEnvironmentsOk

`func (o *ProjectCollection) GetTotalEnvironmentsOk() (*int32, bool)`

GetTotalEnvironmentsOk returns a tuple with the TotalEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnvironments

`func (o *ProjectCollection) SetTotalEnvironments(v int32)`

SetTotalEnvironments sets TotalEnvironments field to given value.

### HasTotalEnvironments

`func (o *ProjectCollection) HasTotalEnvironments() bool`

HasTotalEnvironments returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


