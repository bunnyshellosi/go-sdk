# SyncPathItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPath** | Pointer to **NullableString** | The local path on the host os to sync from. | [optional] [readonly] 
**RemotePath** | Pointer to **string** | The remote path on the container to sync to. | [optional] [readonly] 

## Methods

### NewSyncPathItem

`func NewSyncPathItem() *SyncPathItem`

NewSyncPathItem instantiates a new SyncPathItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPathItemWithDefaults

`func NewSyncPathItemWithDefaults() *SyncPathItem`

NewSyncPathItemWithDefaults instantiates a new SyncPathItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPath

`func (o *SyncPathItem) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *SyncPathItem) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *SyncPathItem) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.

### HasLocalPath

`func (o *SyncPathItem) HasLocalPath() bool`

HasLocalPath returns a boolean if a field has been set.

### SetLocalPathNil

`func (o *SyncPathItem) SetLocalPathNil(b bool)`

 SetLocalPathNil sets the value for LocalPath to be an explicit nil

### UnsetLocalPath
`func (o *SyncPathItem) UnsetLocalPath()`

UnsetLocalPath ensures that no value is present for LocalPath, not even an explicit nil
### GetRemotePath

`func (o *SyncPathItem) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *SyncPathItem) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *SyncPathItem) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *SyncPathItem) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


