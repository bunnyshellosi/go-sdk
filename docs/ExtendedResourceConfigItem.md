# ExtendedResourceConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extended** | Pointer to **bool** |  | [optional] [readonly] 
**Namespace** | Pointer to **string** | The resource namespace. | [optional] [readonly] 
**Kind** | Pointer to **string** | The resource kind. | [optional] [readonly] 
**Name** | Pointer to **string** | The resource name. | [optional] [readonly] 
**Containers** | Pointer to [**map[string]ContainerConfigItem**](ContainerConfigItem.md) | A list of container configs. | [optional] [readonly] 

## Methods

### NewExtendedResourceConfigItem

`func NewExtendedResourceConfigItem() *ExtendedResourceConfigItem`

NewExtendedResourceConfigItem instantiates a new ExtendedResourceConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedResourceConfigItemWithDefaults

`func NewExtendedResourceConfigItemWithDefaults() *ExtendedResourceConfigItem`

NewExtendedResourceConfigItemWithDefaults instantiates a new ExtendedResourceConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtended

`func (o *ExtendedResourceConfigItem) GetExtended() bool`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *ExtendedResourceConfigItem) GetExtendedOk() (*bool, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *ExtendedResourceConfigItem) SetExtended(v bool)`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *ExtendedResourceConfigItem) HasExtended() bool`

HasExtended returns a boolean if a field has been set.

### GetNamespace

`func (o *ExtendedResourceConfigItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ExtendedResourceConfigItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ExtendedResourceConfigItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ExtendedResourceConfigItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetKind

`func (o *ExtendedResourceConfigItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExtendedResourceConfigItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExtendedResourceConfigItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExtendedResourceConfigItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ExtendedResourceConfigItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedResourceConfigItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedResourceConfigItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedResourceConfigItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContainers

`func (o *ExtendedResourceConfigItem) GetContainers() map[string]ContainerConfigItem`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ExtendedResourceConfigItem) GetContainersOk() (*map[string]ContainerConfigItem, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ExtendedResourceConfigItem) SetContainers(v map[string]ContainerConfigItem)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ExtendedResourceConfigItem) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


