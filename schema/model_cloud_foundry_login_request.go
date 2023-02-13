// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CloudFoundryLoginRequest struct for CloudFoundryLoginRequest
type CloudFoundryLoginRequest struct {
	// The full body of the file available at the CF_INSTANCE_CERT path on the CF instance.
	CfInstanceCert string `json:"cf_instance_cert"`

	// The name of the role to authenticate against.
	Role string `json:"role"`

	// The signature generated by the client certificate's private key.
	Signature string `json:"signature"`

	// The date and time used to construct the signature.
	SigningTime string `json:"signing_time"`
}

// NewCloudFoundryLoginRequestWithDefaults instantiates a new CloudFoundryLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFoundryLoginRequestWithDefaults() *CloudFoundryLoginRequest {
	var this CloudFoundryLoginRequest

	return &this
}

func (o CloudFoundryLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cf_instance_cert"] = o.CfInstanceCert
	toSerialize["role"] = o.Role
	toSerialize["signature"] = o.Signature
	toSerialize["signing_time"] = o.SigningTime

	return json.Marshal(toSerialize)
}
