package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceServerAutoScalePolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSServerAutoScalePolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_serverautoscalepolicy.testServerAutoScalePolicy", "name", "test-ssp-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSServerAutoScalePolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_serverautoscalepolicy" "testServerAutoScalePolicy" {
	name = "test-ssp-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	scalein_alertconfig_refs = []
	scaleout_alertconfig_refs = []
}

data "avi_serverautoscalepolicy" "testServerAutoScalePolicy" {
    name= "${avi_serverautoscalepolicy.testServerAutoScalePolicy.name}"
}
`
