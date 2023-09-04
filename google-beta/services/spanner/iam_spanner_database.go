// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package spanner

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/iam_spanner_database.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/spanner/v1"
)

var IamSpannerDatabaseSchema = map[string]*schema.Schema{
	"instance": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"database": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},

	"project": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
}

type SpannerDatabaseIamUpdater struct {
	project  string
	instance string
	database string
<<<<<<< HEAD:google/iam_spanner_database.go
	d        TerraformResourceData
	Config   *Config
}

func NewSpannerDatabaseIamUpdater(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	project, err := getProject(d, config)
=======
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func NewSpannerDatabaseIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	if err != nil {
		return nil, err
	}

	return &SpannerDatabaseIamUpdater{
		project:  project,
		instance: d.Get("instance").(string),
		database: d.Get("database").(string),
		d:        d,
		Config:   config,
	}, nil
}

<<<<<<< HEAD:google/iam_spanner_database.go
func SpannerDatabaseIdParseFunc(d *schema.ResourceData, config *Config) error {
	return parseImportId([]string{"(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<database>[^/]+)"}, d, config)
}

func (u *SpannerDatabaseIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
func SpannerDatabaseIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	return tpgresource.ParseImportId([]string{"(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<database>[^/]+)"}, d, config)
}

func (u *SpannerDatabaseIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD:google/iam_spanner_database.go
	p, err := u.Config.NewSpannerClient(userAgent).Projects.Instances.Databases.GetIamPolicy(spannerDatabaseId{
=======
	p, err := u.Config.NewSpannerClient(userAgent).Projects.Instances.Databases.GetIamPolicy(SpannerDatabaseId{
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
		Project:  u.project,
		Database: u.database,
		Instance: u.instance,
	}.databaseUri(), &spanner.GetIamPolicyRequest{
<<<<<<< HEAD:google/iam_spanner_database.go
		Options: &spanner.GetPolicyOptions{RequestedPolicyVersion: IamPolicyVersion},
=======
		Options: &spanner.GetPolicyOptions{RequestedPolicyVersion: tpgiamresource.IamPolicyVersion},
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	}).Do()

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	cloudResourcePolicy, err := spannerToResourceManagerPolicy(p)

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

<<<<<<< HEAD:google/iam_spanner_database.go
	cloudResourcePolicy.Version = IamPolicyVersion
=======
	cloudResourcePolicy.Version = tpgiamresource.IamPolicyVersion
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go

	return cloudResourcePolicy, nil
}

func (u *SpannerDatabaseIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	spannerPolicy, err := resourceManagerToSpannerPolicy(policy)

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

<<<<<<< HEAD:google/iam_spanner_database.go
	spannerPolicy.Version = IamPolicyVersion

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	spannerPolicy.Version = tpgiamresource.IamPolicyVersion

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/iam_spanner_database.go
	_, err = u.Config.NewSpannerClient(userAgent).Projects.Instances.Databases.SetIamPolicy(spannerDatabaseId{
=======
	_, err = u.Config.NewSpannerClient(userAgent).Projects.Instances.Databases.SetIamPolicy(SpannerDatabaseId{
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
		Project:  u.project,
		Database: u.database,
		Instance: u.instance,
	}.databaseUri(), &spanner.SetIamPolicyRequest{
		Policy: spannerPolicy,
	}).Do()

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *SpannerDatabaseIamUpdater) GetResourceId() string {
	return SpannerDatabaseId{
		Project:  u.project,
		Instance: u.instance,
		Database: u.database,
	}.TerraformId()
}

func (u *SpannerDatabaseIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-spanner-database-%s-%s-%s", u.project, u.instance, u.database)
}

func (u *SpannerDatabaseIamUpdater) DescribeResource() string {
	return fmt.Sprintf("Spanner Database: %s/%s/%s", u.project, u.instance, u.database)
}

func resourceManagerToSpannerPolicy(p *cloudresourcemanager.Policy) (*spanner.Policy, error) {
	out := &spanner.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a resourcemanager policy to a spanner policy: {{err}}", err)
	}
	return out, nil
}

func spannerToResourceManagerPolicy(p *spanner.Policy) (*cloudresourcemanager.Policy, error) {
	out := &cloudresourcemanager.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a spanner policy to a resourcemanager policy: {{err}}", err)
	}
	return out, nil
}

<<<<<<< HEAD:google/iam_spanner_database.go
type spannerDatabaseId struct {
=======
type SpannerDatabaseId struct {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	Project  string
	Instance string
	Database string
}

<<<<<<< HEAD:google/iam_spanner_database.go
func (s spannerDatabaseId) terraformId() string {
	return fmt.Sprintf("%s/%s/%s", s.Project, s.Instance, s.Database)
}

func (s spannerDatabaseId) parentProjectUri() string {
	return fmt.Sprintf("projects/%s", s.Project)
}

func (s spannerDatabaseId) parentInstanceUri() string {
	return fmt.Sprintf("%s/instances/%s", s.parentProjectUri(), s.Instance)
}

func (s spannerDatabaseId) databaseUri() string {
=======
func (s SpannerDatabaseId) TerraformId() string {
	return fmt.Sprintf("%s/%s/%s", s.Project, s.Instance, s.Database)
}

func (s SpannerDatabaseId) parentProjectUri() string {
	return fmt.Sprintf("projects/%s", s.Project)
}

func (s SpannerDatabaseId) parentInstanceUri() string {
	return fmt.Sprintf("%s/instances/%s", s.parentProjectUri(), s.Instance)
}

func (s SpannerDatabaseId) databaseUri() string {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/spanner/iam_spanner_database.go
	return fmt.Sprintf("%s/databases/%s", s.parentInstanceUri(), s.Database)
}
