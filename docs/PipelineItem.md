# PipelineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Pipeline identifier. | [optional] [readonly] 
**Description** | Pointer to **string** | Pipeline description. | [optional] [readonly] 
**Status** | Pointer to **string** | Pipeline status. | [optional] [readonly] 
**Stages** | Pointer to [**[]StageItem**](StageItem.md) | Stage identifier. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Environment identifier. | [optional] [readonly] 
**Event** | Pointer to **string** | Event identifier. | [optional] [readonly] 
**Organization** | Pointer to **NullableString** | Organization identifier. | [optional] [readonly] 

## Methods

### NewPipelineItem

`func NewPipelineItem() *PipelineItem`

NewPipelineItem instantiates a new PipelineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineItemWithDefaults

`func NewPipelineItemWithDefaults() *PipelineItem`

NewPipelineItemWithDefaults instantiates a new PipelineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PipelineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *PipelineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PipelineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PipelineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PipelineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PipelineItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PipelineItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStages

`func (o *PipelineItem) GetStages() []StageItem`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *PipelineItem) GetStagesOk() (*[]StageItem, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *PipelineItem) SetStages(v []StageItem)`

SetStages sets Stages field to given value.

### HasStages

`func (o *PipelineItem) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetEnvironment

`func (o *PipelineItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PipelineItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PipelineItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PipelineItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *PipelineItem) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *PipelineItem) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEvent

`func (o *PipelineItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PipelineItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PipelineItem) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PipelineItem) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOrganization

`func (o *PipelineItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PipelineItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PipelineItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PipelineItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *PipelineItem) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PipelineItem) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


