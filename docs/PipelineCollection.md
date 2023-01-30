# PipelineCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Pipeline identifier. | [optional] [readonly] 
**Description** | Pointer to **string** | Pipeline description. | [optional] [readonly] 
**Status** | Pointer to **string** | Pipeline status. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Environment identifier. | [optional] [readonly] 
**Event** | Pointer to **string** | Event identifier. | [optional] [readonly] 
**Organization** | Pointer to **NullableString** | Organization identifier. | [optional] [readonly] 

## Methods

### NewPipelineCollection

`func NewPipelineCollection() *PipelineCollection`

NewPipelineCollection instantiates a new PipelineCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCollectionWithDefaults

`func NewPipelineCollectionWithDefaults() *PipelineCollection`

NewPipelineCollectionWithDefaults instantiates a new PipelineCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PipelineCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *PipelineCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PipelineCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PipelineCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PipelineCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PipelineCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineCollection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PipelineCollection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *PipelineCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PipelineCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PipelineCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PipelineCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *PipelineCollection) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *PipelineCollection) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEvent

`func (o *PipelineCollection) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PipelineCollection) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PipelineCollection) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PipelineCollection) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOrganization

`func (o *PipelineCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PipelineCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PipelineCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PipelineCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *PipelineCollection) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PipelineCollection) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


