# \ComponentApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentList**](ComponentApi.md#ComponentList) | **Get** /v1/components | List service components matching any selected filters
[**ComponentResources**](ComponentApi.md#ComponentResources) | **Get** /v1/components/{id}/resources | Get kubernetes resources
[**ComponentView**](ComponentApi.md#ComponentView) | **Get** /v1/components/{id} | View a specific service component



## ComponentList

> PaginatedComponentCollection ComponentList(ctx).Page(page).Environment(environment).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Organization(organization).Project(project).Execute()

List service components matching any selected filters



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
    operationStatus := "draft" // string | Filter by operationStatus (optional)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentResources

> []ComponentResourceItem ComponentResources(ctx, id).Execute()

Get kubernetes resources



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
    resp, r, err := apiClient.ComponentApi.ComponentResources(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentApi.ComponentResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentResources`: []ComponentResourceItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentApi.ComponentResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ComponentResourceItem**](ComponentResourceItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentView

> ComponentItem ComponentView(ctx, id).Execute()

View a specific service component



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

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

