/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceHealthMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_duplicate_monitors": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"authentication": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorAuthInfoSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"disable_quickstart": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"dns_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorDNSSchema(),
		},
		"external_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorExternalSchema(),
		},
		"failed_checks": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"http_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"https_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"imap_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorImapSchema(),
		},
		"imaps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorImapSchema(),
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"monitor_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pop3_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorPop3Schema(),
		},
		"pop3s_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorPop3Schema(),
		},
		"radius_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorRadiusSchema(),
		},
		"receive_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"send_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"sip_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSIPSchema(),
		},
		"smtp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSmtpSchema(),
		},
		"smtps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSmtpSchema(),
		},
		"successful_checks": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"tcp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorTcpSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"udp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorUdpSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHealthMonitorCreate,
		Read:   ResourceAviHealthMonitorRead,
		Update: resourceAviHealthMonitorUpdate,
		Delete: resourceAviHealthMonitorDelete,
		Schema: ResourceHealthMonitorSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHealthMonitorImporter,
		},
	}
}

func ResourceHealthMonitorImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHealthMonitorSchema()
	return ResourceImporter(d, m, "healthmonitor", s)
}

func ResourceAviHealthMonitorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := APIRead(d, meta, "healthmonitor", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHealthMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := APICreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "healthmonitor"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviHealthMonitorDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
