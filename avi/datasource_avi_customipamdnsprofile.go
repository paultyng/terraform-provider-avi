/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

//nolint
func dataSourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCustomIpamDnsProfileRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_params": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"script_uri": {
				Type:     schema.TypeString,
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
