# PaginatedLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**PaginatedLink**](PaginatedLink.md) |  | [optional] 
**First** | Pointer to [**NullablePaginatedLinksFirst**](PaginatedLinksFirst.md) |  | [optional] 
**Prev** | Pointer to [**NullablePaginatedLinksFirst**](PaginatedLinksFirst.md) |  | [optional] 
**Next** | Pointer to [**NullablePaginatedLinksFirst**](PaginatedLinksFirst.md) |  | [optional] 
**Last** | Pointer to [**NullablePaginatedLinksFirst**](PaginatedLinksFirst.md) |  | [optional] 
**Item** | Pointer to [**[]PaginatedLink**](PaginatedLink.md) |  | [optional] 

## Methods

### NewPaginatedLinks

`func NewPaginatedLinks() *PaginatedLinks`

NewPaginatedLinks instantiates a new PaginatedLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedLinksWithDefaults

`func NewPaginatedLinksWithDefaults() *PaginatedLinks`

NewPaginatedLinksWithDefaults instantiates a new PaginatedLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *PaginatedLinks) GetSelf() PaginatedLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginatedLinks) GetSelfOk() (*PaginatedLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginatedLinks) SetSelf(v PaginatedLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PaginatedLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetFirst

`func (o *PaginatedLinks) GetFirst() PaginatedLinksFirst`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginatedLinks) GetFirstOk() (*PaginatedLinksFirst, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginatedLinks) SetFirst(v PaginatedLinksFirst)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginatedLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *PaginatedLinks) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *PaginatedLinks) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetPrev

`func (o *PaginatedLinks) GetPrev() PaginatedLinksFirst`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginatedLinks) GetPrevOk() (*PaginatedLinksFirst, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginatedLinks) SetPrev(v PaginatedLinksFirst)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginatedLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### SetPrevNil

`func (o *PaginatedLinks) SetPrevNil(b bool)`

 SetPrevNil sets the value for Prev to be an explicit nil

### UnsetPrev
`func (o *PaginatedLinks) UnsetPrev()`

UnsetPrev ensures that no value is present for Prev, not even an explicit nil
### GetNext

`func (o *PaginatedLinks) GetNext() PaginatedLinksFirst`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedLinks) GetNextOk() (*PaginatedLinksFirst, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedLinks) SetNext(v PaginatedLinksFirst)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedLinks) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedLinks) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetLast

`func (o *PaginatedLinks) GetLast() PaginatedLinksFirst`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginatedLinks) GetLastOk() (*PaginatedLinksFirst, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginatedLinks) SetLast(v PaginatedLinksFirst)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginatedLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *PaginatedLinks) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *PaginatedLinks) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetItem

`func (o *PaginatedLinks) GetItem() []PaginatedLink`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PaginatedLinks) GetItemOk() (*[]PaginatedLink, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PaginatedLinks) SetItem(v []PaginatedLink)`

SetItem sets Item field to given value.

### HasItem

`func (o *PaginatedLinks) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


