# BooleanValueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBooleanValueItem

`func NewBooleanValueItem() *BooleanValueItem`

NewBooleanValueItem instantiates a new BooleanValueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanValueItemWithDefaults

`func NewBooleanValueItemWithDefaults() *BooleanValueItem`

NewBooleanValueItemWithDefaults instantiates a new BooleanValueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BooleanValueItem) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BooleanValueItem) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BooleanValueItem) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *BooleanValueItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *BooleanValueItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooleanValueItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooleanValueItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BooleanValueItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


