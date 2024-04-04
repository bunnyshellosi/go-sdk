# IntegerValueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewIntegerValueItem

`func NewIntegerValueItem() *IntegerValueItem`

NewIntegerValueItem instantiates a new IntegerValueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegerValueItemWithDefaults

`func NewIntegerValueItemWithDefaults() *IntegerValueItem`

NewIntegerValueItemWithDefaults instantiates a new IntegerValueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IntegerValueItem) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IntegerValueItem) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IntegerValueItem) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *IntegerValueItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *IntegerValueItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegerValueItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegerValueItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegerValueItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


