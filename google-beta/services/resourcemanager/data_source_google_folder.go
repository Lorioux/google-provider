// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_folder.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_folder.go
)

func DataSourceGoogleFolder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFolderRead,
		Schema: map[string]*schema.Schema{
			"folder": {
				Type:     schema.TypeString,
				Required: true,
			},
			"folder_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_organization": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"organization": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFolderRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_folder.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_folder.go
	if err != nil {
		return err
	}

	d.SetId(canonicalFolderName(d.Get("folder").(string)))
	if err := resourceGoogleFolderRead(d, meta); err != nil {
		return err
	}
	// If resource doesn't exist, read will not set ID and we should return error.
	if d.Id() == "" {
		return nil
	}

	if v, ok := d.GetOk("lookup_organization"); ok && v.(bool) {
		organization, err := lookupOrganizationName(d.Id(), userAgent, d, config)
		if err != nil {
			return err
		}

		if err := d.Set("organization", organization); err != nil {
			return fmt.Errorf("Error setting organization: %s", err)
		}
	}

	return nil
}

func canonicalFolderName(ba string) string {
	if strings.HasPrefix(ba, "folders/") {
		return ba
	}

	return "folders/" + ba
}

<<<<<<< HEAD:google/data_source_google_folder.go
func lookupOrganizationName(parent, userAgent string, d *schema.ResourceData, config *Config) (string, error) {
=======
func lookupOrganizationName(parent, userAgent string, d *schema.ResourceData, config *transport_tpg.Config) (string, error) {
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/resourcemanager/data_source_google_folder.go
	if parent == "" || strings.HasPrefix(parent, "organizations/") {
		return parent, nil
	} else if strings.HasPrefix(parent, "folders/") {
		parentFolder, err := getGoogleFolder(parent, userAgent, d, config)
		if err != nil {
			return "", fmt.Errorf("Error getting parent folder '%s': %s", parent, err)
		}
		return lookupOrganizationName(parentFolder.Parent, userAgent, d, config)
	} else {
		return "", fmt.Errorf("Unknown parent type '%s' on folder '%s'", parent, d.Id())
	}
}
