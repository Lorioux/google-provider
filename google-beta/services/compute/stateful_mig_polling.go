<<<<<<< HEAD:google/stateful_mig_polling.go
package google
=======
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
<<<<<<< HEAD:google/stateful_mig_polling.go
)

// PerInstanceConfig needs both regular operation polling AND custom polling for deletion which is why this is not generated
func resourceComputePerInstanceConfigPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)
		userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// PerInstanceConfig needs both regular operation polling AND custom polling for deletion which is why this is not generated
func resourceComputePerInstanceConfigPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)
		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return nil, err
		}

<<<<<<< HEAD:google/stateful_mig_polling.go
		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/listPerInstanceConfigs")
=======
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/listPerInstanceConfigs")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return nil, err
		}

<<<<<<< HEAD:google/stateful_mig_polling.go
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		res, err := SendRequest(config, "POST", project, url, userAgent, nil)
=======
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, err
		}
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    url,
			UserAgent: userAgent,
		})
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return res, err
		}
		res, err = flattenNestedComputePerInstanceConfig(d, meta, res)
		if err != nil {
			return nil, err
		}

		// Returns nil res if nested object is not found
		return res, nil
	}
}

// RegionPerInstanceConfig needs both regular operation polling AND custom polling for deletion which is why this is not generated
<<<<<<< HEAD:google/stateful_mig_polling.go
func resourceComputeRegionPerInstanceConfigPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)
		userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
func resourceComputeRegionPerInstanceConfigPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)
		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return nil, err
		}

<<<<<<< HEAD:google/stateful_mig_polling.go
		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/listPerInstanceConfigs")
=======
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/listPerInstanceConfigs")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return nil, err
		}

<<<<<<< HEAD:google/stateful_mig_polling.go
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		res, err := SendRequest(config, "POST", project, url, userAgent, nil)
=======
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, err
		}
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    url,
			UserAgent: userAgent,
		})
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return res, err
		}
		res, err = flattenNestedComputeRegionPerInstanceConfig(d, meta, res)
		if err != nil {
			return nil, err
		}

		// Returns nil res if nested object is not found
		return res, nil
	}
}

// Returns an instance name in the form zones/{zone}/instances/{instance} for the managed
// instance matching the name of a PerInstanceConfig
<<<<<<< HEAD:google/stateful_mig_polling.go
func findInstanceName(d *schema.ResourceData, config *Config) (string, error) {
	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/listManagedInstances")
=======
func findInstanceName(d *schema.ResourceData, config *transport_tpg.Config) (string, error) {
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/listManagedInstances")
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
	if err != nil {
		return "", err
	}

<<<<<<< HEAD:google/stateful_mig_polling.go
	userAgent, err := generateUserAgentString(d, config.UserAgent)
=======
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
	if err != nil {
		return "", err
	}

<<<<<<< HEAD:google/stateful_mig_polling.go
	project, err := getProject(d, config)
=======
	project, err := tpgresource.GetProject(d, config)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
	if err != nil {
		return "", err
	}
	instanceNameToFind := fmt.Sprintf("/%s", d.Get("name").(string))

	token := ""
	for paginate := true; paginate; {
		urlWithToken := ""
		if token != "" {
			urlWithToken = fmt.Sprintf("%s?maxResults=1&pageToken=%s", url, token)
		} else {
			urlWithToken = fmt.Sprintf("%s?maxResults=1", url)
		}
<<<<<<< HEAD:google/stateful_mig_polling.go
		res, err := SendRequest(config, "POST", project, urlWithToken, userAgent, nil)
=======
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   project,
			RawURL:    urlWithToken,
			UserAgent: userAgent,
		})
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
		if err != nil {
			return "", err
		}

		managedInstances, ok := res["managedInstances"]
		if !ok {
			return "", fmt.Errorf("Failed to parse response for listManagedInstances for %s", d.Id())
		}

		managedInstancesArr := managedInstances.([]interface{})
		for _, managedInstanceRaw := range managedInstancesArr {
			instance := managedInstanceRaw.(map[string]interface{})
			name, ok := instance["instance"]
			if !ok {
				return "", fmt.Errorf("Failed to read instance name for managed instance: %#v", instance)
			}
			if strings.HasSuffix(name.(string), instanceNameToFind) {
				return name.(string), nil
			}
		}

		tokenRaw, paginate := res["nextPageToken"]
		if paginate {
			token = tokenRaw.(string)
		}
	}

	return "", fmt.Errorf("Failed to find managed instance with name: %s", instanceNameToFind)
}

<<<<<<< HEAD:google/stateful_mig_polling.go
func PollCheckInstanceConfigDeleted(resp map[string]interface{}, respErr error) PollResult {
	if respErr != nil {
		return ErrorPollResult(respErr)
=======
func PollCheckInstanceConfigDeleted(resp map[string]interface{}, respErr error) transport_tpg.PollResult {
	if respErr != nil {
		return transport_tpg.ErrorPollResult(respErr)
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
	}

	// Nested object 404 appears as nil response
	if resp == nil {
		// Config no longer exists
<<<<<<< HEAD:google/stateful_mig_polling.go
		return SuccessPollResult()
=======
		return transport_tpg.SuccessPollResult()
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
	}

	// Read status
	status := resp["status"].(string)
	if status == "DELETING" {
<<<<<<< HEAD:google/stateful_mig_polling.go
		return PendingStatusPollResult("Still deleting")
	}
	return ErrorPollResult(fmt.Errorf("Expected PerInstanceConfig to be deleting but status is: %s", status))
=======
		return transport_tpg.PendingStatusPollResult("Still deleting")
	}
	return transport_tpg.ErrorPollResult(fmt.Errorf("Expected PerInstanceConfig to be deleting but status is: %s", status))
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/services/compute/stateful_mig_polling.go
}
