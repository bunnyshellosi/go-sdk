# EnvironmentEditComponentsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**EnvironmentEditComponentsActionFilter**](EnvironmentEditComponentsActionFilter.md) |  | 
**Target** | [**GitInfo**](GitInfo.md) |  | 

## Methods

### NewEnvironmentEditComponentsAction

`func NewEnvironmentEditComponentsAction(filter EnvironmentEditComponentsActionFilter, target GitInfo, ) *EnvironmentEditComponentsAction`

NewEnvironmentEditComponentsAction instantiates a new EnvironmentEditComponentsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditComponentsActionWithDefaults

`func NewEnvironmentEditComponentsActionWithDefaults() *EnvironmentEditComponentsAction`

NewEnvironmentEditComponentsActionWithDefaults instantiates a new EnvironmentEditComponentsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *EnvironmentEditComponentsAction) GetFilter() EnvironmentEditComponentsActionFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EnvironmentEditComponentsAction) GetFilterOk() (*EnvironmentEditComponentsActionFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EnvironmentEditComponentsAction) SetFilter(v EnvironmentEditComponentsActionFilter)`

SetFilter sets Filter field to given value.


### GetTarget

`func (o *EnvironmentEditComponentsAction) GetTarget() GitInfo`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EnvironmentEditComponentsAction) GetTargetOk() (*GitInfo, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EnvironmentEditComponentsAction) SetTarget(v GitInfo)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


