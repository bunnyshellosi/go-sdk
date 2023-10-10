# \TemplateAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateDefinition**](TemplateAPI.md#TemplateDefinition) | **Get** /v1/templates/{id}/definition | View the environment definition.
[**TemplateList**](TemplateAPI.md#TemplateList) | **Get** /v1/templates | List templates matching any selected filters.
[**TemplateValidate**](TemplateAPI.md#TemplateValidate) | **Post** /v1/templates/validate | Validates a given template from an external source.
[**TemplateView**](TemplateAPI.md#TemplateView) | **Get** /v1/templates/{id} | View a specific template.



## TemplateDefinition

> map[string]interface{} TemplateDefinition(ctx, id).Execute()

View the environment definition.



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
    resp, r, err := apiClient.TemplateAPI.TemplateDefinition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.TemplateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.TemplateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x+yaml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateList

> PaginatedTemplateCollection TemplateList(ctx).Page(page).Organization(organization).TemplatesRepository(templatesRepository).Source(source).Search(search).Execute()

List templates matching any selected filters.



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
    templatesRepository := "templatesRepository_example" // string | Filter by templatesRepository (optional)
    source := "all" // string | Filter used to fetch templates by source.  * `public` - contains Bunnyshell curated templates  * `private` - contains your organizations templates  * `all` - contains templates for every source (optional)
    search := "search_example" // string | Filter by search (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.TemplateList(context.Background()).Page(page).Organization(organization).TemplatesRepository(templatesRepository).Source(source).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.TemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateList`: PaginatedTemplateCollection
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.TemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **templatesRepository** | **string** | Filter by templatesRepository | 
 **source** | **string** | Filter used to fetch templates by source.  * &#x60;public&#x60; - contains Bunnyshell curated templates  * &#x60;private&#x60; - contains your organizations templates  * &#x60;all&#x60; - contains templates for every source | 
 **search** | **string** | Filter by search | 

### Return type

[**PaginatedTemplateCollection**](PaginatedTemplateCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateValidate

> TemplateCollection TemplateValidate(ctx).TemplateValidateAction(templateValidateAction).Execute()

Validates a given template from an external source.



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
    templateValidateAction := *openapiclient.NewTemplateValidateAction() // TemplateValidateAction | The new template resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.TemplateValidate(context.Background()).TemplateValidateAction(templateValidateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.TemplateValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateValidate`: TemplateCollection
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.TemplateValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateValidateAction** | [**TemplateValidateAction**](TemplateValidateAction.md) | The new template resource | 

### Return type

[**TemplateCollection**](TemplateCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateView

> TemplateItem TemplateView(ctx, id).Execute()

View a specific template.



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
    resp, r, err := apiClient.TemplateAPI.TemplateView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.TemplateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateView`: TemplateItem
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.TemplateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateItem**](TemplateItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

