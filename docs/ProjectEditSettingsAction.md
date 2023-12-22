# ProjectEditSettingsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Labels** | Pointer to [**NullableEdit**](Edit.md) |  | [optional] 

## Methods

### NewProjectEditSettingsAction

`func NewProjectEditSettingsAction(name NullableString, ) *ProjectEditSettingsAction`

NewProjectEditSettingsAction instantiates a new ProjectEditSettingsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEditSettingsActionWithDefaults

`func NewProjectEditSettingsActionWithDefaults() *ProjectEditSettingsAction`

NewProjectEditSettingsActionWithDefaults instantiates a new ProjectEditSettingsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectEditSettingsAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectEditSettingsAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectEditSettingsAction) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProjectEditSettingsAction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectEditSettingsAction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabels

`func (o *ProjectEditSettingsAction) GetLabels() Edit`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ProjectEditSettingsAction) GetLabelsOk() (*Edit, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ProjectEditSettingsAction) SetLabels(v Edit)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ProjectEditSettingsAction) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ProjectEditSettingsAction) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ProjectEditSettingsAction) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


