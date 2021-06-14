<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_networksecuritypolicy"
sidebar_current: "docs-avi-datasource-networksecuritypolicy"
description: |-
  Get information of Avi NetworkSecurityPolicy.
---

# avi_networksecuritypolicy

This data source is used to to get avi_networksecuritypolicy objects.

## Example Usage

```hcl
data "avi_networksecuritypolicy" "foo_networksecuritypolicy" {
    uuid = "networksecuritypolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NetworkSecurityPolicy by name.
* `uuid` - (Optional) Search NetworkSecurityPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of cloud configuration for network sec policy. Internally set by cloud connector.
* `created_by` - Creator name.
* `description` - User defined description for the object.
* `ip_reputation_db_ref` - Ip reputation database. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the object.
* `rules` - List of list.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

