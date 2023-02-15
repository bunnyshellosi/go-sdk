# \ComponentEndpointApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentEndpointList**](ComponentEndpointApi.md#ComponentEndpointList) | **Get** /v1/components/endpoint | List endpoints for service components matching any selected filters
[**ComponentEndpointView**](ComponentEndpointApi.md#ComponentEndpointView) | **Get** /v1/components/{id}/endpoint | View endpoints for a specific service component



## ComponentEndpointList

> PaginatedComponentEndpointCollection ComponentEndpointList(ctx).Page(page).Organization(organization).Project(project).Environment(environment).Name(name).Execute()

List endpoints for service components matching any selected filters



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
    organization := "organization_example" // string | Filter by organization (optional)
    project := "project_example" // string | Filter by project (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    name := "name_example" // string | Filter by name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentEndpointApi.ComponentEndpointList(context.Background()).Page(page).Organization(organization).Project(project).Environment(environment).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentEndpointApi.ComponentEndpointList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentEndpointList`: PaginatedComponentEndpointCollection
    fmt.Fprintf(os.Stdout, "Response from `ComponentEndpointApi.ComponentEndpointList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentEndpointListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **project** | **string** | Filter by project | 
 **environment** | **string** | Filter by environment | 
 **name** | **string** | Filter by name | 

### Return type

[**PaginatedComponentEndpointCollection**](PaginatedComponentEndpointCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentEndpointView

> ComponentEndpointItem ComponentEndpointView(ctx, id).Execute()

View endpoints for a specific service component



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
    resp, r, err := apiClient.ComponentEndpointApi.ComponentEndpointView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentEndpointApi.ComponentEndpointView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentEndpointView`: ComponentEndpointItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentEndpointApi.ComponentEndpointView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentEndpointViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentEndpointItem**](ComponentEndpointItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

