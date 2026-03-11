# WorkflowJobWorkflowJobLogsOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowJob** | Pointer to [**WorkflowJobOutputItem**](WorkflowJobOutputItem.md) |  | [optional] 
**Steps** | Pointer to [**[]WorkflowStepOutputItem**](WorkflowStepOutputItem.md) | Workflow job steps and logs. | [optional] [readonly] 

## Methods

### NewWorkflowJobWorkflowJobLogsOutputItem

`func NewWorkflowJobWorkflowJobLogsOutputItem() *WorkflowJobWorkflowJobLogsOutputItem`

NewWorkflowJobWorkflowJobLogsOutputItem instantiates a new WorkflowJobWorkflowJobLogsOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJobWorkflowJobLogsOutputItemWithDefaults

`func NewWorkflowJobWorkflowJobLogsOutputItemWithDefaults() *WorkflowJobWorkflowJobLogsOutputItem`

NewWorkflowJobWorkflowJobLogsOutputItemWithDefaults instantiates a new WorkflowJobWorkflowJobLogsOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowJob

`func (o *WorkflowJobWorkflowJobLogsOutputItem) GetWorkflowJob() WorkflowJobOutputItem`

GetWorkflowJob returns the WorkflowJob field if non-nil, zero value otherwise.

### GetWorkflowJobOk

`func (o *WorkflowJobWorkflowJobLogsOutputItem) GetWorkflowJobOk() (*WorkflowJobOutputItem, bool)`

GetWorkflowJobOk returns a tuple with the WorkflowJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowJob

`func (o *WorkflowJobWorkflowJobLogsOutputItem) SetWorkflowJob(v WorkflowJobOutputItem)`

SetWorkflowJob sets WorkflowJob field to given value.

### HasWorkflowJob

`func (o *WorkflowJobWorkflowJobLogsOutputItem) HasWorkflowJob() bool`

HasWorkflowJob returns a boolean if a field has been set.

### GetSteps

`func (o *WorkflowJobWorkflowJobLogsOutputItem) GetSteps() []WorkflowStepOutputItem`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowJobWorkflowJobLogsOutputItem) GetStepsOk() (*[]WorkflowStepOutputItem, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowJobWorkflowJobLogsOutputItem) SetSteps(v []WorkflowStepOutputItem)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkflowJobWorkflowJobLogsOutputItem) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


