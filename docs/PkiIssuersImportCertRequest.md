# PkiIssuersImportCertRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PemBundle** | Pointer to **string** | PEM-format, concatenated unencrypted secret-key (optional) and certificates. | [optional] 



## Methods


### NewPkiIssuersImportCertRequest

`func NewPkiIssuersImportCertRequest() *PkiIssuersImportCertRequest`

NewPkiIssuersImportCertRequest instantiates a new PkiIssuersImportCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuersImportCertRequestWithDefaults

`func NewPkiIssuersImportCertRequestWithDefaults() *PkiIssuersImportCertRequest`

NewPkiIssuersImportCertRequestWithDefaults instantiates a new PkiIssuersImportCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPemBundle

`func (o *PkiIssuersImportCertRequest) GetPemBundle() string`

GetPemBundle returns the PemBundle field if non-nil, zero value otherwise.

### GetPemBundleOk

`func (o *PkiIssuersImportCertRequest) GetPemBundleOk() (*string, bool)`

GetPemBundleOk returns a tuple with the PemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemBundle

`func (o *PkiIssuersImportCertRequest) SetPemBundle(v string)`

SetPemBundle sets PemBundle field to given value.


### HasPemBundle

`func (o *PkiIssuersImportCertRequest) HasPemBundle() bool`

HasPemBundle returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

