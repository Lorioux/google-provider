<<<<<<< HEAD:google/data_source_google_compute_router_nat.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_nat.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_router_nat.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_nat.go
)

func DataSourceGoogleComputeRouterNat() *schema.Resource {

<<<<<<< HEAD:google/data_source_google_compute_router_nat.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeRouterNat().Schema)

	addRequiredFieldsToSchema(dsSchema, "name", "router")
	addOptionalFieldsToSchema(dsSchema, "project", "region")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeRouterNat().Schema)

	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name", "router")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project", "region")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_nat.go

	return &schema.Resource{
		Read:   dataSourceGoogleComputeRouterNatRead,
		Schema: dsSchema,
	}

}

func dataSourceGoogleComputeRouterNatRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_router_nat.go
	config := meta.(*Config)

	id, err := replaceVars(d, config, "{{project}}/{{region}}/{{router}}/{{name}}")
=======
	config := meta.(*transport_tpg.Config)

	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{region}}/{{router}}/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_nat.go
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceComputeRouterNatRead(d, meta)
}
