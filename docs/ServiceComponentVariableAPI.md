# \ServiceComponentVariableAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceComponentVariableCreate**](ServiceComponentVariableAPI.md#ServiceComponentVariableCreate) | **Post** /v1/service_component_variables | Create a specific component variable.
[**ServiceComponentVariableDelete**](ServiceComponentVariableAPI.md#ServiceComponentVariableDelete) | **Delete** /v1/service_component_variables/{id} | Delete a specific component variable.
[**ServiceComponentVariableEdit**](ServiceComponentVariableAPI.md#ServiceComponentVariableEdit) | **Patch** /v1/service_component_variables/{id} | Edit a specific component variable.
[**ServiceComponentVariableList**](ServiceComponentVariableAPI.md#ServiceComponentVariableList) | **Get** /v1/service_component_variables | List component variables matching any selected filters.
[**ServiceComponentVariableView**](ServiceComponentVariableAPI.md#ServiceComponentVariableView) | **Get** /v1/service_component_variables/{id} | View a specific component variable.



## ServiceComponentVariableCreate

> ServiceComponentVariableItem ServiceComponentVariableCreate(ctx).ServiceComponentVariableCreateAction(serviceComponentVariableCreateAction).Execute()

Create a specific component variable.



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
    serviceComponentVariableCreateAction := *openapiclient.NewServiceComponentVariableCreateAction("Name_example", "Value_example", "ServiceComponent_example") // ServiceComponentVariableCreateAction | The new service_component_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceComponentVariableAPI.ServiceComponentVariableCreate(context.Background()).ServiceComponentVariableCreateAction(serviceComponentVariableCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceComponentVariableAPI.ServiceComponentVariableCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceComponentVariableCreate`: ServiceComponentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ServiceComponentVariableAPI.ServiceComponentVariableCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceComponentVariableCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceComponentVariableCreateAction** | [**ServiceComponentVariableCreateAction**](ServiceComponentVariableCreateAction.md) | The new service_component_variable resource | 

### Return type

[**ServiceComponentVariableItem**](ServiceComponentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceComponentVariableDelete

> ServiceComponentVariableDelete(ctx, id).Body(body).Execute()

Delete a specific component variable.



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
    r, err := apiClient.ServiceComponentVariableAPI.ServiceComponentVariableDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceComponentVariableAPI.ServiceComponentVariableDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiServiceComponentVariableDeleteRequest struct via the builder pattern


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


## ServiceComponentVariableEdit

> ServiceComponentVariableItem ServiceComponentVariableEdit(ctx, id).ServiceComponentVariableEditAction(serviceComponentVariableEditAction).Execute()

Edit a specific component variable.



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
    serviceComponentVariableEditAction := *openapiclient.NewServiceComponentVariableEditAction("Value_example") // ServiceComponentVariableEditAction | The updated service_component_variable resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceComponentVariableAPI.ServiceComponentVariableEdit(context.Background(), id).ServiceComponentVariableEditAction(serviceComponentVariableEditAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceComponentVariableAPI.ServiceComponentVariableEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceComponentVariableEdit`: ServiceComponentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ServiceComponentVariableAPI.ServiceComponentVariableEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceComponentVariableEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceComponentVariableEditAction** | [**ServiceComponentVariableEditAction**](ServiceComponentVariableEditAction.md) | The updated service_component_variable resource | 

### Return type

[**ServiceComponentVariableItem**](ServiceComponentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceComponentVariableList

> PaginatedServiceComponentVariableCollection ServiceComponentVariableList(ctx).Page(page).Name(name).ServiceComponent(serviceComponent).ServiceComponentName(serviceComponentName).Environment(environment).Project(project).Organization(organization).Execute()

List component variables matching any selected filters.



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
    serviceComponent := "serviceComponent_example" // string | Filter by serviceComponent (optional)
    serviceComponentName := "serviceComponentName_example" // string | Filter by serviceComponentName (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    project := "project_example" // string | Filter by project (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceComponentVariableAPI.ServiceComponentVariableList(context.Background()).Page(page).Name(name).ServiceComponent(serviceComponent).ServiceComponentName(serviceComponentName).Environment(environment).Project(project).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceComponentVariableAPI.ServiceComponentVariableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceComponentVariableList`: PaginatedServiceComponentVariableCollection
    fmt.Fprintf(os.Stdout, "Response from `ServiceComponentVariableAPI.ServiceComponentVariableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceComponentVariableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **name** | **string** | Filter by name | 
 **serviceComponent** | **string** | Filter by serviceComponent | 
 **serviceComponentName** | **string** | Filter by serviceComponentName | 
 **environment** | **string** | Filter by environment | 
 **project** | **string** | Filter by project | 
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedServiceComponentVariableCollection**](PaginatedServiceComponentVariableCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceComponentVariableView

> ServiceComponentVariableItem ServiceComponentVariableView(ctx, id).Execute()

View a specific component variable.



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
    resp, r, err := apiClient.ServiceComponentVariableAPI.ServiceComponentVariableView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceComponentVariableAPI.ServiceComponentVariableView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceComponentVariableView`: ServiceComponentVariableItem
    fmt.Fprintf(os.Stdout, "Response from `ServiceComponentVariableAPI.ServiceComponentVariableView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceComponentVariableViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceComponentVariableItem**](ServiceComponentVariableItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

