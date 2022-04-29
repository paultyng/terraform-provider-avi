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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `hsm` - Hardware security module configuration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name of the hsm group configuration object. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the hsm group configuration object. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

