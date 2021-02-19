package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceAlertEmailConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAlertEmailConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testAlertEmailConfig", "name", "test-aec-abc"),
				),
			},
		},
	})

}

const testAccAVIDSAlertEmailConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_alertemailconfig" "testAlertEmailConfig" {
	name = "test-aec-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	cc_emails = "admin@avicontroller.net"
	description = "test alert email"
	to_emails = "admin@avicontroller.net"
}

data "avi_alertemailconfig" "testAlertEmailConfig" {
    name= "${avi_alertemailconfig.testAlertEmailConfig.name}"
}
`
