# WorkflowCollection

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
**WebUrl** | Pointer to **string** | Workflow web URL. | [optional] [readonly] 

## Methods

### NewWorkflowCollection

`func NewWorkflowCollection() *WorkflowCollection`

NewWorkflowCollection instantiates a new WorkflowCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCollectionWithDefaults

`func NewWorkflowCollectionWithDefaults() *WorkflowCollection`

NewWorkflowCollectionWithDefaults instantiates a new WorkflowCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowCollection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowCollection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *WorkflowCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WorkflowCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WorkflowCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *WorkflowCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *WorkflowCollection) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *WorkflowCollection) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEvent

`func (o *WorkflowCollection) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowCollection) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowCollection) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowCollection) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetDuration

`func (o *WorkflowCollection) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkflowCollection) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkflowCollection) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkflowCollection) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *WorkflowCollection) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *WorkflowCollection) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetJobsCount

`func (o *WorkflowCollection) GetJobsCount() int32`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *WorkflowCollection) GetJobsCountOk() (*int32, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *WorkflowCollection) SetJobsCount(v int32)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *WorkflowCollection) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetCompletedJobsCount

`func (o *WorkflowCollection) GetCompletedJobsCount() int32`

GetCompletedJobsCount returns the CompletedJobsCount field if non-nil, zero value otherwise.

### GetCompletedJobsCountOk

`func (o *WorkflowCollection) GetCompletedJobsCountOk() (*int32, bool)`

GetCompletedJobsCountOk returns a tuple with the CompletedJobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobsCount

`func (o *WorkflowCollection) SetCompletedJobsCount(v int32)`

SetCompletedJobsCount sets CompletedJobsCount field to given value.

### HasCompletedJobsCount

`func (o *WorkflowCollection) HasCompletedJobsCount() bool`

HasCompletedJobsCount returns a boolean if a field has been set.

### GetWebUrl

`func (o *WorkflowCollection) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *WorkflowCollection) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *WorkflowCollection) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *WorkflowCollection) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


