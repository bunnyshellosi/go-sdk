# \WorkflowAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowList**](WorkflowAPI.md#WorkflowList) | **Get** /v1/workflows | List workflows matching any selected filters.
[**WorkflowView**](WorkflowAPI.md#WorkflowView) | **Get** /v1/workflows/{id} | View a specific Workflow.



## WorkflowList

> PaginatedWorkflowCollection WorkflowList(ctx).Page(page).Event(event).Environment(environment).Organization(organization).Status(status).Execute()

List workflows matching any selected filters.



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
    event := "event_example" // string | Filter by event (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    organization := "organization_example" // string | Filter by organization (optional)
    status := "queued" // string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowAPI.WorkflowList(context.Background()).Page(page).Event(event).Environment(environment).Organization(organization).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAPI.WorkflowList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowList`: PaginatedWorkflowCollection
    fmt.Fprintf(os.Stdout, "Response from `WorkflowAPI.WorkflowList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **event** | **string** | Filter by event | 
 **environment** | **string** | Filter by environment | 
 **organization** | **string** | Filter by organization | 
 **status** | **string** | Filter by status | 

### Return type

[**PaginatedWorkflowCollection**](PaginatedWorkflowCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowView

> WorkflowItem WorkflowView(ctx, id).Execute()

View a specific Workflow.



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
    resp, r, err := apiClient.WorkflowAPI.WorkflowView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAPI.WorkflowView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowView`: WorkflowItem
    fmt.Fprintf(os.Stdout, "Response from `WorkflowAPI.WorkflowView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowItem**](WorkflowItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

