// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"errors"
<<<<<<< HEAD:google/metadata.go
	"fmt"
	"log"
	"sort"

	"google.golang.org/api/compute/v1"
)

const METADATA_FINGERPRINT_RETRIES = 10

=======
	"sort"

	compute "google.golang.org/api/compute/v0.beta"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/metadata.go
// Since the google compute API uses optimistic locking, there is a chance
// we need to resubmit our updated metadata. To do this, you need to provide
// an update function that attempts to submit your metadata
func MetadataRetryWrapper(update func() error) error {
<<<<<<< HEAD:google/metadata.go
	attempt := 0
	for attempt < METADATA_FINGERPRINT_RETRIES {
		err := update()
		if err == nil {
			return nil
		}

		if ok, _ := isFingerprintError(err); !ok {
			// Something else went wrong, don't retry
			return err
		}

		log.Printf("[DEBUG] Dismissed an error as retryable as a fingerprint mismatch: %s", err)
		attempt++
	}
	return fmt.Errorf("Failed to update metadata after %d retries", attempt)
=======
	return transport_tpg.MetadataRetryWrapper(update)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/metadata.go
}

// Update the metadata (serverMD) according to the provided diff (oldMDMap v
// newMDMap).
func MetadataUpdate(oldMDMap map[string]interface{}, newMDMap map[string]interface{}, serverMD *compute.Metadata) {
	curMDMap := make(map[string]string)
	// Load metadata on server into map
	for _, kv := range serverMD.Items {
		// If the server state has a key that we had in our old
		// state, but not in our new state, we should delete it
		_, okOld := oldMDMap[kv.Key]
		_, okNew := newMDMap[kv.Key]
		if okOld && !okNew {
			continue
		} else {
			curMDMap[kv.Key] = *kv.Value
		}
	}

	// Insert new metadata into existing metadata (overwriting when needed)
	for key, val := range newMDMap {
		curMDMap[key] = val.(string)
	}

	// Reformat old metadata into a list
	serverMD.Items = nil
	for key, val := range curMDMap {
		v := val
		serverMD.Items = append(serverMD.Items, &compute.MetadataItems{
			Key:   key,
			Value: &v,
		})
	}
}

// Update the beta metadata (serverMD) according to the provided diff (oldMDMap v
// newMDMap).
func BetaMetadataUpdate(oldMDMap map[string]interface{}, newMDMap map[string]interface{}, serverMD *compute.Metadata) {
	curMDMap := make(map[string]string)
	// Load metadata on server into map
	for _, kv := range serverMD.Items {
		// If the server state has a key that we had in our old
		// state, but not in our new state, we should delete it
		_, okOld := oldMDMap[kv.Key]
		_, okNew := newMDMap[kv.Key]
		if okOld && !okNew {
			continue
		} else {
			curMDMap[kv.Key] = *kv.Value
		}
	}

	// Insert new metadata into existing metadata (overwriting when needed)
	for key, val := range newMDMap {
		curMDMap[key] = val.(string)
	}

	// Reformat old metadata into a list
	serverMD.Items = nil
	for key, val := range curMDMap {
		v := val
		serverMD.Items = append(serverMD.Items, &compute.MetadataItems{
			Key:   key,
			Value: &v,
		})
	}
}

func expandComputeMetadata(m map[string]interface{}) []*compute.MetadataItems {
	metadata := make([]*compute.MetadataItems, len(m))
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// Append new metadata to existing metadata
	for _, key := range keys {
		v := m[key].(string)
		metadata = append(metadata, &compute.MetadataItems{
			Key:   key,
			Value: &v,
		})
	}

	return metadata
}

func flattenMetadataBeta(metadata *compute.Metadata) map[string]string {
	metadataMap := make(map[string]string)
	for _, item := range metadata.Items {
		metadataMap[item.Key] = *item.Value
	}
	return metadataMap
}

// This function differs from flattenMetadataBeta only in that it takes
// compute.metadata rather than compute.metadata as an argument. It should
// be removed in favour of flattenMetadataBeta if/when all resources using it get
// beta support.
<<<<<<< HEAD:google/metadata.go
func flattenMetadata(metadata *compute.Metadata) map[string]interface{} {
=======
func FlattenMetadata(metadata *compute.Metadata) map[string]interface{} {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/metadata.go
	metadataMap := make(map[string]interface{})
	for _, item := range metadata.Items {
		metadataMap[item.Key] = *item.Value
	}
	return metadataMap
}

<<<<<<< HEAD:google/metadata.go
func resourceInstanceMetadata(d TerraformResourceData) (*compute.Metadata, error) {
=======
func resourceInstanceMetadata(d tpgresource.TerraformResourceData) (*compute.Metadata, error) {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/metadata.go
	m := &compute.Metadata{}
	mdMap := d.Get("metadata").(map[string]interface{})
	if v, ok := d.GetOk("metadata_startup_script"); ok && v.(string) != "" {
		if w, ok := mdMap["startup-script"]; ok {
			// metadata.startup-script could be from metadata_startup_script in the first place
			if v != w {
				return nil, errors.New("Cannot provide both metadata_startup_script and metadata.startup-script.")
			}
		}
		mdMap["startup-script"] = v
	}
	if len(mdMap) > 0 {
		m.Items = make([]*compute.MetadataItems, 0, len(mdMap))
		var keys []string
		for k := range mdMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := mdMap[k].(string)
			m.Items = append(m.Items, &compute.MetadataItems{
				Key:   k,
				Value: &v,
			})
		}

		// Set the fingerprint. If the metadata has never been set before
		// then this will just be blank.
		m.Fingerprint = d.Get("metadata_fingerprint").(string)
	}

	return m, nil
}
