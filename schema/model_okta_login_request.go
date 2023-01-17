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

// OktaLoginRequest struct for OktaLoginRequest
type OktaLoginRequest struct {
	// Nonce provided if performing login that requires number verification challenge. Logins through the vault login CLI command will automatically generate a nonce.
	Nonce string `json:"nonce"`
	// Password for this user.
	Password string `json:"password"`
	// Preferred factor provider.
	Provider string `json:"provider"`
	// TOTP passcode.
	Totp string `json:"totp"`
}

// NewOktaLoginRequestWithDefaults instantiates a new OktaLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaLoginRequestWithDefaults() *OktaLoginRequest {
	var this OktaLoginRequest

	return &this
}

func (o OktaLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["nonce"] = o.Nonce
	toSerialize["password"] = o.Password
	toSerialize["provider"] = o.Provider
	toSerialize["totp"] = o.Totp

	return json.Marshal(toSerialize)
}