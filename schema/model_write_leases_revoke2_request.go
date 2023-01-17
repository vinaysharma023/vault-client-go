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

// WriteLeasesRevoke2Request struct for WriteLeasesRevoke2Request
type WriteLeasesRevoke2Request struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewWriteLeasesRevoke2RequestWithDefaults instantiates a new WriteLeasesRevoke2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesRevoke2RequestWithDefaults() *WriteLeasesRevoke2Request {
	var this WriteLeasesRevoke2Request

	this.Sync = true

	return &this
}

func (o WriteLeasesRevoke2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease_id"] = o.LeaseId
	toSerialize["sync"] = o.Sync

	return json.Marshal(toSerialize)
}