# WorkflowStepLogOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Workflow step log creation date and time. | [optional] [readonly] 
**Log** | Pointer to **string** | Workflow step log message. | [optional] [readonly] 

## Methods

### NewWorkflowStepLogOutputItem

`func NewWorkflowStepLogOutputItem() *WorkflowStepLogOutputItem`

NewWorkflowStepLogOutputItem instantiates a new WorkflowStepLogOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepLogOutputItemWithDefaults

`func NewWorkflowStepLogOutputItemWithDefaults() *WorkflowStepLogOutputItem`

NewWorkflowStepLogOutputItemWithDefaults instantiates a new WorkflowStepLogOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *WorkflowStepLogOutputItem) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WorkflowStepLogOutputItem) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WorkflowStepLogOutputItem) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *WorkflowStepLogOutputItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetLog

`func (o *WorkflowStepLogOutputItem) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *WorkflowStepLogOutputItem) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *WorkflowStepLogOutputItem) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *WorkflowStepLogOutputItem) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


