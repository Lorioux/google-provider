<<<<<<< HEAD:google/data_source_google_compute_ha_vpn_gateway.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ha_vpn_gateway.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_ha_vpn_gateway.go
)

func DataSourceGoogleComputeHaVpnGateway() *schema.Resource {
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeHaVpnGateway().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "region")
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleComputeHaVpnGateway() *schema.Resource {
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeHaVpnGateway().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "region")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ha_vpn_gateway.go

	return &schema.Resource{
		Read:   dataSourceGoogleComputeHaVpnGatewayRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleComputeHaVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_ha_vpn_gateway.go
	config := meta.(*Config)

	name := d.Get("name").(string)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	name := d.Get("name").(string)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ha_vpn_gateway.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_ha_vpn_gateway.go
	region, err := getRegion(d, config)
=======
	region, err := tpgresource.GetRegion(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ha_vpn_gateway.go
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("projects/%s/regions/%s/vpnGateways/%s", project, region, name))

	return resourceComputeHaVpnGatewayRead(d, meta)
}
