/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviAlertConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAlertConfigRead,
		Schema: map[string]*schema.Schema{
			"action_group_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alert_rule": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAlertRuleSchema(),
			},
			"autoscale_alert": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"expiry_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recommendation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rolling_window": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"throttle": {
				Type:     schema.TypeInt,
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
