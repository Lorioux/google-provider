// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFirebaseHostingSite_firebasehostingSiteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseHostingSiteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingSite_firebasehostingSiteBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_site.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id"},
			},
		},
	})
}

func testAccFirebaseHostingSite_firebasehostingSiteBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-no-app%{random_suffix}"
}
`, context)
}

func TestAccFirebaseHostingSite_firebasehostingSiteFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseHostingSiteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingSite_firebasehostingSiteFullExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_site.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id"},
			},
		},
	})
}

func testAccFirebaseHostingSite_firebasehostingSiteFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_web_app" "default" {
  provider = google-beta
  project  = "%{project_id}"
  display_name = "Test web app for Firebase Hosting"
  deletion_policy = "DELETE"
}

resource "google_firebase_hosting_site" "full" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-app%{random_suffix}"
  app_id = google_firebase_web_app.default.app_id
}
`, context)
}

func testAccCheckFirebaseHostingSiteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_hosting_site" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseHostingBasePath}}projects/{{project}}/sites/{{site_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("FirebaseHostingSite still exists at %s", url)
			}
		}

		return nil
	}
}
