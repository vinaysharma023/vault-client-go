# GcpKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyAlgorithm** | Pointer to **string** | Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\&quot; | [optional] [default to "KEY_ALG_RSA_2048"]
**KeyType** | Pointer to **string** | Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\&quot; | [optional] [default to "TYPE_GOOGLE_CREDENTIALS_FILE"]
**Ttl** | Pointer to **int32** | Lifetime of the service account key | [optional] 

## Methods

### NewGcpKeyRequest

`func NewGcpKeyRequest() *GcpKeyRequest`

NewGcpKeyRequest instantiates a new GcpKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpKeyRequestWithDefaults

`func NewGcpKeyRequestWithDefaults() *GcpKeyRequest`

NewGcpKeyRequestWithDefaults instantiates a new GcpKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyAlgorithm

`func (o *GcpKeyRequest) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *GcpKeyRequest) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *GcpKeyRequest) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *GcpKeyRequest) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetKeyType

`func (o *GcpKeyRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GcpKeyRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GcpKeyRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *GcpKeyRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetTtl

`func (o *GcpKeyRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GcpKeyRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GcpKeyRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GcpKeyRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

