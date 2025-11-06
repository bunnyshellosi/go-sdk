# ComponentSSHCredentialsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The SSH domain. | [optional] [readonly] 
**Port** | Pointer to **int32** | The SSH port. | [optional] [readonly] 
**User** | Pointer to **string** | The SSH username. | [optional] [readonly] 
**SshKeyFingerprints** | Pointer to **map[string]string** | The SSH keys fingerprints. | [optional] [readonly] 
**UserCanManageSSHKeys** | Pointer to **bool** | If the user can manage SSH keys. | [optional] [readonly] 
**EnvironmentId** | Pointer to **string** | The environment identifier. | [optional] [readonly] 
**ComponentId** | Pointer to **string** | The service component identifier. | [optional] [readonly] 

## Methods

### NewComponentSSHCredentialsItem

`func NewComponentSSHCredentialsItem() *ComponentSSHCredentialsItem`

NewComponentSSHCredentialsItem instantiates a new ComponentSSHCredentialsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentSSHCredentialsItemWithDefaults

`func NewComponentSSHCredentialsItemWithDefaults() *ComponentSSHCredentialsItem`

NewComponentSSHCredentialsItemWithDefaults instantiates a new ComponentSSHCredentialsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *ComponentSSHCredentialsItem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ComponentSSHCredentialsItem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ComponentSSHCredentialsItem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ComponentSSHCredentialsItem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *ComponentSSHCredentialsItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ComponentSSHCredentialsItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ComponentSSHCredentialsItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ComponentSSHCredentialsItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUser

`func (o *ComponentSSHCredentialsItem) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ComponentSSHCredentialsItem) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ComponentSSHCredentialsItem) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ComponentSSHCredentialsItem) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSshKeyFingerprints

`func (o *ComponentSSHCredentialsItem) GetSshKeyFingerprints() map[string]string`

GetSshKeyFingerprints returns the SshKeyFingerprints field if non-nil, zero value otherwise.

### GetSshKeyFingerprintsOk

`func (o *ComponentSSHCredentialsItem) GetSshKeyFingerprintsOk() (*map[string]string, bool)`

GetSshKeyFingerprintsOk returns a tuple with the SshKeyFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyFingerprints

`func (o *ComponentSSHCredentialsItem) SetSshKeyFingerprints(v map[string]string)`

SetSshKeyFingerprints sets SshKeyFingerprints field to given value.

### HasSshKeyFingerprints

`func (o *ComponentSSHCredentialsItem) HasSshKeyFingerprints() bool`

HasSshKeyFingerprints returns a boolean if a field has been set.

### GetUserCanManageSSHKeys

`func (o *ComponentSSHCredentialsItem) GetUserCanManageSSHKeys() bool`

GetUserCanManageSSHKeys returns the UserCanManageSSHKeys field if non-nil, zero value otherwise.

### GetUserCanManageSSHKeysOk

`func (o *ComponentSSHCredentialsItem) GetUserCanManageSSHKeysOk() (*bool, bool)`

GetUserCanManageSSHKeysOk returns a tuple with the UserCanManageSSHKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCanManageSSHKeys

`func (o *ComponentSSHCredentialsItem) SetUserCanManageSSHKeys(v bool)`

SetUserCanManageSSHKeys sets UserCanManageSSHKeys field to given value.

### HasUserCanManageSSHKeys

`func (o *ComponentSSHCredentialsItem) HasUserCanManageSSHKeys() bool`

HasUserCanManageSSHKeys returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ComponentSSHCredentialsItem) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ComponentSSHCredentialsItem) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ComponentSSHCredentialsItem) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ComponentSSHCredentialsItem) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetComponentId

`func (o *ComponentSSHCredentialsItem) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentSSHCredentialsItem) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentSSHCredentialsItem) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentSSHCredentialsItem) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


