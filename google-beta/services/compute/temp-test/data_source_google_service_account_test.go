// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_service_account_test.go
========
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_service_account_test.go
)

func TestAccDatasourceGoogleServiceAccount_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.google_service_account.acceptance"
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_service_account_test.go
	account := fmt.Sprintf("tf-test-%d", RandInt(t))

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	account := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_service_account_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleServiceAccount_basic(account),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_service_account_test.go
						resourceName, "id", fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", GetTestProjectFromEnv(), account, GetTestProjectFromEnv())),
========
						resourceName, "id", fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", envvar.GetTestProjectFromEnv(), account, envvar.GetTestProjectFromEnv())),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_service_account_test.go
					resource.TestCheckResourceAttrSet(resourceName, "email"),
					resource.TestCheckResourceAttrSet(resourceName, "unique_id"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "member"),
				),
			},
		},
	})
}

func testAccCheckGoogleServiceAccount_basic(account string) string {
	return fmt.Sprintf(`
resource "google_service_account" "acceptance" {
  account_id   = "%s"
  display_name = "Testing Account"
}

data "google_service_account" "acceptance" {
  account_id = google_service_account.acceptance.account_id
}
`, account)
}
