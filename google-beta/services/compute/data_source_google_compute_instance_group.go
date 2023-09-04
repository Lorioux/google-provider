// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_instance_group.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group.go
)

func DataSourceGoogleComputeInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComputeInstanceGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"self_link"},
			},

			"self_link": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name", "zone"},
			},

			"zone": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"self_link"},
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"instances": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"named_port": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceComputeInstanceGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	if name, ok := d.GetOk("name"); ok {
		zone, err := tpgresource.GetZone(d, config)
		if err != nil {
			return err
		}
<<<<<<< HEAD:google/data_source_google_compute_instance_group.go
		project, err := getProject(d, config)
=======
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s", project, zone, name.(string)))
	} else if selfLink, ok := d.GetOk("self_link"); ok {
		parsed, err := tpgresource.ParseInstanceGroupFieldValue(selfLink.(string), d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group.go
		if err != nil {
			return fmt.Errorf("InstanceGroup name, zone or project could not be parsed from %s", selfLink)
		}
<<<<<<< HEAD:google/data_source_google_compute_instance_group.go
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s", project, zone, name.(string)))
	} else if selfLink, ok := d.GetOk("self_link"); ok {
		parsed, err := ParseInstanceGroupFieldValue(selfLink.(string), d, config)
		if err != nil {
			return fmt.Errorf("InstanceGroup name, zone or project could not be parsed from %s", selfLink)
		}
=======
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_group.go
		if err := d.Set("name", parsed.Name); err != nil {
			return fmt.Errorf("Error setting name: %s", err)
		}
		if err := d.Set("zone", parsed.Zone); err != nil {
			return fmt.Errorf("Error setting zone: %s", err)
		}
		if err := d.Set("project", parsed.Project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/instanceGroups/%s", parsed.Project, parsed.Zone, parsed.Name))
	} else {
		return errors.New("Must provide either `self_link` or `zone/name`")
	}

	return resourceComputeInstanceGroupRead(d, meta)
}
