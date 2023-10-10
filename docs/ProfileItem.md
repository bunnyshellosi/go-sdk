# ProfileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | The command to be run when starting the container. | [optional] [readonly] 
**PortMapping** | Pointer to **[]string** | The port mapping for the container. | [optional] [readonly] 
**Environ** | Pointer to **map[string]string** | The environ for the container. | [optional] [readonly] 
**SyncPaths** | Pointer to [**[]SyncPathItem**](SyncPathItem.md) | The sync paths for the container. | [optional] [readonly] 
**Requirements** | Pointer to [**NullableResourceRequirementItem**](ResourceRequirementItem.md) |  | [optional] 

## Methods

### NewProfileItem

`func NewProfileItem() *ProfileItem`

NewProfileItem instantiates a new ProfileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileItemWithDefaults

`func NewProfileItemWithDefaults() *ProfileItem`

NewProfileItemWithDefaults instantiates a new ProfileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ProfileItem) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ProfileItem) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ProfileItem) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ProfileItem) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetPortMapping

`func (o *ProfileItem) GetPortMapping() []string`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *ProfileItem) GetPortMappingOk() (*[]string, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *ProfileItem) SetPortMapping(v []string)`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *ProfileItem) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetEnviron

`func (o *ProfileItem) GetEnviron() map[string]string`

GetEnviron returns the Environ field if non-nil, zero value otherwise.

### GetEnvironOk

`func (o *ProfileItem) GetEnvironOk() (*map[string]string, bool)`

GetEnvironOk returns a tuple with the Environ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnviron

`func (o *ProfileItem) SetEnviron(v map[string]string)`

SetEnviron sets Environ field to given value.

### HasEnviron

`func (o *ProfileItem) HasEnviron() bool`

HasEnviron returns a boolean if a field has been set.

### GetSyncPaths

`func (o *ProfileItem) GetSyncPaths() []SyncPathItem`

GetSyncPaths returns the SyncPaths field if non-nil, zero value otherwise.

### GetSyncPathsOk

`func (o *ProfileItem) GetSyncPathsOk() (*[]SyncPathItem, bool)`

GetSyncPathsOk returns a tuple with the SyncPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPaths

`func (o *ProfileItem) SetSyncPaths(v []SyncPathItem)`

SetSyncPaths sets SyncPaths field to given value.

### HasSyncPaths

`func (o *ProfileItem) HasSyncPaths() bool`

HasSyncPaths returns a boolean if a field has been set.

### GetRequirements

`func (o *ProfileItem) GetRequirements() ResourceRequirementItem`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ProfileItem) GetRequirementsOk() (*ResourceRequirementItem, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ProfileItem) SetRequirements(v ResourceRequirementItem)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ProfileItem) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *ProfileItem) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *ProfileItem) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


