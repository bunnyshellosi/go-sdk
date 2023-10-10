# \EventAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventList**](EventAPI.md#EventList) | **Get** /v1/events | List events matching any selected filters.
[**EventView**](EventAPI.md#EventView) | **Get** /v1/events/{id} | View a specific event.



## EventList

> PaginatedEventCollection EventList(ctx).Page(page).Type_(type_).Status(status).Environment(environment).Organization(organization).Execute()

List events matching any selected filters.



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
    type_ := "env_create" // string | Filter by type (optional)
    status := "new" // string | Filter by status (optional)
    environment := "environment_example" // string | Filter by environment (optional)
    organization := "organization_example" // string | Filter by organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventAPI.EventList(context.Background()).Page(page).Type_(type_).Status(status).Environment(environment).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.EventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventList`: PaginatedEventCollection
    fmt.Fprintf(os.Stdout, "Response from `EventAPI.EventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **type_** | **string** | Filter by type | 
 **status** | **string** | Filter by status | 
 **environment** | **string** | Filter by environment | 
 **organization** | **string** | Filter by organization | 

### Return type

[**PaginatedEventCollection**](PaginatedEventCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventView

> EventItem EventView(ctx, id).Execute()

View a specific event.



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
    resp, r, err := apiClient.EventAPI.EventView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.EventView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventView`: EventItem
    fmt.Fprintf(os.Stdout, "Response from `EventAPI.EventView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventItem**](EventItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

