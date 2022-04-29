<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_inventoryfaultconfig"
sidebar_current: "docs-avi-resource-inventoryfaultconfig"
description: |-
  Creates and manages Avi InventoryFaultConfig.
---

# avi_inventoryfaultconfig

The InventoryFaultConfig resource allows the creation and management of Avi InventoryFaultConfig

## Example Usage

```hcl
resource "avi_inventoryfaultconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `controller_faults` - (Optional) Configure controller faults. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Optional) Name. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `serviceengine_faults` - (Optional) Configure serviceengine faults. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant. It is a reference to an object of type tenant. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `virtualservice_faults` - (Optional) Configure virtualservice faults. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid auto generated. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.

