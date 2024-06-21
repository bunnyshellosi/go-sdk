# EditPrimaryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhitelistEnabled** | Pointer to **NullableBool** |  | [optional] 
**BranchWhitelist** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEditPrimaryOptions

`func NewEditPrimaryOptions() *EditPrimaryOptions`

NewEditPrimaryOptions instantiates a new EditPrimaryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditPrimaryOptionsWithDefaults

`func NewEditPrimaryOptionsWithDefaults() *EditPrimaryOptions`

NewEditPrimaryOptionsWithDefaults instantiates a new EditPrimaryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelistEnabled

`func (o *EditPrimaryOptions) GetWhitelistEnabled() bool`

GetWhitelistEnabled returns the WhitelistEnabled field if non-nil, zero value otherwise.

### GetWhitelistEnabledOk

`func (o *EditPrimaryOptions) GetWhitelistEnabledOk() (*bool, bool)`

GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEnabled

`func (o *EditPrimaryOptions) SetWhitelistEnabled(v bool)`

SetWhitelistEnabled sets WhitelistEnabled field to given value.

### HasWhitelistEnabled

`func (o *EditPrimaryOptions) HasWhitelistEnabled() bool`

HasWhitelistEnabled returns a boolean if a field has been set.

### SetWhitelistEnabledNil

`func (o *EditPrimaryOptions) SetWhitelistEnabledNil(b bool)`

 SetWhitelistEnabledNil sets the value for WhitelistEnabled to be an explicit nil

### UnsetWhitelistEnabled
`func (o *EditPrimaryOptions) UnsetWhitelistEnabled()`

UnsetWhitelistEnabled ensures that no value is present for WhitelistEnabled, not even an explicit nil
### GetBranchWhitelist

`func (o *EditPrimaryOptions) GetBranchWhitelist() string`

GetBranchWhitelist returns the BranchWhitelist field if non-nil, zero value otherwise.

### GetBranchWhitelistOk

`func (o *EditPrimaryOptions) GetBranchWhitelistOk() (*string, bool)`

GetBranchWhitelistOk returns a tuple with the BranchWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchWhitelist

`func (o *EditPrimaryOptions) SetBranchWhitelist(v string)`

SetBranchWhitelist sets BranchWhitelist field to given value.

### HasBranchWhitelist

`func (o *EditPrimaryOptions) HasBranchWhitelist() bool`

HasBranchWhitelist returns a boolean if a field has been set.

### SetBranchWhitelistNil

`func (o *EditPrimaryOptions) SetBranchWhitelistNil(b bool)`

 SetBranchWhitelistNil sets the value for BranchWhitelist to be an explicit nil

### UnsetBranchWhitelist
`func (o *EditPrimaryOptions) UnsetBranchWhitelist()`

UnsetBranchWhitelist ensures that no value is present for BranchWhitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


