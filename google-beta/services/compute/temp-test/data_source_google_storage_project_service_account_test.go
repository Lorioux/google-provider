// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package storage_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_storage_project_service_account_test.go
========
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/storage/data_source_google_storage_project_service_account_test.go
)

func TestAccDataSourceGoogleStorageProjectServiceAccount_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.google_storage_project_service_account.gcs_account"

<<<<<<<< HEAD:google-beta/services/compute/temp-test/data_source_google_storage_project_service_account_test.go
	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
========
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/storage/data_source_google_storage_project_service_account_test.go
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleStorageProjectServiceAccount_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "email_address"),
					resource.TestCheckResourceAttrSet(resourceName, "member"),
				),
			},
		},
	})
}

const testAccCheckGoogleStorageProjectServiceAccount_basic = `
data "google_storage_project_service_account" "gcs_account" {
}
`
