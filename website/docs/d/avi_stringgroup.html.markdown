<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_stringgroup"
sidebar_current: "docs-avi-datasource-stringgroup"
description: |-
  Get information of Avi StringGroup.
---

# avi_stringgroup

This data source is used to to get avi_stringgroup objects.

## Example Usage

```hcl
data "avi_stringgroup" "foo_stringgroup" {
    uuid = "stringgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search StringGroup by name.
* `uuid` - (Optional) Search StringGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `kv` - Configure key value in the string group. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `longest_match` - Enable the longest match, default is the shortest match. Field introduced in 18.2.8. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name of the string group. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - Type of stringgroup. Enum options - SG_TYPE_STRING, SG_TYPE_KEYVAL. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the string group. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

