# ComponentProfileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | The command to be run when starting the container. | [optional] [readonly] 
**PortMapping** | Pointer to **[]string** | The port mapping for the container. | [optional] [readonly] 
**Environ** | Pointer to **map[string]string** | The environ for the container. | [optional] [readonly] 
**SyncPaths** | Pointer to [**[]SyncPathItem**](SyncPathItem.md) | The sync paths for the container. | [optional] [readonly] 
**Requirements** | Pointer to [**NullableProfileItemRequirements**](ProfileItemRequirements.md) |  | [optional] 

## Methods

### NewComponentProfileItem

`func NewComponentProfileItem() *ComponentProfileItem`

NewComponentProfileItem instantiates a new ComponentProfileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentProfileItemWithDefaults

`func NewComponentProfileItemWithDefaults() *ComponentProfileItem`

NewComponentProfileItemWithDefaults instantiates a new ComponentProfileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ComponentProfileItem) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ComponentProfileItem) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ComponentProfileItem) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ComponentProfileItem) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetPortMapping

`func (o *ComponentProfileItem) GetPortMapping() []string`

GetPortMapping returns the PortMapping field if non-nil, zero value otherwise.

### GetPortMappingOk

`func (o *ComponentProfileItem) GetPortMappingOk() (*[]string, bool)`

GetPortMappingOk returns a tuple with the PortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMapping

`func (o *ComponentProfileItem) SetPortMapping(v []string)`

SetPortMapping sets PortMapping field to given value.

### HasPortMapping

`func (o *ComponentProfileItem) HasPortMapping() bool`

HasPortMapping returns a boolean if a field has been set.

### GetEnviron

`func (o *ComponentProfileItem) GetEnviron() map[string]string`

GetEnviron returns the Environ field if non-nil, zero value otherwise.

### GetEnvironOk

`func (o *ComponentProfileItem) GetEnvironOk() (*map[string]string, bool)`

GetEnvironOk returns a tuple with the Environ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnviron

`func (o *ComponentProfileItem) SetEnviron(v map[string]string)`

SetEnviron sets Environ field to given value.

### HasEnviron

`func (o *ComponentProfileItem) HasEnviron() bool`

HasEnviron returns a boolean if a field has been set.

### GetSyncPaths

`func (o *ComponentProfileItem) GetSyncPaths() []SyncPathItem`

GetSyncPaths returns the SyncPaths field if non-nil, zero value otherwise.

### GetSyncPathsOk

`func (o *ComponentProfileItem) GetSyncPathsOk() (*[]SyncPathItem, bool)`

GetSyncPathsOk returns a tuple with the SyncPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPaths

`func (o *ComponentProfileItem) SetSyncPaths(v []SyncPathItem)`

SetSyncPaths sets SyncPaths field to given value.

### HasSyncPaths

`func (o *ComponentProfileItem) HasSyncPaths() bool`

HasSyncPaths returns a boolean if a field has been set.

### GetRequirements

`func (o *ComponentProfileItem) GetRequirements() ProfileItemRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ComponentProfileItem) GetRequirementsOk() (*ProfileItemRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ComponentProfileItem) SetRequirements(v ProfileItemRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ComponentProfileItem) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *ComponentProfileItem) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *ComponentProfileItem) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


