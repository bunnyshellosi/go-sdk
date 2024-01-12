# SecretEncryptedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | Encrypted expression. | [optional] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] 

## Methods

### NewSecretEncryptedItem

`func NewSecretEncryptedItem() *SecretEncryptedItem`

NewSecretEncryptedItem instantiates a new SecretEncryptedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretEncryptedItemWithDefaults

`func NewSecretEncryptedItemWithDefaults() *SecretEncryptedItem`

NewSecretEncryptedItemWithDefaults instantiates a new SecretEncryptedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *SecretEncryptedItem) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SecretEncryptedItem) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SecretEncryptedItem) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SecretEncryptedItem) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOrganization

`func (o *SecretEncryptedItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SecretEncryptedItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SecretEncryptedItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SecretEncryptedItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


