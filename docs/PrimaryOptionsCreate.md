# PrimaryOptionsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhitelistEnabled** | Pointer to **bool** |  | [optional] 
**BranchWhitelist** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrimaryOptionsCreate

`func NewPrimaryOptionsCreate() *PrimaryOptionsCreate`

NewPrimaryOptionsCreate instantiates a new PrimaryOptionsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryOptionsCreateWithDefaults

`func NewPrimaryOptionsCreateWithDefaults() *PrimaryOptionsCreate`

NewPrimaryOptionsCreateWithDefaults instantiates a new PrimaryOptionsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelistEnabled

`func (o *PrimaryOptionsCreate) GetWhitelistEnabled() bool`

GetWhitelistEnabled returns the WhitelistEnabled field if non-nil, zero value otherwise.

### GetWhitelistEnabledOk

`func (o *PrimaryOptionsCreate) GetWhitelistEnabledOk() (*bool, bool)`

GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEnabled

`func (o *PrimaryOptionsCreate) SetWhitelistEnabled(v bool)`

SetWhitelistEnabled sets WhitelistEnabled field to given value.

### HasWhitelistEnabled

`func (o *PrimaryOptionsCreate) HasWhitelistEnabled() bool`

HasWhitelistEnabled returns a boolean if a field has been set.

### GetBranchWhitelist

`func (o *PrimaryOptionsCreate) GetBranchWhitelist() string`

GetBranchWhitelist returns the BranchWhitelist field if non-nil, zero value otherwise.

### GetBranchWhitelistOk

`func (o *PrimaryOptionsCreate) GetBranchWhitelistOk() (*string, bool)`

GetBranchWhitelistOk returns a tuple with the BranchWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchWhitelist

`func (o *PrimaryOptionsCreate) SetBranchWhitelist(v string)`

SetBranchWhitelist sets BranchWhitelist field to given value.

### HasBranchWhitelist

`func (o *PrimaryOptionsCreate) HasBranchWhitelist() bool`

HasBranchWhitelist returns a boolean if a field has been set.

### SetBranchWhitelistNil

`func (o *PrimaryOptionsCreate) SetBranchWhitelistNil(b bool)`

 SetBranchWhitelistNil sets the value for BranchWhitelist to be an explicit nil

### UnsetBranchWhitelist
`func (o *PrimaryOptionsCreate) UnsetBranchWhitelist()`

UnsetBranchWhitelist ensures that no value is present for BranchWhitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


