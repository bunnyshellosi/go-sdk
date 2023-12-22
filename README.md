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

```golang
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
*ComponentAPI* | [**ComponentList**](docs/ComponentAPI.md#componentlist) | **Get** /v1/components | List service components matching any selected filters
*ComponentAPI* | [**ComponentRemoteDevConfig**](docs/ComponentAPI.md#componentremotedevconfig) | **Get** /v1/components/{id}/remotedev/config | Get remote dev config
*ComponentAPI* | [**ComponentRemoteDevProfile**](docs/ComponentAPI.md#componentremotedevprofile) | **Post** /v1/components/{id}/remotedev/profile | Parse, validate and interpolate the provided remoteDevProfile
*ComponentAPI* | [**ComponentResources**](docs/ComponentAPI.md#componentresources) | **Get** /v1/components/{id}/resources | Get kubernetes resources
*ComponentAPI* | [**ComponentView**](docs/ComponentAPI.md#componentview) | **Get** /v1/components/{id} | View a specific service component
*ComponentEndpointAPI* | [**ComponentEndpointList**](docs/ComponentEndpointAPI.md#componentendpointlist) | **Get** /v1/components/endpoint | List endpoints for service components matching any selected filters
*ComponentEndpointAPI* | [**ComponentEndpointView**](docs/ComponentEndpointAPI.md#componentendpointview) | **Get** /v1/components/{id}/endpoint | View endpoints for a specific service component
*ComponentGitAPI* | [**ComponentGitList**](docs/ComponentGitAPI.md#componentgitlist) | **Get** /v1/components/gitinfo | List git info for service components matching any selected filters
*ComponentGitAPI* | [**ComponentGitView**](docs/ComponentGitAPI.md#componentgitview) | **Get** /v1/components/{id}/gitinfo | View git info for a specific service component
*EnvironmentAPI* | [**EnvironmentClone**](docs/EnvironmentAPI.md#environmentclone) | **Post** /v1/environments/{id}/clone | Clone an environment.
*EnvironmentAPI* | [**EnvironmentCreate**](docs/EnvironmentAPI.md#environmentcreate) | **Post** /v1/environments | Creates a new environment.
*EnvironmentAPI* | [**EnvironmentDefinition**](docs/EnvironmentAPI.md#environmentdefinition) | **Get** /v1/environments/{id}/definition | View the bunnyshell manifest for the environment
*EnvironmentAPI* | [**EnvironmentDelete**](docs/EnvironmentAPI.md#environmentdelete) | **Post** /v1/environments/{id}/delete | Delete a specific environment.
*EnvironmentAPI* | [**EnvironmentDeploy**](docs/EnvironmentAPI.md#environmentdeploy) | **Post** /v1/environments/{id}/deploy | Deploy an environment.
*EnvironmentAPI* | [**EnvironmentEditBuildSettings**](docs/EnvironmentAPI.md#environmenteditbuildsettings) | **Patch** /v1/environments/{id}/build-settings | Edit the build settings of an environment.
*EnvironmentAPI* | [**EnvironmentEditComponents**](docs/EnvironmentAPI.md#environmenteditcomponents) | **Put** /v1/environments/{id}/components | Edit the components of an environment.
*EnvironmentAPI* | [**EnvironmentEditConfiguration**](docs/EnvironmentAPI.md#environmenteditconfiguration) | **Put** /v1/environments/{id}/configuration | Edit an environment.
*EnvironmentAPI* | [**EnvironmentEditSettings**](docs/EnvironmentAPI.md#environmenteditsettings) | **Put** /v1/environments/{id}/settings | Edit an environment.
*EnvironmentAPI* | [**EnvironmentKubeConfig**](docs/EnvironmentAPI.md#environmentkubeconfig) | **Get** /v1/environments/{id}/kube-config | Download Kubernetes Config File
*EnvironmentAPI* | [**EnvironmentList**](docs/EnvironmentAPI.md#environmentlist) | **Get** /v1/environments | List environments matching any selected filters.
*EnvironmentAPI* | [**EnvironmentStart**](docs/EnvironmentAPI.md#environmentstart) | **Post** /v1/environments/{id}/start | Start an environment.
*EnvironmentAPI* | [**EnvironmentStop**](docs/EnvironmentAPI.md#environmentstop) | **Post** /v1/environments/{id}/stop | Stop an environment.
*EnvironmentAPI* | [**EnvironmentView**](docs/EnvironmentAPI.md#environmentview) | **Get** /v1/environments/{id} | View a specific environment.
*EnvironmentVariableAPI* | [**EnvironmentVariableCreate**](docs/EnvironmentVariableAPI.md#environmentvariablecreate) | **Post** /v1/environment_variables | Create a specific environment variable.
*EnvironmentVariableAPI* | [**EnvironmentVariableDelete**](docs/EnvironmentVariableAPI.md#environmentvariabledelete) | **Delete** /v1/environment_variables/{id} | Delete a specific environment variable.
*EnvironmentVariableAPI* | [**EnvironmentVariableEdit**](docs/EnvironmentVariableAPI.md#environmentvariableedit) | **Patch** /v1/environment_variables/{id} | Edit a specific environment variable.
*EnvironmentVariableAPI* | [**EnvironmentVariableList**](docs/EnvironmentVariableAPI.md#environmentvariablelist) | **Get** /v1/environment_variables | List environment variables matching any selected filters.
*EnvironmentVariableAPI* | [**EnvironmentVariableView**](docs/EnvironmentVariableAPI.md#environmentvariableview) | **Get** /v1/environment_variables/{id} | View a specific environment variable.
*EventAPI* | [**EventList**](docs/EventAPI.md#eventlist) | **Get** /v1/events | List events matching any selected filters.
*EventAPI* | [**EventView**](docs/EventAPI.md#eventview) | **Get** /v1/events/{id} | View a specific event.
*KubernetesIntegrationAPI* | [**KubernetesIntegrationList**](docs/KubernetesIntegrationAPI.md#kubernetesintegrationlist) | **Get** /v1/kubernetes_integrations | List Kubernetes integrations matching any selected filters.
*KubernetesIntegrationAPI* | [**KubernetesIntegrationView**](docs/KubernetesIntegrationAPI.md#kubernetesintegrationview) | **Get** /v1/kubernetes_integrations/{id} | View a specific Kubernetes integration.
*OrganizationAPI* | [**OrganizationList**](docs/OrganizationAPI.md#organizationlist) | **Get** /v1/organizations | List organization matching any selected filters.
*OrganizationAPI* | [**OrganizationView**](docs/OrganizationAPI.md#organizationview) | **Get** /v1/organizations/{id} | View a specific organization.
*PipelineAPI* | [**PipelineList**](docs/PipelineAPI.md#pipelinelist) | **Get** /v1/pipelines | List pipelines matching any selected filters.
*PipelineAPI* | [**PipelineView**](docs/PipelineAPI.md#pipelineview) | **Get** /v1/pipelines/{id} | View a specific Pipeline.
*ProjectAPI* | [**ProjectCreate**](docs/ProjectAPI.md#projectcreate) | **Post** /v1/projects | Creates new project.
*ProjectAPI* | [**ProjectDelete**](docs/ProjectAPI.md#projectdelete) | **Delete** /v1/projects/{id} | Delete a specific project.
*ProjectAPI* | [**ProjectEditBuildSettings**](docs/ProjectAPI.md#projecteditbuildsettings) | **Patch** /v1/projects/{id}/build-settings | Edit the build settings of a project.
*ProjectAPI* | [**ProjectEditSettings**](docs/ProjectAPI.md#projecteditsettings) | **Patch** /v1/projects/{id}/settings | Edit a project.
*ProjectAPI* | [**ProjectList**](docs/ProjectAPI.md#projectlist) | **Get** /v1/projects | List projects matching any selected filters.
*ProjectAPI* | [**ProjectView**](docs/ProjectAPI.md#projectview) | **Get** /v1/projects/{id} | View a specific project.
*ProjectVariableAPI* | [**ProjectVariableCreate**](docs/ProjectVariableAPI.md#projectvariablecreate) | **Post** /v1/project_variables | Create a specific project variable.
*ProjectVariableAPI* | [**ProjectVariableDelete**](docs/ProjectVariableAPI.md#projectvariabledelete) | **Delete** /v1/project_variables/{id} | Delete a specific project variable.
*ProjectVariableAPI* | [**ProjectVariableEdit**](docs/ProjectVariableAPI.md#projectvariableedit) | **Patch** /v1/project_variables/{id} | Edit a specific project variable.
*ProjectVariableAPI* | [**ProjectVariableList**](docs/ProjectVariableAPI.md#projectvariablelist) | **Get** /v1/project_variables | List project variables matching any selected filters.
*ProjectVariableAPI* | [**ProjectVariableView**](docs/ProjectVariableAPI.md#projectvariableview) | **Get** /v1/project_variables/{id} | View a specific project variable.
*RegistryIntegrationAPI* | [**RegistryIntegrationList**](docs/RegistryIntegrationAPI.md#registryintegrationlist) | **Get** /v1/registry_integrations | List Registry integrations matching any selected filters.
*RegistryIntegrationAPI* | [**RegistryIntegrationView**](docs/RegistryIntegrationAPI.md#registryintegrationview) | **Get** /v1/registry_integrations/{id} | View a specific Registry integration.
*TemplateAPI* | [**TemplateDefinition**](docs/TemplateAPI.md#templatedefinition) | **Get** /v1/templates/{id}/definition | View the environment definition.
*TemplateAPI* | [**TemplateList**](docs/TemplateAPI.md#templatelist) | **Get** /v1/templates | List templates matching any selected filters.
*TemplateAPI* | [**TemplateValidate**](docs/TemplateAPI.md#templatevalidate) | **Post** /v1/templates/validate | Validates a given template from an external source.
*TemplateAPI* | [**TemplateView**](docs/TemplateAPI.md#templateview) | **Get** /v1/templates/{id} | View a specific template.
*TemplatesRepositoryAPI* | [**TemplatesRepositoryList**](docs/TemplatesRepositoryAPI.md#templatesrepositorylist) | **Get** /v1/templates_repositories | List templates repositories matching any selected filters.
*TemplatesRepositoryAPI* | [**TemplatesRepositoryView**](docs/TemplatesRepositoryAPI.md#templatesrepositoryview) | **Get** /v1/templates_repositories/{id} | View a specific templates repository.
*WorkflowAPI* | [**WorkflowList**](docs/WorkflowAPI.md#workflowlist) | **Get** /v1/workflows | List workflows matching any selected filters.
*WorkflowAPI* | [**WorkflowView**](docs/WorkflowAPI.md#workflowview) | **Get** /v1/workflows/{id} | View a specific Workflow.


## Documentation For Models

 - [BuildSettingsItem](docs/BuildSettingsItem.md)
 - [ClusterKubeConfigRead](docs/ClusterKubeConfigRead.md)
 - [ClusterWrapperKubeConfigRead](docs/ClusterWrapperKubeConfigRead.md)
 - [ComponentCollection](docs/ComponentCollection.md)
 - [ComponentConfigItem](docs/ComponentConfigItem.md)
 - [ComponentConfigItemConfig](docs/ComponentConfigItemConfig.md)
 - [ComponentEndpointCollection](docs/ComponentEndpointCollection.md)
 - [ComponentEndpointItem](docs/ComponentEndpointItem.md)
 - [ComponentGitCollection](docs/ComponentGitCollection.md)
 - [ComponentGitItem](docs/ComponentGitItem.md)
 - [ComponentItem](docs/ComponentItem.md)
 - [ComponentProfileItem](docs/ComponentProfileItem.md)
 - [ComponentResourceItem](docs/ComponentResourceItem.md)
 - [ContainerConfigItem](docs/ContainerConfigItem.md)
 - [ContainerConfigItemProfile](docs/ContainerConfigItemProfile.md)
 - [ContextKubeConfigRead](docs/ContextKubeConfigRead.md)
 - [ContextWrapperKubeConfigRead](docs/ContextWrapperKubeConfigRead.md)
 - [Edit](docs/Edit.md)
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
 - [EmbeddedProjectVariableCollection](docs/EmbeddedProjectVariableCollection.md)
 - [EmbeddedRegistryIntegrationCollection](docs/EmbeddedRegistryIntegrationCollection.md)
 - [EmbeddedTemplateCollection](docs/EmbeddedTemplateCollection.md)
 - [EmbeddedTemplatesRepositoryCollection](docs/EmbeddedTemplatesRepositoryCollection.md)
 - [EmbeddedWorkflowCollection](docs/EmbeddedWorkflowCollection.md)
 - [EnvironmentCloneAction](docs/EnvironmentCloneAction.md)
 - [EnvironmentCollection](docs/EnvironmentCollection.md)
 - [EnvironmentCreateAction](docs/EnvironmentCreateAction.md)
 - [EnvironmentCreateActionGenesis](docs/EnvironmentCreateActionGenesis.md)
 - [EnvironmentEditBuildSettingsAction](docs/EnvironmentEditBuildSettingsAction.md)
 - [EnvironmentEditComponentsAction](docs/EnvironmentEditComponentsAction.md)
 - [EnvironmentEditComponentsActionFilter](docs/EnvironmentEditComponentsActionFilter.md)
 - [EnvironmentEditConfiguration](docs/EnvironmentEditConfiguration.md)
 - [EnvironmentEditConfigurationConfiguration](docs/EnvironmentEditConfigurationConfiguration.md)
 - [EnvironmentEditSettings](docs/EnvironmentEditSettings.md)
 - [EnvironmentEditSettingsEdit](docs/EnvironmentEditSettingsEdit.md)
 - [EnvironmentItem](docs/EnvironmentItem.md)
 - [EnvironmentKubeConfigKubeConfigRead](docs/EnvironmentKubeConfigKubeConfigRead.md)
 - [EnvironmentPartialAction](docs/EnvironmentPartialAction.md)
 - [EnvironmentPartialDeployAction](docs/EnvironmentPartialDeployAction.md)
 - [EnvironmentPartialStartAction](docs/EnvironmentPartialStartAction.md)
 - [EnvironmentVariableCollection](docs/EnvironmentVariableCollection.md)
 - [EnvironmentVariableCreateAction](docs/EnvironmentVariableCreateAction.md)
 - [EnvironmentVariableEditAction](docs/EnvironmentVariableEditAction.md)
 - [EnvironmentVariableItem](docs/EnvironmentVariableItem.md)
 - [Ephemeral](docs/Ephemeral.md)
 - [EventCollection](docs/EventCollection.md)
 - [EventItem](docs/EventItem.md)
 - [ExtendedResourceConfigItem](docs/ExtendedResourceConfigItem.md)
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
 - [PaginatedOrganizationCollection](docs/PaginatedOrganizationCollection.md)
 - [PaginatedPipelineCollection](docs/PaginatedPipelineCollection.md)
 - [PaginatedProjectCollection](docs/PaginatedProjectCollection.md)
 - [PaginatedProjectVariableCollection](docs/PaginatedProjectVariableCollection.md)
 - [PaginatedRegistryIntegrationCollection](docs/PaginatedRegistryIntegrationCollection.md)
 - [PaginatedTemplateCollection](docs/PaginatedTemplateCollection.md)
 - [PaginatedTemplatesRepositoryCollection](docs/PaginatedTemplatesRepositoryCollection.md)
 - [PaginatedWorkflowCollection](docs/PaginatedWorkflowCollection.md)
 - [PipelineCollection](docs/PipelineCollection.md)
 - [PipelineItem](docs/PipelineItem.md)
 - [Primary](docs/Primary.md)
 - [ProblemGeneric](docs/ProblemGeneric.md)
 - [ProblemViolation](docs/ProblemViolation.md)
 - [ProfileItem](docs/ProfileItem.md)
 - [ProjectCollection](docs/ProjectCollection.md)
 - [ProjectCreateAction](docs/ProjectCreateAction.md)
 - [ProjectEditBuildSettingsAction](docs/ProjectEditBuildSettingsAction.md)
 - [ProjectEditSettingsAction](docs/ProjectEditSettingsAction.md)
 - [ProjectItem](docs/ProjectItem.md)
 - [ProjectVariableCollection](docs/ProjectVariableCollection.md)
 - [ProjectVariableCreateAction](docs/ProjectVariableCreateAction.md)
 - [ProjectVariableEditAction](docs/ProjectVariableEditAction.md)
 - [ProjectVariableItem](docs/ProjectVariableItem.md)
 - [RegistryIntegrationCollection](docs/RegistryIntegrationCollection.md)
 - [RegistryIntegrationItem](docs/RegistryIntegrationItem.md)
 - [ResourceListItem](docs/ResourceListItem.md)
 - [ResourceRequirementItem](docs/ResourceRequirementItem.md)
 - [SimpleResourceConfigItem](docs/SimpleResourceConfigItem.md)
 - [StageItem](docs/StageItem.md)
 - [SyncPathItem](docs/SyncPathItem.md)
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
 - [WorkflowCollection](docs/WorkflowCollection.md)
 - [WorkflowItem](docs/WorkflowItem.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Auth-Token and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"X-Auth-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### JWT

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


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

