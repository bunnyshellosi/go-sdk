# ComponentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Service component name | [optional] [readonly] 
**ClusterStatus** | Pointer to **string** | Service component cluster status | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Service component operation status | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**SyncPath** | Pointer to **NullableString** | Service component sync path for remote development | [optional] [readonly] 

## Methods

### NewComponentCollection

`func NewComponentCollection() *ComponentCollection`

NewComponentCollection instantiates a new ComponentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentCollectionWithDefaults

`func NewComponentCollectionWithDefaults() *ComponentCollection`

NewComponentCollectionWithDefaults instantiates a new ComponentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterStatus

`func (o *ComponentCollection) GetClusterStatus() string`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *ComponentCollection) GetClusterStatusOk() (*string, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *ComponentCollection) SetClusterStatus(v string)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *ComponentCollection) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetOperationStatus

`func (o *ComponentCollection) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *ComponentCollection) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *ComponentCollection) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *ComponentCollection) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *ComponentCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ComponentCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ComponentCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ComponentCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSyncPath

`func (o *ComponentCollection) GetSyncPath() string`

GetSyncPath returns the SyncPath field if non-nil, zero value otherwise.

### GetSyncPathOk

`func (o *ComponentCollection) GetSyncPathOk() (*string, bool)`

GetSyncPathOk returns a tuple with the SyncPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPath

`func (o *ComponentCollection) SetSyncPath(v string)`

SetSyncPath sets SyncPath field to given value.

### HasSyncPath

`func (o *ComponentCollection) HasSyncPath() bool`

HasSyncPath returns a boolean if a field has been set.

### SetSyncPathNil

`func (o *ComponentCollection) SetSyncPathNil(b bool)`

 SetSyncPathNil sets the value for SyncPath to be an explicit nil

### UnsetSyncPath
`func (o *ComponentCollection) UnsetSyncPath()`

UnsetSyncPath ensures that no value is present for SyncPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


