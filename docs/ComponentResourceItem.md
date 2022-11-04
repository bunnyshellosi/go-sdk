# ComponentResourceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kubernetes resource kind. | [optional] 
**Name** | Pointer to **string** | Kubernetes resource name. | [optional] 
**Namespace** | Pointer to **NullableString** | Kubernetes resource namespace. | [optional] 

## Methods

### NewComponentResourceItem

`func NewComponentResourceItem() *ComponentResourceItem`

NewComponentResourceItem instantiates a new ComponentResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentResourceItemWithDefaults

`func NewComponentResourceItemWithDefaults() *ComponentResourceItem`

NewComponentResourceItemWithDefaults instantiates a new ComponentResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ComponentResourceItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComponentResourceItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComponentResourceItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComponentResourceItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ComponentResourceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentResourceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentResourceItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentResourceItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ComponentResourceItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ComponentResourceItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ComponentResourceItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ComponentResourceItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ComponentResourceItem) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ComponentResourceItem) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


