# \EnvironmentVariableApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentVariableEdit**](EnvironmentVariableApi.md#EnvironmentVariableEdit) | **Patch** /v1/environment_variables/{id} | Updates the environment_variable resource.
[**EnvironmentVariableList**](EnvironmentVariableApi.md#EnvironmentVariableList) | **Get** /v1/environment_variables | Retrieves the collection of environment_variable resources.
[**EnvironmentVariableView**](EnvironmentVariableApi.md#EnvironmentVariableView) | **Get** /v1/environment_variables/{id} | Retrieves a environment_variable resource.



## EnvironmentVariableEdit

> EnvironmentVariableItem EnvironmentVariableEdit(ctx, id).EnvironmentVariableEdit(environmentVariableEdit).Execute()

Updates the environment_variable resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Resource identifier
    environmentVariableEdit := *openapiclient.NewEnvironmentVariableEdit() // EnvironmentVariableEdit | The updated environment_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableApi.EnvironmentVariableEdit(context.Background(), id).EnvironmentVariableEdit(environmentVariableEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.EnvironmentVariableEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableEdit`: EnvironmentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.EnvironmentVariableEdit`: %v\n", resp)
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

 **environmentVariableEdit** | [**EnvironmentVariableEdit**](EnvironmentVariableEdit.md) | The updated environment_variable resource | 

### Return type

[**EnvironmentVariableItem**](EnvironmentVariableItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariableList

> PaginatedEnvironmentVariableCollection EnvironmentVariableList(ctx).Page(page).Name(name).Environment(environment).Organization(organization).Execute()

Retrieves the collection of environment_variable resources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    name := "name_example" // string | Filter by name (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableApi.EnvironmentVariableList(context.Background()).Page(page).Name(name).Environment(environment).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.EnvironmentVariableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableList`: PaginatedEnvironmentVariableCollection
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.EnvironmentVariableList`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariableView

> EnvironmentVariableItem EnvironmentVariableView(ctx, id).Execute()

Retrieves a environment_variable resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Resource identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentVariableApi.EnvironmentVariableView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.EnvironmentVariableView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentVariableView`: EnvironmentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.EnvironmentVariableView`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

