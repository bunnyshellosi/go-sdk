# \ProjectAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectCreate**](ProjectAPI.md#ProjectCreate) | **Post** /v1/projects | Creates new project.
[**ProjectDelete**](ProjectAPI.md#ProjectDelete) | **Delete** /v1/projects/{id} | Delete a specific project.
[**ProjectEditBuildSettings**](ProjectAPI.md#ProjectEditBuildSettings) | **Patch** /v1/projects/{id}/build-settings | Edit the build settings of a project.
[**ProjectEditSettings**](ProjectAPI.md#ProjectEditSettings) | **Patch** /v1/projects/{id}/settings | Edit a project.
[**ProjectList**](ProjectAPI.md#ProjectList) | **Get** /v1/projects | List projects matching any selected filters.
[**ProjectView**](ProjectAPI.md#ProjectView) | **Get** /v1/projects/{id} | View a specific project.



## ProjectCreate

> ProjectItem ProjectCreate(ctx).ProjectCreateAction(projectCreateAction).Execute()

Creates new project.



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
    projectCreateAction := *openapiclient.NewProjectCreateAction("Name_example", "Organization_example") // ProjectCreateAction | The new project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.ProjectCreate(context.Background()).ProjectCreateAction(projectCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectCreate`: ProjectItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateAction** | [**ProjectCreateAction**](ProjectCreateAction.md) | The new project resource | 

### Return type

[**ProjectItem**](ProjectItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> ProjectDelete(ctx, id).Body(body).Execute()

Delete a specific project.



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
    r, err := apiClient.ProjectAPI.ProjectDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | No Request Body | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectEditBuildSettings

> ProjectItem ProjectEditBuildSettings(ctx, id).ProjectEditBuildSettingsAction(projectEditBuildSettingsAction).Execute()

Edit the build settings of a project.



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
    projectEditBuildSettingsAction := *openapiclient.NewProjectEditBuildSettingsAction("RegistryIntegration_example", "KubernetesIntegration_example", "Cpu_example", NullableInt32(123), NullableInt32(123)) // ProjectEditBuildSettingsAction | The updated project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.ProjectEditBuildSettings(context.Background(), id).ProjectEditBuildSettingsAction(projectEditBuildSettingsAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectEditBuildSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectEditBuildSettings`: ProjectItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectEditBuildSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectEditBuildSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectEditBuildSettingsAction** | [**ProjectEditBuildSettingsAction**](ProjectEditBuildSettingsAction.md) | The updated project resource | 

### Return type

[**ProjectItem**](ProjectItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectEditSettings

> ProjectItem ProjectEditSettings(ctx, id).ProjectEditSettingsAction(projectEditSettingsAction).Execute()

Edit a project.



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
    projectEditSettingsAction := *openapiclient.NewProjectEditSettingsAction("Name_example") // ProjectEditSettingsAction | The updated project resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.ProjectEditSettings(context.Background(), id).ProjectEditSettingsAction(projectEditSettingsAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ProjectEditSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectEditSettings`: ProjectItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ProjectEditSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectEditSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectEditSettingsAction** | [**ProjectEditSettingsAction**](ProjectEditSettingsAction.md) | The updated project resource | 

### Return type

[**ProjectItem**](ProjectItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

