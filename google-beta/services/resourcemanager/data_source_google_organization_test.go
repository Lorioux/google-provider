// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager_test

import (
	"fmt"
	"regexp"
	"testing"

<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_organization_test.go
========
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_organization_test.go
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleOrganization_byFullName(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_organization_test.go
	orgId := GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	orgId := envvar.GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_organization_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleOrganization_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_organization.org", "id", name),
					resource.TestCheckResourceAttr("data.google_organization.org", "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleOrganization_byShortName(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_organization_test.go
	orgId := GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	orgId := envvar.GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_organization_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleOrganization_byName(orgId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_organization.org", "id", name),
					resource.TestCheckResourceAttr("data.google_organization.org", "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleOrganization_byDomain(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_organization_test.go
	name := RandString(t, 16) + ".com"

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	name := acctest.RandString(t, 16) + ".com"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_organization_test.go
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleOrganization_byDomain(name),
				ExpectError: regexp.MustCompile("Organization not found: " + name),
			},
		},
	})
}

func testAccCheckGoogleOrganization_byName(name string) string {
	return fmt.Sprintf(`
data "google_organization" "org" {
  organization = "%s"
}
`, name)
}

func testAccCheckGoogleOrganization_byDomain(name string) string {
	return fmt.Sprintf(`
data "google_organization" "org" {
  domain = "%s"
}
`, name)
}
