# \TemplatesRepositoryAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesRepositoryList**](TemplatesRepositoryAPI.md#TemplatesRepositoryList) | **Get** /v1/templates_repositories | List templates repositories matching any selected filters.
[**TemplatesRepositoryView**](TemplatesRepositoryAPI.md#TemplatesRepositoryView) | **Get** /v1/templates_repositories/{id} | View a specific templates repository.



## TemplatesRepositoryList

> PaginatedTemplatesRepositoryCollection TemplatesRepositoryList(ctx).Page(page).Organization(organization).Execute()

List templates repositories matching any selected filters.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesRepositoryAPI.TemplatesRepositoryList(context.Background()).Page(page).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesRepositoryAPI.TemplatesRepositoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesRepositoryList`: PaginatedTemplatesRepositoryCollection
    fmt.Fprintf(os.Stdout, "Response from `TemplatesRepositoryAPI.TemplatesRepositoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesRepositoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedTemplatesRepositoryCollection**](PaginatedTemplatesRepositoryCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesRepositoryView

> TemplatesRepositoryItem TemplatesRepositoryView(ctx, id).Execute()

View a specific templates repository.



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
    resp, r, err := apiClient.TemplatesRepositoryAPI.TemplatesRepositoryView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesRepositoryAPI.TemplatesRepositoryView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesRepositoryView`: TemplatesRepositoryItem
    fmt.Fprintf(os.Stdout, "Response from `TemplatesRepositoryAPI.TemplatesRepositoryView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesRepositoryViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplatesRepositoryItem**](TemplatesRepositoryItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

