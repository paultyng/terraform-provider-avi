############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_botconfigconsolidator"
sidebar_current: "docs-avi-datasource-botconfigconsolidator"
description: |-
  Get information of Avi BotConfigConsolidator.
---

# avi_botconfigconsolidator

This data source is used to to get avi_botconfigconsolidator objects.

## Example Usage

```hcl
data "avi_botconfigconsolidator" "foo_botconfigconsolidator" {
    uuid = "botconfigconsolidator-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BotConfigConsolidator by name.
* `uuid` - (Optional) Search BotConfigConsolidator by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Human-readable description of this consolidator. Field introduced in 21.1.1.
* `name` - The name of this consolidator. Field introduced in 21.1.1.
* `script` - Script that consolidates results from all components. Field introduced in 21.1.1.
* `tenant_ref` - The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.
* `uuid` - A unique identifier to this consolidator. Field introduced in 21.1.1.

