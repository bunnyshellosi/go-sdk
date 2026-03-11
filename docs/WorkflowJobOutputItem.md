# WorkflowJobOutputItem

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

## Methods

### NewWorkflowJobOutputItem

`func NewWorkflowJobOutputItem() *WorkflowJobOutputItem`

NewWorkflowJobOutputItem instantiates a new WorkflowJobOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJobOutputItemWithDefaults

`func NewWorkflowJobOutputItemWithDefaults() *WorkflowJobOutputItem`

NewWorkflowJobOutputItemWithDefaults instantiates a new WorkflowJobOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowJobOutputItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowJobOutputItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowJobOutputItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowJobOutputItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowJobOutputItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowJobOutputItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowJobOutputItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowJobOutputItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkflowJobOutputItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowJobOutputItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowJobOutputItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowJobOutputItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowJobOutputItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowJobOutputItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowJobOutputItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowJobOutputItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllowedToFail

`func (o *WorkflowJobOutputItem) GetAllowedToFail() bool`

GetAllowedToFail returns the AllowedToFail field if non-nil, zero value otherwise.

### GetAllowedToFailOk

`func (o *WorkflowJobOutputItem) GetAllowedToFailOk() (*bool, bool)`

GetAllowedToFailOk returns a tuple with the AllowedToFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToFail

`func (o *WorkflowJobOutputItem) SetAllowedToFail(v bool)`

SetAllowedToFail sets AllowedToFail field to given value.

### HasAllowedToFail

`func (o *WorkflowJobOutputItem) HasAllowedToFail() bool`

HasAllowedToFail returns a boolean if a field has been set.

### GetStartedAt

`func (o *WorkflowJobOutputItem) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkflowJobOutputItem) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkflowJobOutputItem) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkflowJobOutputItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *WorkflowJobOutputItem) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *WorkflowJobOutputItem) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *WorkflowJobOutputItem) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *WorkflowJobOutputItem) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *WorkflowJobOutputItem) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *WorkflowJobOutputItem) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *WorkflowJobOutputItem) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *WorkflowJobOutputItem) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetDuration

`func (o *WorkflowJobOutputItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkflowJobOutputItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkflowJobOutputItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkflowJobOutputItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *WorkflowJobOutputItem) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *WorkflowJobOutputItem) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTimeoutSeconds

`func (o *WorkflowJobOutputItem) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WorkflowJobOutputItem) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WorkflowJobOutputItem) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *WorkflowJobOutputItem) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


