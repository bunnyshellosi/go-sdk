# ComponentEndpointCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Service component name | [optional] [readonly] 
**Endpoints** | Pointer to **[]string** | Service component URLs | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 

## Methods

### NewComponentEndpointCollection

`func NewComponentEndpointCollection() *ComponentEndpointCollection`

NewComponentEndpointCollection instantiates a new ComponentEndpointCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentEndpointCollectionWithDefaults

`func NewComponentEndpointCollectionWithDefaults() *ComponentEndpointCollection`

NewComponentEndpointCollectionWithDefaults instantiates a new ComponentEndpointCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentEndpointCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentEndpointCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentEndpointCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentEndpointCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentEndpointCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentEndpointCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentEndpointCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentEndpointCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpoints

`func (o *ComponentEndpointCollection) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ComponentEndpointCollection) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ComponentEndpointCollection) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ComponentEndpointCollection) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEnvironment

`func (o *ComponentEndpointCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ComponentEndpointCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ComponentEndpointCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ComponentEndpointCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


