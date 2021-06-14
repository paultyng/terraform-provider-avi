<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_hardwaresecuritymodulegroup"
sidebar_current: "docs-avi-datasource-hardwaresecuritymodulegroup"
description: |-
  Get information of Avi HardwareSecurityModuleGroup.
---

# avi_hardwaresecuritymodulegroup

This data source is used to to get avi_hardwaresecuritymodulegroup objects.

## Example Usage

```hcl
data "avi_hardwaresecuritymodulegroup" "foo_hardwaresecuritymodulegroup" {
    uuid = "hardwaresecuritymodulegroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search HardwareSecurityModuleGroup by name.
* `uuid` - (Optional) Search HardwareSecurityModuleGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `hsm` - Hardware security module configuration.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the hsm group configuration object.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the hsm group configuration object.

