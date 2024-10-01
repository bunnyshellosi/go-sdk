# EnvironItemItem

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

### NewEnvironItemItem

`func NewEnvironItemItem() *EnvironItemItem`

NewEnvironItemItem instantiates a new EnvironItemItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironItemItemWithDefaults

`func NewEnvironItemItemWithDefaults() *EnvironItemItem`

NewEnvironItemItemWithDefaults instantiates a new EnvironItemItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironItemItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironItemItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironItemItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironItemItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupName

`func (o *EnvironItemItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EnvironItemItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EnvironItemItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *EnvironItemItem) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetName

`func (o *EnvironItemItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironItemItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironItemItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironItemItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *EnvironItemItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironItemItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironItemItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironItemItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EnvironItemItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnvironItemItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *EnvironItemItem) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EnvironItemItem) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EnvironItemItem) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *EnvironItemItem) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvironItemItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironItemItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironItemItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironItemItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetOrganization

`func (o *EnvironItemItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EnvironItemItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EnvironItemItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EnvironItemItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


