# \PipelineApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PipelineList**](PipelineApi.md#PipelineList) | **Get** /v1/pipelines | List pipelines matching any selected filters.
[**PipelineView**](PipelineApi.md#PipelineView) | **Get** /v1/pipelines/{id} | View a specific Pipeline.



## PipelineList

> PaginatedPipelineCollection PipelineList(ctx).Page(page).Environment(environment).Event(event).Organization(organization).Status(status).Execute()

List pipelines matching any selected filters.



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
    event := "event_example" // string | Filter by event (optional)
    organization := "organization_example" // string | Filter by organization (optional)
    status := "success" // string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.PipelineList(context.Background()).Page(page).Environment(environment).Event(event).Organization(organization).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.PipelineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PipelineList`: PaginatedPipelineCollection
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.PipelineList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPipelineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **environment** | **string** | Filter by environment | 
 **event** | **string** | Filter by event | 
 **organization** | **string** | Filter by organization | 
 **status** | **string** | Filter by status | 

### Return type

[**PaginatedPipelineCollection**](PaginatedPipelineCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineView

> PipelineItem PipelineView(ctx, id).Execute()

View a specific Pipeline.



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
    resp, r, err := apiClient.PipelineApi.PipelineView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.PipelineView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PipelineView`: PipelineItem
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.PipelineView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PipelineItem**](PipelineItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

