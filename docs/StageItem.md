# StageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Stage identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Stage name. | [optional] [readonly] 
**Status** | Pointer to **string** | Stage status. | [optional] [readonly] 
**Duration** | Pointer to **NullableInt32** | Stage duration. | [optional] [readonly] 
**JobsCount** | Pointer to **int32** | Stage total jobs. | [optional] [readonly] 
**CompletedJobsCount** | Pointer to **int32** | Stage completed jobs. | [optional] [readonly] 
**Pipeline** | Pointer to **string** | Event identifier. | [optional] [readonly] 

## Methods

### NewStageItem

`func NewStageItem() *StageItem`

NewStageItem instantiates a new StageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageItemWithDefaults

`func NewStageItemWithDefaults() *StageItem`

NewStageItemWithDefaults instantiates a new StageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StageItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StageItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StageItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StageItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StageItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *StageItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StageItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuration

`func (o *StageItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *StageItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *StageItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *StageItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *StageItem) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *StageItem) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetJobsCount

`func (o *StageItem) GetJobsCount() int32`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *StageItem) GetJobsCountOk() (*int32, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *StageItem) SetJobsCount(v int32)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *StageItem) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetCompletedJobsCount

`func (o *StageItem) GetCompletedJobsCount() int32`

GetCompletedJobsCount returns the CompletedJobsCount field if non-nil, zero value otherwise.

### GetCompletedJobsCountOk

`func (o *StageItem) GetCompletedJobsCountOk() (*int32, bool)`

GetCompletedJobsCountOk returns a tuple with the CompletedJobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobsCount

`func (o *StageItem) SetCompletedJobsCount(v int32)`

SetCompletedJobsCount sets CompletedJobsCount field to given value.

### HasCompletedJobsCount

`func (o *StageItem) HasCompletedJobsCount() bool`

HasCompletedJobsCount returns a boolean if a field has been set.

### GetPipeline

`func (o *StageItem) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *StageItem) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *StageItem) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *StageItem) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


