# PaginatedEnvironmentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Embedded** | Pointer to [**EmbeddedEnvironmentCollection**](EmbeddedEnvironmentCollection.md) |  | [optional] 

## Methods

### NewPaginatedEnvironmentCollection

`func NewPaginatedEnvironmentCollection() *PaginatedEnvironmentCollection`

NewPaginatedEnvironmentCollection instantiates a new PaginatedEnvironmentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEnvironmentCollectionWithDefaults

`func NewPaginatedEnvironmentCollectionWithDefaults() *PaginatedEnvironmentCollection`

NewPaginatedEnvironmentCollectionWithDefaults instantiates a new PaginatedEnvironmentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedEnvironmentCollection) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedEnvironmentCollection) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedEnvironmentCollection) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedEnvironmentCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalItems

`func (o *PaginatedEnvironmentCollection) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginatedEnvironmentCollection) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginatedEnvironmentCollection) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *PaginatedEnvironmentCollection) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *PaginatedEnvironmentCollection) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *PaginatedEnvironmentCollection) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *PaginatedEnvironmentCollection) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *PaginatedEnvironmentCollection) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEmbedded

`func (o *PaginatedEnvironmentCollection) GetEmbedded() EmbeddedEnvironmentCollection`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PaginatedEnvironmentCollection) GetEmbeddedOk() (*EmbeddedEnvironmentCollection, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PaginatedEnvironmentCollection) SetEmbedded(v EmbeddedEnvironmentCollection)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PaginatedEnvironmentCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


