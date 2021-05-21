<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_alertemailconfig"
sidebar_current: "docs-avi-resource-alertemailconfig"
description: |-
  Creates and manages Avi AlertEmailConfig.
---

# avi_alertemailconfig

The AlertEmailConfig resource allows the creation and management of Avi AlertEmailConfig

## Example Usage

```hcl
resource "avi_alertemailconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name of the email notification service.
* `to_emails` - (Required) Alerts are sent to the comma separated list of  email recipients.
* `cc_emails` - (Optional) Alerts are copied to the comma separated list of  email recipients.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) User defined description for the object.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

