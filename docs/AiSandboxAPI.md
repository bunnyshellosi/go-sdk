# \AiSandboxAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiSandboxConvertDockerCompose**](AiSandboxAPI.md#AiSandboxConvertDockerCompose) | **Post** /v1/convert/docker-compose | Convert docker compose to bunnyshell.yaml.



## AiSandboxConvertDockerCompose

> map[string]interface{} AiSandboxConvertDockerCompose(ctx).AiSandboxConvertDockerComposeAction(aiSandboxConvertDockerComposeAction).Execute()

Convert docker compose to bunnyshell.yaml.



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
    aiSandboxConvertDockerComposeAction := *openapiclient.NewAiSandboxConvertDockerComposeAction() // AiSandboxConvertDockerComposeAction | The new ai_sandbox resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AiSandboxAPI.AiSandboxConvertDockerCompose(context.Background()).AiSandboxConvertDockerComposeAction(aiSandboxConvertDockerComposeAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AiSandboxAPI.AiSandboxConvertDockerCompose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AiSandboxConvertDockerCompose`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AiSandboxAPI.AiSandboxConvertDockerCompose`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSandboxConvertDockerComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSandboxConvertDockerComposeAction** | [**AiSandboxConvertDockerComposeAction**](AiSandboxConvertDockerComposeAction.md) | The new ai_sandbox resource | 

### Return type

**map[string]interface{}**

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

