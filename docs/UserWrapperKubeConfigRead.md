# UserWrapperKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**User** | Pointer to [**UserKubeConfigRead**](UserKubeConfigRead.md) |  | [optional] 

## Methods

### NewUserWrapperKubeConfigRead

`func NewUserWrapperKubeConfigRead(name string, ) *UserWrapperKubeConfigRead`

NewUserWrapperKubeConfigRead instantiates a new UserWrapperKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWrapperKubeConfigReadWithDefaults

`func NewUserWrapperKubeConfigReadWithDefaults() *UserWrapperKubeConfigRead`

NewUserWrapperKubeConfigReadWithDefaults instantiates a new UserWrapperKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserWrapperKubeConfigRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWrapperKubeConfigRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWrapperKubeConfigRead) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *UserWrapperKubeConfigRead) GetUser() UserKubeConfigRead`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserWrapperKubeConfigRead) GetUserOk() (*UserKubeConfigRead, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserWrapperKubeConfigRead) SetUser(v UserKubeConfigRead)`

SetUser sets User field to given value.

### HasUser

`func (o *UserWrapperKubeConfigRead) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


