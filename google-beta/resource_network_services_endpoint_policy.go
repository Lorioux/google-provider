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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkServicesEndpointPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesEndpointPolicyCreate,
		Read:   resourceNetworkServicesEndpointPolicyRead,
		Update: resourceNetworkServicesEndpointPolicyUpdate,
		Delete: resourceNetworkServicesEndpointPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesEndpointPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"endpoint_matcher": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Required. A matcher that selects endpoints to which the policies should be applied.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata_label_matcher": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The matcher is based on node metadata presented by xDS clients.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metadata_label_match_criteria": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validateEnum([]string{"MATCH_ANY", "MATCH_ALL"}),
										Description:  `Specifies how matching should be done. Possible values: ["MATCH_ANY", "MATCH_ALL"]`,
									},
									"metadata_labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"label_name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `Required. Label name presented as key in xDS Node Metadata.`,
												},
												"label_value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `Required. Label value presented as value corresponding to the above key, in xDS Node Metadata.`,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the EndpointPolicy resource.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEnum([]string{"SIDECAR_PROXY", "GRPC_SERVER"}),
				Description:  `The type of endpoint policy. This is primarily used to validate the configuration. Possible values: ["SIDECAR_PROXY", "GRPC_SERVER"]`,
			},
			"authorization_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints.`,
			},
			"client_tls_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the TcpRoute resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"server_tls_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends.`,
			},
			"traffic_port_selector": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `List of ports. Can be port numbers or port range (example, [80-90] specifies all ports from 80 to 90, including 80 and 90) or named ports or * to specify all ports. If the list is empty, all ports are selected.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the TcpRoute was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the TcpRoute was updated in UTC.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkServicesEndpointPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkServicesEndpointPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkServicesEndpointPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	authorizationPolicyProp, err := expandNetworkServicesEndpointPolicyAuthorizationPolicy(d.Get("authorization_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorization_policy"); !isEmptyValue(reflect.ValueOf(authorizationPolicyProp)) && (ok || !reflect.DeepEqual(v, authorizationPolicyProp)) {
		obj["authorizationPolicy"] = authorizationPolicyProp
	}
	serverTlsPolicyProp, err := expandNetworkServicesEndpointPolicyServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !isEmptyValue(reflect.ValueOf(serverTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	clientTlsPolicyProp, err := expandNetworkServicesEndpointPolicyClientTlsPolicy(d.Get("client_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_tls_policy"); !isEmptyValue(reflect.ValueOf(clientTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, clientTlsPolicyProp)) {
		obj["clientTlsPolicy"] = clientTlsPolicyProp
	}
	typeProp, err := expandNetworkServicesEndpointPolicyType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	trafficPortSelectorProp, err := expandNetworkServicesEndpointPolicyTrafficPortSelector(d.Get("traffic_port_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("traffic_port_selector"); !isEmptyValue(reflect.ValueOf(trafficPortSelectorProp)) && (ok || !reflect.DeepEqual(v, trafficPortSelectorProp)) {
		obj["trafficPortSelector"] = trafficPortSelectorProp
	}
	endpointMatcherProp, err := expandNetworkServicesEndpointPolicyEndpointMatcher(d.Get("endpoint_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint_matcher"); !isEmptyValue(reflect.ValueOf(endpointMatcherProp)) && (ok || !reflect.DeepEqual(v, endpointMatcherProp)) {
		obj["endpointMatcher"] = endpointMatcherProp
	}

	url, err := ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/endpointPolicies?endpointPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EndpointPolicy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EndpointPolicy: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/endpointPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating EndpointPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EndpointPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EndpointPolicy %q: %#v", d.Id(), res)

	return resourceNetworkServicesEndpointPolicyRead(d, meta)
}

func resourceNetworkServicesEndpointPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/endpointPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NetworkServicesEndpointPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkServicesEndpointPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesEndpointPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesEndpointPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesEndpointPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("authorization_policy", flattenNetworkServicesEndpointPolicyAuthorizationPolicy(res["authorizationPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("server_tls_policy", flattenNetworkServicesEndpointPolicyServerTlsPolicy(res["serverTlsPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("client_tls_policy", flattenNetworkServicesEndpointPolicyClientTlsPolicy(res["clientTlsPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("type", flattenNetworkServicesEndpointPolicyType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("traffic_port_selector", flattenNetworkServicesEndpointPolicyTrafficPortSelector(res["trafficPortSelector"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}
	if err := d.Set("endpoint_matcher", flattenNetworkServicesEndpointPolicyEndpointMatcher(res["endpointMatcher"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointPolicy: %s", err)
	}

	return nil
}

func resourceNetworkServicesEndpointPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkServicesEndpointPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkServicesEndpointPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	authorizationPolicyProp, err := expandNetworkServicesEndpointPolicyAuthorizationPolicy(d.Get("authorization_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorization_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorizationPolicyProp)) {
		obj["authorizationPolicy"] = authorizationPolicyProp
	}
	serverTlsPolicyProp, err := expandNetworkServicesEndpointPolicyServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	clientTlsPolicyProp, err := expandNetworkServicesEndpointPolicyClientTlsPolicy(d.Get("client_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_tls_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientTlsPolicyProp)) {
		obj["clientTlsPolicy"] = clientTlsPolicyProp
	}
	typeProp, err := expandNetworkServicesEndpointPolicyType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	trafficPortSelectorProp, err := expandNetworkServicesEndpointPolicyTrafficPortSelector(d.Get("traffic_port_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("traffic_port_selector"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trafficPortSelectorProp)) {
		obj["trafficPortSelector"] = trafficPortSelectorProp
	}
	endpointMatcherProp, err := expandNetworkServicesEndpointPolicyEndpointMatcher(d.Get("endpoint_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint_matcher"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, endpointMatcherProp)) {
		obj["endpointMatcher"] = endpointMatcherProp
	}

	url, err := ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/endpointPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EndpointPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("authorization_policy") {
		updateMask = append(updateMask, "authorizationPolicy")
	}

	if d.HasChange("server_tls_policy") {
		updateMask = append(updateMask, "serverTlsPolicy")
	}

	if d.HasChange("client_tls_policy") {
		updateMask = append(updateMask, "clientTlsPolicy")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("traffic_port_selector") {
		updateMask = append(updateMask, "trafficPortSelector")
	}

	if d.HasChange("endpoint_matcher") {
		updateMask = append(updateMask, "endpointMatcher")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating EndpointPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating EndpointPolicy %q: %#v", d.Id(), res)
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Updating EndpointPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkServicesEndpointPolicyRead(d, meta)
}

func resourceNetworkServicesEndpointPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointPolicy: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/endpointPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EndpointPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EndpointPolicy")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting EndpointPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EndpointPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesEndpointPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/endpointPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/endpointPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesEndpointPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyAuthorizationPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyServerTlsPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyClientTlsPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyTrafficPortSelector(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ports"] =
		flattenNetworkServicesEndpointPolicyTrafficPortSelectorPorts(original["ports"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesEndpointPolicyTrafficPortSelectorPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyEndpointMatcher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["metadata_label_matcher"] =
		flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher(original["metadataLabelMatcher"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["metadata_label_match_criteria"] =
		flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteria(original["metadataLabelMatchCriteria"], d, config)
	transformed["metadata_labels"] =
		flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(original["metadataLabels"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteria(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"label_name":  flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelName(original["labelName"], d, config),
			"label_value": flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelValue(original["labelValue"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkServicesEndpointPolicyLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesEndpointPolicyDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyAuthorizationPolicy(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyServerTlsPolicy(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyClientTlsPolicy(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyTrafficPortSelector(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPorts, err := expandNetworkServicesEndpointPolicyTrafficPortSelectorPorts(original["ports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
		transformed["ports"] = transformedPorts
	}

	return transformed, nil
}

func expandNetworkServicesEndpointPolicyTrafficPortSelectorPorts(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcher(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetadataLabelMatcher, err := expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher(original["metadata_label_matcher"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadataLabelMatcher); val.IsValid() && !isEmptyValue(val) {
		transformed["metadataLabelMatcher"] = transformedMetadataLabelMatcher
	}

	return transformed, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetadataLabelMatchCriteria, err := expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteria(original["metadata_label_match_criteria"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadataLabelMatchCriteria); val.IsValid() && !isEmptyValue(val) {
		transformed["metadataLabelMatchCriteria"] = transformedMetadataLabelMatchCriteria
	}

	transformedMetadataLabels, err := expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(original["metadata_labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadataLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["metadataLabels"] = transformedMetadataLabels
	}

	return transformed, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteria(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedLabelName, err := expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelName(original["label_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabelName); val.IsValid() && !isEmptyValue(val) {
			transformed["labelName"] = transformedLabelName
		}

		transformedLabelValue, err := expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelValue(original["label_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabelValue); val.IsValid() && !isEmptyValue(val) {
			transformed["labelValue"] = transformedLabelValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsLabelValue(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
