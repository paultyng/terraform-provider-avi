############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_stringgroup"
sidebar_current: "docs-avi-resource-stringgroup"
description: |-
  Creates and manages Avi StringGroup.
---

# avi_stringgroup

The StringGroup resource allows the creation and management of Avi StringGroup

## Example Usage

```hcl
resource "avi_stringgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the string group.
* `type` - (Required) Type of stringgroup. Enum options - SG_TYPE_STRING, SG_TYPE_KEYVAL.
* `description` - (Optional) User defined description for the object.
* `kv` - (Optional) Configure key value in the string group.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `longest_match` - (Optional) Enable the longest match, default is the shortest match. Field introduced in 18.2.8.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the string group.

