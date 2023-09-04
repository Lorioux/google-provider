// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package kms

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/iam_kms_crypto_key.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IamKmsCryptoKeySchema = map[string]*schema.Schema{
	"crypto_key_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
}

type KmsCryptoKeyIamUpdater struct {
	resourceId string
<<<<<<< HEAD:google/iam_kms_crypto_key.go
	d          TerraformResourceData
	Config     *Config
}

func NewKmsCryptoKeyIamUpdater(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
=======
	d          tpgresource.TerraformResourceData
	Config     *transport_tpg.Config
}

func NewKmsCryptoKeyIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
	cryptoKey := d.Get("crypto_key_id").(string)
	cryptoKeyId, err := ParseKmsCryptoKeyId(cryptoKey, config)

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error parsing resource ID for %s: {{err}}", cryptoKey), err)
	}

	return &KmsCryptoKeyIamUpdater{
<<<<<<< HEAD:google/iam_kms_crypto_key.go
		resourceId: cryptoKeyId.cryptoKeyId(),
=======
		resourceId: cryptoKeyId.CryptoKeyId(),
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
		d:          d,
		Config:     config,
	}, nil
}

func CryptoIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	cryptoKeyId, err := ParseKmsCryptoKeyId(d.Id(), config)
	if err != nil {
		return err
	}
<<<<<<< HEAD:google/iam_kms_crypto_key.go
	if err := d.Set("crypto_key_id", cryptoKeyId.cryptoKeyId()); err != nil {
		return fmt.Errorf("Error setting crypto_key_id: %s", err)
	}
	d.SetId(cryptoKeyId.cryptoKeyId())
=======
	if err := d.Set("crypto_key_id", cryptoKeyId.CryptoKeyId()); err != nil {
		return fmt.Errorf("Error setting crypto_key_id: %s", err)
	}
	d.SetId(cryptoKeyId.CryptoKeyId())
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
	return nil
}

func (u *KmsCryptoKeyIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
<<<<<<< HEAD:google/iam_kms_crypto_key.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD:google/iam_kms_crypto_key.go
	p, err := u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(u.resourceId).OptionsRequestedPolicyVersion(IamPolicyVersion).Do()
=======
	p, err := u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.CryptoKeys.GetIamPolicy(u.resourceId).OptionsRequestedPolicyVersion(tpgiamresource.IamPolicyVersion).Do()
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	cloudResourcePolicy, err := kmsToResourceManagerPolicy(p)

	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return cloudResourcePolicy, nil
}

func (u *KmsCryptoKeyIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
<<<<<<< HEAD:google/iam_kms_crypto_key.go
	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/kms/iam_kms_crypto_key.go
	if err != nil {
		return err
	}

	kmsPolicy, err := resourceManagerToKmsPolicy(policy)

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Invalid IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	_, err = u.Config.NewKmsClient(userAgent).Projects.Locations.KeyRings.CryptoKeys.SetIamPolicy(u.resourceId, &cloudkms.SetIamPolicyRequest{
		Policy: kmsPolicy,
	}).Do()

	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *KmsCryptoKeyIamUpdater) GetResourceId() string {
	return u.resourceId
}

func (u *KmsCryptoKeyIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-kms-crypto-key-%s", u.resourceId)
}

func (u *KmsCryptoKeyIamUpdater) DescribeResource() string {
	return fmt.Sprintf("KMS CryptoKey %q", u.resourceId)
}
