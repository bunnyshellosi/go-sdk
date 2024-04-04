# FloatTypeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to [**NullableFloatValueItem**](FloatValueItem.md) |  | [optional] 
**Name** | Pointer to **string** | A variable used within the template. | [optional] 
**Description** | Pointer to **string** | The variable description | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewFloatTypeItem

`func NewFloatTypeItem() *FloatTypeItem`

NewFloatTypeItem instantiates a new FloatTypeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatTypeItemWithDefaults

`func NewFloatTypeItemWithDefaults() *FloatTypeItem`

NewFloatTypeItemWithDefaults instantiates a new FloatTypeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *FloatTypeItem) GetDefaultValue() FloatValueItem`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FloatTypeItem) GetDefaultValueOk() (*FloatValueItem, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FloatTypeItem) SetDefaultValue(v FloatValueItem)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FloatTypeItem) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *FloatTypeItem) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *FloatTypeItem) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetName

`func (o *FloatTypeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloatTypeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloatTypeItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FloatTypeItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FloatTypeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FloatTypeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FloatTypeItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FloatTypeItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *FloatTypeItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatTypeItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatTypeItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FloatTypeItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


