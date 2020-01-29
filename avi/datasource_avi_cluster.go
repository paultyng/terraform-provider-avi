/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviCluster() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceClusterNodeSchema(),
			},
			"rejoin_nodes_automatically": {
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
			"virtual_ip": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"cluster_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"up_since": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"progress": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
