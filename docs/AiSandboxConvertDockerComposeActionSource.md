# AiSandboxConvertDockerComposeActionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] [default to "docker-compose.yaml"]
**DockerComposePath** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "string"]
**Url** | **string** |  | 
**Ref** | **string** |  | 
**OrganizationId** | **string** |  | 
**DockerComposeYaml** | **string** |  | 
**DotEnvFileContent** | Pointer to **string** |  | [optional] 

## Methods

### NewAiSandboxConvertDockerComposeActionSource

`func NewAiSandboxConvertDockerComposeActionSource(dockerComposePath string, url string, ref string, organizationId string, dockerComposeYaml string, ) *AiSandboxConvertDockerComposeActionSource`

NewAiSandboxConvertDockerComposeActionSource instantiates a new AiSandboxConvertDockerComposeActionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSandboxConvertDockerComposeActionSourceWithDefaults

`func NewAiSandboxConvertDockerComposeActionSourceWithDefaults() *AiSandboxConvertDockerComposeActionSource`

NewAiSandboxConvertDockerComposeActionSourceWithDefaults instantiates a new AiSandboxConvertDockerComposeActionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AiSandboxConvertDockerComposeActionSource) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AiSandboxConvertDockerComposeActionSource) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AiSandboxConvertDockerComposeActionSource) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDockerComposePath

`func (o *AiSandboxConvertDockerComposeActionSource) GetDockerComposePath() string`

GetDockerComposePath returns the DockerComposePath field if non-nil, zero value otherwise.

### GetDockerComposePathOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetDockerComposePathOk() (*string, bool)`

GetDockerComposePathOk returns a tuple with the DockerComposePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerComposePath

`func (o *AiSandboxConvertDockerComposeActionSource) SetDockerComposePath(v string)`

SetDockerComposePath sets DockerComposePath field to given value.


### GetType

`func (o *AiSandboxConvertDockerComposeActionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiSandboxConvertDockerComposeActionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiSandboxConvertDockerComposeActionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AiSandboxConvertDockerComposeActionSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiSandboxConvertDockerComposeActionSource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRef

`func (o *AiSandboxConvertDockerComposeActionSource) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AiSandboxConvertDockerComposeActionSource) SetRef(v string)`

SetRef sets Ref field to given value.


### GetOrganizationId

`func (o *AiSandboxConvertDockerComposeActionSource) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AiSandboxConvertDockerComposeActionSource) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetDockerComposeYaml

`func (o *AiSandboxConvertDockerComposeActionSource) GetDockerComposeYaml() string`

GetDockerComposeYaml returns the DockerComposeYaml field if non-nil, zero value otherwise.

### GetDockerComposeYamlOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetDockerComposeYamlOk() (*string, bool)`

GetDockerComposeYamlOk returns a tuple with the DockerComposeYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerComposeYaml

`func (o *AiSandboxConvertDockerComposeActionSource) SetDockerComposeYaml(v string)`

SetDockerComposeYaml sets DockerComposeYaml field to given value.


### GetDotEnvFileContent

`func (o *AiSandboxConvertDockerComposeActionSource) GetDotEnvFileContent() string`

GetDotEnvFileContent returns the DotEnvFileContent field if non-nil, zero value otherwise.

### GetDotEnvFileContentOk

`func (o *AiSandboxConvertDockerComposeActionSource) GetDotEnvFileContentOk() (*string, bool)`

GetDotEnvFileContentOk returns a tuple with the DotEnvFileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotEnvFileContent

`func (o *AiSandboxConvertDockerComposeActionSource) SetDotEnvFileContent(v string)`

SetDotEnvFileContent sets DotEnvFileContent field to given value.

### HasDotEnvFileContent

`func (o *AiSandboxConvertDockerComposeActionSource) HasDotEnvFileContent() bool`

HasDotEnvFileContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


