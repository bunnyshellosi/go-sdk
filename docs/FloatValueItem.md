# FloatValueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewFloatValueItem

`func NewFloatValueItem() *FloatValueItem`

NewFloatValueItem instantiates a new FloatValueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatValueItemWithDefaults

`func NewFloatValueItemWithDefaults() *FloatValueItem`

NewFloatValueItemWithDefaults instantiates a new FloatValueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FloatValueItem) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FloatValueItem) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FloatValueItem) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *FloatValueItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *FloatValueItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatValueItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatValueItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FloatValueItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


