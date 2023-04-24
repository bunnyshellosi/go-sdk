# ResourceRequirementItemRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **NullableString** | The CPU resources for the container. | [optional] [readonly] 
**Memory** | Pointer to **NullableString** | The Memory resources for the container. | [optional] [readonly] 

## Methods

### NewResourceRequirementItemRequests

`func NewResourceRequirementItemRequests() *ResourceRequirementItemRequests`

NewResourceRequirementItemRequests instantiates a new ResourceRequirementItemRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementItemRequestsWithDefaults

`func NewResourceRequirementItemRequestsWithDefaults() *ResourceRequirementItemRequests`

NewResourceRequirementItemRequestsWithDefaults instantiates a new ResourceRequirementItemRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceRequirementItemRequests) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceRequirementItemRequests) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceRequirementItemRequests) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceRequirementItemRequests) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ResourceRequirementItemRequests) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ResourceRequirementItemRequests) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ResourceRequirementItemRequests) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceRequirementItemRequests) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceRequirementItemRequests) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceRequirementItemRequests) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ResourceRequirementItemRequests) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ResourceRequirementItemRequests) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


