# EnvironmentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | Environment type: primary or ephemeral. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment name. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Environment k8s namespace. | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Environment operation status. | [optional] [readonly] 
**Project** | Pointer to **string** | Project identifier. | [optional] [readonly] 

## Methods

### NewEnvironmentCollection

`func NewEnvironmentCollection() *EnvironmentCollection`

NewEnvironmentCollection instantiates a new EnvironmentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCollectionWithDefaults

`func NewEnvironmentCollectionWithDefaults() *EnvironmentCollection`

NewEnvironmentCollectionWithDefaults instantiates a new EnvironmentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentCollection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentCollection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EnvironmentCollection) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EnvironmentCollection) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EnvironmentCollection) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EnvironmentCollection) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOperationStatus

`func (o *EnvironmentCollection) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *EnvironmentCollection) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *EnvironmentCollection) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *EnvironmentCollection) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetProject

`func (o *EnvironmentCollection) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentCollection) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentCollection) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentCollection) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


