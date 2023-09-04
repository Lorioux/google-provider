<<<<<<< HEAD:google/data_source_google_compute_router_status.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_router_status.go
	"google.golang.org/api/compute/v1"
)

func DataSourceGoogleComputeRouterStatus() *schema.Resource {
	routeElemSchema := datasourceSchemaFromResourceSchema(ResourceComputeRoute().Schema)
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/compute/v0.beta"
)

func DataSourceGoogleComputeRouterStatus() *schema.Resource {
	routeElemSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeRoute().Schema)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go

	return &schema.Resource{
		Read: dataSourceComputeRouterStatusRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the router to query.",
				Required:    true,
				Computed:    false,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project ID of the target router.",
				Optional:    true,
				Computed:    false,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the target router.",
				Optional:    true,
				Computed:    true,
			},
			"network": {
				Type:        schema.TypeString,
				Description: "URI of the network to which this router belongs.",
				Computed:    true,
			},
			"best_routes": {
				Type:        schema.TypeList,
				Description: "Best routes for this router's network.",
				Elem: &schema.Resource{
					Schema: routeElemSchema,
				},
				Computed: true,
			},
			"best_routes_for_router": {
				Type:        schema.TypeList,
				Description: "Best routes learned by this router.",
				Elem: &schema.Resource{
					Schema: routeElemSchema,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceComputeRouterStatusRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_router_status.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_router_status.go
	project, err := getProject(d, config)
=======
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_router_status.go
	region, err := getRegion(d, config)
=======
	region, err := tpgresource.GetRegion(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go
	if err != nil {
		return err
	}

	var name string
	if n, ok := d.GetOk("name"); ok {
		name = n.(string)
	}

	resp, err := config.NewComputeClient(userAgent).Routers.GetRouterStatus(project, region, name).Do()
	if err != nil {
		return err
	}

	status := resp.Result

	if err := d.Set("network", status.Network); err != nil {
		return fmt.Errorf("Error setting network: %s", err)
	}

	if err := d.Set("best_routes", flattenRoutes(status.BestRoutes)); err != nil {
		return fmt.Errorf("Error setting best_routes: %s", err)
	}

	if err := d.Set("best_routes_for_router", flattenRoutes(status.BestRoutesForRouter)); err != nil {
		return fmt.Errorf("Error setting best_routes_for_router: %s", err)
	}

<<<<<<< HEAD:google/data_source_google_compute_router_status.go
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{name}}")
=======
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/routers/{{name}}")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_router_status.go
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return nil
}

func flattenRoutes(routes []*compute.Route) []map[string]interface{} {
	results := make([]map[string]interface{}, len(routes))

	for i, route := range routes {
		results[i] = map[string]interface{}{
			"dest_range":          route.DestRange,
			"name":                route.Name,
			"network":             route.Network,
			"description":         route.Description,
			"next_hop_gateway":    route.NextHopGateway,
			"next_hop_ilb":        route.NextHopIlb,
			"next_hop_ip":         route.NextHopIp,
			"next_hop_vpn_tunnel": route.NextHopVpnTunnel,
			"priority":            route.Priority,
			"tags":                route.Tags,
			"next_hop_network":    route.NextHopNetwork,
		}
	}

	return results
}
