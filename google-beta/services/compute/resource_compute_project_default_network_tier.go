<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go

import (
	"fmt"
	"log"
	"time"

<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
	"google.golang.org/api/compute/v1"
=======
	compute "google.golang.org/api/compute/v0.beta"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
)

func ResourceComputeProjectDefaultNetworkTier() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeProjectDefaultNetworkTierCreateOrUpdate,
		Read:   resourceComputeProjectDefaultNetworkTierRead,
		Update: resourceComputeProjectDefaultNetworkTierCreateOrUpdate,
		Delete: resourceComputeProjectDefaultNetworkTierDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
		},

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"network_tier": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  `The default network tier to be configured for the project. This field can take the following values: PREMIUM or STANDARD.`,
				ValidateFunc: validation.StringInSlice([]string{"PREMIUM", "STANDARD"}, false),
			},

			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeProjectDefaultNetworkTierCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
	projectID, err := getProject(d, config)
=======
	projectID, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
	if err != nil {
		return err
	}

	request := &compute.ProjectsSetDefaultNetworkTierRequest{
		NetworkTier: d.Get("network_tier").(string),
	}
	op, err := config.NewComputeClient(userAgent).Projects.SetDefaultNetworkTier(projectID, request).Do()
	if err != nil {
		return fmt.Errorf("SetDefaultNetworkTier failed: %s", err)
	}

	log.Printf("[DEBUG] SetDefaultNetworkTier: %d (%s)", op.Id, op.SelfLink)
	err = ComputeOperationWaitTime(config, op, projectID, "SetDefaultNetworkTier", userAgent, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("SetDefaultNetworkTier failed: %s", err)
	}

	d.SetId(projectID)

	return resourceComputeProjectDefaultNetworkTierRead(d, meta)
}

func resourceComputeProjectDefaultNetworkTierRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
	if err != nil {
		return err
	}

	projectId := d.Id()

	project, err := config.NewComputeClient(userAgent).Projects.Get(projectId).Do()
	if err != nil {
<<<<<<< HEAD:google/resource_compute_project_default_network_tier.go
		return handleNotFoundError(err, d, fmt.Sprintf("Project data for project %q", projectId))
=======
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("Project data for project %q", projectId))
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_project_default_network_tier.go
	}

	err = d.Set("network_tier", project.DefaultNetworkTier)
	if err != nil {
		return fmt.Errorf("Error setting default network tier: %s", err)
	}

	if err := d.Set("project", projectId); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}

	return nil
}

func resourceComputeProjectDefaultNetworkTierDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[WARNING] Default Network Tier will be only removed from Terraform state, but will be left intact on GCP.")

	return schema.RemoveFromState(d, meta)
}
