# KubernetesRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasNameSource** | Pointer to **string** | Source to use when deriving the Alias name. valid choices: \&quot;serviceaccount_uid\&quot; : &lt;token.uid&gt; e.g. 474b11b5-0f20-4f9d-8ca5-65715ab325e0 (most secure choice) \&quot;serviceaccount_name\&quot; : &lt;namespace&gt;/&lt;serviceaccount&gt; e.g. vault/vault-agent default: \&quot;serviceaccount_uid\&quot; | [optional] [default to "serviceaccount_uid"]
**Audience** | Pointer to **string** | Optional Audience claim to verify in the jwt. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**BoundServiceAccountNames** | Pointer to **[]string** | List of service account names able to access this role. If set to \&quot;*\&quot; all names are allowed. | [optional] 
**BoundServiceAccountNamespaces** | Pointer to **[]string** | List of namespaces allowed to access this role. If set to \&quot;*\&quot; all namespaces are allowed. | [optional] 
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**NumUses** | Pointer to **int32** | Use \&quot;token_num_uses\&quot; instead. If this and \&quot;token_num_uses\&quot; are both specified, only \&quot;token_num_uses\&quot; will be used. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 

## Methods

### NewKubernetesRoleRequest

`func NewKubernetesRoleRequest() *KubernetesRoleRequest`

NewKubernetesRoleRequest instantiates a new KubernetesRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRoleRequestWithDefaults

`func NewKubernetesRoleRequestWithDefaults() *KubernetesRoleRequest`

NewKubernetesRoleRequestWithDefaults instantiates a new KubernetesRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasNameSource

`func (o *KubernetesRoleRequest) GetAliasNameSource() string`

GetAliasNameSource returns the AliasNameSource field if non-nil, zero value otherwise.

### GetAliasNameSourceOk

`func (o *KubernetesRoleRequest) GetAliasNameSourceOk() (*string, bool)`

GetAliasNameSourceOk returns a tuple with the AliasNameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasNameSource

`func (o *KubernetesRoleRequest) SetAliasNameSource(v string)`

SetAliasNameSource sets AliasNameSource field to given value.

### HasAliasNameSource

`func (o *KubernetesRoleRequest) HasAliasNameSource() bool`

HasAliasNameSource returns a boolean if a field has been set.

### GetAudience

`func (o *KubernetesRoleRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *KubernetesRoleRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *KubernetesRoleRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *KubernetesRoleRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundCidrs

`func (o *KubernetesRoleRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *KubernetesRoleRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *KubernetesRoleRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *KubernetesRoleRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetBoundServiceAccountNames

`func (o *KubernetesRoleRequest) GetBoundServiceAccountNames() []string`

GetBoundServiceAccountNames returns the BoundServiceAccountNames field if non-nil, zero value otherwise.

### GetBoundServiceAccountNamesOk

`func (o *KubernetesRoleRequest) GetBoundServiceAccountNamesOk() (*[]string, bool)`

GetBoundServiceAccountNamesOk returns a tuple with the BoundServiceAccountNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccountNames

`func (o *KubernetesRoleRequest) SetBoundServiceAccountNames(v []string)`

SetBoundServiceAccountNames sets BoundServiceAccountNames field to given value.

### HasBoundServiceAccountNames

`func (o *KubernetesRoleRequest) HasBoundServiceAccountNames() bool`

HasBoundServiceAccountNames returns a boolean if a field has been set.

### GetBoundServiceAccountNamespaces

`func (o *KubernetesRoleRequest) GetBoundServiceAccountNamespaces() []string`

GetBoundServiceAccountNamespaces returns the BoundServiceAccountNamespaces field if non-nil, zero value otherwise.

### GetBoundServiceAccountNamespacesOk

`func (o *KubernetesRoleRequest) GetBoundServiceAccountNamespacesOk() (*[]string, bool)`

GetBoundServiceAccountNamespacesOk returns a tuple with the BoundServiceAccountNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccountNamespaces

`func (o *KubernetesRoleRequest) SetBoundServiceAccountNamespaces(v []string)`

SetBoundServiceAccountNamespaces sets BoundServiceAccountNamespaces field to given value.

### HasBoundServiceAccountNamespaces

`func (o *KubernetesRoleRequest) HasBoundServiceAccountNamespaces() bool`

HasBoundServiceAccountNamespaces returns a boolean if a field has been set.

### GetMaxTtl

`func (o *KubernetesRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *KubernetesRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *KubernetesRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *KubernetesRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetNumUses

`func (o *KubernetesRoleRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *KubernetesRoleRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *KubernetesRoleRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.

### HasNumUses

`func (o *KubernetesRoleRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.

### GetPeriod

`func (o *KubernetesRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *KubernetesRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *KubernetesRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *KubernetesRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *KubernetesRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *KubernetesRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *KubernetesRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *KubernetesRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *KubernetesRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *KubernetesRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *KubernetesRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *KubernetesRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *KubernetesRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *KubernetesRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *KubernetesRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *KubernetesRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *KubernetesRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *KubernetesRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *KubernetesRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *KubernetesRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *KubernetesRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *KubernetesRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *KubernetesRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *KubernetesRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *KubernetesRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *KubernetesRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *KubernetesRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *KubernetesRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *KubernetesRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *KubernetesRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *KubernetesRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *KubernetesRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *KubernetesRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *KubernetesRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *KubernetesRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *KubernetesRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *KubernetesRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *KubernetesRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *KubernetesRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *KubernetesRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *KubernetesRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *KubernetesRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *KubernetesRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *KubernetesRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *KubernetesRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KubernetesRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KubernetesRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KubernetesRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

