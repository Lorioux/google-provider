// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package billing_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_billing_account_test.go
)

func TestAccDataSourceGoogleBillingAccount_byFullName(t *testing.T) {
	billingId := GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleBillingAccount_byFullName(t *testing.T) {
	billingId := envvar.GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/billing/data_source_google_billing_account_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleBillingAccount_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "id", billingId),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "name", name),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "open", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byShortName(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_billing_account_test.go
	billingId := GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	billingId := envvar.GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/billing/data_source_google_billing_account_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleBillingAccount_byName(billingId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "id", billingId),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "name", name),
					resource.TestCheckResourceAttr("data.google_billing_account.acct", "open", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byFullNameClosed(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_billing_account_test.go
	billingId := GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	billingId := envvar.GetTestMasterBillingAccountFromEnv(t)
	name := "billingAccounts/" + billingId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/billing/data_source_google_billing_account_test.go
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleBillingAccount_byNameClosed(name),
				ExpectError: regexp.MustCompile("Billing account not found: " + name),
			},
		},
	})
}

func TestAccDataSourceGoogleBillingAccount_byDisplayName(t *testing.T) {
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_billing_account_test.go
	name := RandString(t, 16)

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	name := acctest.RandString(t, 16)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/billing/data_source_google_billing_account_test.go
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleBillingAccount_byDisplayName(name),
				ExpectError: regexp.MustCompile("Billing account not found: " + name),
			},
		},
	})
}

func testAccCheckGoogleBillingAccount_byName(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  billing_account = "%s"
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_billing_account_test.go
========
  lookup_projects = false
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/billing/data_source_google_billing_account_test.go
}
`, name)
}

func testAccCheckGoogleBillingAccount_byNameClosed(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  billing_account = "%s"
  open            = false
}
`, name)
}

func testAccCheckGoogleBillingAccount_byDisplayName(name string) string {
	return fmt.Sprintf(`
data "google_billing_account" "acct" {
  display_name = "%s"
}
`, name)
}
