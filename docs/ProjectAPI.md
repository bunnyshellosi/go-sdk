# \ProjectAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectList**](ProjectAPI.md#ProjectList) | **Get** /v1/projects | List projects matching any selected filters.
[**ProjectView**](ProjectAPI.md#ProjectView) | **Get** /v1/projects/{id} | View a specific project.



## ProjectList

> PaginatedProjectCollection ProjectList(ctx).Page(page).Organization(organization).Search(search).Labels(labels).Execute()

List projects matching any selected filters.



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
    organization := "organization_example" // string | Filter by organization (optional)
    search := "search_example" // string | Filter by search (optional)
    labels := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Filter by label key-value pair. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.ProjectList(context.Background()).Page(page).Organization(organization).Search(search).Labels(labels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectList`: PaginatedProjectCollection
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **search** | **string** | Filter by search | 
 **labels** | **map[string]map[string]string** | Filter by label key-value pair. | 

### Return type

[**PaginatedProjectCollection**](PaginatedProjectCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectView

> ProjectItem ProjectView(ctx, id).Execute()

View a specific project.



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
    resp, r, err := apiClient.ProjectAPI.ProjectView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectView`: ProjectItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectItem**](ProjectItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

