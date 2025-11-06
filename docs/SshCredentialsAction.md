# SshCredentialsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentSshHost** | **string** |  | 
**WithSSHKeysFingerprints** | Pointer to **bool** |  | [optional] 

## Methods

### NewSshCredentialsAction

`func NewSshCredentialsAction(componentSshHost string, ) *SshCredentialsAction`

NewSshCredentialsAction instantiates a new SshCredentialsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshCredentialsActionWithDefaults

`func NewSshCredentialsActionWithDefaults() *SshCredentialsAction`

NewSshCredentialsActionWithDefaults instantiates a new SshCredentialsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentSshHost

`func (o *SshCredentialsAction) GetComponentSshHost() string`

GetComponentSshHost returns the ComponentSshHost field if non-nil, zero value otherwise.

### GetComponentSshHostOk

`func (o *SshCredentialsAction) GetComponentSshHostOk() (*string, bool)`

GetComponentSshHostOk returns a tuple with the ComponentSshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentSshHost

`func (o *SshCredentialsAction) SetComponentSshHost(v string)`

SetComponentSshHost sets ComponentSshHost field to given value.


### GetWithSSHKeysFingerprints

`func (o *SshCredentialsAction) GetWithSSHKeysFingerprints() bool`

GetWithSSHKeysFingerprints returns the WithSSHKeysFingerprints field if non-nil, zero value otherwise.

### GetWithSSHKeysFingerprintsOk

`func (o *SshCredentialsAction) GetWithSSHKeysFingerprintsOk() (*bool, bool)`

GetWithSSHKeysFingerprintsOk returns a tuple with the WithSSHKeysFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithSSHKeysFingerprints

`func (o *SshCredentialsAction) SetWithSSHKeysFingerprints(v bool)`

SetWithSSHKeysFingerprints sets WithSSHKeysFingerprints field to given value.

### HasWithSSHKeysFingerprints

`func (o *SshCredentialsAction) HasWithSSHKeysFingerprints() bool`

HasWithSSHKeysFingerprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


