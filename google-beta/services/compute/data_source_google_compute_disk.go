<<<<<<< HEAD:google/data_source_google_compute_disk.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_disk.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_disk.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_disk.go
)

func DataSourceGoogleComputeDisk() *schema.Resource {

<<<<<<< HEAD:google/data_source_google_compute_disk.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeDisk().Schema)
	addRequiredFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "zone")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeDisk().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "zone")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_disk.go

	return &schema.Resource{
		Read:   dataSourceGoogleComputeDiskRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleComputeDiskRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_disk.go
	config := meta.(*Config)

	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/disks/{{name}}")
=======
	config := meta.(*transport_tpg.Config)

	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/disks/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_disk.go
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceComputeDiskRead(d, meta)
}
