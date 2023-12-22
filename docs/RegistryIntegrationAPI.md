# \RegistryIntegrationAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryIntegrationList**](RegistryIntegrationAPI.md#RegistryIntegrationList) | **Get** /v1/registry_integrations | List Registry integrations matching any selected filters.
[**RegistryIntegrationView**](RegistryIntegrationAPI.md#RegistryIntegrationView) | **Get** /v1/registry_integrations/{id} | View a specific Registry integration.



## RegistryIntegrationList

> PaginatedRegistryIntegrationCollection RegistryIntegrationList(ctx).Page(page).Organization(organization).Execute()

List Registry integrations matching any selected filters.



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
    resp, r, err := apiClient.RegistryIntegrationAPI.RegistryIntegrationList(context.Background()).Page(page).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryIntegrationAPI.RegistryIntegrationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistryIntegrationList`: PaginatedRegistryIntegrationCollection
    fmt.Fprintf(os.Stdout, "Response from `RegistryIntegrationAPI.RegistryIntegrationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryIntegrationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedRegistryIntegrationCollection**](PaginatedRegistryIntegrationCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryIntegrationView

> RegistryIntegrationItem RegistryIntegrationView(ctx, id).Execute()

View a specific Registry integration.



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
    resp, r, err := apiClient.RegistryIntegrationAPI.RegistryIntegrationView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryIntegrationAPI.RegistryIntegrationView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistryIntegrationView`: RegistryIntegrationItem
    fmt.Fprintf(os.Stdout, "Response from `RegistryIntegrationAPI.RegistryIntegrationView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryIntegrationViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryIntegrationItem**](RegistryIntegrationItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

