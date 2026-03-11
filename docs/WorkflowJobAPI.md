# \WorkflowJobAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowJobList**](WorkflowJobAPI.md#WorkflowJobList) | **Get** /v1/workflow_jobs | List workflow jobs matching any selected filters.
[**WorkflowJobLogs**](WorkflowJobAPI.md#WorkflowJobLogs) | **Get** /v1/workflow_jobs/{id}/logs | View logs for a specific workflow job.
[**WorkflowJobView**](WorkflowJobAPI.md#WorkflowJobView) | **Get** /v1/workflow_jobs/{id} | View a specific workflow job.



## WorkflowJobList

> PaginatedWorkflowJobCollection WorkflowJobList(ctx).Workflow(workflow).Page(page).Status(status).Execute()

List workflow jobs matching any selected filters.



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
    workflow := "workflow_example" // string | Filter workflow jobs by workflow identifier.
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    status := []string{"Status_example"} // []string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowJobAPI.WorkflowJobList(context.Background()).Workflow(workflow).Page(page).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobAPI.WorkflowJobList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowJobList`: PaginatedWorkflowJobCollection
    fmt.Fprintf(os.Stdout, "Response from `WorkflowJobAPI.WorkflowJobList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflow** | **string** | Filter workflow jobs by workflow identifier. | 
 **page** | **int32** | The collection page number | [default to 1]
 **status** | **[]string** | Filter by status | 

### Return type

[**PaginatedWorkflowJobCollection**](PaginatedWorkflowJobCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowJobLogs

> WorkflowJobWorkflowJobLogsOutputItem WorkflowJobLogs(ctx, id).StepStatus(stepStatus).Execute()

View logs for a specific workflow job.



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
    id := "id_example" // string | Workflow job identifier.
    stepStatus := []string{"StepStatus_example"} // []string | Filter workflow steps by status. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowJobAPI.WorkflowJobLogs(context.Background(), id).StepStatus(stepStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobAPI.WorkflowJobLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowJobLogs`: WorkflowJobWorkflowJobLogsOutputItem
    fmt.Fprintf(os.Stdout, "Response from `WorkflowJobAPI.WorkflowJobLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowJobLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stepStatus** | **[]string** | Filter workflow steps by status. | 

### Return type

[**WorkflowJobWorkflowJobLogsOutputItem**](WorkflowJobWorkflowJobLogsOutputItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowJobView

> WorkflowJobItem WorkflowJobView(ctx, id).Execute()

View a specific workflow job.



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
    id := "id_example" // string | Workflow job identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowJobAPI.WorkflowJobView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowJobAPI.WorkflowJobView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowJobView`: WorkflowJobItem
    fmt.Fprintf(os.Stdout, "Response from `WorkflowJobAPI.WorkflowJobView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowJobViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowJobItem**](WorkflowJobItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

