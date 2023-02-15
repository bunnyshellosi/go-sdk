# \KubernetesIntegrationApi

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesIntegrationList**](KubernetesIntegrationApi.md#KubernetesIntegrationList) | **Get** /v1/kubernetes_integrations | List Kubernetes integrations matching any selected filters.
[**KubernetesIntegrationView**](KubernetesIntegrationApi.md#KubernetesIntegrationView) | **Get** /v1/kubernetes_integrations/{id} | View a specific Kubernetes integration.



## KubernetesIntegrationList

> PaginatedKubernetesIntegrationCollection KubernetesIntegrationList(ctx).Page(page).Organization(organization).Environment(environment).Status(status).CloudProvider(cloudProvider).Execute()

List Kubernetes integrations matching any selected filters.



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
    environment := "environment_example" // string | Filter by environment (optional)
    status := "new" // string | Filter by status (optional)
    cloudProvider := "cloudProvider_example" // string | Filter by cloudProvider (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesIntegrationApi.KubernetesIntegrationList(context.Background()).Page(page).Organization(organization).Environment(environment).Status(status).CloudProvider(cloudProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesIntegrationApi.KubernetesIntegrationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesIntegrationList`: PaginatedKubernetesIntegrationCollection
    fmt.Fprintf(os.Stdout, "Response from `KubernetesIntegrationApi.KubernetesIntegrationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIntegrationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **environment** | **string** | Filter by environment | 
 **status** | **string** | Filter by status | 
 **cloudProvider** | **string** | Filter by cloudProvider | 

### Return type

[**PaginatedKubernetesIntegrationCollection**](PaginatedKubernetesIntegrationCollection.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesIntegrationView

> KubernetesIntegrationItem KubernetesIntegrationView(ctx, id).Execute()

View a specific Kubernetes integration.



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
    resp, r, err := apiClient.KubernetesIntegrationApi.KubernetesIntegrationView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesIntegrationApi.KubernetesIntegrationView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KubernetesIntegrationView`: KubernetesIntegrationItem
    fmt.Fprintf(os.Stdout, "Response from `KubernetesIntegrationApi.KubernetesIntegrationView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIntegrationViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesIntegrationItem**](KubernetesIntegrationItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

