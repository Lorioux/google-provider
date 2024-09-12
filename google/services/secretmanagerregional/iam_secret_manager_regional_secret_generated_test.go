// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package secretmanagerregional_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
)

func TestAccSecretManagerRegionalRegionalSecretIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecretManagerRegionalRegionalSecretIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecretManagerRegionalRegionalSecretIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_secret_manager_regional_secret_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s roles/secretmanager.secretAccessor user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecretManagerRegionalRegionalSecretIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/secretmanager.secretAccessor",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_secret_manager_regional_secret_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_secret_manager_regional_secret_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/secrets/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-tf-reg-secret%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_member" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_secret_manager_regional_secret_iam_policy" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_secret_manager_regional_secret_iam_policy" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  depends_on = [
    google_secret_manager_regional_secret_iam_policy.foo
  ]
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_secret_manager_regional_secret_iam_policy" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_binding" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_binding" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_binding" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_binding" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_secret_manager_regional_secret_iam_binding" "foo2" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_secret_manager_regional_secret_iam_binding" "foo3" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_member" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

resource "google_secret_manager_regional_secret_iam_member" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_secret_manager_regional_secret_iam_member" "foo2" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_secret_manager_regional_secret_iam_member" "foo3" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccSecretManagerRegionalRegionalSecretIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "regional-secret-basic" {
  secret_id = "tf-test-tf-reg-secret%{random_suffix}"
  location = "us-central1"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "value1",
    key2 = "value2",
    key3 = "value3"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_secret_manager_regional_secret_iam_policy" "foo" {
  project = google_secret_manager_regional_secret.regional-secret-basic.project
  location = google_secret_manager_regional_secret.regional-secret-basic.location
  secret_id = google_secret_manager_regional_secret.regional-secret-basic.secret_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}