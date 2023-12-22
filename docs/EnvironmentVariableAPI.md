# \EnvironmentVariableAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentVariableCreate**](EnvironmentVariableAPI.md#EnvironmentVariableCreate) | **Post** /v1/environment_variables | Create a specific environment variable.
[**EnvironmentVariableDelete**](EnvironmentVariableAPI.md#EnvironmentVariableDelete) | **Delete** /v1/environment_variables/{id} | Delete a specific environment variable.
[**EnvironmentVariableEdit**](EnvironmentVariableAPI.md#EnvironmentVariableEdit) | **Patch** /v1/environment_variables/{id} | Edit a specific environment variable.
[**EnvironmentVariableList**](EnvironmentVariableAPI.md#EnvironmentVariableList) | **Get** /v1/environment_variables | List environment variables matching any selected filters.
[**EnvironmentVariableView**](EnvironmentVariableAPI.md#EnvironmentVariableView) | **Get** /v1/environment_variables/{id} | View a specific environment variable.



## EnvironmentVariableCreate

> EnvironmentVariableItem EnvironmentVariableCreate(ctx).EnvironmentVariableCreateAction(environmentVariableCreateAction).Execute()

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
    environmentVariableCreateAction := *openapiclient.NewEnvironmentVariableCreateAction("Name_example", "Value_example", "Environment_example") // EnvironmentVariableCreateAction | The new environment_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableAPI.EnvironmentVariableCreate(context.Background()).EnvironmentVariableCreateAction(environmentVariableCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EnvironmentVariableCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableCreate`: EnvironmentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.EnvironmentVariableCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentVariableCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentVariableCreateAction** | [**EnvironmentVariableCreateAction**](EnvironmentVariableCreateAction.md) | The new environment_variable resource | 

### Return type

[**EnvironmentVariableItem**](EnvironmentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariableDelete

> EnvironmentVariableDelete(ctx, id).Body(body).Execute()

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
    r, err := apiClient.EnvironmentVariableAPI.EnvironmentVariableDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EnvironmentVariableDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentVariableDeleteRequest struct via the builder pattern


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


## EnvironmentVariableEdit

> EnvironmentVariableItem EnvironmentVariableEdit(ctx, id).EnvironmentVariableEditAction(environmentVariableEditAction).Execute()

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
    environmentVariableEditAction := *openapiclient.NewEnvironmentVariableEditAction("Value_example") // EnvironmentVariableEditAction | The updated environment_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableAPI.EnvironmentVariableEdit(context.Background(), id).EnvironmentVariableEditAction(environmentVariableEditAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EnvironmentVariableEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableEdit`: EnvironmentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.EnvironmentVariableEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentVariableEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableEditAction** | [**EnvironmentVariableEditAction**](EnvironmentVariableEditAction.md) | The updated environment_variable resource | 

### Return type

[**EnvironmentVariableItem**](EnvironmentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariableList

> PaginatedEnvironmentVariableCollection EnvironmentVariableList(ctx).Page(page).Name(name).Environment(environment).Organization(organization).Execute()

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
    environment := "environment_example" // string | Filter by environment (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableAPI.EnvironmentVariableList(context.Background()).Page(page).Name(name).Environment(environment).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EnvironmentVariableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableList`: PaginatedEnvironmentVariableCollection
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.EnvironmentVariableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentVariableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **name** | **string** | Filter by name | 
 **environment** | **string** | Filter by environment | 
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedEnvironmentVariableCollection**](PaginatedEnvironmentVariableCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariableView

> EnvironmentVariableItem EnvironmentVariableView(ctx, id).Execute()

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
    resp, r, err := apiClient.EnvironmentVariableAPI.EnvironmentVariableView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EnvironmentVariableView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableView`: EnvironmentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.EnvironmentVariableView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentVariableViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableItem**](EnvironmentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

