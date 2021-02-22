package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVIPoolBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIPoolConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testPool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "name", "test-Pool-abc"),
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "ignore_servers", "false"),
				),
			},
			{
				Config: testAccAVIPoolupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIPoolExists("avi_pool.testPool"),
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "name", "test-Pool-updated"),
					resource.TestCheckResourceAttr(
						"avi_pool.testPool", "ignore_servers", "false"),
				),
			},
			{
				ResourceName:      "avi_pool.testPool",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIPoolConfig,
			},
		},
	})

}

func testAccCheckAVIPoolExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI Pool ID is set")
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

func testAccCheckAVIPoolDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_pool" {
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
			return fmt.Errorf("AVI Pool still exists")
		}
	}
	return nil
}

const testAccAVIPoolConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
data "avi_healthmonitor" "default_monitor" {
    name= "System-HTTP"
}
resource "avi_pool" "testPool" {
	name = "test-Pool-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	ignore_servers = false
	health_monitor_refs = [data.avi_healthmonitor.default_monitor.id]
	servers {
	ratio = "1"
	ip {
		type = "V4"
		addr = "10.90.64.66"
	}
	hostname = "10.90.64.66"
	port = "8080"
	enabled = true
}
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
}
resource "avi_pool" "testPoolHmRef" {
        name = "test-pool-hm-ref"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	ignore_servers = false
	servers {
	ratio = "1"
	ip {
		type = "V4"
		addr = "10.90.64.66"
	}
	hostname = "10.90.64.66"
	port = "8080"
	enabled = true
}
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
}
`

const testAccAVIPoolupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
data "avi_healthmonitor" "default_monitor" {
    name= "System-HTTP"
}
resource "avi_pool" "testPool" {
	name = "test-Pool-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	ignore_servers = false
	health_monitor_refs = [data.avi_healthmonitor.default_monitor.id]
	servers {
	ratio = "1"
	ip {
		type = "V4"
		addr = "10.90.64.66"
	}
	hostname = "10.90.64.66"
	port = "8080"
	enabled = true
}
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
}
resource "avi_pool" "testPoolHmRef" {
        name = "test-pool-hm-ref-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	ignore_servers = false
	servers {
	ratio = "1"
	ip {
		type = "V4"
		addr = "10.90.64.66"
	}
	hostname = "10.90.64.66"
	port = "8080"
	enabled = true
}
	fail_action {
		type = "FAIL_ACTION_CLOSE_CONN"
	}
}
`
