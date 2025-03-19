# \EnvironmentAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentClone**](EnvironmentAPI.md#EnvironmentClone) | **Post** /v1/environments/{id}/clone | Clone an environment.
[**EnvironmentCreate**](EnvironmentAPI.md#EnvironmentCreate) | **Post** /v1/environments | Creates a new environment.
[**EnvironmentDefinition**](EnvironmentAPI.md#EnvironmentDefinition) | **Get** /v1/environments/{id}/definition | View the bunnyshell manifest for the environment
[**EnvironmentDelete**](EnvironmentAPI.md#EnvironmentDelete) | **Post** /v1/environments/{id}/delete | Delete a specific environment.
[**EnvironmentDeploy**](EnvironmentAPI.md#EnvironmentDeploy) | **Post** /v1/environments/{id}/deploy | Deploy an environment.
[**EnvironmentEditBuildSettings**](EnvironmentAPI.md#EnvironmentEditBuildSettings) | **Patch** /v1/environments/{id}/build-settings | Edit the build settings of an environment.
[**EnvironmentEditComponents**](EnvironmentAPI.md#EnvironmentEditComponents) | **Put** /v1/environments/{id}/components | Edit the components of an environment.
[**EnvironmentEditConfiguration**](EnvironmentAPI.md#EnvironmentEditConfiguration) | **Put** /v1/environments/{id}/configuration | Edit an environment.
[**EnvironmentEditSettings**](EnvironmentAPI.md#EnvironmentEditSettings) | **Put** /v1/environments/{id}/settings | Edit an environment.
[**EnvironmentKubeConfig**](EnvironmentAPI.md#EnvironmentKubeConfig) | **Get** /v1/environments/{id}/kube-config | Download Kubernetes Config File
[**EnvironmentList**](EnvironmentAPI.md#EnvironmentList) | **Get** /v1/environments | List environments matching any selected filters.
[**EnvironmentStart**](EnvironmentAPI.md#EnvironmentStart) | **Post** /v1/environments/{id}/start | Start an environment.
[**EnvironmentStop**](EnvironmentAPI.md#EnvironmentStop) | **Post** /v1/environments/{id}/stop | Stop an environment.
[**EnvironmentView**](EnvironmentAPI.md#EnvironmentView) | **Get** /v1/environments/{id} | View a specific environment.



## EnvironmentClone

> EnvironmentItem EnvironmentClone(ctx, id).EnvironmentCloneAction(environmentCloneAction).Execute()

Clone an environment.



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
    environmentCloneAction := *openapiclient.NewEnvironmentCloneAction("Name_example") // EnvironmentCloneAction | The new environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentClone(context.Background(), id).EnvironmentCloneAction(environmentCloneAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentClone`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentCloneAction** | [**EnvironmentCloneAction**](EnvironmentCloneAction.md) | The new environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentCreate

> EnvironmentItem EnvironmentCreate(ctx).EnvironmentCreateAction(environmentCreateAction).Execute()

Creates a new environment.



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
    environmentCreateAction := *openapiclient.NewEnvironmentCreateAction("Name_example", "Project_example") // EnvironmentCreateAction | The new environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentCreate(context.Background()).EnvironmentCreateAction(environmentCreateAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentCreate`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentCreateAction** | [**EnvironmentCreateAction**](EnvironmentCreateAction.md) | The new environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentDefinition

> map[string]interface{} EnvironmentDefinition(ctx, id).Execute()

View the bunnyshell manifest for the environment



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
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentDefinition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentDefinitionRequest struct via the builder pattern


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


## EnvironmentDelete

> EventItem EnvironmentDelete(ctx, id).Body(body).Execute()

Delete a specific environment.



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
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentDelete`: EventItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | No Request Body | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentDeploy

> EventItem EnvironmentDeploy(ctx, id).EnvironmentPartialDeployAction(environmentPartialDeployAction).Execute()

Deploy an environment.



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
    environmentPartialDeployAction := *openapiclient.NewEnvironmentPartialDeployAction() // EnvironmentPartialDeployAction | The new environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentDeploy(context.Background(), id).EnvironmentPartialDeployAction(environmentPartialDeployAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentDeploy`: EventItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentPartialDeployAction** | [**EnvironmentPartialDeployAction**](EnvironmentPartialDeployAction.md) | The new environment resource | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentEditBuildSettings

> EnvironmentItem EnvironmentEditBuildSettings(ctx, id).EnvironmentEditBuildSettingsAction(environmentEditBuildSettingsAction).Execute()

Edit the build settings of an environment.



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
    environmentEditBuildSettingsAction := *openapiclient.NewEnvironmentEditBuildSettingsAction("RegistryIntegration_example", "KubernetesIntegration_example", "Cpu_example", NullableInt32(123), NullableInt32(123)) // EnvironmentEditBuildSettingsAction | The updated environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentEditBuildSettings(context.Background(), id).EnvironmentEditBuildSettingsAction(environmentEditBuildSettingsAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentEditBuildSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentEditBuildSettings`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentEditBuildSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentEditBuildSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEditBuildSettingsAction** | [**EnvironmentEditBuildSettingsAction**](EnvironmentEditBuildSettingsAction.md) | The updated environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentEditComponents

> EnvironmentItem EnvironmentEditComponents(ctx, id).EnvironmentEditComponentsAction(environmentEditComponentsAction).Execute()

Edit the components of an environment.



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
    environmentEditComponentsAction := *openapiclient.NewEnvironmentEditComponentsAction(openapiclient.environment_EditComponentsAction_filter{FilterGit: openapiclient.NewFilterGit()}, *openapiclient.NewGitInfo()) // EnvironmentEditComponentsAction | The updated environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentEditComponents(context.Background(), id).EnvironmentEditComponentsAction(environmentEditComponentsAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentEditComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentEditComponents`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentEditComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentEditComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEditComponentsAction** | [**EnvironmentEditComponentsAction**](EnvironmentEditComponentsAction.md) | The updated environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentEditConfiguration

> EnvironmentItem EnvironmentEditConfiguration(ctx, id).EnvironmentEditConfiguration(environmentEditConfiguration).Execute()

Edit an environment.



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
    environmentEditConfiguration := *openapiclient.NewEnvironmentEditConfiguration() // EnvironmentEditConfiguration | The updated environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentEditConfiguration(context.Background(), id).EnvironmentEditConfiguration(environmentEditConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentEditConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentEditConfiguration`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentEditConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentEditConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEditConfiguration** | [**EnvironmentEditConfiguration**](EnvironmentEditConfiguration.md) | The updated environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentEditSettings

> EnvironmentItem EnvironmentEditSettings(ctx, id).EnvironmentEditSettings(environmentEditSettings).Execute()

Edit an environment.



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
    environmentEditSettings := *openapiclient.NewEnvironmentEditSettings() // EnvironmentEditSettings | The updated environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentEditSettings(context.Background(), id).EnvironmentEditSettings(environmentEditSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentEditSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentEditSettings`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentEditSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentEditSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEditSettings** | [**EnvironmentEditSettings**](EnvironmentEditSettings.md) | The updated environment resource | 

### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentKubeConfig

> EnvironmentKubeConfigKubeConfigRead EnvironmentKubeConfig(ctx, id).Execute()

Download Kubernetes Config File



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
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentKubeConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentKubeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentKubeConfig`: EnvironmentKubeConfigKubeConfigRead
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentKubeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentKubeConfigKubeConfigRead**](EnvironmentKubeConfigKubeConfigRead.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x+yaml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentList

> PaginatedEnvironmentCollection EnvironmentList(ctx).Page(page).Organization(organization).KubernetesIntegration(kubernetesIntegration).Type_(type_).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Search(search).ComponentGitRepository(componentGitRepository).ComponentGitBranch(componentGitBranch).Project(project).Labels(labels).Execute()

List environments matching any selected filters.



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
    kubernetesIntegration := "kubernetesIntegration_example" // string | Filter by kubernetesIntegration (optional)
    type_ := "primary" // string | Filter by type (optional)
    operationStatus := "draft" // string | Filter by operationStatus (optional)
    clusterStatus := "running_ok" // string | Filter by clusterStatus (optional)
    search := "search_example" // string | Filter by search (optional)
    componentGitRepository := "componentGitRepository_example" // string | Filter by componentGitRepository (optional)
    componentGitBranch := "componentGitBranch_example" // string | Filter by componentGitBranch (optional)
    project := "project_example" // string | Filter by project (optional)
    labels := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | Filter by label key-value pair. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentList(context.Background()).Page(page).Organization(organization).KubernetesIntegration(kubernetesIntegration).Type_(type_).OperationStatus(operationStatus).ClusterStatus(clusterStatus).Search(search).ComponentGitRepository(componentGitRepository).ComponentGitBranch(componentGitBranch).Project(project).Labels(labels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentList`: PaginatedEnvironmentCollection
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **organization** | **string** | Filter by organization | 
 **kubernetesIntegration** | **string** | Filter by kubernetesIntegration | 
 **type_** | **string** | Filter by type | 
 **operationStatus** | **string** | Filter by operationStatus | 
 **clusterStatus** | **string** | Filter by clusterStatus | 
 **search** | **string** | Filter by search | 
 **componentGitRepository** | **string** | Filter by componentGitRepository | 
 **componentGitBranch** | **string** | Filter by componentGitBranch | 
 **project** | **string** | Filter by project | 
 **labels** | **map[string]map[string]string** | Filter by label key-value pair. | 

### Return type

[**PaginatedEnvironmentCollection**](PaginatedEnvironmentCollection.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentStart

> EventItem EnvironmentStart(ctx, id).EnvironmentPartialStartAction(environmentPartialStartAction).Execute()

Start an environment.



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
    environmentPartialStartAction := *openapiclient.NewEnvironmentPartialStartAction() // EnvironmentPartialStartAction | The new environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentStart(context.Background(), id).EnvironmentPartialStartAction(environmentPartialStartAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentStart`: EventItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentPartialStartAction** | [**EnvironmentPartialStartAction**](EnvironmentPartialStartAction.md) | The new environment resource | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentStop

> EventItem EnvironmentStop(ctx, id).EnvironmentPartialAction(environmentPartialAction).Execute()

Stop an environment.



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
    environmentPartialAction := *openapiclient.NewEnvironmentPartialAction() // EnvironmentPartialAction | The new environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentStop(context.Background(), id).EnvironmentPartialAction(environmentPartialAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentStop`: EventItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentPartialAction** | [**EnvironmentPartialAction**](EnvironmentPartialAction.md) | The new environment resource | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentView

> EnvironmentItem EnvironmentView(ctx, id).Execute()

View a specific environment.



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
    resp, r, err := apiClient.EnvironmentAPI.EnvironmentView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentView`: EnvironmentItem
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentItem**](EnvironmentItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

