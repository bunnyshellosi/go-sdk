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
*ComponentApi* | [**ComponentList**](docs/ComponentApi.md#componentlist) | **Get** /v1/components | List service components matching any selected filters
*ComponentApi* | [**ComponentResources**](docs/ComponentApi.md#componentresources) | **Get** /v1/components/{id}/resources | Get kubernetes resources
*ComponentApi* | [**ComponentView**](docs/ComponentApi.md#componentview) | **Get** /v1/components/{id} | View a specific service component
*ComponentEndpointApi* | [**ComponentEndpointList**](docs/ComponentEndpointApi.md#componentendpointlist) | **Get** /v1/components/endpoint | List endpoints for service components matching any selected filters
*ComponentEndpointApi* | [**ComponentEndpointView**](docs/ComponentEndpointApi.md#componentendpointview) | **Get** /v1/components/{id}/endpoint | View endpoints for a specific service component
*ComponentGitApi* | [**ComponentGitList**](docs/ComponentGitApi.md#componentgitlist) | **Get** /v1/components/gitinfo | List git info for service components matching any selected filters
*ComponentGitApi* | [**ComponentGitView**](docs/ComponentGitApi.md#componentgitview) | **Get** /v1/components/{id}/gitinfo | View git info for a specific service component
*EnvironmentApi* | [**EnvironmentClone**](docs/EnvironmentApi.md#environmentclone) | **Post** /v1/environments/{id}/clone | Clone an environment.
*EnvironmentApi* | [**EnvironmentCreate**](docs/EnvironmentApi.md#environmentcreate) | **Post** /v1/environments | Creates a new environment.
*EnvironmentApi* | [**EnvironmentDefinition**](docs/EnvironmentApi.md#environmentdefinition) | **Get** /v1/environments/{id}/definition | View the bunnyshell manifest for the environment
*EnvironmentApi* | [**EnvironmentDelete**](docs/EnvironmentApi.md#environmentdelete) | **Post** /v1/environments/{id}/delete | Delete a specific environment.
*EnvironmentApi* | [**EnvironmentDeploy**](docs/EnvironmentApi.md#environmentdeploy) | **Post** /v1/environments/{id}/deploy | Deploy an environment.
*EnvironmentApi* | [**EnvironmentEdit**](docs/EnvironmentApi.md#environmentedit) | **Put** /v1/environments/{id} | Edit an environment.
*EnvironmentApi* | [**EnvironmentEditComponents**](docs/EnvironmentApi.md#environmenteditcomponents) | **Put** /v1/environments/{id}/components | Edit the components of an environment.
*EnvironmentApi* | [**EnvironmentKubeConfig**](docs/EnvironmentApi.md#environmentkubeconfig) | **Get** /v1/environments/{id}/kube-config | Download Kubernetes Config File
*EnvironmentApi* | [**EnvironmentList**](docs/EnvironmentApi.md#environmentlist) | **Get** /v1/environments | List environments matching any selected filters.
*EnvironmentApi* | [**EnvironmentStart**](docs/EnvironmentApi.md#environmentstart) | **Post** /v1/environments/{id}/start | Start an environment.
*EnvironmentApi* | [**EnvironmentStop**](docs/EnvironmentApi.md#environmentstop) | **Post** /v1/environments/{id}/stop | Stop an environment.
*EnvironmentApi* | [**EnvironmentView**](docs/EnvironmentApi.md#environmentview) | **Get** /v1/environments/{id} | View a specific environment.
*EnvironmentVariableApi* | [**EnvironmentVariableEdit**](docs/EnvironmentVariableApi.md#environmentvariableedit) | **Patch** /v1/environment_variables/{id} | Edit a specific environment variable.
*EnvironmentVariableApi* | [**EnvironmentVariableList**](docs/EnvironmentVariableApi.md#environmentvariablelist) | **Get** /v1/environment_variables | List environment variables matching any selected filters.
*EnvironmentVariableApi* | [**EnvironmentVariableView**](docs/EnvironmentVariableApi.md#environmentvariableview) | **Get** /v1/environment_variables/{id} | View a specific environment variable.
*EventApi* | [**EventList**](docs/EventApi.md#eventlist) | **Get** /v1/events | List events matching any selected filters.
*EventApi* | [**EventView**](docs/EventApi.md#eventview) | **Get** /v1/events/{id} | View a specific event.
*KubernetesIntegrationApi* | [**KubernetesIntegrationList**](docs/KubernetesIntegrationApi.md#kubernetesintegrationlist) | **Get** /v1/kubernetes_integrations | List Kubernetes integrations matching any selected filters.
*KubernetesIntegrationApi* | [**KubernetesIntegrationView**](docs/KubernetesIntegrationApi.md#kubernetesintegrationview) | **Get** /v1/kubernetes_integrations/{id} | View a specific Kubernetes integration.
*OrganizationApi* | [**OrganizationList**](docs/OrganizationApi.md#organizationlist) | **Get** /v1/organizations | List organization matching any selected filters.
*OrganizationApi* | [**OrganizationView**](docs/OrganizationApi.md#organizationview) | **Get** /v1/organizations/{id} | View a specific organization.
*PipelineApi* | [**PipelineList**](docs/PipelineApi.md#pipelinelist) | **Get** /v1/pipelines | List pipelines matching any selected filters.
*PipelineApi* | [**PipelineView**](docs/PipelineApi.md#pipelineview) | **Get** /v1/pipelines/{id} | View a specific Pipeline.
*ProjectApi* | [**ProjectList**](docs/ProjectApi.md#projectlist) | **Get** /v1/projects | List projects matching any selected filters.
*ProjectApi* | [**ProjectView**](docs/ProjectApi.md#projectview) | **Get** /v1/projects/{id} | View a specific project.
*TemplateApi* | [**TemplateDefinition**](docs/TemplateApi.md#templatedefinition) | **Get** /v1/templates/{id}/definition | View the environment definition.
*TemplateApi* | [**TemplateList**](docs/TemplateApi.md#templatelist) | **Get** /v1/templates | List templates matching any selected filters.
*TemplateApi* | [**TemplateValidate**](docs/TemplateApi.md#templatevalidate) | **Post** /v1/templates/validate | Validates a given template from an external source.
*TemplateApi* | [**TemplateView**](docs/TemplateApi.md#templateview) | **Get** /v1/templates/{id} | View a specific template.
*TemplatesRepositoryApi* | [**TemplatesRepositoryList**](docs/TemplatesRepositoryApi.md#templatesrepositorylist) | **Get** /v1/templates_repositories | List templates repositories matching any selected filters.
*TemplatesRepositoryApi* | [**TemplatesRepositoryView**](docs/TemplatesRepositoryApi.md#templatesrepositoryview) | **Get** /v1/templates_repositories/{id} | View a specific templates repository.


## Documentation For Models

 - [ClusterKubeConfigRead](docs/ClusterKubeConfigRead.md)
 - [ClusterWrapperKubeConfigRead](docs/ClusterWrapperKubeConfigRead.md)
 - [ComponentCollection](docs/ComponentCollection.md)
 - [ComponentEndpointCollection](docs/ComponentEndpointCollection.md)
 - [ComponentEndpointItem](docs/ComponentEndpointItem.md)
 - [ComponentGitCollection](docs/ComponentGitCollection.md)
 - [ComponentGitItem](docs/ComponentGitItem.md)
 - [ComponentItem](docs/ComponentItem.md)
 - [ComponentResourceItem](docs/ComponentResourceItem.md)
 - [ContextKubeConfigRead](docs/ContextKubeConfigRead.md)
 - [ContextWrapperKubeConfigRead](docs/ContextWrapperKubeConfigRead.md)
 - [EmbeddedComponentCollection](docs/EmbeddedComponentCollection.md)
 - [EmbeddedComponentEndpointCollection](docs/EmbeddedComponentEndpointCollection.md)
 - [EmbeddedComponentGitCollection](docs/EmbeddedComponentGitCollection.md)
 - [EmbeddedEnvironmentCollection](docs/EmbeddedEnvironmentCollection.md)
 - [EmbeddedEnvironmentVariableCollection](docs/EmbeddedEnvironmentVariableCollection.md)
 - [EmbeddedEventCollection](docs/EmbeddedEventCollection.md)
 - [EmbeddedKubernetesIntegrationCollection](docs/EmbeddedKubernetesIntegrationCollection.md)
 - [EmbeddedOrganizationCollection](docs/EmbeddedOrganizationCollection.md)
 - [EmbeddedPipelineCollection](docs/EmbeddedPipelineCollection.md)
 - [EmbeddedProjectCollection](docs/EmbeddedProjectCollection.md)
 - [EmbeddedTemplateCollection](docs/EmbeddedTemplateCollection.md)
 - [EmbeddedTemplatesRepositoryCollection](docs/EmbeddedTemplatesRepositoryCollection.md)
 - [EnvironmentCloneAction](docs/EnvironmentCloneAction.md)
 - [EnvironmentCollection](docs/EnvironmentCollection.md)
 - [EnvironmentCreateAction](docs/EnvironmentCreateAction.md)
 - [EnvironmentCreateActionGenesis](docs/EnvironmentCreateActionGenesis.md)
 - [EnvironmentEditAction](docs/EnvironmentEditAction.md)
 - [EnvironmentEditActionEdit](docs/EnvironmentEditActionEdit.md)
 - [EnvironmentEditActionGenesis](docs/EnvironmentEditActionGenesis.md)
 - [EnvironmentEditComponentsAction](docs/EnvironmentEditComponentsAction.md)
 - [EnvironmentEditComponentsActionFilter](docs/EnvironmentEditComponentsActionFilter.md)
 - [EnvironmentItem](docs/EnvironmentItem.md)
 - [EnvironmentKubeConfigKubeConfigRead](docs/EnvironmentKubeConfigKubeConfigRead.md)
 - [EnvironmentVariableCollection](docs/EnvironmentVariableCollection.md)
 - [EnvironmentVariableEdit](docs/EnvironmentVariableEdit.md)
 - [EnvironmentVariableItem](docs/EnvironmentVariableItem.md)
 - [Ephemeral](docs/Ephemeral.md)
 - [EventCollection](docs/EventCollection.md)
 - [EventItem](docs/EventItem.md)
 - [FilterGit](docs/FilterGit.md)
 - [FilterName](docs/FilterName.md)
 - [FromGit](docs/FromGit.md)
 - [FromGitSpec](docs/FromGitSpec.md)
 - [FromString](docs/FromString.md)
 - [FromTemplate](docs/FromTemplate.md)
 - [GitInfo](docs/GitInfo.md)
 - [KubernetesIntegrationCollection](docs/KubernetesIntegrationCollection.md)
 - [KubernetesIntegrationItem](docs/KubernetesIntegrationItem.md)
 - [OrganizationCollection](docs/OrganizationCollection.md)
 - [OrganizationItem](docs/OrganizationItem.md)
 - [PaginatedComponentCollection](docs/PaginatedComponentCollection.md)
 - [PaginatedComponentEndpointCollection](docs/PaginatedComponentEndpointCollection.md)
 - [PaginatedComponentGitCollection](docs/PaginatedComponentGitCollection.md)
 - [PaginatedEnvironmentCollection](docs/PaginatedEnvironmentCollection.md)
 - [PaginatedEnvironmentVariableCollection](docs/PaginatedEnvironmentVariableCollection.md)
 - [PaginatedEventCollection](docs/PaginatedEventCollection.md)
 - [PaginatedKubernetesIntegrationCollection](docs/PaginatedKubernetesIntegrationCollection.md)
 - [PaginatedLink](docs/PaginatedLink.md)
 - [PaginatedLinks](docs/PaginatedLinks.md)
 - [PaginatedLinksFirst](docs/PaginatedLinksFirst.md)
 - [PaginatedOrganizationCollection](docs/PaginatedOrganizationCollection.md)
 - [PaginatedPipelineCollection](docs/PaginatedPipelineCollection.md)
 - [PaginatedProjectCollection](docs/PaginatedProjectCollection.md)
 - [PaginatedTemplateCollection](docs/PaginatedTemplateCollection.md)
 - [PaginatedTemplatesRepositoryCollection](docs/PaginatedTemplatesRepositoryCollection.md)
 - [PipelineCollection](docs/PipelineCollection.md)
 - [PipelineItem](docs/PipelineItem.md)
 - [Primary](docs/Primary.md)
 - [ProblemGeneric](docs/ProblemGeneric.md)
 - [ProblemViolation](docs/ProblemViolation.md)
 - [ProjectCollection](docs/ProjectCollection.md)
 - [ProjectItem](docs/ProjectItem.md)
 - [StageItem](docs/StageItem.md)
 - [TemplateCollection](docs/TemplateCollection.md)
 - [TemplateItem](docs/TemplateItem.md)
 - [TemplateValidateAction](docs/TemplateValidateAction.md)
 - [TemplateValidateActionSource](docs/TemplateValidateActionSource.md)
 - [TemplatesRepositoryCollection](docs/TemplatesRepositoryCollection.md)
 - [TemplatesRepositoryItem](docs/TemplatesRepositoryItem.md)
 - [UserKubeConfigRead](docs/UserKubeConfigRead.md)
 - [UserWrapperKubeConfigRead](docs/UserWrapperKubeConfigRead.md)
 - [ValidateSourceGit](docs/ValidateSourceGit.md)
 - [ValidateSourceString](docs/ValidateSourceString.md)


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

