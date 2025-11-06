# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "workspace"]
**TerminationProtection** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace() *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Workspace) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Workspace) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Workspace) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Workspace) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerminationProtection

`func (o *Workspace) GetTerminationProtection() bool`

GetTerminationProtection returns the TerminationProtection field if non-nil, zero value otherwise.

### GetTerminationProtectionOk

`func (o *Workspace) GetTerminationProtectionOk() (*bool, bool)`

GetTerminationProtectionOk returns a tuple with the TerminationProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtection

`func (o *Workspace) SetTerminationProtection(v bool)`

SetTerminationProtection sets TerminationProtection field to given value.

### HasTerminationProtection

`func (o *Workspace) HasTerminationProtection() bool`

HasTerminationProtection returns a boolean if a field has been set.

### SetTerminationProtectionNil

`func (o *Workspace) SetTerminationProtectionNil(b bool)`

 SetTerminationProtectionNil sets the value for TerminationProtection to be an explicit nil

### UnsetTerminationProtection
`func (o *Workspace) UnsetTerminationProtection()`

UnsetTerminationProtection ensures that no value is present for TerminationProtection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


