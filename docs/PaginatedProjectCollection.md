# PaginatedProjectCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Embedded** | Pointer to [**EmbeddedProjectCollection**](EmbeddedProjectCollection.md) |  | [optional] 

## Methods

### NewPaginatedProjectCollection

`func NewPaginatedProjectCollection() *PaginatedProjectCollection`

NewPaginatedProjectCollection instantiates a new PaginatedProjectCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProjectCollectionWithDefaults

`func NewPaginatedProjectCollectionWithDefaults() *PaginatedProjectCollection`

NewPaginatedProjectCollectionWithDefaults instantiates a new PaginatedProjectCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedProjectCollection) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedProjectCollection) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedProjectCollection) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedProjectCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalItems

`func (o *PaginatedProjectCollection) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginatedProjectCollection) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginatedProjectCollection) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *PaginatedProjectCollection) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *PaginatedProjectCollection) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *PaginatedProjectCollection) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *PaginatedProjectCollection) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *PaginatedProjectCollection) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEmbedded

`func (o *PaginatedProjectCollection) GetEmbedded() EmbeddedProjectCollection`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PaginatedProjectCollection) GetEmbeddedOk() (*EmbeddedProjectCollection, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PaginatedProjectCollection) SetEmbedded(v EmbeddedProjectCollection)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PaginatedProjectCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


