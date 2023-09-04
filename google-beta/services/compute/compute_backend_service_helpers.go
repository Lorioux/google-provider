<<<<<<< HEAD:google/compute_backend_service_helpers.go
package google

import (
	"google.golang.org/api/compute/v1"
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	compute "google.golang.org/api/compute/v0.beta"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/compute_backend_service_helpers.go
)

// Incredibly hacky way of getting a reference to an SPR of the right type into
// the generated BackendService code. goimports will always import `compute`, so
// we need to provide the import manually to be able to switch libraries. Since
// this is a problem exactly once, just provide a function in a file where we
// *can* easily pick the imported copy, and return the correct struct.
func emptySecurityPolicyReference() *compute.SecurityPolicyReference {
	return &compute.SecurityPolicyReference{}
}
