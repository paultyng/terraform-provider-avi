############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_botmapping"
sidebar_current: "docs-avi-datasource-botmapping"
description: |-
  Get information of Avi BotMapping.
---

# avi_botmapping

This data source is used to to get avi_botmapping objects.

## Example Usage

```hcl
data "avi_botmapping" "foo_botmapping" {
    uuid = "botmapping-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BotMapping by name.
* `uuid` - (Optional) Search BotMapping by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `mapping_rules` - Rules for bot classification. Field introduced in 21.1.1. Minimum of 1 items required.
* `name` - The name of this mapping. Field introduced in 21.1.1.
* `tenant_ref` - The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.
* `uuid` - A unique identifier of this mapping. Field introduced in 21.1.1.

