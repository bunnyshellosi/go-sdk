# \SshAPI

All URIs are relative to *https://api.environments.bunnyshell.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshConfig**](SshAPI.md#SshConfig) | **Post** /v1/ssh/credentials | Get SSH config for given SSH domain.



## SshConfig

> SshSSHCredentialsItem SshConfig(ctx).SshCredentialsAction(sshCredentialsAction).Execute()

Get SSH config for given SSH domain.



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
    sshCredentialsAction := *openapiclient.NewSshCredentialsAction("ComponentSshHost_example") // SshCredentialsAction | The new ssh resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshAPI.SshConfig(context.Background()).SshCredentialsAction(sshCredentialsAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshAPI.SshConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshConfig`: SshSSHCredentialsItem
    fmt.Fprintf(os.Stdout, "Response from `SshAPI.SshConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshCredentialsAction** | [**SshCredentialsAction**](SshCredentialsAction.md) | The new ssh resource | 

### Return type

[**SshSSHCredentialsItem**](SshSSHCredentialsItem.md)

### Authorization

[JWT](../README.md#JWT), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

