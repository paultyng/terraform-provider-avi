package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIErrorPageBodyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIErrorPageBodyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIErrorPageBodyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIErrorPageBodyExists("avi_errorpagebody.testErrorPageBody"),
					resource.TestCheckResourceAttr(
						"avi_errorpagebody.testErrorPageBody", "name", "test-Custom-Error-Page-abc"),
				),
			},
			{
				Config: testAccAVIErrorPageBodyupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIErrorPageBodyExists("avi_errorpagebody.testErrorPageBody"),
					resource.TestCheckResourceAttr(
						"avi_errorpagebody.testErrorPageBody", "name", "test-Custom-Error-Page-updated"),
				),
			},
			{
				ResourceName:      "avi_errorpagebody.testErrorPageBody",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIErrorPageBodyConfig,
			},
		},
	})

}

func testAccCheckAVIErrorPageBodyExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ErrorPageBody ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIErrorPageBodyDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_errorpagebody" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI ErrorPageBody still exists")
		}
	}
	return nil
}

const testAccAVIErrorPageBodyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpagebody" "testErrorPageBody" {
	name = "test-Custom-Error-Page-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	error_page_body = "<!DOCTYPE html><html><head></head><body><div><p> Please contact our technical support</p></div></body></html>"
}
`

const testAccAVIErrorPageBodyupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpagebody" "testErrorPageBody" {
	name = "test-Custom-Error-Page-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	error_page_body = "<!DOCTYPE html><html><head></head><body><div><p> Please contact our technical support team</p></div></body></html>"
}
`
