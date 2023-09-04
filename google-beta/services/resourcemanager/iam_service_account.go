// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import (
	"fmt"

<<<<<<< HEAD:google/iam_service_account.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"

>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

var IamServiceAccountSchema = map[string]*schema.Schema{
	"service_account_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: verify.ValidateRegexp(verify.ServiceAccountLinkRegex),
	},
}

type ServiceAccountIamUpdater struct {
	serviceAccountId string
<<<<<<< HEAD:google/iam_service_account.go
	d                TerraformResourceData
	Config           *Config
}

func NewServiceAccountIamUpdater(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
=======
	d                tpgresource.TerraformResourceData
	Config           *transport_tpg.Config
}

func NewServiceAccountIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go
	return &ServiceAccountIamUpdater{
		serviceAccountId: d.Get("service_account_id").(string),
		d:                d,
		Config:           config,
	}, nil
}

<<<<<<< HEAD:google/iam_service_account.go
func ServiceAccountIdParseFunc(d *schema.ResourceData, _ *Config) error {
=======
func ServiceAccountIdParseFunc(d *schema.ResourceData, _ *transport_tpg.Config) error {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go
	if err := d.Set("service_account_id", d.Id()); err != nil {
		return fmt.Errorf("Error setting service_account_id: %s", err)
	}
	return nil
}

func (u *ServiceAccountIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
<<<<<<< HEAD:google/iam_service_account.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD:google/iam_service_account.go
	p, err := u.Config.NewIamClient(userAgent).Projects.ServiceAccounts.GetIamPolicy(u.serviceAccountId).OptionsRequestedPolicyVersion(IamPolicyVersion).Do()
=======
	p, err := u.Config.NewIamClient(userAgent).Projects.ServiceAccounts.GetIamPolicy(u.serviceAccountId).OptionsRequestedPolicyVersion(tpgiamresource.IamPolicyVersion).Do()
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	cloudResourcePolicy, err := iamToResourceManagerPolicy(p)
	if err != nil {
		return nil, err
	}

	return cloudResourcePolicy, nil
}

func (u *ServiceAccountIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	iamPolicy, err := resourceManagerToIamPolicy(policy)
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/iam_service_account.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/iam_service_account.go
	if err != nil {
		return err
	}

	_, err = u.Config.NewIamClient(userAgent).Projects.ServiceAccounts.SetIamPolicy(u.GetResourceId(), &iam.SetIamPolicyRequest{
		Policy: iamPolicy,
	}).Do()

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ServiceAccountIamUpdater) GetResourceId() string {
	return u.serviceAccountId
}

func (u *ServiceAccountIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-service-account-%s", u.serviceAccountId)
}

func (u *ServiceAccountIamUpdater) DescribeResource() string {
	return fmt.Sprintf("service account '%s'", u.serviceAccountId)
}

func resourceManagerToIamPolicy(p *cloudresourcemanager.Policy) (*iam.Policy, error) {
	out := &iam.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a v1 policy to a iam policy: {{err}}", err)
	}
	return out, nil
}

func iamToResourceManagerPolicy(p *iam.Policy) (*cloudresourcemanager.Policy, error) {
	out := &cloudresourcemanager.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a iam policy to a v1 policy: {{err}}", err)
	}
	return out, nil
}
