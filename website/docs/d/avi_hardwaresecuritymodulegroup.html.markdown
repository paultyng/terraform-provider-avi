############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Name of the hsm group configuration object.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the hsm group configuration object.

