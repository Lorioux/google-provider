<<<<<<< HEAD:google/data_source_google_global_compute_forwarding_rule.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_global_compute_forwarding_rule.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_global_compute_forwarding_rule.go
)

func DataSourceGoogleComputeGlobalForwardingRule() *schema.Resource {
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeGlobalForwardingRule().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleComputeGlobalForwardingRule() *schema.Resource {
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeGlobalForwardingRule().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_global_compute_forwarding_rule.go

	return &schema.Resource{
		Read:   dataSourceGoogleComputeGlobalForwardingRuleRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleComputeGlobalForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_global_compute_forwarding_rule.go
	config := meta.(*Config)

	name := d.Get("name").(string)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	name := d.Get("name").(string)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_global_compute_forwarding_rule.go
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("projects/%s/global/forwardingRules/%s", project, name))

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}
