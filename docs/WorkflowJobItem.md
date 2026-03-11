# WorkflowJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Workflow job identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Workflow job name. | [optional] [readonly] 
**Type** | Pointer to **string** | Workflow job type. | [optional] [readonly] 
**Status** | Pointer to **string** | Workflow job status. | [optional] [readonly] 
**AllowedToFail** | Pointer to **bool** | Whether this workflow job is allowed to fail. | [optional] [readonly] 
**StartedAt** | Pointer to **NullableTime** | Workflow job start date and time. | [optional] [readonly] 
**EndedAt** | Pointer to **NullableTime** | Workflow job end date and time. | [optional] [readonly] 
**Duration** | Pointer to **NullableInt32** | Workflow job duration in seconds. | [optional] [readonly] 
**TimeoutSeconds** | Pointer to **int32** | Workflow job timeout in seconds. | [optional] [readonly] 
**Workflow** | Pointer to **string** | Workflow identifier. | [optional] [readonly] 

## Methods

### NewWorkflowJobItem

`func NewWorkflowJobItem() *WorkflowJobItem`

NewWorkflowJobItem instantiates a new WorkflowJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJobItemWithDefaults

`func NewWorkflowJobItemWithDefaults() *WorkflowJobItem`

NewWorkflowJobItemWithDefaults instantiates a new WorkflowJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowJobItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowJobItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowJobItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowJobItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowJobItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowJobItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkflowJobItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowJobItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowJobItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowJobItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowJobItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllowedToFail

`func (o *WorkflowJobItem) GetAllowedToFail() bool`

GetAllowedToFail returns the AllowedToFail field if non-nil, zero value otherwise.

### GetAllowedToFailOk

`func (o *WorkflowJobItem) GetAllowedToFailOk() (*bool, bool)`

GetAllowedToFailOk returns a tuple with the AllowedToFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToFail

`func (o *WorkflowJobItem) SetAllowedToFail(v bool)`

SetAllowedToFail sets AllowedToFail field to given value.

### HasAllowedToFail

`func (o *WorkflowJobItem) HasAllowedToFail() bool`

HasAllowedToFail returns a boolean if a field has been set.

### GetStartedAt

`func (o *WorkflowJobItem) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkflowJobItem) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkflowJobItem) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkflowJobItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *WorkflowJobItem) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *WorkflowJobItem) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *WorkflowJobItem) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *WorkflowJobItem) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *WorkflowJobItem) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *WorkflowJobItem) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *WorkflowJobItem) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *WorkflowJobItem) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetDuration

`func (o *WorkflowJobItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkflowJobItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkflowJobItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkflowJobItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *WorkflowJobItem) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *WorkflowJobItem) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTimeoutSeconds

`func (o *WorkflowJobItem) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WorkflowJobItem) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WorkflowJobItem) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *WorkflowJobItem) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetWorkflow

`func (o *WorkflowJobItem) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowJobItem) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowJobItem) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WorkflowJobItem) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


