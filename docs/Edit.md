# Edit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **map[string]string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] [default to "merge"]

## Methods

### NewEdit

`func NewEdit() *Edit`

NewEdit instantiates a new Edit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditWithDefaults

`func NewEditWithDefaults() *Edit`

NewEditWithDefaults instantiates a new Edit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *Edit) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Edit) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Edit) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Edit) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStrategy

`func (o *Edit) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Edit) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Edit) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Edit) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


