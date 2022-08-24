# ContextWrapperKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Context** | Pointer to [**ContextKubeConfigRead**](ContextKubeConfigRead.md) |  | [optional] 

## Methods

### NewContextWrapperKubeConfigRead

`func NewContextWrapperKubeConfigRead(name string, ) *ContextWrapperKubeConfigRead`

NewContextWrapperKubeConfigRead instantiates a new ContextWrapperKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWrapperKubeConfigReadWithDefaults

`func NewContextWrapperKubeConfigReadWithDefaults() *ContextWrapperKubeConfigRead`

NewContextWrapperKubeConfigReadWithDefaults instantiates a new ContextWrapperKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContextWrapperKubeConfigRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextWrapperKubeConfigRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextWrapperKubeConfigRead) SetName(v string)`

SetName sets Name field to given value.


### GetContext

`func (o *ContextWrapperKubeConfigRead) GetContext() ContextKubeConfigRead`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContextWrapperKubeConfigRead) GetContextOk() (*ContextKubeConfigRead, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContextWrapperKubeConfigRead) SetContext(v ContextKubeConfigRead)`

SetContext sets Context field to given value.

### HasContext

`func (o *ContextWrapperKubeConfigRead) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


