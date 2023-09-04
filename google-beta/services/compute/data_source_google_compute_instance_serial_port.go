<<<<<<< HEAD:google/data_source_google_compute_instance_serial_port.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_serial_port.go

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/data_source_google_compute_instance_serial_port.go
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_serial_port.go
)

func DataSourceGoogleComputeInstanceSerialPort() *schema.Resource {
	return &schema.Resource{
		Read: computeInstanceSerialPortRead,
		Schema: map[string]*schema.Schema{
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func computeInstanceSerialPortRead(d *schema.ResourceData, meta interface{}) error {
<<<<<<< HEAD:google/data_source_google_compute_instance_serial_port.go
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_serial_port.go
	if err != nil {
		return err
	}

<<<<<<< HEAD:google/data_source_google_compute_instance_serial_port.go
	project, err := getProject(d, config)
=======
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_serial_port.go
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
<<<<<<< HEAD:google/data_source_google_compute_instance_serial_port.go
	zone, err := getZone(d, config)
=======
	zone, err := tpgresource.GetZone(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/data_source_google_compute_instance_serial_port.go
	if err != nil {
		return err
	}
	if err := d.Set("zone", zone); err != nil {
		return fmt.Errorf("Error setting zone: %s", err)
	}

	port := int64(d.Get("port").(int))
	output, err := config.NewComputeClient(userAgent).Instances.GetSerialPortOutput(project, zone, d.Get("instance").(string)).Port(port).Do()
	if err != nil {
		return err
	}

	if err := d.Set("contents", output.Contents); err != nil {
		return fmt.Errorf("Error setting contents: %s", err)
	}
	d.SetId(output.SelfLink)
	return nil
}
