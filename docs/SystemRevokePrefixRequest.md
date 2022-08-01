# SystemRevokePrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sync** | Pointer to **bool** | Whether or not to perform the revocation synchronously | [optional] [default to true]

## Methods

### NewSystemRevokePrefixRequest

`func NewSystemRevokePrefixRequest() *SystemRevokePrefixRequest`

NewSystemRevokePrefixRequest instantiates a new SystemRevokePrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRevokePrefixRequestWithDefaults

`func NewSystemRevokePrefixRequestWithDefaults() *SystemRevokePrefixRequest`

NewSystemRevokePrefixRequestWithDefaults instantiates a new SystemRevokePrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSync

`func (o *SystemRevokePrefixRequest) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *SystemRevokePrefixRequest) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *SystemRevokePrefixRequest) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *SystemRevokePrefixRequest) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

