# OrganizationCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Organization identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Organization name as defined in Bunnyshell UI. | [optional] [readonly] 
**Timezone** | Pointer to **string** | Organization timezone as defined in Bunnyshell UI. | [optional] [readonly] 

## Methods

### NewOrganizationCollection

`func NewOrganizationCollection() *OrganizationCollection`

NewOrganizationCollection instantiates a new OrganizationCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCollectionWithDefaults

`func NewOrganizationCollectionWithDefaults() *OrganizationCollection`

NewOrganizationCollectionWithDefaults instantiates a new OrganizationCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimezone

`func (o *OrganizationCollection) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OrganizationCollection) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OrganizationCollection) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *OrganizationCollection) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


