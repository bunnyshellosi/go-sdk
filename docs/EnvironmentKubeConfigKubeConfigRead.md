# EnvironmentKubeConfigKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] [default to "v1"]
**Kind** | Pointer to **string** |  | [optional] [default to "Config"]
**Preferences** | Pointer to **map[string]interface{}** |  | [optional] 
**Clusters** | Pointer to [**[]ClusterWrapperKubeConfigRead**](ClusterWrapperKubeConfigRead.md) |  | [optional] 
**Users** | Pointer to [**[]UserWrapperKubeConfigRead**](UserWrapperKubeConfigRead.md) |  | [optional] 
**Contexts** | Pointer to [**[]ContextWrapperKubeConfigRead**](ContextWrapperKubeConfigRead.md) |  | [optional] 
**CurrentContext** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentKubeConfigKubeConfigRead

`func NewEnvironmentKubeConfigKubeConfigRead() *EnvironmentKubeConfigKubeConfigRead`

NewEnvironmentKubeConfigKubeConfigRead instantiates a new EnvironmentKubeConfigKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentKubeConfigKubeConfigReadWithDefaults

`func NewEnvironmentKubeConfigKubeConfigReadWithDefaults() *EnvironmentKubeConfigKubeConfigRead`

NewEnvironmentKubeConfigKubeConfigReadWithDefaults instantiates a new EnvironmentKubeConfigKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EnvironmentKubeConfigKubeConfigRead) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentKubeConfigKubeConfigRead) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EnvironmentKubeConfigKubeConfigRead) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *EnvironmentKubeConfigKubeConfigRead) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentKubeConfigKubeConfigRead) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EnvironmentKubeConfigKubeConfigRead) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPreferences

`func (o *EnvironmentKubeConfigKubeConfigRead) GetPreferences() map[string]interface{}`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetPreferencesOk() (*map[string]interface{}, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *EnvironmentKubeConfigKubeConfigRead) SetPreferences(v map[string]interface{})`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *EnvironmentKubeConfigKubeConfigRead) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetClusters

`func (o *EnvironmentKubeConfigKubeConfigRead) GetClusters() []ClusterWrapperKubeConfigRead`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetClustersOk() (*[]ClusterWrapperKubeConfigRead, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *EnvironmentKubeConfigKubeConfigRead) SetClusters(v []ClusterWrapperKubeConfigRead)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *EnvironmentKubeConfigKubeConfigRead) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetUsers

`func (o *EnvironmentKubeConfigKubeConfigRead) GetUsers() []UserWrapperKubeConfigRead`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetUsersOk() (*[]UserWrapperKubeConfigRead, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EnvironmentKubeConfigKubeConfigRead) SetUsers(v []UserWrapperKubeConfigRead)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EnvironmentKubeConfigKubeConfigRead) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetContexts

`func (o *EnvironmentKubeConfigKubeConfigRead) GetContexts() []ContextWrapperKubeConfigRead`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetContextsOk() (*[]ContextWrapperKubeConfigRead, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *EnvironmentKubeConfigKubeConfigRead) SetContexts(v []ContextWrapperKubeConfigRead)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *EnvironmentKubeConfigKubeConfigRead) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetCurrentContext

`func (o *EnvironmentKubeConfigKubeConfigRead) GetCurrentContext() string`

GetCurrentContext returns the CurrentContext field if non-nil, zero value otherwise.

### GetCurrentContextOk

`func (o *EnvironmentKubeConfigKubeConfigRead) GetCurrentContextOk() (*string, bool)`

GetCurrentContextOk returns a tuple with the CurrentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentContext

`func (o *EnvironmentKubeConfigKubeConfigRead) SetCurrentContext(v string)`

SetCurrentContext sets CurrentContext field to given value.

### HasCurrentContext

`func (o *EnvironmentKubeConfigKubeConfigRead) HasCurrentContext() bool`

HasCurrentContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


