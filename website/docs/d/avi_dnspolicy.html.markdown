<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_dnspolicy"
sidebar_current: "docs-avi-datasource-dnspolicy"
description: |-
  Get information of Avi DnsPolicy.
---

# avi_dnspolicy

This data source is used to to get avi_dnspolicy objects.

## Example Usage

```hcl
data "avi_dnspolicy" "foo_dnspolicy" {
    uuid = "dnspolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search DnsPolicy by name.
* `uuid` - (Optional) Search DnsPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `created_by` - Creator name. Field introduced in 17.1.1.
* `description` - Field introduced in 17.1.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the dns policy. Field introduced in 17.1.1.
* `rule` - Dns rules. Field introduced in 17.1.1.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1.
* `uuid` - Uuid of the dns policy. Field introduced in 17.1.1.

