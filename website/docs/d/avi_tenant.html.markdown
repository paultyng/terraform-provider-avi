############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_tenant"
sidebar_current: "docs-avi-datasource-tenant"
description: |-
  Get information of Avi Tenant.
---

# avi_tenant

This data source is used to to get avi_tenant objects.

## Example Usage

```hcl
data "avi_tenant" "foo_tenant" {
    uuid = "tenant-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Tenant by name.
* `uuid` - (Optional) Search Tenant by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `config_settings` - Dict settings for tenant.
* `created_by` - Creator of this tenant.
* `description` - User defined description for the object.
* `local` - Boolean flag to set local.
* `name` - Name of the object.
* `suggested_object_labels` - Suggestive pool of key value pairs for recommending assignment of labels to objects in the user interface. Every entry is unique in both key and value. Field introduced in 20.1.2. Maximum of 256 items allowed.
* `uuid` - Unique object identifier of the object.

