# ResourceRequirementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**NullableResourceRequirementItemRequests**](ResourceRequirementItemRequests.md) |  | [optional] 
**Limits** | Pointer to [**NullableResourceRequirementItemLimits**](ResourceRequirementItemLimits.md) |  | [optional] 

## Methods

### NewResourceRequirementItem

`func NewResourceRequirementItem() *ResourceRequirementItem`

NewResourceRequirementItem instantiates a new ResourceRequirementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementItemWithDefaults

`func NewResourceRequirementItemWithDefaults() *ResourceRequirementItem`

NewResourceRequirementItemWithDefaults instantiates a new ResourceRequirementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *ResourceRequirementItem) GetRequests() ResourceRequirementItemRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ResourceRequirementItem) GetRequestsOk() (*ResourceRequirementItemRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ResourceRequirementItem) SetRequests(v ResourceRequirementItemRequests)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ResourceRequirementItem) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *ResourceRequirementItem) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *ResourceRequirementItem) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil
### GetLimits

`func (o *ResourceRequirementItem) GetLimits() ResourceRequirementItemLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ResourceRequirementItem) GetLimitsOk() (*ResourceRequirementItemLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ResourceRequirementItem) SetLimits(v ResourceRequirementItemLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ResourceRequirementItem) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *ResourceRequirementItem) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *ResourceRequirementItem) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


