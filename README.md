# Go API client for sdk

Interact with Bunnyshell Platform

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 0.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.bunnyshell.com/](https://www.bunnyshell.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sdk "bunnyshell.com/sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.environments.bunnyshell.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ComponentApi* | [**ComponentList**](docs/ComponentApi.md#componentlist) | **Get** /v1/components | Retrieves the collection of component resources.
*ComponentApi* | [**ComponentResources**](docs/ComponentApi.md#componentresources) | **Get** /v1/components/{id}/resources | Retrieves a component resource.
*ComponentApi* | [**ComponentView**](docs/ComponentApi.md#componentview) | **Get** /v1/components/{id} | Retrieves a component resource.
*EnvironmentApi* | [**EnvironmentClone**](docs/EnvironmentApi.md#environmentclone) | **Post** /v1/environments/{id}/clone | Creates a environment resource.
*EnvironmentApi* | [**EnvironmentDelete**](docs/EnvironmentApi.md#environmentdelete) | **Post** /v1/environments/{id}/delete | Delete a specific environment.
*EnvironmentApi* | [**EnvironmentDeploy**](docs/EnvironmentApi.md#environmentdeploy) | **Post** /v1/environments/{id}/deploy | Creates a environment resource.
*EnvironmentApi* | [**EnvironmentKubeConfig**](docs/EnvironmentApi.md#environmentkubeconfig) | **Get** /v1/environments/{id}/kube-config | Retrieves a environment resource.
*EnvironmentApi* | [**EnvironmentList**](docs/EnvironmentApi.md#environmentlist) | **Get** /v1/environments | Retrieves the collection of environment resources.
*EnvironmentApi* | [**EnvironmentStart**](docs/EnvironmentApi.md#environmentstart) | **Post** /v1/environments/{id}/start | Start an environment.
*EnvironmentApi* | [**EnvironmentStop**](docs/EnvironmentApi.md#environmentstop) | **Post** /v1/environments/{id}/stop | Creates a environment resource.
*EnvironmentApi* | [**EnvironmentView**](docs/EnvironmentApi.md#environmentview) | **Get** /v1/environments/{id} | Retrieves a environment resource.
*EnvironmentVariableApi* | [**EnvironmentVariableEdit**](docs/EnvironmentVariableApi.md#environmentvariableedit) | **Patch** /v1/environment_variables/{id} | Updates the environment_variable resource.
*EnvironmentVariableApi* | [**EnvironmentVariableList**](docs/EnvironmentVariableApi.md#environmentvariablelist) | **Get** /v1/environment_variables | Retrieves the collection of environment_variable resources.
*EnvironmentVariableApi* | [**EnvironmentVariableView**](docs/EnvironmentVariableApi.md#environmentvariableview) | **Get** /v1/environment_variables/{id} | Retrieves a environment_variable resource.
*EventApi* | [**EventList**](docs/EventApi.md#eventlist) | **Get** /v1/events | Retrieves the collection of event resources.
*EventApi* | [**EventView**](docs/EventApi.md#eventview) | **Get** /v1/events/{id} | Retrieves a event resource.
*OrganizationApi* | [**OrganizationList**](docs/OrganizationApi.md#organizationlist) | **Get** /v1/organizations | Retrieves the collection of organization resources.
*OrganizationApi* | [**OrganizationView**](docs/OrganizationApi.md#organizationview) | **Get** /v1/organizations/{id} | Retrieves a organization resource.
*ProjectApi* | [**ProjectList**](docs/ProjectApi.md#projectlist) | **Get** /v1/projects | Retrieves the collection of project resources.
*ProjectApi* | [**ProjectView**](docs/ProjectApi.md#projectview) | **Get** /v1/projects/{id} | Retrieves a project resource.


## Documentation For Models

 - [ClusterKubeConfigRead](docs/ClusterKubeConfigRead.md)
 - [ClusterWrapperKubeConfigRead](docs/ClusterWrapperKubeConfigRead.md)
 - [ComponentCollection](docs/ComponentCollection.md)
 - [ComponentItem](docs/ComponentItem.md)
 - [ComponentResourceItem](docs/ComponentResourceItem.md)
 - [ContextKubeConfigRead](docs/ContextKubeConfigRead.md)
 - [ContextWrapperKubeConfigRead](docs/ContextWrapperKubeConfigRead.md)
 - [EmbeddedComponentCollection](docs/EmbeddedComponentCollection.md)
 - [EmbeddedEnvironmentCollection](docs/EmbeddedEnvironmentCollection.md)
 - [EmbeddedEnvironmentVariableCollection](docs/EmbeddedEnvironmentVariableCollection.md)
 - [EmbeddedEventCollection](docs/EmbeddedEventCollection.md)
 - [EmbeddedOrganizationCollection](docs/EmbeddedOrganizationCollection.md)
 - [EmbeddedProjectCollection](docs/EmbeddedProjectCollection.md)
 - [EnvironmentCloneAction](docs/EnvironmentCloneAction.md)
 - [EnvironmentCollection](docs/EnvironmentCollection.md)
 - [EnvironmentItem](docs/EnvironmentItem.md)
 - [EnvironmentKubeConfigKubeConfigRead](docs/EnvironmentKubeConfigKubeConfigRead.md)
 - [EnvironmentVariableCollection](docs/EnvironmentVariableCollection.md)
 - [EnvironmentVariableEdit](docs/EnvironmentVariableEdit.md)
 - [EnvironmentVariableItem](docs/EnvironmentVariableItem.md)
 - [EventCollection](docs/EventCollection.md)
 - [EventItem](docs/EventItem.md)
 - [OrganizationCollection](docs/OrganizationCollection.md)
 - [OrganizationItem](docs/OrganizationItem.md)
 - [PaginatedComponentCollection](docs/PaginatedComponentCollection.md)
 - [PaginatedEnvironmentCollection](docs/PaginatedEnvironmentCollection.md)
 - [PaginatedEnvironmentVariableCollection](docs/PaginatedEnvironmentVariableCollection.md)
 - [PaginatedEventCollection](docs/PaginatedEventCollection.md)
 - [PaginatedLinkSelf](docs/PaginatedLinkSelf.md)
 - [PaginatedLinks](docs/PaginatedLinks.md)
 - [PaginatedOrganizationCollection](docs/PaginatedOrganizationCollection.md)
 - [PaginatedProjectCollection](docs/PaginatedProjectCollection.md)
 - [ProblemGeneric](docs/ProblemGeneric.md)
 - [ProjectCollection](docs/ProjectCollection.md)
 - [ProjectItem](docs/ProjectItem.md)
 - [UserKubeConfigRead](docs/UserKubeConfigRead.md)
 - [UserWrapperKubeConfigRead](docs/UserWrapperKubeConfigRead.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Auth-Token and passed in as the auth context for each request.


### JWT

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

osi+support@bunnyshell.com

