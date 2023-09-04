// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package kms

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/iam_kms_key_ring.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IamKmsKeyRingSchema = map[string]*schema.Schema{
	"key_ring_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
}

type KmsKeyRingIamUpdater struct {
	resourceId string
<<<<<<< HEAD:google/iam_kms_key_ring.go
	d          TerraformResourceData
	Config     *Config
}

func NewKmsKeyRingIamUpdater(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
=======
	d          tpgresource.TerraformResourceData
	Config     *transport_tpg.Config
}

func NewKmsKeyRingIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
	keyRing := d.Get("key_ring_id").(string)
	keyRingId, err := parseKmsKeyRingId(keyRing, config)

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error parsing resource ID for %s: {{err}}", keyRing), err)
	}

	return &KmsKeyRingIamUpdater{
<<<<<<< HEAD:google/iam_kms_key_ring.go
		resourceId: keyRingId.keyRingId(),
=======
		resourceId: keyRingId.KeyRingId(),
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
		d:          d,
		Config:     config,
	}, nil
}

func KeyRingIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	keyRingId, err := parseKmsKeyRingId(d.Id(), config)
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/iam_kms_key_ring.go
	if err := d.Set("key_ring_id", keyRingId.keyRingId()); err != nil {
		return fmt.Errorf("Error setting key_ring_id: %s", err)
	}
	d.SetId(keyRingId.keyRingId())
=======
	if err := d.Set("key_ring_id", keyRingId.KeyRingId()); err != nil {
		return fmt.Errorf("Error setting key_ring_id: %s", err)
	}
	d.SetId(keyRingId.KeyRingId())
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
	return nil
}

func (u *KmsKeyRingIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
<<<<<<< HEAD:google/iam_kms_key_ring.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD:google/iam_kms_key_ring.go
	p, err := u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.GetIamPolicy(u.resourceId).OptionsRequestedPolicyVersion(IamPolicyVersion).Do()
=======
	p, err := u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.GetIamPolicy(u.resourceId).OptionsRequestedPolicyVersion(tpgiamresource.IamPolicyVersion).Do()
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	cloudResourcePolicy, err := kmsToResourceManagerPolicy(p)

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return cloudResourcePolicy, nil
}

func (u *KmsKeyRingIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	kmsPolicy, err := resourceManagerToKmsPolicy(policy)

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

<<<<<<< HEAD:google/iam_kms_key_ring.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_key_ring.go
	if err != nil {
		return err
	}

	_, err = u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.SetIamPolicy(u.resourceId, &cloudkms.SetIamPolicyRequest{
		Policy: kmsPolicy,
	}).Do()

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *KmsKeyRingIamUpdater) GetResourceId() string {
	return u.resourceId
}

func (u *KmsKeyRingIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-kms-key-ring-%s", u.resourceId)
}

func (u *KmsKeyRingIamUpdater) DescribeResource() string {
	return fmt.Sprintf("KMS KeyRing %q", u.resourceId)
}

func resourceManagerToKmsPolicy(p *cloudresourcemanager.Policy) (*cloudkms.Policy, error) {
	out := &cloudkms.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a v1 policy to a kms policy: {{err}}", err)
	}
	return out, nil
}

func kmsToResourceManagerPolicy(p *cloudkms.Policy) (*cloudresourcemanager.Policy, error) {
	out := &cloudresourcemanager.Policy{}
	err := tpgresource.Convert(p, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a kms policy to a v1 policy: {{err}}", err)
	}
	return out, nil
}
