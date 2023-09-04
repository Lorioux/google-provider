<<<<<<< HEAD:google/data_source_google_compute_region_ssl_certificate.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_ssl_certificate.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_region_ssl_certificate.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_ssl_certificate.go
)

func DataSourceGoogleRegionComputeSslCertificate() *schema.Resource {
	// Generate datasource schema from resource
<<<<<<< HEAD:google/data_source_google_compute_region_ssl_certificate.go
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeRegionSslCertificate().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")
	addOptionalFieldsToSchema(dsSchema, "region")
=======
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceComputeRegionSslCertificate().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "region")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_ssl_certificate.go

	return &schema.Resource{
		Read:   dataSourceComputeRegionSslCertificateRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeRegionSslCertificateRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_region_ssl_certificate.go
	config := meta.(*Config)

	project, region, name, err := GetRegionalResourcePropertiesFromSelfLinkOrSchema(d, config)
=======
	config := meta.(*transport_tpg.Config)

	project, region, name, err := tpgresource.GetRegionalResourcePropertiesFromSelfLinkOrSchema(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_region_ssl_certificate.go
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("projects/%s/regions/%s/sslCertificates/%s", project, region, name))

	return resourceComputeRegionSslCertificateRead(d, meta)
}
