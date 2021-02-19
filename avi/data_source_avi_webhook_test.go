package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceWebhookBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSWebhookConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_webhook.testWebhook", "name", "test-wb-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSWebhookConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_webhook" "testWebhook" {
	name = "test-wb-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_webhook" "testWebhook" {
    name= "${avi_webhook.testWebhook.name}"
}
`
