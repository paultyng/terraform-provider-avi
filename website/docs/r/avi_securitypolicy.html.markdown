############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_securitypolicy"
sidebar_current: "docs-avi-resource-securitypolicy"
description: |-
  Creates and manages Avi SecurityPolicy.
---

# avi_securitypolicy

The SecurityPolicy resource allows the creation and management of Avi SecurityPolicy

## Example Usage

```hcl
resource "avi_securitypolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the security policy. Field introduced in 18.2.1.
* `description` - (Optional) Security policy is used to specify various configuration information used to perform distributed denial of service (ddos) attacks detection and mitigation. Field introduced in 18.2.1.
* `dns_attacks` - (Optional) Attacks utilizing the dns protocol operations. Field introduced in 18.2.1.
* `dns_policy_index` - (Optional) Index of the dns policy to use for the mitigation rules applied to the dns attacks. Field introduced in 18.2.1.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `network_security_policy_index` - (Optional) Index of the network security policy to use for the mitigation rules applied to the attacks. Field introduced in 18.2.1.
* `oper_mode` - (Optional) Mode of dealing with the attacks - perform detection only, or detect and mitigate the attacks. Enum options - DETECTION, MITIGATION. Field introduced in 18.2.1.
* `tcp_attacks` - (Optional) Attacks utilizing the tcp protocol operations. Field introduced in 18.2.1.
* `tenant_ref` - (Optional) Tenancy of the security policy. It is a reference to an object of type tenant. Field introduced in 18.2.1.
* `udp_attacks` - (Optional) Attacks utilizing the udp protocol operations. Field introduced in 18.2.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  The uuid of the security policy. Field introduced in 18.2.1.

