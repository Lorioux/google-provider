<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go

import (
	"fmt"
	"sort"

<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/compute/v1"
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/compute/v0.beta"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go
)

func DataSourceGoogleComputeSnapshot() *schema.Resource {

	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeSnapshot().Schema)
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeSnapshot().Schema)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go

	dsSchema["filter"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	dsSchema["most_recent"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
	}

	// Set 'Optional' schema elements
<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
	addOptionalFieldsToSchema(dsSchema, "name", "filter", "most_recent", "project")
=======
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "name", "filter", "most_recent", "project")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go

	dsSchema["name"].ExactlyOneOf = []string{"name", "filter"}
	dsSchema["filter"].ExactlyOneOf = []string{"name", "filter"}

	return &schema.Resource{
		Read:   dataSourceGoogleComputeSnapshotRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleComputeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
	config := meta.(*Config)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go
	if err != nil {
		return err
	}

	if v, ok := d.GetOk("name"); ok {
		return retrieveSnapshot(d, meta, project, v.(string))
	}

	if v, ok := d.GetOk("filter"); ok {
<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
		userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go
		if err != nil {
			return err
		}

		projectGetCall := config.NewResourceManagerClient(userAgent).Projects.Get(project)

		if config.UserProjectOverride {
			billingProject := project

			// err == nil indicates that the billing_project value was found
<<<<<<< HEAD:google/data_source_google_compute_snapshot.go
			if bp, err := getBillingProject(d, config); err == nil {
=======
			if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_snapshot.go
				billingProject = bp
			}
			projectGetCall.Header().Add("X-Goog-User-Project", billingProject)
		}

		//handling the pagination locally
		allSnapshots := make([]*compute.Snapshot, 0)
		token := ""
		for paginate := true; paginate; {
			snapshots, err := config.NewComputeClient(userAgent).Snapshots.List(project).Filter(v.(string)).PageToken(token).Do()
			if err != nil {
				return fmt.Errorf("error retrieving list of snapshots: %s", err)

			}
			allSnapshots = append(allSnapshots, snapshots.Items...)

			token = snapshots.NextPageToken
			paginate = token != ""
		}

		mostRecent := d.Get("most_recent").(bool)
		if mostRecent {
			sort.Sort(ByCreationTimestampOfSnapshot(allSnapshots))
		}

		count := len(allSnapshots)
		if count == 1 || count > 1 && mostRecent {
			return retrieveSnapshot(d, meta, project, allSnapshots[0].Name)
		}

		return fmt.Errorf("your filter has returned %d snapshot(s). Please refine your filter or set most_recent to return exactly one snapshot", len(allSnapshots))

	}

	return fmt.Errorf("one of name or filter must be set")
}

func retrieveSnapshot(d *schema.ResourceData, meta interface{}, project, name string) error {
	d.SetId("projects/" + project + "/global/snapshots/" + name)
	d.Set("name", name)
	return resourceComputeSnapshotRead(d, meta)
}

// ByCreationTimestamp implements sort.Interface for []*Snapshot based on
// the CreationTimestamp field.
type ByCreationTimestampOfSnapshot []*compute.Snapshot

func (a ByCreationTimestampOfSnapshot) Len() int      { return len(a) }
func (a ByCreationTimestampOfSnapshot) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCreationTimestampOfSnapshot) Less(i, j int) bool {
	return a[i].CreationTimestamp > a[j].CreationTimestamp
}
