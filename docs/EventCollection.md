# EventCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Event identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | Event type. | [optional] [readonly] 
**Status** | Pointer to **string** | Event operation status. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Environment identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewEventCollection

`func NewEventCollection() *EventCollection`

NewEventCollection instantiates a new EventCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCollectionWithDefaults

`func NewEventCollectionWithDefaults() *EventCollection`

NewEventCollectionWithDefaults instantiates a new EventCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EventCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventCollection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventCollection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *EventCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventCollection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventCollection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *EventCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EventCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EventCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EventCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EventCollection) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EventCollection) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOrganization

`func (o *EventCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EventCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EventCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EventCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


