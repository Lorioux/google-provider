<<<<<<< HEAD:google/data_source_google_compute_node_types.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_node_types.go

import (
	"fmt"
	"log"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_node_types.go

	"google.golang.org/api/compute/v1"
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	compute "google.golang.org/api/compute/v0.beta"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_node_types.go
)

func DataSourceGoogleComputeNodeTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeNodeTypesRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGoogleComputeNodeTypesRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_node_types.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_node_types.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_node_types.go
	project, err := getProject(d, config)
=======
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_node_types.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_node_types.go
	zone, err := getZone(d, config)
=======
	zone, err := tpgresource.GetZone(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_node_types.go
	if err != nil {
		return fmt.Errorf("Please specify zone to get appropriate node types for zone. Unable to get zone: %s", err)
	}

	resp, err := config.NewComputeClient(userAgent).NodeTypes.List(project, zone).Do()
	if err != nil {
		return err
	}
	nodeTypes := flattenComputeNodeTypes(resp.Items)
	log.Printf("[DEBUG] Received Google Compute Regions: %q", nodeTypes)

	if err := d.Set("names", nodeTypes); err != nil {
		return fmt.Errorf("Error setting names: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("zone", zone); err != nil {
		return fmt.Errorf("Error setting zone: %s", err)
	}
	d.SetId(fmt.Sprintf("projects/%s/zones/%s", project, zone))

	return nil
}

func flattenComputeNodeTypes(nodeTypes []*compute.NodeType) []string {
	result := make([]string, len(nodeTypes))
	for i, nodeType := range nodeTypes {
		result[i] = nodeType.Name
	}
	sort.Strings(result)
	return result
}
