// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHWriteRoleRequest struct for SSHWriteRoleRequest
type SSHWriteRoleRequest struct {
	// [Required for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Admin user at remote host. The shared key being registered should be for this user and should have root privileges. Everytime a dynamic credential is being generated for other users, Vault uses this admin username to login to remote host and install the generated credential for the other user.
	AdminUser string `json:"admin_user"`

	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512, default, or the empty string.
	AlgorithmSigner string `json:"algorithm_signer"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use the base domains listed in \"allowed_domains\", e.g. \"example.com\". This is a separate option as in some cases this can be considered a security threat.
	AllowBareDomains bool `json:"allow_bare_domains"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates bool `json:"allow_host_certificates"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use subdomains of those listed in \"allowed_domains\".
	AllowSubdomains bool `json:"allow_subdomains"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates bool `json:"allow_user_certificates"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If true, users can override the key ID for a signed certificate with the \"key_id\" field. When false, the key ID will always be the token display name. The key ID is logged by the SSH server and can be useful for auditing.
	AllowUserKeyIds bool `json:"allow_user_key_ids"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] A comma-separated list of critical options that certificates can have when signed. To allow any critical options, set this to an empty string.
	AllowedCriticalOptions string `json:"allowed_critical_options"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If this option is not specified, client can request for a signed certificate for any valid host. If only certain domains are allowed, then this list enforces it.
	AllowedDomains string `json:"allowed_domains"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate bool `json:"allowed_domains_template"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] A comma-separated list of extensions that certificates can have when signed. An empty list means that no extension overrides are allowed by an end-user; explicitly specify '*' to allow any extensions to be set.
	AllowedExtensions string `json:"allowed_extensions"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, allows the enforcement of key types and minimum key sizes to be signed.
	AllowedUserKeyLengths map[string]interface{} `json:"allowed_user_key_lengths"`

	// [Optional for all types] [Works differently for CA type] If this option is not specified, or is '*', client can request a credential for any valid user at the remote host, including the admin user. If only certain usernames are to be allowed, then this list enforces it. If this field is set, then credentials can only be created for default_user and usernames present in this list. Setting this option will enable all the users with access to this role to fetch credentials for all other usernames in this list. Use with caution. N.B.: with the CA type, an empty list means that no users are allowed; explicitly specify '*' to allow any user.
	AllowedUsers string `json:"allowed_users"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Allowed users can be specified using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate bool `json:"allowed_users_template"`

	// [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks for which the role is applicable for. CIDR blocks can belong to more than one role.
	CidrList string `json:"cidr_list"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] Critical options certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \"allowed_critical_options\". Defaults to none.
	DefaultCriticalOptions map[string]interface{} `json:"default_critical_options"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] Extensions certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \"allowed_extensions\". Defaults to none.
	DefaultExtensions map[string]interface{} `json:"default_extensions"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Default extension values can be specified using identity template policies. Non-templated extension values are also permitted.
	DefaultExtensionsTemplate bool `json:"default_extensions_template"`

	// [Required for Dynamic type] [Required for OTP type] [Optional for CA type] Default username for which a credential will be generated. When the endpoint 'creds/' is used without a username, this value will be used as default username.
	DefaultUser string `json:"default_user"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Default user can be specified using identity template policies. Non-templated users are also permitted.
	DefaultUserTemplate bool `json:"default_user_template"`

	// [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks. IP addresses belonging to these blocks are not accepted by the role. This is particularly useful when big CIDR blocks are being used by the role and certain parts of it needs to be kept out.
	ExcludeCidrList string `json:"exclude_cidr_list"`

	// [Optional for Dynamic type] [Not-applicable for OTP type] [Not applicable for CA type] Script used to install and uninstall public keys in the target machine. The inbuilt default install script will be for Linux hosts. For sample script, refer the project documentation website.
	InstallScript string `json:"install_script"`

	// [Required for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Name of the registered key in Vault. Before creating the role, use the 'keys/' endpoint to create a named key.
	Key string `json:"key"`

	// [Optional for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Length of the RSA dynamic key in bits. It is 1024 by default or it can be 2048.
	KeyBits int32 `json:"key_bits"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a custom format for the key id of a signed certificate. The following variables are available for use: '{{token_display_name}}' - The display name of the token used to make the request. '{{role_name}}' - The name of the role signing the request. '{{public_key_hash}}' - A SHA256 checksum of the public key that is being signed.
	KeyIdFormat string `json:"key_id_format"`

	// [Optional for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Comma separated option specifications which will be prefixed to RSA key in authorized_keys file. Options should be valid and comply with authorized_keys file format and should not contain spaces.
	KeyOptionSpecs string `json:"key_option_specs"`

	// [Required for all types] Type of key used to login to hosts. It can be either 'otp', 'dynamic' or 'ca'. 'otp' type requires agent to be installed in remote hosts.
	KeyType string `json:"key_type"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] The maximum allowed lease duration
	MaxTtl int32 `json:"max_ttl"`

	// The duration that the SSH certificate should be backdated by at issuance.
	NotBeforeDuration int32 `json:"not_before_duration"`

	// [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Port number for SSH connection. Default is '22'. Port number does not play any role in creation of OTP. For 'otp' type, this is just a way to inform client about the port number to use. Port number will be returned to client by Vault server along with OTP.
	Port int32 `json:"port"`

	// [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] The lease duration if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the value of max_ttl.
	Ttl int32 `json:"ttl"`
}

// NewSSHWriteRoleRequestWithDefaults instantiates a new SSHWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteRoleRequestWithDefaults() *SSHWriteRoleRequest {
	var this SSHWriteRoleRequest

	this.AllowHostCertificates = false
	this.AllowUserCertificates = false
	this.AllowedDomainsTemplate = false
	this.AllowedUsersTemplate = false
	this.DefaultExtensionsTemplate = false
	this.DefaultUserTemplate = false
	this.NotBeforeDuration = 30

	return &this
}

func (o SSHWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["admin_user"] = o.AdminUser
	toSerialize["algorithm_signer"] = o.AlgorithmSigner
	toSerialize["allow_bare_domains"] = o.AllowBareDomains
	toSerialize["allow_host_certificates"] = o.AllowHostCertificates
	toSerialize["allow_subdomains"] = o.AllowSubdomains
	toSerialize["allow_user_certificates"] = o.AllowUserCertificates
	toSerialize["allow_user_key_ids"] = o.AllowUserKeyIds
	toSerialize["allowed_critical_options"] = o.AllowedCriticalOptions
	toSerialize["allowed_domains"] = o.AllowedDomains
	toSerialize["allowed_domains_template"] = o.AllowedDomainsTemplate
	toSerialize["allowed_extensions"] = o.AllowedExtensions
	toSerialize["allowed_user_key_lengths"] = o.AllowedUserKeyLengths
	toSerialize["allowed_users"] = o.AllowedUsers
	toSerialize["allowed_users_template"] = o.AllowedUsersTemplate
	toSerialize["cidr_list"] = o.CidrList
	toSerialize["default_critical_options"] = o.DefaultCriticalOptions
	toSerialize["default_extensions"] = o.DefaultExtensions
	toSerialize["default_extensions_template"] = o.DefaultExtensionsTemplate
	toSerialize["default_user"] = o.DefaultUser
	toSerialize["default_user_template"] = o.DefaultUserTemplate
	toSerialize["exclude_cidr_list"] = o.ExcludeCidrList
	toSerialize["install_script"] = o.InstallScript
	toSerialize["key"] = o.Key
	toSerialize["key_bits"] = o.KeyBits
	toSerialize["key_id_format"] = o.KeyIdFormat
	toSerialize["key_option_specs"] = o.KeyOptionSpecs
	toSerialize["key_type"] = o.KeyType
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["not_before_duration"] = o.NotBeforeDuration
	toSerialize["port"] = o.Port
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
