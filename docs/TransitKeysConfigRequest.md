# TransitKeysConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPlaintextBackup** | Pointer to **bool** | Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled. | [optional] 
**AutoRotatePeriod** | Pointer to **int32** | Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the key. | [optional] 
**DeletionAllowed** | Pointer to **bool** | Whether to allow deletion of the key | [optional] 
**Exportable** | Pointer to **bool** | Enables export of the key. Once set, this cannot be disabled. | [optional] 
**MinDecryptionVersion** | Pointer to **int32** | If set, the minimum version of the key allowed to be decrypted. For signing keys, the minimum version allowed to be used for verification. | [optional] 
**MinEncryptionVersion** | Pointer to **int32** | If set, the minimum version of the key allowed to be used for encryption; or for signing keys, to be used for signing. If set to zero, only the latest version of the key is allowed. | [optional] 

## Methods

### NewTransitKeysConfigRequest

`func NewTransitKeysConfigRequest() *TransitKeysConfigRequest`

NewTransitKeysConfigRequest instantiates a new TransitKeysConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitKeysConfigRequestWithDefaults

`func NewTransitKeysConfigRequestWithDefaults() *TransitKeysConfigRequest`

NewTransitKeysConfigRequestWithDefaults instantiates a new TransitKeysConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPlaintextBackup

`func (o *TransitKeysConfigRequest) GetAllowPlaintextBackup() bool`

GetAllowPlaintextBackup returns the AllowPlaintextBackup field if non-nil, zero value otherwise.

### GetAllowPlaintextBackupOk

`func (o *TransitKeysConfigRequest) GetAllowPlaintextBackupOk() (*bool, bool)`

GetAllowPlaintextBackupOk returns a tuple with the AllowPlaintextBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaintextBackup

`func (o *TransitKeysConfigRequest) SetAllowPlaintextBackup(v bool)`

SetAllowPlaintextBackup sets AllowPlaintextBackup field to given value.

### HasAllowPlaintextBackup

`func (o *TransitKeysConfigRequest) HasAllowPlaintextBackup() bool`

HasAllowPlaintextBackup returns a boolean if a field has been set.

### GetAutoRotatePeriod

`func (o *TransitKeysConfigRequest) GetAutoRotatePeriod() int32`

GetAutoRotatePeriod returns the AutoRotatePeriod field if non-nil, zero value otherwise.

### GetAutoRotatePeriodOk

`func (o *TransitKeysConfigRequest) GetAutoRotatePeriodOk() (*int32, bool)`

GetAutoRotatePeriodOk returns a tuple with the AutoRotatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotatePeriod

`func (o *TransitKeysConfigRequest) SetAutoRotatePeriod(v int32)`

SetAutoRotatePeriod sets AutoRotatePeriod field to given value.

### HasAutoRotatePeriod

`func (o *TransitKeysConfigRequest) HasAutoRotatePeriod() bool`

HasAutoRotatePeriod returns a boolean if a field has been set.

### GetDeletionAllowed

`func (o *TransitKeysConfigRequest) GetDeletionAllowed() bool`

GetDeletionAllowed returns the DeletionAllowed field if non-nil, zero value otherwise.

### GetDeletionAllowedOk

`func (o *TransitKeysConfigRequest) GetDeletionAllowedOk() (*bool, bool)`

GetDeletionAllowedOk returns a tuple with the DeletionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionAllowed

`func (o *TransitKeysConfigRequest) SetDeletionAllowed(v bool)`

SetDeletionAllowed sets DeletionAllowed field to given value.

### HasDeletionAllowed

`func (o *TransitKeysConfigRequest) HasDeletionAllowed() bool`

HasDeletionAllowed returns a boolean if a field has been set.

### GetExportable

`func (o *TransitKeysConfigRequest) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *TransitKeysConfigRequest) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *TransitKeysConfigRequest) SetExportable(v bool)`

SetExportable sets Exportable field to given value.

### HasExportable

`func (o *TransitKeysConfigRequest) HasExportable() bool`

HasExportable returns a boolean if a field has been set.

### GetMinDecryptionVersion

`func (o *TransitKeysConfigRequest) GetMinDecryptionVersion() int32`

GetMinDecryptionVersion returns the MinDecryptionVersion field if non-nil, zero value otherwise.

### GetMinDecryptionVersionOk

`func (o *TransitKeysConfigRequest) GetMinDecryptionVersionOk() (*int32, bool)`

GetMinDecryptionVersionOk returns a tuple with the MinDecryptionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDecryptionVersion

`func (o *TransitKeysConfigRequest) SetMinDecryptionVersion(v int32)`

SetMinDecryptionVersion sets MinDecryptionVersion field to given value.

### HasMinDecryptionVersion

`func (o *TransitKeysConfigRequest) HasMinDecryptionVersion() bool`

HasMinDecryptionVersion returns a boolean if a field has been set.

### GetMinEncryptionVersion

`func (o *TransitKeysConfigRequest) GetMinEncryptionVersion() int32`

GetMinEncryptionVersion returns the MinEncryptionVersion field if non-nil, zero value otherwise.

### GetMinEncryptionVersionOk

`func (o *TransitKeysConfigRequest) GetMinEncryptionVersionOk() (*int32, bool)`

GetMinEncryptionVersionOk returns a tuple with the MinEncryptionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEncryptionVersion

`func (o *TransitKeysConfigRequest) SetMinEncryptionVersion(v int32)`

SetMinEncryptionVersion sets MinEncryptionVersion field to given value.

### HasMinEncryptionVersion

`func (o *TransitKeysConfigRequest) HasMinEncryptionVersion() bool`

HasMinEncryptionVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

