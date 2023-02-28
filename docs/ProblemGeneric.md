# ProblemGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Violations** | Pointer to [**[]ProblemViolation**](ProblemViolation.md) |  | [optional] 

## Methods

### NewProblemGeneric

`func NewProblemGeneric() *ProblemGeneric`

NewProblemGeneric instantiates a new ProblemGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemGenericWithDefaults

`func NewProblemGenericWithDefaults() *ProblemGeneric`

NewProblemGenericWithDefaults instantiates a new ProblemGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ProblemGeneric) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemGeneric) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemGeneric) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProblemGeneric) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ProblemGeneric) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemGeneric) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemGeneric) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ProblemGeneric) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetViolations

`func (o *ProblemGeneric) GetViolations() []ProblemViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ProblemGeneric) GetViolationsOk() (*[]ProblemViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ProblemGeneric) SetViolations(v []ProblemViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *ProblemGeneric) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### SetViolationsNil

`func (o *ProblemGeneric) SetViolationsNil(b bool)`

 SetViolationsNil sets the value for Violations to be an explicit nil

### UnsetViolations
`func (o *ProblemGeneric) UnsetViolations()`

UnsetViolations ensures that no value is present for Violations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


