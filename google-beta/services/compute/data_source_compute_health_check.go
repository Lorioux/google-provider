<<<<<<< HEAD:google/data_source_compute_health_check.go
package google

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceGoogleComputeHealthCheck() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeHealthCheck().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleComputeHealthCheck() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeHealthCheck().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_health_check.go

	return &schema.Resource{
		Read:   dataSourceGoogleComputeHealthCheckRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleComputeHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_compute_health_check.go
	id, err := replaceVars(d, meta.(*Config), "projects/{{project}}/global/healthChecks/{{name}}")
=======
	id, err := tpgresource.ReplaceVars(d, meta.(*transport_tpg.Config), "projects/{{project}}/global/healthChecks/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_health_check.go
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceComputeHealthCheckRead(d, meta)
}
