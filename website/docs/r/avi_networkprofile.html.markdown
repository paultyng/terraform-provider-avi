############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_networkprofile"
sidebar_current: "docs-avi-resource-networkprofile"
description: |-
  Creates and manages Avi NetworkProfile.
---

# avi_networkprofile

The NetworkProfile resource allows the creation and management of Avi NetworkProfile

## Example Usage

```hcl
resource "avi_networkprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the network profile.
* `profile` - (Required) Dict settings for networkprofile.
* `connection_mirror` - (Optional) When enabled, avi mirrors all tcp fastpath connections to standby. Applicable only in legacy ha mode. Field introduced in 18.1.3,18.2.1.
* `description` - (Optional) User defined description for the object.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the network profile.

