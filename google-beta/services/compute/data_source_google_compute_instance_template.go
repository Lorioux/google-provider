<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go

import (
	"fmt"
	"sort"

<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"google.golang.org/api/compute/v1"
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	compute "google.golang.org/api/compute/v0.beta"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
)

func DataSourceGoogleComputeInstanceTemplate() *schema.Resource {
	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeInstanceTemplate().Schema)
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeInstanceTemplate().Schema)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go

	dsSchema["filter"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
=======
	dsSchema["self_link_unique"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
	dsSchema["most_recent"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
	}

	// Set 'Optional' schema elements
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
	addOptionalFieldsToSchema(dsSchema, "name", "filter", "most_recent", "project")

	dsSchema["name"].ExactlyOneOf = []string{"name", "filter"}
	dsSchema["filter"].ExactlyOneOf = []string{"name", "filter"}
=======
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "name", "filter", "most_recent", "project", "self_link_unique")

	mutuallyExclusive := []string{"name", "filter", "self_link_unique"}
	for _, n := range mutuallyExclusive {
		dsSchema[n].ExactlyOneOf = mutuallyExclusive
	}
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go

	return &schema.Resource{
		Read:   datasourceComputeInstanceTemplateRead,
		Schema: dsSchema,
	}
}

func datasourceComputeInstanceTemplateRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
	config := meta.(*Config)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
	if err != nil {
		return err
	}

	if v, ok := d.GetOk("name"); ok {
		return retrieveInstance(d, meta, project, v.(string))
	}
	if v, ok := d.GetOk("filter"); ok {
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
		userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
		if err != nil {
			return err
		}

		templates, err := config.NewComputeClient(userAgent).InstanceTemplates.List(project).Filter(v.(string)).Do()
		if err != nil {
			return fmt.Errorf("error retrieving list of instance templates: %s", err)
		}

		mostRecent := d.Get("most_recent").(bool)
		if mostRecent {
			sort.Sort(ByCreationTimestamp(templates.Items))
		}

		count := len(templates.Items)
		if count == 1 || count > 1 && mostRecent {
			return retrieveInstance(d, meta, project, templates.Items[0].Name)
		}

		return fmt.Errorf("your filter has returned %d instance template(s). Please refine your filter or set most_recent to return exactly one instance template", len(templates.Items))
	}
<<<<<<< HEAD:google/data_source_google_compute_instance_template.go

	return fmt.Errorf("one of name or filters must be set")
=======
	if v, ok := d.GetOk("self_link_unique"); ok {
		return retrieveInstanceFromUniqueId(d, meta, project, v.(string))
	}

	return fmt.Errorf("one of name, filters or self_link_unique must be set")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
}

func retrieveInstance(d *schema.ResourceData, meta interface{}, project, name string) error {
	d.SetId("projects/" + project + "/global/instanceTemplates/" + name)

	return resourceComputeInstanceTemplateRead(d, meta)
}

<<<<<<< HEAD:google/data_source_google_compute_instance_template.go
=======
func retrieveInstanceFromUniqueId(d *schema.ResourceData, meta interface{}, project, self_link_unique string) error {
	normalId, _ := parseUniqueId(self_link_unique)
	d.SetId(normalId)
	d.Set("self_link_unique", self_link_unique)

	return resourceComputeInstanceTemplateRead(d, meta)
}

>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_template.go
// ByCreationTimestamp implements sort.Interface for []*InstanceTemplate based on
// the CreationTimestamp field.
type ByCreationTimestamp []*compute.InstanceTemplate

func (a ByCreationTimestamp) Len() int      { return len(a) }
func (a ByCreationTimestamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCreationTimestamp) Less(i, j int) bool {
	return a[i].CreationTimestamp > a[j].CreationTimestamp
}
