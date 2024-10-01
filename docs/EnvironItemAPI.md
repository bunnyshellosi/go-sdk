# \EnvironItemAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironItemCreate**](EnvironItemAPI.md#EnvironItemCreate) | **Post** /v1/environ_items | Create a specific environment variable.
[**EnvironItemDelete**](EnvironItemAPI.md#EnvironItemDelete) | **Delete** /v1/environ_items/{id} | Delete a specific environment variable.
[**EnvironItemEdit**](EnvironItemAPI.md#EnvironItemEdit) | **Patch** /v1/environ_items/{id} | Edit a specific environment variable.
[**EnvironItemList**](EnvironItemAPI.md#EnvironItemList) | **Get** /v1/environ_items | List environment variables matching any selected filters.
[**EnvironItemView**](EnvironItemAPI.md#EnvironItemView) | **Get** /v1/environ_items/{id} | View a specific environment variable.



## EnvironItemCreate

> EnvironItemItem EnvironItemCreate(ctx).EnvironItemCreateAction(environItemCreateAction).Execute()

Create a specific environment variable.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    environItemCreateAction := *openapiclient.NewEnvironItemCreateAction("GroupName_example", "Name_example", "Value_example", "Environment_example") // EnvironItemCreateAction | The new environ_item resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironItemAPI.EnvironItemCreate(context.Background()).EnvironItemCreateAction(environItemCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironItemAPI.EnvironItemCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironItemCreate`: EnvironItemItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironItemAPI.EnvironItemCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironItemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environItemCreateAction** | [**EnvironItemCreateAction**](EnvironItemCreateAction.md) | The new environ_item resource | 

### Return type

[**EnvironItemItem**](EnvironItemItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironItemDelete

> EnvironItemDelete(ctx, id).Body(body).Execute()

Delete a specific environment variable.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    id := "id_example" // string | Resource identifier
    body := interface{}(987) // interface{} | No Request Body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironItemAPI.EnvironItemDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironItemAPI.EnvironItemDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironItemDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | No Request Body | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironItemEdit

> EnvironItemItem EnvironItemEdit(ctx, id).EnvironItemEditAction(environItemEditAction).Execute()

Edit a specific environment variable.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    id := "id_example" // string | Resource identifier
    environItemEditAction := *openapiclient.NewEnvironItemEditAction() // EnvironItemEditAction | The updated environ_item resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironItemAPI.EnvironItemEdit(context.Background(), id).EnvironItemEditAction(environItemEditAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironItemAPI.EnvironItemEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironItemEdit`: EnvironItemItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironItemAPI.EnvironItemEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironItemEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environItemEditAction** | [**EnvironItemEditAction**](EnvironItemEditAction.md) | The updated environ_item resource | 

### Return type

[**EnvironItemItem**](EnvironItemItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironItemList

> PaginatedEnvironItemCollection EnvironItemList(ctx).Page(page).Name(name).GroupName(groupName).Environment(environment).Organization(organization).Execute()

List environment variables matching any selected filters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    name := "name_example" // string | Filter by name (optional)
    groupName := "groupName_example" // string | Filter by groupName (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironItemAPI.EnvironItemList(context.Background()).Page(page).Name(name).GroupName(groupName).Environment(environment).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironItemAPI.EnvironItemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironItemList`: PaginatedEnvironItemCollection
    fmt.Fprintf(os.Stdout, "Response from `EnvironItemAPI.EnvironItemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **name** | **string** | Filter by name | 
 **groupName** | **string** | Filter by groupName | 
 **environment** | **string** | Filter by environment | 
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedEnvironItemCollection**](PaginatedEnvironItemCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironItemView

> EnvironItemItem EnvironItemView(ctx, id).Execute()

View a specific environment variable.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    id := "id_example" // string | Resource identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironItemAPI.EnvironItemView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironItemAPI.EnvironItemView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironItemView`: EnvironItemItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironItemAPI.EnvironItemView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironItemViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironItemItem**](EnvironItemItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

