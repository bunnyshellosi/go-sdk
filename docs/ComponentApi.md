# \ComponentApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentList**](ComponentApi.md#ComponentList) | **Get** /v1/components | Retrieves the collection of component resources.
[**ComponentRemoteDevelopmentDown**](ComponentApi.md#ComponentRemoteDevelopmentDown) | **Post** /v1/components/{id}/remote-development/down | Creates a component resource.
[**ComponentRemoteDevelopmentUp**](ComponentApi.md#ComponentRemoteDevelopmentUp) | **Post** /v1/components/{id}/remote-development/up | Creates a component resource.
[**ComponentView**](ComponentApi.md#ComponentView) | **Get** /v1/components/{id} | Retrieves a component resource.



## ComponentList

> PaginatedComponentCollection ComponentList(ctx).Page(page).Environment(environment).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Organization(organization).Project(project).Execute()

Retrieves the collection of component resources.



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
    environment := "environment_example" // string | Filter by environment (optional)
    operationStatus := "ready" // string | Filter by operationStatus (optional)
    clusterStatus := "not_available" // string | Filter by clusterStatus (optional)
    organization := "organization_example" // string | Filter by organization (optional)
    project := "project_example" // string | Filter by project (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.ComponentList(context.Background()).Page(page).Environment(environment).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Organization(organization).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentList`: PaginatedComponentCollection
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **environment** | **string** | Filter by environment | 
 **operationStatus** | **string** | Filter by operationStatus | 
 **clusterStatus** | **string** | Filter by clusterStatus | 
 **organization** | **string** | Filter by organization | 
 **project** | **string** | Filter by project | 

### Return type

[**PaginatedComponentCollection**](PaginatedComponentCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentRemoteDevelopmentDown

> ComponentItem ComponentRemoteDevelopmentDown(ctx, id).Body(body).Execute()

Creates a component resource.



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
    body := interface{}(987) // interface{} | No Request Body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.ComponentRemoteDevelopmentDown(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentRemoteDevelopmentDown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentRemoteDevelopmentDown`: ComponentItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentRemoteDevelopmentDown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentRemoteDevelopmentDownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | No Request Body | 

### Return type

[**ComponentItem**](ComponentItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentRemoteDevelopmentUp

> ComponentItem ComponentRemoteDevelopmentUp(ctx, id).ComponentRemoteDevelopmentUp(componentRemoteDevelopmentUp).Execute()

Creates a component resource.



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
    componentRemoteDevelopmentUp := *openapiclient.NewComponentRemoteDevelopmentUp() // ComponentRemoteDevelopmentUp | The new component resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentApi.ComponentRemoteDevelopmentUp(context.Background(), id).ComponentRemoteDevelopmentUp(componentRemoteDevelopmentUp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentRemoteDevelopmentUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentRemoteDevelopmentUp`: ComponentItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentRemoteDevelopmentUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentRemoteDevelopmentUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentRemoteDevelopmentUp** | [**ComponentRemoteDevelopmentUp**](ComponentRemoteDevelopmentUp.md) | The new component resource | 

### Return type

[**ComponentItem**](ComponentItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentView

> ComponentItem ComponentView(ctx, id).Execute()

Retrieves a component resource.



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
    resp, r, err := apiClient.ComponentApi.ComponentView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentView`: ComponentItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentItem**](ComponentItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

