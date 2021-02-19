############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_protocolparser"
sidebar_current: "docs-avi-datasource-protocolparser"
description: |-
  Get information of Avi ProtocolParser.
---

# avi_protocolparser

This data source is used to to get avi_protocolparser objects.

## Example Usage

```hcl
data "avi_protocolparser" "foo_protocolparser" {
    uuid = "protocolparser-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ProtocolParser by name.
* `uuid` - (Optional) Search ProtocolParser by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Description of the protocol parser. Field introduced in 18.2.3.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Name of the protocol parser. Field introduced in 18.2.3.
* `parser_code` - Command script provided inline. Field introduced in 18.2.3.
* `tenant_ref` - Tenant uuid of the protocol parser. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `uuid` - Uuid of the protocol parser. Field introduced in 18.2.3.

