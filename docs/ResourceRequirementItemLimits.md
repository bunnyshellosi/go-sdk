# ResourceRequirementItemLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **NullableString** | The CPU resources for the container. | [optional] [readonly] 
**Memory** | Pointer to **NullableString** | The Memory resources for the container. | [optional] [readonly] 

## Methods

### NewResourceRequirementItemLimits

`func NewResourceRequirementItemLimits() *ResourceRequirementItemLimits`

NewResourceRequirementItemLimits instantiates a new ResourceRequirementItemLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementItemLimitsWithDefaults

`func NewResourceRequirementItemLimitsWithDefaults() *ResourceRequirementItemLimits`

NewResourceRequirementItemLimitsWithDefaults instantiates a new ResourceRequirementItemLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceRequirementItemLimits) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceRequirementItemLimits) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceRequirementItemLimits) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceRequirementItemLimits) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ResourceRequirementItemLimits) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ResourceRequirementItemLimits) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ResourceRequirementItemLimits) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceRequirementItemLimits) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceRequirementItemLimits) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceRequirementItemLimits) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ResourceRequirementItemLimits) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ResourceRequirementItemLimits) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


