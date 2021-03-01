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

func ResourceNetworkProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection_mirror": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"profile": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceNetworkProfileUnionSchema(),
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
	}
}

func resourceAviNetworkProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkProfileCreate,
		Read:   ResourceAviNetworkProfileRead,
		Update: resourceAviNetworkProfileUpdate,
		Delete: resourceAviNetworkProfileDelete,
		Schema: ResourceNetworkProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkProfileImporter,
		},
	}
}

func ResourceNetworkProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkProfileSchema()
	return ResourceImporter(d, m, "networkprofile", s)
}

func ResourceAviNetworkProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := APIRead(d, meta, "networkprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := APICreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networkprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
