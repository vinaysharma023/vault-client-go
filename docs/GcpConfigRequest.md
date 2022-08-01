# GcpConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | Google credentials JSON that Vault will use to verify users against GCP APIs. If not specified, will use application default credentials | [optional] 
**CustomEndpoint** | Pointer to **map[string]interface{}** | Specifies overrides for various Google API Service Endpoints used in requests. | [optional] 
**GceAlias** | Pointer to **string** | Indicates what value to use when generating an alias for GCE authentications. | [optional] [default to "role_id"]
**GceMetadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#39;default&#39;, includes: instance_creation_timestamp, instance_id, instance_name, project_id, project_number, role, service_account_id, service_account_email, zone. Not editing this field means the &#39;default&#39; fields are included. Explicitly setting this field to empty overrides the &#39;default&#39; and means no metadata will be included. If not using &#39;default&#39;, explicit fields must be sent like: &#39;field1,field2&#39;. | [optional] [default to ["default"]]
**GoogleCertsEndpoint** | Pointer to **string** | Deprecated. This field does nothing and be removed in a future release | [optional] 
**IamAlias** | Pointer to **string** | Indicates what value to use when generating an alias for IAM authentications. | [optional] [default to "role_id"]
**IamMetadata** | Pointer to **[]string** | The metadata to include on the aliases and audit logs generated by this plugin. When set to &#39;default&#39;, includes: project_id, role, service_account_id, service_account_email. Not editing this field means the &#39;default&#39; fields are included. Explicitly setting this field to empty overrides the &#39;default&#39; and means no metadata will be included. If not using &#39;default&#39;, explicit fields must be sent like: &#39;field1,field2&#39;. | [optional] [default to ["default"]]

## Methods

### NewGcpConfigRequest

`func NewGcpConfigRequest() *GcpConfigRequest`

NewGcpConfigRequest instantiates a new GcpConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpConfigRequestWithDefaults

`func NewGcpConfigRequestWithDefaults() *GcpConfigRequest`

NewGcpConfigRequestWithDefaults instantiates a new GcpConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GcpConfigRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GcpConfigRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GcpConfigRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GcpConfigRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCustomEndpoint

`func (o *GcpConfigRequest) GetCustomEndpoint() map[string]interface{}`

GetCustomEndpoint returns the CustomEndpoint field if non-nil, zero value otherwise.

### GetCustomEndpointOk

`func (o *GcpConfigRequest) GetCustomEndpointOk() (*map[string]interface{}, bool)`

GetCustomEndpointOk returns a tuple with the CustomEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEndpoint

`func (o *GcpConfigRequest) SetCustomEndpoint(v map[string]interface{})`

SetCustomEndpoint sets CustomEndpoint field to given value.

### HasCustomEndpoint

`func (o *GcpConfigRequest) HasCustomEndpoint() bool`

HasCustomEndpoint returns a boolean if a field has been set.

### GetGceAlias

`func (o *GcpConfigRequest) GetGceAlias() string`

GetGceAlias returns the GceAlias field if non-nil, zero value otherwise.

### GetGceAliasOk

`func (o *GcpConfigRequest) GetGceAliasOk() (*string, bool)`

GetGceAliasOk returns a tuple with the GceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceAlias

`func (o *GcpConfigRequest) SetGceAlias(v string)`

SetGceAlias sets GceAlias field to given value.

### HasGceAlias

`func (o *GcpConfigRequest) HasGceAlias() bool`

HasGceAlias returns a boolean if a field has been set.

### GetGceMetadata

`func (o *GcpConfigRequest) GetGceMetadata() []string`

GetGceMetadata returns the GceMetadata field if non-nil, zero value otherwise.

### GetGceMetadataOk

`func (o *GcpConfigRequest) GetGceMetadataOk() (*[]string, bool)`

GetGceMetadataOk returns a tuple with the GceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceMetadata

`func (o *GcpConfigRequest) SetGceMetadata(v []string)`

SetGceMetadata sets GceMetadata field to given value.

### HasGceMetadata

`func (o *GcpConfigRequest) HasGceMetadata() bool`

HasGceMetadata returns a boolean if a field has been set.

### GetGoogleCertsEndpoint

`func (o *GcpConfigRequest) GetGoogleCertsEndpoint() string`

GetGoogleCertsEndpoint returns the GoogleCertsEndpoint field if non-nil, zero value otherwise.

### GetGoogleCertsEndpointOk

`func (o *GcpConfigRequest) GetGoogleCertsEndpointOk() (*string, bool)`

GetGoogleCertsEndpointOk returns a tuple with the GoogleCertsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCertsEndpoint

`func (o *GcpConfigRequest) SetGoogleCertsEndpoint(v string)`

SetGoogleCertsEndpoint sets GoogleCertsEndpoint field to given value.

### HasGoogleCertsEndpoint

`func (o *GcpConfigRequest) HasGoogleCertsEndpoint() bool`

HasGoogleCertsEndpoint returns a boolean if a field has been set.

### GetIamAlias

`func (o *GcpConfigRequest) GetIamAlias() string`

GetIamAlias returns the IamAlias field if non-nil, zero value otherwise.

### GetIamAliasOk

`func (o *GcpConfigRequest) GetIamAliasOk() (*string, bool)`

GetIamAliasOk returns a tuple with the IamAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAlias

`func (o *GcpConfigRequest) SetIamAlias(v string)`

SetIamAlias sets IamAlias field to given value.

### HasIamAlias

`func (o *GcpConfigRequest) HasIamAlias() bool`

HasIamAlias returns a boolean if a field has been set.

### GetIamMetadata

`func (o *GcpConfigRequest) GetIamMetadata() []string`

GetIamMetadata returns the IamMetadata field if non-nil, zero value otherwise.

### GetIamMetadataOk

`func (o *GcpConfigRequest) GetIamMetadataOk() (*[]string, bool)`

GetIamMetadataOk returns a tuple with the IamMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamMetadata

`func (o *GcpConfigRequest) SetIamMetadata(v []string)`

SetIamMetadata sets IamMetadata field to given value.

### HasIamMetadata

`func (o *GcpConfigRequest) HasIamMetadata() bool`

HasIamMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

