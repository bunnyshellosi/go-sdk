# EventItem

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

### NewEventItem

`func NewEventItem() *EventItem`

NewEventItem instantiates a new EventItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventItemWithDefaults

`func NewEventItemWithDefaults() *EventItem`

NewEventItemWithDefaults instantiates a new EventItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EventItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *EventItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *EventItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EventItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EventItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EventItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EventItem) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EventItem) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOrganization

`func (o *EventItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EventItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EventItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EventItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EventItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EventItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


