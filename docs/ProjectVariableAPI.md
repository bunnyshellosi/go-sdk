# \ProjectVariableAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectVariableCreate**](ProjectVariableAPI.md#ProjectVariableCreate) | **Post** /v1/project_variables | Create a specific project variable.
[**ProjectVariableDelete**](ProjectVariableAPI.md#ProjectVariableDelete) | **Delete** /v1/project_variables/{id} | Delete a specific project variable.
[**ProjectVariableEdit**](ProjectVariableAPI.md#ProjectVariableEdit) | **Patch** /v1/project_variables/{id} | Edit a specific project variable.
[**ProjectVariableList**](ProjectVariableAPI.md#ProjectVariableList) | **Get** /v1/project_variables | List project variables matching any selected filters.
[**ProjectVariableView**](ProjectVariableAPI.md#ProjectVariableView) | **Get** /v1/project_variables/{id} | View a specific project variable.



## ProjectVariableCreate

> ProjectVariableItem ProjectVariableCreate(ctx).ProjectVariableCreateAction(projectVariableCreateAction).Execute()

Create a specific project variable.



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
    projectVariableCreateAction := *openapiclient.NewProjectVariableCreateAction("Name_example", "Value_example", "Project_example") // ProjectVariableCreateAction | The new project_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectVariableAPI.ProjectVariableCreate(context.Background()).ProjectVariableCreateAction(projectVariableCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectVariableAPI.ProjectVariableCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectVariableCreate`: ProjectVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectVariableAPI.ProjectVariableCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectVariableCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectVariableCreateAction** | [**ProjectVariableCreateAction**](ProjectVariableCreateAction.md) | The new project_variable resource | 

### Return type

[**ProjectVariableItem**](ProjectVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectVariableDelete

> ProjectVariableDelete(ctx, id).Body(body).Execute()

Delete a specific project variable.



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
    r, err := apiClient.ProjectVariableAPI.ProjectVariableDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectVariableAPI.ProjectVariableDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectVariableDeleteRequest struct via the builder pattern


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


## ProjectVariableEdit

> ProjectVariableItem ProjectVariableEdit(ctx, id).ProjectVariableEditAction(projectVariableEditAction).Execute()

Edit a specific project variable.



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
    projectVariableEditAction := *openapiclient.NewProjectVariableEditAction("Value_example") // ProjectVariableEditAction | The updated project_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectVariableAPI.ProjectVariableEdit(context.Background(), id).ProjectVariableEditAction(projectVariableEditAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectVariableAPI.ProjectVariableEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectVariableEdit`: ProjectVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectVariableAPI.ProjectVariableEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectVariableEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectVariableEditAction** | [**ProjectVariableEditAction**](ProjectVariableEditAction.md) | The updated project_variable resource | 

### Return type

[**ProjectVariableItem**](ProjectVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectVariableList

> PaginatedProjectVariableCollection ProjectVariableList(ctx).Page(page).Name(name).Project(project).Organization(organization).Execute()

List project variables matching any selected filters.



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
    project := "project_example" // string | Filter by project (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectVariableAPI.ProjectVariableList(context.Background()).Page(page).Name(name).Project(project).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectVariableAPI.ProjectVariableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectVariableList`: PaginatedProjectVariableCollection
    fmt.Fprintf(os.Stdout, "Response from `ProjectVariableAPI.ProjectVariableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectVariableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **name** | **string** | Filter by name | 
 **project** | **string** | Filter by project | 
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedProjectVariableCollection**](PaginatedProjectVariableCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectVariableView

> ProjectVariableItem ProjectVariableView(ctx, id).Execute()

View a specific project variable.



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
    resp, r, err := apiClient.ProjectVariableAPI.ProjectVariableView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectVariableAPI.ProjectVariableView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectVariableView`: ProjectVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ProjectVariableAPI.ProjectVariableView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectVariableViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectVariableItem**](ProjectVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

