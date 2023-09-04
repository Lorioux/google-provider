<<<<<<< HEAD:google/data_source_google_compute_instance_group_manager.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group_manager.go

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_instance_group_manager.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group_manager.go
)

func DataSourceGoogleComputeInstanceGroupManager() *schema.Resource {

<<<<<<< HEAD:google/data_source_google_compute_instance_group_manager.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeInstanceGroupManager().Schema)
	addOptionalFieldsToSchema(dsSchema, "name", "self_link", "project", "zone")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeInstanceGroupManager().Schema)
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "name", "self_link", "project", "zone")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group_manager.go

	return &schema.Resource{
		Read:   dataSourceComputeInstanceGroupManagerRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_instance_group_manager.go
	config := meta.(*Config)
	if selfLink, ok := d.GetOk("self_link"); ok {
		parsed, err := ParseInstanceGroupFieldValue(selfLink.(string), d, config)
=======
	config := meta.(*transport_tpg.Config)
	if selfLink, ok := d.GetOk("self_link"); ok {
		parsed, err := tpgresource.ParseInstanceGroupFieldValue(selfLink.(string), d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group_manager.go
		if err != nil {
			return fmt.Errorf("InstanceGroup name, zone or project could not be parsed from %s", selfLink)
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
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/instanceGroupManagers/%s", parsed.Project, parsed.Zone, parsed.Name))
	} else if name, ok := d.GetOk("name"); ok {
<<<<<<< HEAD:google/data_source_google_compute_instance_group_manager.go
		zone, err := getZone(d, config)
		if err != nil {
			return err
		}
		project, err := getProject(d, config)
=======
		zone, err := tpgresource.GetZone(d, config)
		if err != nil {
			return err
		}
		project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group_manager.go
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/instanceGroupManagers/%s", project, zone, name.(string)))
	} else {
		return errors.New("Must provide either `self_link` or `zone/name`")
	}

	err := resourceComputeInstanceGroupManagerRead(d, meta)

	if err != nil {
		return err
	}
	if d.Id() == "" {
		return errors.New("Instance Manager Group not found")
	}
	return nil
}
