<<<<<<< HEAD:google/data_source_compute_network_endpoint_group.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_endpoint_group.go

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_compute_network_endpoint_group.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_endpoint_group.go
)

func DataSourceGoogleComputeNetworkEndpointGroup() *schema.Resource {
	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_compute_network_endpoint_group.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeNetworkEndpointGroup().Schema)

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "zone")
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "self_link")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeNetworkEndpointGroup().Schema)

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "name")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "zone")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "self_link")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_endpoint_group.go

	return &schema.Resource{
		Read:   dataSourceComputeNetworkEndpointGroupRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_compute_network_endpoint_group.go
	config := meta.(*Config)
	if name, ok := d.GetOk("name"); ok {
		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		zone, err := getZone(d, config)
=======
	config := meta.(*transport_tpg.Config)
	if name, ok := d.GetOk("name"); ok {
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return err
		}
		zone, err := tpgresource.GetZone(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_endpoint_group.go
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", project, zone, name.(string)))
	} else if selfLink, ok := d.GetOk("self_link"); ok {
<<<<<<< HEAD:google/data_source_compute_network_endpoint_group.go
		parsed, err := ParseNetworkEndpointGroupFieldValue(selfLink.(string), d, config)
=======
		parsed, err := tpgresource.ParseNetworkEndpointGroupFieldValue(selfLink.(string), d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_compute_network_endpoint_group.go
		if err != nil {
			return err
		}
		if err := d.Set("name", parsed.Name); err != nil {
			return fmt.Errorf("Error setting name: %s", err)
		}
		if err := d.Set("zone", parsed.Zone); err != nil {
			return fmt.Errorf("Error setting zone: %s", err)
		}
		if err := d.Set("project", parsed.Project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", parsed.Project, parsed.Zone, parsed.Name))
	} else {
		return errors.New("Must provide either `self_link` or `zone/name`")
	}

	return resourceComputeNetworkEndpointGroupRead(d, meta)
}
