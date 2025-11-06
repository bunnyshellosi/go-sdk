# SshSSHCredentialsItem

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

### NewSshSSHCredentialsItem

`func NewSshSSHCredentialsItem() *SshSSHCredentialsItem`

NewSshSSHCredentialsItem instantiates a new SshSSHCredentialsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSSHCredentialsItemWithDefaults

`func NewSshSSHCredentialsItemWithDefaults() *SshSSHCredentialsItem`

NewSshSSHCredentialsItemWithDefaults instantiates a new SshSSHCredentialsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *SshSSHCredentialsItem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SshSSHCredentialsItem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SshSSHCredentialsItem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SshSSHCredentialsItem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *SshSSHCredentialsItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SshSSHCredentialsItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SshSSHCredentialsItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SshSSHCredentialsItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUser

`func (o *SshSSHCredentialsItem) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SshSSHCredentialsItem) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SshSSHCredentialsItem) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SshSSHCredentialsItem) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSshKeyFingerprints

`func (o *SshSSHCredentialsItem) GetSshKeyFingerprints() map[string]string`

GetSshKeyFingerprints returns the SshKeyFingerprints field if non-nil, zero value otherwise.

### GetSshKeyFingerprintsOk

`func (o *SshSSHCredentialsItem) GetSshKeyFingerprintsOk() (*map[string]string, bool)`

GetSshKeyFingerprintsOk returns a tuple with the SshKeyFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyFingerprints

`func (o *SshSSHCredentialsItem) SetSshKeyFingerprints(v map[string]string)`

SetSshKeyFingerprints sets SshKeyFingerprints field to given value.

### HasSshKeyFingerprints

`func (o *SshSSHCredentialsItem) HasSshKeyFingerprints() bool`

HasSshKeyFingerprints returns a boolean if a field has been set.

### GetUserCanManageSSHKeys

`func (o *SshSSHCredentialsItem) GetUserCanManageSSHKeys() bool`

GetUserCanManageSSHKeys returns the UserCanManageSSHKeys field if non-nil, zero value otherwise.

### GetUserCanManageSSHKeysOk

`func (o *SshSSHCredentialsItem) GetUserCanManageSSHKeysOk() (*bool, bool)`

GetUserCanManageSSHKeysOk returns a tuple with the UserCanManageSSHKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCanManageSSHKeys

`func (o *SshSSHCredentialsItem) SetUserCanManageSSHKeys(v bool)`

SetUserCanManageSSHKeys sets UserCanManageSSHKeys field to given value.

### HasUserCanManageSSHKeys

`func (o *SshSSHCredentialsItem) HasUserCanManageSSHKeys() bool`

HasUserCanManageSSHKeys returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SshSSHCredentialsItem) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SshSSHCredentialsItem) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SshSSHCredentialsItem) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SshSSHCredentialsItem) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetComponentId

`func (o *SshSSHCredentialsItem) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *SshSSHCredentialsItem) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *SshSSHCredentialsItem) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *SshSSHCredentialsItem) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


