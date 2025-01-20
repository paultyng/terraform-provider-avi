// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviFileObject() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviFileObjectRead,
		Schema: map[string]*schema.Schema{
			"checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"child_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"compressed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crl_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCRLSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceFileObjectEventMapSchema(),
			},
			"expires_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gslb_geodb_format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_parent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_version_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"read_only": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"restrict_download": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
