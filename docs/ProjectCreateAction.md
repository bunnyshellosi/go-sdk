# ProjectCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Organization** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectCreateAction

`func NewProjectCreateAction(name string, organization string, ) *ProjectCreateAction`

NewProjectCreateAction instantiates a new ProjectCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateActionWithDefaults

`func NewProjectCreateActionWithDefaults() *ProjectCreateAction`

NewProjectCreateActionWithDefaults instantiates a new ProjectCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *ProjectCreateAction) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectCreateAction) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectCreateAction) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetLabels

`func (o *ProjectCreateAction) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ProjectCreateAction) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ProjectCreateAction) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ProjectCreateAction) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


