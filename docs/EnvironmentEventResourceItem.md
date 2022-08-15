# EnvironmentEventResourceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Event identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | Event type. | [optional] [readonly] 
**Status** | Pointer to **string** | Event operation status. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Environment identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Event creation time. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Event last update time. | [optional] [readonly] 

## Methods

### NewEnvironmentEventResourceItem

`func NewEnvironmentEventResourceItem() *EnvironmentEventResourceItem`

NewEnvironmentEventResourceItem instantiates a new EnvironmentEventResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEventResourceItemWithDefaults

`func NewEnvironmentEventResourceItemWithDefaults() *EnvironmentEventResourceItem`

NewEnvironmentEventResourceItemWithDefaults instantiates a new EnvironmentEventResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentEventResourceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentEventResourceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentEventResourceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentEventResourceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentEventResourceItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentEventResourceItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentEventResourceItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentEventResourceItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentEventResourceItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentEventResourceItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentEventResourceItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentEventResourceItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvironmentEventResourceItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentEventResourceItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentEventResourceItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentEventResourceItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentEventResourceItem) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentEventResourceItem) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOrganization

`func (o *EnvironmentEventResourceItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EnvironmentEventResourceItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EnvironmentEventResourceItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EnvironmentEventResourceItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnvironmentEventResourceItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentEventResourceItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentEventResourceItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentEventResourceItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentEventResourceItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentEventResourceItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentEventResourceItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentEventResourceItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


