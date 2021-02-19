package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceStringGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSStringGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_stringgroup.testStringGroup", "name", "test-System-Compressible-Content-Types-abc"),
				),
			},
		},
	})

}

const testAccAVIDSStringGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_stringgroup" "testStringGroup" {
	name = "test-System-Compressible-Content-Types-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	kv {
	key = "text/html"
}
kv {
	key = "text/xml"
}
kv {
	key = "text/plain"
}
kv {
	key = "text/css"
}
kv {
	key = "text/javascript"
}
kv {
	key = "application/javascript"
}
kv {
	key = "application/x-javascript"
}
kv {
	key = "application/xml"
}
kv {
	key = "application/pdf"
}
	type = "SG_TYPE_STRING"
}

data "avi_stringgroup" "testStringGroup" {
    name= "${avi_stringgroup.testStringGroup.name}"
}
`
