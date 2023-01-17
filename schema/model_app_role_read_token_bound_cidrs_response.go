/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenBoundCIDRsResponse struct for AppRoleReadTokenBoundCIDRsResponse
type AppRoleReadTokenBoundCIDRsResponse struct {
	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`
}

// NewAppRoleReadTokenBoundCIDRsResponseWithDefaults instantiates a new AppRoleReadTokenBoundCIDRsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenBoundCIDRsResponseWithDefaults() *AppRoleReadTokenBoundCIDRsResponse {
	var this AppRoleReadTokenBoundCIDRsResponse

	return &this
}

func (o AppRoleReadTokenBoundCIDRsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs

	return json.Marshal(toSerialize)
}