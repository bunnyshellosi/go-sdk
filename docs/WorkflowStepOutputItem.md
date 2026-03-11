# WorkflowStepOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Workflow step name. | [optional] [readonly] 
**Status** | Pointer to **string** | Workflow step status. | [optional] [readonly] 
**ExitCode** | Pointer to **int32** | Workflow step exit code. | [optional] [readonly] 
**Time** | Pointer to **time.Time** | Workflow step execution date and time. | [optional] [readonly] 
**Logs** | Pointer to [**[]WorkflowStepLogOutputItem**](WorkflowStepLogOutputItem.md) | Workflow step logs. | [optional] [readonly] 

## Methods

### NewWorkflowStepOutputItem

`func NewWorkflowStepOutputItem() *WorkflowStepOutputItem`

NewWorkflowStepOutputItem instantiates a new WorkflowStepOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepOutputItemWithDefaults

`func NewWorkflowStepOutputItemWithDefaults() *WorkflowStepOutputItem`

NewWorkflowStepOutputItemWithDefaults instantiates a new WorkflowStepOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowStepOutputItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStepOutputItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStepOutputItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowStepOutputItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowStepOutputItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowStepOutputItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowStepOutputItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowStepOutputItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExitCode

`func (o *WorkflowStepOutputItem) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *WorkflowStepOutputItem) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *WorkflowStepOutputItem) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *WorkflowStepOutputItem) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetTime

`func (o *WorkflowStepOutputItem) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WorkflowStepOutputItem) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WorkflowStepOutputItem) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *WorkflowStepOutputItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetLogs

`func (o *WorkflowStepOutputItem) GetLogs() []WorkflowStepLogOutputItem`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *WorkflowStepOutputItem) GetLogsOk() (*[]WorkflowStepLogOutputItem, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *WorkflowStepOutputItem) SetLogs(v []WorkflowStepLogOutputItem)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *WorkflowStepOutputItem) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


