############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_networkservice"
sidebar_current: "docs-avi-resource-networkservice"
description: |-
  Creates and manages Avi NetworkService.
---

# avi_networkservice

The NetworkService resource allows the creation and management of Avi NetworkService

## Example Usage

```hcl
resource "avi_networkservice" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the networkservice. Field introduced in 18.2.5.
* `se_group_ref` - (Required) Service engine group to which the service is applied. It is a reference to an object of type serviceenginegroup. Field introduced in 18.2.5.
* `service_type` - (Required) Indicates the type of networkservice. Enum options - ROUTING_SERVICE. Field introduced in 18.2.5.
* `vrf_ref` - (Required) Vrf context to which the service is scoped. It is a reference to an object of type vrfcontext. Field introduced in 18.2.5.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Field introduced in 18.2.5.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `routing_service` - (Optional) Routing information of the networkservice. Field introduced in 18.2.5.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 18.2.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the networkservice. Field introduced in 18.2.5.

