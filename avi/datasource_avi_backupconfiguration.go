/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviBackupConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBackupConfigurationRead,
		Schema: map[string]*schema.Schema{
			"aws_access_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_bucket_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_secret_access": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_file_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maximum_backups_stored": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"save_local": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ssh_user_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_to_remote_host": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"upload_to_s3": {
				Type:     schema.TypeBool,
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
