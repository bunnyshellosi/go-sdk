# WorkflowItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Workflow identifier. | [optional] [readonly] 
**Description** | Pointer to **string** | Workflow description. | [optional] [readonly] 
**Status** | Pointer to **string** | Workflow status. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Environment identifier. | [optional] [readonly] 
**Event** | Pointer to **string** | Event identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 
**Duration** | Pointer to **NullableInt32** | Workflow duration in seconds. | [optional] [readonly] 
**JobsCount** | Pointer to **int32** | Number of jobs. | [optional] [readonly] 
**CompletedJobsCount** | Pointer to **int32** | Number of completed jobs. | [optional] [readonly] 

## Methods

### NewWorkflowItem

`func NewWorkflowItem() *WorkflowItem`

NewWorkflowItem instantiates a new WorkflowItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowItemWithDefaults

`func NewWorkflowItemWithDefaults() *WorkflowItem`

NewWorkflowItemWithDefaults instantiates a new WorkflowItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *WorkflowItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WorkflowItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WorkflowItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *WorkflowItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *WorkflowItem) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *WorkflowItem) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEvent

`func (o *WorkflowItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowItem) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowItem) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetDuration

`func (o *WorkflowItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkflowItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkflowItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkflowItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *WorkflowItem) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *WorkflowItem) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetJobsCount

`func (o *WorkflowItem) GetJobsCount() int32`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *WorkflowItem) GetJobsCountOk() (*int32, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *WorkflowItem) SetJobsCount(v int32)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *WorkflowItem) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetCompletedJobsCount

`func (o *WorkflowItem) GetCompletedJobsCount() int32`

GetCompletedJobsCount returns the CompletedJobsCount field if non-nil, zero value otherwise.

### GetCompletedJobsCountOk

`func (o *WorkflowItem) GetCompletedJobsCountOk() (*int32, bool)`

GetCompletedJobsCountOk returns a tuple with the CompletedJobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobsCount

`func (o *WorkflowItem) SetCompletedJobsCount(v int32)`

SetCompletedJobsCount sets CompletedJobsCount field to given value.

### HasCompletedJobsCount

`func (o *WorkflowItem) HasCompletedJobsCount() bool`

HasCompletedJobsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


