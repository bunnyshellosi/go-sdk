# \ComponentAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentList**](ComponentAPI.md#ComponentList) | **Get** /v1/components | List service components matching any selected filters
[**ComponentRemoteDevConfig**](ComponentAPI.md#ComponentRemoteDevConfig) | **Get** /v1/components/{id}/remotedev/config | Get remote dev config
[**ComponentRemoteDevProfile**](ComponentAPI.md#ComponentRemoteDevProfile) | **Post** /v1/components/{id}/remotedev/profile | Parse, validate and interpolate the provided remoteDevProfile
[**ComponentResources**](ComponentAPI.md#ComponentResources) | **Get** /v1/components/{id}/resources | Get kubernetes resources
[**ComponentView**](ComponentAPI.md#ComponentView) | **Get** /v1/components/{id} | View a specific service component



## ComponentList

> PaginatedComponentCollection ComponentList(ctx).Page(page).Name(name).Environment(environment).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Organization(organization).Project(project).GitRepository(gitRepository).GitBranch(gitBranch).Execute()

List service components matching any selected filters



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
    name := "name_example" // string | Filter by name (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    operationStatus := "draft" // string | Filter by operationStatus (optional)
    clusterStatus := "not_available" // string | Filter by clusterStatus (optional)
    organization := "organization_example" // string | Filter by organization (optional)
    project := "project_example" // string | Filter by project (optional)
    gitRepository := "gitRepository_example" // string | Filter by gitRepository (optional)
    gitBranch := "gitBranch_example" // string | Filter by gitBranch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentAPI.ComponentList(context.Background()).Page(page).Name(name).Environment(environment).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Organization(organization).Project(project).GitRepository(gitRepository).GitBranch(gitBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.ComponentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentList`: PaginatedComponentCollection
    fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.ComponentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **name** | **string** | Filter by name | 
 **environment** | **string** | Filter by environment | 
 **operationStatus** | **string** | Filter by operationStatus | 
 **clusterStatus** | **string** | Filter by clusterStatus | 
 **organization** | **string** | Filter by organization | 
 **project** | **string** | Filter by project | 
 **gitRepository** | **string** | Filter by gitRepository | 
 **gitBranch** | **string** | Filter by gitBranch | 

### Return type

[**PaginatedComponentCollection**](PaginatedComponentCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentRemoteDevConfig

> ComponentConfigItem ComponentRemoteDevConfig(ctx, id).Execute()

Get remote dev config



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
    resp, r, err := apiClient.ComponentAPI.ComponentRemoteDevConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.ComponentRemoteDevConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentRemoteDevConfig`: ComponentConfigItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.ComponentRemoteDevConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentRemoteDevConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentConfigItem**](ComponentConfigItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentRemoteDevProfile

> ComponentProfileItem ComponentRemoteDevProfile(ctx, id).Body(body).Execute()

Parse, validate and interpolate the provided remoteDevProfile



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
    body := interface{}(987) // interface{} | No Request Body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentAPI.ComponentRemoteDevProfile(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.ComponentRemoteDevProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentRemoteDevProfile`: ComponentProfileItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.ComponentRemoteDevProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentRemoteDevProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | No Request Body | 

### Return type

[**ComponentProfileItem**](ComponentProfileItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/x+yaml
- **Accept**: application/json, application/problem+json

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
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    id := "id_example" // string | Resource identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentAPI.ComponentResources(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.ComponentResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentResources`: []ComponentResourceItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.ComponentResources`: %v\n", resp)
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

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

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
    openapiclient "bunnyshell.com/sdk"
)

func main() {
    id := "id_example" // string | Resource identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentAPI.ComponentView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.ComponentView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentView`: ComponentItem
    fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.ComponentView`: %v\n", resp)
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

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

