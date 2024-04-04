# TemplateItemVariablesSchemaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator property | 
**DefaultValue** | Pointer to [**EnumTypeItemDefaultValue**](EnumTypeItemDefaultValue.md) |  | [optional] 
**Name** | Pointer to **string** | A variable used within the template. | [optional] 
**Description** | Pointer to **string** | The variable description | [optional] 
**Values** | Pointer to [**[]EnumTypeItemValuesInner**](EnumTypeItemValuesInner.md) | The available values for the enum. | [optional] 

## Methods

### NewTemplateItemVariablesSchemaInner

`func NewTemplateItemVariablesSchemaInner(type_ string, ) *TemplateItemVariablesSchemaInner`

NewTemplateItemVariablesSchemaInner instantiates a new TemplateItemVariablesSchemaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateItemVariablesSchemaInnerWithDefaults

`func NewTemplateItemVariablesSchemaInnerWithDefaults() *TemplateItemVariablesSchemaInner`

NewTemplateItemVariablesSchemaInnerWithDefaults instantiates a new TemplateItemVariablesSchemaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateItemVariablesSchemaInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateItemVariablesSchemaInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateItemVariablesSchemaInner) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultValue

`func (o *TemplateItemVariablesSchemaInner) GetDefaultValue() EnumTypeItemDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *TemplateItemVariablesSchemaInner) GetDefaultValueOk() (*EnumTypeItemDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *TemplateItemVariablesSchemaInner) SetDefaultValue(v EnumTypeItemDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *TemplateItemVariablesSchemaInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetName

`func (o *TemplateItemVariablesSchemaInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateItemVariablesSchemaInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateItemVariablesSchemaInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateItemVariablesSchemaInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateItemVariablesSchemaInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateItemVariablesSchemaInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateItemVariablesSchemaInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateItemVariablesSchemaInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValues

`func (o *TemplateItemVariablesSchemaInner) GetValues() []EnumTypeItemValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TemplateItemVariablesSchemaInner) GetValuesOk() (*[]EnumTypeItemValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TemplateItemVariablesSchemaInner) SetValues(v []EnumTypeItemValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *TemplateItemVariablesSchemaInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


