/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVrfContextRead,
		Schema: map[string]*schema.Schema{
			"attrs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"bfd_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBfdProfileSchema(),
			},
			"bgp_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBgpProfileSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"debugvrfcontext": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDebugVrfContextSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mon": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGatewayMonitorSchema(),
			},
			"internal_gateway_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceInternalGatewayMonitorSchema(),
			},
			"lldp_enable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceStaticRouteSchema(),
			},
			"system_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
