# EnumTypeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to [**EnumTypeItemDefaultValue**](EnumTypeItemDefaultValue.md) |  | [optional] 
**Values** | Pointer to [**[]EnumTypeItemValuesInner**](EnumTypeItemValuesInner.md) | The available values for the enum. | [optional] 
**Name** | Pointer to **string** | A variable used within the template. | [optional] 
**Description** | Pointer to **string** | The variable description | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEnumTypeItem

`func NewEnumTypeItem() *EnumTypeItem`

NewEnumTypeItem instantiates a new EnumTypeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumTypeItemWithDefaults

`func NewEnumTypeItemWithDefaults() *EnumTypeItem`

NewEnumTypeItemWithDefaults instantiates a new EnumTypeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *EnumTypeItem) GetDefaultValue() EnumTypeItemDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EnumTypeItem) GetDefaultValueOk() (*EnumTypeItemDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EnumTypeItem) SetDefaultValue(v EnumTypeItemDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EnumTypeItem) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetValues

`func (o *EnumTypeItem) GetValues() []EnumTypeItemValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EnumTypeItem) GetValuesOk() (*[]EnumTypeItemValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EnumTypeItem) SetValues(v []EnumTypeItemValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *EnumTypeItem) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetName

`func (o *EnumTypeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumTypeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumTypeItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumTypeItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EnumTypeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumTypeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumTypeItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnumTypeItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EnumTypeItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnumTypeItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnumTypeItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnumTypeItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


