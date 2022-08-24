# ComponentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Service component name | [optional] [readonly] 
**ClusterStatus** | Pointer to **string** | Service component cluster status | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Service component operation status | [optional] [readonly] 
**PublicURLs** | Pointer to **[]string** | Service component URLs | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**SyncPath** | Pointer to **NullableString** | Service component sync path for remote development | [optional] [readonly] 

## Methods

### NewComponentItem

`func NewComponentItem() *ComponentItem`

NewComponentItem instantiates a new ComponentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentItemWithDefaults

`func NewComponentItemWithDefaults() *ComponentItem`

NewComponentItemWithDefaults instantiates a new ComponentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterStatus

`func (o *ComponentItem) GetClusterStatus() string`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *ComponentItem) GetClusterStatusOk() (*string, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *ComponentItem) SetClusterStatus(v string)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *ComponentItem) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetOperationStatus

`func (o *ComponentItem) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *ComponentItem) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *ComponentItem) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *ComponentItem) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetPublicURLs

`func (o *ComponentItem) GetPublicURLs() []string`

GetPublicURLs returns the PublicURLs field if non-nil, zero value otherwise.

### GetPublicURLsOk

`func (o *ComponentItem) GetPublicURLsOk() (*[]string, bool)`

GetPublicURLsOk returns a tuple with the PublicURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicURLs

`func (o *ComponentItem) SetPublicURLs(v []string)`

SetPublicURLs sets PublicURLs field to given value.

### HasPublicURLs

`func (o *ComponentItem) HasPublicURLs() bool`

HasPublicURLs returns a boolean if a field has been set.

### GetEnvironment

`func (o *ComponentItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ComponentItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ComponentItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ComponentItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSyncPath

`func (o *ComponentItem) GetSyncPath() string`

GetSyncPath returns the SyncPath field if non-nil, zero value otherwise.

### GetSyncPathOk

`func (o *ComponentItem) GetSyncPathOk() (*string, bool)`

GetSyncPathOk returns a tuple with the SyncPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPath

`func (o *ComponentItem) SetSyncPath(v string)`

SetSyncPath sets SyncPath field to given value.

### HasSyncPath

`func (o *ComponentItem) HasSyncPath() bool`

HasSyncPath returns a boolean if a field has been set.

### SetSyncPathNil

`func (o *ComponentItem) SetSyncPathNil(b bool)`

 SetSyncPathNil sets the value for SyncPath to be an explicit nil

### UnsetSyncPath
`func (o *ComponentItem) UnsetSyncPath()`

UnsetSyncPath ensures that no value is present for SyncPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


