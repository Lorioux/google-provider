<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
package google
=======
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
=======
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgdclresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
)

func ResourceComputeNetworkFirewallPolicyAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkFirewallPolicyAssociationCreate,
		Read:   resourceComputeNetworkFirewallPolicyAssociationRead,
		Delete: resourceComputeNetworkFirewallPolicyAssociationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkFirewallPolicyAssociationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"attachment_target": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
				DiffSuppressFunc: compareSelfLinkOrResourceName,
=======
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
				Description:      "The target that the firewall policy is attached to.",
			},

			"firewall_policy": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
				DiffSuppressFunc: compareSelfLinkOrResourceName,
=======
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
				Description:      "The firewall policy ID of the association.",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name for an association.",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
				DiffSuppressFunc: compareSelfLinkOrResourceName,
=======
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
				Description:      "The project for the resource",
			},

			"short_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The short name of the firewall policy of the association.",
			},
		},
	}
}

func resourceComputeNetworkFirewallPolicyAssociationCreate(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	config := meta.(*Config)
	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}

	obj := &compute.NetworkFirewallPolicyAssociation{
		AttachmentTarget: dcl.String(d.Get("attachment_target").(string)),
		FirewallPolicy:   dcl.String(d.Get("firewall_policy").(string)),
		Name:             dcl.String(d.Get("name").(string)),
		Project:          dcl.String(project),
	}

<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	id, err := replaceVarsForId(d, config, "projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
=======
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	directive := CreateDirective
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	directive := tpgdclresource.CreateDirective
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
=======
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyNetworkFirewallPolicyAssociation(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating NetworkFirewallPolicyAssociation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkFirewallPolicyAssociation %q: %#v", d.Id(), res)

	return resourceComputeNetworkFirewallPolicyAssociationRead(d, meta)
}

func resourceComputeNetworkFirewallPolicyAssociationRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	config := meta.(*Config)
	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}

	obj := &compute.NetworkFirewallPolicyAssociation{
		AttachmentTarget: dcl.String(d.Get("attachment_target").(string)),
		FirewallPolicy:   dcl.String(d.Get("firewall_policy").(string)),
		Name:             dcl.String(d.Get("name").(string)),
		Project:          dcl.String(project),
	}

<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
=======
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetNetworkFirewallPolicyAssociation(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("ComputeNetworkFirewallPolicyAssociation %q", d.Id())
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
		return handleNotFoundDCLError(err, d, resourceName)
=======
		return tpgdclresource.HandleNotFoundDCLError(err, d, resourceName)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	}

	if err = d.Set("attachment_target", res.AttachmentTarget); err != nil {
		return fmt.Errorf("error setting attachment_target in state: %s", err)
	}
	if err = d.Set("firewall_policy", res.FirewallPolicy); err != nil {
		return fmt.Errorf("error setting firewall_policy in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("short_name", res.ShortName); err != nil {
		return fmt.Errorf("error setting short_name in state: %s", err)
	}

	return nil
}

func resourceComputeNetworkFirewallPolicyAssociationDelete(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	config := meta.(*Config)
	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}

	obj := &compute.NetworkFirewallPolicyAssociation{
		AttachmentTarget: dcl.String(d.Get("attachment_target").(string)),
		FirewallPolicy:   dcl.String(d.Get("firewall_policy").(string)),
		Name:             dcl.String(d.Get("name").(string)),
		Project:          dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting NetworkFirewallPolicyAssociation %q", d.Id())
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
=======
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLComputeClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	if err := client.DeleteNetworkFirewallPolicyAssociation(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting NetworkFirewallPolicyAssociation: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting NetworkFirewallPolicyAssociation %q", d.Id())
	return nil
}

func resourceComputeNetworkFirewallPolicyAssociationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	config := meta.(*Config)

	if err := parseImportId([]string{
=======
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
		"projects/(?P<project>[^/]+)/global/firewallPolicies/(?P<firewall_policy>[^/]+)/associations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<firewall_policy>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
<<<<<<< HEAD:google/resource_compute_network_firewall_policy_association.go
	id, err := replaceVarsForId(d, config, "projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
=======
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/resource_compute_network_firewall_policy_association.go
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
