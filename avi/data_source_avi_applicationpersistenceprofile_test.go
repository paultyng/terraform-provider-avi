package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceApplicationPersistenceProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSApplicationPersistenceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "name", "test-System-Persistence-Client-IP-abc"),
					resource.TestCheckResourceAttr(
						"avi_applicationpersistenceprofile.testApplicationPersistenceProfile", "is_federated", "false"),
				),
			},
		},
	})

}

const testAccAVIDSApplicationPersistenceProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_applicationpersistenceprofile" "testApplicationPersistenceProfile" {
	name = "test-System-Persistence-Client-IP-abc"
	is_federated = false
	tenant_ref = data.avi_tenant.default_tenant.id
	persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
	server_hm_down_recovery = "HM_DOWN_PICK_NEW_SERVER"
	ip_persistence_profile {
		ip_persistent_timeout = "5"
	}
}

data "avi_applicationpersistenceprofile" "testApplicationPersistenceProfile" {
    name= "${avi_applicationpersistenceprofile.testApplicationPersistenceProfile.name}"
}
`
