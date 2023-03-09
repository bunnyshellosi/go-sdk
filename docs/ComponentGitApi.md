# \ComponentGitApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentGitList**](ComponentGitApi.md#ComponentGitList) | **Get** /v1/components/gitinfo | List git info for service components matching any selected filters
[**ComponentGitView**](ComponentGitApi.md#ComponentGitView) | **Get** /v1/components/{id}/gitinfo | View git info for a specific service component



## ComponentGitList

> PaginatedComponentGitCollection ComponentGitList(ctx).Page(page).Organization(organization).Project(project).Environment(environment).Name(name).Repository(repository).Branch(branch).Execute()

List git info for service components matching any selected filters



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
    repository := "repository_example" // string | Filter by repository (optional)
    branch := "branch_example" // string | Filter by branch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentGitApi.ComponentGitList(context.Background()).Page(page).Organization(organization).Project(project).Environment(environment).Name(name).Repository(repository).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGitApi.ComponentGitList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentGitList`: PaginatedComponentGitCollection
    fmt.Fprintf(os.Stdout, "Response from `ComponentGitApi.ComponentGitList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentGitListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **project** | **string** | Filter by project | 
 **environment** | **string** | Filter by environment | 
 **name** | **string** | Filter by name | 
 **repository** | **string** | Filter by repository | 
 **branch** | **string** | Filter by branch | 

### Return type

[**PaginatedComponentGitCollection**](PaginatedComponentGitCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentGitView

> ComponentGitItem ComponentGitView(ctx, id).Execute()

View git info for a specific service component



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
    resp, r, err := apiClient.ComponentGitApi.ComponentGitView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGitApi.ComponentGitView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentGitView`: ComponentGitItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentGitApi.ComponentGitView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentGitViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentGitItem**](ComponentGitItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

