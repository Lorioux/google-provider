<<<<<<< HEAD:google/data_source_google_compute_ssl_certificate.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ssl_certificate.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_ssl_certificate.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ssl_certificate.go
)

func DataSourceGoogleComputeSslCertificate() *schema.Resource {
	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_google_compute_ssl_certificate.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeSslCertificate().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeSslCertificate().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ssl_certificate.go

	return &schema.Resource{
		Read:   dataSourceComputeSslCertificateRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeSslCertificateRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_ssl_certificate.go
	config := meta.(*Config)

	project, err := getProject(d, config)
=======
	config := meta.(*transport_tpg.Config)

	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_ssl_certificate.go
	if err != nil {
		return err
	}
	certificateName := d.Get("name").(string)

	d.SetId(fmt.Sprintf("projects/%s/global/sslCertificates/%s", project, certificateName))

	return resourceComputeSslCertificateRead(d, meta)
}
