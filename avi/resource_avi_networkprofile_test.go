package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAVINetworkProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVINetworkProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVINetworkProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkProfileExists("avi_networkprofile.testNetworkProfile"),
					resource.TestCheckResourceAttr(
						"avi_networkprofile.testNetworkProfile", "name", "test-System-TCP-Proxy-abc"),
				),
			},
			{
				Config: testAccAVINetworkProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVINetworkProfileExists("avi_networkprofile.testNetworkProfile"),
					resource.TestCheckResourceAttr(
						"avi_networkprofile.testNetworkProfile", "name", "test-System-TCP-Proxy-updated"),
				),
			},
			{
				ResourceName:      "avi_networkprofile.testNetworkProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVINetworkProfileConfig,
			},
		},
	})

}

func testAccCheckAVINetworkProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI NetworkProfile ID is set")
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

func testAccCheckAVINetworkProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_networkprofile" {
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
			return fmt.Errorf("AVI NetworkProfile still exists")
		}
	}
	return nil
}

const testAccAVINetworkProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networkprofile" "testNetworkProfile" {
	name = "test-System-TCP-Proxy-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	profile {
		tcp_proxy_profile {
			receive_window = "64"
			time_wait_delay = "2000"
			cc_algo = "CC_ALGO_NEW_RENO"
			nagles_algorithm = false
			max_syn_retransmissions = "8"
			ignore_time_wait = false
			idle_connection_type = "KEEP_ALIVE"
			aggressive_congestion_avoidance = false
			idle_connection_timeout = "600"
			max_retransmissions = "8"
			ip_dscp = "0"
			automatic = true
			use_interface_mtu = true
			reorder_threshold = "10"
			min_rexmt_timeout = "50"
		}
		type = "PROTOCOL_TYPE_TCP_PROXY"
	}
}
`

const testAccAVINetworkProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networkprofile" "testNetworkProfile" {
	name = "test-System-TCP-Proxy-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	profile {
		tcp_proxy_profile {
			receive_window = "64"
			time_wait_delay = "2000"
			cc_algo = "CC_ALGO_NEW_RENO"
			nagles_algorithm = false
			max_syn_retransmissions = "8"
			ignore_time_wait = false
			idle_connection_type = "KEEP_ALIVE"
			aggressive_congestion_avoidance = false
			idle_connection_timeout = "600"
			max_retransmissions = "8"
			ip_dscp = "0"
			automatic = true
			use_interface_mtu = true
			reorder_threshold = "10"
			min_rexmt_timeout = "50"
		}
		type = "PROTOCOL_TYPE_TCP_PROXY"
	}
}
`
