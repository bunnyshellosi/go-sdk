# ClusterWrapperKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cluster** | Pointer to [**ClusterKubeConfigRead**](ClusterKubeConfigRead.md) |  | [optional] 

## Methods

### NewClusterWrapperKubeConfigRead

`func NewClusterWrapperKubeConfigRead(name string, ) *ClusterWrapperKubeConfigRead`

NewClusterWrapperKubeConfigRead instantiates a new ClusterWrapperKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWrapperKubeConfigReadWithDefaults

`func NewClusterWrapperKubeConfigReadWithDefaults() *ClusterWrapperKubeConfigRead`

NewClusterWrapperKubeConfigReadWithDefaults instantiates a new ClusterWrapperKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterWrapperKubeConfigRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterWrapperKubeConfigRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterWrapperKubeConfigRead) SetName(v string)`

SetName sets Name field to given value.


### GetCluster

`func (o *ClusterWrapperKubeConfigRead) GetCluster() ClusterKubeConfigRead`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterWrapperKubeConfigRead) GetClusterOk() (*ClusterKubeConfigRead, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterWrapperKubeConfigRead) SetCluster(v ClusterKubeConfigRead)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterWrapperKubeConfigRead) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


