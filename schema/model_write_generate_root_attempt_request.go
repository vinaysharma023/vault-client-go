// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteGenerateRootAttemptRequest struct for WriteGenerateRootAttemptRequest
type WriteGenerateRootAttemptRequest struct {
	// Specifies a base64-encoded PGP public key.
	PgpKey string `json:"pgp_key"`
}

// NewWriteGenerateRootAttemptRequestWithDefaults instantiates a new WriteGenerateRootAttemptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteGenerateRootAttemptRequestWithDefaults() *WriteGenerateRootAttemptRequest {
	var this WriteGenerateRootAttemptRequest

	return &this
}

func (o WriteGenerateRootAttemptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["pgp_key"] = o.PgpKey

	return json.Marshal(toSerialize)
}
