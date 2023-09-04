<<<<<<< HEAD:google/data_source_google_compute_backend_bucket.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_backend_bucket.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_backend_bucket.go
)

func DataSourceGoogleComputeBackendBucket() *schema.Resource {
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeBackendBucket().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleComputeBackendBucket() *schema.Resource {
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeBackendBucket().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_backend_bucket.go

	return &schema.Resource{
		Read:   dataSourceComputeBackendBucketRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeBackendBucketRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_backend_bucket.go
	config := meta.(*Config)

	backendBucketName := d.Get("name").(string)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	backendBucketName := d.Get("name").(string)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_backend_bucket.go
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("projects/%s/global/backendBuckets/%s", project, backendBucketName))

	return resourceComputeBackendBucketRead(d, meta)
}
