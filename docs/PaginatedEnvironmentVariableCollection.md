# PaginatedEnvironmentVariableCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**TotalItems** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Embedded** | Pointer to [**EmbeddedEnvironmentVariableCollection**](EmbeddedEnvironmentVariableCollection.md) |  | [optional] 

## Methods

### NewPaginatedEnvironmentVariableCollection

`func NewPaginatedEnvironmentVariableCollection() *PaginatedEnvironmentVariableCollection`

NewPaginatedEnvironmentVariableCollection instantiates a new PaginatedEnvironmentVariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEnvironmentVariableCollectionWithDefaults

`func NewPaginatedEnvironmentVariableCollectionWithDefaults() *PaginatedEnvironmentVariableCollection`

NewPaginatedEnvironmentVariableCollectionWithDefaults instantiates a new PaginatedEnvironmentVariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedEnvironmentVariableCollection) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedEnvironmentVariableCollection) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedEnvironmentVariableCollection) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedEnvironmentVariableCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalItems

`func (o *PaginatedEnvironmentVariableCollection) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginatedEnvironmentVariableCollection) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginatedEnvironmentVariableCollection) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *PaginatedEnvironmentVariableCollection) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetPage

`func (o *PaginatedEnvironmentVariableCollection) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedEnvironmentVariableCollection) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedEnvironmentVariableCollection) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedEnvironmentVariableCollection) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *PaginatedEnvironmentVariableCollection) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *PaginatedEnvironmentVariableCollection) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *PaginatedEnvironmentVariableCollection) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *PaginatedEnvironmentVariableCollection) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEmbedded

`func (o *PaginatedEnvironmentVariableCollection) GetEmbedded() EmbeddedEnvironmentVariableCollection`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PaginatedEnvironmentVariableCollection) GetEmbeddedOk() (*EmbeddedEnvironmentVariableCollection, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PaginatedEnvironmentVariableCollection) SetEmbedded(v EmbeddedEnvironmentVariableCollection)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PaginatedEnvironmentVariableCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


