# EnvironItemCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Environment variable identifier. | [optional] [readonly] 
**GroupName** | Pointer to **string** | Environ group name. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment variable name. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | Environment variable value. | [optional] [readonly] 
**Secret** | Pointer to **bool** | Environment variable marked as secret. | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewEnvironItemCollection

`func NewEnvironItemCollection() *EnvironItemCollection`

NewEnvironItemCollection instantiates a new EnvironItemCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironItemCollectionWithDefaults

`func NewEnvironItemCollectionWithDefaults() *EnvironItemCollection`

NewEnvironItemCollectionWithDefaults instantiates a new EnvironItemCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironItemCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironItemCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironItemCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironItemCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupName

`func (o *EnvironItemCollection) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EnvironItemCollection) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EnvironItemCollection) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *EnvironItemCollection) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetName

`func (o *EnvironItemCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironItemCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironItemCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironItemCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *EnvironItemCollection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironItemCollection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironItemCollection) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironItemCollection) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EnvironItemCollection) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnvironItemCollection) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *EnvironItemCollection) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EnvironItemCollection) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EnvironItemCollection) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *EnvironItemCollection) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvironItemCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironItemCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironItemCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironItemCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetOrganization

`func (o *EnvironItemCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EnvironItemCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EnvironItemCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EnvironItemCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


