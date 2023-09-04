<<<<<<< HEAD:google/data_source_compute_network_peering.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_peering.go

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_compute_network_peering.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_peering.go
)

const regexGCEName = "^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$"

func DataSourceComputeNetworkPeering() *schema.Resource {

<<<<<<< HEAD:google/data_source_compute_network_peering.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeNetworkPeering().Schema)
	addRequiredFieldsToSchema(dsSchema, "name", "network")

	dsSchema["name"].ValidateFunc = validateRegexp(regexGCEName)
	dsSchema["network"].ValidateFunc = validateRegexp(peerNetworkLinkRegex)
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeNetworkPeering().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name", "network")

	dsSchema["name"].ValidateFunc = verify.ValidateRegexp(regexGCEName)
	dsSchema["network"].ValidateFunc = verify.ValidateRegexp(peerNetworkLinkRegex)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_peering.go
	return &schema.Resource{
		Read:   dataSourceComputeNetworkPeeringRead,
		Schema: dsSchema,
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(4 * time.Minute),
		},
	}
}

func dataSourceComputeNetworkPeeringRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_compute_network_peering.go
	config := meta.(*Config)

	networkFieldValue, err := ParseNetworkFieldValue(d.Get("network").(string), d, config)
=======
	config := meta.(*transport_tpg.Config)

	networkFieldValue, err := tpgresource.ParseNetworkFieldValue(d.Get("network").(string), d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_peering.go
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", networkFieldValue.Name, d.Get("name").(string)))

	return resourceComputeNetworkPeeringRead(d, meta)
}
