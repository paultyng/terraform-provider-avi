############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_errorpageprofile"
sidebar_current: "docs-avi-resource-errorpageprofile"
description: |-
  Creates and manages Avi ErrorPageProfile.
---

# avi_errorpageprofile

The ErrorPageProfile resource allows the creation and management of Avi ErrorPageProfile

## Example Usage

```hcl
resource "avi_errorpageprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Field introduced in 17.2.4.
* `error_pages` - (Optional) Defined error pages for http status codes. Field introduced in 17.2.4.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.4.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.4.

