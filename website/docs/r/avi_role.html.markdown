############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_role"
sidebar_current: "docs-avi-resource-role"
description: |-
  Creates and manages Avi Role.
---

# avi_role

The Role resource allows the creation and management of Avi Role

## Example Usage

```hcl
resource "avi_role" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `filters` - (Optional) Filters for granular object access control based on object labels. Multiple filters are merged using the and operator. If empty, all objects according to the privileges will be accessible to the user. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `privileges` - (Optional) List of list.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

