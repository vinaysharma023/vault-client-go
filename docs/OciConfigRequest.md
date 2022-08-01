# OciConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomeTenancyId** | Pointer to **string** | The tenancy id of the account. | [optional] 

## Methods

### NewOciConfigRequest

`func NewOciConfigRequest() *OciConfigRequest`

NewOciConfigRequest instantiates a new OciConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciConfigRequestWithDefaults

`func NewOciConfigRequestWithDefaults() *OciConfigRequest`

NewOciConfigRequestWithDefaults instantiates a new OciConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHomeTenancyId

`func (o *OciConfigRequest) GetHomeTenancyId() string`

GetHomeTenancyId returns the HomeTenancyId field if non-nil, zero value otherwise.

### GetHomeTenancyIdOk

`func (o *OciConfigRequest) GetHomeTenancyIdOk() (*string, bool)`

GetHomeTenancyIdOk returns a tuple with the HomeTenancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTenancyId

`func (o *OciConfigRequest) SetHomeTenancyId(v string)`

SetHomeTenancyId sets HomeTenancyId field to given value.

### HasHomeTenancyId

`func (o *OciConfigRequest) HasHomeTenancyId() bool`

HasHomeTenancyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

