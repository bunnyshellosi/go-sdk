# ComponentManageKeysAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** |  | 
**SshKeyName** | **string** |  | 
**SshKeyValue** | Pointer to **NullableString** |  | [optional] 
**WithVMSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewComponentManageKeysAction

`func NewComponentManageKeysAction(operation string, sshKeyName string, ) *ComponentManageKeysAction`

NewComponentManageKeysAction instantiates a new ComponentManageKeysAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentManageKeysActionWithDefaults

`func NewComponentManageKeysActionWithDefaults() *ComponentManageKeysAction`

NewComponentManageKeysActionWithDefaults instantiates a new ComponentManageKeysAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ComponentManageKeysAction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ComponentManageKeysAction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ComponentManageKeysAction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetSshKeyName

`func (o *ComponentManageKeysAction) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *ComponentManageKeysAction) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *ComponentManageKeysAction) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetSshKeyValue

`func (o *ComponentManageKeysAction) GetSshKeyValue() string`

GetSshKeyValue returns the SshKeyValue field if non-nil, zero value otherwise.

### GetSshKeyValueOk

`func (o *ComponentManageKeysAction) GetSshKeyValueOk() (*string, bool)`

GetSshKeyValueOk returns a tuple with the SshKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyValue

`func (o *ComponentManageKeysAction) SetSshKeyValue(v string)`

SetSshKeyValue sets SshKeyValue field to given value.

### HasSshKeyValue

`func (o *ComponentManageKeysAction) HasSshKeyValue() bool`

HasSshKeyValue returns a boolean if a field has been set.

### SetSshKeyValueNil

`func (o *ComponentManageKeysAction) SetSshKeyValueNil(b bool)`

 SetSshKeyValueNil sets the value for SshKeyValue to be an explicit nil

### UnsetSshKeyValue
`func (o *ComponentManageKeysAction) UnsetSshKeyValue()`

UnsetSshKeyValue ensures that no value is present for SshKeyValue, not even an explicit nil
### GetWithVMSync

`func (o *ComponentManageKeysAction) GetWithVMSync() bool`

GetWithVMSync returns the WithVMSync field if non-nil, zero value otherwise.

### GetWithVMSyncOk

`func (o *ComponentManageKeysAction) GetWithVMSyncOk() (*bool, bool)`

GetWithVMSyncOk returns a tuple with the WithVMSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVMSync

`func (o *ComponentManageKeysAction) SetWithVMSync(v bool)`

SetWithVMSync sets WithVMSync field to given value.

### HasWithVMSync

`func (o *ComponentManageKeysAction) HasWithVMSync() bool`

HasWithVMSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


