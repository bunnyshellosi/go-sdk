# SimpleResourceConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Simple** | Pointer to **bool** |  | [optional] [readonly] 
**Containers** | Pointer to [**map[string]ContainerConfigItem**](ContainerConfigItem.md) | A list of container configs. | [optional] [readonly] 

## Methods

### NewSimpleResourceConfigItem

`func NewSimpleResourceConfigItem() *SimpleResourceConfigItem`

NewSimpleResourceConfigItem instantiates a new SimpleResourceConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleResourceConfigItemWithDefaults

`func NewSimpleResourceConfigItemWithDefaults() *SimpleResourceConfigItem`

NewSimpleResourceConfigItemWithDefaults instantiates a new SimpleResourceConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimple

`func (o *SimpleResourceConfigItem) GetSimple() bool`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *SimpleResourceConfigItem) GetSimpleOk() (*bool, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *SimpleResourceConfigItem) SetSimple(v bool)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *SimpleResourceConfigItem) HasSimple() bool`

HasSimple returns a boolean if a field has been set.

### GetContainers

`func (o *SimpleResourceConfigItem) GetContainers() map[string]ContainerConfigItem`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *SimpleResourceConfigItem) GetContainersOk() (*map[string]ContainerConfigItem, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *SimpleResourceConfigItem) SetContainers(v map[string]ContainerConfigItem)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *SimpleResourceConfigItem) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


