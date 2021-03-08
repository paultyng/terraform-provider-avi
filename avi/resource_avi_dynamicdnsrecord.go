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

//nolint
func ResourceDynamicDnsRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DNS_RECORD_RESPONSE_ROUND_ROBIN",
		},
		"cname": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDnsCnameRdataSchema(),
		},
		"delegated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_vs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"fqdn": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip6_address": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsAAAARdataSchema(),
		},
		"ip_address": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsARdataSchema(),
		},
		"metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"mx_records": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsMxRdataSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ns": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsNsRdataSchema(),
		},
		"num_records_in_response": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"service_locators": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsSrvRdataSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"txt_records": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsTxtRdataSchema(),
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"wildcard_match": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

//nolint
func resourceAviDynamicDnsRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviDynamicDnsRecordCreate,
		Read:   ResourceAviDynamicDnsRecordRead,
		Update: resourceAviDynamicDnsRecordUpdate,
		Delete: resourceAviDynamicDnsRecordDelete,
		Schema: ResourceDynamicDnsRecordSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceDynamicDnsRecordImporter,
		},
	}
}

//nolint
func ResourceDynamicDnsRecordImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceDynamicDnsRecordSchema()
	return ResourceImporter(d, m, "dynamicdnsrecord", s)
}

//nolint
func ResourceAviDynamicDnsRecordRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDynamicDnsRecordSchema()
	err := APIRead(d, meta, "dynamicdnsrecord", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

//nolint
func resourceAviDynamicDnsRecordCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDynamicDnsRecordSchema()
	err := APICreateOrUpdate(d, meta, "dynamicdnsrecord", s)
	if err == nil {
		err = ResourceAviDynamicDnsRecordRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDynamicDnsRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDynamicDnsRecordSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "dynamicdnsrecord", s)
	if err == nil {
		err = ResourceAviDynamicDnsRecordRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDynamicDnsRecordDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "dynamicdnsrecord"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviDynamicDnsRecordDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
