<<<<<<< HEAD:google/data_source_google_compute_region_network_endpoint_group.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_network_endpoint_group.go

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_region_network_endpoint_group.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_network_endpoint_group.go
)

func DataSourceGoogleComputeRegionNetworkEndpointGroup() *schema.Resource {
	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_google_compute_region_network_endpoint_group.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeRegionNetworkEndpointGroup().Schema)

	addOptionalFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "region")
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "self_link")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeRegionNetworkEndpointGroup().Schema)

	tpgresource.AddOptionalFieldsToSchema(dsSchema, "name")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "region")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "self_link")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_network_endpoint_group.go

	return &schema.Resource{
		Read:   dataSourceComputeRegionNetworkEndpointGroupRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeRegionNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_region_network_endpoint_group.go
	config := meta.(*Config)
	if name, ok := d.GetOk("name"); ok {
		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		region, err := getRegion(d, config)
=======
	config := meta.(*transport_tpg.Config)
	if name, ok := d.GetOk("name"); ok {
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return err
		}
		region, err := tpgresource.GetRegion(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_network_endpoint_group.go
		if err != nil {
			return err
		}

		d.SetId(fmt.Sprintf("projects/%s/regions/%s/networkEndpointGroups/%s", project, region, name.(string)))
	} else if selfLink, ok := d.GetOk("self_link"); ok {
<<<<<<< HEAD:google/data_source_google_compute_region_network_endpoint_group.go
		parsed, err := ParseNetworkEndpointGroupRegionalFieldValue(selfLink.(string), d, config)
=======
		parsed, err := tpgresource.ParseNetworkEndpointGroupRegionalFieldValue(selfLink.(string), d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_network_endpoint_group.go
		if err != nil {
			return err
		}
		if err := d.Set("name", parsed.Name); err != nil {
			return fmt.Errorf("Error setting name: %s", err)
		}
		if err := d.Set("project", parsed.Project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("region", parsed.Region); err != nil {
			return fmt.Errorf("Error setting region: %s", err)
		}

		d.SetId(fmt.Sprintf("projects/%s/regions/%s/networkEndpointGroups/%s", parsed.Project, parsed.Region, parsed.Name))
	} else {
		return errors.New("Must provide either `self_link` or `region/name`")
	}

	return resourceComputeRegionNetworkEndpointGroupRead(d, meta)
}
