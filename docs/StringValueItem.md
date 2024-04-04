# StringValueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewStringValueItem

`func NewStringValueItem() *StringValueItem`

NewStringValueItem instantiates a new StringValueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringValueItemWithDefaults

`func NewStringValueItemWithDefaults() *StringValueItem`

NewStringValueItemWithDefaults instantiates a new StringValueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *StringValueItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringValueItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringValueItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringValueItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *StringValueItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringValueItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringValueItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StringValueItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


