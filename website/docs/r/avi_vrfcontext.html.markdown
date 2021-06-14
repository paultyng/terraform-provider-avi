<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_vrfcontext"
sidebar_current: "docs-avi-resource-vrfcontext"
description: |-
  Creates and manages Avi VrfContext.
---

# avi_vrfcontext

The VrfContext resource allows the creation and management of Avi VrfContext

## Example Usage

```hcl
resource "avi_vrfcontext" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `attrs` - (Optional) Key/value vrfcontext attributes. Field introduced in 20.1.2. Allowed in basic edition, essentials edition, enterprise edition.
* `bfd_profile` - (Optional) Bfd configuration profile. Field introduced in 20.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `bgp_profile` - (Optional) Bgp local and peer info.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `debugvrfcontext` - (Optional) Configure debug flags for vrf. Field introduced in 17.1.1.
* `description` - (Optional) User defined description for the object.
* `gateway_mon` - (Optional) Configure ping based heartbeat check for gateway in service engines of vrf. Maximum of 32 items allowed.
* `internal_gateway_monitor` - (Optional) Configure ping based heartbeat check for all default gateways in service engines of vrf. Field introduced in 17.1.1.
* `lldp_enable` - (Optional) Enable lldp. Field introduced in 18.2.10, 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `static_routes` - (Optional) List of list.
* `system_default` - (Optional) Boolean flag to set system_default.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

