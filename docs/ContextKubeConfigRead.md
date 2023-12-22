# ContextKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**User** | **string** |  | 

## Methods

### NewContextKubeConfigRead

`func NewContextKubeConfigRead(cluster string, user string, ) *ContextKubeConfigRead`

NewContextKubeConfigRead instantiates a new ContextKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextKubeConfigReadWithDefaults

`func NewContextKubeConfigReadWithDefaults() *ContextKubeConfigRead`

NewContextKubeConfigReadWithDefaults instantiates a new ContextKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ContextKubeConfigRead) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ContextKubeConfigRead) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ContextKubeConfigRead) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetNamespace

`func (o *ContextKubeConfigRead) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ContextKubeConfigRead) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ContextKubeConfigRead) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ContextKubeConfigRead) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUser

`func (o *ContextKubeConfigRead) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ContextKubeConfigRead) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ContextKubeConfigRead) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


